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
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Outbox interface {
	Insert(ctx context.Context, message []Message) error
}

func Insert[T interface {
	Topic() string
}](ctx context.Context, outbox Outbox, messages ...T) error {
	if len(messages) == 0 {
		return nil
	}
	var outboxMessages []Message
	for _, message := range messages {
		payload, err := json.Marshal(message)
		if err != nil {
			return err
		}
		outboxMessages = append(outboxMessages, Message{
			ID:          uuid.NewString(),
			Topic:       message.Topic(),
			Payload:     payload,
			Error:       nil,
			ProcessedAt: nil,
			CreatedAt:   time.Now(),
		})
	}

	return outbox.Insert(ctx, outboxMessages)
}
