package main

import (
	"fmt"
	"github.com/mhewedy/ews/ews"
)

func main() {

	client := &ews.Client{
		EWSAddr:  "https://outlook.office365.com/EWS/Exchange.asmx",
		Username: "example@mhewedy.onmicrosoft.com",
		Password: "systemsystem@123",
	}

	err := ews.CreateItem(client,
		"example@mhewedy.onmicrosoft.com",
		[]string{"mhewedy@gmail.com", "someone@else.com"},
		"An email subject",
		[]byte("The email body, as plain text"))

	if err != nil {
		// handle err
	}

	fmt.Println("mail sent")
}
