package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type TestData struct {
	Name string `json:"name"`
	Age  int    `json:"age,omitempty"`
}

func TestEncode(t *testing.T) {
	tests := []struct {
		name         string
		status       int
		data         interface{}
		expectedBody string
		expectedErr  string
	}{
		{
			name:         "Valid data",
			status:       http.StatusOK,
			data:         TestData{Name: "John", Age: 30},
			expectedBody: `{"name":"John","age":30}`,
			expectedErr:  "",
		},
		{
			name:         "Invalid data",
			status:       http.StatusBadRequest,
			data:         make(chan int), // unsupported type for encoding
			expectedBody: "",
			expectedErr:  "failed to encode data:",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			err := encode(w, tt.status, tt.data)

			assert.Equal(t, tt.expectedErr, fmt.Sprintf("%v", err)[:len(tt.expectedErr)])
			assert.Equal(t, tt.status, w.Code)
			assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
			assert.Equal(t, strings.TrimSpace(tt.expectedBody), strings.TrimSpace(w.Body.String()))
		})
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		name          string
		requestBody   interface{}
		targetType    interface{}
		expected      interface{}
		expectedError string
	}{
		{
			name:        "Valid JSON",
			requestBody: TestData{Name: "Alice", Age: 25},
			targetType:  TestData{},
			expected:    TestData{Name: "Alice", Age: 25},
		},
		{
			name:          "Invalid Type",
			requestBody:   `{"name": "Bob", "age": "thirty"}`, // age should be an int, but it's a string
			targetType:    TestData{},
			expected:      TestData{},
			expectedError: "failed to decode data",
		},
		{
			name:          "Invalid JSON",
			requestBody:   `{"name": "Charlie", "age": 40`, // missing closing brace
			targetType:    TestData{},
			expected:      TestData{},
			expectedError: "failed to decode data",
		},
		{
			name:          "Empty JSON",
			requestBody:   `{}`,
			targetType:    TestData{},
			expected:      TestData{},
			expectedError: "failed to decode data",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Marshal the request body into JSON
			requestBody, err := json.Marshal(tt.requestBody)
			assert.NoError(t, err)

			// Create a mock HTTP request with the JSON body
			req, err := http.NewRequest("POST", "/endpoint", bytes.NewBuffer(requestBody))
			assert.NoError(t, err)

			// Call the decode function with the appropriate type
			decoded, err := decode[TestData](req)

			// Check the output and error against expectations
			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, decoded)
			}
		})
	}
}
