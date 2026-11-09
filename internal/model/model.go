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
	IDKelompok   string `json:"id_group"`
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

type DbGroup struct {
	ID        int
	Class     string
	GroupName string
	CreatedAt time.Time
}

type DbVote struct {
	ID            int
	IDParticipant int
	IDGroup       int
	GroupName     string
	Class         string
	Score         int
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
