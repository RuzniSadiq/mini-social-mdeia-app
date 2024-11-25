package auth

// UserCredentials represents the user's login credentials
type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}