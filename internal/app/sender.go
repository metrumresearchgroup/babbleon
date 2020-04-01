package app

import (
	"github.com/metrumresearchgroup/babbleon/internal/repository"
	log "github.com/sirupsen/logrus"
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
	babylondetails *repository.BabylonModelDetails
}

func NewSlackNotificationService(details *repository.BabylonModelDetails,
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
	Additional []string `mapstructure:"additional_message_values"`
}


