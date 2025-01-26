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

package example

import (
	"context"
	"fmt"

	exampleeventv1 "github.com/notjustmoney/protobox/examples/event-dispatching/gen/proto/event/example/v1"
)

type PublishExampleCreatedOnExampleCreatedHandler struct{}

func NewPublishExampleCreatedOnExampleCreatedHandler() *PublishExampleCreatedOnExampleCreatedHandler {
	return &PublishExampleCreatedOnExampleCreatedHandler{}
}

func (h *PublishExampleCreatedOnExampleCreatedHandler) Handle(ctx context.Context, event *exampleeventv1.ExampleCreated) error {
	fmt.Println("PublishExampleCreatedOnExampleCreatedHandler: ", event)
	return nil
}
