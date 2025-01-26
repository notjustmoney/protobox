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
	"fmt"
	"os"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"

	"github.com/notjustmoney/protobox/internal/proto/generator"
)

func main() {
	protogen.Options{
		ParamFunc:                    nil,
		ImportRewriteFunc:            nil,
		InternalStripForEditionsDiff: nil,
		DefaultAPILevel:              0,
	}.Run(func(plugin *protogen.Plugin) error {
		plugin.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		if err := generator.Generate(plugin); err != nil {
			fmt.Fprintf(os.Stderr, "protobox: %v\n", err)
			plugin.Error(err)
		}
		return nil
	})
}
