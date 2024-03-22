package responses

type MuhohoResponseI interface {
	GetFace() string
}

type muhohoResponse struct {
	face string
}

func NewMuhohoResponse(face string) MuhohoResponseI {
	return &muhohoResponse{face}
}

func (res *muhohoResponse) GetFace() string {
	return res.face
}
