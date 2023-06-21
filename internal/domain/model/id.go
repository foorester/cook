package model

import "github.com/google/uuid"

type (
	ID struct {
		val uuid.UUID
	}
)

func (i *ID) GenID(id ...uuid.UUID) {
	if i.val != uuid.Nil {
		return // already has a value assigned
	}

	if len(id) > 0 {
		i.val = id[0] // If value is provided, use it
		return
	}

	i.val = uuid.New()
}

func (i *ID) Val() uuid.UUID {
	return i.val
}

func (i *ID) String() string {
	return i.val.String()
}
