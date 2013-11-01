package models

type User struct {
    Name string
    Password string
}

func (u *User) Verify() error {

    return nil
}
