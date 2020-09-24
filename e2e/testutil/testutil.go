package testutil

import (
	"e2e/models"
	"encoding/json"
	"fmt"
	"github.com/dghubble/sling"
	"io/ioutil"
	"net/http"
	"testing"
)

func AssertStatusCode(t *testing.T, actual, expect int) {
	if actual != expect {
		t.Errorf("statusCode = %v, want %v", actual, expect)
	}
}

func AssertService(t *testing.T, actual, expect string) {
	if actual != expect {
		t.Errorf("service = %v, want %v", actual, expect)
	}
}

func AssertMessage(t *testing.T, actual, expect string) {
	if actual != expect {
		t.Errorf("message = %v, want %v", actual, expect)
	}
}

func AssertHost(t *testing.T, actual, expect string) {
	if actual != expect {
		t.Errorf("host = %v, want %v", actual, expect)
	}
}

func AssertResponseHeader(t *testing.T, actual http.Header, expect map[string]string) {
	var headerCount int
	for k, v := range expect {
		if actual.Get(k) == v {
			headerCount++
			continue
		} else {
			t.Errorf("header key:%v = %s, want %s", k, actual.Get(k), v)
		}
	}
	if headerCount != len(expect) {
		t.Error("some headers are not match")
	}
}

func AssertResponseCookie(t *testing.T, actual []*http.Cookie, expect map[string]string) {
	var cookieCount int
	for k, v := range expect {
		for _, cookie := range actual {
			if cookie.Name == k {
				if cookie.Value == v {
					cookieCount++
					continue
				} else {
					t.Errorf("cookie key:%v = %s, want %s", k, cookie.Value, v)
				}
			}
		}
	}
	if cookieCount != len(expect) {
		t.Error("some cookies are not match")
	}
}

func HasBalanceIdCookie(t *testing.T, actual []*http.Cookie) {
	for _, cookie := range actual {
		if cookie.Name == "balanceid" {
			return
		}
	}
	t.Error("There isn't balanceid in cookie")
}

func UnhealthyRequest(t *testing.T, fields models.Fields) (*http.Response, models.ServiceResponse) {
	s := sling.New().Base("http://localhost:8080")
	req, err := s.New().Get("/unhealthy").Request()
	if err != nil {
		t.Errorf("error: %v", err)
	}
	return request(t, req)
}

func HealthyRequest(t *testing.T, fields models.Fields) (*http.Response, models.ServiceResponse) {
	s := sling.New().Base("http://localhost:8080")
	req, err := s.New().Get("/healthy").Request()
	if err != nil {
		t.Errorf("error: %v", err)
	}
	return request(t, req)
}

func Request(t *testing.T, fields models.Fields) (*http.Response, models.ServiceResponse) {
	s := sling.New().Base(fmt.Sprintf("%s://localhost", fields.Scheme))
	req, _ := s.New().Get(fields.Path).Request()
	req.Host = fields.Host
	if len(fields.RequestCookie) > 0 {
		for k, v := range fields.RequestCookie {
			req.AddCookie(&http.Cookie{
				Name:  k,
				Value: v,
			})
		}
	}
	return request(t, req)
}

func request(t *testing.T, req *http.Request) (*http.Response, models.ServiceResponse) {
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		t.Errorf("error: %v", err)
	}
	var m models.ServiceResponse
	b, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(b, &m)
	return resp, m
}
