package test

import (
	"github.com/yakob-abada/backend-match/pkg/pagination"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPageTokenStruct(t *testing.T) {
	t.Parallel()
	type pageToken struct {
		Int    int
		String string
	}
	for _, tt := range []struct {
		name string
		in   pageToken
	}{
		{
			name: "all set",
			in: pageToken{
				Int:    42,
				String: "foo",
			},
		},
		{
			name: "default value",
			in: pageToken{
				String: "foo",
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			str := pagination.EncodePageTokenStruct(tt.in)
			var out pageToken
			assert.NoError(t, pagination.DecodePageTokenStruct(str, &out), str)
			assert.Equal(t, tt.in, out)
		})
	}
}
