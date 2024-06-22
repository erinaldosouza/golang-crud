package request

type UserRequest struct {
	Email string `json:"email,omitempty"`
	Pwd   string `json:"pwd,omitempty"`
	Name  string `json:"name,omitempty"`
	Age   int8   `json:"age,omitempty"`
}
