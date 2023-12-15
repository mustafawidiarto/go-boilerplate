package main

import (
	"github.com/mustafawidiarto/go-boilerplate/command"

	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// The main function initializes the root command and executes it.
func main() {
	if err := NewRootCommand().Execute(); err != nil {
		log.Fatal().Msgf("failed run app: %s", err.Error())
	}
}

// NewRootCommand function returns a Cobra command that serves as the root command
// for the service-go application.
func NewRootCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:   "integration-go",
		Short: "Run service",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	command.AddCommand(serverCmd(), cronCmd())
	return command
}

// The cronCmd returns a new instance of cobra.Command that represents the "cron" command
// for run the cron job
func cronCmd() *cobra.Command {
	var command = &cobra.Command{
		Use:   "cron",
		Short: "Run cron",
		Run: func(cmd *cobra.Command, args []string) {
			cron := command.NewCron()
			cron.Run()
		},
	}

	return command
}

// The serverCmd function returns a Cobra command that starts a server on the specified port
// or the default port if no port is specified, defaults to 8080
func serverCmd() *cobra.Command {
	var port int
	var command = &cobra.Command{
		Use:   "server",
		Short: "Run server",
		Run: func(cmd *cobra.Command, args []string) {
			srv := command.NewServer()
			srv.Run(port)
		},
	}

	command.Flags().IntVar(&port, "port", 8080, "Listen on given port")
	return command
}
