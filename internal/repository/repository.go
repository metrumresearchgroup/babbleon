package repository

import (
	"bytes"
	"os"
	"strconv"
	"text/template"
)

type NotificationRepository interface {
	SendNotification(notification string, destination string) error
}


type BabylonSlackMessage struct {
	ModelDetails *BabylonModelDetails
	Additional []string
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

func (b *BabylonSlackMessage) PopulatedTemplate(templateString string) (string, error) {
	t, err := template.New("babylon_details").Parse(templateString)
	outbytes := new(bytes.Buffer)
	if err != nil {
		return "", err
	}

	err = t.Execute(outbytes,b)

	if err != nil {
		return "", err
	}

	return outbytes.String(), nil
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