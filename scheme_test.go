package ilpaddr_test

import (
	"testing"

	"github.com/reiver/go-ilpaddr"
)

func TestScheme(t *testing.T) {

	tests := []struct{
		ILPAddr string
		Expected string
	}{
		{
		},



		{
			ILPAddr: "",
			Expected: "",
		},



		{
			ILPAddr: "g",
			Expected: "",
		},
		{
			ILPAddr: "private",
			Expected: "",
		},
		{
			ILPAddr: "example",
			Expected: "",
		},
		{
			ILPAddr: "test",
			Expected: "",
		},
		{
			ILPAddr: "test1",
			Expected: "",
		},
		{
			ILPAddr: "test2",
			Expected: "",
		},
		{
			ILPAddr: "test3",
			Expected: "",
		},
		{
			ILPAddr: "local",
			Expected: "",
		},
		{
			ILPAddr: "peer",
			Expected: "",
		},
		{
			ILPAddr: "self",
			Expected: "",
		},



		{
			ILPAddr: "g.",
			Expected: "",
		},
		{
			ILPAddr: "private.",
			Expected: "",
		},
		{
			ILPAddr: "example.",
			Expected: "",
		},
		{
			ILPAddr: "test.",
			Expected: "",
		},
		{
			ILPAddr: "test1.",
			Expected: "",
		},
		{
			ILPAddr: "test2.",
			Expected: "",
		},
		{
			ILPAddr: "test3.",
			Expected: "",
		},
		{
			ILPAddr: "local.",
			Expected: "",
		},
		{
			ILPAddr: "peer.",
			Expected: "",
		},
		{
			ILPAddr: "self.",
			Expected: "",
		},



		{
			ILPAddr: "g.e",
			Expected: ilpaddr.SchemeGlobal,
		},
		{
			ILPAddr: "private.e",
			Expected: ilpaddr.SchemePrivate,
		},
		{
			ILPAddr: "example.e",
			Expected: ilpaddr.SchemeExample,
		},
		{
			ILPAddr: "test.e",
			Expected: ilpaddr.SchemeTest,
		},
		{
			ILPAddr: "test1.e",
			Expected: ilpaddr.SchemeTest1,
		},
		{
			ILPAddr: "test2.e",
			Expected: ilpaddr.SchemeTest2,
		},
		{
			ILPAddr: "test3.e",
			Expected: ilpaddr.SchemeTest3,
		},
		{
			ILPAddr: "local.e",
			Expected: ilpaddr.SchemeLocal,
		},
		{
			ILPAddr: "peer.e",
			Expected: ilpaddr.SchemePeer,
		},
		{
			ILPAddr: "self.e",
			Expected: ilpaddr.SchemeSelf,
		},



		{
			ILPAddr: "g.banana",
			Expected: ilpaddr.SchemeGlobal,
		},
		{
			ILPAddr: "private.banana",
			Expected: ilpaddr.SchemePrivate,
		},
		{
			ILPAddr: "example.banana",
			Expected: ilpaddr.SchemeExample,
		},
		{
			ILPAddr: "test.banana",
			Expected: ilpaddr.SchemeTest,
		},
		{
			ILPAddr: "test1.banana",
			Expected: ilpaddr.SchemeTest1,
		},
		{
			ILPAddr: "test2.banana",
			Expected: ilpaddr.SchemeTest2,
		},
		{
			ILPAddr: "test3.banana",
			Expected: ilpaddr.SchemeTest3,
		},
		{
			ILPAddr: "local.banana",
			Expected: ilpaddr.SchemeLocal,
		},
		{
			ILPAddr: "peer.banana",
			Expected: ilpaddr.SchemePeer,
		},
		{
			ILPAddr: "self.banana",
			Expected: ilpaddr.SchemeSelf,
		},



		{
			ILPAddr:  "g.acme.bob",
			Expected: ilpaddr.SchemeGlobal,
		},



		{
			ILPAddr: "g.once",
			Expected: ilpaddr.SchemeGlobal,
		},
		{
			ILPAddr: "g.once.twice",
			Expected: ilpaddr.SchemeGlobal,
		},
		{
			ILPAddr: "g.once.twice.thrice",
			Expected: ilpaddr.SchemeGlobal,
		},
		{
			ILPAddr: "g.once.twice.thrice.fource",
			Expected: ilpaddr.SchemeGlobal,
		},



		{
			ILPAddr: "local.acme.bob",
			Expected: ilpaddr.SchemeLocal,
		},
	}

	for testNumber, test := range tests {

		actual := ilpaddr.Scheme(test.ILPAddr)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual scheme is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("ILP-ADDR: %q", test.ILPAddr)
			continue
		}
	}
}
