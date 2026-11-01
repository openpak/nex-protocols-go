package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	news_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/news/types"
)

// GetNewsMessagesByTypeRequest holds the request parameters of News::GetNewsMessagesByType
type GetNewsMessagesByTypeRequest struct {
	NewsChannelType types.String
	Range           types.ResultRange
}

// WriteTo writes the GetNewsMessagesByTypeRequest to the given writable
func (m GetNewsMessagesByTypeRequest) WriteTo(writable types.Writable) {
	m.NewsChannelType.WriteTo(writable)
	m.Range.WriteTo(writable)
}

// ExtractFrom extracts the GetNewsMessagesByTypeRequest from the given readable
func (m *GetNewsMessagesByTypeRequest) ExtractFrom(readable types.Readable) error {
	if err := m.NewsChannelType.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetNewsMessagesByType.newsChannelType. %s", err.Error())
	}

	if err := m.Range.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetNewsMessagesByType.range. %s", err.Error())
	}

	return nil
}

// GetNewsMessagesByTypeResponse holds the response parameters of News::GetNewsMessagesByType
type GetNewsMessagesByTypeResponse struct {
	NewsMessages types.List[news_types.NewsMessage]
}

// WriteTo writes the GetNewsMessagesByTypeResponse to the given writable
func (m GetNewsMessagesByTypeResponse) WriteTo(writable types.Writable) {
	m.NewsMessages.WriteTo(writable)
}

// ExtractFrom extracts the GetNewsMessagesByTypeResponse from the given readable
func (m *GetNewsMessagesByTypeResponse) ExtractFrom(readable types.Readable) error {
	if err := m.NewsMessages.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetNewsMessagesByType.newsMessages. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleGetNewsMessagesByType(packet nex.PacketInterface) {
	if protocol.GetNewsMessagesByType == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "News::GetNewsMessagesByType not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params GetNewsMessagesByTypeRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.GetNewsMessagesByType(err, packet, callID, params.NewsChannelType, params.Range)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetNewsMessagesByType(nil, packet, callID, params.NewsChannelType, params.Range)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
