package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

// GetContentDetailsRequest holds the request parameters of Title Storage::GetContentDetails
type GetContentDetailsRequest struct {
	ContentID types.UInt32
}

// WriteTo writes the GetContentDetailsRequest to the given writable
func (m GetContentDetailsRequest) WriteTo(writable types.Writable) {
	m.ContentID.WriteTo(writable)
}

// ExtractFrom extracts the GetContentDetailsRequest from the given readable
func (m *GetContentDetailsRequest) ExtractFrom(readable types.Readable) error {
	if err := m.ContentID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetContentDetails.contentID. %s", err.Error())
	}

	return nil
}

// GetContentDetailsResponse holds the response parameters of Title Storage::GetContentDetails
type GetContentDetailsResponse struct {
	Details types.DataHolder
}

// WriteTo writes the GetContentDetailsResponse to the given writable
func (m GetContentDetailsResponse) WriteTo(writable types.Writable) {
	m.Details.WriteTo(writable)
}

// ExtractFrom extracts the GetContentDetailsResponse from the given readable
func (m *GetContentDetailsResponse) ExtractFrom(readable types.Readable) error {
	if err := m.Details.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetContentDetails.details. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleGetContentDetails(packet nex.PacketInterface) {
	if protocol.GetContentDetails == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Title Storage::GetContentDetails not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params GetContentDetailsRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.GetContentDetails(err, packet, callID, params.ContentID)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetContentDetails(nil, packet, callID, params.ContentID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
