// Copyright (c) 2017 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"time"
)

// MessageType represents the type of a VPP message.
// Note: this is currently derived from the message header (fields),
// and in many cases it does not represent the actual type of VPP message.
// This means that some replies can be identified as requests, etc.
// TODO: use services to identify type of message
type MessageType int

const (
	// RequestMessage represents a VPP request message
	RequestMessage MessageType = iota
	// ReplyMessage represents a VPP reply message
	ReplyMessage
	// EventMessage represents a VPP event message
	EventMessage
	// OtherMessage represents other VPP message
	OtherMessage
)

// Message is an interface that is implemented by all VPP Binary API messages generated by the binapigenerator.
type Message interface {
	// GetMessageName returns the original VPP name of the message, as defined in the VPP API.
	GetMessageName() string

	// GetCrcString returns the string with CRC checksum of the message definition (the string represents a hexadecimal number).
	GetCrcString() string

	// GetMessageType returns the type of the VPP message.
	GetMessageType() MessageType
}

// DataType is an interface that is implemented by all VPP Binary API data types by the binapi_generator.
type DataType interface {
	// GetTypeName returns the original VPP name of the data type, as defined in the VPP API.
	GetTypeName() string
}

// ChannelProvider provides the communication channel with govpp core.
type ChannelProvider interface {
	// NewAPIChannel returns a new channel for communication with VPP via govpp core.
	// It uses default buffer sizes for the request and reply Go channels.
	NewAPIChannel() (Channel, error)

	// NewAPIChannelBuffered returns a new channel for communication with VPP via govpp core.
	// It allows to specify custom buffer sizes for the request and reply Go channels.
	NewAPIChannelBuffered(reqChanBufSize, replyChanBufSize int) (Channel, error)
}

// Channel provides methods for direct communication with VPP channel.
type Channel interface {
	// SendRequest asynchronously sends a request to VPP. Returns a request context, that can be used to call ReceiveReply.
	// In case of any errors by sending, the error will be delivered to ReplyChan (and returned by ReceiveReply).
	SendRequest(msg Message) RequestCtx

	// SendMultiRequest asynchronously sends a multipart request (request to which multiple responses are expected) to VPP.
	// Returns a multipart request context, that can be used to call ReceiveReply.
	// In case of any errors by sending, the error will be delivered to ReplyChan (and returned by ReceiveReply).
	SendMultiRequest(msg Message) MultiRequestCtx

	// SubscribeNotification subscribes for receiving of the specified notification messages via provided Go channel.
	// Note that the caller is responsible for creating the Go channel with preferred buffer size. If the channel's
	// buffer is full, the notifications will not be delivered into it.
	SubscribeNotification(notifChan chan Message, event Message) (SubscriptionCtx, error)

	// SetReplyTimeout sets the timeout for replies from VPP. It represents the maximum time the API waits for a reply
	// from VPP before returning an error.
	SetReplyTimeout(timeout time.Duration)

	// CheckCompatibility checks the compatiblity for the given messages.
	// It will return an error if any of the given messages are not compatible.
	CheckCompatiblity(msgs ...Message) error

	// Close closes the API channel and releases all API channel-related resources in the ChannelProvider.
	Close()
}

// RequestCtx is helper interface which allows to receive reply on request.
type RequestCtx interface {
	// ReceiveReply receives a reply from VPP (blocks until a reply is delivered from VPP, or until an error occurs).
	// The reply will be decoded into the msg argument. Error will be returned if the response cannot be received or decoded.
	ReceiveReply(msg Message) error
}

// MultiRequestCtx is helper interface which allows to receive reply on multi-request.
type MultiRequestCtx interface {
	// ReceiveReply receives a reply from VPP (blocks until a reply is delivered from VPP, or until an error occurs).
	// The reply will be decoded into the msg argument. If the last reply has been already consumed, lastReplyReceived is
	// set to true. Do not use the message itself if lastReplyReceived is true - it won't be filled with actual data.
	// Error will be returned if the response cannot be received or decoded.
	ReceiveReply(msg Message) (lastReplyReceived bool, err error)
}

// SubscriptionCtx is helper interface which allows to control subscription for notification events.
type SubscriptionCtx interface {
	// Unsubscribe unsubscribes from receiving the notifications tied to the subscription context.
	Unsubscribe() error
}

var registeredMessages = make(map[string]Message)

// RegisterMessage is called from generated code to register message.
func RegisterMessage(x Message, name string) {
	name = x.GetMessageName() + "_" + x.GetCrcString()
	/*if _, ok := registeredMessages[name]; ok {
		panic(fmt.Errorf("govpp: duplicate message registered: %s (%s)", name, x.GetCrcString()))
	}*/
	registeredMessages[name] = x
}

// GetRegisteredMessages returns list of all registered messages.
func GetRegisteredMessages() map[string]Message {
	return registeredMessages
}

// GoVppAPIPackageIsVersion1 is referenced from generated binapi files
// to assert that that code is compatible with this version of the GoVPP api package.
const GoVppAPIPackageIsVersion1 = true