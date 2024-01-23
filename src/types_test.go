package src

import (
	"reflect"
	"testing"
)

func TestPelajaran_init(t *testing.T) {
	type fields struct {
		Nama string
		Kode kodePelaj
		Guru []kodeGuru
	}
	type args struct {
		nama string
		kode kodePelaj
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name: "Nahwu",
			fields: fields{
				Nama: "Nahwu",
				Kode: kodePelaj("NHW"),
				Guru: []kodeGuru{kodeGuru("NZR"), kodeGuru("FTR")},
			},
			args: args{
				nama: "Nahwu",
				kode: "NHW",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pelajaran{
				Nama: tt.fields.Nama,
				Kode: tt.fields.Kode,
				Guru: tt.fields.Guru,
			}
			p.init(tt.args.nama, tt.args.kode)
		})
	}
}

func TestKelas_init(t *testing.T) {
	type fields struct {
		Nama string
		Kode kodeKelas
		Guru []kodeGuru
	}
	type args struct {
		nama string
		kode kodeKelas
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name: "Bahasa Arab",
			fields: fields{
				Nama: "Bahasa Arab",
				Kode: kodeKelas("BA"),
				Guru: []kodeGuru{kodeGuru("NZR"), kodeGuru("FTR")},
			},
			args: args{
				nama: "Bahasa Arab",
				kode: "BA",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &Kelas{
				Nama: tt.fields.Nama,
				Kode: tt.fields.Kode,
				Guru: tt.fields.Guru,
			}
			k.init(tt.args.nama, tt.args.kode)
		})
	}
}

func TestGuru_init(t *testing.T) {
	type fields struct {
		Nama      string
		Kode      kodeGuru
		Kelas     []kodeKelas
		Pelajaran []kodePelaj
		Kepe      map[kodeKelas][]kodePelaj
	}
	type args struct {
		nama string
		kode kodeGuru
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name: "Nahwu",
			fields: fields{
				Nama:      "Nahwu",
				Kode:      kodeGuru("NZR"),
				Kelas:     []kodeKelas{kodeKelas("BA")},
				Pelajaran: []kodePelaj{kodePelaj("NHW")},
				Kepe: map[kodeKelas][]kodePelaj{
					kodeKelas("BA"): {kodePelaj("NHW")},
				},
			},
			args: args{
				nama: "Nahwu",
				kode: "NZR",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Guru{
				Nama:      tt.fields.Nama,
				Kode:      tt.fields.Kode,
				Kelas:     tt.fields.Kelas,
				Pelajaran: tt.fields.Pelajaran,
				Kepe:      tt.fields.Kepe,
			}
			g.init(tt.args.nama, tt.args.kode)
		})
	}
}

func TestGuru_ToString(t *testing.T) {
	type fields struct {
		Nama      string
		Kode      kodeGuru
		Kelas     []kodeKelas
		Pelajaran []kodePelaj
		Kepe      map[kodeKelas][]kodePelaj
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{
			name: "Nahwu",
			fields: fields{
				Nama:      "Nahwu",
				Kode:      kodeGuru("NZR"),
				Kelas:     []kodeKelas{kodeKelas("BA")},
				Pelajaran: []kodePelaj{kodePelaj("NHW")},
				Kepe: map[kodeKelas][]kodePelaj{
					kodeKelas("BA"): {kodePelaj("NHW")},
				},
			},
			want: "Ust. Nahwu:\nKode:      NZR\nKelas:     [BA]\nPelajaran: [NHW]\nKepe:      map[BA:[NHW]]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Guru{
				Nama:      tt.fields.Nama,
				Kode:      tt.fields.Kode,
				Kelas:     tt.fields.Kelas,
				Pelajaran: tt.fields.Pelajaran,
				Kepe:      tt.fields.Kepe,
			}
			if got := g.ToString(); got != tt.want {
				t.Errorf("Guru.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGuru_mapKepe(t *testing.T) {
	type fields struct {
		Nama      string
		Kode      kodeGuru
		Kelas     []kodeKelas
		Pelajaran []kodePelaj
		Kepe      map[kodeKelas][]kodePelaj
	}
	type args struct {
		k kodeKelas
		p kodePelaj
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name: "Bahasa Arab",
			fields: fields{
				Nama:      "Bahasa Arab",
				Kode:      kodeGuru("NZR"),
				Kelas:     []kodeKelas{kodeKelas("BA")},
				Pelajaran: []kodePelaj{kodePelaj("NHW")},
				Kepe: map[kodeKelas][]kodePelaj{
					kodeKelas("BA"): {kodePelaj("NHW")},
				},
			},
			args: args{
				k: kodeKelas("BA"),
				p: kodePelaj("NHW"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Guru{
				Nama:      tt.fields.Nama,
				Kode:      tt.fields.Kode,
				Kelas:     tt.fields.Kelas,
				Pelajaran: tt.fields.Pelajaran,
				Kepe:      tt.fields.Kepe,
			}
			g.mapKepe(tt.args.k, tt.args.p)
		})
	}
}

func TestNewKodeKelas(t *testing.T) {
	type args struct {
		kode string
	}
	tests := []struct {
		name string
		args args
		want kodeKelas
	}{
		// TODO: Add test cases.
		{
			name: "Bahasa Arab",
			args: args{
				kode: "BA",
			},
			want: kodeKelas("BA"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewKodeKelas(tt.args.kode); got != tt.want {
				t.Errorf("NewKodeKelas() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewKodePelaj(t *testing.T) {
	type args struct {
		kode string
	}
	tests := []struct {
		name string
		args args
		want kodePelaj
	}{
		// TODO: Add test cases.
		{
			name: "Bahasa Arab",
			args: args{
				kode: "NHW",
			},
			want: kodePelaj("NHW"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewKodePelaj(tt.args.kode); got != tt.want {
				t.Errorf("NewKodePelaj() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewKodeGuru(t *testing.T) {
	type args struct {
		kode string
	}
	tests := []struct {
		name string
		args args
		want kodeGuru
	}{
		// TODO: Add test cases.
		{
			name: "Bahasa Arab",
			args: args{
				kode: "NZR",
			},
			want: kodeGuru("NZR"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewKodeGuru(tt.args.kode); got != tt.want {
				t.Errorf("NewKodeGuru() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPelajaran(t *testing.T) {
	type args struct {
		nama string
		kode kodePelaj
	}
	tests := []struct {
		name string
		args args
		want *Pelajaran
	}{
		// TODO: Add test cases.
		{
			name: "Bahasa Arab",
			args: args{
				nama: "Bahasa Arab",
				kode: kodePelaj("NHW"),
			},
			want: &Pelajaran{
				Nama: "Bahasa Arab",
				Kode: kodePelaj("NHW"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPelajaran(tt.args.nama, tt.args.kode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPelajaran() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewKelas(t *testing.T) {
	type args struct {
		nama string
		kode kodeKelas
	}
	tests := []struct {
		name string
		args args
		want *Kelas
	}{
		// TODO: Add test cases.
		{
			name: "Bahasa Arab",
			args: args{
				nama: "Bahasa Arab",
				kode: kodeKelas("BA"),
			},
			want: &Kelas{
				Nama: "Bahasa Arab",
				Kode: kodeKelas("BA"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewKelas(tt.args.nama, tt.args.kode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewKelas() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewGuru(t *testing.T) {
	type args struct {
		nama string
		kode kodeGuru
	}
	tests := []struct {
		name string
		args args
		want *Guru
	}{
		// TODO: Add test cases.
		{
			name: "Bahasa Arab",
			args: args{
				nama: "Bahasa Arab",
				kode: kodeGuru("NZR"),
			},
			want: &Guru{
				Nama: "Bahasa Arab",
				Kode: kodeGuru("NZR"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGuru(tt.args.nama, tt.args.kode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGuru() = %v, want %v", got, tt.want)
			}
		})
	}
}
