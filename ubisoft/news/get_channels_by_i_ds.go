package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	news_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/news/types"
)

// GetChannelsByIDsRequest holds the request parameters of News::GetChannelsByIDs
type GetChannelsByIDsRequest struct {
	NewsChannelIDs types.List[types.UInt32]
}

// WriteTo writes the GetChannelsByIDsRequest to the given writable
func (m GetChannelsByIDsRequest) WriteTo(writable types.Writable) {
	m.NewsChannelIDs.WriteTo(writable)
}

// ExtractFrom extracts the GetChannelsByIDsRequest from the given readable
func (m *GetChannelsByIDsRequest) ExtractFrom(readable types.Readable) error {
	if err := m.NewsChannelIDs.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetChannelsByIDs.newsChannelIDs. %s", err.Error())
	}

	return nil
}

// GetChannelsByIDsResponse holds the response parameters of News::GetChannelsByIDs
type GetChannelsByIDsResponse struct {
	Channels types.List[news_types.NewsChannel]
}

// WriteTo writes the GetChannelsByIDsResponse to the given writable
func (m GetChannelsByIDsResponse) WriteTo(writable types.Writable) {
	m.Channels.WriteTo(writable)
}

// ExtractFrom extracts the GetChannelsByIDsResponse from the given readable
func (m *GetChannelsByIDsResponse) ExtractFrom(readable types.Readable) error {
	if err := m.Channels.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetChannelsByIDs.channels. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleGetChannelsByIDs(packet nex.PacketInterface) {
	if protocol.GetChannelsByIDs == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "News::GetChannelsByIDs not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params GetChannelsByIDsRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.GetChannelsByIDs(err, packet, callID, params.NewsChannelIDs)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetChannelsByIDs(nil, packet, callID, params.NewsChannelIDs)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
