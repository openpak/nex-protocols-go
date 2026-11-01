package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleGetVirtualCurrencyUserBalance(packet nex.PacketInterface) {
	if protocol.GetVirtualCurrencyUserBalance == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "UplayWin::GetVirtualCurrencyUserBalance not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	platformCode, err := parseGetVirtualCurrencyUserBalanceRequest(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetVirtualCurrencyUserBalance(err, packet, callID, platformCode)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetVirtualCurrencyUserBalance(nil, packet, callID, platformCode)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

// parseGetVirtualCurrencyUserBalanceRequest decodes the parameters of Uplay Win::GetVirtualCurrencyUserBalance from the given stream
func parseGetVirtualCurrencyUserBalanceRequest(stream *nex.ByteStreamIn) (types.String, error) {
	var platformCode types.String

	if err := platformCode.ExtractFrom(stream); err != nil {
		return platformCode, fmt.Errorf("failed to read platformCode from parameters. %s", err.Error())
	}

	return platformCode, nil
}
