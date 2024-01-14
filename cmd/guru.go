package cmd

import (
	"azareal/src"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var guruCmd = &cobra.Command{
	Use:     "guru",
	Short:   "print nama guru berdasarkan kode",
	Long:    `mencetak guru berdsarkan argumen yang dimasukkan.`,
	Example: "azareal guru [kode guru]",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		kodeGuru := src.NewKodeGuru(args[0])
		guru, ok := src.SemuaGuru[kodeGuru]
		if !ok {
			log.Fatalln("ERROR: kode guru tidak ditemukan")
		}
		nama := guru.Nama
		fmt.Println(nama)
	},
}

func init() {
	root.AddCommand(guruCmd)
}
