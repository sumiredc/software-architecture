package use_cases

import (
	"fmt"

	"github.com/Greek-Academy/software-architecture/clean-architecture/application/use_cases/requests"
	"github.com/Greek-Academy/software-architecture/clean-architecture/application/use_cases/responses"
	"github.com/Greek-Academy/software-architecture/clean-architecture/domain/entities"
)

type MakeFaceI interface {
	Handle() responses.MuhohoResponseI
}

type makeFace struct {
	req requests.MuhohoRequestI
}

func NewMakeFace(req requests.MuhohoRequestI) MakeFaceI {
	return &makeFace{req}
}

func (mf *makeFace) Handle() responses.MuhohoResponseI {
	m := entities.NewMuhoho(mf.req.GetEye(), mf.req.GetMouth())
	face := fmt.Sprintf("%s %s%s%s%s", m.LeftCheek, m.Eye, m.Mouth, m.Eye, m.RightCheek)
	return responses.NewMuhohoResponse(face)
}
