package cmd

import (
	"azareal/src"
	"fmt"

	"github.com/spf13/cobra"
)

var guruCmd = &cobra.Command{
	Use:   "guru",
	Short: "print nama guru berdasarkan kode",
	Long:  `mencetak guru berdsarkan argumen yang dimasukkan.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		kodeGuru := src.NewKodeGuru(args[0])
		fmt.Println(src.SemuaGuru[kodeGuru].Nama)
	},
}

func init() {
	root.AddCommand(guruCmd)
}
