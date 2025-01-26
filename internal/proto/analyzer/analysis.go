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
	"google.golang.org/protobuf/compiler/protogen"

	"github.com/notjustmoney/protobox/internal/proto/analyzer/internal/registry"
	"github.com/notjustmoney/protobox/internal/proto/analyzer/types"
)

type Analysis struct {
	dispatcherRegistry *registry.DispatcherRegistry
	eventRegistry      *registry.EventRegistry
}

func NewAnalysis(
	dispatcherRegistry *registry.DispatcherRegistry,
	eventRegistry *registry.EventRegistry,
) *Analysis {
	return &Analysis{
		dispatcherRegistry: dispatcherRegistry,
		eventRegistry:      eventRegistry,
	}
}

func (a *Analysis) EventsByFile(file *protogen.File) []*types.Event {
	return a.eventRegistry.EventsByFile(file)
}

func (a *Analysis) DispatchersByFile(file *protogen.File) []*types.Dispatcher {
	return a.dispatcherRegistry.DispatchersByFile(file)
}
