package telemetryhandler

import (
	"bytes"
	"encoding/binary"
	enums "f1-telemetry/internal/model/enums"
	packets "f1-telemetry/internal/model/packets"
	"fmt"
)

// TODO:
// sync pool.
// Return pointer
func ParsePacket(data []byte) (any, error, bool) {
	buf := bytes.NewReader(data)

	var header packets.PacketHeader
	if err := binary.Read(buf, binary.LittleEndian, &header); err != nil {
		return nil, fmt.Errorf("read header: %w", err), false
	}

	switch header.PacketId {
	case uint8(enums.CarTelemetry):
		var pkt packets.CarTelemetryPacket
		pkt.Header = header

		if err := binary.Read(buf, binary.LittleEndian, &pkt.CarTelemetryData); err != nil {
			return nil, err, false
		}
		if err := binary.Read(buf, binary.LittleEndian, &pkt.MfdPanelIndex); err != nil {
			return nil, err, false
		}
		if err := binary.Read(buf, binary.LittleEndian, &pkt.MfdPanelIndexSecondaryPlayer); err != nil {
			return nil, err, false
		}
		if err := binary.Read(buf, binary.LittleEndian, &pkt.SuggestedGear); err != nil {
			return nil, err, false
		}

		return pkt, nil, true
	case uint8(enums.Lap):
		var pkt packets.LapPacket
		pkt.Header = header

		if err := binary.Read(buf, binary.LittleEndian, &pkt.LapData); err != nil {
			return nil, err, false
		}

		if err := binary.Read(buf, binary.LittleEndian, &pkt.TimeTrialPBCarIdx); err != nil {
			return nil, err, false
		}

		if err := binary.Read(buf, binary.LittleEndian, &pkt.TimeTrialRivalCarIdx); err != nil {
			return nil, err, false
		}

		return pkt, nil, true

	case uint8(enums.Session):
		var pkt packets.SessionPacket

		err := binary.Read(buf, binary.LittleEndian, &pkt)
		if err != nil {
			return nil, err, false
		}

		return pkt, nil, true
	default:
		return nil, nil, false
	}
}
