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

// Call_GetStorageConfigurations forwards the call to dev.CallMethod() then parses the payload of the reply as a GetStorageConfigurationsResponse.
func Call_GetStorageConfigurations(ctx context.Context, dev *onvif.Device, request device.GetStorageConfigurations) (device.GetStorageConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetStorageConfigurationsResponse device.GetStorageConfigurationsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetStorageConfigurationsResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetStorageConfigurations")
		return reply.Body.GetStorageConfigurationsResponse, errors.Annotate(err, "reply")
	}
}
