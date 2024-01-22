package utils

import (
	"encoding/json"

	"github.com/hibiken/asynq"
	"github.com/umer-emumba/BudgetBuddy/types"
)

func NewEmailDeliveryTask(mailOptions types.MailOptions) (*asynq.Task, error) {
	payload, err := json.Marshal(mailOptions)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(types.TypeEmailDelivery, payload), nil
}
