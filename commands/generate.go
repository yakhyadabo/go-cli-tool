package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate the code",
	Long:  `It generates the code that can be later used for deployment.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is the generate command!")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
