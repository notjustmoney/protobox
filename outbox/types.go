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

package outbox

import (
	"context"
	"time"
)

type Message struct {
	ID          string
	Topic       string
	Payload     []byte
	Error       *string
	CreatedAt   time.Time
	ProcessedAt *time.Time
}

type Publisher interface {
	Publish(ctx context.Context, message *Message) error
}

type PublisherFunc func(ctx context.Context, message Message) error

func (f PublisherFunc) Publish(ctx context.Context, message Message) error {
	return f(ctx, message)
}

type PollConfig struct {
	BatchSize int
}

type Poller interface {
	Poll(ctx context.Context, cfg PollConfig) ([]Message, error)
}

type PollerFunc func(ctx context.Context, cfg PollConfig) ([]Message, error)

func (f PollerFunc) Poll(ctx context.Context, cfg PollConfig) ([]Message, error) {
	return f(ctx, cfg)
}

type PublishedMessage struct {
	ID          string
	Error       *string
	ProcessedAt time.Time
}

type Marker interface {
	Mark(ctx context.Context, messages []PublishedMessage) (int, error)
}

type MarkerFunc func(ctx context.Context, messages []PublishedMessage) (int, error)

func (f MarkerFunc) Mark(ctx context.Context, messages []PublishedMessage) (int, error) {
	return f(ctx, messages)
}
