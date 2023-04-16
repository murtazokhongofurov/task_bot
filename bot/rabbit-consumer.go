package bot

import "fmt"

func (h *BotHandler) GetMessageFromQueue(queueName string) (string, error) {
	msgs, err := h.ch.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return "", fmt.Errorf("failed to consume message from queue: %v", err)
	}
	for msg := range msgs {
		return string(msg.Body), nil
	}
	return "", fmt.Errorf("failed to receive message from queue")
}