package entity

type AuthResp struct {
	Token     string `json:"token"`
	ExpiresIn int    `json:"expires_in"`
}
