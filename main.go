package main

import (
	"fmt"
	"net"
	"net/smtp"
	"strings"
)

func main() {
	for true {
		var input string
		fmt.Print("Enter email address:")
		fmt.Scanln(&input)
		CheckEmail(input)
	}

}

func smtpc() (*smtp.Client, error) {
	mxrecords, err := net.LookupMX("gmail.com")
	if err != nil {
		return nil, err
	}
	mxrecord := mxrecords[0]
	smtpserver, err := smtp.Dial(mxrecord.Host + ":25")
	if err != nil {
		return nil, err
	}
	err = smtpserver.Hello("majhcc.com")
	if err != nil {
		return nil, err
	}
	return smtpserver, nil

}

func CheckEmail(input string) {
	smtpserver, err := smtpc()
	if err != nil {
		fmt.Println(err)
		return
	}
	smtpserver.Mail("randomemail@gmail.com")
	err = smtpserver.Rcpt(input)
	if err != nil {
		if strings.Contains(string(err.Error()), "550") {
			fmt.Println("Email not found")
		} else if strings.Contains(string(err.Error()), "421 please try again later") {
			fmt.Println("Service not available")
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Email found")
	}
}
