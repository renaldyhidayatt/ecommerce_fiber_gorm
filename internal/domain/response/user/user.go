package user

type UserResponse struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	IsStaff    bool   `json:"is_staff"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}
