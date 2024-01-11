package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"azareal/src"
)

var verCmd = &cobra.Command{
	Use:   "version",
	Short: "prints version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("azareal v%s - LitFill\n", src.Version)
	},
}

func init() {
	root.AddCommand(verCmd)
}
