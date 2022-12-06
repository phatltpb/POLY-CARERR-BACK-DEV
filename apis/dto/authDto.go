package dto

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"Password"`
}

type Forgotpassword struct {
	Email string `json:"email"`
}

type ResetPassword struct {
	Password string `json:"password"`
}

type ChangePassword struct {
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}
