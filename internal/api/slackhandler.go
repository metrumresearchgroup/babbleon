package api

import (
	"github.com/metrumresearchgroup/babbleon/internal/app"
	log "github.com/sirupsen/logrus"
)

type Slackhandler struct {
	logger *log.Logger
	Notifier app.NotificationService
}

func NewSlackHandler(log *log.Logger, notificationService app.NotificationService) *Slackhandler {
	return &Slackhandler{
		logger:   log,
		Notifier: notificationService,
	}
}

func (s *Slackhandler) Process() error {
	return s.Notifier.Notify()
}
