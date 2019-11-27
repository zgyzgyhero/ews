package main

import (
	"fmt"
	"github.com/mhewedy/ews/ews"
	"log"
)

func main() {

	client := ews.NewClientWithConfig(
		"https://outlook.office365.com/EWS/Exchange.asmx",
		"example@mhewedy.onmicrosoft.com",
		"systemsystem@123",
		&ews.Config{Dump: true},
	)

	err := ews.CreateItem(client,
		"example@mhewedy.onmicrosoft.com",
		[]string{"mhewedy@gmail.com", "someone@else.com"},
		"An email subject",
		"The email body, as plain text",
	)

	if err != nil {
		log.Fatal("err: ", err)
	}

	fmt.Println("mail sent")
}
