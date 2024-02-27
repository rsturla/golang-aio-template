package http

import (
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
		responseBody  interface{}
		targetType    interface{}
		expected      interface{}
		expectedError string
	}{
		{
			name:         "Valid JSON",
			responseBody: TestData{Name: "Alice", Age: 25},
			targetType:   TestData{},
			expected:     TestData{Name: "Alice", Age: 25},
		},
		{
			name:          "Valid Missing Field",
			responseBody:  `{"name": "David"}`, // age is missing
			targetType:    TestData{},
			expected:      TestData{Name: "David"},
			expectedError: "",
		},
		{
			name:          "Invalid Type",
			responseBody:  `{"name": "Bob", "age": "thirty"}`, // age should be an int, but it's a string
			targetType:    TestData{},
			expected:      TestData{},
			expectedError: "failed to decode data",
		},
		{
			name:          "Invalid JSON",
			responseBody:  `{"name": "Charlie", "age": 40`, // missing closing brace
			targetType:    TestData{},
			expected:      TestData{},
			expectedError: "failed to decode data",
		},
		{
			name:          "Empty JSON",
			responseBody:  `{}`,
			targetType:    TestData{},
			expected:      TestData{},
			expectedError: "failed to decode data",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Marshal the response body into JSON
			responseBody, err := json.Marshal(tt.responseBody)
			assert.NoError(t, err)

			// Create a mock HTTP response with the JSON body
			resp := httptest.NewRecorder()
			_, err = resp.Write(responseBody)
			assert.NoError(t, err)

			// Call the decode function with the appropriate type
			decoded, err := decode[TestData](resp.Result())

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
