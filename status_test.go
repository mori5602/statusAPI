package statusAPI_test

import (
	"github.com/mori5602/statusAPI"
	"log"
	"path/filepath"
	"testing"
)

func TestStatusJSON_ReadFile(t *testing.T) {
	status := statusAPI.NewStatus()
	err := status.ReadFile(filepath.Join("testdata", "ajax.Json"))
	if err != nil {
		t.Fatal(err)
	}
	log.Println(status)
}
