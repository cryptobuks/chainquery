package processing

import (
	"encoding/hex"
	"encoding/json"

	"github.com/lbryio/chainquery/lbrycrd"
	"github.com/lbryio/chainquery/model"
	"github.com/lbryio/lbry.go/errors"
	"github.com/lbryio/lbry.go/util"
	"github.com/lbryio/lbryschema.go/pb"

	"github.com/lbryio/chainquery/datastore"
	"github.com/sirupsen/logrus"
)

func processAsClaim(script []byte, vout model.Output, tx model.Transaction) (address *string, err error) {
	var pubkeyscript []byte
	var name string
	var claimid string
	if lbrycrd.IsClaimNameScript(script) {
		name, claimid, pubkeyscript, err = processClaimNameScript(&script, vout, tx)
		if err != nil {
			return nil, err
		}
		return nil, nil
	} else if lbrycrd.IsClaimSupportScript(script) {
		name, claimid, pubkeyscript, err = processClaimSupportScript(&script, vout, tx)
		if err != nil {
			return nil, err
		}
		return nil, nil
	} else if lbrycrd.IsClaimUpdateScript(script) {
		name, claimid, pubkeyscript, err = processClaimUpdateScript(&script, vout, tx)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	pksAddress := lbrycrd.GetAddressFromPublicKeyScript(pubkeyscript)
	address = &pksAddress
	logrus.Debug("Handled Claim: ", " Name ", name, ", ClaimID ", claimid)

	return nil, errors.Base("Not a claim -- " + hex.EncodeToString(script))
}

func processClaimNameScript(script *[]byte, vout model.Output, tx model.Transaction) (name string, claimid string, pkscript []byte, err error) {
	claimid, err = util.ClaimIDFromOutpoint(vout.TransactionHash, int(vout.Vout))
	if err != nil {
		return name, "", pkscript, err
	}
	name, value, pkscript, err := lbrycrd.ParseClaimNameScript(*script)
	if err != nil {
		errors.Prefix("Claim name script parsing error: ", err)
		return name, claimid, pkscript, err
	}
	pbClaim, err := lbrycrd.DecodeClaimValue(name, value)
	if err != nil {
		logrus.Warning("saving non-conforming claim - Name: ", name, " ClaimId: ", claimid)
		saveUnknownClaim(name, claimid, false, value, vout, tx)
		return name, claimid, pkscript, nil
	}
	if pbClaim != nil && err == nil {
		claim := datastore.GetClaim(claimid)
		claim, err := processClaim(pbClaim, claim, value, vout, tx)
		if err != nil {
			return name, claimid, pkscript, err
		}
		claim.ClaimID = claimid
		claim.Name = name
		claim.TransactionTime = tx.TransactionTime
		datastore.PutClaim(claim)
	}

	return name, claimid, pkscript, err
}

func processClaimSupportScript(script *[]byte, vout model.Output, tx model.Transaction) (name string, claimid string, pubkeyscript []byte, err error) {
	name, claimid, pubkeyscript, err = lbrycrd.ParseClaimSupportScript(*script)
	if err != nil {
		errors.Prefix("Claim support processing error: ", err)
		return name, claimid, pubkeyscript, err
	}
	support := datastore.GetSupport(tx.Hash, vout.Vout)
	support = processSupport(claimid, support, vout, tx)
	datastore.PutSupport(support)

	return name, claimid, pubkeyscript, err
}

func processClaimUpdateScript(script *[]byte, vout model.Output, tx model.Transaction) (name string, claimId string, pubkeyscript []byte, err error) {
	name, claimId, value, pubkeyscript, err := lbrycrd.ParseClaimUpdateScript(*script)
	if err != nil {
		errors.Prefix("Claim update processing error: ", err)
		return name, claimId, pubkeyscript, err
	}
	pbClaim, err := lbrycrd.DecodeClaimValue(name, value)
	if err != nil {
		logrus.Warning("saving non-conforming claim - Update: ", name, " ClaimId: ", claimId)
		saveUnknownClaim(name, claimId, true, value, vout, tx)
		return name, claimId, pubkeyscript, nil
	}
	if pbClaim != nil && err == nil {
		claim := datastore.GetClaim(claimId)
		claim, err := processUpdateClaim(pbClaim, claim, value)
		if err != nil {
			return name, claimId, pubkeyscript, err
		}
		if claim == nil {
			logrus.Warning("ClaimUpdate for non-existent claim! ", claimId)
			return name, claimId, pubkeyscript, err
		}
		datastore.PutClaim(claim)
	}
	return name, claimId, pubkeyscript, err
}

func processClaim(pbClaim *pb.Claim, claim *model.Claim, value []byte, output model.Output, tx model.Transaction) (*model.Claim, error) {
	if claim == nil {
		claim = &model.Claim{}
	}
	claim.SetTransactionByHashG(false, &tx)
	claim.Vout = output.Vout
	claim.Version = pbClaim.GetVersion().String()
	claim.ValueAsHex = hex.EncodeToString(value)

	var js map[string]interface{} //JSON Map
	if json.Unmarshal(value, &js) == nil {
		claim.ValueAsJSON.String = string(value)
		claim.ValueAsJSON.Valid = true
	}

	setSourceInfo(claim, pbClaim)
	setMetaDataInfo(claim, pbClaim)
	setCertificateInfo(claim, pbClaim)

	return claim, nil
}

func processSupport(claimID string, support *model.Support, output model.Output, tx model.Transaction) *model.Support {
	if support == nil {
		support = &model.Support{}
	}

	support.TransactionHash.String = tx.Hash
	support.TransactionHash.Valid = true
	support.Vout = output.Vout
	support.SupportAmount = output.Value.Float64
	support.SupportedClaimID = claimID

	return support

}

func processUpdateClaim(pbClaim *pb.Claim, claim *model.Claim, value []byte) (*model.Claim, error) {
	if claim == nil {
		return nil, nil
	}
	claim.Version = pbClaim.GetVersion().String()
	claim.ValueAsHex = hex.EncodeToString(value)

	var js map[string]interface{} //JSON Map
	if json.Unmarshal(value, &js) == nil {
		claim.ValueAsJSON.String = string(value)
		claim.ValueAsJSON.Valid = true
	}

	setSourceInfo(claim, pbClaim)
	setMetaDataInfo(claim, pbClaim)
	setPublisherInfo(claim, pbClaim)
	setCertificateInfo(claim, pbClaim)

	return claim, nil
}

func setPublisherInfo(claim *model.Claim, pbClaim *pb.Claim) {
	if pbClaim.GetPublisherSignature() != nil {
		publisherClaimId := hex.EncodeToString(pbClaim.GetPublisherSignature().GetCertificateId())
		claim.PublisherID.String = publisherClaimId
		claim.PublisherID.Valid = true
		claim.PublisherSig.String = hex.EncodeToString(pbClaim.GetPublisherSignature().GetSignature())
		claim.PublisherSig.Valid = true
	}
}

func setCertificateInfo(claim *model.Claim, pbClaim *pb.Claim) {

	if pbClaim.GetClaimType() == pb.Claim_certificateType {
		certificate := pbClaim.GetCertificate()
		certBytes, err := json.Marshal(certificate)
		if err != nil {
			logrus.Error("Could not form json from certificate")
		}
		claim.Certificate.String = string(certBytes)
		claim.Certificate.Valid = true
	}
}

func setMetaDataInfo(claim *model.Claim, pbClaim *pb.Claim) {
	stream := pbClaim.GetStream()
	if stream != nil {
		metadata := stream.GetMetadata()
		if metadata != nil {
			claim.Title.String = metadata.GetTitle()
			claim.Title.Valid = true //

			claim.Description.String = metadata.GetDescription()
			claim.Description.Valid = true

			claim.Language.String = metadata.GetLanguage().String()
			claim.Language.Valid = true

			claim.Author.String = metadata.GetAuthor()
			claim.Author.Valid = true

			claim.ThumbnailURL.String = metadata.GetThumbnail()
			claim.ThumbnailURL.Valid = true

			fee := metadata.GetFee()
			if fee != nil {
				claim.FeeCurrency.String = fee.GetCurrency().String()
				claim.FeeCurrency.Valid = true

				claim.Fee = float64(fee.GetAmount())
			}
		}
	}
}

func setSourceInfo(claim *model.Claim, pbClaim *pb.Claim) {
	stream := pbClaim.GetStream()
	if stream != nil {
		source := stream.GetSource()
		if source != nil {
			contentType := source.GetContentType()
			if contentType != "" {
				claim.ContentType.String = contentType
				claim.ContentType.Valid = true
				if source.GetSourceType() == pb.Source_lbry_sd_hash {
					claim.SDHash.String = hex.EncodeToString(source.GetSource())
					claim.SDHash.Valid = true
				}
			}
		}
	}
}

func saveUnknownClaim(name string, claimid string, isUpdate bool, value []byte, vout model.Output, tx model.Transaction) {
	unknownClaim := model.UnknownClaim{}
	unknownClaim.Vout = vout.Vout
	unknownClaim.Name = name
	unknownClaim.ClaimID = claimid
	unknownClaim.IsUpdate = isUpdate
	unknownClaim.TransactionHash.String = vout.TransactionHash
	unknownClaim.TransactionHash.Valid = true
	unknownClaim.ValueAsHex = hex.EncodeToString(value)
	unknownClaim.BlockHash = tx.BlockByHashID

	var js map[string]interface{} //JSON Map
	if json.Unmarshal(value, &js) == nil {
		unknownClaim.ValueAsJSON.String = string(value)
		unknownClaim.ValueAsJSON.Valid = true
	}

	unknownClaim.SetOutputG(false, &vout)
	unknownClaim.InsertG()

}
