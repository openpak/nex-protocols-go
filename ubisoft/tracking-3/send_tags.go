package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	tracking_3_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/tracking-3/types"
)

// SendTagsRequest holds the request parameters of Tracking 3::SendTags
type SendTagsRequest struct {
	TagData types.List[tracking_3_types.TrackingTag]
}

// WriteTo writes the SendTagsRequest to the given writable
func (m SendTagsRequest) WriteTo(writable types.Writable) {
	m.TagData.WriteTo(writable)
}

// ExtractFrom extracts the SendTagsRequest from the given readable
func (m *SendTagsRequest) ExtractFrom(readable types.Readable) error {
	if err := m.TagData.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SendTags.tagData. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleSendTags(packet nex.PacketInterface) {
	if protocol.SendTags == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Tracking 3::SendTags not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params SendTagsRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.SendTags(err, packet, callID, params.TagData)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.SendTags(nil, packet, callID, params.TagData)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
