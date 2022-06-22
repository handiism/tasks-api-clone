package request

type EmailChecking struct {
	Email string `binding:"required,email" json:"email"`
}
