package socks5

import (
	"testing"
)

func TestStaticCredentials(t *testing.T) {
	creds := StaticCredentials{
		"foo": "bar",
		"baz": "",
	}

	if !creds.Valid("foo", "bar") {
		t.Fatal("expect valid")
	}

	if !creds.Valid("baz", "") {
		t.Fatal("expect valid")
	}

	if creds.Valid("foo", "") {
		t.Fatal("expect invalid")
	}
}
