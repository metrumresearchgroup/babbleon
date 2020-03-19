package repository

import (
	log "github.com/sirupsen/logrus"
	"github.com/slack-go/slack"
)

type SlackRepository struct {
	logger *log.Logger
	token string //For authentication
}

func NewSlackRepository(l *log.Logger, t string ) *SlackRepository {
	return &SlackRepository{
		logger: l,
		token:  t,
	}
}

func (s *SlackRepository) SendNotification(notification string, destination string) error {
	api := slack.New(s.token)

	response , _ , err  := api.PostMessage(destination,slack.MsgOptionText(notification,false))

	log.WithFields(log.Fields{
		"response" : response,
	}).Debug("Sent notification")

	return err
}