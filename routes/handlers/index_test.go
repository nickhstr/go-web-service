package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nickhstr/go-web-service/utils/test"
)

func TestIndex(t *testing.T) {
	respRec := httptest.NewRecorder()

	t.Log("Given the Index handler")
	{
		handler := http.HandlerFunc(Index)
		req, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		t.Logf("\tWhen making a %s request to '%s'", req.Method, req.URL)
		{
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
