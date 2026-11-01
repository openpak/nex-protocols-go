package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	user_storage_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/user-storage/types"
)

// SaveContentAndGetUploadInfoRequest holds the request parameters of User Storage::SaveContentAndGetUploadInfo
type SaveContentAndGetUploadInfoRequest struct {
	Properties types.List[user_storage_types.ContentProperty]
	Size       types.UInt32
	ContentKey user_storage_types.UserContentKey
}

// WriteTo writes the SaveContentAndGetUploadInfoRequest to the given writable
func (m SaveContentAndGetUploadInfoRequest) WriteTo(writable types.Writable) {
	m.Properties.WriteTo(writable)
	m.Size.WriteTo(writable)
	m.ContentKey.WriteTo(writable)
}

// ExtractFrom extracts the SaveContentAndGetUploadInfoRequest from the given readable
func (m *SaveContentAndGetUploadInfoRequest) ExtractFrom(readable types.Readable) error {
	if err := m.Properties.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SaveContentAndGetUploadInfo.properties. %s", err.Error())
	}

	if err := m.Size.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SaveContentAndGetUploadInfo.size. %s", err.Error())
	}

	if err := m.ContentKey.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SaveContentAndGetUploadInfo.contentKey. %s", err.Error())
	}

	return nil
}

// SaveContentAndGetUploadInfoResponse holds the response parameters of User Storage::SaveContentAndGetUploadInfo
type SaveContentAndGetUploadInfoResponse struct {
	UploadInfo user_storage_types.UserContentURL
	PendingID  types.UInt64
	Headers    types.List[types.String]
}

// WriteTo writes the SaveContentAndGetUploadInfoResponse to the given writable
func (m SaveContentAndGetUploadInfoResponse) WriteTo(writable types.Writable) {
	m.UploadInfo.WriteTo(writable)
	m.PendingID.WriteTo(writable)
	m.Headers.WriteTo(writable)
}

// ExtractFrom extracts the SaveContentAndGetUploadInfoResponse from the given readable
func (m *SaveContentAndGetUploadInfoResponse) ExtractFrom(readable types.Readable) error {
	if err := m.UploadInfo.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SaveContentAndGetUploadInfo.uploadInfo. %s", err.Error())
	}

	if err := m.PendingID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SaveContentAndGetUploadInfo.pendingID. %s", err.Error())
	}

	if err := m.Headers.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SaveContentAndGetUploadInfo.headers. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleSaveContentAndGetUploadInfo(packet nex.PacketInterface) {
	if protocol.SaveContentAndGetUploadInfo == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "User Storage::SaveContentAndGetUploadInfo not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params SaveContentAndGetUploadInfoRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.SaveContentAndGetUploadInfo(err, packet, callID, params.Properties, params.Size, params.ContentKey)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.SaveContentAndGetUploadInfo(nil, packet, callID, params.Properties, params.Size, params.ContentKey)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
