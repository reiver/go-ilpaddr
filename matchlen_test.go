package ilpaddr_test

import (
	"testing"

	"github.com/reiver/go-ilpaddr"
)

func TestMatchLen(t *testing.T) {

	tests := []struct{
		ILPAddr1 string
		ILPAddr2 string
		Expected uint
	}{
		{
		},



		{
			ILPAddr1: "g",
			ILPAddr2: "g",
			Expected: 0,
		},
		{
			ILPAddr1: "private",
			ILPAddr2: "private",
			Expected: 0,
		},
		{
			ILPAddr1: "example",
			ILPAddr2: "example",
			Expected: 0,
		},
		{
			ILPAddr1: "test",
			ILPAddr2: "test",
			Expected: 0,
		},
		{
			ILPAddr1: "test1",
			ILPAddr2: "test1",
			Expected: 0,
		},
		{
			ILPAddr1: "test2",
			ILPAddr2: "test2",
			Expected: 0,
		},
		{
			ILPAddr1: "test3",
			ILPAddr2: "test3",
			Expected: 0,
		},
		{
			ILPAddr1: "local",
			ILPAddr2: "local",
			Expected: 0,
		},
		{
			ILPAddr1: "peer",
			ILPAddr2: "peer",
			Expected: 0,
		},
		{
			ILPAddr1: "self",
			ILPAddr2: "self",
			Expected: 0,
		},



		{
			ILPAddr1: "g.",
			ILPAddr2: "g.",
			Expected: 0,
		},
		{
			ILPAddr1: "private.",
			ILPAddr2: "private.",
			Expected: 0,
		},
		{
			ILPAddr1: "example.",
			ILPAddr2: "example.",
			Expected: 0,
		},
		{
			ILPAddr1: "test.",
			ILPAddr2: "test.",
			Expected: 0,
		},
		{
			ILPAddr1: "test1.",
			ILPAddr2: "test1.",
			Expected: 0,
		},
		{
			ILPAddr1: "test2.",
			ILPAddr2: "test2.",
			Expected: 0,
		},
		{
			ILPAddr1: "test3.",
			ILPAddr2: "test3.",
			Expected: 0,
		},
		{
			ILPAddr1: "local.",
			ILPAddr2: "local.",
			Expected: 0,
		},
		{
			ILPAddr1: "peer.",
			ILPAddr2: "peer.",
			Expected: 0,
		},
		{
			ILPAddr1: "self.",
			ILPAddr2: "self.",
			Expected: 0,
		},



		{
			ILPAddr1: "g.e",
			ILPAddr2: "g.e",
			Expected: 2,
		},
		{
			ILPAddr1: "private.e",
			ILPAddr2: "private.e",
			Expected: 2,
		},
		{
			ILPAddr1: "example.e",
			ILPAddr2: "example.e",
			Expected: 2,
		},
		{
			ILPAddr1: "test.e",
			ILPAddr2: "test.e",
			Expected: 2,
		},
		{
			ILPAddr1: "test1.e",
			ILPAddr2: "test1.e",
			Expected: 2,
		},
		{
			ILPAddr1: "test2.e",
			ILPAddr2: "test2.e",
			Expected: 2,
		},
		{
			ILPAddr1: "test3.e",
			ILPAddr2: "test3.e",
			Expected: 2,
		},
		{
			ILPAddr1: "local.e",
			ILPAddr2: "local.e",
			Expected: 2,
		},
		{
			ILPAddr1: "peer.e",
			ILPAddr2: "peer.e",
			Expected: 2,
		},
		{
			ILPAddr1: "self.e",
			ILPAddr2: "self.e",
			Expected: 2,
		},



		{
			ILPAddr1: "banana.e",
			ILPAddr2: "banana.e",
			Expected: 0,
		},



		{
			ILPAddr1: "g.us-fed.ach.0.acmebank.swx0a0.acmecorp.sales.199.~ipr.cdfa5e16-e759-4ba3-88f6-8b9dc83c1868.2",
			ILPAddr2: "g.us-fed.ach.0.acmebank.swx0a0.acmecorp.sales.199.~ipr.cdfa5e16-e759-4ba3-88f6-8b9dc83c1868.2",
			Expected: 12,
		},



		{
			ILPAddr1: "g.e.",
			ILPAddr2: "g.e.",
			Expected: 0,
		},
		{
			ILPAddr1: "private.e.",
			ILPAddr2: "private.e.",
			Expected: 0,
		},
		{
			ILPAddr1: "example.e.",
			ILPAddr2: "example.e.",
			Expected: 0,
		},
		{
			ILPAddr1: "test.e.",
			ILPAddr2: "test.e.",
			Expected: 0,
		},
		{
			ILPAddr1: "test1.e.",
			ILPAddr2: "test1.e.",
			Expected: 0,
		},
		{
			ILPAddr1: "test2.e.",
			ILPAddr2: "test2.e.",
			Expected: 0,
		},
		{
			ILPAddr1: "test3.e.",
			ILPAddr2: "test3.e.",
			Expected: 0,
		},
		{
			ILPAddr1: "local.e.",
			ILPAddr2: "local.e.",
			Expected: 0,
		},
		{
			ILPAddr1: "peer.e.",
			ILPAddr2: "peer.e.",
			Expected: 0,
		},
		{
			ILPAddr1: "self.e.",
			ILPAddr2: "self.e.",
			Expected: 0,
		},



		{
			ILPAddr1: "g..e",
			ILPAddr2: "g..e",
			Expected: 0,
		},
		{
			ILPAddr1: "private..e",
			ILPAddr2: "private..e",
			Expected: 0,
		},
		{
			ILPAddr1: "example..e",
			ILPAddr2: "example..e",
			Expected: 0,
		},
		{
			ILPAddr1: "test..e",
			ILPAddr2: "test..e",
			Expected: 0,
		},
		{
			ILPAddr1: "test1..e",
			ILPAddr2: "test1..e",
			Expected: 0,
		},
		{
			ILPAddr1: "test2..e",
			ILPAddr2: "test2..e",
			Expected: 0,
		},
		{
			ILPAddr1: "test3..e",
			ILPAddr2: "test3..e",
			Expected: 0,
		},
		{
			ILPAddr1: "local..e",
			ILPAddr2: "local..e",
			Expected: 0,
		},
		{
			ILPAddr1: "peer..e",
			ILPAddr2: "peer..e",
			Expected: 0,
		},
		{
			ILPAddr1: "self..e",
			ILPAddr2: "self..e",
			Expected: 0,
		},



		{
			ILPAddr1: "g.us-fed.ach.0.acmebank.swx0a0.acmecorp.sales.199.~ipr.cdfa5e16-e759-4ba3-88f6-8b9dc83c1868.2",
			ILPAddr2: "g.once.twice.thrice.fource",
			Expected: 1,
		},
		{
			ILPAddr1: "g.us-fed.ach.0.acmebank.swx0a0.acmecorp.sales.199.~ipr.cdfa5e16-e759-4ba3-88f6-8b9dc83c1868.2",
			ILPAddr2: "g.us-fed.apple.banana.cherry",
			Expected: 2,
		},
		{
			ILPAddr1: "g.us-fed.ach.0.acmebank.swx0a0.acmecorp.sales.199.~ipr.cdfa5e16-e759-4ba3-88f6-8b9dc83c1868.2",
			ILPAddr2: "g.us-fed.ach.1.Hello.World",
			Expected: 3,
		},



		{
			ILPAddr1: "g",
			ILPAddr2: "g._a.~b.-c.d.e.ff.g.h--.i.j.k.l.m.n.o.p.q.r.s.t.u.v.w.x.y.z",
			Expected: 0,
		},
		{
			ILPAddr1: "g._a",
			ILPAddr2: "g._a.~b.-c.d.e.ff.g.h--.i.j.k.l.m.n.o.p.q.r.s.t.u.v.w.x.y.z",
			Expected: 2,
		},
		{
			ILPAddr1: "g._a.~b",
			ILPAddr2: "g._a.~b.-c.d.e.ff.g.h--.i.j.k.l.m.n.o.p.q.r.s.t.u.v.w.x.y.z",
			Expected: 3,
		},
		{
			ILPAddr1: "g._a.~b.-c",
			ILPAddr2: "g._a.~b.-c.d.e.ff.g.h--.i.j.k.l.m.n.o.p.q.r.s.t.u.v.w.x.y.z",
			Expected: 4,
		},
		{
			ILPAddr1: "g._a.~b.-c.d",
			ILPAddr2: "g._a.~b.-c.d.e.ff.g.h--.i.j.k.l.m.n.o.p.q.r.s.t.u.v.w.x.y.z",
			Expected: 5,
		},
		{
			ILPAddr1: "g._a.~b.-c.d.e",
			ILPAddr2: "g._a.~b.-c.d.e.ff.g.h--.i.j.k.l.m.n.o.p.q.r.s.t.u.v.w.x.y.z",
			Expected: 6,
		},
		{
			ILPAddr1: "g._a.~b.-c.d.e.ff",
			ILPAddr2: "g._a.~b.-c.d.e.ff.g.h--.i.j.k.l.m.n.o.p.q.r.s.t.u.v.w.x.y.z",
			Expected: 7,
		},
		{
			ILPAddr1: "g._a.~b.-c.d.e.ff.g",
			ILPAddr2: "g._a.~b.-c.d.e.ff.g.h--.i.j.k.l.m.n.o.p.q.r.s.t.u.v.w.x.y.z",
			Expected: 8,
		},
		{
			ILPAddr1: "g._a.~b.-c.d.e.ff.g.h--",
			ILPAddr2: "g._a.~b.-c.d.e.ff.g.h--.i.j.k.l.m.n.o.p.q.r.s.t.u.v.w.x.y.z",
			Expected: 9,
		},
		{
			ILPAddr1: "g._a.~b.-c.d.e.ff.g.h--.i",
			ILPAddr2: "g._a.~b.-c.d.e.ff.g.h--.i.j.k.l.m.n.o.p.q.r.s.t.u.v.w.x.y.z",
			Expected: 10,
		},

		{
			ILPAddr1: "g._a.~b.-c.d.e.ff.g.h--.i.j.k.l.m.n.o.p.q.r.s.t.u.v.w.x.y",
			ILPAddr2: "g._a.~b.-c.d.e.ff.g.h--.i.j.k.l.m.n.o.p.q.r.s.t.u.v.w.x.y.z",
			Expected: 26,
		},
		{
			ILPAddr1: "g._a.~b.-c.d.e.ff.g.h--.i.j.k.l.m.n.o.p.q.r.s.t.u.v.w.x.y.z",
			ILPAddr2: "g._a.~b.-c.d.e.ff.g.h--.i.j.k.l.m.n.o.p.q.r.s.t.u.v.w.x.y.z",
			Expected: 27,
		},



		{
			ILPAddr1: "g._a.~b.-c.d.e.ff.g.h--.i.j.k.l.m.n.o.p.q.r.s.t.u.v.w.x.y.z",
			ILPAddr2: "g._a.~b.-c.d.e.ff.g.h--.i.j.k.l.m.n.o.p.q.r.s.t.u.v.w.x.y",
			Expected: 26,
		},
	}

	for testNumber, test := range tests {

		actual := ilpaddr.MatchLen(test.ILPAddr1, test.ILPAddr2)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual prefix-match-length is not what was expected.", testNumber)
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			t.Logf("ILP-ADDRESS-1: %q", test.ILPAddr1)
			t.Logf("ILP-ADDRESS-2: %q", test.ILPAddr2)
			continue
		}
	}
}
