package repository

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/thaynaCaixeta/lucky-admin/internal/domain"
)

type Repository interface {
	SaveGame(numRounds int, closesAt time.Time, createdBy string) (*domain.Game, error)
	CloseConnection()
}

type repo struct {
	conn *sqlx.DB
}

func NewRepository(conn *sqlx.DB) Repository {
	return &repo{
		conn: conn,
	}
}

func (r *repo) CloseConnection() {
	r.conn.Close()
}

func (r *repo) SaveGame(numRounds int, closesAt time.Time, createdBy string) (*domain.Game, error) {
	tx, err := r.conn.Beginx()
	if err != nil {
		log.Printf("Failed to begin transaction: %v", err)
		return nil, NewDatabaseError("failed to begin transaction", err)
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		}
	}()

	// Retrieve the admin id associated with the given username
	stmt, err := tx.PrepareNamed("SELECT id FROM admins WHERE username = :created_by")
	if err != nil {
		log.Printf("Failed to prepare admin query: %v", err)
		tx.Rollback()
		return nil, NewDatabaseError("failed to prepare admin query", err)
	}
	defer stmt.Close()

	var adminId uuid.UUID
	err = stmt.Get(&adminId, map[string]interface{}{"created_by": createdBy})
	if err != nil {
		log.Printf("Admin not found: %v", err)
		tx.Rollback()
		return nil, NewAdminNotFoundError(createdBy)
	}

	gameId, err := uuid.NewRandom()
	if err != nil {
		tx.Rollback()
		return nil, NewUUIDGenerationError(err)
	}

	newGame := domain.NewGame(
		gameId.String(),
		numRounds,
		time.Now().UTC(),
		closesAt.UTC(),
		domain.GameStatus(0),
		adminId.String(),
	)
	_, err = tx.Exec(`
		INSERT INTO games(id, num_rounds, created_at, closes_at, completion_status, created_by)
		VALUES ($1, $2, $3, $4, $5, $6)`,
		newGame.Id,
		newGame.NumRounds,
		newGame.CreatedAt,
		newGame.ClosesAt,
		newGame.Status.String(),
		newGame.CreatedBy,
	)
	if err != nil {
		tx.Rollback()
		log.Printf("Failed to insert game: %v", err)
		return nil, NewDatabaseError("failed to insert game", err)
	}
	if err = tx.Commit(); err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		return nil, NewTransactionCommitError(err)
	}

	// Set the username back before returning to keep the response consistency with the request
	newGame.CreatedBy = createdBy
	return &newGame, nil
}
