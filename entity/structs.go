package entity

type User struct {
	Uid string `json:"uid"`
}

type Token struct {
	IDToken      string `json:"idToken"`
	RefreshToken string `json:"refreshToken"`
}
