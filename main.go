package main

import (
	entities "github.com/beecorrea/midas/src/entities/private"
	"github.com/beecorrea/midas/src/services"
)

func main(){
	// var svc services.Account
	pv := services.GetPrivateServices()
	pv.Account.GetFunds(&entities.Account{})
}