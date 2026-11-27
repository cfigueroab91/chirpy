package main

import "testing"

func TestReplaceProfaneWords(t *testing.T) {
	cases := map[string]string{
		"This is a kerfuffle opinion": "This is a **** opinion",
		"SHARBERT and fornax":         "**** and ****",
		"Sharbert! stays":             "Sharbert! stays",
		"nothing to clean":            "nothing to clean",
	}
	for input, want := range cases {
		if got := replaceProfaneWords(input); got != want {
			t.Errorf("replaceProfaneWords(%q) = %q, want %q", input, got, want)
		}
	}
}
