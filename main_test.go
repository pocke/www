package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

func TestCheckCertAndKey(t *testing.T) {
	testCases := []struct {
		cert     string
		key      string
		expected bool
	}{
		{cert: "", key: "", expected: false},
		{cert: "cert.pem", key: "", expected: true},
		{cert: "", key: "key.pem", expected: true},
		{cert: "cert.pem", key: "key.pem", expected: false},
	}

	for _, tc := range testCases {
		got := checkCertAndKey(tc.cert, tc.key)
		if got != tc.expected {
			t.Errorf("checkCertAndKey(%q, %q) = %v, want %v", tc.cert, tc.key, got, tc.expected)
		}
	}
}

func TestLoadConfigFile(t *testing.T) {
	testCases := []struct {
		name     string
		content  *string
		expected []string
	}{
		{name: "not exist", content: nil, expected: []string{}},
		{name: "with options", content: ptr("-p 8888\n"), expected: []string{"-p", "8888"}},
		{name: "empty", content: ptr("\n"), expected: []string{""}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Chdir(t.TempDir())
			if tc.content != nil {
				if err := os.WriteFile(".www", []byte(*tc.content), 0644); err != nil {
					t.Fatal(err)
				}
			}

			got, err := loadConfigFile()
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("loadConfigFile() = %#v, want %#v", got, tc.expected)
			}
		})
	}
}

func TestHandler(t *testing.T) {
	t.Chdir(t.TempDir())
	content := "<html><body>hello</body></html>"
	if err := os.WriteFile("index.html", []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	srv := httptest.NewTestServer(t, http.HandlerFunc(handler))

	for _, path := range []string{"/index.html", "/"} {
		resp, err := srv.Client().Get(srv.URL + path)
		if err != nil {
			t.Fatal(err)
		}
		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			t.Fatal(err)
		}

		if string(body) != content {
			t.Errorf("GET %s: body = %q, want %q", path, body, content)
		}
		if got := resp.Header.Get("Cache-Control"); got != "no-store" {
			t.Errorf("GET %s: Cache-Control = %q, want %q", path, got, "no-store")
		}
	}
}

func TestMain_Version(t *testing.T) {
	t.Chdir(t.TempDir())
	if err := Main([]string{"www", "-v"}); err != nil {
		t.Errorf("Main() = %v, want nil", err)
	}
}

func TestMain_CertWithoutKey(t *testing.T) {
	t.Chdir(t.TempDir())
	err := Main([]string{"www", "--cert", "x"})
	if err == nil {
		t.Fatal("Main() = nil, want error")
	}
	expected := "you must specify both --cert and --key"
	if err.Error() != expected {
		t.Errorf("Main() = %q, want %q", err, expected)
	}
}

func TestMain_UnknownFlag(t *testing.T) {
	t.Chdir(t.TempDir())
	if err := Main([]string{"www", "--nonexistent"}); err == nil {
		t.Fatal("Main() = nil, want error")
	}
}

func ptr(s string) *string {
	return &s
}
