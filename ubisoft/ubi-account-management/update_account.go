package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	ubi_account_management_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/ubi-account-management/types"
)

func (protocol *Protocol) handleUpdateAccount(packet nex.PacketInterface) {
	if protocol.UpdateAccount == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "UbiAccountManagement::UpdateAccount not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	ubiAccount, err := parseUpdateAccountRequest(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateAccount(err, packet, callID, ubiAccount)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.UpdateAccount(nil, packet, callID, ubiAccount)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

// parseUpdateAccountRequest decodes the parameters of Ubi Account Management::UpdateAccount from the given stream
func parseUpdateAccountRequest(stream *nex.ByteStreamIn) (ubi_account_management_types.UbiAccount, error) {
	var ubiAccount ubi_account_management_types.UbiAccount

	if err := ubiAccount.ExtractFrom(stream); err != nil {
		return ubiAccount, fmt.Errorf("failed to read ubiAccount from parameters. %s", err.Error())
	}

	return ubiAccount, nil
}
