# Design docs

## Entities

### External (belong to the outer part of the system)
- Account owner: 
	* responsible for expressing desire for performing a bank transfer
	* communicates with the accountant to initialize the transfer
	
### Public (reside at the border of the system, interfacing the internal and private parts)
- Cashier: 
	* initializes a bank transfer context (sender acc, receiver acc, amount)
	* executes the transfer to the bank system
	
- Account Manager:
	* opens/closes accounts.
	
### Private (belong to the inner part of the system. accessible only from the border, towards the inside)
- Transfer:
	* holds the context necessary to perform a transfer
	* is passed to the banking system
- Account:
	* stores the total funds of an account
  * communicates with the bookkeeper to log the transactions performed
	
### Internal (belong to the core of the system. interfaces only with the system itself)
- Bookkeeper:
	* logs a transaction as a side effect of the account

---

## Services (which actions the system can do)

### External
- Send money to an account
	
### Public 
- Create account
- Create a transfer
- Check account funds
- Draw money from account
- Deposit money in account

### Private
- Create account
- Create transfer
- Inform account funds
- Add/subtract funds to/from account

### Internal
- Log account transaction

---

## Processes
- Processes are quick and pre-elaborated means of offering a set of services within a specific context and state (i.e. an use-case). They are well-defined,  consume certain rules and can be partially and completely observed by external and public/private/internal entities, respectively.
- The states feed a rule which validates a context. The context denotes a step in the process (i.e. a step in the algorithm), and the state represents a snapshot of the context at a certain point during the process. 
  - E.g. inside the "draw money from one account and deposit it in another account" context from the "Transfer money" process, the sender is A, the receiver is B and the amount is 100.

### Transfer money
- Entities
  - Account Owner
  - Cashier
  - Transfer
  - Account
  - Bookkeeper

### Open account
- Entities
  - Account Owner
  - Cashier
  - Account Manager
  - Transfer
  - Account 

---
## Rules
- Rules are simply what can and cannot be done inside a system.
- They are bound to a context and need a state to validate if something can or cannot be done during a process.
  
### Transfer money
- Does the account exist?
- Is the requester the owner of this account?
- Is this account open?
- Can this amount be drawn from this account?

### Open account
- Does the account exist?
- Is this account open?