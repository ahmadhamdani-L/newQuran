package utils

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
	"golang.org/x/crypto/bcrypt"
)

// Connect with Kafka ...
func InitKafkaConn() *kafka.Conn {
	// config.LoadConfig()
	protocol := os.Getenv("KAFKA_PROTOCOL")
	broker := os.Getenv("BBROKER")
	topic := os.Getenv("TOPIC")
	conn, err := kafka.DialLeader(context.Background(), protocol, broker, topic, 0)
	if err != nil {
		log.Fatal("Failed to connect with kafka", err.Error())
	}
	_ = conn.SetWriteDeadline(time.Time{})
	log.Println("Connected with Kafka server successfully !!")
	return conn
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}