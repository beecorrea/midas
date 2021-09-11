package private

import "github.com/google/uuid"

type Account struct {
	Id uuid.UUID
	Owner uuid.UUID
	Funds int
}