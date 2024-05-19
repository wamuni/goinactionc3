package main

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"chapter3.wamuni/m/src/utils"
)

// init function is executed before main function
func init() {
	// if the number of the arguments is not 2, finish the code
	if len(os.Args) != 2 {
		fmt.Printf("Usage: ./interface <url>\n")
		os.Exit(-1)
	}
}

// this is a simple curl program
func main() {
	bufferWithioCopy()
	utils.PrintSeprater()
	interfaceImplementation()
	// get response from web service
	// res, err := http.Get(os.Args[1])
	// // if there is err returned back, then finish the program
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// io.Copy(os.Stdout, res.Body)
	// if err := res.Body.Close(); err != nil {
	// 	fmt.Println(err)
	// }
	counter := utils.New(10)
	fmt.Printf("Counter: %d\n", counter)
	utils.PrintSeprater()
	a := utils.Admin{Right: 10}
	a.Name = "Eddie"
	a.Email = "Eddie@Gmail.com"

	fmt.Printf("User: %v\n", a)
}

func bufferWithioCopy() {
	var b bytes.Buffer

	b.Write([]byte("Hello")) // bytes.Buffer has implement io.Write interface

	fmt.Fprintf(&b, " World\n")

	// Copy the data in b to os.Stdout
	io.Copy(os.Stdout, &b)

}

type notifier interface {
	notify()
}

type User struct {
	name  string
	email string
}

// As long as User has the function that interface required, then it is a notifier
func (u *User) notify() {
	fmt.Printf("Sending notification to user: %s<%s>\n", u.name, u.email)
}

type Admin struct {
	name  string
	email string
}

func (a *Admin) notify() {
	fmt.Printf("Sending notifiction to admin: %s<%s>\n", a.name, a.email)
}

type Embed_Admin struct {
	User
	level string
}

func (ea *Embed_Admin) notify() {
	fmt.Printf("Sending notification to embed admin %s<%s>\n", ea.name, ea.email)
}

func sendNotification(n notifier) {
	n.notify()
}

func interfaceImplementation() {
	interface_user := User{"Eddie", "Eddie@Gmail.com"}
	sendNotification(&interface_user)
	interface_admin := Admin{"EddieBoss", "Eddie.Boss@Gmail.com"}
	sendNotification(&interface_admin)
	interface_embed_admin := Embed_Admin{interface_user, "super"}
	interface_embed_admin.User.notify()
	interface_embed_admin.notify() // if embed admin implements interface functions, then it won't use inner type's implement interface function
	sendNotification(&interface_embed_admin)
}
