package models

import (
	"encoding/json"
	"strings"
)

type Password string

type User struct {
	Username string `json:"username"`
	Password Password `json:"password"`
}

func (p Password) MarshalJSON() ([]byte, error) {
	s := string(p)
	masked := s[:1] + strings.Repeat("*", len(s)-2) + s[len(s)-1:]
	return json.Marshal(masked)
}