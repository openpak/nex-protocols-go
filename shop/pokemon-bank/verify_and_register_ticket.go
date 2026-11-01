// Package protocol implements the Pokemon Bank Shop protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

// extractVerifyAndRegisterTicketParameters reads the VerifyAndRegisterTicket request parameters from the given stream
func extractVerifyAndRegisterTicketParameters(stream *nex.ByteStreamIn) (ticketEnvelope types.QBuffer, purchasedTime types.DateTime, err error) {
	if err = ticketEnvelope.ExtractFrom(stream); err != nil {
		err = fmt.Errorf("failed to read ticketEnvelope from parameters. %s", err.Error())
		return
	}

	if err = purchasedTime.ExtractFrom(stream); err != nil {
		err = fmt.Errorf("failed to read purchasedTime from parameters. %s", err.Error())
		return
	}

	return
}

func (protocol *Protocol) handleVerifyAndRegisterTicket(packet nex.PacketInterface) {
	if protocol.VerifyAndRegisterTicket == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "ShopPokemonBank::VerifyAndRegisterTicket not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	ticketEnvelope, purchasedTime, err := extractVerifyAndRegisterTicketParameters(parametersStream)
	if err != nil {
		_, rmcError := protocol.VerifyAndRegisterTicket(err, packet, callID, ticketEnvelope, purchasedTime)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.VerifyAndRegisterTicket(nil, packet, callID, ticketEnvelope, purchasedTime)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
