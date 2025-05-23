package notification

import (
	"encoding/json"
	"fmt"
	"go-ecommerce-app/config"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type NotificationClient interface {
	SendSMS(phone, message string) error
}

type notificationClient struct {
	config config.AppConfig
}

func NewNotificationClient(config config.AppConfig) NotificationClient {
	return &notificationClient{
		config: config,
	}
}

func (c notificationClient) SendSMS(phone, message string) error {
	accountSid := c.config.TwilioAccountSID
	authToken := c.config.TwilioAccountToken

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(phone)
	params.SetFrom(c.config.TwilioFromPhone)
	params.SetBody(message)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println("Error sending SMS message: " + err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}

	return nil
}
