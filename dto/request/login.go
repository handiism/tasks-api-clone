package request

type Login struct {
	Email    string `binding:"required,email" json:"email" faker:"email"`
	Password string `binding:"required,alphanum,min=8" json:"password" faker:"password"`
}
