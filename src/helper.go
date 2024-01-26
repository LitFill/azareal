package src

func regGuruKelas(g *Guru, ks ...*Kelas) {
	for _, k := range ks {
		g.Kelas = append(g.Kelas, k.Kode)
		k.Guru = append(k.Guru, g.Kode)
	}
}

func regGuruPelaj(g *Guru, ps ...*Pelajaran) {
	for _, p := range ps {
		g.Pelajaran = append(g.Pelajaran, p.Kode)
		p.Guru = append(p.Guru, g.Kode)
	}
}

