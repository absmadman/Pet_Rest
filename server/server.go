package server

import "errors"

// NEED TO USE MANNERS PACKAGE FOR SHUT DOWN SERVER
// http.Request.Method
// net/url
// "github.com/braintree/manners"

type AuthHandler struct {
	data *UsersData
}

type UsersData struct {
	users map[string]User
}

type User struct {
	login string
	pass  string
}

// Register func return nil if ok, else returns error
func (hnd *AuthHandler) Register(log string, pas string) error {
	if _, have := hnd.data.users[log]; have {
		return errors.New("user is already exist")
	}
	hnd.data.users[log] = User{
		login: log,
		pass:  pas,
	}
	return nil
}

func (hnd *AuthHandler) Login(log string, pas string) error {
	userData, have := hnd.data.users[log]
	if !have {
		return errors.New("user is not exist")
	}
	if userData.pass != pas {
		return errors.New("invalid password")
	}

	return nil
}
