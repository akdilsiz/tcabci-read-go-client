// Package tcabci_read_go_client
//
// Copyright 2013-2018 TransferChain A.G
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tcabci_read_go_client

// MessageType ..
type MessageType string

const (
	// Subscribe message
	Subscribe MessageType = "subscribe"
	// Unsubscribe message
	Unsubscribe MessageType = "unsubscribe"
)

// Message ..
type Message struct {
	IsWeb bool        `json:"is_web"`
	Type  MessageType `json:"type"`
	Addrs []string    `json:"addrs"`
}