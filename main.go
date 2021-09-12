package main

import (
	"fmt"

	"github.com/beecorrea/midas/src/entities/external"
	"github.com/beecorrea/midas/src/processes"
	"github.com/beecorrea/midas/src/services"
)

func main(){
	fmt.Println("Beginning... Creating services")
	// Init services
	private := services.GetPrivateServices()
	public := services.GetPublicServices(private)
	fmt.Printf("Services created: %v \t\t %v\n", private, public)

	// Init processors
	accman := processes.GetAccManager(public)
	cashier := processes.GetCashier(public)
	fmt.Printf("Processors created: %v \t\t %v\n", accman, cashier)
	
	// Get random account owners
	alice := external.CreateAccOwner()
	bob := external.CreateAccOwner()
	fmt.Printf("Account Owners created: %v \t\t %v\n", alice, bob)
	

	// Create accounts
	accman.OpenAccount(alice)
	accman.OpenAccount(bob)
	fmt.Printf("Accounts created: %v \t\t %v\n", alice.Account, bob.Account)
	
	// Give them some money
	// This is an example where having atomic operations (i.e. services)
	// facilitates direct interventions.
	public.Cashier.Deposit(alice.Account, 500)
	public.Cashier.Deposit(bob.Account, 1000)
	fmt.Printf("Money deposited: alice: %v \t\t bob: %v\n", alice.Account.Funds, bob.Account.Funds)

	cashier.TransferMoney(bob, alice.Account.Id, 500)
	fmt.Printf("Money transferred from A to B: %v \t\t %v", alice.Account.Funds, bob.Account.Funds)
}

