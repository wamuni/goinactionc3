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

type email_user struct {
	name  string
	email string
}

func (eu email_user) notify() {
	fmt.Printf("Sending email to %s<%s>\n", eu.name, eu.email)
}

func (eu *email_user) changeEmail(email string) {
	eu.email = email
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
	utils.PrintSeprater()
	bill_email_user := email_user{"bill", "bill@email.com"}
	bill_email_user.notify()
	lisa_email_user := email_user{"lisa", "lisa@email.com"}
	lisa_email_user.notify()
	bill_email_user.changeEmail("bill.changed@email.com")
	lisa_email_user.notify()
	bill_email_user.notify()
}

func printUser(u User) {
	fmt.Printf("Name:\t\t%s,\nEmail:\t\t%s,\nExt:\t\t%d,\nPrivileged:\t%t,\n", u.name, u.email, u.ext, u.privileged)
}

func printAdmin(ad admin) {
	printUser(ad.person)
	fmt.Printf("Level:\t\t%s\n", ad.level)
}
