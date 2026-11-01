package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleSetActionCompleted(packet nex.PacketInterface) {
	if protocol.SetActionCompleted == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "UplayWin::SetActionCompleted not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	actionCode, cultureName, platformCode, gameCode, err := parseSetActionCompletedRequest(parametersStream)
	if err != nil {
		_, rmcError := protocol.SetActionCompleted(err, packet, callID, actionCode, cultureName, platformCode, gameCode)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.SetActionCompleted(nil, packet, callID, actionCode, cultureName, platformCode, gameCode)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

// parseSetActionCompletedRequest decodes the parameters of Uplay Win::SetActionCompleted from the given stream
func parseSetActionCompletedRequest(stream *nex.ByteStreamIn) (types.String, types.String, types.String, types.String, error) {
	var actionCode types.String
	var cultureName types.String
	var platformCode types.String
	var gameCode types.String

	if err := actionCode.ExtractFrom(stream); err != nil {
		return actionCode, cultureName, platformCode, gameCode, fmt.Errorf("failed to read actionCode from parameters. %s", err.Error())
	}

	if err := cultureName.ExtractFrom(stream); err != nil {
		return actionCode, cultureName, platformCode, gameCode, fmt.Errorf("failed to read cultureName from parameters. %s", err.Error())
	}

	if err := platformCode.ExtractFrom(stream); err != nil {
		return actionCode, cultureName, platformCode, gameCode, fmt.Errorf("failed to read platformCode from parameters. %s", err.Error())
	}

	if err := gameCode.ExtractFrom(stream); err != nil {
		return actionCode, cultureName, platformCode, gameCode, fmt.Errorf("failed to read gameCode from parameters. %s", err.Error())
	}

	return actionCode, cultureName, platformCode, gameCode, nil
}
