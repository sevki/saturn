package atlas

import (
	"testing"
)

func TestAtlas(t *testing.T) {
	tests := []struct {
		url     string
		guesses []string
	}{
		{
			"/gmsl",
			[]string{"/gmsl.md", "/gmsl/index.html", "/gmsl/index.md"},
		},
	}
	for _, test := range tests {
		t.Run("urltest"+test.url, func(t *testing.T) {
			got, expecting := guessPossibleNames(test.url), test.guesses
			fail := func() {
				t.Logf("was expecting %q got %q instead", expecting, got)
				t.FailNow()
			}
			if len(got) != len(expecting) {
				fail()
			}
			for i, x := range got {
				if expecting[i] != x {
					fail()
				}
			}
		})
	}
}
