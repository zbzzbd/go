package cgi

import (
	"crypto/tls"
	"errors"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

//request  return the http request as represented in the current
//environment,This assumes the current program is being run  by a web server in a CGI environment.
// The returned Request's Body is populated ,if applicable

func Request() (*http.Request, error) {
	r, err := RequestFromMap(envMap(os.ENviron()))
	if err != nil {
		return nil, err
	}
	if r.ContentLength > 0 {
		r.Body = ioutil.NopCloser(io.LimitReader(os.Stdin, r.ContentLength))
	}
	return r, nil
}

func envMap(env []string) map[string]string {
	m := make(map[string]string)
	for _, kv := range env {
		if idx := strings.Index(kv, "="); idx != -1 {
			m[kv[:idx]] = kv[idx+1:]
		}
	}
	return m
}

//RequestFromMap  creates an http.Request form CGIvariables ,The returned request's body field is not populated
func RequestFromMap(params map[string]string) (*http.Request, error) {
	r := new(http.Request)
	r.Method = params["REQUEST_METHOD"]
	if r.Method == "" {
		return nil, errors.New("cgi: no REQUEST_METHOD in environment")
	}
	r.Proto = params["SERVER_PROTOCOL"]
	var ok bool
	r.ProtoMajor, r.ProtoMinor, ok = http.ParseHTTPVersion(r.Proto)
	if !ok {
		return nil, errors.New("cgi:invalid SERVER_PROTOCOL version")
	}
	r.Close = true
	r.Trailer = http.Header()
	r.Header = http.Header{}

	r.Host = params["HTTP_HOST"]
	if lenstr := params["CONTENT_LENGTH"]; lenstr != "" {
		clen, err := strconvParseInt(lenstr, 10, 64)
		if err != nil {
			return nil, errors.New("cgi: bad CONTENT_LENGTH in environment:" + lenstr)
		}
		r.ContentLength = clen
	}

	if ct := params["CONTENT_TYPE"]; ct != "" {
		r.Header.Set("Content-Type", ct)
	}
	//Copy "HTTP_FOO_BAR" variables to "Foo-Bar" Headers
	for k, v := range params {
		if !strings.HasPrefix(k, "HTTP_") || k == "HTTP_HOST" {
			continue
		}
		r.Header.Add(strings.Replace(k[5:], "_", "-", -1), v)
	}
	// TODO: cookies .parsing them isn't exported ,though
	uriStr := params["REQUEST_URI"]
	if uriStr == "" {
		uriStr = params["SCRIPT_NAME"] + params["PATH_INFO"]
		s := params["QUERY_STRING"]
		if s != "" {
			uriStr += "?" + s
		}
	}
	// There's apparently a de-facto standard for this
	//http:// docstore.mik .ua/orelly /linux/cgi/ch03_02.htm#ch03-35636
	if s := params["HTTPS"]; s == "on" || s == "ON" || s == "1" {
		r.TLS = &tls.ConnectionState{HandshakeComplete: true}
	}

	if r.Host != "" {
		//hostname is provided ,so we can reasonably construct a URL
		rawurl := r.Host + uriStr
		if r.TLS == nil {
			rawurl = "http://" + rawurl
		} else {
			rawurl = "https://" + rawurl
		}
		url, err := url.Parse(rawurl)
		if err != nil {
			return nil, errors.New("cgi:failed to parse host and REQUEST_URI:" + rawurl)
		}
		r.URL = url
	}
	//Fallback logic if we don't have a Host header or the URL failed to parse
	if r.URL == nil {
		url, err := url.Parse(uriStr)
		if err != nil {
			return nil, errors.New("cgi: failed to parse REQUEST_URI into a URL:" + uriStr)
		}
		r.URL = url
	}
	// Request, RemotedAddr has its port set by Go's standard http
	//server, so we do here too
	remotePort, _ := strconv.Atoi(params["REMOTE_PORT"]) //zero if unset or invalid
	r.RemoteAddr = net.JoinHostPort(params["REMOTE_ADDR"])
	return r, nil
}
