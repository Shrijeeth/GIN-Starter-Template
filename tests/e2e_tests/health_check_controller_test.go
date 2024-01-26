package tests

import (
	"encoding/json"
	"log"
	"mymodule/config"
	"mymodule/tests"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func beforeAll(tb testing.TB, configs *config.Config) {
	log.Println("Before All")
}

func afterAll(tb testing.TB, configs *config.Config) {
	log.Println("After All")
}

func beforeEach(tb testing.TB, configs *config.Config) {
	log.Println("Before Each")
}

func afterEach(tb testing.TB, configs *config.Config) {
	log.Println("After Each")
}

func setupTestSuite(tb testing.TB, configs *config.Config) func(tb testing.TB) {
	beforeAll(tb, configs)
	return func(tb testing.TB) {
		afterAll(tb, configs)
	}
}

func setupTestCaseSuite(tb testing.TB, configs *config.Config) func(tb testing.TB) {
	beforeEach(tb, configs)
	return func(tb testing.TB) {
		afterEach(tb, configs)
	}
}

func TestHealthCheckController(t *testing.T) {
	testRouter, testConfig := tests.SetupTestRouter()
	recorder := httptest.NewRecorder()

	testSuite := setupTestSuite(t, testConfig)
	defer testSuite(t)

	t.Run("Health Check", func(t *testing.T) {
		t.Run("success case for healthy server and database", func(t *testing.T) {
			testCase := setupTestCaseSuite(t, testConfig)
			defer testCase(t)

			expectedResponse := map[string]interface{}{
				"server": "Healthy",
				"database": "Healthy",
			}
		
			req, _ := http.NewRequest("GET", "/health-check/", nil)
			var res map[string]interface{}
		
			testRouter.ServeHTTP(recorder, req)
			_ = json.Unmarshal(recorder.Body.Bytes(), &res)
		
			assert.Equal(t, 200, recorder.Code)
			assert.Equal(t, expectedResponse, res)
		})
	})
}