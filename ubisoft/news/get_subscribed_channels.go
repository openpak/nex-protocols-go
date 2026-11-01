package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	news_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/news/types"
)

// GetSubscribedChannelsRequest holds the request parameters of News::GetSubscribedChannels
type GetSubscribedChannelsRequest struct {
	ResultRange types.ResultRange
}

// WriteTo writes the GetSubscribedChannelsRequest to the given writable
func (m GetSubscribedChannelsRequest) WriteTo(writable types.Writable) {
	m.ResultRange.WriteTo(writable)
}

// ExtractFrom extracts the GetSubscribedChannelsRequest from the given readable
func (m *GetSubscribedChannelsRequest) ExtractFrom(readable types.Readable) error {
	if err := m.ResultRange.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetSubscribedChannels.resultRange. %s", err.Error())
	}

	return nil
}

// GetSubscribedChannelsResponse holds the response parameters of News::GetSubscribedChannels
type GetSubscribedChannelsResponse struct {
	Channels types.List[news_types.NewsChannel]
}

// WriteTo writes the GetSubscribedChannelsResponse to the given writable
func (m GetSubscribedChannelsResponse) WriteTo(writable types.Writable) {
	m.Channels.WriteTo(writable)
}

// ExtractFrom extracts the GetSubscribedChannelsResponse from the given readable
func (m *GetSubscribedChannelsResponse) ExtractFrom(readable types.Readable) error {
	if err := m.Channels.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetSubscribedChannels.channels. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleGetSubscribedChannels(packet nex.PacketInterface) {
	if protocol.GetSubscribedChannels == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "News::GetSubscribedChannels not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params GetSubscribedChannelsRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.GetSubscribedChannels(err, packet, callID, params.ResultRange)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetSubscribedChannels(nil, packet, callID, params.ResultRange)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
