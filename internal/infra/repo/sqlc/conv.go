package sqlc

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/foorester/cook/internal/domain/model"
	"github.com/foorester/cook/internal/sys/errors"
)

func toInsertBookParams(b model.Book) (p InsertBookParams, err error) {
	bookID, err := toPgUUID(b.ID.Val())
	if err != nil {
		return p, errors.Wrap(err, "invalid book ID")
	}

	ownerID, err := toPgUUID(b.Owner.ID.Val())
	if err != nil {
		return p, errors.Wrap(err, "invalid owner ID")
	}

	return InsertBookParams{
		ID:          bookID,
		Name:        b.Name,
		Description: b.Description,
		OwnerID:     ownerID,
	}, nil
}

func toPgUUID(uid uuid.UUID) (pgUUID pgtype.UUID, err error) {
	err = pgUUID.Scan(uid)
	if err != nil {
		return pgUUID, errors.Wrap(err, InvalidResourceIDErr.Error())
	}

	return pgUUID, nil
}

// NOTE: For sure there is a more straightforward way to do this exists.
func toUUID(pgUUID pgtype.UUID) (uid uuid.UUID, err error) {
	if !pgUUID.Valid {
		return uuid.Nil, InvalidResourceIDErr
	}

	pgUUIDStr := fmt.Sprintf("%x-%x-%x-%x-%x", pgUUID.Bytes[0:4], pgUUID.Bytes[4:6], pgUUID.Bytes[6:8], pgUUID.Bytes[8:10], pgUUID.Bytes[10:16])

	uid, err = uuid.Parse(pgUUIDStr)
	if err != nil {
		return uuid.Nil, err
	}

	return uid, nil
}

func toID(pgUUID pgtype.UUID) (id model.ID, err error) {
	uid, err := toUUID(pgUUID)
	if err != nil {
		return id, err
	}

	return model.NewID(uid), nil
}

func toTime(pgTime pgtype.Timestamp) (t time.Time, err error) {
	if !pgTime.Valid {
		return t, nil
	}

	return pgTime.Time, nil
}

func toAudit(createdAt, updatedAt pgtype.Timestamp) model.Audit {
	ca, err := toTime(createdAt)
	if err != nil {
		ca = time.Time{}
	}

	ua, err := toTime(updatedAt)
	if err != nil {
		ca = time.Time{}
	}

	return model.Audit{
		CreatedAt: ca,
		UpdatedAt: ua,
	}
}
