package src

import (
	"fmt"
)

var (
	SemuaPelajaran = make(map[kodePelaj]*Pelajaran)
	SemuaKelas     = make(map[kodeKelas]*Kelas)
	SemuaGuru      = make(map[kodeGuru]*Guru)
)

var (
	Nzr = NewGuru("Nizar", "NZR")
	Ftr = NewGuru("Fathur", "FTR")
)

func PrintNzr() {
	fmt.Println(Nzr)
}
