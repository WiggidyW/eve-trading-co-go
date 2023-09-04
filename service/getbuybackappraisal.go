package service

import (
	"context"

	protoclient "github.com/WiggidyW/etco-go/client/proto"
	"github.com/WiggidyW/etco-go/proto"
	"github.com/WiggidyW/etco-go/protoutil"
	"github.com/WiggidyW/etco-go/staticdb"
)

func (s *Service) GetBuybackAppraisal(
	ctx context.Context,
	req *proto.GetBuybackAppraisalRequest,
) (
	rep *proto.GetBuybackAppraisalResponse,
	err error,
) {
	rep = &proto.GetBuybackAppraisalResponse{}

	var isAdmin bool
	if req.Admin {
		var ok bool
		_, _, _, rep.Auth, rep.Error, ok =
			s.TryAuthenticate(
				ctx,
				req.Auth,
				"new-buyback-appraisal",
				true,
			)
		if !ok {
			return rep, nil
		}
		isAdmin = true
	}

	typeNamingSession := protoutil.
		MaybeNewLocalTypeNamingSession(req.IncludeTypeNaming)
	appraisalRep, err := s.getBuybackAppraisalClient.Fetch(
		ctx,
		protoclient.PBGetAppraisalParams[*staticdb.LocalIndexMap]{
			TypeNamingSession: typeNamingSession,
			AppraisalCode:     req.Code,
		},
	)
	if err != nil {
		rep.Error = NewProtoErrorRep(
			proto.ErrorCode_SERVER_ERROR,
			err.Error(),
		)
		return rep, nil
	}

	rep.Appraisal = appraisalRep.Appraisal
	rep.TypeNamingLists = protoutil.MaybeFinishTypeNamingSession(
		typeNamingSession,
	)
	if isAdmin {
		rep.CharacterId = appraisalRep.CharacterId
	} else {
		rep.HashCharacterId = protoutil.NewPBObfuscateCharacterID(
			appraisalRep.CharacterId,
		)
	}

	return rep, nil
}