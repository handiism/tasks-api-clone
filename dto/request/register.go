package request

type Register struct {
	Name     string `binding:"required" json:"name" faker:"name"`
	Email    string `binding:"required,email" json:"email" faker:"email"`
	Password string `binding:"required,alphanum,min=8" json:"password" faker:"password"`
}
