package uuid

import (
	"encoding/json"
	"fmt"
	googleUUID "github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	u := New()

	assert.Equal(t, u.UUID.String(), u.Str)
	assert.NotEqual(t, u.UUID, googleUUID.Nil)
}

func TestParse(t *testing.T) {
	_uGood := `39662fb2-f1a0-4142-b936-a8e4fe52db65`
	_uBad := `39662fb2-f1a0-4142-b936-a8e4fe52dbxx`

	u, err := Parse(_uGood)

	assert.Nil(t, err)
	assert.Equal(t, u.Str, _uGood)
	assert.Equal(t, u.UUID.String(), _uGood)

	u, err = Parse(_uBad)

	assert.Error(t, err)
	assert.Equal(t, googleUUID.Nil, u.UUID)
	assert.Equal(t, googleUUID.Nil.String(), u.Str)
}

func TestUUID_Scan(t *testing.T) {
	_uGood := `39662fb2-f1a0-4142-b936-a8e4fe52db65`
	_uBad := `39662fb2-f1a0-4142-b936-a8e4fe52dbxx`

	u := &UUID{}
	err := u.Scan(_uGood)

	assert.Equal(t, _uGood, u.Str)
	assert.Nil(t, err)

	u = &UUID{}
	err = u.Scan(_uBad)

	assert.Equal(t, googleUUID.Nil.String(), u.Str)
	assert.Error(t, err)
}

func TestUUID_UnmarshalText(t *testing.T) {
	_uGood := []byte(`39662fb2-f1a0-4142-b936-a8e4fe52db65`)
	_uBad := []byte(`39662fb2-f1a0-4142-b936-a8e4fe52dbxx`)

	u := &UUID{}
	err := u.UnmarshalText(_uGood)

	assert.Equal(t, string(_uGood), u.Str)
	assert.Nil(t, err)

	u = &UUID{}
	err = u.UnmarshalText(_uBad)

	assert.Equal(t, googleUUID.Nil.String(), u.Str)
	assert.Error(t, err)
}

func TestUUID_UnmarshalBinary(t *testing.T) {
	_uGood, _ := googleUUID.Parse("39662fb2-f1a0-4142-b936-a8e4fe52db65")
	_uBad, _ := googleUUID.Parse(`39662fb2-f1a0-4142-b936-a8e4fe52dbxx`)

	_bGood, _ := _uGood.MarshalBinary()
	_bBad, _ := _uBad.MarshalBinary()

	u := &UUID{}
	err := u.UnmarshalBinary(_bGood)

	assert.Equal(t, _uGood.String(), u.Str)
	assert.Nil(t, err)

	u = &UUID{}
	err = u.UnmarshalText(_bBad)

	assert.Equal(t, googleUUID.Nil.String(), u.Str)
	assert.Error(t, err)
}

func TestUUID_UnmarshalJSON(t *testing.T) {
	_uGood := `"39662fb2-f1a0-4142-b936-a8e4fe52db65"`
	_uBad := `"39662fb2-f1a0-4142-b936-a8e4fe52dbxx"`

	u := UUID{}
	err := json.Unmarshal([]byte(_uGood), &u)

	assert.Equal(t, _uGood, fmt.Sprintf(`"%s"`, u.Str))
	assert.Nil(t, err)

	u = UUID{}
	err = json.Unmarshal([]byte(_uBad), &u)

	assert.Equal(t, googleUUID.Nil.String(), u.Str)
	assert.Error(t, err)
}
