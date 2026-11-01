package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	user_storage_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/user-storage/types"
)

// SaveContentDBRequest holds the request parameters of User Storage::SaveContentDB
type SaveContentDBRequest struct {
	Properties types.List[user_storage_types.ContentProperty]
	Data       types.Buffer
	ContentKey user_storage_types.UserContentKey
}

// WriteTo writes the SaveContentDBRequest to the given writable
func (m SaveContentDBRequest) WriteTo(writable types.Writable) {
	m.Properties.WriteTo(writable)
	m.Data.WriteTo(writable)
	m.ContentKey.WriteTo(writable)
}

// ExtractFrom extracts the SaveContentDBRequest from the given readable
func (m *SaveContentDBRequest) ExtractFrom(readable types.Readable) error {
	if err := m.Properties.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SaveContentDB.properties. %s", err.Error())
	}

	if err := m.Data.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SaveContentDB.data. %s", err.Error())
	}

	if err := m.ContentKey.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SaveContentDB.contentKey. %s", err.Error())
	}

	return nil
}

// SaveContentDBResponse holds the response parameters of User Storage::SaveContentDB
type SaveContentDBResponse struct {
	ContentKey user_storage_types.UserContentKey
}

// WriteTo writes the SaveContentDBResponse to the given writable
func (m SaveContentDBResponse) WriteTo(writable types.Writable) {
	m.ContentKey.WriteTo(writable)
}

// ExtractFrom extracts the SaveContentDBResponse from the given readable
func (m *SaveContentDBResponse) ExtractFrom(readable types.Readable) error {
	if err := m.ContentKey.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SaveContentDB.contentKey. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleSaveContentDB(packet nex.PacketInterface) {
	if protocol.SaveContentDB == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "User Storage::SaveContentDB not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params SaveContentDBRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.SaveContentDB(err, packet, callID, params.Properties, params.Data, params.ContentKey)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.SaveContentDB(nil, packet, callID, params.Properties, params.Data, params.ContentKey)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
