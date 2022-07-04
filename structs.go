package main

const (
	noMatch = "🔒"
	match   = "🔓"
	boom    = "💥"
	working = "💚"
)

type Config struct {
	pass    string
	encoded string
	verbose bool
}

var (
	Version = "0.0"
)
