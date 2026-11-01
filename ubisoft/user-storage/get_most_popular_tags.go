package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	user_storage_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/user-storage/types"
)

// GetMostPopularTagsRequest holds the request parameters of User Storage::GetMostPopularTags
type GetMostPopularTagsRequest struct {
	ContentKey user_storage_types.UserContentKey
}

// WriteTo writes the GetMostPopularTagsRequest to the given writable
func (m GetMostPopularTagsRequest) WriteTo(writable types.Writable) {
	m.ContentKey.WriteTo(writable)
}

// ExtractFrom extracts the GetMostPopularTagsRequest from the given readable
func (m *GetMostPopularTagsRequest) ExtractFrom(readable types.Readable) error {
	if err := m.ContentKey.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetMostPopularTags.contentKey. %s", err.Error())
	}

	return nil
}

// GetMostPopularTagsResponse holds the response parameters of User Storage::GetMostPopularTags
type GetMostPopularTagsResponse struct {
	Tags                  types.List[user_storage_types.WeightedTag]
	TotalNumberOfTaggings types.UInt32
}

// WriteTo writes the GetMostPopularTagsResponse to the given writable
func (m GetMostPopularTagsResponse) WriteTo(writable types.Writable) {
	m.Tags.WriteTo(writable)
	m.TotalNumberOfTaggings.WriteTo(writable)
}

// ExtractFrom extracts the GetMostPopularTagsResponse from the given readable
func (m *GetMostPopularTagsResponse) ExtractFrom(readable types.Readable) error {
	if err := m.Tags.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetMostPopularTags.tags. %s", err.Error())
	}

	if err := m.TotalNumberOfTaggings.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetMostPopularTags.totalNumberOfTaggings. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleGetMostPopularTags(packet nex.PacketInterface) {
	if protocol.GetMostPopularTags == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "User Storage::GetMostPopularTags not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params GetMostPopularTagsRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.GetMostPopularTags(err, packet, callID, params.ContentKey)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetMostPopularTags(nil, packet, callID, params.ContentKey)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
