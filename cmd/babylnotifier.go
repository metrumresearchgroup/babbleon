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
		viper.BindEnv("MESSAGE")
		viper.BindEnv("TARGET")
		viper.BindEnv("DEBUG")


		var config app.Configuration

		viper.Unmarshal(&config)

		logger := log.New()

		if config.Debug {
			logger.SetLevel(log.DebugLevel)
		}

		babylonDetails := app.NewBabylonModelDetails()
		slackrepo := repository.NewSlackRepository(logger, config.OAuthToken)
		service := app.NewSlackNotificationService(babylonDetails,
			&config,
			slackrepo,
			logger,
			config.Message,
			config.Target)

		application := api.NewSlackHandler(logger,service)


		err := application.Process()

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
	const messageIdentifier string = "message"
	rootCmd.Flags().String(messageIdentifier,"", "The person / number / identifier to which the " +
		"message will be sent")
	viper.BindPFlag(messageIdentifier, rootCmd.Flags().Lookup(messageIdentifier))

	const  targetIdentifier string = "target"
	rootCmd.Flags().String(targetIdentifier, "" , "The message to send out")
	viper.BindPFlag(targetIdentifier, rootCmd.Flags().Lookup(targetIdentifier))
}
