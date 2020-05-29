package listeners

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/saiprasadkrishnamurthy/interviews-api/models"
)

// TranscodingCompletedReceivedEventListener foo listener.
type TranscodingCompletedReceivedEventListener struct {
	*BaseListener
}

// Handle on message function of all the listeners.
func (l *TranscodingCompletedReceivedEventListener) Handle(msg *nats.Msg) {
	event := models.TranscodingCompletedEvent{}
	json.Unmarshal(msg.Data, &event)
	log.Println("\n\n  =============================== ")
	log.Printf("Received in TranscodingCompletedReceivedEventListener: %+v ", event)
	// TODO Call Audio -> Text transcription API and save the text in the same directory alongside the video file.
	// Emit a analyse event that analyses the results.
}
