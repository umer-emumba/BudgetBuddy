package config

import (
	"log"

	"github.com/adjust/rmq/v5"
)

var RmqConnection rmq.Connection

func ConnectRMQ() error {
	conn, err := rmq.OpenConnection("queue service", "tcp", AppCfg.RedistConfig.Host, 1, nil)
	if err != nil {
		return err
	}
	RmqConnection = conn

	cleaner := rmq.NewCleaner(conn)

	returned, err := cleaner.Clean()
	if err != nil {
		log.Printf("failed to clean: %s", err)

	}
	log.Printf("cleaned %d", returned)

	return nil
}
