package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	news_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/news/types"
)

// GetNumberOfNewsRequest holds the request parameters of News::GetNumberOfNews
type GetNumberOfNewsRequest struct {
	Recipient news_types.NewsRecipient
}

// WriteTo writes the GetNumberOfNewsRequest to the given writable
func (m GetNumberOfNewsRequest) WriteTo(writable types.Writable) {
	m.Recipient.WriteTo(writable)
}

// ExtractFrom extracts the GetNumberOfNewsRequest from the given readable
func (m *GetNumberOfNewsRequest) ExtractFrom(readable types.Readable) error {
	if err := m.Recipient.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetNumberOfNews.recipient. %s", err.Error())
	}

	return nil
}

// GetNumberOfNewsResponse holds the response parameters of News::GetNumberOfNews
type GetNumberOfNewsResponse struct {
	NumberOfNews types.UInt32
}

// WriteTo writes the GetNumberOfNewsResponse to the given writable
func (m GetNumberOfNewsResponse) WriteTo(writable types.Writable) {
	m.NumberOfNews.WriteTo(writable)
}

// ExtractFrom extracts the GetNumberOfNewsResponse from the given readable
func (m *GetNumberOfNewsResponse) ExtractFrom(readable types.Readable) error {
	if err := m.NumberOfNews.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetNumberOfNews.numberOfNews. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleGetNumberOfNews(packet nex.PacketInterface) {
	if protocol.GetNumberOfNews == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "News::GetNumberOfNews not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params GetNumberOfNewsRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.GetNumberOfNews(err, packet, callID, params.Recipient)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetNumberOfNews(nil, packet, callID, params.Recipient)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
