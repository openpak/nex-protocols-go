package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleLinkAccount(packet nex.PacketInterface) {
	if protocol.LinkAccount == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "UbiAccountManagement::LinkAccount not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	ubiAccountUsername, ubiAccountPassword, err := parseLinkAccountRequest(parametersStream)
	if err != nil {
		_, rmcError := protocol.LinkAccount(err, packet, callID, ubiAccountUsername, ubiAccountPassword)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.LinkAccount(nil, packet, callID, ubiAccountUsername, ubiAccountPassword)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

// parseLinkAccountRequest decodes the parameters of Ubi Account Management::LinkAccount from the given stream
func parseLinkAccountRequest(stream *nex.ByteStreamIn) (types.String, types.String, error) {
	var ubiAccountUsername types.String
	var ubiAccountPassword types.String

	if err := ubiAccountUsername.ExtractFrom(stream); err != nil {
		return ubiAccountUsername, ubiAccountPassword, fmt.Errorf("failed to read ubiAccountUsername from parameters. %s", err.Error())
	}

	if err := ubiAccountPassword.ExtractFrom(stream); err != nil {
		return ubiAccountUsername, ubiAccountPassword, fmt.Errorf("failed to read ubiAccountPassword from parameters. %s", err.Error())
	}

	return ubiAccountUsername, ubiAccountPassword, nil
}
