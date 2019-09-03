package config

import "testing"

func TestName(t *testing.T) {
	LoadGlobalConfig("config.toml")
}
