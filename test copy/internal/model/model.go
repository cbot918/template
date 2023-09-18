package model

type SignupRequest struct {
	Email    string `json:"email"`
	Name     string `json:"Name"`
	Password string `json:"password"`
}

type SignupResponse struct {
	Email string `json:"email"`
	Name  string `json:"Name"`
}
