package describe

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/go-playground/assert/v2"
)

func randomString(l int) string {
	charPool := []rune("._!,")
	for c := 'a'; c <= 'z'; c++ {
		charPool = append(charPool, c)
	}
	for c := 'A'; c <= 'Z'; c++ {
		charPool = append(charPool, c)
	}
	var res []rune

	for i := 0; i < l; i++ {
		res = append(res, charPool[rand.Intn(len(charPool))])
	}
	return string(res)
}

func Test_parseDescription(t *testing.T) {
	pref := randomString(100)
	postf := randomString(100)
	descr := randomString(30)
	tag := fmt.Sprintf("{{Short description|%s}}", descr)
	all := pref + tag + postf
	res, ok := parseDescription(all)
	assert.Equal(t, true, ok)
	assert.Equal(t, descr, res)
}
