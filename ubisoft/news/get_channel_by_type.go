package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	news_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/news/types"
)

// GetChannelByTypeRequest holds the request parameters of News::GetChannelByType
type GetChannelByTypeRequest struct {
	NewsChannelType types.String
}

// WriteTo writes the GetChannelByTypeRequest to the given writable
func (m GetChannelByTypeRequest) WriteTo(writable types.Writable) {
	m.NewsChannelType.WriteTo(writable)
}

// ExtractFrom extracts the GetChannelByTypeRequest from the given readable
func (m *GetChannelByTypeRequest) ExtractFrom(readable types.Readable) error {
	if err := m.NewsChannelType.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetChannelByType.newsChannelType. %s", err.Error())
	}

	return nil
}

// GetChannelByTypeResponse holds the response parameters of News::GetChannelByType
type GetChannelByTypeResponse struct {
	Channel news_types.NewsChannel
}

// WriteTo writes the GetChannelByTypeResponse to the given writable
func (m GetChannelByTypeResponse) WriteTo(writable types.Writable) {
	m.Channel.WriteTo(writable)
}

// ExtractFrom extracts the GetChannelByTypeResponse from the given readable
func (m *GetChannelByTypeResponse) ExtractFrom(readable types.Readable) error {
	if err := m.Channel.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetChannelByType.channel. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleGetChannelByType(packet nex.PacketInterface) {
	if protocol.GetChannelByType == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "News::GetChannelByType not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params GetChannelByTypeRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.GetChannelByType(err, packet, callID, params.NewsChannelType)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetChannelByType(nil, packet, callID, params.NewsChannelType)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
