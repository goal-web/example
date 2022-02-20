package tests

import (
	"fmt"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCacheFactory(t *testing.T) {
	app := initApp()

	cacheFactory := app.Get("cache").(contracts.CacheFactory)

	value := utils.RandStr(50)
	err := cacheFactory.Store().Forever("a", value)
	assert.Nil(t, err, err)
	valueFromCache := cacheFactory.Store().Get("a")
	fmt.Println(valueFromCache)
	assert.True(t, valueFromCache == value)
}
