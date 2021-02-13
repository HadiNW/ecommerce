package user

// UserFormatter ...
type UserFormatter struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Token    string `json:"token,omitempty"`
}

// FormatUser ...
func FormatUser(user User, token string) UserFormatter {
	return UserFormatter{
		ID:       user.ID,
		FullName: user.FullName,
		Username: user.Username,
		Token:    token,
	}
}
