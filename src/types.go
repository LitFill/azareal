package src

import (
	// "azareal/utils"
	"fmt"
	"log"
	"slices"
)

type (
	kodeKelas string
	kodePelaj string
	kodeGuru  string
)

type Pelajaran struct {
	Nama string
	Kode kodePelaj
	Guru []kodeGuru
}

func (p *Pelajaran) init(nama string, kode kodePelaj) {
	p.Nama = nama
	p.Kode = kode
}

type Kelas struct {
	Nama string
	Kode kodeKelas
	Guru []kodeGuru
}

func (k *Kelas) init(nama string, kode kodeKelas) {
	k.Nama = nama
	k.Kode = kode
}

type Guru struct {
	Nama      string
	Kode      kodeGuru
	Kelas     []kodeKelas
	Pelajaran []kodePelaj
	Kepe      map[kodeKelas][]kodePelaj
}

func (g *Guru) init(nama string, kode kodeGuru) {
	g.Nama = nama
	g.Kode = kode
	g.Kepe = make(map[kodeKelas][]kodePelaj)
}

func (g *Guru) ToString() string {
	str := fmt.Sprintf(`Ust. %s:
	Kode:      %s
	Kelas:     %s
	Pelajaran: %s
	Kepe:      %v`,
		g.Nama, g.Kode, g.Kelas, g.Pelajaran, g.Kepe)
	return str
}

func (g *Guru) mapKepe(k kodeKelas, p kodePelaj) {
	if !slices.Contains(g.Kelas, k) {
		// log.Fatalf("ERROR: Ust. %s tidak mengajar di kelas %s\n", g.Nama, k)
		log.Panicf("ERROR: Ust. %s tidak mengajar di kelas %s\n", g.Nama, k)
		return
	}

	if !slices.Contains(g.Pelajaran, p) {
		// log.Fatalf("ERROR: Ust. %s tidak mengajar pelajaran %s", g.Nama, p)
		log.Panicf("ERROR: Ust. %s tidak mengajar pelajaran %s", g.Nama, p)
		return
	}

	g.Kepe[k] = append(g.Kepe[k], p)
}

/***** constructors *****/

func NewKodeKelas(kode string) kodeKelas {
	return kodeKelas(kode)
}

func NewKodePelaj(kode string) kodePelaj {
	return kodePelaj(kode)
}

func NewKodeGuru(kode string) kodeGuru {
	return kodeGuru(kode)
}

func NewPelajaran(nama string, kode kodePelaj) *Pelajaran {
	pelajaran := &Pelajaran{}
	pelajaran.init(nama, kode)
	SemuaPelajaran[kode] = pelajaran
	return pelajaran
}

func NewKelas(nama string, kode kodeKelas) *Kelas {
	kelas := &Kelas{}
	kelas.init(nama, kode)
	SemuaKelas[kode] = kelas
	return kelas
}
func NewGuru(nama string, kode kodeGuru) *Guru {
	guru := &Guru{}
	guru.init(nama, kode)
	SemuaGuru[kode] = guru
	return guru
}
