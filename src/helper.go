package src

func regGuruKelas(g *Guru, k *Kelas) {
	g.Kelas = append(g.Kelas, k.Kode)
	k.Guru = append(k.Guru, g.Kode)
}
