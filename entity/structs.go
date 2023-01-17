package entity

type User struct {
	Uid   string `json:"uid"`
	Email string `json:"email"`
}

type Token struct {
	IDToken      string `json:"idToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshToken struct {
	IDToken      string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
	UID          string `json:"user_id"`
}
