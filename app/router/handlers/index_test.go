package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nickhstr/go-web-service/app/router/handlers"
	"github.com/nickhstr/go-web-service/app/utils/test"
)

func TestIndex(t *testing.T) {
	respRec := httptest.NewRecorder()

	tests := []struct {
		method     string
		url        string
		statusCode int
	}{
		{"GET", "/", http.StatusOK},
	}

	t.Log("Given the Index handler")
	{
		for i, tt := range tests {
			t.Logf("\tTest %d: When making a %s request to '%s'", i, tt.method, tt.url)
			{
				handler := http.HandlerFunc(handlers.Index)
				req, err := http.NewRequest(tt.method, tt.url, nil)
				if err != nil {
					t.Fatal(err)
				}

				handler.ServeHTTP(respRec, req)

				if status := respRec.Code; status != http.StatusOK {
					t.Fatalf("\t%s\tThe status code should be %v", test.FAILURE, http.StatusOK)
				}
				t.Logf("\t%s\tThe status code should be %v", test.SUCCESS, http.StatusOK)

				expectedBody := "Hello World!"
				if body := respRec.Body.String(); body != expectedBody {
					t.Logf("\t%s\tThe body should match '%s'", test.FAILURE, expectedBody)
				}
				t.Logf("\t%s\tThe body should match '%s'", test.SUCCESS, expectedBody)
			}
		}
	}
}
