package screenshot

import (
	"testing"
)

func TestLs(t *testing.T) {
	localServer(":8080")
}

func TestLh(t *testing.T) {
	localhtml()
}
