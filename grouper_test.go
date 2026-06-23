package mutate

import (
	"net/url"
	"reflect"
	"testing"
)

func TestGroupURLs(t *testing.T) {
	urls := []string{
		"https://example.com/foo",
		"https://example.com/bar?a=1",
		"https://user:pass@example.org/baz?x=1",
	}

	want := URLGrpmap{
		"https://example.com": {
			{
				Path:          "/foo",
				PathWithQuery: "",
			},
			{
				Path:          "/bar",
				PathWithQuery: "/bar?a=1",
			},
		},
		"https://user:pass@example.org": {
			{
				Path:          "/baz",
				PathWithQuery: "/baz?x=1",
			},
		},
	}

	got := groupURLs(urls)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("groupURLs() = %#v, want %#v", got, want)
	}
}

func TestGetBase(t *testing.T) {
	tests := []struct {
		name string
		raw  string
		want string
	}{
		{
			name: "without user info",
			raw:  "https://example.com/foo",
			want: "https://example.com",
		},
		{
			name: "with user info",
			raw:  "https://user:pass@example.com/foo",
			want: "https://user:pass@example.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := url.Parse(tt.raw)
			if err != nil {
				t.Fatal(err)
			}

			if got := getBase(u); got != tt.want {
				t.Fatalf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestGetURLValues(t *testing.T) {
	tests := []struct {
		name string
		raw  string
		want UrlValues
	}{
		{
			name: "path only",
			raw:  "https://example.com/foo",
			want: UrlValues{
				Path: "/foo",
			},
		},
		{
			name: "path with query",
			raw:  "https://example.com/foo?a=1&b=2",
			want: UrlValues{
				Path:          "/foo",
				PathWithQuery: "/foo?a=1&b=2",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := url.Parse(tt.raw)
			if err != nil {
				t.Fatal(err)
			}

			if got := getURLValues(u); !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got %#v, want %#v", got, tt.want)
			}
		})
	}
}
