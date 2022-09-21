package nulls

import (
	"encoding/json"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/require"
)

func Test_UUID_UnmarshalJSON(t *testing.T) {
	r := require.New(t)
	id, err := uuid.NewV4()
	r.NoError(err)

	b, err := json.Marshal(id)
	r.NoError(err)

	nid := &UUID{}

	r.NoError(nid.UnmarshalJSON(b))

	r.True(nid.Valid)
	r.Equal(id.String(), nid.UUID.String())
}

func Test_UUID_NewUUIDValid(t *testing.T) {
	r := require.New(t)
	id, err := uuid.NewV4()
	r.NoError(err)

	nid := NewUUID(id)

	r.True(nid.Valid)
	r.Equal(id.String(), nid.UUID.String())
}

func Test_UUID_NewUUIDInvalid(t *testing.T) {
	r := require.New(t)

	nid := NewUUID(uuid.UUID{})

	r.False(nid.Valid)

	nid = NewUUID(uuid.Nil)

	r.False(nid.Valid)
}
