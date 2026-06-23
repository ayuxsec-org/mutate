package mutate

import (
	"reflect"
	"slices"
	"testing"
)

func TestMutateURLs(t *testing.T) {
	m := New([]string{
		"https://example.com/foo/bar",
		"https://example.com/search?q=test&page=2",
	}, DefaultIdentifier, NewOpts(true, true))

	got := m.MutateURLs()

	want := []string{
		"https://example.com/" + m.Identifier,
		"https://example.com/foo/" + m.Identifier,
		"https://example.com/foo/bar/" + m.Identifier,

		"https://example.com/search?page=2&q=" + m.Identifier,
		"https://example.com/search?page=" + m.Identifier + "&q=test",

		"https://example.com/search/" + m.Identifier,
	}

	for _, expected := range want {
		if !slices.Contains(got, expected) {
			t.Errorf("expected mutated URL %q not found in result", expected)
		}
	}

	if len(got) != len(want) {
		t.Errorf("got %d URLs, want %d\n%v", len(got), len(want), got)
	}
	t.Log(want)
}

func TestMutatePaths(t *testing.T) {
	tests := []struct {
		name string
		path string
		id   string
		want []string
	}{
		{
			name: "nested path",
			path: "/foo/bar",
			id:   "MUTATED",
			want: []string{
				"/MUTATED",
				"/foo/MUTATED",
				"/foo/bar/MUTATED",
			},
		},
		{
			name: "single segment",
			path: "/foo",
			id:   "MUTATED",
			want: []string{
				"/MUTATED",
				"/foo/MUTATED",
			},
		},
		{
			name: "emtpy segment",
			path: "",
			id:   "MUTATED",
			want: []string{
				"/MUTATED",
			},
		},
	}

	for _, tt := range tests {
		// if tt.name != "emtpy segment" {
		// 	continue
		// }
		t.Run(tt.name, func(t *testing.T) {
			got := mutatePaths(tt.path, tt.id)

			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("mutatePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMutatePathWithQuery(t *testing.T) {
	got := mutatePathWithQuery(
		"/search?q=test&page=2",
		"MUTATED",
	)

	want := []string{
		"/search?page=2&q=MUTATED",
		"/search?page=MUTATED&q=test",
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("mutatePathWithQuery() = %v, want %v", got, want)
	}
}
