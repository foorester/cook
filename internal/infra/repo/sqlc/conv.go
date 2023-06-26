package sqlc

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/foorester/cook/internal/domain/model"
	"github.com/foorester/cook/internal/sys/errors"
)

func toInsertBookParams(b model.Book) (p InsertBookParams, err error) {
	bookID, ok := toPgUUID(b.ID.Val())
	if !ok {
		return p, errors.Wrap(InvalidResourceIDErr, "invalid book ID")
	}

	ownerID, ok := toPgUUID(b.Owner.ID)
	if !ok {
		return p, errors.Wrap(InvalidResourceIDErr, "invalid owner ID")
	}

	return InsertBookParams{
		ID:          bookID,
		Name:        b.Name,
		Description: b.Description,
		OwnerID:     ownerID,
	}, nil
}

func toPgUUID(uid uuid.UUID) (pgUUID pgtype.UUID, ok bool) {
	err := pgUUID.Scan(uid)
	if err != nil {
		return pgUUID, false
	}

	return pgUUID, true
}
