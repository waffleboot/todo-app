package todo

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
