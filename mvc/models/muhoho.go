package models

import (
	"fmt"
	"net/http"
)

type Muhoho struct {
	leftCheek  string
	rightCheek string
	Eye        string
	Mouth      string
}

func NewMuhoho(writer *http.ResponseWriter, eye string, mouth string) *Muhoho {
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
		leftCheek:  LEFT_CHEEK,
		rightCheek: RIGHT_CHEEK,
		Eye:        eye,
		Mouth:      mouth,
	}
}

func (m *Muhoho) Face() string {
	return fmt.Sprintf("%s %s%s%s%s", m.leftCheek, m.Eye, m.Mouth, m.Eye, m.rightCheek)
}
