package web

import (
  "fmt"
  //"io/ioutil"
  "net/http"
  "net/http/httptest"
  "net/url"
  "regexp"
  "testing"
)

func TestHandlers(t *testing.T){
  fmt.Println("running test from web pkg")
  tests := []struct{
    Description string
    Handler func(http.ResponsWriter, *http.Request)
    Path string
    Params url.Values
    Status int
    Match map[string]bool
  }{
    {
      Description: "RequestInspector",
      Handler: RequestInspector,
      Path: "/",
      Status: http.StatusOK,
      Match: map[string]bool{
        "URL":true,
      },
    },
  }

  for _, test := range tests {
    record := httptest.NewRecorder()
    req := &http.Request{
      Method: "GET",
      URL: &url.URL{Path: test.Path},
      Form: test.Params,
    }
    test.Handler(record, req)
    fmt.Println("Description: ", test.Description)
    fmt.Println("StatusCode: ", record.Code)
    fmt.Println("Body: ", record.Body)

    if got, want := record.Code, test.Status; got != want{
      t.Errorf("%s: response code = %d, want %d", test.Description, got, want)
    }

    for re, match := range test.Match{
      if got := regexp.MustCompile(re).Match(record.Body.Bytes()); got != match{
        t.Errorf("%s: %q ~ /%s/ = %v, want %v", test.Description, record.Body, re, got, want)
      }
    }
  }
}

func TestAppHandler(t *testing.T){
  App.SetRoute()
  srv := httptest.NewServer(App.Router)
  defer srv.Close()
  srvURL, err := url.Parse(srv.URL)
  if err != nil{
    t.Fatal(err)
  }
  fmt.Println("Test server URL", srvURL)
  tests := []struct{
    Path string
    Method string
    StatusCode int
  }{
    {
      Path: srv.URL + "/pages/foo/",
      Method: "GET",
      StatusCode: 200,
    },
		{
			Path:       srv.URL + "/pages/",
			Method:     "GET",
			StatusCode: 200,
		},
		{
			Path:       srv.URL + "/",
			Method:     "GET",
			StatusCode: 200,
		},
		{
			Path:       srv.URL + "/NotFound/",
			Method:     "GET",
			StatusCode: 404,
		},
  }

  for _, test := range tests{
    fmt.Println("Request : ", test.Method, test.Path, test.StatusCode)
    req, err := http.NewRequest(test.Method, test.Path, nil)
    if err != nil{
      t.Fatal(err)
    }
    client := &http.client{}
    resp, err := client.Do(req)
    if err != nil{
      t.Fatal(err)
    }
    defer resp.Body.Close()
    if resp.StatusCode != test.StatusCode{
      t.Error("Error the status code do not match the expectation: ", resp.StatusCode)
    }
  }
}
