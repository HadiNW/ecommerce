package user

// UserFormatter ...
type UserFormatter struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Username string `json:"username"`
}

// FormatUser ...
func FormatUser(user User) UserFormatter {
	return UserFormatter{
		ID:       user.ID,
		FullName: user.FullName,
		Username: user.Username,
	}
}
