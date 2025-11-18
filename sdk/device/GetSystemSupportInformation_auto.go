// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package device

import (
	"context"

	"github.com/gowvp/onvif"
	"github.com/gowvp/onvif/device"
	"github.com/gowvp/onvif/sdk"
	"github.com/juju/errors"
)

// Call_GetSystemSupportInformation forwards the call to dev.CallMethod() then parses the payload of the reply as a GetSystemSupportInformationResponse.
func Call_GetSystemSupportInformation(ctx context.Context, dev *onvif.Device, request device.GetSystemSupportInformation) (device.GetSystemSupportInformationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetSystemSupportInformationResponse device.GetSystemSupportInformationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetSystemSupportInformationResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetSystemSupportInformation")
		return reply.Body.GetSystemSupportInformationResponse, errors.Annotate(err, "reply")
	}
}
