package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Test the code",
	Long:  `It tests the generated code before the deployment.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is the test command!")
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
