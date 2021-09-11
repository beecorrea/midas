package private

type Transfer struct {
	From   *Account
	To     *Account
	Amount int
}