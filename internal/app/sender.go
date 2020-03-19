package app

import (
	"github.com/metrumresearchgroup/babbleon/internal/repository"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
	"text/template"
)

type NotificationService interface {
	Notify() error
}

type SlackNotificationService struct {
	messageContent string //What we're actually going to send
	messageTarget string
	messageTemplate template.Template
	notifier repository.NotificationRepository
	configuration *Configuration
	logger *log.Logger
	babylondetails *BabylonModelDetails
}

func NewSlackNotificationService(details *BabylonModelDetails,
	configuration *Configuration,
	notifier repository.NotificationRepository,
	logger *log.Logger,
	message string,
	target string) *SlackNotificationService {
	return &SlackNotificationService{
		messageContent: message,
		messageTarget:  target,
		notifier:       notifier,
		configuration:  configuration,
		logger:         logger,
		babylondetails: details,
	}
}



func (n *SlackNotificationService) Notify() error {
	return n.notifier.SendNotification(n.messageContent, n.messageTarget)
}


type Configuration struct {
	OAuthToken string `mapstructure:"OAUTH_TOKEN"`
	Debug bool `mapstructure:"DEBUG"`
	Message string `mapstructure:"message"`
	Target string `mapstructure:"target"`
}

type BabylonModelDetails struct {
	Path string
	Model string
	Filename string
	Extension string
	OutputDirectory string
	Successful bool
	Error string //Because we're only going to see it expressed as a string in the env
}

func NewBabylonModelDetails() *BabylonModelDetails {

	parsed, _ := strconv.ParseBool(os.Getenv("BABYLON_SUCCESSFUL"))

	return &BabylonModelDetails{
		Path:            os.Getenv("BABYLON_MODEL_PATH"),
		Model:           os.Getenv("BABYLON_MODEL"),
		Filename:        os.Getenv("BABYLON_MODEL_FILENAME"),
		Extension:       os.Getenv("BABYLON_MODEL_EXT"),
		OutputDirectory: os.Getenv("BABYLON_OUTPUT_DIR"),
		Successful:      parsed,
		Error:           os.Getenv("BABYLON_ERROR"),
	}
}
