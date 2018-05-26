package constant

// AccountType represents the type of account source
type AccountType string

const (
	Undefined AccountType = "undefined"
	Email     AccountType = "email"
	Kakao     AccountType = "kakao"
	Facebook  AccountType = "facebook"
	Gmail     AccountType = "gmail"
)
