package client

const (
	EmailConfigType              = "emailConfig"
	EmailConfigFieldReceiver     = "receiver"
	EmailConfigFieldRequireTLS   = "requireTLS"
	EmailConfigFieldSmtpHost     = "smtpHost"
	EmailConfigFieldSmtpPassword = "smtpPassword"
	EmailConfigFieldSmtpPort     = "smtpPort"
	EmailConfigFieldSmtpUsername = "smtpUsername"
)

type EmailConfig struct {
	Receiver     string `json:"receiver,omitempty"`
	RequireTLS   string `json:"requireTLS,omitempty"`
	SmtpHost     string `json:"smtpHost,omitempty"`
	SmtpPassword string `json:"smtpPassword,omitempty"`
	SmtpPort     string `json:"smtpPort,omitempty"`
	SmtpUsername string `json:"smtpUsername,omitempty"`
}
