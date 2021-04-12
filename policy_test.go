package main

import (
	"testing"
)

// basic case, exception case, fixed caseCompareEvent
// input, expected output, actual output

func TestValidWebhookUrl(t *testing.T) {
	var inUrl string
	var outErr error

	inUrl = "https://foo.bar/entity"
	outErr = ValidUrl(inUrl)
	if outErr != nil {
		t.Error(outErr)
	}

	inUrl = "http://localhost:3000/entity"
	outErr = ValidUrl(inUrl)
	if outErr == nil {
		t.Error("missing valid local domain")
	}

	inUrl = "http://foo.bar.local/entity"
	outErr = ValidUrl(inUrl)
	if outErr == nil {
		t.Error("missing valid local domain")
	}

	inUrl = "http://127.0.0.1/entity"
	outErr = ValidUrl(inUrl)
	if outErr == nil {
		t.Error("missing valid private ip")
	}
}
