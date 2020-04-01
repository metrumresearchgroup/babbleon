package cmd

import (
	"fmt"
	"github.com/metrumresearchgroup/babbleon/internal/api"
	"github.com/metrumresearchgroup/babbleon/internal/app"
	"github.com/metrumresearchgroup/babbleon/internal/repository"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "babbleon",
	Short: "babbleon is an easy-to-use notification tool",
	Long: `babbleon is an easy-to-use notification tool that pairs
very well with Babylon. This facilitates the notification of users after the termination 
of a phase of execution`,
	Run: func(cmd *cobra.Command, args []string) {
		viper.SetEnvPrefix("babbleon")
		viper.AutomaticEnv()
		//Bindings
		viper.BindEnv("OAUTH_TOKEN")
		viper.BindEnv("TARGET")
		viper.BindEnv("DEBUG")


		var config app.Configuration

		viper.Unmarshal(&config)

		logger := log.New()

		if config.Debug {
			logger.SetLevel(log.DebugLevel)
		}

		babylonDetails := repository.NewBabylonModelDetails()
		babylonMessage := repository.BabylonSlackMessage{
			ModelDetails: babylonDetails,
			Additional:   config.Additional,
		}
		message, err := babylonMessage.PopulatedTemplate(repository.SlackMessageTemplate)

		if err != nil {
			log.Fatalf("Failure to process the message from template: %s", err)
		}

		slackrepo := repository.NewSlackRepository(logger, config.OAuthToken)
		service := app.NewSlackNotificationService(babylonDetails,
			&config,
			slackrepo,
			logger,
			message,
			config.Target)

		application := api.NewSlackHandler(logger,service)


		err = application.Process()

		if err != nil {
			logger.WithFields(log.Fields{
				"to" : config.Target,
				"message" : config.Message,
			}).Errorf("An error occurred sending the requested data")
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init(){
	const  targetIdentifier string = "target"
	rootCmd.Flags().String(targetIdentifier, "" , "The message to send out")
	viper.BindPFlag(targetIdentifier, rootCmd.Flags().Lookup(targetIdentifier))

	const additionalValuesIdentifier string = "additional_message_values"
	rootCmd.Flags().StringSlice(additionalValuesIdentifier, []string{} , "Additional information to include in the message placed exactly as provided at the end of the message")
	viper.BindPFlag(additionalValuesIdentifier,rootCmd.Flags().Lookup(additionalValuesIdentifier))
}
