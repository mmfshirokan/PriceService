package consumer

import (
	"context"
	"fmt"
	"io"

	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

type consumer struct {
	brokerURL string
	topic     string
}

type Reader interface {
	Read(ctx context.Context)
}

func New(brokerURL string, topic string) Reader {
	return &consumer{
		brokerURL: brokerURL,
		topic:     topic,
	}
}

func (cons *consumer) Read(ctx context.Context) {
	for i := 0; i < 3; i++ {
		go func(partition int) {
			reader := kafka.NewReader(kafka.ReaderConfig{
				Brokers:   []string{cons.brokerURL},
				Topic:     cons.topic,
				Partition: i,
			})

			gorutineNum := reader.Config().Partition
			log.Info("Sucssesful start ", gorutineNum)

			for {
				msg, err := reader.ReadMessage(ctx)
				if err == io.EOF {
					log.Info(fmt.Sprintf("exiting kafka read gorutine-%d", gorutineNum))
					break
				}
				if err != nil {
					log.Error(fmt.Sprintf("error occured in gorutine-%d: %v", gorutineNum, err))
					break
				}

				log.WithFields(log.Fields{
					"key":   string(msg.Key),
					"value": string(msg.Value),
				})
			}

		}(i)
	}
}
