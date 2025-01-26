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

package main

import (
	"context"
	"fmt"

	"github.com/notjustmoney/protobox"
	exampleeventv1 "github.com/notjustmoney/protobox/examples/event-dispatching/gen/proto/event/example/v1"
	examplerpcv1 "github.com/notjustmoney/protobox/examples/event-dispatching/gen/proto/rpc/example/v1"
	"github.com/notjustmoney/protobox/examples/event-dispatching/internal/example"
)

func main() {
	sendNotificationOnExampleCreatedHandler := example.NewSendNotificationOnExampleCreatedHandler()
	publishExampleCreatedOnExampleCreatedHandler := example.NewPublishExampleCreatedOnExampleCreatedHandler()

	exampleDispatcher := exampleeventv1.NewSequentialEventDispatcher(
		exampleeventv1.WithExampleCreatedEventHandler(
			sendNotificationOnExampleCreatedHandler,
			publishExampleCreatedOnExampleCreatedHandler,
		),
	)
	examplePublisher := protobox.NewPublisher(exampleDispatcher)
	createExampleHandler := example.NewCreateExampleHandler(examplePublisher)
	resp, err := createExampleHandler.Handle(context.Background(), &examplerpcv1.CreateExampleRequest{
		Name:        "example",
		Description: "example description",
	})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Response:", resp)
}
