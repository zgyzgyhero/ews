## EWS Exchange Web Service
Is a exchange web service client in golang 

very dumb, hacky, possibly flaky package to send emails from an Exchange server via EWS (in the event yours doesn't expose a SMTP server)

usage:
```
func main() {

	client := ews.NewClientWithConfig(
		"https://outlook.office365.com/EWS/Exchange.asmx",
		"example@mhewedy.onmicrosoft.com",
		"systemsystem@123",
		&ews.Config{Dump: true},
	)

	err := ews.CreateItem(client,
		[]string{"mhewedy@gmail.com", "someone@else.com"},
		"An email subject",
		"The email body, as plain text",
	)

	if err != nil {
		log.Fatal("err: ", err.Error())
	}

	fmt.Println("mail sent")
}
```
the other exported types are just the raw data structures for the request XML; you can ignore them

I'm not sure if I'll develop this further; feel free to (warning: here be SOAP)
some resources I used are in comments in the code

TODOs
- figure out why UTF-8 isn't used for email bodies, or how to force the encoding
