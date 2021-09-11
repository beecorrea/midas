package services

import (
	"github.com/beecorrea/midas/src/services/private"
	"github.com/beecorrea/midas/src/services/public"
)

func GetPrivateServices() private.PrivateServices {
	return private.PrivateServices{
		Account: private.NewPrivateAccountSvc(),
		Transfer: private.NewPrivateTransferSvc(),
	}
}

func GetPublicServices(prv private.PrivateServices) public.PublicServices{
	return public.PublicServices{
		Cashier: public.NewPublicCashierService(prv),
		AccountManager: public.NewPublicAccManagerSvc(prv),
	}
}