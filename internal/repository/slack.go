package repository

import (
	log "github.com/sirupsen/logrus"
	"github.com/slack-go/slack"
)

const SlackMessageTemplate string = `Hello! Job with identifier of ` + "`" + `{{ .ModelDetails.Filename  }}` + "`" +  ` has completed.

Was it successful? *{{ .ModelDetails.Successful }}* {{ if eq .ModelDetails.Successful false}} :red_circle: :red_circle: {{ end }} {{ if eq .ModelDetails.Successful true }} :white_check_mark: :white_check_mark: {{ end }}

• Output Directory:	{{ .ModelDetails.OutputDirectory }}
• Model:			{{ .ModelDetails.Model }}
• Original Path:	{{ .ModelDetails.Path }} 
{{ if not .ModelDetails.Successful }}• Error details for the failure are: {{ .ModelDetails.Error }}
{{- end }}

{{ $length := len .Additional }} {{ if gt $length 0 }}
Additional Values:
	{{- range .Additional }}
• {{ . }}
	{{- end}}
{{- end}}
`

type SlackRepository struct {
	logger *log.Logger
	token string //For authentication
	babylonDetails *BabylonModelDetails
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

	s.logger.WithFields(log.Fields{
		"response" : response,
	}).Debug("Sent notification")

	return err
}