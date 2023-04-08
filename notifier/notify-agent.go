package notifier

import (
	"errors"
	"fmt"
	"time"

	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

type NotifyAgent struct {
	Sender     string
	Receiver   string
	RestClient *twilio.RestClient
}

var (
	ErrArgNil error = errors.New("Argument may not be nil.")
)

func (n *NotifyAgent) SetSender(s string) error {
	if s == "" {
		return ErrArgNil
	}
	n.Sender = s

	return nil
}

func (n *NotifyAgent) SetReceiver(s string) error {
	if s == "" {
		return ErrArgNil
	}
	n.Receiver = s

	return nil
}

func (n *NotifyAgent) SetRestClient(r *twilio.RestClient) error {
	if r == nil {
		return ErrArgNil
	}
	n.RestClient = r

	return nil
}

func (n NotifyAgent) SendMessage(msg string) error {
	msg = fmt.Sprintf("[%v] %s", time.Now().Format("2-1-2006 15:4:5"), msg)

	params := &api.CreateMessageParams{}
	params.SetBody(msg)
	params.SetFrom(n.Sender)
	params.SetTo(n.Receiver)

	_, err := n.RestClient.Api.CreateMessage(params)

	return err
}
