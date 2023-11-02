package motor

import (
	"github.com/google/uuid"
)

// Gera um UUIDv4 aleatório
func GenerateRandomUUIDv4() string {
	u, err := uuid.NewRandom()
	if err != nil {
		return err.Error()
	}

	return u.String()
}
