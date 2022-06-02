package models

type User struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}

type UserVisible struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u *User) UserToUserVisible() UserVisible {
	return UserVisible{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}
