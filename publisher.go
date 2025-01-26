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

package protobox

import (
	"context"

	"google.golang.org/protobuf/proto"
)

type Publisher interface {
	Publish(ctx context.Context, message proto.Message) error
}

type publisher struct {
	dispatchers []Dispatcher
}

func NewPublisher(dispatchers ...Dispatcher) Publisher {
	return &publisher{
		dispatchers: dispatchers,
	}
}

func (p *publisher) Publish(ctx context.Context, message proto.Message) error {
	for _, dispatcher := range p.dispatchers {
		if dispatcher == nil {
			continue
		}
		if err := dispatcher.Dispatch(ctx, message); err != nil {
			return err
		}
	}
	return nil
}
