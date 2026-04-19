package generator

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"kafka-pipeline/internal/model"
)

var continents = []string{
	"North America",
	"Asia",
	"South America",
	"Europe",
	"Africa",
	"Australia",
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randLetters(n int) string {
	var b strings.Builder
	for i := 0; i < n; i++ {
		b.WriteByte(byte(rand.Intn(26) + 97))
	}
	return b.String()
}

func randAddress() string {
	return fmt.Sprintf("%d %s %s",
		rand.Intn(9999),
		randLetters(5),
		randLetters(7),
	)
}

func Generate() model.Record {
	return model.Record{
		ID:        rand.Int(),
		Name:      randLetters(10 + rand.Intn(6)),
		Address:   randAddress(),
		Continent: continents[rand.Intn(len(continents))],
	}
}
