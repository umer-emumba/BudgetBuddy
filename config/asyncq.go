package config

import (
	"github.com/hibiken/asynq"
)

func CreateAsyncQServer() *asynq.Server {
	return asynq.NewServer(
		asynq.RedisClientOpt{Addr: AppCfg.RedisConfig.Host},
		asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency: 10,
			// Optionally specify multiple queues with different priority.
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
		},
	)

}

func CreateAsynqClient() *asynq.Client {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: AppCfg.RedisConfig.Host})
	return client
}
