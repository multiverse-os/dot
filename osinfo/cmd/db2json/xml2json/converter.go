package xml2json

import (
	"bytes"
	"io"
)

func Convert(r io.Reader, ps ...plugin) (*bytes.Buffer, error) {
	root := &Node{}
	err := NewDecoder(r, ps...).Decode(root)
	if err != nil {
		return nil, err
	}
	buffer := new(bytes.Buffer)
	jsonBytes, err := json.MarshalIndent(&buffer, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	return jsonBytes, nil
}
