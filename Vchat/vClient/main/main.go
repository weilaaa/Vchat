package main

import (
	"Vchat/vClient/process"
	"fmt"
	"os"
)

var userID int
var userPW string
var userNAME string

func main() {

	var key int

	//display first menu page
	for true {
		fmt.Println("---->")
		fmt.Println("ARRIVED VCHAT ZONE")
		fmt.Println("PICK YOUR POISON")
		fmt.Println("---->")
		fmt.Println("1.LOGIN TO VCHAT")
		fmt.Println("2.REGISTER NOW")
		fmt.Println("3.RUN AWAY")
		fmt.Println("---->")

		// your choice
		fmt.Scanln(&key)
		switch key {
		case 1:
			fmt.Println("LOGGING IN")
			up := process.UserProcess{}
			err := up.Login(userID, userPW)
			if err != nil {
				fmt.Println("LOGIN FAILED")
				return
			}
		case 2:
			fmt.Println("CUSTOMER YOURSELF")
			up := process.UserProcess{}
			err := up.Register(userID, userPW, userNAME)
			if err != nil {
				fmt.Println("REGISTER FAILED")
				return
			}
		case 3:
			fmt.Println("GO BACK HOME, COWARD")
			os.Exit(0)
		default:
			fmt.Println("WHAT YOU GIVE A SHIT")
		}
	}

}
