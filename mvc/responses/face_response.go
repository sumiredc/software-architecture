package responses

type FaceResponse struct {
	Face string
}

func NewFaceResponse(face string) *FaceResponse {
	return &FaceResponse{
		Face: face,
	}
}
