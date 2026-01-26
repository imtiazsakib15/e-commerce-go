package repo

type User struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"isShopOwner"`
}

type UserRepo interface {
	Create(user User) (*User, error)
	Find(email string, password string) (*User, error)
}

type userRepo struct {
	users []User
}

func NewUserRepo() UserRepo {
	return &userRepo{}
}

func (r *userRepo) Create(user User) (*User, error) {
	user.ID = len(r.users) + 1
	r.users = append(r.users, user)
	return &user, nil
}

func (r *userRepo) Find(email string, password string) (*User, error) {
	for idx, user := range r.users {
		if user.Email == email && user.Password == password {
			return &r.users[idx], nil
		}
	}
	return nil, nil
}