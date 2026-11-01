package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	user_storage_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/user-storage/types"
)

// SearchContentsRequest holds the request parameters of User Storage::SearchContents
type SearchContentsRequest struct {
	Query user_storage_types.UserStorageQuery
}

// WriteTo writes the SearchContentsRequest to the given writable
func (m SearchContentsRequest) WriteTo(writable types.Writable) {
	m.Query.WriteTo(writable)
}

// ExtractFrom extracts the SearchContentsRequest from the given readable
func (m *SearchContentsRequest) ExtractFrom(readable types.Readable) error {
	if err := m.Query.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SearchContents.query. %s", err.Error())
	}

	return nil
}

// SearchContentsResponse holds the response parameters of User Storage::SearchContents
type SearchContentsResponse struct {
	SearchResults types.List[user_storage_types.UserContent]
}

// WriteTo writes the SearchContentsResponse to the given writable
func (m SearchContentsResponse) WriteTo(writable types.Writable) {
	m.SearchResults.WriteTo(writable)
}

// ExtractFrom extracts the SearchContentsResponse from the given readable
func (m *SearchContentsResponse) ExtractFrom(readable types.Readable) error {
	if err := m.SearchResults.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SearchContents.searchResults. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleSearchContents(packet nex.PacketInterface) {
	if protocol.SearchContents == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "User Storage::SearchContents not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params SearchContentsRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.SearchContents(err, packet, callID, params.Query)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.SearchContents(nil, packet, callID, params.Query)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
