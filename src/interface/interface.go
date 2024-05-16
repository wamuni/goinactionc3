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
	fmt.Printf("Sending notification to user: %s<%s>", u.name, u.email)
}

func sendNotification(n notifier) {
	n.notify()
}

func interfaceImplementation() {
	interface_user := User{"Eddie", "Eddie@Gmail.com"}
	sendNotification(&interface_user)
}
