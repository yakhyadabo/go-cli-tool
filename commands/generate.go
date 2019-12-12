package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate the code",
	Long:  `It generates the code that can be later used for deployment.`,
	Run: generate ,
}

var component string
var environment string
var region string

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringVarP(&component, "comp", "c", "", "The name of the component")
	generateCmd.Flags().StringVarP(&environment, "env", "e", "", "The environment of the component")
	generateCmd.Flags().StringVarP(&region, "reg", "r", "", "The region where the component will be deployed")
}

func generate (cmd *cobra.Command, args []string) {
	fmt.Println("Generate " +  component + "-" + environment + " for deployment in " + region)
}