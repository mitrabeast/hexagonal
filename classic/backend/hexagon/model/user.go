package model

import "encoding/json"

type User struct {
	ID        int64  `json:"userId"`
	Username  string `json:"username" validate:"required,gte=5,lte=12"`
	FirstName string `json:"firstName" validate:"required,gte=2,lte=20"`
	LastName  string `json:"lastName" validate:"gte=5,lte=20"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,gte=8,lte=50"`
}

func (u *User) PublicInfo() json.RawMessage {
	data, _ := json.Marshal(struct {
		ID        int64  `json:"userId"`
		Username  string `json:"username"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName,omitempty"`
		Email     string `json:"email"`
	}{
		u.ID, u.Username, u.FirstName, u.LastName, u.Email,
	})
	return data
}
