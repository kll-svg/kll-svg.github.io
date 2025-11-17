package main

import "time"

type BaseUser struct {
	ID        int
	UserName  string
	Email     string
	CreatedAt time.Time
}

func (u *BaseUser) GetCreatedAt() string {
	return u.CreatedAt.Format("2006-01-02 15:04:05")
}

func main() {

}
