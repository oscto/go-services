package metadata

type UserCreateBody struct {
	Email    string `json:"email"`
	Salt     string `json:"salt"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}
