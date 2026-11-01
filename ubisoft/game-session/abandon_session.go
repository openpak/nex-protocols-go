package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	game_session_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/game-session/types"
)

// AbandonSessionRequest holds the request parameters of Game Session::AbandonSession
type AbandonSessionRequest struct {
	GameSessionKey game_session_types.GameSessionKey
}

// WriteTo writes the AbandonSessionRequest to the given writable
func (m AbandonSessionRequest) WriteTo(writable types.Writable) {
	m.GameSessionKey.WriteTo(writable)
}

// ExtractFrom extracts the AbandonSessionRequest from the given readable
func (m *AbandonSessionRequest) ExtractFrom(readable types.Readable) error {
	if err := m.GameSessionKey.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read AbandonSession.gameSessionKey. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleAbandonSession(packet nex.PacketInterface) {
	if protocol.AbandonSession == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Game Session::AbandonSession not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params AbandonSessionRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.AbandonSession(err, packet, callID, params.GameSessionKey)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.AbandonSession(nil, packet, callID, params.GameSessionKey)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
