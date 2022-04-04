package main

import (
	"crypto/tls"
	browser "github.com/EDDYCJY/fake-useragent"
	"net/http"
)

// roundTripper is a custom round tripper which adds the validated request headers.
type roundTripper struct {
	inner     http.RoundTripper
	userAgent string
}

// return required headers
func getTLSConfiguration(inner http.RoundTripper) http.RoundTripper {
	if trans, ok := inner.(*http.Transport); ok {
		trans.TLSClientConfig = getCloudFlareTLSConfiguration()
	}

	return &roundTripper{
		inner:     inner,
		userAgent: browser.Firefox(),
	}
}

// adds the required request headers
func (ug *roundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	if r.Header.Get("Accept-Language") == "" {
		r.Header.Set("Accept-Language", "en-US,en;q=0.5")
	}

	if r.Header.Get("Accept") == "" {
		// Accept-Encoding header
		r.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	}

	// only use fake user agent if no custom user agent is defined already
	if r.Header.Get("User-Agent") == "" {
		r.Header.Set("User-Agent", ug.userAgent)
	}

	// in case we don't have an inner transport layer from the round tripper
	if ug.inner == nil {
		return (&http.Transport{
			TLSClientConfig: getCloudFlareTLSConfiguration(),
		}).RoundTrip(r)
	}

	return ug.inner.RoundTrip(r)
}

// getCloudFlareTLSConfiguration returns an accepted client TLS configuration to not get detected by CloudFlare directly
func getCloudFlareTLSConfiguration() *tls.Config {
	return &tls.Config{
		PreferServerCipherSuites: false,
		CurvePreferences:         []tls.CurveID{tls.CurveP256, tls.CurveP384, tls.CurveP521, tls.X25519},
	}
}
