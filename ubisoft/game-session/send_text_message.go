package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	game_session_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/game-session/types"
)

// SendTextMessageRequest holds the request parameters of Game Session::SendTextMessage
type SendTextMessageRequest struct {
	GameSessionMessage game_session_types.GameSessionMessage
}

// WriteTo writes the SendTextMessageRequest to the given writable
func (m SendTextMessageRequest) WriteTo(writable types.Writable) {
	m.GameSessionMessage.WriteTo(writable)
}

// ExtractFrom extracts the SendTextMessageRequest from the given readable
func (m *SendTextMessageRequest) ExtractFrom(readable types.Readable) error {
	if err := m.GameSessionMessage.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SendTextMessage.gameSessionMessage. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleSendTextMessage(packet nex.PacketInterface) {
	if protocol.SendTextMessage == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Game Session::SendTextMessage not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params SendTextMessageRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.SendTextMessage(err, packet, callID, params.GameSessionMessage)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.SendTextMessage(nil, packet, callID, params.GameSessionMessage)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
