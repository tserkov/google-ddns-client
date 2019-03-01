package record

import (
	"errors"
)

type Record struct {
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
	Hostname string `json:"hostname" yaml:"hostname"`
}

func (r Record) Valid() error {
	if r.Username == "" {
		return errors.New("username not provided")
	}

	if r.Password == "" {
		return errors.New("password not provided")
	}

	if r.Hostname == "" {
		return errors.New("hostname not provided")
	}

	return nil
}
