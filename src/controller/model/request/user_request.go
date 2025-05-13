package request

type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,containsany=!@#$%^&*"`
	Username string `json:"username" binding:"required,min=4,max=50"`
	Age      int8   `json:"age" binding:"min=18,max=100"` // Assuming age is between 18 and 100
}
