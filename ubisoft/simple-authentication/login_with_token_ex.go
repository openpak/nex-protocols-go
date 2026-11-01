package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleLoginWithTokenEx(packet nex.PacketInterface) {
	if protocol.LoginWithTokenEx == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "SimpleAuthentication::LoginWithTokenEx not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	strToken, oAnyData, err := parseLoginWithTokenExRequest(parametersStream)
	if err != nil {
		_, rmcError := protocol.LoginWithTokenEx(err, packet, callID, strToken, oAnyData)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.LoginWithTokenEx(nil, packet, callID, strToken, oAnyData)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

// parseLoginWithTokenExRequest decodes the parameters of Simple Authentication::LoginWithTokenEx from the given stream
func parseLoginWithTokenExRequest(stream *nex.ByteStreamIn) (types.String, types.DataHolder, error) {
	var strToken types.String
	var oAnyData types.DataHolder

	if err := strToken.ExtractFrom(stream); err != nil {
		return strToken, oAnyData, fmt.Errorf("failed to read strToken from parameters. %s", err.Error())
	}

	if err := oAnyData.ExtractFrom(stream); err != nil {
		return strToken, oAnyData, fmt.Errorf("failed to read oAnyData from parameters. %s", err.Error())
	}

	return strToken, oAnyData, nil
}
