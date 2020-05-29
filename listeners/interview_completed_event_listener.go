package listeners

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"github.com/nats-io/nats.go"
	"github.com/saiprasadkrishnamurthy/interviews-api/config"
	"github.com/saiprasadkrishnamurthy/interviews-api/models"
	"github.com/saiprasadkrishnamurthy/interviews-api/utils"
	"github.com/xfrr/goffmpeg/transcoder"
)

// InterviewCompletedReceivedEventListener foo listener.
type InterviewCompletedReceivedEventListener struct {
	*BaseListener
}

// Handle on message function of all the listeners.
func (l *InterviewCompletedReceivedEventListener) Handle(msg *nats.Msg) {
	event := models.InterviewCompletedEvent{}
	json.Unmarshal(msg.Data, &event)
	log.Printf("Received in InterviewCompletedEventListener: %+v ", event)
	inputDir := filepath.Join(config.GetConfig().VideoStore.WorkDir, event.SessionID, event.CandidateID)

	dirs, _ := ioutil.ReadDir(inputDir)

	chans := []chan models.TranscodingResult{}
	for _, dir := range dirs {
		if dir.IsDir() {
			c := make(chan models.TranscodingResult)
			chans = append(chans, c)
			baseDir := filepath.Join(inputDir, dir.Name())
			inputPath := filepath.Join(inputDir, dir.Name(), "answer.webm")
			outputPath := filepath.Join(inputDir, dir.Name(), "answer.mp3")
			go transcodeToAudio(event, baseDir, inputPath, outputPath, c)
		}
	}

	transcodingResults := []models.TranscodingResult{}
	for _, c := range chans {
		transcodingResult := utils.ExtractTranscodingResult(c, config.GetConfig().Transcoding.TimeoutInSeconds)
		fmt.Printf("%+v", transcodingResult)
		transcodingResults = append(transcodingResults, transcodingResult)
	}
	natsConn := config.GetNatsConnection()
	natsSubject := config.GetConfig().Nats.TranscodingCompletedSubject
	transcodingCompletedEvent := models.TranscodingCompletedEvent{
		SessionID:          event.SessionID,
		CandidateID:        event.CandidateID,
		TranscodingResults: transcodingResults,
	}
	doc, _ := json.Marshal(transcodingCompletedEvent)
	err := natsConn.Publish(natsSubject, doc)
	if err != nil {
		log.Println(err)
	}
	log.Printf("Published message %s on subject %s ", doc, natsSubject)
}

func transcodeToAudio(event models.InterviewCompletedEvent, baseDir string, inputPath string, outputPath string, c chan models.TranscodingResult) {
	fmt.Println(" Input Path: ", inputPath)
	// Create new instance of transcoder
	trans := new(transcoder.Transcoder)

	// Initialize transcoder passing the input file path and output file path

	e := trans.InitializeEmptyTranscoder()
	if e != nil {
		fmt.Println(e)
	}
	trans.SetInputPath(inputPath)
	trans.SetOutputPath(outputPath)

	// Start transcoder process without checking progress
	done := trans.Run(false)

	lastIndexOfSlash := strings.LastIndex(baseDir, "/") + 1
	runes := []rune(baseDir)
	fmt.Println(" Runes: ", runes)
	question := string(runes[lastIndexOfSlash:])
	fmt.Println(" Question: ", question)
	// This channel is used to wait for the process to end
	result := "Success"
	if err := <-done; err != nil {
		fmt.Println("Error while transcoding: ", err)
		result = "Error"
	} else {
		fmt.Println("Successfully Transcoded to: ", outputPath)
		result = "Success"
	}
	c <- models.TranscodingResult{
		SessionID:   event.SessionID,
		CandidateID: event.CandidateID,
		Question:    question,
		Result:      result,
	}
}
