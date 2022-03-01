package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

const connStrEnvName = "CONNECTION_STRING"
const sqlCommandEnvName = "SQL_COMMAND"

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "This package will connect to a given postgres database and execute a series of commands. Set the env vars " + connStrEnvName + "and " + sqlCommandEnvName + " to configure",
	RunE: func(cmd *cobra.Command, args []string) error {
		connStr := os.Getenv(connStrEnvName)
		command := os.Getenv(sqlCommandEnvName)
		if err := validateInput(connStr, command); err != nil {
			return err
		}

		return ConnectAndExec(connStr, command)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

const expectedConnStr = "expected connection string format: gcppostgres://user:password@myproject/us-central1/instanceId/mydb"

func validateInput(connStr string, command string) error {
	if connStr == "" {
		return fmt.Errorf("env variable %s must be set, %s", connStrEnvName, expectedConnStr)
	}
	if _, err := url.Parse(connStr); err != nil {
		return fmt.Errorf("parsing connection URL, %s", expectedConnStr)
	}
	if command == "" {
		return fmt.Errorf("env variable %s must be set", sqlCommandEnvName)
	}
	return nil
}
