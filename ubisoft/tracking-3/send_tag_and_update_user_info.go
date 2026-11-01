package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

// SendTagAndUpdateUserInfoRequest holds the request parameters of Tracking 3::SendTagAndUpdateUserInfo
type SendTagAndUpdateUserInfoRequest struct {
	TrackingID types.UInt32
	Tag        types.String
	Attributes types.String
	DeltaTime  types.UInt32
	UserID     types.String
}

// WriteTo writes the SendTagAndUpdateUserInfoRequest to the given writable
func (m SendTagAndUpdateUserInfoRequest) WriteTo(writable types.Writable) {
	m.TrackingID.WriteTo(writable)
	m.Tag.WriteTo(writable)
	m.Attributes.WriteTo(writable)
	m.DeltaTime.WriteTo(writable)
	m.UserID.WriteTo(writable)
}

// ExtractFrom extracts the SendTagAndUpdateUserInfoRequest from the given readable
func (m *SendTagAndUpdateUserInfoRequest) ExtractFrom(readable types.Readable) error {
	if err := m.TrackingID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SendTagAndUpdateUserInfo.trackingID. %s", err.Error())
	}

	if err := m.Tag.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SendTagAndUpdateUserInfo.tag. %s", err.Error())
	}

	if err := m.Attributes.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SendTagAndUpdateUserInfo.attributes. %s", err.Error())
	}

	if err := m.DeltaTime.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SendTagAndUpdateUserInfo.deltaTime. %s", err.Error())
	}

	if err := m.UserID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SendTagAndUpdateUserInfo.userID. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleSendTagAndUpdateUserInfo(packet nex.PacketInterface) {
	if protocol.SendTagAndUpdateUserInfo == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Tracking 3::SendTagAndUpdateUserInfo not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params SendTagAndUpdateUserInfoRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.SendTagAndUpdateUserInfo(err, packet, callID, params.TrackingID, params.Tag, params.Attributes, params.DeltaTime, params.UserID)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.SendTagAndUpdateUserInfo(nil, packet, callID, params.TrackingID, params.Tag, params.Attributes, params.DeltaTime, params.UserID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
