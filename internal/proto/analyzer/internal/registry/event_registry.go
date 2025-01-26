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

package registry

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"

	"github.com/notjustmoney/protobox/internal/proto/analyzer/types"
)

type EventRegistry struct {
	events []*types.Event
	topics map[string]struct{}
}

func NewEventRegistry() *EventRegistry {
	return &EventRegistry{
		topics: make(map[string]struct{}),
	}
}

func (r *EventRegistry) Events() []*types.Event {
	return r.events
}

func (r *EventRegistry) EventsByFile(file *protogen.File) []*types.Event {
	var events []*types.Event
	for _, event := range r.events {
		if event.File == file {
			events = append(events, event)
		}
	}
	return events
}

func (r *EventRegistry) AddEvent(event *types.Event) error {
	if _, ok := r.topics[event.Topic]; ok {
		return fmt.Errorf("event with topic %s already exists", event.Topic)
	}
	r.events = append(r.events, event)
	r.topics[event.Topic] = struct{}{}
	return nil
}

func (r *EventRegistry) HasTopic(topic string) bool {
	_, ok := r.topics[topic]
	return ok
}
