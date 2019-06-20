package uuid

import (
	"github.com/satori/go.uuid"
	"regexp"
)

func IsValidUUID(uuid string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}

func Generate() string {
	u1 := uuid.Must(uuid.NewV4())
	return u1.String()
}
