package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	user_storage_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/user-storage/types"
)

// GetContentDBRequest holds the request parameters of User Storage::GetContentDB
type GetContentDBRequest struct {
	ContentKey user_storage_types.UserContentKey
}

// WriteTo writes the GetContentDBRequest to the given writable
func (m GetContentDBRequest) WriteTo(writable types.Writable) {
	m.ContentKey.WriteTo(writable)
}

// ExtractFrom extracts the GetContentDBRequest from the given readable
func (m *GetContentDBRequest) ExtractFrom(readable types.Readable) error {
	if err := m.ContentKey.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetContentDB.contentKey. %s", err.Error())
	}

	return nil
}

// GetContentDBResponse holds the response parameters of User Storage::GetContentDB
type GetContentDBResponse struct {
	Data types.Buffer
}

// WriteTo writes the GetContentDBResponse to the given writable
func (m GetContentDBResponse) WriteTo(writable types.Writable) {
	m.Data.WriteTo(writable)
}

// ExtractFrom extracts the GetContentDBResponse from the given readable
func (m *GetContentDBResponse) ExtractFrom(readable types.Readable) error {
	if err := m.Data.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetContentDB.data. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleGetContentDB(packet nex.PacketInterface) {
	if protocol.GetContentDB == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "User Storage::GetContentDB not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params GetContentDBRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.GetContentDB(err, packet, callID, params.ContentKey)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetContentDB(nil, packet, callID, params.ContentKey)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
