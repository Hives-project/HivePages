package consumer

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/Hives-project/HivePages/pkg/config"
	"github.com/Hives-project/HivePages/pkg/page"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
)

func StartKafkaConsumer(cfg config.KafkaConfig, pageSvc page.PageService) {
	scram, err := scram.Mechanism(scram.SHA512, cfg.User, cfg.Password)
	if err != nil {
		log.Fatal(err)
	}
	dialer := &kafka.Dialer{
		Timeout:       10 * time.Second,
		DualStack:     true,
		SASLMechanism: scram,
		TLS:           &tls.Config{},
	}
	readerConfig := kafka.ReaderConfig{
		Brokers:     []string{fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)},
		GroupTopics: []string{"sndosdzx-createPage", "sndosdzx-getUsername"},
		GroupID:     "pages",
		Dialer:      dialer,
	}

	reader := kafka.NewReader(readerConfig)

	go func() {
		for {
			m, err := reader.ReadMessage(context.Background())
			if err != nil {
				fmt.Println("Error", err)
				continue
			}
			RouteMessagesFromTopics(pageSvc, m)
		}
	}()

	log.Println("service is running..")

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)
	<-c

	log.Println("service is shutting down..")

	os.Exit(0)
}
