package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleLogin(packet nex.PacketInterface) {
	if protocol.Login == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "SimpleAuthentication::Login not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	strUsername, strPassword, err := parseLoginRequest(parametersStream)
	if err != nil {
		_, rmcError := protocol.Login(err, packet, callID, strUsername, strPassword)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.Login(nil, packet, callID, strUsername, strPassword)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

// parseLoginRequest decodes the parameters of Simple Authentication::Login from the given stream
func parseLoginRequest(stream *nex.ByteStreamIn) (types.String, types.String, error) {
	var strUsername types.String
	var strPassword types.String

	if err := strUsername.ExtractFrom(stream); err != nil {
		return strUsername, strPassword, fmt.Errorf("failed to read strUsername from parameters. %s", err.Error())
	}

	if err := strPassword.ExtractFrom(stream); err != nil {
		return strUsername, strPassword, fmt.Errorf("failed to read strPassword from parameters. %s", err.Error())
	}

	return strUsername, strPassword, nil
}
