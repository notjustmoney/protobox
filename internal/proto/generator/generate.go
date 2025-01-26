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

package generator

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"

	"github.com/notjustmoney/protobox/internal/proto/analyzer"
	"github.com/notjustmoney/protobox/internal/proto/generator/source"
)

func Generate(plugin *protogen.Plugin) error {
	analysis, err := analyzer.Analyze(plugin)
	if err != nil {
		return fmt.Errorf("generator: %w", err)
	}
	if err := generateProtobox(plugin, analysis); err != nil {
		return fmt.Errorf("generator: %w", err)
	}
	return nil
}

func generateProtobox(plugin *protogen.Plugin, analysis *analyzer.Analysis) error {
	for _, file := range plugin.Files {
		events := analysis.EventsByFile(file)
		if len(events) == 0 {
			continue
		}
		params := source.ProtoboxParams{
			Package: string(file.GoPackageName),
		}
		for _, event := range events {
			params.Events = append(params.Events, struct {
				Name    string
				Topic   string
				Version string
			}{
				Name:    event.Message.GoIdent.GoName,
				Topic:   event.Topic,
				Version: event.Version,
			})
		}
		for _, dispatcher := range analysis.DispatchersByFile(file) {
			dispatcherParams := struct {
				Name       string
				Messages   []string
				Sequential bool
				Parallel   bool
			}{
				Name:       dispatcher.Message.GoIdent.GoName,
				Sequential: dispatcher.Sequential,
				Parallel:   dispatcher.Parallel,
			}
			for _, dispatcherEvents := range dispatcher.Events() {
				dispatcherParams.Messages = append(dispatcherParams.Messages, dispatcherEvents.GoIdent.GoName)
			}
			params.Dispatchers = append(params.Dispatchers, dispatcherParams)
		}
		sourceCode, err := source.Protobox(params)
		if err != nil {
			return fmt.Errorf("generate protobox: %w", err)
		}
		filename := file.GeneratedFilenamePrefix + "_protobox.pb.go"
		gen := plugin.NewGeneratedFile(filename, file.GoImportPath)
		gen.P(sourceCode)
	}
	return nil
}
