package entity

type User struct {
	ID   *uint32
	Name string
	Type string
}

func NewUser(name string, types string) (*User, error) {
	t := &User{
		Name: name,
		Type: types,
	}

	err := t.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}

	return t, nil
}

func (u *User) Validate() error {
	if u.Name == "" || u.Type == "" {
		return ErrInvalidEntity
	}

	return nil
}
