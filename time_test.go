package nulls_test

import (
	"testing"
	"time"

	"github.com/gobuffalo/nulls"
	"github.com/stretchr/testify/require"
)

func Test_TimeUnmarshalJSON(t *testing.T) {
	r := require.New(t)

	tn := time.Now()
	tnStr := tn.Format(time.RFC3339)

	ns := nulls.Time{}
	r.NoError(ns.UnmarshalJSON([]byte("\"" + tnStr + "\"")))
	r.True(ns.Valid)
	r.Equal(tn.Format(time.RFC3339), ns.Time.Format(time.RFC3339))
}

func Test_TimeUnmarshalJSON_Overwrite(t *testing.T) {
	r := require.New(t)

	ns := nulls.NewTime(time.Now())

	r.NoError(ns.UnmarshalJSON([]byte("null")))
	r.False(ns.Valid)
	r.True(ns.Time.IsZero())
}
