package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

// SubscribeChannelRequest holds the request parameters of News::SubscribeChannel
type SubscribeChannelRequest struct {
	NewsChannelID types.UInt32
}

// WriteTo writes the SubscribeChannelRequest to the given writable
func (m SubscribeChannelRequest) WriteTo(writable types.Writable) {
	m.NewsChannelID.WriteTo(writable)
}

// ExtractFrom extracts the SubscribeChannelRequest from the given readable
func (m *SubscribeChannelRequest) ExtractFrom(readable types.Readable) error {
	if err := m.NewsChannelID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SubscribeChannel.newsChannelID. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleSubscribeChannel(packet nex.PacketInterface) {
	if protocol.SubscribeChannel == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "News::SubscribeChannel not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params SubscribeChannelRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.SubscribeChannel(err, packet, callID, params.NewsChannelID)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.SubscribeChannel(nil, packet, callID, params.NewsChannelID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
