package e2e_test

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	adminToken = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mjc4NTE5NzIsInJvbGUiOjB9.2My2wlmg6qvCFI-87nRahcPNr7H11vYpI5asyP8Qfwc"
	userToken  = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mjc4NTE5NzIsInJvbGUiOjF9.cZQq1MmrSo5yEPkv7_cGwCWnGLYwuWFKOSGNBlc_FK0"
)

func createRequest(method, route, body string, headers [][2]string) (*http.Request, error) {
	req, err := http.NewRequest(method, "http://localhost:8080"+route, strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	for _, header := range headers {
		req.Header.Add(header[0], header[1])
	}

	return req, nil
}

func sendRequest(
	t *testing.T,
	client *http.Client,
	req *http.Request,
	expectedStatus int,
	parseResp bool,
	respBody interface{},
) {
	t.Helper()
	resp, err := client.Do(req)
	require.NoError(t, err)

	require.Equal(t, expectedStatus, resp.StatusCode)

	if parseResp {
		err = json.NewDecoder(resp.Body).Decode(respBody)
		require.NoError(t, err)
	}

	resp.Body.Close()
}

func TestCreateBanner(t *testing.T) {
	client := http.Client{}

	tests := []struct {
		Name           string
		Method         string
		route          string
		body           string
		headers        [][2]string
		expectedStatus int
		parseResp      bool
		respBody       interface{}
	}{
		{
			Name:           "user create banner",
			Method:         http.MethodPost,
			route:          "/banner",
			body:           "{\"tag_ids\": [1001, 1002],\"feature_id\": 5001,\"content\": {\"title\": \"some_title\",\"text\": \"some_text\",\"url\": \"some_url\"},\"is_active\": true}",
			headers:        [][2]string{{"Content-Type", "application/json"}},
			expectedStatus: http.StatusForbidden,
			parseResp:      false,
			respBody:       nil,
		},
		{
			Name:           "admin create banner",
			Method:         http.MethodPost,
			route:          "/banner",
			body:           "{\"tag_ids\": [1003, 1004],\"feature_id\": 5002,\"content\": {\"title\": \"some_title\",\"text\": \"some_text\",\"url\": \"some_url\"},\"is_active\": true}",
			headers:        [][2]string{{"Content-Type", "application/json"}},
			expectedStatus: http.StatusCreated,
			parseResp:      false,
			respBody:       nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			var token string
			if strings.Split(tt.Name, " ")[0] == "user" {
				token = userToken
			} else {
				token = adminToken
			}
			req, err := createRequest(tt.Method, tt.route, tt.body, append(tt.headers, [2]string{"auth-x", token}))
			require.NoError(t, err)

			sendRequest(t, &client, req, tt.expectedStatus, tt.parseResp, tt.respBody)
		})
	}
}

func TestGetBanner(t *testing.T) {
	client := http.Client{}

	var resp json.RawMessage
	tests := []struct {
		Name           string
		Method         string
		route          string
		body           string
		headers        [][2]string
		expectedStatus int
		parseResp      bool
		respBody       interface{}
	}{
		{
			Name:           "user get banner not found",
			Method:         http.MethodGet,
			route:          "/user_banner?tag_id=500&feature_id=500",
			body:           "",
			headers:        [][2]string{{"Content-Type", "application/json"}},
			expectedStatus: http.StatusNotFound,
			parseResp:      true,
			respBody:       &resp,
		},
		{
			Name:           "admin create banner not active",
			Method:         http.MethodPost,
			route:          "/banner",
			body:           "{\"tag_ids\": [1005, 1006],\"feature_id\": 5005,\"content\": {\"title\": \"some_title\",\"text\": \"some_text\",\"url\": \"some_url\"},\"is_active\": false}",
			headers:        [][2]string{{"Content-Type", "application/json"}},
			expectedStatus: http.StatusCreated,
			parseResp:      true,
			respBody:       &resp,
		},
		{
			Name:           "user get banner not active",
			Method:         http.MethodGet,
			route:          "/user_banner?tag_id=1005&feature_id=5005",
			body:           "",
			headers:        [][2]string{{"Content-Type", "application/json"}},
			expectedStatus: http.StatusNotFound,
			parseResp:      true,
			respBody:       &resp,
		},
		{
			Name:           "admin get banner not active",
			Method:         http.MethodGet,
			route:          "/user_banner?tag_id=1005&feature_id=5005",
			body:           "",
			headers:        [][2]string{{"Content-Type", "application/json"}},
			expectedStatus: http.StatusOK,
			parseResp:      true,
			respBody:       &resp,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			var token string
			if strings.Split(tt.Name, " ")[0] == "user" {
				token = userToken
			} else {
				token = adminToken
			}
			req, err := createRequest(tt.Method, tt.route, tt.body, append(tt.headers, [2]string{"auth-x", token}))
			require.NoError(t, err)

			sendRequest(t, &client, req, tt.expectedStatus, tt.parseResp, tt.respBody)
		})
	}
}

func TestUpdateBanner(t *testing.T) {
	client := http.Client{}

	var resp json.RawMessage
	tests := []struct {
		Name           string
		Method         string
		route          string
		body           string
		headers        [][2]string
		expectedStatus int
		parseResp      bool
		respBody       interface{}
	}{
		{
			Name:           "admin update not existing banner",
			Method:         http.MethodPatch,
			route:          "/banner/555",
			body:           "{\"tag_ids\": [1008], \"feature_id\": 5008, \"content\": {\"title\": \"new_title\", \"text\": \"new_text\", \"url\": \"new_url\"},\"is_active\": false}",
			headers:        [][2]string{{"Content-Type", "application/json"}},
			expectedStatus: http.StatusNotFound,
			parseResp:      true,
			respBody:       &resp,
		},
		{
			Name:           "admin update ok",
			Method:         http.MethodPatch,
			route:          "/banner/600",
			body:           "{\"tag_ids\": [1008], \"feature_id\": 5008, \"content\": {\"title\": \"new_title\", \"text\": \"new_text\", \"url\": \"new_url\"},\"is_active\": true}",
			headers:        [][2]string{{"Content-Type", "application/json"}},
			expectedStatus: http.StatusOK,
			parseResp:      true,
			respBody:       &resp,
		},
		{
			Name:           "user update token err",
			Method:         http.MethodPatch,
			route:          "/banner/4",
			body:           "{\"tag_ids\": [1009], \"feature_id\": 5009, \"content\": {\"title\": \"new_title\", \"text\": \"new_text\", \"url\": \"new_url\"},\"is_active\": false}",
			headers:        [][2]string{{"Content-Type", "application/json"}},
			expectedStatus: http.StatusForbidden,
			parseResp:      false,
			respBody:       nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			var token string
			if strings.Split(tt.Name, " ")[0] == "user" {
				token = userToken
			} else {
				token = adminToken
			}
			req, err := createRequest(tt.Method, tt.route, tt.body, append(tt.headers, [2]string{"auth-x", token}))
			require.NoError(t, err)

			sendRequest(t, &client, req, tt.expectedStatus, tt.parseResp, tt.respBody)
		})
	}
}

func TestDeleteBanner(t *testing.T) {
	client := http.Client{}

	tests := []struct {
		Name           string
		Method         string
		route          string
		body           string
		headers        [][2]string
		expectedStatus int
		parseResp      bool
		respBody       interface{}
	}{
		{
			Name:           "user delete token err",
			Method:         http.MethodDelete,
			route:          "/banner/4",
			body:           "",
			headers:        [][2]string{{"Content-Type", "application/json"}},
			expectedStatus: http.StatusForbidden,
			parseResp:      false,
			respBody:       nil,
		},
		{
			Name:           "admin delete ok",
			Method:         http.MethodDelete,
			route:          "/banner/4",
			body:           "",
			headers:        [][2]string{{"Content-Type", "application/json"}},
			expectedStatus: http.StatusNoContent,
			parseResp:      false,
			respBody:       nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			var token string
			if strings.Split(tt.Name, " ")[0] == "user" {
				token = userToken
			} else {
				token = adminToken
			}
			req, err := createRequest(tt.Method, tt.route, tt.body, append(tt.headers, [2]string{"auth-x", token}))
			require.NoError(t, err)

			sendRequest(t, &client, req, tt.expectedStatus, tt.parseResp, tt.respBody)
		})
	}
}
