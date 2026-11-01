package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	ubi_account_management_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/ubi-account-management/types"
)

func (protocol *Protocol) handleCreateAccount(packet nex.PacketInterface) {
	if protocol.CreateAccount == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "UbiAccountManagement::CreateAccount not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	ubiAccount, err := parseCreateAccountRequest(parametersStream)
	if err != nil {
		_, rmcError := protocol.CreateAccount(err, packet, callID, ubiAccount)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.CreateAccount(nil, packet, callID, ubiAccount)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

// parseCreateAccountRequest decodes the parameters of Ubi Account Management::CreateAccount from the given stream
func parseCreateAccountRequest(stream *nex.ByteStreamIn) (ubi_account_management_types.UbiAccount, error) {
	var ubiAccount ubi_account_management_types.UbiAccount

	if err := ubiAccount.ExtractFrom(stream); err != nil {
		return ubiAccount, fmt.Errorf("failed to read ubiAccount from parameters. %s", err.Error())
	}

	return ubiAccount, nil
}
