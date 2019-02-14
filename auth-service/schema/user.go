package schema

//User Data strutcure
type User struct {
	ID       string `json:"id" storm:"index"`
	Name     string `json:"name" storm:"index"`
	Email    string `json:"email" storm:"unique"`
	Password string `json:"password"`
	Role     string `json:"role" storm:"index"`
}
