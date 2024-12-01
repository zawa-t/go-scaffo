package inmemory

import (
	"errors"

	"{{ .BaseImportPath }}/internal/domain/user"
)

type UserRepository struct {
	users  map[int]user.User
	nextID int
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users:  make(map[int]user.User),
		nextID: 1,
	}
}

func (r *UserRepository) FindBy(id int) (*user.User, error) {
	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (r *UserRepository) Save(user user.User) error {
	user.ID = r.nextID
	r.nextID++
	r.users[user.ID] = user
	return nil
}
