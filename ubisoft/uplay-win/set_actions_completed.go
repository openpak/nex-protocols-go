package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleSetActionsCompleted(packet nex.PacketInterface) {
	if protocol.SetActionsCompleted == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "UplayWin::SetActionsCompleted not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	actionCodeList, cultureName, platformCode, gameCode, err := parseSetActionsCompletedRequest(parametersStream)
	if err != nil {
		_, rmcError := protocol.SetActionsCompleted(err, packet, callID, actionCodeList, cultureName, platformCode, gameCode)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.SetActionsCompleted(nil, packet, callID, actionCodeList, cultureName, platformCode, gameCode)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

// parseSetActionsCompletedRequest decodes the parameters of Uplay Win::SetActionsCompleted from the given stream
func parseSetActionsCompletedRequest(stream *nex.ByteStreamIn) (types.List[types.String], types.String, types.String, types.String, error) {
	var actionCodeList types.List[types.String]
	var cultureName types.String
	var platformCode types.String
	var gameCode types.String

	if err := actionCodeList.ExtractFrom(stream); err != nil {
		return actionCodeList, cultureName, platformCode, gameCode, fmt.Errorf("failed to read actionCodeList from parameters. %s", err.Error())
	}

	if err := cultureName.ExtractFrom(stream); err != nil {
		return actionCodeList, cultureName, platformCode, gameCode, fmt.Errorf("failed to read cultureName from parameters. %s", err.Error())
	}

	if err := platformCode.ExtractFrom(stream); err != nil {
		return actionCodeList, cultureName, platformCode, gameCode, fmt.Errorf("failed to read platformCode from parameters. %s", err.Error())
	}

	if err := gameCode.ExtractFrom(stream); err != nil {
		return actionCodeList, cultureName, platformCode, gameCode, fmt.Errorf("failed to read gameCode from parameters. %s", err.Error())
	}

	return actionCodeList, cultureName, platformCode, gameCode, nil
}
