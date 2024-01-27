package types

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ffajarpratama/pos-wash-api/pkg/hash"
	"github.com/sirupsen/logrus"
)

type PhoneNumber string
type Password string

func (phone PhoneNumber) Format() PhoneNumber {
	phoneStr := string(phone)
	if strings.HasPrefix(phoneStr, "0") {
		return PhoneNumber("62" + phoneStr[1:])
	}

	return PhoneNumber(phoneStr)
}

func (pwd Password) Hash() Password {
	pass, err := hash.HashAndSalt([]byte(pwd))
	if err != nil {
		panic(err)
	}

	return Password(pass)
}

// follows github.com/sirupsen/logrus@v1.9.3/json_formatter.go
type CustomJSONFormatter struct{}

// Format renders a single log entry
func (f *CustomJSONFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	data := make(map[string]interface{}, len(entry.Data)+4)
	for k, v := range entry.Data {
		switch v := v.(type) {
		case error:
			// Otherwise errors are ignored by `encoding/json`
			// https://github.com/sirupsen/logrus/issues/137
			data[k] = v.Error()
		default:
			data[k] = v
		}
	}

	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	encoder := json.NewEncoder(b)
	encoder.SetIndent("", "  ")
	err := encoder.Encode(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal fields to JSON, %w", err)
	}

	return b.Bytes(), nil
}
