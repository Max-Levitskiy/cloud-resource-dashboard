package session

import (
	"context"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/conf"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/logger"
	"google.golang.org/api/cloudresourcemanager/v1"
	"google.golang.org/api/option"
	"io/ioutil"
	"os"
	"strings"
)

var (
	files    = getCredentialsFiles()
	sessions = fetchSessions()
)

func GetSessions() []*Session {
	return sessions
}

type Session struct {
	credentialFilePath string
	project            *cloudresourcemanager.Project
}

func (s *Session) GetProject() *cloudresourcemanager.Project {
	return s.project
}
func (s *Session) GetProjectId() string {
	return s.project.ProjectId
}

func (s *Session) GetProjectName() string {
	return s.project.Name
}

func (s *Session) GetCredentialsOption() option.ClientOption {
	return option.WithCredentialsFile(s.credentialFilePath)
}

func fetchSessions() []*Session {
	sessionsChannel := make(chan []*Session, len(files))

	for _, file := range files {
		credentialFilePath := getCredentialsPath() + "/" + file.Name()
		if service, err := cloudresourcemanager.NewService(context.Background(), option.WithCredentialsFile(credentialFilePath)); err == nil {
			go fetchProjectsAndAppendSessions(sessionsChannel, credentialFilePath, service)
		} else {
			logger.Error.Println(err)
		}
	}

	var sessions []*Session
	for i := 0; i < len(files); i++ {
		sessions = append(sessions, <-sessionsChannel...)
	}

	close(sessionsChannel)

	return sessions
}

func fetchProjectsAndAppendSessions(ch chan<- []*Session, credentialFilePath string, manager *cloudresourcemanager.Service) {
	var sessions []*Session
	if response, err := manager.Projects.List().Do(); err == nil {
		for _, project := range response.Projects {
			sessions = append(sessions, &Session{
				credentialFilePath: credentialFilePath,
				project:            project,
			})
		}
		ch <- sessions
	} else {
		logger.Error.Println(err)
	}
}

func getCredentialsFiles() []os.FileInfo {
	credentialsPath := getCredentialsPath()

	var files []os.FileInfo
	if fileInfos, err := ioutil.ReadDir(credentialsPath); err == nil {
		for _, file := range fileInfos {
			if !file.IsDir() && strings.HasSuffix(file.Name(), ".json") {
				files = append(files, file)
			}
		}
	} else {
		logger.Warn.Println(err)
	}
	return files
}

func getCredentialsPath() string {
	var credentialsPath string
	if conf.Inst.GCP.CredentialsPath == nil {
		credentialsPath = conf.Inst.HomeDir + "/.gcp/credentials"
	} else {
		credentialsPath = *conf.Inst.GCP.CredentialsPath
	}
	return credentialsPath
}
