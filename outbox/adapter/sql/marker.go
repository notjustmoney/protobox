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
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/samber/lo"

	"github.com/notjustmoney/protobox/outbox"
)

type MarkerFunc func(ctx context.Context, messages []outbox.PublishedMessage) (int, error)

func (f MarkerFunc) Mark(ctx context.Context, messages []outbox.PublishedMessage) (int, error) {
	return f(ctx, messages)
}

func Marker(
	dbCtx *dbCtx,
	cfg Config,
) MarkerFunc {
	return func(ctx context.Context, messages []outbox.PublishedMessage) (int, error) {
		args := make([]interface{}, 0, len(messages)*3)
		for i := 0; i < len(messages); i++ {
			args = append(args, messages[i].ID)
			args = append(args, messages[i].ProcessedAt.UTC())
			args = append(args, messages[i].Error)
		}

		if _, err := sq.
			Update(cfg.TableName).
			Set("processed_at", sq.Expr("v.processed_at")).
			Set("error", sq.Expr("v.error")).
			Suffix("FROM").
			SuffixExpr(
				sq.Alias(
					sq.Expr(fmt.Sprintf("VALUES %s", strings.Join(lo.Map(messages, func(_ outbox.PublishedMessage, index int) string {
						return fmt.Sprintf("($%d, $%d::timestamp with time zone, $%d)", index*3+1, index*3+2, index*3+3)
					}), ",")), args...),
					"v(id, processed_at, error)",
				),
			).
			Suffix("WHERE " + cfg.TableName + ".id = v.id::uuid").
			PlaceholderFormat(sq.Dollar).
			RunWith(dbCtx.executor(ctx)).
			ExecContext(ctx); err != nil {
			return 0, err
		}

		return len(messages), nil
	}
}
