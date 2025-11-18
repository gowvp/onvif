// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package event

import (
	"context"

	"github.com/gowvp/onvif"
	"github.com/gowvp/onvif/event"
	"github.com/gowvp/onvif/sdk"
	"github.com/juju/errors"
)

// Call_CreatePullPointSubscription forwards the call to dev.CallMethod() then parses the payload of the reply as a CreatePullPointSubscriptionResponse.
func Call_CreatePullPointSubscription(ctx context.Context, dev *onvif.Device, request event.CreatePullPointSubscription) (event.CreatePullPointSubscriptionResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			CreatePullPointSubscriptionResponse event.CreatePullPointSubscriptionResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.CreatePullPointSubscriptionResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "CreatePullPointSubscription")
		return reply.Body.CreatePullPointSubscriptionResponse, errors.Annotate(err, "reply")
	}
}
