package utils

import (
	"crypto/sha256"
	"encoding/binary"
)

func UuidToUint64(uuid string) (int64, error) {
	// Tạo một băm từ chuỗi UUID
	hash := sha256.Sum256([]byte(uuid))

	// Lấy 8 byte cuối cùng của băm
	bytes := hash[24:]

	// Chuyển 8 byte thành một số nguyên int64
	number := int64(binary.BigEndian.Uint64(bytes))

	return number, nil
}
