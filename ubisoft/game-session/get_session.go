package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	game_session_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/game-session/types"
)

// GetSessionRequest holds the request parameters of Game Session::GetSession
type GetSessionRequest struct {
	GameSessionKey game_session_types.GameSessionKey
}

// WriteTo writes the GetSessionRequest to the given writable
func (m GetSessionRequest) WriteTo(writable types.Writable) {
	m.GameSessionKey.WriteTo(writable)
}

// ExtractFrom extracts the GetSessionRequest from the given readable
func (m *GetSessionRequest) ExtractFrom(readable types.Readable) error {
	if err := m.GameSessionKey.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetSession.gameSessionKey. %s", err.Error())
	}

	return nil
}

// GetSessionResponse holds the response parameters of Game Session::GetSession
type GetSessionResponse struct {
	SearchResult game_session_types.GameSessionSearchResult
}

// WriteTo writes the GetSessionResponse to the given writable
func (m GetSessionResponse) WriteTo(writable types.Writable) {
	m.SearchResult.WriteTo(writable)
}

// ExtractFrom extracts the GetSessionResponse from the given readable
func (m *GetSessionResponse) ExtractFrom(readable types.Readable) error {
	if err := m.SearchResult.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetSession.searchResult. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleGetSession(packet nex.PacketInterface) {
	if protocol.GetSession == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Game Session::GetSession not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params GetSessionRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.GetSession(err, packet, callID, params.GameSessionKey)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetSession(nil, packet, callID, params.GameSessionKey)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
