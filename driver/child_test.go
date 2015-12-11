package driver

import (
"testing"
)

func TestRequest(t *testing.T) {
	
   env:=map[string]string{
    "SERVER_PROTOCOL" :"HTTP/1.1",
    "REQUEST_METHOD" :"GET",
    "HTTP_HOST" :"example.com",
    "HTTP_REFERER" : "elsewhere",
    "HTTP_USER_AGENT" : "goclient",
    "HTTP_FOO_BAR" :"baz",
    "REQUEST_URI" :"/path?a=b",
    "CONTENT_LENGTH" :"123"
    "CONTENT_TYPE" :"text/xml",
    "REMOTE_ADDR" : "5.6.7.8"
    "REMOTE_PORT"  :"54321"
   }

   req,err:= RequestFromMap(env)
   if err !=nil {
   	t.Fatalf("RequestFromMap: %v", env)  //用"%v"标志，它可以以适当的格式输出任意的类型（包括数组和结构）
   }
   if g,e:=req.UserAgent(),"goclient";e!=g {
     t.Errorf("expected agent %q;got %q", e,g)//q代表什么意思
   }
   if g,e:=req.Method,"GET";e!=g {
   	t.Errorf("expected Method %q got %q",e,g)
   }
   if g,e:=req.Header.Get("Content-Type"),"text/xml";e!=g {
       t.Errorf("expetcted content-Type %q;got %q",e,g)
   }
   if g,e:=req.ContentLength,int64(123);e!=g{
   	t.Errorf("expected contentLength %q;got %q",e,g)
   }
   if g,e:=req.Referer(),"elsewhere";e!=g{
   	t.Errorf("expected referer %q;got %q", e,g)
   }
   if g,e:=req.Header ==nil {
   	t.Fatalf("unexpected nil Header")
   }
   if g,e :=req.Header.Get("Foo-Bar"),"baz";e!=g {
   	t.Errorf("expected Foo-Bar %q;got %q", e,g)
   }
   if g,e :=req.URL.String(),"http://example.com/path?a=b";e !=g {
   	t.Errorf("expected URL %q;got %q", e,g)
   }
   if g,e := req.FormValue("a"),"b"; e!=g {
   	t.Errorf("expected FormValue(a) %q;got %q", e,g)
   }
   
}

