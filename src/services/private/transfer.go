package private

import (
	"github.com/beecorrea/midas/src/entities/private"
)

type PrivateTransferService struct {
}

func NewPrivateTransferSvc() *PrivateTransferService {
	return &PrivateTransferService{}
}

func (pts *PrivateTransferService) NewTransfer(from *private.Account, to *private.Account, amount int) *private.Transfer {
	tf := &private.Transfer{
		From: from,
		To: to,
		Amount: amount,
	}
	return tf
}