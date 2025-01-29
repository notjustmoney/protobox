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
	"fmt"
	"time"

	"github.com/samber/lo"
	"golang.org/x/sync/errgroup"

	"github.com/notjustmoney/protobox/internal/collections"
)

type RelayConfig struct {
	BatchSize int
}

type RelayHandler struct {
	poller    Poller
	publisher Publisher
	marker    Marker
}

func NewRelayHandler(
	poller Poller,
	publisher Publisher,
	marker Marker,
) *RelayHandler {
	return &RelayHandler{
		publisher: publisher,
		poller:    poller,
		marker:    marker,
	}
}

func (h *RelayHandler) Handle(ctx context.Context, cfg RelayConfig) (int, error) {
	batchSize := cfg.BatchSize
	if batchSize <= 0 {
		batchSize = 100
	}
	records, err := h.poller.Poll(ctx, PollConfig{
		BatchSize: cfg.BatchSize,
	})
	if err != nil {
		fmt.Errorf("outbox: %w", err)
	}
	queue := collections.NewConcurrentQueue[PublishedRecord]()
	g, _ := errgroup.WithContext(ctx)
	for _, record := range records {
		g.Go(func() error {
			h.publish(ctx, &record, queue)
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return 0, fmt.Errorf("outbox: %w", err)
	}

	if queue.IsEmpty() {
		return 0, nil
	}

	processed, err := h.marker.Mark(ctx, queue.DequeueAll())
	if err != nil {
		return 0, fmt.Errorf("outbox: %w", err)
	}

	return processed, nil
}

func (h *RelayHandler) publish(ctx context.Context, record *Record, queue *collections.ConcurrentQueue[PublishedRecord]) {
	err := h.publisher.Publish(ctx, record)
	if err != nil {
		queue.Enqueue(PublishedRecord{
			ID:          record.ID,
			ProcessedAt: time.Now(),
			Error: lo.TernaryF(err != nil, func() *string {
				return lo.ToPtr(err.Error())
			}, nil),
		})
	}
	queue.Enqueue(PublishedRecord{
		ID:          record.ID,
		ProcessedAt: time.Now(),
		Error:       nil,
	})
}
