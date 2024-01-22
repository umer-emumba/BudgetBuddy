package types

type QueueJob struct {
	Name     string      `json:"name"`
	MailData MailOptions `json:"mail_data"`
}
