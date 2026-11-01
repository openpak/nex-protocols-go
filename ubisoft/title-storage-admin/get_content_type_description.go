package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	title_storage_admin_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/title-storage-admin/types"
)

// GetContentTypeDescriptionRequest holds the request parameters of Title Storage Admin::GetContentTypeDescription
type GetContentTypeDescriptionRequest struct {
	ContentType types.String
}

// WriteTo writes the GetContentTypeDescriptionRequest to the given writable
func (m GetContentTypeDescriptionRequest) WriteTo(writable types.Writable) {
	m.ContentType.WriteTo(writable)
}

// ExtractFrom extracts the GetContentTypeDescriptionRequest from the given readable
func (m *GetContentTypeDescriptionRequest) ExtractFrom(readable types.Readable) error {
	if err := m.ContentType.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetContentTypeDescription.contentType. %s", err.Error())
	}

	return nil
}

// GetContentTypeDescriptionResponse holds the response parameters of Title Storage Admin::GetContentTypeDescription
type GetContentTypeDescriptionResponse struct {
	ExtensionFields types.List[title_storage_admin_types.ContentTypeExtensionField]
}

// WriteTo writes the GetContentTypeDescriptionResponse to the given writable
func (m GetContentTypeDescriptionResponse) WriteTo(writable types.Writable) {
	m.ExtensionFields.WriteTo(writable)
}

// ExtractFrom extracts the GetContentTypeDescriptionResponse from the given readable
func (m *GetContentTypeDescriptionResponse) ExtractFrom(readable types.Readable) error {
	if err := m.ExtensionFields.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetContentTypeDescription.extensionFields. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleGetContentTypeDescription(packet nex.PacketInterface) {
	if protocol.GetContentTypeDescription == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Title Storage Admin::GetContentTypeDescription not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params GetContentTypeDescriptionRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.GetContentTypeDescription(err, packet, callID, params.ContentType)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetContentTypeDescription(nil, packet, callID, params.ContentType)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
