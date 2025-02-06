package fedifinger

import (
	"io"
	"net/http"
	liburl "net/url"

	"github.com/reiver/go-erorr"
	"github.com/reiver/go-errhttp"
)

const (
	errNilHTTPResponse     = erorr.Error("fedifinger: nil http-response")
	errNilHTTPResponseBody = erorr.Error("fedifinger: nil http-response-body")
	errNilParsedURL        = erorr.Error("fedifinger: nil parsed-url")
)

// Get gets the activity-JSON for a **Fediverse-ID**.
//
// For example:
//
//	bytes, err := fedifinger.Get("@reiver@mastodon.social")
func Get(fediverseID string) ([]byte, error) {

	var activityURL string
	{
		var err error
		activityURL, err = Resolve(fediverseID)
		if nil != err {
			var nada []byte
			return nada, err
		}
		if "" == activityURL {
			var nada []byte
			return nada, erorr.Errorf("fedifinger: fediverse-id %q resolved to an empty URL", fediverseID)
		}
	}

	var httprequest http.Request
	{
		var urloc *liburl.URL
		{
			var err error
			urloc, err = liburl.Parse(activityURL)
			if nil != err {
				var nada []byte
				return nada, erorr.Errorf("fedifinger: problem parsing HTTP(S) URL %q: %w", activityURL, err)
			}
			if nil == urloc {
				var nada []byte
				return nada, errNilParsedURL
			}
		}

		switch urloc.Scheme {
		case "http","https":
			// nothing here
		default:
			var nada []byte
			return nada, erorr.Errorf("fedifinger: not an HTTP(S) URL — %q", activityURL)
		}

		var header http.Header = http.Header{}
		header.Add("Accept", contentTypeActivity)

		httprequest = http.Request{
			Method: http.MethodGet,
			URL: urloc,
			Header: header,
		}
	}

	var httpresponse *http.Response
	{
		var err error
		httpresponse, err = http.DefaultClient.Do(&httprequest)
		if nil != err {
			var nada []byte
			return nada, erorr.Errorf("fedifinger: problem making HTTP(S) request to %q: %w", activityURL, err)
		}
		if nil == httpresponse {
			var nada []byte
			return nada, errNilHTTPResponse
		}
	}

	{
		if 400 <= httpresponse.StatusCode && httpresponse.StatusCode <= 599 {
			var nada []byte
			return nada, errhttp.Return(httpresponse.StatusCode)
		}
		if 200 != httpresponse.StatusCode {
			var nada []byte
			return nada, erorr.Errorf("fedifinger: not HTTP 200 response from %q — %d", activityURL, httpresponse.StatusCode)
		}
	}

	var bytes []byte
	{
		var err error
		bytes, err = io.ReadAll(httpresponse.Body)
		if nil != err {
			var nada []byte
			return nada, erorr.Errorf("fedifinger: problem reading-all from HTTP-response-body from %q: %w", activityURL, err)
		}
		httpresponse.Body.Close()
	}

	return bytes, nil
}
