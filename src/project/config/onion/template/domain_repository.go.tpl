package user

type UserRepository interface {
	Save(user User) error
	FindBy(id int) (*User, error)
}
