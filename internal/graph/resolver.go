package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver implements GraphQL API logic
type Resolver struct {
	Logger
	*ResolverSettings
}

// Logger defines interface to log messages
type Logger interface {
	Println(v ...interface{})
}
