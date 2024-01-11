/*
Copyright © 2024 LitFill <marrazzy54@gmail.com>
*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// root represents the base command when called without any subcommands
var root = &cobra.Command{
	Use:   "azareal",
	Short: "Azareal - cli untuk jadwal madin",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) {
	// fmt.Println("root cmd.")
	// },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := root.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.azareal.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	root.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
