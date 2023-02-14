package authrequest

type Token struct {
	Id    string `validate:"required" json:"id"`
	Name  string `validate:"required" json:"name"`
	Email string `validate:"required" json:"email"`
}
