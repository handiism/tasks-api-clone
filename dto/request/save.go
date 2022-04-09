package request

type SaveUser struct {
	Name     string `binding:"required"                json:"name"`
	Email    string `binding:"required,email"          json:"email"`
	Password string `binding:"required,alphanum,min=8" json:"password"`
}

type SaveList struct {
	Title string `binding:"required,min=1" json:"title"`
}
