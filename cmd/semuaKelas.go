package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"azareal/src"
)

var semuaKelasCmd = &cobra.Command{
	Use:   "semua-kelas",
	Short: "prints all classes",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Semua Kelas:")
		for _, kelas := range src.SemuaKelas {
			fmt.Println(kelas)
		}
	},
}

func init() {
	root.AddCommand(semuaKelasCmd)
}
