package data

type User struct {
	Username string
	Email    string
	UserRole UserRole
}

type UserRole string

const (
	Citizen UserRole = "Citizen"
)

func (u User) Equals(user User) bool {
	return u.Username == user.Username
}
