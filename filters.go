package crawler

import "net/url"

type FilterFunc func(*url.URL) (ok bool)

func Unique() FilterFunc {
	been := make(map[string]bool)
	return func(u *url.URL) bool {
		_, exist := been[u.String()]
		if exist {
			return false
		}

		been[u.String()] = true

		return true
	}
}

func SameHost(base *url.URL) FilterFunc {
	return func(u *url.URL) bool {
		return base.Hostname() == u.Hostname()
	}
}
