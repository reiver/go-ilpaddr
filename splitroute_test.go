package ilpaddr

import (
	"testing"
)

func TestSplitRoute(t *testing.T) {

	tests := []struct{
		Route string
		ExpectedSegment string
		ExpectedRemainingRoute string
	}{
		{
		},



		{
			Route:               "g.acme.bob",
			ExpectedSegment:     "g",
			ExpectedRemainingRoute: "acme.bob",
		},
		{
			Route:            "acme.bob",
			ExpectedSegment:  "acme",
			ExpectedRemainingRoute: "bob",
		},
		{
			Route:              "bob",
			ExpectedSegment:    "bob",
			ExpectedRemainingRoute: "",
		},



		{
			Route:                "g.us-fed.ach.0.acmebank.swx0a0.acmecorp.sales.199.~ipr.cdfa5e16-e759-4ba3-88f6-8b9dc83c1868.2",
			ExpectedSegment:      "g",
			ExpectedRemainingRoute: "us-fed.ach.0.acmebank.swx0a0.acmecorp.sales.199.~ipr.cdfa5e16-e759-4ba3-88f6-8b9dc83c1868.2",
		},
		{
			Route:           "us-fed.ach.0.acmebank.swx0a0.acmecorp.sales.199.~ipr.cdfa5e16-e759-4ba3-88f6-8b9dc83c1868.2",
			ExpectedSegment: "us-fed",
			ExpectedRemainingRoute: "ach.0.acmebank.swx0a0.acmecorp.sales.199.~ipr.cdfa5e16-e759-4ba3-88f6-8b9dc83c1868.2",
		},
		{
			Route:              "ach.0.acmebank.swx0a0.acmecorp.sales.199.~ipr.cdfa5e16-e759-4ba3-88f6-8b9dc83c1868.2",
			ExpectedSegment:    "ach",
			ExpectedRemainingRoute: "0.acmebank.swx0a0.acmecorp.sales.199.~ipr.cdfa5e16-e759-4ba3-88f6-8b9dc83c1868.2",
		},
		{
			Route:                "0.acmebank.swx0a0.acmecorp.sales.199.~ipr.cdfa5e16-e759-4ba3-88f6-8b9dc83c1868.2",
			ExpectedSegment:      "0",
			ExpectedRemainingRoute: "acmebank.swx0a0.acmecorp.sales.199.~ipr.cdfa5e16-e759-4ba3-88f6-8b9dc83c1868.2",
		},
		{
			Route:           "acmebank.swx0a0.acmecorp.sales.199.~ipr.cdfa5e16-e759-4ba3-88f6-8b9dc83c1868.2",
			ExpectedSegment: "acmebank",
			ExpectedRemainingRoute:   "swx0a0.acmecorp.sales.199.~ipr.cdfa5e16-e759-4ba3-88f6-8b9dc83c1868.2",
		},
		{
			Route:           "swx0a0.acmecorp.sales.199.~ipr.cdfa5e16-e759-4ba3-88f6-8b9dc83c1868.2",
			ExpectedSegment: "swx0a0",
			ExpectedRemainingRoute: "acmecorp.sales.199.~ipr.cdfa5e16-e759-4ba3-88f6-8b9dc83c1868.2",
		},
		{
			Route:           "acmecorp.sales.199.~ipr.cdfa5e16-e759-4ba3-88f6-8b9dc83c1868.2",
			ExpectedSegment: "acmecorp",
			ExpectedRemainingRoute:   "sales.199.~ipr.cdfa5e16-e759-4ba3-88f6-8b9dc83c1868.2",
		},
		{
			Route:            "sales.199.~ipr.cdfa5e16-e759-4ba3-88f6-8b9dc83c1868.2",
			ExpectedSegment:  "sales",
			ExpectedRemainingRoute: "199.~ipr.cdfa5e16-e759-4ba3-88f6-8b9dc83c1868.2",
		},
		{
			Route:             "199.~ipr.cdfa5e16-e759-4ba3-88f6-8b9dc83c1868.2",
			ExpectedSegment:   "199",
			ExpectedRemainingRoute: "~ipr.cdfa5e16-e759-4ba3-88f6-8b9dc83c1868.2",
		},
		{
			Route:             "~ipr.cdfa5e16-e759-4ba3-88f6-8b9dc83c1868.2",
			ExpectedSegment:   "~ipr",
			ExpectedRemainingRoute: "cdfa5e16-e759-4ba3-88f6-8b9dc83c1868.2",
		},
		{
			Route:           "cdfa5e16-e759-4ba3-88f6-8b9dc83c1868.2",
			ExpectedSegment: "cdfa5e16-e759-4ba3-88f6-8b9dc83c1868",
			ExpectedRemainingRoute:                               "2",
		},
		{
			Route:                "2",
			ExpectedSegment:      "2",
			ExpectedRemainingRoute: "",
		},



		{
			Route:               "g.acme.bob.",
			ExpectedSegment:     "g",
			ExpectedRemainingRoute: "acme.bob.",
		},
		{
			Route:             "acme.bob.",
			ExpectedSegment:   "acme",
			ExpectedRemainingRoute: "bob.",
		},
		{
			Route:             "bob.",
			ExpectedSegment:   "bob",
			ExpectedRemainingRoute: "",
		},



		{
			Route:                 ".g.acme.bob",
			ExpectedSegment:       "",
			ExpectedRemainingRoute: "g.acme.bob",
		},



		{
			Route:                 ".g.acme.bob.",
			ExpectedSegment:       "",
			ExpectedRemainingRoute: "g.acme.bob.",
		},
	}

	for testNumber, test := range tests {

		actualSegment, actualRemainingRoute := splitRoute(test.Route)

		{
			actual := actualSegment
			expected := test.ExpectedSegment

			if expected != actual {
				t.Errorf("For test #%d, the actual 'segment' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("ROUTE: %q", test.Route)
				continue
			}
		}

		{
			actual := actualRemainingRoute
			expected := test.ExpectedRemainingRoute

			if expected != actual {
				t.Errorf("For test #%d, the actual 'remaining-route' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("ROUTE: %q", test.Route)
				continue
			}
		}
	}
}
