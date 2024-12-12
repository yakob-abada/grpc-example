package pagination

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"hash/crc32"
)

type Request interface {
	proto.Message
	GetPaginationToken() string
}

// calculateRequestChecksum calculates a checksum for all fields of the request that must be the same across calls.
func calculateRequestChecksum(request Request) (uint32, error) {
	// Clone the original request, clear fields that may vary across calls, then checksum the resulting message.
	clonedRequest := proto.Clone(request)
	r := clonedRequest.ProtoReflect()
	r.Clear(r.Descriptor().Fields().ByName("pagination_token"))

	data, err := proto.Marshal(clonedRequest)
	if err != nil {
		return 0, fmt.Errorf("calculate request checksum: %w", err)
	}
	return crc32.ChecksumIEEE(data), nil
}
