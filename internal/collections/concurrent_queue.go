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

package collections

import (
	"sync"
	"sync/atomic"
)

type ConcurrentQueue[T any] struct {
	items []T
	mu    sync.Mutex
	size  atomic.Uint64
}

func NewConcurrentQueue[T any]() *ConcurrentQueue[T] {
	return &ConcurrentQueue[T]{
		items: make([]T, 0),
	}
}

func (q *ConcurrentQueue[T]) Enqueue(item T) {
	q.mu.Lock()
	q.items = append(q.items, item)
	q.mu.Unlock()
	q.size.Add(1)
}

func (q *ConcurrentQueue[T]) Dequeue() T {
	q.mu.Lock()
	defer q.mu.Unlock()

	item := q.items[0]
	q.items = q.items[1:]
	q.size.Add(^uint64(0))
	return item
}

func (q *ConcurrentQueue[T]) DequeueAll() []T {
	q.mu.Lock()
	defer q.mu.Unlock()

	items := make([]T, len(q.items))
	copy(items, q.items)
	q.items = make([]T, 0)
	q.size.Store(0)
	return items
}

func (q *ConcurrentQueue[T]) IsEmpty() bool {
	return q.size.Load() == 0
}
