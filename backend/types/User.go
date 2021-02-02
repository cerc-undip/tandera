package types

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"-" form:"password"`
}
