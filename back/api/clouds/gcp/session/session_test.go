package session

import (
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/logger"
	"testing"
)

func Test_fetchAvailableProjects(t *testing.T) {
	for _, session := range sessions {
		logger.Info.Println(session.credentialFilePath, session.project.Name, session.project.ProjectId)
	}
}
