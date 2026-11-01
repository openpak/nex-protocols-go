package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	game_session_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/game-session/types"
)

// SearchSessionsRequest holds the request parameters of Game Session::SearchSessions
type SearchSessionsRequest struct {
	GameSessionQuery game_session_types.GameSessionQuery
}

// WriteTo writes the SearchSessionsRequest to the given writable
func (m SearchSessionsRequest) WriteTo(writable types.Writable) {
	m.GameSessionQuery.WriteTo(writable)
}

// ExtractFrom extracts the SearchSessionsRequest from the given readable
func (m *SearchSessionsRequest) ExtractFrom(readable types.Readable) error {
	if err := m.GameSessionQuery.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SearchSessions.gameSessionQuery. %s", err.Error())
	}

	return nil
}

// SearchSessionsResponse holds the response parameters of Game Session::SearchSessions
type SearchSessionsResponse struct {
	SearchResults types.List[game_session_types.GameSessionSearchResult]
}

// WriteTo writes the SearchSessionsResponse to the given writable
func (m SearchSessionsResponse) WriteTo(writable types.Writable) {
	m.SearchResults.WriteTo(writable)
}

// ExtractFrom extracts the SearchSessionsResponse from the given readable
func (m *SearchSessionsResponse) ExtractFrom(readable types.Readable) error {
	if err := m.SearchResults.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SearchSessions.searchResults. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleSearchSessions(packet nex.PacketInterface) {
	if protocol.SearchSessions == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Game Session::SearchSessions not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params SearchSessionsRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.SearchSessions(err, packet, callID, params.GameSessionQuery)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.SearchSessions(nil, packet, callID, params.GameSessionQuery)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
