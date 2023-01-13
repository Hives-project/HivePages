package producer

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"log"
	"time"

	"github.com/Hives-project/HivePages/pkg/config"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
)

var (
	writer kafka.Writer
)

func Init(cfg config.KafkaConfig) {
	scram, err := scram.Mechanism(scram.SHA512, cfg.User, cfg.Password)
	if err != nil {
		log.Println(err)
		return
	}
	dialer := &kafka.Dialer{
		Timeout:       10 * time.Second,
		DualStack:     true,
		SASLMechanism: scram,
		TLS:           &tls.Config{},
	}
	writer = *kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{cfg.Host + ":" + cfg.Port},
		Dialer:  dialer,
	})
}

func UpdateUsername(username string, krabbelId string) error {
	writer.Topic = "sndosdzx-updateUsername"
	body, err := json.Marshal(map[string]string{"username": username, "krabbel_id": krabbelId})
	if err != nil {
		return err
	}
	return writer.WriteMessages(context.Background(), kafka.Message{Value: body})
}
