// Package protocol implements the RankingSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handlePutCommunityCompetitionScore(packet nex.PacketInterface) {
	if protocol.PutCommunityCompetitionScore == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "RankingSuperSmashBros4::PutCommunityCompetitionScore not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	globals.Logger.Warning("RankingSuperSmashBros4::PutCommunityCompetitionScore STUBBED")

	request := packet.RMCMessage()
	callID := request.CallID

	rmcMessage, rmcError := protocol.PutCommunityCompetitionScore(nil, packet, callID, packet.Payload())
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
