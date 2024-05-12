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

// declare a struct within a struct
type admin struct {
	person User
	level  string
}

// declare a new type for existing type
type Duration int64

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

	eddie := User{"Eddie", "eddie@email.com", 124, true}
	utils.PrintSeprater()
	printUser(eddie)

	// how to declare embeded struct
	fred := admin{
		person: User{
			name:       "Lisa",
			email:      "lisa@email.com",
			ext:        123,
			privileged: true,
		},
		level: "super",
	}
	utils.PrintSeprater()
	printAdmin(fred)
}

func printUser(u User) {
	fmt.Printf("Name:\t\t%s,\nEmail:\t\t%s,\nExt:\t\t%d,\nPrivileged:\t%t,\n", u.name, u.email, u.ext, u.privileged)
}

func printAdmin(ad admin) {
	printUser(ad.person)
	fmt.Printf("Level:\t\t%s\n", ad.level)
}
