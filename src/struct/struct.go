package main

import (
	"fmt"

	"chapter3.wamuni/m/src/utils"
)

type User struct {
	name       string
	email      string
	ext        int
	privileged bool
}

func main() {
	fmt.Printf("Hello MotherFucker\n")
	utils.PrintSeprater()
	var bill User // initialize with nothing, then it will only get zero-value
	printUser(bill)
	lisa := User{
		name:       "Lisa",
		email:      "lisa@email.com",
		ext:        123,
		privileged: true,
	}
	utils.PrintSeprater()
	printUser(lisa)
}

func printUser(u User) {
	fmt.Printf("Name:\t\t%s,\nEmail:\t\t%s,\nExt:\t\t%d,\nPrivileged:\t%t,\n", u.name, u.email, u.ext, u.privileged)
}
