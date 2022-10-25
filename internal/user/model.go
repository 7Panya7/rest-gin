package user

type Person struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Age       int8   `json:"age" binding:"gte=1,lte=130"`
	Email     string `json:"email" binding:"required,email"`
}
