package private

import "github.com/google/uuid"

type Transfer struct {
	From uuid.UUID
	To uuid.UUID
	Amount int
}