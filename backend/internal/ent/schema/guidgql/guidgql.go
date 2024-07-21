package guidgql

import (
	"database/sql/driver"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/jxskiss/base62"
	"github.com/teris-io/shortid"
)

type GUID struct {
	Type Table
	ID   string
}

func New(t Table) func() GUID {
	return func() GUID {
		return GUID{
			Type: t,
			ID:   shortid.MustGenerate(),
		}
	}
}

func (guid GUID) String() string {
	return base62.EncodeToString([]byte(fmt.Sprintf("%d,%s", guid.Type, guid.ID)))
}

func UnmarshalGUID(src interface{}) (g GUID, err error) {
	switch v := src.(type) {
	case []byte:
		s, err := base62.Decode(v)
		if err != nil {
			return g, err
		}
		res := strings.Split(string(s), ",")
		t, err := strconv.Atoi(res[0])
		if err != nil {
			return g, err
		}

		return GUID{
			Type: Table(t),
			ID:   res[1],
		}, nil
	case string:
		var s []byte
		s, err = base62.DecodeString(v)
		if err != nil {
			return g, err
		}
		res := strings.Split(string(s), ",")
		t, err := strconv.Atoi(res[0])
		if err != nil {
			return g, err
		}

		return GUID{
			Type: Table(t),
			ID:   res[1],
		}, nil
	case GUID:
		return v, nil
	default:
		err = fmt.Errorf("invalid type %T, expect []byte or string", src)
		return g, err
	}
}

func (guid GUID) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(guid.String()))
}

func (guid *GUID) UnmarshalGQL(v interface{}) error {
	g, err := UnmarshalGUID(v)
	if err != nil {
		return err
	}

	*guid = g

	return nil
}

func (guid *GUID) Scan(src interface{}) error {
	return guid.UnmarshalGQL(src)
}

func (guid GUID) Value() (driver.Value, error) {
	return guid.String(), nil
}

func (guid GUID) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(guid.String())), nil
}

func (guid *GUID) UnmarshalJSON(data []byte) error {
	unquoted, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}

	return guid.UnmarshalGQL(unquoted)
}
