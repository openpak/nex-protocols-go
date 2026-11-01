package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

// CreateContentAndGetUploadInformationRequest holds the request parameters of Title Storage Admin::CreateContentAndGetUploadInformation
type CreateContentAndGetUploadInformationRequest struct {
	Name          types.String
	Description   types.String
	Key           types.String
	DateAvailable types.DateTime
	DateExpired   types.DateTime
	Flags         types.UInt32
	ExtendedFlags types.UInt32
	ContentType   types.String
	Extension     types.DataHolder
	Mime          types.String
}

// WriteTo writes the CreateContentAndGetUploadInformationRequest to the given writable
func (m CreateContentAndGetUploadInformationRequest) WriteTo(writable types.Writable) {
	m.Name.WriteTo(writable)
	m.Description.WriteTo(writable)
	m.Key.WriteTo(writable)
	m.DateAvailable.WriteTo(writable)
	m.DateExpired.WriteTo(writable)
	m.Flags.WriteTo(writable)
	m.ExtendedFlags.WriteTo(writable)
	m.ContentType.WriteTo(writable)
	m.Extension.WriteTo(writable)
	m.Mime.WriteTo(writable)
}

// ExtractFrom extracts the CreateContentAndGetUploadInformationRequest from the given readable
func (m *CreateContentAndGetUploadInformationRequest) ExtractFrom(readable types.Readable) error {
	if err := m.Name.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read CreateContentAndGetUploadInformation.name. %s", err.Error())
	}

	if err := m.Description.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read CreateContentAndGetUploadInformation.description. %s", err.Error())
	}

	if err := m.Key.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read CreateContentAndGetUploadInformation.key. %s", err.Error())
	}

	if err := m.DateAvailable.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read CreateContentAndGetUploadInformation.dateAvailable. %s", err.Error())
	}

	if err := m.DateExpired.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read CreateContentAndGetUploadInformation.dateExpired. %s", err.Error())
	}

	if err := m.Flags.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read CreateContentAndGetUploadInformation.flags. %s", err.Error())
	}

	if err := m.ExtendedFlags.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read CreateContentAndGetUploadInformation.extendedFlags. %s", err.Error())
	}

	if err := m.ContentType.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read CreateContentAndGetUploadInformation.contentType. %s", err.Error())
	}

	if err := m.Extension.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read CreateContentAndGetUploadInformation.extension. %s", err.Error())
	}

	if err := m.Mime.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read CreateContentAndGetUploadInformation.mime. %s", err.Error())
	}

	return nil
}

// CreateContentAndGetUploadInformationResponse holds the response parameters of Title Storage Admin::CreateContentAndGetUploadInformation
type CreateContentAndGetUploadInformationResponse struct {
	ContentId      types.UInt32
	UploadURL      types.String
	UploadMimeType types.String
	UploadBody     types.Map[types.String, types.String]
}

// WriteTo writes the CreateContentAndGetUploadInformationResponse to the given writable
func (m CreateContentAndGetUploadInformationResponse) WriteTo(writable types.Writable) {
	m.ContentId.WriteTo(writable)
	m.UploadURL.WriteTo(writable)
	m.UploadMimeType.WriteTo(writable)
	m.UploadBody.WriteTo(writable)
}

// ExtractFrom extracts the CreateContentAndGetUploadInformationResponse from the given readable
func (m *CreateContentAndGetUploadInformationResponse) ExtractFrom(readable types.Readable) error {
	if err := m.ContentId.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read CreateContentAndGetUploadInformation.contentId. %s", err.Error())
	}

	if err := m.UploadURL.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read CreateContentAndGetUploadInformation.uploadURL. %s", err.Error())
	}

	if err := m.UploadMimeType.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read CreateContentAndGetUploadInformation.uploadMimeType. %s", err.Error())
	}

	if err := m.UploadBody.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read CreateContentAndGetUploadInformation.uploadBody. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleCreateContentAndGetUploadInformation(packet nex.PacketInterface) {
	if protocol.CreateContentAndGetUploadInformation == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Title Storage Admin::CreateContentAndGetUploadInformation not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params CreateContentAndGetUploadInformationRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.CreateContentAndGetUploadInformation(err, packet, callID, params.Name, params.Description, params.Key, params.DateAvailable, params.DateExpired, params.Flags, params.ExtendedFlags, params.ContentType, params.Extension, params.Mime)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.CreateContentAndGetUploadInformation(nil, packet, callID, params.Name, params.Description, params.Key, params.DateAvailable, params.DateExpired, params.Flags, params.ExtendedFlags, params.ContentType, params.Extension, params.Mime)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
