package request

type UserRequest struct {
	Email string `json:"email" binding:"required,email"`
	Pwd   string `json:"pwd"   binding:"required,min=6,max=6,containsany=@#$*"`
	Name  string `json:"name"  binding:"required,min=4,max=50"`
	Age   int8   `json:"age"   binding:"required,min=18,max=140"`
}
