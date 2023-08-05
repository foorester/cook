package transport

import (
	"github.com/google/uuid"
)

func toUUID(uuidStr string) (uid uuid.UUID, err error) {
	uid, err = uuid.Parse(uuidStr)
	if err != nil {
		return uid, InvalidID
	}

	return uid, nil
}
