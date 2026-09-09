package main

type User struct {
	Name string
	Age  int
}

func isValid(u User) bool {
	if u.Name == "" {
		return false
	}
	if u.Age < 0 {
		return false
	}
	return true
}
