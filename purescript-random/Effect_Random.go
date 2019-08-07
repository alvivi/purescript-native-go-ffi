package purescript_random

import (
	. "github.com/purescript-native/go-runtime"
	"math/rand"
	"time"
)

func init() {
	exports := Foreign("Effect.Random")
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	exports["random"] = func() Any {
		return r.Float64()
	}
}
