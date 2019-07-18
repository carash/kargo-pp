package user

type User struct {
}

func (b User) Less(o User) bool {
	return true
}

type Role struct {
}

func (v Role) Less(o Role) bool {
	return true
}
