package cmd

import (
	"fmt"
	"strings"

	"github.com/semaphoreci/cli/cmd/utils"
	"github.com/semaphoreci/cli/config"
	"github.com/spf13/cobra"

	client "github.com/semaphoreci/cli/api/client"
)

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to a Semaphore endpoint",
	Args:  cobra.ExactArgs(2),
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		host := args[0]
		token := args[1]

		baseClient := client.NewBaseClient(token, host, "")
		client := client.NewProjectV1AlphaApiWithCustomClient(baseClient)

		_, err := client.ListProjects()

		if err != nil {
			fmt.Fprintf(cmd.OutOrStderr(), "%s", err)
			utils.Exit(1)
		}

		name := strings.Replace(host, ".", "_", -1)

		config.SetActiveContext(name)

		err = config.SetAuth(token)
		if err != nil {
			fmt.Fprintf(cmd.OutOrStderr(), "%s", err)
		}

		err = config.SetHost(host)
		if err != nil {
			fmt.Fprintf(cmd.OutOrStderr(), "%s", err)
		}

		cmd.Printf("connected to %s\n", host)
	},
}

func init() {
	RootCmd.AddCommand(connectCmd)
}
