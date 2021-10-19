package model

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yaoapp/gou"
	"github.com/yaoapp/xiang/config"
	"github.com/yaoapp/xiang/share"
)

func TestLoad(t *testing.T) {
	share.DBConnect(config.Conf.Database)
	gou.Models = make(map[string]*gou.Model)
	Load(config.Conf)
	LoadFrom("not a path", "404.")
	check(t)
}

func check(t *testing.T) {
	keys := []string{}
	for key := range gou.Models {
		keys = append(keys, key)
	}
	assert.Equal(t, 9, len(keys))
}