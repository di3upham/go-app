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

	inUrl = "https://foo.bar:80/entity"
	outErr = ValidUrl(inUrl)
	if outErr != nil {
		t.Error(outErr)
	}

	inUrl = "https://142.250.204.36:443"
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

	inUrl = "http://169.254.0.0/16"
	outErr = ValidUrl(inUrl)
	if outErr == nil {
		t.Error("missing valid private ip")
	}

	inUrl = "https://abc.xyz:6060"
	outErr = ValidUrl(inUrl)
	if outErr == nil {
		t.Error("missing valid port")
	}
}
