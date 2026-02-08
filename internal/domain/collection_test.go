package domain

import "testing"

func TestCompileDomain(t *testing.T) {
	collection := Collection{
		Name: "Sample",
		Requests: []RequestSpec{
			{
				Name:   "Get users",
				Method: "GET",
				URL:    "https://api.example.com/users",
				Headers: []Header{
					{Name: "Accept", Value: "application/json"},
				},
				Body: BodySpec{
					JSON: map[string]any{"active": true},
				},
				Assertions: AssertionsSpec{
					Status:          200,
					MaxMilliseconds: 500,
					JSONPathExists:  []string{"$.data"},
				},
				Extract: []ExtractSpec{
					{JSONPath: "$.data[0].id", Var: "USER_ID"},
				},
			},
		},
	}

	if collection.Requests[0].Method == "" {
		t.Fatal("expected method to be set")
	}
}
