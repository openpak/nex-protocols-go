package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

// UnsubscribeChannelRequest holds the request parameters of News::UnsubscribeChannel
type UnsubscribeChannelRequest struct {
	NewsChannelID types.UInt32
}

// WriteTo writes the UnsubscribeChannelRequest to the given writable
func (m UnsubscribeChannelRequest) WriteTo(writable types.Writable) {
	m.NewsChannelID.WriteTo(writable)
}

// ExtractFrom extracts the UnsubscribeChannelRequest from the given readable
func (m *UnsubscribeChannelRequest) ExtractFrom(readable types.Readable) error {
	if err := m.NewsChannelID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read UnsubscribeChannel.newsChannelID. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleUnsubscribeChannel(packet nex.PacketInterface) {
	if protocol.UnsubscribeChannel == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "News::UnsubscribeChannel not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params UnsubscribeChannelRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.UnsubscribeChannel(err, packet, callID, params.NewsChannelID)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.UnsubscribeChannel(nil, packet, callID, params.NewsChannelID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
