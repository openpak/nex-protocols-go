package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleLoginWithTokenCafeEx(packet nex.PacketInterface) {
	if protocol.LoginWithTokenCafeEx == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "SimpleAuthentication::LoginWithTokenCafeEx not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	strNintendoToken, oAnyData, err := parseLoginWithTokenCafeExRequest(parametersStream)
	if err != nil {
		_, rmcError := protocol.LoginWithTokenCafeEx(err, packet, callID, strNintendoToken, oAnyData)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.LoginWithTokenCafeEx(nil, packet, callID, strNintendoToken, oAnyData)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

// parseLoginWithTokenCafeExRequest decodes the parameters of Simple Authentication::LoginWithTokenCafeEx from the given stream
func parseLoginWithTokenCafeExRequest(stream *nex.ByteStreamIn) (types.String, types.DataHolder, error) {
	var strNintendoToken types.String
	var oAnyData types.DataHolder

	if err := strNintendoToken.ExtractFrom(stream); err != nil {
		return strNintendoToken, oAnyData, fmt.Errorf("failed to read strNintendoToken from parameters. %s", err.Error())
	}

	if err := oAnyData.ExtractFrom(stream); err != nil {
		return strNintendoToken, oAnyData, fmt.Errorf("failed to read oAnyData from parameters. %s", err.Error())
	}

	return strNintendoToken, oAnyData, nil
}
