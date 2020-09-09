package handlers

import "net/http"

type Users struct {
	username *string
}

func NewUsers(u *string) *Users {
	return &Users(u)
}

func (u *Users) ServeHTTP(rw http.ResponseWriter, *http.Request) {
	fmt.Fprintf(rw, "Hello user %s", u)
}