package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy the code",
	Long:  `It deploys the code that has been generated.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is the deploy command!")
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
}
