package config

type Page struct {
	MaxLimit     int `ini:"maxLimit"`
	MinLimit     int `ini:"minLimit"`
	DefaultLimit int `ini:"defaultLimit"`
}
