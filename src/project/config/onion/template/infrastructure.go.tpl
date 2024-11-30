package infrastructure

import (
	"errors"

	"{{ .BaseImportPath }}/internal/domain/user"
)

type InMemoryUserRepository struct {
	users  map[int]user.User
	nextID int
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users:  make(map[int]user.User),
		nextID: 1,
	}
}

func (r *InMemoryUserRepository) FindBy(id int) (*user.User, error) {
	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (r *InMemoryUserRepository) Save(user user.User) error {
	user.ID = r.nextID
	r.nextID++
	r.users[user.ID] = user
	return nil
}
