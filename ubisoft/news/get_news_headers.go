package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	news_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/news/types"
)

// GetNewsHeadersRequest holds the request parameters of News::GetNewsHeaders
type GetNewsHeadersRequest struct {
	Recipient news_types.NewsRecipient
	Range     types.ResultRange
}

// WriteTo writes the GetNewsHeadersRequest to the given writable
func (m GetNewsHeadersRequest) WriteTo(writable types.Writable) {
	m.Recipient.WriteTo(writable)
	m.Range.WriteTo(writable)
}

// ExtractFrom extracts the GetNewsHeadersRequest from the given readable
func (m *GetNewsHeadersRequest) ExtractFrom(readable types.Readable) error {
	if err := m.Recipient.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetNewsHeaders.recipient. %s", err.Error())
	}

	if err := m.Range.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetNewsHeaders.range. %s", err.Error())
	}

	return nil
}

// GetNewsHeadersResponse holds the response parameters of News::GetNewsHeaders
type GetNewsHeadersResponse struct {
	NewsHeaders types.List[news_types.NewsHeader]
}

// WriteTo writes the GetNewsHeadersResponse to the given writable
func (m GetNewsHeadersResponse) WriteTo(writable types.Writable) {
	m.NewsHeaders.WriteTo(writable)
}

// ExtractFrom extracts the GetNewsHeadersResponse from the given readable
func (m *GetNewsHeadersResponse) ExtractFrom(readable types.Readable) error {
	if err := m.NewsHeaders.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetNewsHeaders.newsHeaders. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleGetNewsHeaders(packet nex.PacketInterface) {
	if protocol.GetNewsHeaders == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "News::GetNewsHeaders not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params GetNewsHeadersRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.GetNewsHeaders(err, packet, callID, params.Recipient, params.Range)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetNewsHeaders(nil, packet, callID, params.Recipient, params.Range)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
