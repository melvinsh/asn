package asn_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/melvinsh/asn/pkg/asn"
)

func TestFindSubnets(t *testing.T) {
	testCases := []struct {
		name        string
		asn         string
		expected    []string
		expectedErr error
	}{
		{
			name:        "Valid ASN",
			asn:         "AS1224",
			expected:    []string{"141.142.0.0/16", "198.17.196.0/24"},
			expectedErr: nil,
		},
		{
			name:        "Empty ASN",
			asn:         "AS2993",
			expected:    nil,
			expectedErr: nil,
		},
		{
			name:        "Invalid ASN",
			asn:         "invalid",
			expected:    nil,
			expectedErr: errors.New("ASN not found (404)"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := asn.FindSubnets(tc.asn)
			assert.Equal(t, actual, tc.expected)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
