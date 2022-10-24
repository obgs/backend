package guidgql

import (
	"database/sql/driver"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
)

type GUID struct {
	Type string `json:"t"`
	UUID string `json:"i"`
}

func New(t string) func() GUID {
	return func() GUID {
		return GUID{
			Type: t,
			UUID: uuid.New().String(),
		}
	}
}

func (guid GUID) String() string {
	j, _ := json.Marshal(guid)
	return base64.StdEncoding.EncodeToString(j)
}

func MarshalGUID(g GUID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(g.String()))
	})
}

func UnmarshalGUID(v interface{}) (g GUID, err error) {
	s, ok := v.(string)
	if !ok {
		return g, fmt.Errorf("invalid type %T, expect string", v)
	}

	j, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return g, fmt.Errorf("invalid base64 string %s", s)
	}

	return g, json.Unmarshal(j, &g)
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
	switch v := src.(type) {
	case []byte:
		var s []byte
		_, err := base64.StdEncoding.Decode(s, v)
		if err != nil {
			return err
		}

		return json.Unmarshal(s, guid)
	case string:
		s, err := base64.StdEncoding.DecodeString(v)
		if err != nil {
			return err
		}

		return json.Unmarshal(s, guid)
	default:
		return fmt.Errorf("invalid type %T, expect []byte or string", src)
	}
}

func (guid GUID) Value() (driver.Value, error) {
	return guid.String(), nil
}
