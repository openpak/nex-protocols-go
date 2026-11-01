package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	user_storage_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/user-storage/types"
)

// SearchContentsByPlayersRequest holds the request parameters of User Storage::SearchContentsByPlayers
type SearchContentsByPlayersRequest struct {
	Pids  types.List[types.UInt32]
	Query user_storage_types.UserStorageQuery
}

// WriteTo writes the SearchContentsByPlayersRequest to the given writable
func (m SearchContentsByPlayersRequest) WriteTo(writable types.Writable) {
	m.Pids.WriteTo(writable)
	m.Query.WriteTo(writable)
}

// ExtractFrom extracts the SearchContentsByPlayersRequest from the given readable
func (m *SearchContentsByPlayersRequest) ExtractFrom(readable types.Readable) error {
	if err := m.Pids.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SearchContentsByPlayers.pids. %s", err.Error())
	}

	if err := m.Query.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SearchContentsByPlayers.query. %s", err.Error())
	}

	return nil
}

// SearchContentsByPlayersResponse holds the response parameters of User Storage::SearchContentsByPlayers
type SearchContentsByPlayersResponse struct {
	SearchResults types.List[user_storage_types.UserContent]
}

// WriteTo writes the SearchContentsByPlayersResponse to the given writable
func (m SearchContentsByPlayersResponse) WriteTo(writable types.Writable) {
	m.SearchResults.WriteTo(writable)
}

// ExtractFrom extracts the SearchContentsByPlayersResponse from the given readable
func (m *SearchContentsByPlayersResponse) ExtractFrom(readable types.Readable) error {
	if err := m.SearchResults.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SearchContentsByPlayers.searchResults. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleSearchContentsByPlayers(packet nex.PacketInterface) {
	if protocol.SearchContentsByPlayers == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "User Storage::SearchContentsByPlayers not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params SearchContentsByPlayersRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.SearchContentsByPlayers(err, packet, callID, params.Pids, params.Query)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.SearchContentsByPlayers(nil, packet, callID, params.Pids, params.Query)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
