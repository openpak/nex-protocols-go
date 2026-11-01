package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	news_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/news/types"
)

// GetChannelsRequest holds the request parameters of News::GetChannels
type GetChannelsRequest struct {
	ResultRange types.ResultRange
}

// WriteTo writes the GetChannelsRequest to the given writable
func (m GetChannelsRequest) WriteTo(writable types.Writable) {
	m.ResultRange.WriteTo(writable)
}

// ExtractFrom extracts the GetChannelsRequest from the given readable
func (m *GetChannelsRequest) ExtractFrom(readable types.Readable) error {
	if err := m.ResultRange.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetChannels.resultRange. %s", err.Error())
	}

	return nil
}

// GetChannelsResponse holds the response parameters of News::GetChannels
type GetChannelsResponse struct {
	Channels types.List[news_types.NewsChannel]
}

// WriteTo writes the GetChannelsResponse to the given writable
func (m GetChannelsResponse) WriteTo(writable types.Writable) {
	m.Channels.WriteTo(writable)
}

// ExtractFrom extracts the GetChannelsResponse from the given readable
func (m *GetChannelsResponse) ExtractFrom(readable types.Readable) error {
	if err := m.Channels.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetChannels.channels. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleGetChannels(packet nex.PacketInterface) {
	if protocol.GetChannels == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "News::GetChannels not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params GetChannelsRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.GetChannels(err, packet, callID, params.ResultRange)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetChannels(nil, packet, callID, params.ResultRange)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
