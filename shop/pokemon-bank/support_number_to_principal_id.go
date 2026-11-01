// Package protocol implements the Pokemon Bank Shop protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

// extractSupportNumberToPrincipalIDParameters reads the SupportNumberToPrincipalID request parameters from the given stream
func extractSupportNumberToPrincipalIDParameters(stream *nex.ByteStreamIn) (supportNumber types.String, err error) {
	if err = supportNumber.ExtractFrom(stream); err != nil {
		err = fmt.Errorf("failed to read supportNumber from parameters. %s", err.Error())
		return
	}

	return
}

func (protocol *Protocol) handleSupportNumberToPrincipalID(packet nex.PacketInterface) {
	if protocol.SupportNumberToPrincipalID == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "ShopPokemonBank::SupportNumberToPrincipalID not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	supportNumber, err := extractSupportNumberToPrincipalIDParameters(parametersStream)
	if err != nil {
		_, rmcError := protocol.SupportNumberToPrincipalID(err, packet, callID, supportNumber)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.SupportNumberToPrincipalID(nil, packet, callID, supportNumber)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
