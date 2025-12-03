// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package device

import (
	"context"

	"github.com/gowvp/onvif"
	"github.com/gowvp/onvif/device"
	"github.com/gowvp/onvif/media"
	"github.com/gowvp/onvif/sdk"
	"github.com/juju/errors"
)

// Call_GetDeviceInformation forwards the call to dev.CallMethod() then parses the payload of the reply as a GetDeviceInformationResponse.
func Call_GetDeviceInformation(ctx context.Context, dev *onvif.Device, request device.GetDeviceInformation) (device.GetDeviceInformationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetDeviceInformationResponse device.GetDeviceInformationResponse
			Fault                        media.GetDeviceInformationFault
		}
	}
	var reply Envelope
	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetDeviceInformationResponse, errors.Annotate(err, "call")
	}
	if err := sdk.ReadAndParse(ctx, httpReply, &reply, "GetDeviceInformation"); err != nil {
		return reply.Body.GetDeviceInformationResponse, errors.Annotate(err, "reply")
	}
	if httpReply.StatusCode == 200 {
		return reply.Body.GetDeviceInformationResponse, nil
	}
	return reply.Body.GetDeviceInformationResponse, errors.New(reply.Body.Fault.Reason.Text)
}
