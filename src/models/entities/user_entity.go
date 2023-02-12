package entities

type UserEntity struct {
	Id           string `json:"id"` 
	IdDetailUser string `json:"id_detail_user"`
	Name 		 string `json:"name"`
	Email 		 string `json:"email"`
	Password     string `json:"password"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}