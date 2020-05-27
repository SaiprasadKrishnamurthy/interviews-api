package repositories

import "github.com/saiprasadkrishnamurthy/interviews-api/models"

// GetSession get session.
func GetSession(sessionID string) models.Session {
	session := models.Session{
		SessionID:          sessionID,
		Name:               "Sample Session",
		Description:        "Description",
		Instructions:       []string{"Instruction 1", "Instruction 2"},
		TotalTimeInSeconds: 60,
	}
	return session
}
