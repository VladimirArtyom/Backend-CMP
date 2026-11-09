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

func InsertGroup(ctx context.Context, nama string, class string) error {
	var query string = "INSERT INTO t_group (group_name, class) VALUES ($1, $2)"
	_, err := DbPool.Exec(ctx, query, nama, class)
	if err != nil {
		return err
	}
	return nil

}

func GetGroup(ctx context.Context, name string) (*model.DbGroup, error) {
	var group model.DbGroup = model.DbGroup{}
	var query string = "SELECT * FROM t_group WHERE group_name=$1"
	err := DbPool.QueryRow(ctx, query, name).Scan(&group.ID,
		&group.GroupName, &group.Class, &group.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &group, nil

}

func GetGroups(ctx context.Context) ([]model.DbGroup, error) {
	var groups []model.DbGroup = []model.DbGroup{}
	var query = "SELECT * FROM t_group"
	rows, err := DbPool.Query(ctx, query)
	if err != nil {

		return nil, err
	}

	for rows.Next() {
		var group model.DbGroup = model.DbGroup{}
		err = rows.Scan(&group.ID, &group.GroupName,
			&group.Class, &group.CreatedAt)
		if err != nil {
			return nil, err
		}

		groups = append(groups, group)
	}

	return groups, nil
}

func DeleteGroup(ctx context.Context, groupName string) error {

	var query string = "DELETE FROM t_group WHERE group_name=$1"
	_, err := DbPool.Exec(ctx, query, groupName)
	if err != nil {
		return err
	}
	return nil
}

func DeleteAllGroup(ctx context.Context) error {
	var query string = "DELETE FROM t_group"
	_, err := DbPool.Exec(ctx, query)
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
func InsertVote(ctx context.Context, id_participant, id_group int, class string, group_name string, score int) error {
	var query = `INSERT INTO 
	votes (id_participant, id_group, class, group_name, score) 
	VALUES ($1, $2, $3, $4, $5)`

	_, err := DbPool.Exec(ctx, query, id_participant, id_group, class, group_name, score)
	if err != nil {
		return err
	}

	return nil

}

func DeleteVote(ctx context.Context, id_participant, id_group int) error {
	var query = "DELETE FROM votes WHERE id_participant=$1 AND id_group=$2"
	_, err := DbPool.Exec(ctx, query, id_participant, id_group)
	if err != nil {
		return err
	}

	return nil

}

func GetVoteById(ctx context.Context, id_participant, id_group int) (*model.DbVote, error) {
	var vote model.DbVote = model.DbVote{}
	var query = "SELECT * FROM votes WHERE id_participant=$1 AND id_group=$2"
	err := DbPool.QueryRow(ctx, query, id_participant, id_group).Scan(&vote.ID,
		&vote.IDParticipant, &vote.IDGroup,
		&vote.GroupName, &vote.Score, &vote.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &vote, nil

}

func GetVotes(ctx context.Context) ([]model.DbVote, error) {

	var query = "SELECT * FROM votes"
	var votes []model.DbVote = []model.DbVote{}

	rows, err := DbPool.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var vote model.DbVote = model.DbVote{}
		err = rows.Scan(&vote.ID, &vote.IDParticipant, &vote.IDGroup,
			&vote.Class, &vote.GroupName, &vote.Score, &vote.CreatedAt)

		if err != nil {
			return nil, err
		}

		votes = append(votes, vote)
	}

	return votes, nil
}
