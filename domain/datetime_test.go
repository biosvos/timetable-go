package domain

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFromString(t *testing.T) {
	for _, date := range []string{"2023-02-16 12:40", "02-16 12:40", "16 12:40", "12:40", "40", ""} {
		t.Run(date, func(t *testing.T) {
			_, err := FromString(date)
			require.NoError(t, err)
		})
	}
}
