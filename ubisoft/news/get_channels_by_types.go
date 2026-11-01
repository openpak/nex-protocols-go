package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	news_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/news/types"
)

// GetChannelsByTypesRequest holds the request parameters of News::GetChannelsByTypes
type GetChannelsByTypesRequest struct {
	NewsChannelTypes types.List[types.String]
	ResultRange      types.ResultRange
}

// WriteTo writes the GetChannelsByTypesRequest to the given writable
func (m GetChannelsByTypesRequest) WriteTo(writable types.Writable) {
	m.NewsChannelTypes.WriteTo(writable)
	m.ResultRange.WriteTo(writable)
}

// ExtractFrom extracts the GetChannelsByTypesRequest from the given readable
func (m *GetChannelsByTypesRequest) ExtractFrom(readable types.Readable) error {
	if err := m.NewsChannelTypes.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetChannelsByTypes.newsChannelTypes. %s", err.Error())
	}

	if err := m.ResultRange.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetChannelsByTypes.resultRange. %s", err.Error())
	}

	return nil
}

// GetChannelsByTypesResponse holds the response parameters of News::GetChannelsByTypes
type GetChannelsByTypesResponse struct {
	Channels types.List[news_types.NewsChannel]
}

// WriteTo writes the GetChannelsByTypesResponse to the given writable
func (m GetChannelsByTypesResponse) WriteTo(writable types.Writable) {
	m.Channels.WriteTo(writable)
}

// ExtractFrom extracts the GetChannelsByTypesResponse from the given readable
func (m *GetChannelsByTypesResponse) ExtractFrom(readable types.Readable) error {
	if err := m.Channels.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetChannelsByTypes.channels. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleGetChannelsByTypes(packet nex.PacketInterface) {
	if protocol.GetChannelsByTypes == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "News::GetChannelsByTypes not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params GetChannelsByTypesRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.GetChannelsByTypes(err, packet, callID, params.NewsChannelTypes, params.ResultRange)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetChannelsByTypes(nil, packet, callID, params.NewsChannelTypes, params.ResultRange)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
