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

package idempotent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"

	"github.com/notjustmoney/protobox"
)

type EventHandler[T proto.Message] struct {
	id         string
	handler    protobox.EventHandler[T]
	store      Store
	identifier Identifier[T]
}

func (h *EventHandler[T]) Handle(ctx context.Context, message T) error {
	messageID := h.identifier.ID(message)
	if messageID == "" {
		return errors.New("inbox: empty message ID from identifier")
	}
	record, err := h.store.Get(ctx, h.id, messageID)
	if err != nil {
		return fmt.Errorf("inbox: %w", err)
	}
	if record != nil {
		return nil
	}

	if err := h.handler.Handle(ctx, message); err != nil {
		var alreadyHandledMessageErr *AlreadyHandledMessageError
		if errors.As(err, &alreadyHandledMessageErr) {
			// TODO: log and ignore the error
			return nil
		}
		return fmt.Errorf("inbox: %w", err)
	}

	record = &Record{
		ID:          uuid.NewString(),
		HandlerID:   h.id,
		MessageID:   messageID,
		ProcessedAt: time.Now(),
	}
	if err := h.store.Create(ctx, record); err != nil {
		return fmt.Errorf("inbox: %w", err)
	}

	return nil
}

type RequestHandler[T proto.Message, R proto.Message] struct {
	id         string
	handler    protobox.RequestHandler[T, R]
	store      Store
	identifier Identifier[T]
}

func (h *RequestHandler[T, R]) Handle(ctx context.Context, message T) (R, error) {
	messageID := h.identifier.ID(message)
	if messageID == "" {
		var empty R
		return empty, errors.New("inbox: empty message ID from identifier")
	}
	record, err := h.store.Get(ctx, h.id, messageID)
	if err != nil {
		var empty R
		return empty, fmt.Errorf("inbox: %w", err)
	}
	if record != nil {
		var empty R
		return empty, nil
	}

	resp, err := h.handler.Handle(ctx, message)
	if err != nil {
		var (
			empty                    R
			alreadyHandledMessageErr *AlreadyHandledMessageError
		)

		if errors.As(err, &alreadyHandledMessageErr) {
			// TODO: log and ignore the error
			return empty, nil
		}
		return empty, fmt.Errorf("inbox: %w", err)
	}

	record = &Record{
		ID:          uuid.NewString(),
		HandlerID:   h.id,
		MessageID:   messageID,
		ProcessedAt: time.Now(),
	}
	if err := h.store.Create(ctx, record); err != nil {
		var empty R
		return empty, fmt.Errorf("inbox: %w", err)
	}

	return resp, nil
}
