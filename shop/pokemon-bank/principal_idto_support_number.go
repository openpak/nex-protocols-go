// Package protocol implements the Pokemon Bank Shop protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

// extractPrincipalIDToSupportNumberParameters reads the PrincipalIDToSupportNumber request parameters from the given stream
func extractPrincipalIDToSupportNumberParameters(stream *nex.ByteStreamIn) (pid types.UInt32, err error) {
	if err = pid.ExtractFrom(stream); err != nil {
		err = fmt.Errorf("failed to read pid from parameters. %s", err.Error())
		return
	}

	return
}

func (protocol *Protocol) handlePrincipalIDToSupportNumber(packet nex.PacketInterface) {
	if protocol.PrincipalIDToSupportNumber == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "ShopPokemonBank::PrincipalIDToSupportNumber not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	pid, err := extractPrincipalIDToSupportNumberParameters(parametersStream)
	if err != nil {
		_, rmcError := protocol.PrincipalIDToSupportNumber(err, packet, callID, pid)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.PrincipalIDToSupportNumber(nil, packet, callID, pid)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
