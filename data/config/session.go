package config

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterVendor struct {
}

type SessionToken struct {
	AccessToken  string `json:"access"`
	RefreshToken string `json:"refresh"`
}

type User struct {
	Credentials
	SessionToken
}
