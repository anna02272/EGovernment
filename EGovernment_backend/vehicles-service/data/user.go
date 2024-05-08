package data

type User struct {
	Username string
	Email    string
	UserRole UserRole
}

type UserRole string

const (
	Policeman UserRole = "Policeman"
)

func (u User) Equals(user User) bool {
	return u.Username == user.Username
}
