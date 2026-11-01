package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleGetRewards(packet nex.PacketInterface) {
	if protocol.GetRewards == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "UplayWin::GetRewards not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	startRowIndex, maximumRows, sortExpression, cultureName, platformCode, gameCode, err := parseGetRewardsRequest(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetRewards(err, packet, callID, startRowIndex, maximumRows, sortExpression, cultureName, platformCode, gameCode)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetRewards(nil, packet, callID, startRowIndex, maximumRows, sortExpression, cultureName, platformCode, gameCode)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

// parseGetRewardsRequest decodes the parameters of Uplay Win::GetRewards from the given stream
func parseGetRewardsRequest(stream *nex.ByteStreamIn) (types.Int32, types.Int32, types.String, types.String, types.String, types.String, error) {
	var startRowIndex types.Int32
	var maximumRows types.Int32
	var sortExpression types.String
	var cultureName types.String
	var platformCode types.String
	var gameCode types.String

	if err := startRowIndex.ExtractFrom(stream); err != nil {
		return startRowIndex, maximumRows, sortExpression, cultureName, platformCode, gameCode, fmt.Errorf("failed to read startRowIndex from parameters. %s", err.Error())
	}

	if err := maximumRows.ExtractFrom(stream); err != nil {
		return startRowIndex, maximumRows, sortExpression, cultureName, platformCode, gameCode, fmt.Errorf("failed to read maximumRows from parameters. %s", err.Error())
	}

	if err := sortExpression.ExtractFrom(stream); err != nil {
		return startRowIndex, maximumRows, sortExpression, cultureName, platformCode, gameCode, fmt.Errorf("failed to read sortExpression from parameters. %s", err.Error())
	}

	if err := cultureName.ExtractFrom(stream); err != nil {
		return startRowIndex, maximumRows, sortExpression, cultureName, platformCode, gameCode, fmt.Errorf("failed to read cultureName from parameters. %s", err.Error())
	}

	if err := platformCode.ExtractFrom(stream); err != nil {
		return startRowIndex, maximumRows, sortExpression, cultureName, platformCode, gameCode, fmt.Errorf("failed to read platformCode from parameters. %s", err.Error())
	}

	if err := gameCode.ExtractFrom(stream); err != nil {
		return startRowIndex, maximumRows, sortExpression, cultureName, platformCode, gameCode, fmt.Errorf("failed to read gameCode from parameters. %s", err.Error())
	}

	return startRowIndex, maximumRows, sortExpression, cultureName, platformCode, gameCode, nil
}
