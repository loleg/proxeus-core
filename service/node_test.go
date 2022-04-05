package service

import (
	"testing"
)

func TestList(t *testing.T) {

	nodeService := NewNodeService(nil)

	if nodeService.List("mailsender") == nil {
		t.Errorf("Missing mailsender")
		// t.Fail()
	}
	if nodeService.List("priceretriever") == nil {
		t.Errorf("Missing priceretriever")
		// t.Fail()
	}
}
