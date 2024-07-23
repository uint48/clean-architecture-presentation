package user

// role values
const (
	RegisteredUserRole = 1
	NotRegisteredRole  = 0
)

type User struct {
	ID       string
	Username string
	Password string
	Email    string
	IsActive bool
	Role     int
	Balance  float64
}
