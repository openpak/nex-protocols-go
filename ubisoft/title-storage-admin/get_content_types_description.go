package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	title_storage_admin_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/title-storage-admin/types"
)

// GetContentTypesDescriptionRequest holds the request parameters of Title Storage Admin::GetContentTypesDescription
type GetContentTypesDescriptionRequest struct{}

// WriteTo writes the GetContentTypesDescriptionRequest to the given writable
func (m GetContentTypesDescriptionRequest) WriteTo(writable types.Writable) {
	_ = writable
}

// ExtractFrom extracts the GetContentTypesDescriptionRequest from the given readable
func (m *GetContentTypesDescriptionRequest) ExtractFrom(readable types.Readable) error {
	_ = readable
	return nil
}

// GetContentTypesDescriptionResponse holds the response parameters of Title Storage Admin::GetContentTypesDescription
type GetContentTypesDescriptionResponse struct {
	ContentDescriptions types.List[title_storage_admin_types.ContentTypeDescription]
}

// WriteTo writes the GetContentTypesDescriptionResponse to the given writable
func (m GetContentTypesDescriptionResponse) WriteTo(writable types.Writable) {
	m.ContentDescriptions.WriteTo(writable)
}

// ExtractFrom extracts the GetContentTypesDescriptionResponse from the given readable
func (m *GetContentTypesDescriptionResponse) ExtractFrom(readable types.Readable) error {
	if err := m.ContentDescriptions.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetContentTypesDescription.contentDescriptions. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleGetContentTypesDescription(packet nex.PacketInterface) {
	if protocol.GetContentTypesDescription == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Title Storage Admin::GetContentTypesDescription not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID

	rmcMessage, rmcError := protocol.GetContentTypesDescription(nil, packet, callID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
