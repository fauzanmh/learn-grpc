package serializer

import (
	"learn-grpc/sample"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "learn-grpc/tmp/laptop.bin"

	laptop1 := sample.NewLaptop()
	err := WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)
}
