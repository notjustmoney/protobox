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

	"github.com/notjustmoney/protobox/outbox"

	sq "github.com/Masterminds/squirrel"
)

func Poller(dbCtx *dbCtx, cfg Config) outbox.PollerFunc {
	return func(ctx context.Context, pollCfg outbox.PollConfig) ([]outbox.Message, error) {
		rows, err := sq.
			Select("id", "topic", "payload").
			From(cfg.TableName).
			Where(sq.Eq{"processed_at": nil}).
			OrderBy("created_at").
			Limit(uint64(pollCfg.BatchSize)).
			Suffix("FOR UPDATE SKIP LOCKED").
			PlaceholderFormat(sq.Dollar).
			RunWith(dbCtx.executor(ctx)).
			QueryContext(ctx)
		if err != nil {
			return nil, err
		}

		var messages []outbox.Message
		for rows.Next() {
			var message outbox.Message
			if err := rows.Scan(&message.ID, &message.Topic, &message.Payload); err != nil {
				return nil, err
			}
			messages = append(messages, message)
		}

		return messages, nil
	}
}
