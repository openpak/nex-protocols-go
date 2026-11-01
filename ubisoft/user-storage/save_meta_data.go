package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	user_storage_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/user-storage/types"
)

// SaveMetaDataRequest holds the request parameters of User Storage::SaveMetaData
type SaveMetaDataRequest struct {
	Properties types.List[user_storage_types.ContentProperty]
	ContentKey user_storage_types.UserContentKey
}

// WriteTo writes the SaveMetaDataRequest to the given writable
func (m SaveMetaDataRequest) WriteTo(writable types.Writable) {
	m.Properties.WriteTo(writable)
	m.ContentKey.WriteTo(writable)
}

// ExtractFrom extracts the SaveMetaDataRequest from the given readable
func (m *SaveMetaDataRequest) ExtractFrom(readable types.Readable) error {
	if err := m.Properties.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SaveMetaData.properties. %s", err.Error())
	}

	if err := m.ContentKey.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SaveMetaData.contentKey. %s", err.Error())
	}

	return nil
}

// SaveMetaDataResponse holds the response parameters of User Storage::SaveMetaData
type SaveMetaDataResponse struct {
	ContentKey user_storage_types.UserContentKey
}

// WriteTo writes the SaveMetaDataResponse to the given writable
func (m SaveMetaDataResponse) WriteTo(writable types.Writable) {
	m.ContentKey.WriteTo(writable)
}

// ExtractFrom extracts the SaveMetaDataResponse from the given readable
func (m *SaveMetaDataResponse) ExtractFrom(readable types.Readable) error {
	if err := m.ContentKey.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SaveMetaData.contentKey. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleSaveMetaData(packet nex.PacketInterface) {
	if protocol.SaveMetaData == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "User Storage::SaveMetaData not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params SaveMetaDataRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.SaveMetaData(err, packet, callID, params.Properties, params.ContentKey)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.SaveMetaData(nil, packet, callID, params.Properties, params.ContentKey)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
