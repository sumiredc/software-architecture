package entities

type Muhoho struct {
	LeftCheek  string
	RightCheek string
	Eye        string
	Mouth      string
}

func NewMuhoho(eye string, mouth string) *Muhoho {
	const LEFT_CHEEK = "("
	const RIGHT_CHEEK = ")"
	const DEFAULT_EYE = "^"
	const DEFAULT_MOUTH = "ω"

	if eye == "" {
		eye = DEFAULT_EYE
	}

	if mouth == "" {
		mouth = DEFAULT_MOUTH
	}

	return &Muhoho{
		LeftCheek:  LEFT_CHEEK,
		RightCheek: RIGHT_CHEEK,
		Eye:        eye,
		Mouth:      mouth,
	}
}
