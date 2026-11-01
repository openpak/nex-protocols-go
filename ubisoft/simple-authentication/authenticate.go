package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleAuthenticate(packet nex.PacketInterface) {
	if protocol.Authenticate == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "SimpleAuthentication::Authenticate not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	strUserName, err := parseAuthenticateRequest(parametersStream)
	if err != nil {
		_, rmcError := protocol.Authenticate(err, packet, callID, strUserName)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.Authenticate(nil, packet, callID, strUserName)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

// parseAuthenticateRequest decodes the parameters of Simple Authentication::Authenticate from the given stream
func parseAuthenticateRequest(stream *nex.ByteStreamIn) (types.String, error) {
	var strUserName types.String

	if err := strUserName.ExtractFrom(stream); err != nil {
		return strUserName, fmt.Errorf("failed to read strUserName from parameters. %s", err.Error())
	}

	return strUserName, nil
}
