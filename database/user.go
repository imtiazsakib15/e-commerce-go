package database

type User struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"isShopOwner"`
}

var users []User

func (u *User) Store(user User) User {
	user.ID = len(users) + 1
	users = append(users, user)
	return user
}

func Find(email string, password string) *User {
	for _, user := range users {
		if user.Email == email && user.Password == password {
			return &user
		}
	}
	return nil
}