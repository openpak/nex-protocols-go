package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	tracking_3_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/tracking-3/types"
)

// SendUserInfoRequest holds the request parameters of Tracking 3::SendUserInfo
type SendUserInfoRequest struct {
	UserInfo  tracking_3_types.TrackingInformation
	DeltaTime types.UInt32
}

// WriteTo writes the SendUserInfoRequest to the given writable
func (m SendUserInfoRequest) WriteTo(writable types.Writable) {
	m.UserInfo.WriteTo(writable)
	m.DeltaTime.WriteTo(writable)
}

// ExtractFrom extracts the SendUserInfoRequest from the given readable
func (m *SendUserInfoRequest) ExtractFrom(readable types.Readable) error {
	if err := m.UserInfo.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SendUserInfo.userInfo. %s", err.Error())
	}

	if err := m.DeltaTime.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SendUserInfo.deltaTime. %s", err.Error())
	}

	return nil
}

// SendUserInfoResponse holds the response parameters of Tracking 3::SendUserInfo
type SendUserInfoResponse struct {
	UserInfo   tracking_3_types.TrackingInformation
	TrackingID types.UInt32
}

// WriteTo writes the SendUserInfoResponse to the given writable
func (m SendUserInfoResponse) WriteTo(writable types.Writable) {
	m.UserInfo.WriteTo(writable)
	m.TrackingID.WriteTo(writable)
}

// ExtractFrom extracts the SendUserInfoResponse from the given readable
func (m *SendUserInfoResponse) ExtractFrom(readable types.Readable) error {
	if err := m.UserInfo.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SendUserInfo.userInfo. %s", err.Error())
	}

	if err := m.TrackingID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SendUserInfo.trackingID. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleSendUserInfo(packet nex.PacketInterface) {
	if protocol.SendUserInfo == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Tracking 3::SendUserInfo not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params SendUserInfoRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.SendUserInfo(err, packet, callID, params.UserInfo, params.DeltaTime)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.SendUserInfo(nil, packet, callID, params.UserInfo, params.DeltaTime)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
