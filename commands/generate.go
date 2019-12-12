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

var component string
var environment string
var region string

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringVarP(&component,"comp", "c","", "Set the name of the component")
	generateCmd.Flags().StringVarP(&environment,"env", "e","", "Set the environment of the component")
	generateCmd.Flags().StringVarP(&region,"reg", "r","", "Set the region where the component will be deployed")
}

