/*
 * Copyright (c) 2025. protobox
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package types

import (
	"google.golang.org/protobuf/compiler/protogen"

	"github.com/notjustmoney/protobox/internal/proto/options"
)

type Dispatcher struct {
	File       *protogen.File
	Message    *protogen.Message
	Sequential bool
	Parallel   bool
}

func (d *Dispatcher) Events() []*protogen.Message {
	var events []*protogen.Message
	for _, message := range d.File.Messages {
		if options.IsEvent(message) {
			events = append(events, message)
		}
	}
	return events
}
