package auth

// UserSession encapsulates the user login session
type UserSession struct {
	Token string `json:"token"`
}
