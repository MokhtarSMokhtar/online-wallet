package enums

type UserType string

const (
	// Admin represents an administrative user with elevated permissions
	Admin UserType = "admin"

	// Customer represents a standard user with basic access
	Customer UserType = "customer"

	// Vendor represents a user who can sell products or services
	Vendor UserType = "vendor"
)
