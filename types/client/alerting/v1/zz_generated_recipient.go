package client

const (
	RecipientType            = "recipient"
	RecipientFieldNotifierId = "notifierId"
	RecipientFieldRecipient  = "recipient"
)

type Recipient struct {
	NotifierId string `json:"notifierId,omitempty"`
	Recipient  string `json:"recipient,omitempty"`
}
