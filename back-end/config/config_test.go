package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_initFromFile(t *testing.T) {
	initFromFile("./testconfig.toml")
	assert.Equal(t, "./data-manager.log", Config.LogConf.FilePath)
}
