package auth

import "encoding/json"

type AuthUserDto struct {
	PassHash string `json:"passHash"`
	Code     int    `json:"code"`
}

func (u AuthUserDto) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

func (u AuthUserDto) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &u)
}
