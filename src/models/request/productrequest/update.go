package productrequest

type Update struct {
	Id   int    `validate:"required" json:"id"`
	Name string `validate:"required,min=4,max=200" json:"name"` // tidak butuh (id Int) karena set auto-increment
}