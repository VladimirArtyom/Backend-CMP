package model

import "time"

type VoteRequest struct {
	Action       string `json:"action"`
	IDKelompok   string `json:"id_kelompok"`
	NamaKelompok string `json:"nama_kelompok"`
	Kelas        string `json:"kelas"`
}

type GradingRequest struct {
	Action       string `json:"action"`
	IDKelompok   string `json:"id_kelompok"`
	NamaKelompok string `json:"nama_kelompok"`
	KodeReferral string `json:"nama_penilai"`
	Kelas        string `json:"kelas"`
	GradeP1      string `json:"grade_p1"`
	GradeP2      string `json:"grade_p2"`
	GradeP3      string `json:"grade_p3"`
	GradeP4      string `json:"grade_p4"`
}

type DbParticipant struct {
	ID        int
	Email     string
	CreatedAt time.Time
}

type DbKelompok struct {
	ID           int
	Kelas        string
	NamaKelompok string
	CreatedAt    time.Time
}

type DbKelompokFavorit struct {
	ID            int
	IDParticipant int
	IDKelompok    int
	NamaKelompok  string
	Kelas         string
	Skor          int
	CreatedAt     time.Time
}

type DbGrade struct {
	ID           int
	KodeReferral string
	IDKelompok   int
	NamaPenilai  string
	Nilai        int
	Tahun        int
	CreatedAt    string
}

type Referall struct {
	ID           int
	KodeReferall string
	NamaPenilai  string
	CreatedAt    string
}
