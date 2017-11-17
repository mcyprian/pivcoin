package uuid

import (
	"testing"
)

func TestGenerateID(t *testing.T) {
	firstId := GenerateID()
	secondId := GenerateID()

	if len(firstId) != len(secondId) {
		t.Errorf("The length of generated ids is not the same.\n")
	}
	if firstId == secondId {
		t.Errorf("Ids are not unique.\n")
	}
}
