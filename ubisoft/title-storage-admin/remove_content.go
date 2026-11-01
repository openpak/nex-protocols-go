package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

// RemoveContentRequest holds the request parameters of Title Storage Admin::RemoveContent
type RemoveContentRequest struct {
	ContentId types.UInt32
}

// WriteTo writes the RemoveContentRequest to the given writable
func (m RemoveContentRequest) WriteTo(writable types.Writable) {
	m.ContentId.WriteTo(writable)
}

// ExtractFrom extracts the RemoveContentRequest from the given readable
func (m *RemoveContentRequest) ExtractFrom(readable types.Readable) error {
	if err := m.ContentId.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read RemoveContent.contentId. %s", err.Error())
	}

	return nil
}

// RemoveContentResponse holds the response parameters of Title Storage Admin::RemoveContent
type RemoveContentResponse struct {
	Retval types.Bool
}

// WriteTo writes the RemoveContentResponse to the given writable
func (m RemoveContentResponse) WriteTo(writable types.Writable) {
	m.Retval.WriteTo(writable)
}

// ExtractFrom extracts the RemoveContentResponse from the given readable
func (m *RemoveContentResponse) ExtractFrom(readable types.Readable) error {
	if err := m.Retval.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read RemoveContent.retval. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleRemoveContent(packet nex.PacketInterface) {
	if protocol.RemoveContent == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Title Storage Admin::RemoveContent not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params RemoveContentRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.RemoveContent(err, packet, callID, params.ContentId)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.RemoveContent(nil, packet, callID, params.ContentId)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
