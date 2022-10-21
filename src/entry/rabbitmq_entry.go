package entry

import (
	"github.com/Fonzeca/Chatline/src/services"
)

func NewRabbitMqDataEntry() {
	channel := services.GlobalChannel

	_, err := channel.QueueDeclare("chatline", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}
}
