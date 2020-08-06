package po2json

import (
	"testing"
)

func TestPO2JSON(t *testing.T) {
	err := PO2JSON("messages", "../locale", "po.json")
	if err != nil {
		t.Error(err)
		return
	}
}
