package src

import (
	"fmt"
)

var (
	SemuaGuru      = make(map[kodeGuru]*Guru)
	SemuaKelas     = make(map[kodeKelas]*Kelas)
	SemuaPelajaran = make(map[kodePelaj]*Pelajaran)
)

var (
	Akhlaq    = NewPelajaran("Akhlaq", "AKH")
	Arab      = NewPelajaran("Bahasa Arab", "ARB")
	Balaghoh  = NewPelajaran("Balaghoh", "BLG")
	Fiqh      = NewPelajaran("Fiqh", "FQH")
	Ilal      = NewPelajaran("I'lal dan I'rob", "ILL")
	Imla      = NewPelajaran("Imla'", "IML")
	Miftah    = NewPelajaran("Al-Miftah", "MFT")
	Mustholah = NewPelajaran("Mustholah Hadis", "MSH")
	Nahwu     = NewPelajaran("Nahwu", "NHW")
	Shorof    = NewPelajaran("Shorof", "SRF")
	Tajwid    = NewPelajaran("Tajwid", "TJW")
	Tasme     = NewPelajaran("Tasme'", "TSM")
	Tauhid    = NewPelajaran("Tauhid", "THD")
	Ushul     = NewPelajaran("Ushul Fiqh", "USF")
)

var (
	T1a = NewKelas("1 Tsanawiyah A", "T1A")
	T1b = NewKelas("1 Tsanawiyah B", "T1B")
	T1c = NewKelas("1 Tsanawiyah C", "T1C")
	T1d = NewKelas("1 Tsanawiyah D", "T1D")
	T1e = NewKelas("1 Tsanawiyah E", "T1E")
	T1f = NewKelas("1 Tsanawiyah F", "T1F")
	T1g = NewKelas("1 Tsanawiyah G", "T1G")
	T2a = NewKelas("2 Tsanawiyah A", "T2A")
	T2b = NewKelas("2 Tsanawiyah B", "T2B")
	T2c = NewKelas("2 Tsanawiyah C", "T2C")
	T2d = NewKelas("2 Tsanawiyah D", "T2D")
	T2e = NewKelas("2 Tsanawiyah E", "T2E")
	T2f = NewKelas("2 Tsanawiyah F", "T2F")
	T2g = NewKelas("2 Tsanawiyah G", "T2G")
	T2h = NewKelas("2 Tsanawiyah H", "T2H")
	T2i = NewKelas("2 Tsanawiyah I", "T2I")
	T2j = NewKelas("2 Tsanawiyah J", "T2J")
	T3a = NewKelas("3 Tsanawiyah A", "T3A")
	T3b = NewKelas("3 Tsanawiyah B", "T3B")
	T3c = NewKelas("3 Tsanawiyah C", "T3C")
	T3d = NewKelas("3 Tsanawiyah D", "T3D")
	T3e = NewKelas("3 Tsanawiyah E", "T3E")
	T3f = NewKelas("3 Tsanawiyah F", "T3F")
	T3g = NewKelas("3 Tsanawiyah G", "T3G")
	T3h = NewKelas("3 Tsanawiyah H", "T3H")
	T3i = NewKelas("3 Tsanawiyah I", "T3I")
	T3j = NewKelas("3 Tsanawiyah J", "T3J")

	A1a = NewKelas("1 Aliyah A", "A1A")
	A1b = NewKelas("1 Aliyah B", "A1B")
	A1c = NewKelas("1 Aliyah C", "A1C")
	A1d = NewKelas("1 Aliyah D", "A1D")
	A1e = NewKelas("1 Aliyah E", "A1E")
	A1f = NewKelas("1 Aliyah F", "A1F")
	A2a = NewKelas("2 Aliyah A", "A2A")
	A2b = NewKelas("2 Aliyah B", "A2B")
	A2c = NewKelas("2 Aliyah C", "A2C")
	A3a = NewKelas("3 Aliyah A", "A3A")
	A3b = NewKelas("3 Aliyah B", "A3B")
	A3c = NewKelas("3 Aliyah C", "A3C")

	Ist = NewKelas("Isti'dad", "IST")

	Mtp = NewKelas("Mutakhorijin Pagi", "MTP")
	Mts = NewKelas("Mutakhorijin Sore", "MTS")
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
