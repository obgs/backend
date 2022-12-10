package stat

import (
	"database/sql/driver"
	"fmt"
	"io"
	"strconv"
)

type StatType int

const (
	Numeric StatType = iota
	Enum
)

// New returns the default value for StatType.
func New() StatType {
	return Numeric
}

func (t StatType) String() string {
	switch t {
	case Numeric:
		return "numeric"
	case Enum:
		return "enum"
	default:
		return "unknown"
	}
}

func (t StatType) Values() []string {
	switch t {
	case Numeric:
		return []string{"numeric"}
	case Enum:
		return []string{"enum"}
	default:
		return []string{"unknown"}
	}
}

func UnmarshalStatType(src interface{}) (StatType, error) {
	switch v := src.(type) {
	case string:
		switch v {
		case "numeric":
			return Numeric, nil
		case "enum":
			return Enum, nil
		default:
			return 0, fmt.Errorf("invalid StatType %q", v)
		}
	default:
		return 0, fmt.Errorf("invalid type %T for StatType", src)
	}
}

func (t StatType) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(t.String()))
}

func (st *StatType) UnmarshalGQL(v interface{}) error {
	s, err := UnmarshalStatType(v)
	if err != nil {
		return err
	}

	*st = s

	return nil
}

func (t *StatType) Scan(src interface{}) error {
	return t.UnmarshalGQL(src)
}

func (t StatType) Value() (driver.Value, error) {
	return t.String(), nil
}
