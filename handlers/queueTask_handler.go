package handlers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/umer-emumba/BudgetBuddy/types"
	"github.com/umer-emumba/BudgetBuddy/utils"
)

type QueueTaskHandler struct {
	helper *utils.Helper
}

func NewQueueTaskHandler() QueueTaskHandler {
	return QueueTaskHandler{
		helper: utils.NewHelper(),
	}
}

func (h QueueTaskHandler) HandleEmailDeliveryTask(ctx context.Context, t *asynq.Task) error {
	var data types.MailOptions
	if err := json.Unmarshal(t.Payload(), &data); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	h.helper.SendMail(data)
	return nil
}
