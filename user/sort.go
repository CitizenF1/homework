package user

import "time"

type User struct {
	ID        string
	CreatedAt time.Time
}

type UserList []User

func (u UserList) Len() int {
	return len(u)
}

func (u UserList) Less(i, j int) bool {
	return u[i].CreatedAt.Before(u[j].CreatedAt)
}

func (u UserList) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}
