package config

// Database ...
type Database struct {
	Uri      string
	Name     string
	TestName string
}

// ENV ...
type ENV struct {
	// AppPort ...
	AppPort string

	// Database ...
	Database Database
}
