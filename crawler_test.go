package crawler_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestCrawler_Crawl(t *testing.T) {
	srv := testServer(t)
	defer srv.Close()
	srv.URL
}

func testServer(t *testing.T) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var f *os.File

		switch r.RequestURI {
		case "https://blog.marcuskarlsson.com/":
			f, _ = os.Open("./testdata/landing_page.html")
		case "https://blog.marcuskarlsson.com/posts":
			f, _ = os.Open("./testdata/posts.html")
		case "https://blog.marcuskarlsson.com/tags/interesting":
			f, _ = os.Open("./testdata/interesting.html")
		case "https://blog.marcuskarlsson.com/tags/post":
			f, _ = os.Open("./testdata/tags_posts.html")
		case "https://blog.marcuskarlsson.com/posts/my-first-post/":
			f, _ = os.Open("./testdata/first_post.html")
		case "https://blog.marcuskarlsson.com/posts/my-first-post/#disqus_thread":
			f, _ = os.Open("./testdata/disqus_post.html")
		default:
			w.WriteHeader(http.StatusNotFound)
			return
		}

		defer f.Close()

		w.WriteHeader(http.StatusOK)
		io.Copy(w, f)
	}))
}
