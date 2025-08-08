package ilpaddr_test

import (
	"testing"

	"github.com/reiver/go-ilpaddr"
)

func TestHasScheme(t *testing.T) {

	tests := []struct{
		ILPAddr string
		Scheme string
		Expected bool
	}{
		{
		},



		{
			ILPAddr: "",
			Scheme: ilpaddr.SchemeGlobal,
			Expected: false,
		},
		{
			ILPAddr: "g.acme.bob",
			Scheme: "",
			Expected: false,
		},



		{
			ILPAddr: "g.acme.bob",
			Scheme: ilpaddr.SchemeGlobal,
			Expected: true,
		},
		{
			ILPAddr: "g.acme.bob",
			Scheme: ilpaddr.SchemeLocal,
			Expected: false,
		},



		{
			ILPAddr: "local.acme.bob",
			Scheme: ilpaddr.SchemeGlobal,
			Expected: false,
		},
		{
			ILPAddr: "local.acme.bob",
			Scheme: ilpaddr.SchemeLocal,
			Expected: true,
		},
	}

	for testNumber, test := range tests {

		actual := ilpaddr.HasScheme(test.ILPAddr, test.Scheme)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual has-scheme result is not what was expected.", testNumber)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("ILP-ADDR: %q", test.ILPAddr)
			t.Logf("SCHEME: %q", test.Scheme)
			continue
		}
	}
}
