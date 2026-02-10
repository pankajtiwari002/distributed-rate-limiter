package config

type RateLimit struct {
	Capacity   int     `yaml:"capacity"`
	RefillRate float64 `yaml:"refill_rate"`
}

type Config struct {
	Mode   string                   `yaml:"mode"`
	Limits map[string]RateLimit     `yaml:"limits"`
}
