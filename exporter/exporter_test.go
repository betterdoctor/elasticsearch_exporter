package exporter

import (
	"testing"
	"time"
)

func TestNewExporter(t *testing.T) {
	cases := []struct {
		uri string
		all bool
		ok  bool
	}{
		{uri: "", all: false, ok: false},
		{uri: "localhost", all: false, ok: false},
		{uri: "http://localhost:9200", all: false, ok: true},
		{uri: "http://es.es:9200", all: false, ok: true},
	}

	for _, test := range cases {
		_, err := NewExporter(test.uri, 5*time.Second, test.all)
		if test.ok && err != nil {
			t.Errorf("expected no error but got %v", err)
		}

		if !test.ok && err == nil {
			t.Error("expected error but got nil")
		}
	}
}
