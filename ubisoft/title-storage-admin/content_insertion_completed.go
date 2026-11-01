package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

// ContentInsertionCompletedRequest holds the request parameters of Title Storage Admin::ContentInsertionCompleted
type ContentInsertionCompletedRequest struct {
	ContentId   types.UInt32
	ContentSize types.UInt32
}

// WriteTo writes the ContentInsertionCompletedRequest to the given writable
func (m ContentInsertionCompletedRequest) WriteTo(writable types.Writable) {
	m.ContentId.WriteTo(writable)
	m.ContentSize.WriteTo(writable)
}

// ExtractFrom extracts the ContentInsertionCompletedRequest from the given readable
func (m *ContentInsertionCompletedRequest) ExtractFrom(readable types.Readable) error {
	if err := m.ContentId.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read ContentInsertionCompleted.contentId. %s", err.Error())
	}

	if err := m.ContentSize.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read ContentInsertionCompleted.contentSize. %s", err.Error())
	}

	return nil
}

// ContentInsertionCompletedResponse holds the response parameters of Title Storage Admin::ContentInsertionCompleted
type ContentInsertionCompletedResponse struct {
	Retval types.Bool
}

// WriteTo writes the ContentInsertionCompletedResponse to the given writable
func (m ContentInsertionCompletedResponse) WriteTo(writable types.Writable) {
	m.Retval.WriteTo(writable)
}

// ExtractFrom extracts the ContentInsertionCompletedResponse from the given readable
func (m *ContentInsertionCompletedResponse) ExtractFrom(readable types.Readable) error {
	if err := m.Retval.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read ContentInsertionCompleted.retval. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleContentInsertionCompleted(packet nex.PacketInterface) {
	if protocol.ContentInsertionCompleted == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Title Storage Admin::ContentInsertionCompleted not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params ContentInsertionCompletedRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.ContentInsertionCompleted(err, packet, callID, params.ContentId, params.ContentSize)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.ContentInsertionCompleted(nil, packet, callID, params.ContentId, params.ContentSize)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
