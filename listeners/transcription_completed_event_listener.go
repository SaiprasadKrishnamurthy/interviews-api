package listeners

import (
	"encoding/json"
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/saiprasadkrishnamurthy/interviews-api/models"
)

// TranscriptionCompletedReceivedEventListener foo listener.
type TranscriptionCompletedReceivedEventListener struct {
	*BaseListener
}

// Handle on message function of all the listeners.
func (l *TranscriptionCompletedReceivedEventListener) Handle(msg *nats.Msg) {
	event := models.TranscriptionCompletedEvent{}
	json.Unmarshal(msg.Data, &event)
	fmt.Printf("\n\n\n\n ==================== Received ================ %+v", event)
	// TODO Analyse and save it to ES as results.
}
