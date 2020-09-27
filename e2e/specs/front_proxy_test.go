package specs

import (
	"e2e/models"
	"e2e/testutil"
	"net/http"
	"testing"
	"time"
)

type TT struct {
	name   string
	fields models.Fields
}

func Test_BasicRouting(t *testing.T) {
	tests := []TT{
		{name: "/にアクセスするとservice1からHello, Worldを返す", fields: models.Fields{Scheme: "https", Host: "www.example.com", Path: "/", Status: http.StatusOK, Service: "service1", Message: "Hello, World"}},
		{name: "/service/2にアクセスするとservice2からHello, Worldを返す", fields: models.Fields{Scheme: "https", Host: "www.example.com", Path: "/service/2", Status: http.StatusOK, Service: "service2", Message: "Hello, World"}},
		{name: "localhostドメインでアクセスすると404を返す", fields: models.Fields{Scheme: "https", Host: "localhost", Path: "/", Status: http.StatusNotFound}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, body := testutil.Request(t, tt.fields)
			testutil.AssertStatusCode(t, resp.StatusCode, tt.fields.Status)
			testutil.AssertService(t, body.Service, tt.fields.Service)
			testutil.AssertMessage(t, body.Message, tt.fields.Message)
		})
	}
}

func Test_Header(t *testing.T) {
	tests := []TT{
		{name: "/にアクセスするとheaderにx-frame-options,x-xss-protectionを付与して返す", fields: models.Fields{Scheme: "https", Host: "www.example.com", Path: "/healthz", Status: http.StatusOK, Service: "service1", Message: "Status is healthy", RespHeader: map[string]string{"x-frame-options": "sameorigin", "x-xss-protection": "1; mode=block"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, body := testutil.Request(t, tt.fields)
			testutil.AssertStatusCode(t, resp.StatusCode, tt.fields.Status)
			testutil.AssertService(t, body.Service, tt.fields.Service)
			testutil.AssertMessage(t, body.Message, tt.fields.Message)
			testutil.AssertResponseHeader(t, resp.Header, tt.fields.RespHeader)
		})
	}
}

func Test_RingHash(t *testing.T) {
	tests := []TT{
		{name: "/service/2にアクセスするとレスポンスにSet-Cookieが付与される", fields: models.Fields{Scheme: "https", Host: "www.example.com", Path: "/service/2", Status: http.StatusOK, Service: "service2", RequestCookie: map[string]string{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, body := testutil.Request(t, tt.fields)
			testutil.HasBalanceIdCookie(t, resp.Cookies())

			for _, cookie := range resp.Cookies() {
				if cookie.Name == "balanceid" {
					tt.fields.RequestCookie[cookie.Name] = cookie.Value
				}
			}

			for i := 0; i < 5; i++ {
				r, b := testutil.Request(t, tt.fields)
				testutil.AssertStatusCode(t, r.StatusCode, tt.fields.Status)
				testutil.AssertService(t, b.Service, tt.fields.Service)
				testutil.AssertHost(t, b.Host, body.Host)
			}
		})
	}
}

func Test_Healthcheck(t *testing.T) {
	test := TT{name: "1つのhostがdownした場合、activeなhostにのみrequestが振り分けられる", fields: models.Fields{Scheme: "https", Host: "www.example.com", Path: "/healthz", Status: http.StatusOK, Service: "service1", Message: "Status is healthy"}}

	t.Run(test.name, func(t *testing.T) {
		_, _ = testutil.UnhealthyRequest(t, test.fields)
		time.Sleep(10 * time.Second)
		_, firstBody := testutil.Request(t, test.fields)
		var i int
		for i = 0; i < 5; i++ {
			resp, body := testutil.Request(t, test.fields)
			testutil.AssertStatusCode(t, resp.StatusCode, test.fields.Status)
			testutil.AssertService(t, body.Service, test.fields.Service)
			testutil.AssertMessage(t, body.Message, test.fields.Message)
			testutil.AssertHost(t, body.Host, firstBody.Host)
		}
		_, _ = testutil.HealthyRequest(t, test.fields)
	})
}
