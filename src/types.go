package src

type (
	kodeKelas string
	kodePelaj string
	kodeGuru  string
)

func NewKodeKelas(kode string) kodeKelas {
	return kodeKelas(kode)
}

func NewKodePelaj(kode string) kodePelaj {
	return kodePelaj(kode)
}

func NewKodeGuru(kode string) kodeGuru {
	return kodeGuru(kode)
}

type Pelajaran struct {
	Nama string
	Kode kodePelaj
	Guru []kodeGuru
}

func (p *Pelajaran) init(nama string, kode kodePelaj) {
	p.Nama = nama
	p.Kode = kode
}

func NewPelajaran(nama string, kode kodePelaj) *Pelajaran {
	pelajaran := &Pelajaran{}
	pelajaran.init(nama, kode)
	SemuaPelajaran[kode] = pelajaran
	return pelajaran
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

func NewKelas(nama string, kode kodeKelas) *Kelas {
	kelas := &Kelas{}
	kelas.init(nama, kode)
	SemuaKelas[kode] = kelas
	return kelas
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
}

func NewGuru(nama string, kode kodeGuru) *Guru {
	guru := &Guru{}
	guru.init(nama, kode)
	SemuaGuru[kode] = guru
	return guru
}
