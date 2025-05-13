package response

type UserResponse struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Age      int8   `json:"age"`
}
