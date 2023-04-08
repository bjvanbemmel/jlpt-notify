package main

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

func main() {
    client := twilio.NewRestClient()
    
    params := &api.CreateMessageParams{}
    params.SetBody("Hello, 世界!")
    params.SetFrom(os.Getenv("SMS_SENDER"))
    params.SetTo(os.Getenv("SMS_RECEIVER"))

    if _, err := client.Api.CreateMessage(params); err != nil {
        log.Fatal(err.Error())
    }
}
