package serializer

import (
	"fmt"
	"io/ioutil"

	"google.golang.org/protobuf/proto"
)

//WriteProtobufToBinaryFile writes protocol buffer message to binary file
func WriteProtobufToBinaryFile(message proto.Message, filename string) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("cannot marsha proto message to binary: %w", err)
	}

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("cannot write binary data to binary file: %w", err)
	}

	return nil
}
