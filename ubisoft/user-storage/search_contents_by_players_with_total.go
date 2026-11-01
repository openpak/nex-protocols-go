package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	user_storage_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/user-storage/types"
)

// SearchContentsByPlayersWithTotalRequest holds the request parameters of User Storage::SearchContentsByPlayersWithTotal
type SearchContentsByPlayersWithTotalRequest struct {
	Pids  types.List[types.UInt32]
	Query user_storage_types.UserStorageQuery
}

// WriteTo writes the SearchContentsByPlayersWithTotalRequest to the given writable
func (m SearchContentsByPlayersWithTotalRequest) WriteTo(writable types.Writable) {
	m.Pids.WriteTo(writable)
	m.Query.WriteTo(writable)
}

// ExtractFrom extracts the SearchContentsByPlayersWithTotalRequest from the given readable
func (m *SearchContentsByPlayersWithTotalRequest) ExtractFrom(readable types.Readable) error {
	if err := m.Pids.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SearchContentsByPlayersWithTotal.pids. %s", err.Error())
	}

	if err := m.Query.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SearchContentsByPlayersWithTotal.query. %s", err.Error())
	}

	return nil
}

// SearchContentsByPlayersWithTotalResponse holds the response parameters of User Storage::SearchContentsByPlayersWithTotal
type SearchContentsByPlayersWithTotalResponse struct {
	SearchResults types.List[user_storage_types.UserContent]
	TotalResults  types.UInt32
}

// WriteTo writes the SearchContentsByPlayersWithTotalResponse to the given writable
func (m SearchContentsByPlayersWithTotalResponse) WriteTo(writable types.Writable) {
	m.SearchResults.WriteTo(writable)
	m.TotalResults.WriteTo(writable)
}

// ExtractFrom extracts the SearchContentsByPlayersWithTotalResponse from the given readable
func (m *SearchContentsByPlayersWithTotalResponse) ExtractFrom(readable types.Readable) error {
	if err := m.SearchResults.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SearchContentsByPlayersWithTotal.searchResults. %s", err.Error())
	}

	if err := m.TotalResults.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SearchContentsByPlayersWithTotal.totalResults. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleSearchContentsByPlayersWithTotal(packet nex.PacketInterface) {
	if protocol.SearchContentsByPlayersWithTotal == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "User Storage::SearchContentsByPlayersWithTotal not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params SearchContentsByPlayersWithTotalRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.SearchContentsByPlayersWithTotal(err, packet, callID, params.Pids, params.Query)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.SearchContentsByPlayersWithTotal(nil, packet, callID, params.Pids, params.Query)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
