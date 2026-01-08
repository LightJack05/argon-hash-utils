package argonhashutils

import (
	"encoding/base64"
	"fmt"
)

// Argon2idPasswordHash represents an Argon2id hashed password with its parameters
type Argon2idPasswordHash struct {
	Version     int
	Salt        []byte
	Hash        []byte
	Time        uint32
	Memory      uint32
	Parallelism uint8
}

func (h *Argon2idPasswordHash) ToString() string {
	return fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", h.Version, h.Memory, h.Time, h.Parallelism, base64.RawStdEncoding.EncodeToString(h.Salt), base64.RawStdEncoding.EncodeToString(h.Hash))
}

