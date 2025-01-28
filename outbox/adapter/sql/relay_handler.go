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

package outboxsql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/notjustmoney/protobox/outbox"
)

type RelayHandler struct {
	db           *sql.DB
	relayHandler *outbox.RelayHandler
}

func NewRelayHandler(
	db *sql.DB,
	publisher outbox.Publisher,
	cfg Config,
) *RelayHandler {
	dbCtx := &dbCtx{db: db}
	return &RelayHandler{
		db: db,
		relayHandler: outbox.NewRelayHandler(
			Poller(dbCtx, cfg),
			publisher,
			Marker(dbCtx, cfg),
		),
	}
}

func (h *RelayHandler) Handle(ctx context.Context) (int, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("outboxsql: %w", err)
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		}
		if err != nil {
			_ = tx.Rollback()
			return
		}
		_ = tx.Commit()
	}()

	processed, err := h.Handle(ctx)
	if err != nil {
		return 0, fmt.Errorf("outboxsql: %w", err)
	}
	return processed, nil
}
