package listeners

import (
	"github.com/nats-io/nats.go"
	"github.com/saiprasadkrishnamurthy/interviews-api/config"
)

// BaseListener base listener.
type BaseListener struct {
	NatsConn *nats.Conn
}

// OnMessage on message function of all the listeners.
func (l *BaseListener) OnMessage(subject string, queue string, msgHandler nats.MsgHandler) {
	l.NatsConn.QueueSubscribe(subject, queue, msgHandler)
}

// InitializeAllListeners initializes all listeners.
func InitializeAllListeners(nats *nats.Conn) {
	baseListener := &BaseListener{NatsConn: config.GetNatsConnection()}
	initializeInterviewCompletedReceivedEventListener(baseListener)

	// List all your listeners here.

}

func initializeInterviewCompletedReceivedEventListener(baseListener *BaseListener) {
	l := &InterviewCompletedReceivedEventListener{BaseListener: baseListener}
	natsSubject := config.GetConfig().Nats.AnswersReceivedSubject
	natsQueue := "queue_for_" + natsSubject
	l.OnMessage(natsSubject, natsQueue, l.Handle)
}
