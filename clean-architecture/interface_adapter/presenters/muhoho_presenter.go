package presenters

import (
	"github.com/Greek-Academy/software-architecture/clean-architecture/application/use_cases/responses"
)

type MuhohoPresenterI interface{}

type muhohoPresenter struct {
	Face string
}

func NewMuhohoPresenter(res responses.MuhohoResponseI) MuhohoPresenterI {
	return &muhohoPresenter{Face: res.GetFace()}
}
