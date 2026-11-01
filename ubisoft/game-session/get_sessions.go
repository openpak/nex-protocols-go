package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	game_session_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/game-session/types"
)

// GetSessionsRequest holds the request parameters of Game Session::GetSessions
type GetSessionsRequest struct {
	GameSessionKeys types.List[game_session_types.GameSessionKey]
}

// WriteTo writes the GetSessionsRequest to the given writable
func (m GetSessionsRequest) WriteTo(writable types.Writable) {
	m.GameSessionKeys.WriteTo(writable)
}

// ExtractFrom extracts the GetSessionsRequest from the given readable
func (m *GetSessionsRequest) ExtractFrom(readable types.Readable) error {
	if err := m.GameSessionKeys.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetSessions.gameSessionKeys. %s", err.Error())
	}

	return nil
}

// GetSessionsResponse holds the response parameters of Game Session::GetSessions
type GetSessionsResponse struct {
	SearchResults types.List[game_session_types.GameSessionSearchResult]
}

// WriteTo writes the GetSessionsResponse to the given writable
func (m GetSessionsResponse) WriteTo(writable types.Writable) {
	m.SearchResults.WriteTo(writable)
}

// ExtractFrom extracts the GetSessionsResponse from the given readable
func (m *GetSessionsResponse) ExtractFrom(readable types.Readable) error {
	if err := m.SearchResults.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetSessions.searchResults. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleGetSessions(packet nex.PacketInterface) {
	if protocol.GetSessions == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Game Session::GetSessions not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params GetSessionsRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.GetSessions(err, packet, callID, params.GameSessionKeys)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetSessions(nil, packet, callID, params.GameSessionKeys)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
