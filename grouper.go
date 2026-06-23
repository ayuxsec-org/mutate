package mutate

import "net/url"

type UrlValues struct {
	PathWithQuery string
	Path          string
}

type URLGrpmap map[string][]UrlValues

func groupURLs(rawURLs []string) URLGrpmap {
	var ugrpMap = make(URLGrpmap)
	for _, rawURL := range rawURLs {
		parsedURL, err := url.Parse(rawURL)
		if err != nil {
			continue // todo
		}
		baseURL := getBase(parsedURL)
		ugrpMap[baseURL] = append(ugrpMap[baseURL], getURLValues(parsedURL))
	}
	return ugrpMap
}

func getBase(u *url.URL) string {
	if u.User != nil {
		return u.Scheme + "://" + u.User.String() + "@" + u.Host
	}
	return u.Scheme + "://" + u.Host
}

func getURLValues(u *url.URL) UrlValues {
	if u.RawQuery != "" {
		return UrlValues{PathWithQuery: u.RequestURI(), Path: u.Path}
	}
	return UrlValues{Path: u.Path}
}
