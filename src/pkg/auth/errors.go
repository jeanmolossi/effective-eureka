package auth

type InvalidCredentialsErr struct {
	Message string `json:"error" example:"Invalid credentials"`
}
