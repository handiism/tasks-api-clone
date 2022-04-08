package request

type Login struct {
	Email    string `binding:"required,email"          json:"email"`
	Password string `binding:"required,alphanum,min=8" json:"password"`
}
