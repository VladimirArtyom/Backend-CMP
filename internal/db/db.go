package db

import (
	"compro-exh/internal/model"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DbPool *pgxpool.Pool

func Init() {
	var ctx context.Context = context.Background()
	var database_url string = "postgres://postgres:password@localhost:5432/compro_26_27_odd?sslmode=disable"
	var err error
	DbPool, err = pgxpool.New(ctx, database_url)
	if err != nil {
		log.Fatalf("Unable to connect to database : %v", err)
		return
	}
	err = DbPool.Ping(ctx)
	if err != nil {
		log.Fatalf("Unable to ping the database: %v\n", err)
		return
	}
	fmt.Println("Successfully connected to PostGreSQL")
}

func InsertParticipant(ctx context.Context, email string) error {
	// To be added timestamps
	var time_now string = time.Now().Format(time.RFC3339)
	var query string = "INSERT INTO participant(email, created_at) VALUES ($1, $2)"
	_, err := DbPool.Exec(ctx, query, email, time_now)
	if err != nil {
		return err
	}
	return nil

}

func GetParticipant(ctx context.Context, email string) (*model.DbParticipant, error) {
	var participant *model.DbParticipant = &model.DbParticipant{}
	var query string = "SELECT * FROM participant where email=$1"
	err := DbPool.QueryRow(ctx, query, email).Scan(&participant.ID, &participant.Email, &participant.CreatedAt)
	if err != nil {
		return nil, err
	}

	return participant, nil

}

func DeleteParticipant(ctx context.Context, email string) error {

	var query string = "DELETE FROM participant WHERE email=$1"
	_, err := DbPool.Exec(ctx, query, email)
	if err != nil {
		return err
	}
	return nil
}

func InsertKelompok(ctx context.Context, nama string, kelas string) error {
	var query string = "INSERT INTO kelompok (nama_kelompok, kelas) VALUES ($1, $2)"
	_, err := DbPool.Exec(ctx, query, nama, kelas)
	if err != nil {
		return err
	}
	return nil

}

func DeleteKelompok(ctx context.Context, id int) error {

	var query string = "DELETE FROM KELOMPOK WHERE id=$1"
	_, err := DbPool.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteAllKelompk(ctx context.Context) error {
	var query string = "DELETE FROM kelompok"
	_, err := DbPool.Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func InsertKelompokFavorit(ctx context.Context) error {
	var query string = "INSERT INTO kelompok_favorit() VALUES ()"
	_, err := DbPool.Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func GetKelompokFavorit(ctx context.Context, id int) error {
	var kel *model.DbKelompokFavorit = &model.DbKelompokFavorit{}
	var query string = "SELECT id, id_participant, id_kelompok, kelas, nama_kelompok, skor, created_at FROM kelompok_favorit"
	err := DbPool.QueryRow(ctx, query, id).Scan(&kel.ID, &kel.IDParticipant, &kel.IDKelompok, &kel.Kelas, &kel.NamaKelompok, &kel.Skor, &kel.CreatedAt)

	if err != nil {
		return err
	}
	return nil

}

func GetGrade(ctx context.Context, id int) error {
	var grade *model.DbGrade = &model.DbGrade{}
	var query string = "SELECT id, kode_referral, id_kelompok, nama_penilai, nilai, tahun, created_at FROM penilaian WHERE id=$1"
	err := DbPool.QueryRow(ctx, query, id).Scan(&grade.ID, &grade.KodeReferral, &grade.IDKelompok, &grade.NamaPenilai, &grade.NamaPenilai, &grade.Nilai, &grade.Tahun, &grade.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}

func GetGrades(ctx context.Context) ([]model.DbGrade, error) {
	var grades []model.DbGrade = []model.DbGrade{}
	var query string = "SELECT id, kode_referral, id_kelompok, nama_penilai, nilai, tahun, created_at FROM penilaian"
	rows, err := DbPool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var grade model.DbGrade = model.DbGrade{}
		err := rows.Scan(&grade.ID,
			&grade.KodeReferral, &grade.IDKelompok,
			&grade.NamaPenilai, &grade.Nilai, &grade.Tahun,
			&grade.CreatedAt)
		if err != nil {
			return nil, err
		}
		grades = append(grades, grade)
	}
	return grades, nil
}

func InsertGrade(ctx context.Context, kode_referral, id_kelompok, nama_penilai string, nilai, tahun int, created_at string) error {

	var query string = `INSERT INTO penilaian 
		(kode_referral, id_kelompok, nama_penilai, nilai, tahun, created_at) VALUES
		($1, $2, $3, $4, $5, $6)`
	_, err := DbPool.Exec(ctx, query,
		kode_referral, id_kelompok,
		nama_penilai, nilai,
		tahun, created_at)

	if err != nil {
		return err
	}

	return nil
}

// Insert participant
// REMOVE PARTICIPANT

// Insert kelompok (Pre generated)
// delete a group
// Delete all groups
// Delete list of groups

// Insert referral judges and lectures have different code referall
// delete a referall --
// delete all referall --
// delete a list of referell --

// insert group favorite
func Insert_Vote_to_A_Group() {

}
