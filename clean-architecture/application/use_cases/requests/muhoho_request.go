package requests

import "net/http"

type MuhohoRequestI interface {
	GetEye() string
	GetMouth() string
}

type muhohoRequest struct {
	Eye   string
	Mouth string
}

func NewMuhohoRequest(r *http.Request) MuhohoRequestI {
	return &muhohoRequest{
		Eye:   r.URL.Query().Get("eye"),
		Mouth: r.URL.Query().Get("mouth"),
	}
}

func (req *muhohoRequest) GetEye() string {
	return req.Eye
}

func (req *muhohoRequest) GetMouth() string {
	return req.Mouth
}
