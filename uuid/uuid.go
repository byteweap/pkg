package uuid

import (
	"strings"

	uuid "github.com/satori/go.uuid"
)

// UUID 不带 - 32位id
func UUID() string {
	id := uuid.NewV4().String()
	return strings.ReplaceAll(id, "-", "")
}
