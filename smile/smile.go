package smile

import (
	"encoding/base64"
	"encoding/json"

	"log"

	"github.com/gocollection/go-smile/decode"
	"github.com/gocollection/go-smile/domain"
)

func DecodeToJSON(smile []byte) (string, error) {
	obj, err := DecodeToObject(smile)
	if err != nil {
		return "", err
	}

	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}

func DecodeToObject(smile []byte) (interface{}, error) {
	header, err := domain.DecodeHeader(smile)
	if err != nil {
		log.Printf(
			"[go-smile] error decoding bytes. err: %s; base64-string: %s",
			err.Error(), base64.StdEncoding.EncodeToString(smile),
		)
		return "", err
	}

	if _, err := domain.DecodeHeader(smile[header.SizeBytes:]); err == nil {
		smile = smile[header.SizeBytes:]
	}

	var d decode.Decoder
	_, b, err := d.DecodeBytes(smile[header.SizeBytes:])
	return b, err
}
