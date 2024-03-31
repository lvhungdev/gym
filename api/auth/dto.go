package auth

type signInReqDto struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type signInResDto struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type signUpReqDto struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type signUpResDto struct {
	Email string `json:"email"`
	Token string `json:"token"`
}
