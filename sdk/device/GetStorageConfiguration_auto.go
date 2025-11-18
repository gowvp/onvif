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

// Call_GetStorageConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a GetStorageConfigurationResponse.
func Call_GetStorageConfiguration(ctx context.Context, dev *onvif.Device, request device.GetStorageConfiguration) (device.GetStorageConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetStorageConfigurationResponse device.GetStorageConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetStorageConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetStorageConfiguration")
		return reply.Body.GetStorageConfigurationResponse, errors.Annotate(err, "reply")
	}
}
