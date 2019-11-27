package main

import (
	"fmt"
	"github.com/mhewedy/ews/ews"
	"io/ioutil"
	"log"
)

func main() {
	b, err := ews.BuildTextEmail(
		"example@mhewedy.onmicrosoft.com",
		[]string{"mhewedy@gmail.com", "someone@else.com"},
		"An email subject",
		[]byte("The email body, as plain text"))
	if err != nil {
		// handle err
	}
	resp, err := ews.Issue(
		"https://outlook.office365.com/EWS/Exchange.asmx",
		"example@mhewedy.onmicrosoft.com",
		"systemsystem@123",
		b)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		bbs, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		log.Fatal(string(bbs))
	}

	fmt.Println("mail sent")
}
