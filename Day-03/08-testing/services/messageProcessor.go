package services

type MessageService interface {
	Send(message string) error
}

type MessageProcessor struct {
	messageService MessageService
}

func (mp *MessageProcessor) Process(msg string) bool {
	/*
		scrammbledMsg := "some dummy test - " + msg
		if err := mp.messageService.Send(scrammbledMsg); err != nil {
			return false
		}
	*/

	if err := mp.messageService.Send(msg); err != nil {
		return false
	}
	return true
}
