package generator

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var continents = []string{
	"North America", "Asia", "South America",
	"Europe", "Africa", "Australia",
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyz")
	var b strings.Builder
	for i := 0; i < n; i++ {
		b.WriteRune(letters[rand.Intn(len(letters))])
	}
	return b.String()
}

func randAddress() string {
	letters := []rune("abcdefghijklmnopqrstuvwxyz")
	length := rand.Intn(6) + 15

	var b strings.Builder
	for i := 0; i < length; i++ {
		switch rand.Intn(3) {
		case 0:
			b.WriteByte(byte(rand.Intn(10) + '0'))
		case 1:
			b.WriteRune(letters[rand.Intn(len(letters))])
		default:
			b.WriteByte(' ')
		}
	}
	return b.String()
}

func GenerateLine() string {
	id := rand.Int31()
	name := randString(rand.Intn(6) + 10)
	address := randAddress()
	continent := continents[rand.Intn(len(continents))]

	return fmt.Sprintf("%d,%s,%s,%s", id, name, address, continent)
}
