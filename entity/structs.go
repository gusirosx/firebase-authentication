// Package entity contains data structures representing different entities in the system
package entity

// User represents a user in the system
type User struct {
	Uid   string `json:"uid"`   // unique identifier for the user
	Email string `json:"email"` // email of the user
}

// Token represents an access and refresh token pair
type Token struct {
	IDToken      string `json:"idToken"`      // access token
	RefreshToken string `json:"refreshToken"` // refresh token
}

// RefreshToken represents a set of tokens and a user identifier
type RefreshToken struct {
	IDToken      string `json:"id_token"`      // access token
	RefreshToken string `json:"refresh_token"` // refresh token
	AccessToken  string `json:"access_token"`  // JWT access token
	UID          string `json:"user_id"`       // user identifier
}
