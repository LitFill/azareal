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
	Nahwu  = NewPelajaran("Nahwu", "NHW")
	Shorof = NewPelajaran("Shorof", "SRF")
)

var (
	T1a = NewKelas("1 Tsanawiyah A", "T1A")
	T1b = NewKelas("1 Tsanawiyah B", "T1B")
)

var (
	Nzr = NewGuru("Nizar", "NZR")
	Ftr = NewGuru("Fathur", "FTR")
)

func Init() {
	/**** registrasi guru kelas ****/
	regGuruKelas(Nzr, T1a)

	/**** registrasi guru pelajaran ****/
	regGuruPelaj(Nzr, Nahwu)

	/**** mapping pelajaran kelas ****/
	Nzr.mapKepe(T1a.Kode, Nahwu.Kode)
}

func PrintNzr() {
	fmt.Println(Nzr)
}
