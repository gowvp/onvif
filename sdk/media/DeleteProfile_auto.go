// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package media

import (
	"context"

	"github.com/gowvp/onvif"
	"github.com/gowvp/onvif/media"
	"github.com/gowvp/onvif/sdk"
	"github.com/juju/errors"
)

// Call_DeleteProfile forwards the call to dev.CallMethod() then parses the payload of the reply as a DeleteProfileResponse.
func Call_DeleteProfile(ctx context.Context, dev *onvif.Device, request media.DeleteProfile) (media.DeleteProfileResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			DeleteProfileResponse media.DeleteProfileResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.DeleteProfileResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "DeleteProfile")
		return reply.Body.DeleteProfileResponse, errors.Annotate(err, "reply")
	}
}
