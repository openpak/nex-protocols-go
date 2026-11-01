package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	news_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/news/types"
)

// GetNewsMessagesRequest holds the request parameters of News::GetNewsMessages
type GetNewsMessagesRequest struct {
	NewsMessageIDs types.List[types.UInt32]
}

// WriteTo writes the GetNewsMessagesRequest to the given writable
func (m GetNewsMessagesRequest) WriteTo(writable types.Writable) {
	m.NewsMessageIDs.WriteTo(writable)
}

// ExtractFrom extracts the GetNewsMessagesRequest from the given readable
func (m *GetNewsMessagesRequest) ExtractFrom(readable types.Readable) error {
	if err := m.NewsMessageIDs.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetNewsMessages.newsMessageIDs. %s", err.Error())
	}

	return nil
}

// GetNewsMessagesResponse holds the response parameters of News::GetNewsMessages
type GetNewsMessagesResponse struct {
	NewsMessages types.List[news_types.NewsMessage]
}

// WriteTo writes the GetNewsMessagesResponse to the given writable
func (m GetNewsMessagesResponse) WriteTo(writable types.Writable) {
	m.NewsMessages.WriteTo(writable)
}

// ExtractFrom extracts the GetNewsMessagesResponse from the given readable
func (m *GetNewsMessagesResponse) ExtractFrom(readable types.Readable) error {
	if err := m.NewsMessages.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetNewsMessages.newsMessages. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleGetNewsMessages(packet nex.PacketInterface) {
	if protocol.GetNewsMessages == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "News::GetNewsMessages not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params GetNewsMessagesRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.GetNewsMessages(err, packet, callID, params.NewsMessageIDs)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetNewsMessages(nil, packet, callID, params.NewsMessageIDs)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
