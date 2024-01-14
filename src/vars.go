package src

import (
	"fmt"
)

var (
	SemuaGuru = make(map[kodeGuru]*Guru)

	SemuaPelajaran []Pelajaran
	SemuaKelas     []Kelas
)

var (
	Nzr = NewGuru("Nizar", "NZR")
	Ftr = NewGuru("Fathur", "FTR")
)

func PrintNzr() {
	fmt.Println(Nzr)
}
