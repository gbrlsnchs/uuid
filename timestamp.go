package uuid

import (
	"crypto/rand"
	"encoding/binary"
	"errors"
	"io"
	"time"
)

const (
	jd       = 2299160                   // days from the starting point of the Julian Date system until Oct 15th 1582
	unix     = 2440587                   // days from the starting point of the Julian Date system until Jan 1st 1970
	epoch    = unix - jd                 // days between Oct 15th 1582 and Jan 1st 1970
	sec      = epoch * 86400             // epoch in seconds
	sec100ns = sec * (time.Second / 100) // epoch in 100s of nanoseconds
)

func timestampUUID() (UUID, error) {
	now := time.Now().UnixNano() / 100 // how many 100s of nanoseconds elapsed since Unix Epoch
	diff := now + int64(sec100ns)      // how many 100s of nanoseconds elapsed since Oct 15 1582
	var guid UUID
	// 32 least significant bits (guid[:4]).
	// Convert the LS 32 bits to a uint32.
	timeLow := uint32(diff & 0xFFFFFFFF)
	// Next 16 least significant bits (guid[4:6]).
	// Clear the LS 32 bits from time low and convert it to a uint16.
	timeMid := uint16(diff >> 32 & 0xFFFF)
	// Next 16 least significant bits (guid[6:8]).
	// Clear the LS 48 bits from both time low and time mid and convert it to a unit16,
	// leaving the MS 4 bits free for OR'ing with the version bits later.
	timeHigh := uint16(diff >> 48 & 0x0FFF)
	binary.BigEndian.PutUint32(guid[:4], timeLow)
	binary.BigEndian.PutUint16(guid[4:6], timeMid)
	binary.BigEndian.PutUint16(guid[6:8], timeHigh)
	// Clock sequence is 16 random bits (guid[8:10]).
	if _, err := io.ReadFull(rand.Reader, guid[8:10]); err != nil {
		return Null, err
	}
	mac, err := macAddr()
	if err != nil {
		return Null, err
	}
	if len(mac) != 6 {
		return Null, errors.New("uuid: MAC address is not a 48-bit slice")
	}
	// 48 bits for the Node ID (guid[10:]).
	copy(guid[10:], mac)
	return guid, nil
}
