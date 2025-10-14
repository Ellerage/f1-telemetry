package telemetryhandler

import (
	"bytes"
	"encoding/binary"
	"f1-telemetry/internal/model"
	"fmt"
)

func ParsePacket(data []byte) (any, error, bool) {
	buf := bytes.NewReader(data)

	var header model.PacketHeader
	if err := binary.Read(buf, binary.LittleEndian, &header); err != nil {
		return nil, fmt.Errorf("read header: %w", err), false
	}

	switch header.PacketId {
	case uint8(model.Lap):
		var pkt model.LapPacket
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

	case uint8(model.Motion):
		var pkt model.MotionPacket
		pkt.Header = header

		if err := binary.Read(buf, binary.LittleEndian, &pkt.CarMotionData); err != nil {
			return nil, err, false
		}

		return pkt, nil, true
	case uint8(model.CarTelemetry):
		var pkt model.CarTelemetryPacket
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
	default:
		return nil, nil, false
	}
}
