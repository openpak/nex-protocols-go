// Package protocol implements the RankingSuperSmashBros4 protocol
//
// The Super Smash Bros. 4 specific methods (14-20) replace the base Ranking
// methods with the same IDs. The wiki only documents their names, not their
// parameters, so every handler receives the raw payload (layout provisional:
// no parameter tables exist for these methods).
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	ranking "github.com/PretendoNetwork/nex-protocols-go/v2/ranking"
	"golang.org/x/exp/slices"
)

const (
	// ProtocolID is the Protocol ID for the Ranking (Super Smash Bros. 4) protocol
	ProtocolID = 0x70

	// MethodGetCommunityCompetitionScore is the method ID for the GetCommunityCompetitionScore method
	MethodGetCommunityCompetitionScore = 0xE

	// MethodGetCommunityCompetitionRankData is the method ID for the GetCommunityCompetitionRankData method
	MethodGetCommunityCompetitionRankData = 0xF

	// MethodGetCommunityCompetitionRanking is the method ID for the GetCommunityCompetitionRanking method
	MethodGetCommunityCompetitionRanking = 0x10

	// MethodPutCommunityCompetitionScore is the method ID for the PutCommunityCompetitionScore method
	MethodPutCommunityCompetitionScore = 0x11

	// MethodDeleteCommunityCompetitionScore is the method ID for the DeleteCommunityCompetitionScore method
	MethodDeleteCommunityCompetitionScore = 0x12

	// MethodDeleteCommunityCompetitionScores is the method ID for the DeleteCommunityCompetitionScores method
	MethodDeleteCommunityCompetitionScores = 0x13

	// MethodDEBUG_PutCommunityCompetitionScores is the method ID for the DEBUG_PutCommunityCompetitionScores method
	MethodDEBUG_PutCommunityCompetitionScores = 0x14
)

var patchedMethods = []uint32{
	MethodGetCommunityCompetitionScore,
	MethodGetCommunityCompetitionRankData,
	MethodGetCommunityCompetitionRanking,
	MethodPutCommunityCompetitionScore,
	MethodDeleteCommunityCompetitionScore,
	MethodDeleteCommunityCompetitionScores,
	MethodDEBUG_PutCommunityCompetitionScores,
}

type rankingProtocol = ranking.Protocol

// Protocol stores all the RMC method handlers for the Ranking (Super Smash Bros. 4) protocol and listens for requests
// Embeds the Ranking protocol
type Protocol struct {
	endpoint nex.EndpointInterface
	rankingProtocol
	GetCommunityCompetitionScore        func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error)
	GetCommunityCompetitionRankData     func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error)
	GetCommunityCompetitionRanking      func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error)
	PutCommunityCompetitionScore        func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error)
	DeleteCommunityCompetitionScore     func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error)
	DeleteCommunityCompetitionScores    func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error)
	DEBUG_PutCommunityCompetitionScores func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error)
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	message := packet.RMCMessage()

	if !message.IsRequest || message.ProtocolID != ProtocolID {
		return
	}

	if !slices.Contains(patchedMethods, message.MethodID) {
		protocol.rankingProtocol.HandlePacket(packet)
		return
	}

	switch message.MethodID {
	case MethodGetCommunityCompetitionScore:
		protocol.handleGetCommunityCompetitionScore(packet)
	case MethodGetCommunityCompetitionRankData:
		protocol.handleGetCommunityCompetitionRankData(packet)
	case MethodGetCommunityCompetitionRanking:
		protocol.handleGetCommunityCompetitionRanking(packet)
	case MethodPutCommunityCompetitionScore:
		protocol.handlePutCommunityCompetitionScore(packet)
	case MethodDeleteCommunityCompetitionScore:
		protocol.handleDeleteCommunityCompetitionScore(packet)
	case MethodDeleteCommunityCompetitionScores:
		protocol.handleDeleteCommunityCompetitionScores(packet)
	case MethodDEBUG_PutCommunityCompetitionScores:
		protocol.handleDEBUG_PutCommunityCompetitionScores(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported Ranking (Super Smash Bros. 4) method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new RankingSuperSmashBros4 protocol
func NewProtocol(endpoint nex.EndpointInterface) *Protocol {
	protocol := &Protocol{endpoint: endpoint}
	protocol.SetEndpoint(endpoint)

	return protocol
}
