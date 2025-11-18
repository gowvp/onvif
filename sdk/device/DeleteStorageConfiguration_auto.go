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

// Call_DeleteStorageConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a DeleteStorageConfigurationResponse.
func Call_DeleteStorageConfiguration(ctx context.Context, dev *onvif.Device, request device.DeleteStorageConfiguration) (device.DeleteStorageConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			DeleteStorageConfigurationResponse device.DeleteStorageConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.DeleteStorageConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "DeleteStorageConfiguration")
		return reply.Body.DeleteStorageConfigurationResponse, errors.Annotate(err, "reply")
	}
}
