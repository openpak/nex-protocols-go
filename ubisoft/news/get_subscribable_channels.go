package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	news_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/news/types"
)

// GetSubscribableChannelsRequest holds the request parameters of News::GetSubscribableChannels
type GetSubscribableChannelsRequest struct {
	ResultRange types.ResultRange
}

// WriteTo writes the GetSubscribableChannelsRequest to the given writable
func (m GetSubscribableChannelsRequest) WriteTo(writable types.Writable) {
	m.ResultRange.WriteTo(writable)
}

// ExtractFrom extracts the GetSubscribableChannelsRequest from the given readable
func (m *GetSubscribableChannelsRequest) ExtractFrom(readable types.Readable) error {
	if err := m.ResultRange.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetSubscribableChannels.resultRange. %s", err.Error())
	}

	return nil
}

// GetSubscribableChannelsResponse holds the response parameters of News::GetSubscribableChannels
type GetSubscribableChannelsResponse struct {
	Channels types.List[news_types.NewsChannel]
}

// WriteTo writes the GetSubscribableChannelsResponse to the given writable
func (m GetSubscribableChannelsResponse) WriteTo(writable types.Writable) {
	m.Channels.WriteTo(writable)
}

// ExtractFrom extracts the GetSubscribableChannelsResponse from the given readable
func (m *GetSubscribableChannelsResponse) ExtractFrom(readable types.Readable) error {
	if err := m.Channels.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetSubscribableChannels.channels. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleGetSubscribableChannels(packet nex.PacketInterface) {
	if protocol.GetSubscribableChannels == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "News::GetSubscribableChannels not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params GetSubscribableChannelsRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.GetSubscribableChannels(err, packet, callID, params.ResultRange)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetSubscribableChannels(nil, packet, callID, params.ResultRange)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
