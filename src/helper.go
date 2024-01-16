package src

func regGuruKelas(g *Guru, k *Kelas) {
	g.Kelas = append(g.Kelas, k.Kode)
	k.Guru = append(k.Guru, g.Kode)
}

func regGuruPelaj(g *Guru, p *Pelajaran) {
	g.Pelajaran = append(g.Pelajaran, p.Kode)
	p.Guru = append(p.Guru, g.Kode)
}