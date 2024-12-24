package cmd

import (
	"strings"
	"testing"
)

func TestResolveUrl(t *testing.T) {
	reqUrl := resolveUrl("www.google.com")
	if !strings.HasPrefix(reqUrl.String(), "http://") {
		t.Fatalf("reqUrl %s does not startsWith http://", reqUrl.String())
	}
	reqUrl2 := resolveUrl("https://www.google.com")
	if !strings.HasPrefix(reqUrl2.String(), "http://") {
		t.Fatalf("reqUrl %s does not startsWith http://", reqUrl2.String())
	}
	reqUrl3 := resolveUrl("http://www.google.com")
	if !strings.HasPrefix(reqUrl3.String(), "http://") {
		t.Fatalf("reqUrl %s does not startsWith http://", reqUrl3.String())
	}
}

func TestQueryArrayV1(t *testing.T) {
	req, buf, err := TestQueryHelper("localhost:1338/latest/meta-data")
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

func TestQueryValueV1(t *testing.T) {
	req, buf, err := TestQueryHelper("localhost:1338/latest/meta-data/ami-id")
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

func TestQueryInvalidV1(t *testing.T) {
	req, buf, err := TestQueryHelper("localhost:1338/latest/meta-data/a")
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

func TestWrongHost(t *testing.T) {
	_, buf, err := TestQueryHelper("google.com")
	if err != nil {
		t.Fatal("test failed at testing query")
	}

	if strings.Count(buf, "\n") <= 1 {
		t.Fatalf("result has lesser elements than expected %d", strings.Count(buf, "\n"))
	}
}
