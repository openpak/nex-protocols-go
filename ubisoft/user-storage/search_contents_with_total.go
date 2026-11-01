package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	user_storage_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/user-storage/types"
)

// SearchContentsWithTotalRequest holds the request parameters of User Storage::SearchContentsWithTotal
type SearchContentsWithTotalRequest struct {
	Query user_storage_types.UserStorageQuery
}

// WriteTo writes the SearchContentsWithTotalRequest to the given writable
func (m SearchContentsWithTotalRequest) WriteTo(writable types.Writable) {
	m.Query.WriteTo(writable)
}

// ExtractFrom extracts the SearchContentsWithTotalRequest from the given readable
func (m *SearchContentsWithTotalRequest) ExtractFrom(readable types.Readable) error {
	if err := m.Query.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SearchContentsWithTotal.query. %s", err.Error())
	}

	return nil
}

// SearchContentsWithTotalResponse holds the response parameters of User Storage::SearchContentsWithTotal
type SearchContentsWithTotalResponse struct {
	SearchResults types.List[user_storage_types.UserContent]
	TotalResults  types.UInt32
}

// WriteTo writes the SearchContentsWithTotalResponse to the given writable
func (m SearchContentsWithTotalResponse) WriteTo(writable types.Writable) {
	m.SearchResults.WriteTo(writable)
	m.TotalResults.WriteTo(writable)
}

// ExtractFrom extracts the SearchContentsWithTotalResponse from the given readable
func (m *SearchContentsWithTotalResponse) ExtractFrom(readable types.Readable) error {
	if err := m.SearchResults.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SearchContentsWithTotal.searchResults. %s", err.Error())
	}

	if err := m.TotalResults.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SearchContentsWithTotal.totalResults. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleSearchContentsWithTotal(packet nex.PacketInterface) {
	if protocol.SearchContentsWithTotal == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "User Storage::SearchContentsWithTotal not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params SearchContentsWithTotalRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.SearchContentsWithTotal(err, packet, callID, params.Query)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.SearchContentsWithTotal(nil, packet, callID, params.Query)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
