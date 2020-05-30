package listeners

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"

	speech "cloud.google.com/go/speech/apiv1"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"

	"github.com/nats-io/nats.go"
	"github.com/saiprasadkrishnamurthy/interviews-api/config"
	"github.com/saiprasadkrishnamurthy/interviews-api/models"
	"github.com/saiprasadkrishnamurthy/interviews-api/utils"
)

// TranscodingCompletedReceivedEventListener foo listener.
type TranscodingCompletedReceivedEventListener struct {
	*BaseListener
}

// Handle on message function of all the listeners.
func (l *TranscodingCompletedReceivedEventListener) Handle(msg *nats.Msg) {
	event := models.TranscodingCompletedEvent{}
	json.Unmarshal(msg.Data, &event)
	chans := []chan models.TranscriptionResult{}
	for _, res := range event.TranscodingResults {
		c := make(chan models.TranscriptionResult)
		chans = append(chans, c)
		go transcript(c, res)
	}
	transcriptionResults := []models.TranscriptionResult{}
	for _, c := range chans {
		transcriptionResult := utils.ExtractTranscriptionResult(c, config.GetConfig().Transcription.TimeoutInSeconds)
		transcriptionResults = append(transcriptionResults, transcriptionResult)
	}
	natsConn := config.GetNatsConnection()
	natsSubject := config.GetConfig().Nats.TranscriptionCompletedSubject
	transcriptionCompletedEvent := models.TranscriptionCompletedEvent{
		SessionID:            event.SessionID,
		CandidateID:          event.CandidateID,
		TranscriptionResults: transcriptionResults,
	}
	doc, _ := json.Marshal(transcriptionCompletedEvent)
	err := natsConn.Publish(natsSubject, doc)
	if err != nil {
		log.Println(err)
	}
}

func transcript(c chan models.TranscriptionResult, result models.TranscodingResult) {
	ctx := context.Background()
	failure := "Success"
	// Creates a client.
	client, err := speech.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		failure = err.Error()
	}

	inputDir := filepath.Join(config.GetConfig().VideoStore.WorkDir, result.SessionID, result.CandidateID, result.Question)
	// Sets the name of the audio file to transcribe.
	filename := filepath.Join(inputDir, "answer.mp3")

	// Reads the audio file into memory.
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
		failure = err.Error()
	}

	// Detects speech in the audio file.
	resp, err := client.Recognize(ctx, &speechpb.RecognizeRequest{
		Config: &speechpb.RecognitionConfig{
			Encoding:        8, // for mp3
			SampleRateHertz: config.GetConfig().Transcription.SampleRateHertz,
			LanguageCode:    config.GetConfig().Transcription.Language,
		},
		Audio: &speechpb.RecognitionAudio{
			AudioSource: &speechpb.RecognitionAudio_Content{Content: data},
		},
	})
	if err != nil {
		log.Fatalf("failed to recognize: %v", err)
		failure = err.Error()
	}

	// Prints the results.
	text := ""
	var confidence float32
	for _, result := range resp.Results {
		for _, alt := range result.Alternatives {
			text += alt.Transcript + " "
			confidence = alt.Confidence
		}
	}
	textFile := filepath.Join(inputDir, "answer.txt")
	b := []byte(text)
	if err := ioutil.WriteFile(textFile, b, 0644); err != nil {
		failure = err.Error()
	}

	tr := models.TranscriptionResult{
		SessionID:   result.SessionID,
		CandidateID: result.CandidateID,
		Question:    result.Question,
		Result:      failure,
		Confidence:  confidence,
	}
	c <- tr
}
