package services

import "testing"

/* implement MessageService (interface) */
type MockMessageService struct {
	sendCalled    bool
	messageToSend string
	returnValue   error
}

func (mms *MockMessageService) Send(message string) error {
	mms.sendCalled = true
	mms.messageToSend = message
	return mms.returnValue
}

func TestMessageProcessor_Process(t *testing.T) {
	tests := []struct {
		name string
		msg  string
		want bool
	}{
		{name: "Emails the given message", msg: "test message", want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			//arrange
			//dependency
			mms := &MockMessageService{
				returnValue: nil,
			}

			//sut
			mp := &MessageProcessor{
				messageService: mms,
			}

			//act
			got := mp.Process(tt.msg)

			//assert
			if !mms.sendCalled {
				t.Fatal("MessageProcessor.Process(), did not send the message")
			}

			if tt.msg != mms.messageToSend {
				t.Errorf("MessageProcessor.Process(), message to send = %q, message sent = %q\n", tt.msg, mms.messageToSend)
			}

			if got != true {
				t.Errorf("MessageProcessor.Process(), Should return %t for successfully sending the message, but retured = %t\n", true, got)
			}
		})
	}
}

/* Create a test to test the behavior of MessageProcessor.Process() when the MessageService.Send() returns an error */
