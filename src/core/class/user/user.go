package user

type User struct {
	ID   int
	Name string
	Role []*Role
}

type Role struct {
	ID   int
	Name string
	Code string
}
