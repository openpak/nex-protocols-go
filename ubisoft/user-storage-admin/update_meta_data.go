package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	user_storage_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/user-storage/types"
)

// UpdateMetaDataRequest holds the request parameters of User Storage Admin::UpdateMetaData
type UpdateMetaDataRequest struct {
	ContentKey user_storage_types.UserContentKey
	Properties types.List[user_storage_types.ContentProperty]
}

// WriteTo writes the UpdateMetaDataRequest to the given writable
func (m UpdateMetaDataRequest) WriteTo(writable types.Writable) {
	m.ContentKey.WriteTo(writable)
	m.Properties.WriteTo(writable)
}

// ExtractFrom extracts the UpdateMetaDataRequest from the given readable
func (m *UpdateMetaDataRequest) ExtractFrom(readable types.Readable) error {
	if err := m.ContentKey.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read UpdateMetaData.contentKey. %s", err.Error())
	}

	if err := m.Properties.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read UpdateMetaData.properties. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleUpdateMetaData(packet nex.PacketInterface) {
	if protocol.UpdateMetaData == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "User Storage Admin::UpdateMetaData not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params UpdateMetaDataRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.UpdateMetaData(err, packet, callID, params.ContentKey, params.Properties)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.UpdateMetaData(nil, packet, callID, params.ContentKey, params.Properties)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
