package langtag

import "testing"

func TestNormalize(t *testing.T) {
	cases := map[string]string{"en_us": "en-US", "PT-br": "pt-BR", " de ": "de"}
	for in, want := range cases {
		got, err := Normalize(in)
		if err != nil || got != want {
			t.Errorf("Normalize(%q) = %q, %v; want %q", in, got, err, want)
		}
	}
}

func TestNormalizeEmpty(t *testing.T) {
	if _, err := Normalize("  "); err != ErrEmpty {
		t.Fatalf("want ErrEmpty, got %v", err)
	}
}

func TestNormalizeInvalid(t *testing.T) {
	if _, err := Normalize("not a tag!"); err == nil {
		t.Fatal("want an error for an invalid tag")
	}
}

func TestBest(t *testing.T) {
	if got := Best("fr-CH, fr;q=0.9, en;q=0.8", "en", "fr"); got != "fr" {
		t.Errorf("Best = %q, want fr", got)
	}
	if got := Best("ja", "en", "fr"); got != "en" {
		t.Errorf("Best fallback = %q, want en", got)
	}
	if got := Best("en"); got != "" {
		t.Errorf("Best with no supported = %q", got)
	}
}

func TestBestInvalidHeader(t *testing.T) {
	if got := Best(";;;q=abc", "en", "fr"); got != "en" {
		t.Errorf("Best with invalid header = %q, want en", got)
	}
}
