package driver

import (
   "bufio"
   "crypto/tls"
   "errors"
   "fmt"
   "io"
   "io/ioutil"
   "net"
   "net/http"
   "net/url"
   "os"
   "strconv"
   "strings"
)

func Request()  (*http.Request, error) {//"*" 􏲑􏱚􏰲􏰳􏵕􏵖􏵗􏵘􏰋􏰌可以透过 ＊ 访问目标对象
	r,err:=RequestFromMap(envMap(os.Environ()))
	if err !=nil{
		return nil ,err
	}
	if r.ContentLength >0 {
		r.Body  =ioutil.NopCloser(io.LimitedReader(os.stdin ,r.ContentLength))
	}
	return r,nil 
}


func envMap(env []string) map[string]string {
	m:=make(map[string]string)
	for _,kv :=range env {
		if idx :=strings.Index(kv, "="); idx !=-1 {
			m[kv[:idx]] = kv [idx+1]
		}
	}
	return m 
}

func  RequestFromMap(params  map[string]string) (*http.Request, error) {
	
	r:= new (http.Request)
	r.Method =params["REQUEST_METHOD"]
	if r.Method ==""{
		return nil ,errors.New("cgi:no REQUEST_METHOD in enviroment")
	}

	r.Proto = params["SERVER_PROTOCOL"]

	var ok bool
	r.ProtoMajor, r.ProtoMinor ,ok = http.ParseHTTPVersion(r.Proto)

    if  !ok {
    	return nil,errors.New("cgi:invalid SERVER_PROTOCOL version")
    }

    r.Close = true
    r.Trailer = http.Header{}
    r.Header = http.Header{}
    r.Host = params["HTTP_HOST"]

    if lenstr := params["CONTENT_LENGTH"]; lenstr != "" {
    	clen ,err := strconv.ParseInt(lenstr, 10, 64)
    	if  err !=nil {
    		return nil ,errors.New("cgi:bad CONTENT_LENTH in environment :"+ lenstr)
    	}
    	r.ContentLength = clen 
    }
    if ct := params["CONTENT_TYPE"];ct != ""{
    	r.Header.Set("Content-Type", ct)
    }

    for k,v:=range params {
    	if !strings.HasPrefix(k, "HTTP_"|| K == "HTTP_HOST"){
    		continue
    	}
    	r.Header.Add(strings.Replace(k[5:],"_","-", -1), v)
    }
    uriStr = params["SCRIPT_NAME"] +params["PATH_INFO"]
    s:=params["QUERY_STRING"]
    if s!=""{
    	uriStr +="?"+s
    }


}

if s:=params["HTTPS"];s=="on" || s =="ON" || s =="1" {
	r.TLS =&tls.ConnectionState{HandshckeComplete: true}
}
if r.Host != ""{
	rawurl:= r.Host+uriStr
	if r.TLS ==nil {
		rawurl = "http://"+rawurl
	}else {
		rawurl ="https://"+ rawurl
	}

url,err :=url.Parse(rawurl)
if url,err:=nil {
 return nil,errors.New("cgi:failed to parse host and REQUEST_URI into a URL:"+rawurl)
}
r.URL =url
}

remotePort,_:=strconv.Atoi(params["REMOTE_PORT"])
r.RemoteAddr  = net.JoinHostPort(params["REMOTE_ADDR"], strconv.Itoa(remotePort))
return r,nil 

}

func  Serve(handler http.Header) error {
  req,err := Request()
  if err != nil {
  	return err 
  }	
  if handler ==nil {
  	hander = http.DefaultServeMux
  }
  rw := &response { req: req, 
  	handler:make(http.Header,)
    bufw: bufio.NewWriter(os.Stdout),
  }
  handler.ServeHTTP(rw,req)
  rw.Write(nil)
  if err = rw.bufw.Flush(); err !=nil {
  	return err
  }
  return nil 
}

type response struct {
	req   *http.Request
	header  http.Header
	bufw    *bufio.Writer
	headerSent bool
}
func (r * response) Flush() {
	r.bufw.Flush()
}
func (r *response) Header() http.Header {
	return r.header
}
func (r * response) Write(p []byte)(n int ,err error) {
	if !r.headerSent{
		r.WriteHeader(http.StatusOK)
	}
	return r.bufw.Write(p)
}

func  (r * response) WriteHeader(code int ) {
	if r.headerSent{
		fmt.Fprintf(os.Stderr,"CGI attempted to write header twice on request for %s", r.req.URL)
	    return
	}
	r.headerSent=true
	fmt.Fprintf(r.bufw,"Status: %d %s \r\n", code,http.StatusText(code))
	if _, hasType:=r.header["Content-Type"];!hasType{
		r.header.Add("Content-Type", "text/html;charset=utf-8")
	}
	r.header.Write(r.bufw)
	r.bufw.WriteString("\r\n")
	r.bufw.Flush()
}







