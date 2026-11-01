package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	user_storage_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/user-storage/types"
)

// GetTagsRequest holds the request parameters of User Storage::GetTags
type GetTagsRequest struct {
	ContentKey user_storage_types.UserContentKey
}

// WriteTo writes the GetTagsRequest to the given writable
func (m GetTagsRequest) WriteTo(writable types.Writable) {
	m.ContentKey.WriteTo(writable)
}

// ExtractFrom extracts the GetTagsRequest from the given readable
func (m *GetTagsRequest) ExtractFrom(readable types.Readable) error {
	if err := m.ContentKey.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetTags.contentKey. %s", err.Error())
	}

	return nil
}

// GetTagsResponse holds the response parameters of User Storage::GetTags
type GetTagsResponse struct {
	TagIds types.List[types.UInt32]
}

// WriteTo writes the GetTagsResponse to the given writable
func (m GetTagsResponse) WriteTo(writable types.Writable) {
	m.TagIds.WriteTo(writable)
}

// ExtractFrom extracts the GetTagsResponse from the given readable
func (m *GetTagsResponse) ExtractFrom(readable types.Readable) error {
	if err := m.TagIds.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetTags.tagIds. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleGetTags(packet nex.PacketInterface) {
	if protocol.GetTags == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "User Storage::GetTags not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params GetTagsRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.GetTags(err, packet, callID, params.ContentKey)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetTags(nil, packet, callID, params.ContentKey)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
