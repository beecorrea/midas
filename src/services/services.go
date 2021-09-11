package services

import "github.com/beecorrea/midas/src/services/private"

type PrivateServices struct {
	Account Account
	Transfer Transfer
}

func GetPrivateServices() PrivateServices {
	return PrivateServices{
		Account: private.NewPrivateAccountSvc(),
		Transfer: private.NewPrivateTransferSvc(),
	}
}

type PublicServices struct {
	Cashier Cashier
	AccountManager AccountManager
}