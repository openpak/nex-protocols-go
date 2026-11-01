package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	news_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/news/types"
)

// GetNewsHeadersByTypeRequest holds the request parameters of News::GetNewsHeadersByType
type GetNewsHeadersByTypeRequest struct {
	NewsChannelType types.String
	Range           types.ResultRange
}

// WriteTo writes the GetNewsHeadersByTypeRequest to the given writable
func (m GetNewsHeadersByTypeRequest) WriteTo(writable types.Writable) {
	m.NewsChannelType.WriteTo(writable)
	m.Range.WriteTo(writable)
}

// ExtractFrom extracts the GetNewsHeadersByTypeRequest from the given readable
func (m *GetNewsHeadersByTypeRequest) ExtractFrom(readable types.Readable) error {
	if err := m.NewsChannelType.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetNewsHeadersByType.newsChannelType. %s", err.Error())
	}

	if err := m.Range.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetNewsHeadersByType.range. %s", err.Error())
	}

	return nil
}

// GetNewsHeadersByTypeResponse holds the response parameters of News::GetNewsHeadersByType
type GetNewsHeadersByTypeResponse struct {
	NewsHeaders types.List[news_types.NewsHeader]
}

// WriteTo writes the GetNewsHeadersByTypeResponse to the given writable
func (m GetNewsHeadersByTypeResponse) WriteTo(writable types.Writable) {
	m.NewsHeaders.WriteTo(writable)
}

// ExtractFrom extracts the GetNewsHeadersByTypeResponse from the given readable
func (m *GetNewsHeadersByTypeResponse) ExtractFrom(readable types.Readable) error {
	if err := m.NewsHeaders.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetNewsHeadersByType.newsHeaders. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleGetNewsHeadersByType(packet nex.PacketInterface) {
	if protocol.GetNewsHeadersByType == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "News::GetNewsHeadersByType not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params GetNewsHeadersByTypeRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.GetNewsHeadersByType(err, packet, callID, params.NewsChannelType, params.Range)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetNewsHeadersByType(nil, packet, callID, params.NewsChannelType, params.Range)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
