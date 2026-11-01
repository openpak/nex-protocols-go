package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	user_storage_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/user-storage/types"
)

// GetMetaDataRequest holds the request parameters of User Storage::GetMetaData
type GetMetaDataRequest struct {
	ContentKey user_storage_types.UserContentKey
}

// WriteTo writes the GetMetaDataRequest to the given writable
func (m GetMetaDataRequest) WriteTo(writable types.Writable) {
	m.ContentKey.WriteTo(writable)
}

// ExtractFrom extracts the GetMetaDataRequest from the given readable
func (m *GetMetaDataRequest) ExtractFrom(readable types.Readable) error {
	if err := m.ContentKey.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetMetaData.contentKey. %s", err.Error())
	}

	return nil
}

// GetMetaDataResponse holds the response parameters of User Storage::GetMetaData
type GetMetaDataResponse struct {
	Content user_storage_types.UserContent
}

// WriteTo writes the GetMetaDataResponse to the given writable
func (m GetMetaDataResponse) WriteTo(writable types.Writable) {
	m.Content.WriteTo(writable)
}

// ExtractFrom extracts the GetMetaDataResponse from the given readable
func (m *GetMetaDataResponse) ExtractFrom(readable types.Readable) error {
	if err := m.Content.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetMetaData.content. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleGetMetaData(packet nex.PacketInterface) {
	if protocol.GetMetaData == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "User Storage::GetMetaData not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params GetMetaDataRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.GetMetaData(err, packet, callID, params.ContentKey)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetMetaData(nil, packet, callID, params.ContentKey)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
