package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

// GetInvitationSentCountRequest holds the request parameters of Game Session::GetInvitationSentCount
type GetInvitationSentCountRequest struct {
	GameSessionTypeID types.UInt32
}

// WriteTo writes the GetInvitationSentCountRequest to the given writable
func (m GetInvitationSentCountRequest) WriteTo(writable types.Writable) {
	m.GameSessionTypeID.WriteTo(writable)
}

// ExtractFrom extracts the GetInvitationSentCountRequest from the given readable
func (m *GetInvitationSentCountRequest) ExtractFrom(readable types.Readable) error {
	if err := m.GameSessionTypeID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetInvitationSentCount.gameSessionTypeID. %s", err.Error())
	}

	return nil
}

// GetInvitationSentCountResponse holds the response parameters of Game Session::GetInvitationSentCount
type GetInvitationSentCountResponse struct {
	Count types.UInt32
}

// WriteTo writes the GetInvitationSentCountResponse to the given writable
func (m GetInvitationSentCountResponse) WriteTo(writable types.Writable) {
	m.Count.WriteTo(writable)
}

// ExtractFrom extracts the GetInvitationSentCountResponse from the given readable
func (m *GetInvitationSentCountResponse) ExtractFrom(readable types.Readable) error {
	if err := m.Count.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetInvitationSentCount.count. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleGetInvitationSentCount(packet nex.PacketInterface) {
	if protocol.GetInvitationSentCount == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Game Session::GetInvitationSentCount not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params GetInvitationSentCountRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.GetInvitationSentCount(err, packet, callID, params.GameSessionTypeID)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetInvitationSentCount(nil, packet, callID, params.GameSessionTypeID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
