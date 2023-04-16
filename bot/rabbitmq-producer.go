package bot

// func (h *BotHandler) ReceiveMessageToQueue(msg string, queueName string) error {
// 	err := h.ch.Publish(
// 		"",
// 		queueName,
// 		false,
// 		false,
// 		amqp.Publishing{
// 			ContentType: "text/plain",
// 			Body: []byte(msg),
// 		},
// 	)
// 	if err != nil {
// 		log.Println("Error publish msg", err.Error())
// 	}
// 	return nil

// }