package tests

import (
	"github.com/goal-web/contracts"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRedisBloomFilter(t *testing.T) {
	var app = initApp()

	app.Call(func(factory contracts.BloomFactory) {
		var filter = factory.Filter("users")

		filter.AddString("goal")

		assert.True(t, filter.TestString("goal"))
	})
}

func BenchmarkRedisBoolFilter(b *testing.B) {
	var app = initApp()

	app.Call(func(factory contracts.BloomFactory) {
		var filter = factory.Filter("users")

		for i := 0; i < b.N; i++ {
			filter.AddString("goal")
		}

	})

}
