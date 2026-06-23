package mutate

import (
	"net/url"
	"strings"
)

type Mutator struct {
	RawURLs    []string
	Identifier string
	Opts       Options
}

type Options struct {
	MutatePaths bool
	MutateKeys  bool
}

// todo (optional): change to api based arguments
func NewOpts(paths bool, keys bool) Options {
	return Options{
		MutatePaths: paths,
		MutateKeys:  keys,
	}
}

func New(rawURLs []string, id string, opts Options) *Mutator {
	return &Mutator{
		RawURLs:    rawURLs,
		Identifier: id,
		Opts:       opts,
	}
}

func (m *Mutator) MutateURLs() []string {
	ugrpMap := groupURLs(m.RawURLs)
	var mutatedURLs []string

	for url, values := range ugrpMap {
		for _, value := range values {
			if m.Opts.MutateKeys {
				if value.PathWithQuery != "" {
					mutatedURLs = append(mutatedURLs, prefixString(url, mutatePathWithQuery(value.PathWithQuery, m.Identifier))...)
				}
			}
			if m.Opts.MutatePaths {
				mutatedURLs = append(mutatedURLs, prefixString(url, mutatePaths(value.Path, m.Identifier))...)
			}
		}
	}

	return dedupeSlice(mutatedURLs)
}

func mutatePathWithQuery(path string, id string) []string {
	parsedPath, _ := url.Parse(path) // todo
	var mutatedPaths []string
	params := parsedPath.Query()

	for key, value := range params {
		params[key] = []string{id}
		parsedPath.RawQuery = params.Encode()
		mutatedPath := parsedPath.String()
		params[key] = value // revert to original value
		mutatedPaths = append(mutatedPaths, mutatedPath)
	}
	return mutatedPaths
}

func mutatePaths(path string, id string) []string {
	var mutatedPaths []string

	path = strings.Trim(path, "/")
	if path == "" {
		return []string{path + "/" + id}
	}

	segments := strings.Split(strings.Trim(path, "/"), "/")

	for i := 0; i <= len(segments); i++ {
		joinedPath := strings.Join(segments[:i], "/")
		mutatedPaths = append(mutatedPaths, "/"+joinedPath+"/"+id)
	}

	return mutatedPaths
}
