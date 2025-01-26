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

package analyzer

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"

	"github.com/notjustmoney/protobox/internal/proto/analyzer/internal/registry"
	"github.com/notjustmoney/protobox/internal/proto/analyzer/types"
	"github.com/notjustmoney/protobox/internal/proto/options"
	"github.com/notjustmoney/protobox/internal/proto/validator"
)

type analyzer struct {
	dispatcherRegistry *registry.DispatcherRegistry
	eventRegistry      *registry.EventRegistry
}

func newAnalyzer() *analyzer {
	return &analyzer{
		dispatcherRegistry: registry.NewDispatcherRegistry(),
		eventRegistry:      registry.NewEventRegistry(),
	}
}

func (a *analyzer) Analyze(plugin *protogen.Plugin) (*Analysis, error) {
	for _, file := range plugin.Files {
		if !file.Generate {
			continue
		}
		if err := a.analyzeFile(file); err != nil {
			return nil, fmt.Errorf("analyzer: %w", err)
		}
	}

	return &Analysis{
		eventRegistry:      a.eventRegistry,
		dispatcherRegistry: a.dispatcherRegistry,
	}, nil
}

func (a *analyzer) analyzeFile(file *protogen.File) error {
	for _, message := range file.Messages {
		if err := validator.ValidateFile(file); err != nil {
			return fmt.Errorf("analyzeFile: %w", err)
		}
		if err := a.analyzeMessage(file, message); err != nil {
			return fmt.Errorf("analyzeFile: %w", err)
		}
	}
	return nil
}

func (a *analyzer) analyzeMessage(file *protogen.File, message *protogen.Message) error {
	if options.IsEvent(message) {
		if err := a.analyzeEvent(file, message); err != nil {
			return fmt.Errorf("analyzeMessage: %w", err)
		}
	}
	if options.IsDispatcher(message) {
		if err := a.analyzeDispatcher(file, message); err != nil {
			return fmt.Errorf("analyzeMessage: %w", err)
		}
	}
	return nil
}

func (a *analyzer) analyzeEvent(file *protogen.File, message *protogen.Message) error {
	event := options.Event(message)
	if event == nil {
		return nil
	}
	if err := a.eventRegistry.AddEvent(&types.Event{
		File:    file,
		Message: message,
		Topic:   options.Topic(message),
		Version: options.Version(message),
	}); err != nil {
		return fmt.Errorf("analyzeEvent: %w", err)
	}
	return nil
}

func (a *analyzer) analyzeDispatcher(file *protogen.File, message *protogen.Message) error {
	dispatcher := options.Dispatcher(message)
	if dispatcher == nil {
		return nil
	}

	a.dispatcherRegistry.AddDispatcher(&types.Dispatcher{
		File:       file,
		Message:    message,
		Sequential: dispatcher.Sequential,
		Parallel:   dispatcher.Parallel,
	})
	return nil
}
