package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

// SendTagRequest holds the request parameters of Tracking 3::SendTag
type SendTagRequest struct {
	TrackingID types.UInt32
	Tag        types.String
	Attributes types.String
	DeltaTime  types.UInt32
}

// WriteTo writes the SendTagRequest to the given writable
func (m SendTagRequest) WriteTo(writable types.Writable) {
	m.TrackingID.WriteTo(writable)
	m.Tag.WriteTo(writable)
	m.Attributes.WriteTo(writable)
	m.DeltaTime.WriteTo(writable)
}

// ExtractFrom extracts the SendTagRequest from the given readable
func (m *SendTagRequest) ExtractFrom(readable types.Readable) error {
	if err := m.TrackingID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SendTag.trackingID. %s", err.Error())
	}

	if err := m.Tag.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SendTag.tag. %s", err.Error())
	}

	if err := m.Attributes.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SendTag.attributes. %s", err.Error())
	}

	if err := m.DeltaTime.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SendTag.deltaTime. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleSendTag(packet nex.PacketInterface) {
	if protocol.SendTag == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Tracking 3::SendTag not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params SendTagRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.SendTag(err, packet, callID, params.TrackingID, params.Tag, params.Attributes, params.DeltaTime)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.SendTag(nil, packet, callID, params.TrackingID, params.Tag, params.Attributes, params.DeltaTime)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
