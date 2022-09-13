package entity_test

import (
	"testing"

	"github.com/anggi-susanto/go-exercise/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewSku(t *testing.T) {
	s, err := entity.NewUser("a", "a")
	assert.Nil(t, err)
	assert.Equal(t, s.Type, "a")
	assert.Equal(t, s.Name, "a")
}

func TestSkuValidate(t *testing.T) {
	type test struct {
		types string
		name  string
		want  error
	}

	tests := []test{
		{
			types: "a",
			name:  "a",
			want:  nil,
		},
		{
			types: "",
			name:  "a",
			want:  entity.ErrInvalidEntity,
		},
		{
			types: "a",
			name:  "",
			want:  entity.ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		_, err := entity.NewUser(tc.name, tc.types)
		assert.Equal(t, err, tc.want)
	}
}
