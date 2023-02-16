package domain

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTimeRecord(t *testing.T) {
	const startTime = "2400-12-31 23:39"
	const endTime = "2400-12-31 23:39"
	const id = "123"
	const memo = "memo"
	start, _ := FromString(startTime)
	record := NewTimeRecord(id, start, memo)
	end, _ := FromString(endTime)
	convertedRecord, _ := record.WithEnd(end)
	convertedRecord.id = "1222"

	require.Equal(t, id, record.id)
	require.Equal(t, startTime, record.start.String())
	require.Nil(t, record.end)
	require.Equal(t, memo, record.memo)
	require.Equal(t, "1222", convertedRecord.id)
	require.Equal(t, startTime, convertedRecord.start.String())
	require.Equal(t, endTime, convertedRecord.end.String())
	require.Equal(t, memo, convertedRecord.memo)
}
