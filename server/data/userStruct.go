package data

// User represents an user as in database
type User struct {
	UserID    int
	BierName  string
	FirstName string
	LastName  string
	Status    string
	Email     string
	Balance   float32 `json:",string"`
	Phone     string
	MaxDebt   int `json:",string"`
}