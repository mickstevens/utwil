package utwil

import (
	"encoding/json"
	"fmt"
	"time"
)

// YMD is used to format time.Time for querying the Twilio REST API
const YMD = "2006-01-02"

// Time is a wrapper around time.Time to support JSON marshalling to/from
// the Twilio REST API, which uses RFC1123Z.
type Time struct {
	time.Time
}

// Marshals time.Time into the time.RFC1123Z format
func (t *Time) MarshalJSON() ([]byte, error) {
	str := t.Format(fmt.Sprintf(`"%s"`, time.RFC1123Z))
	return []byte(str), nil
}

// Unmarshals time.Time from the time.RFC1123Z format
func (t *Time) UnmarshalJSON(data []byte) error {
	ot, err := time.Parse(fmt.Sprintf(`"%s"`, time.RFC1123Z), string(data))
	if err != nil {
		return err
	}
	t.Time = ot
	return nil
}

func decodeJSON(buf []byte, result interface{}) error {
	err := checkJSON(buf)
	if err != nil {
		return err
	}

	return json.Unmarshal(buf, result)
}

func printJSON(data interface{}) {
	bs, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(bs))

}
