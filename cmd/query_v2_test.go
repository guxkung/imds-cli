package cmd

import (
	"strings"
	"testing"
)

func TestQueryArrayV2(t *testing.T) {
	TestGetTokenHelper()
	req, buf, err := TestQueryHelperV2("localhost:1338/latest/meta-data")
	if err != nil {
		t.Fatal("test failed at testing query")
	}

	if req.URL.RequestURI() != "/latest/meta-data" {
		t.Fatalf("request URI mismatch %s from definition (/latest/meta-data)", req.URL.RequestURI())
	}
	if strings.Count(buf, "\n") <= 1 {
		t.Fatalf("result has lesser elements than expected %d", strings.Count(buf, "\n"))
	}
}

func TestQueryArrayV2InjectToken(t *testing.T) {
	storeToken("A")
	_, _, err := TestQueryHelperV2("localhost:1338/latest/meta-data")
	if err != nil {
		t.Fatal("test failed at testing query")
	}
}

func TestQueryValueV2(t *testing.T) {
	TestGetTokenHelper()
	req, buf, err := TestQueryHelperV2("localhost:1338/latest/meta-data/ami-id")
	if err != nil {
		t.Fatal("test failed at testing query")
	}

	if req.URL.RequestURI() != "/latest/meta-data/ami-id" {
		t.Fatalf("request URI mismatch %s from definition (/latest/meta-data/ami-id)", req.URL.RequestURI())
	}
	if strings.Count(buf, "\n") > 0 {
		t.Fatalf("result has number of elements differed than expected %d", strings.Count(buf, "\n"))
	}
}

func TestQueryInvalidV2(t *testing.T) {
	TestGetTokenHelper()
	req, buf, err := TestQueryHelperV2("localhost:1338/latest/meta-data/a")
	if err != nil {
		t.Fatal("test failed at testing query")
	}

	if req.URL.RequestURI() != "/latest/meta-data/a" {
		t.Fatalf("request URI mismatch %s from definition (/latest/meta-data/a)", req.URL.RequestURI())
	}
	if strings.Count(buf, "\n") <= 1 {
		t.Fatalf("result has lesser elements than expected %d", strings.Count(buf, "\n"))
	}
}
