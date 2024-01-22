package services

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/adjust/rmq/v5"
	config "github.com/umer-emumba/BudgetBuddy/configs"
	"github.com/umer-emumba/BudgetBuddy/types"
	"github.com/umer-emumba/BudgetBuddy/utils"
)

type QueueService struct {
	queues map[string]rmq.Queue
	helper *utils.Helper
}

var instance *QueueService

func GetQueueService() (*QueueService, error) {
	if instance == nil {
		var err error
		instance, err = NewQueueService()
		if err != nil {
			return nil, err
		}
	}
	return instance, nil
}

func NewQueueService() (*QueueService, error) {

	service := &QueueService{
		queues: make(map[string]rmq.Queue),
		helper: utils.NewHelper(),
	}

	if err := service.instantiateQueue("default"); err != nil {
		return nil, err
	}

	if err := service.startConsumer("default"); err != nil {
		return nil, err
	}

	return service, nil
}

func (s *QueueService) instantiateQueue(queueName string) error {
	queue, err := config.RmqConnection.OpenQueue(queueName)
	if err != nil {
		return err
	}
	fmt.Println(queue)
	s.queues[queueName] = queue
	return nil
}

func (s *QueueService) GetQueue(name string) rmq.Queue {
	return s.queues[name]
}

func (s *QueueService) startConsumer(queueName string) error {
	queue := s.GetQueue(queueName)
	if queue == nil {
		return fmt.Errorf("queue not found: %s", queueName)
	}
	queue.StartConsuming(10, time.Second)
	queue.AddConsumerFunc("default-consumer", func(delivery rmq.Delivery) {
		var job types.QueueJob
		if err := json.Unmarshal([]byte(delivery.Payload()), &job); err != nil {
			fmt.Println(err)
			delivery.Reject()
		}

		switch job.Name {
		case "send-email":
			s.helper.SendMail(job.MailData)
			delivery.Ack()

		}
	})
	return nil
}
