package infrastructure

import (
	"errors"

	"{{ .ModuleName }}/{{ .AppName }}/internal/domain/user"
)

type inMemoryUserRepository struct {
	users  map[int]user.User
	nextID int
}

func NewInMemoryUserRepository() user.UserRepository {
	return &inMemoryUserRepository{
		users:  make(map[int]user.User),
		nextID: 1,
	}
}

func (r *inMemoryUserRepository) FindBy(id int) (*user.User, error) {
	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (r *inMemoryUserRepository) Save(user user.User) error {
	user.ID = r.nextID
	r.nextID++
	r.users[user.ID] = user
	return nil
}
