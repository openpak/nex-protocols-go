package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	user_storage_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/user-storage/types"
)

// UploadEndRequest holds the request parameters of User Storage::UploadEnd
type UploadEndRequest struct {
	PendingID  types.UInt64
	Result     types.Bool
	ContentKey user_storage_types.UserContentKey
}

// WriteTo writes the UploadEndRequest to the given writable
func (m UploadEndRequest) WriteTo(writable types.Writable) {
	m.PendingID.WriteTo(writable)
	m.Result.WriteTo(writable)
	m.ContentKey.WriteTo(writable)
}

// ExtractFrom extracts the UploadEndRequest from the given readable
func (m *UploadEndRequest) ExtractFrom(readable types.Readable) error {
	if err := m.PendingID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read UploadEnd.pendingID. %s", err.Error())
	}

	if err := m.Result.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read UploadEnd.result. %s", err.Error())
	}

	if err := m.ContentKey.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read UploadEnd.contentKey. %s", err.Error())
	}

	return nil
}

// UploadEndResponse holds the response parameters of User Storage::UploadEnd
type UploadEndResponse struct {
	ContentKey user_storage_types.UserContentKey
}

// WriteTo writes the UploadEndResponse to the given writable
func (m UploadEndResponse) WriteTo(writable types.Writable) {
	m.ContentKey.WriteTo(writable)
}

// ExtractFrom extracts the UploadEndResponse from the given readable
func (m *UploadEndResponse) ExtractFrom(readable types.Readable) error {
	if err := m.ContentKey.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read UploadEnd.contentKey. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleUploadEnd(packet nex.PacketInterface) {
	if protocol.UploadEnd == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "User Storage::UploadEnd not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params UploadEndRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.UploadEnd(err, packet, callID, params.PendingID, params.Result, params.ContentKey)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.UploadEnd(nil, packet, callID, params.PendingID, params.Result, params.ContentKey)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
