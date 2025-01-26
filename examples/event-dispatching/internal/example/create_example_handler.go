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
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/notjustmoney/protobox"
	exampleeventv1 "github.com/notjustmoney/protobox/examples/event-dispatching/gen/proto/event/example/v1"
	examplerpcv1 "github.com/notjustmoney/protobox/examples/event-dispatching/gen/proto/rpc/example/v1"
)

type CreateExampleHandler struct {
	publisher protobox.Publisher
}

func NewCreateExampleHandler(
	publisher protobox.Publisher,
) *CreateExampleHandler {
	return &CreateExampleHandler{
		publisher: publisher,
	}
}

func (h *CreateExampleHandler) Handle(ctx context.Context, req *examplerpcv1.CreateExampleRequest) (*examplerpcv1.CreateExampleResponse, error) {
	id := uuid.NewString()
	now := time.Now()
	if err := h.publisher.Publish(ctx, &exampleeventv1.ExampleCreated{
		Id:          id,
		Name:        req.Name,
		Description: req.Description,
		CreatedAt:   timestamppb.New(now),
	}); err != nil {
		return nil, err
	}
	return &examplerpcv1.CreateExampleResponse{
		Id:          id,
		Name:        req.Name,
		Description: req.Description,
		CreatedAt:   timestamppb.New(now),
	}, nil
}
