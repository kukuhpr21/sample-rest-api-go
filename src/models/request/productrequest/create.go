package productrequest

type Create struct {
	Name string `validate:"required,min=4,max=200" json:"name"`
}
