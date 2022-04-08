package request

type SaveUser struct {
	Name     string `binding:"required"                json:"name"`
	Email    string `binding:"required,email"          json:"email"`
	Password string `binding:"required,alphanum,min=8" json:"password"`
}
