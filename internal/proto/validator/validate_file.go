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

package validator

import (
	"errors"

	"google.golang.org/protobuf/compiler/protogen"

	"github.com/notjustmoney/protobox/internal/proto/options"
)

func ValidateFile(file *protogen.File) error {
	var dispatcher *protogen.Message
	for _, message := range file.Messages {
		isDispatcher := options.IsDispatcher(message)
		isEvent := options.IsEvent(message)
		if isDispatcher && isEvent {
			return errors.New("dispatcher message cannot be an event")
		}
		if isDispatcher {
			if dispatcher != nil {
				return errors.New("multiple dispatcher messages found")
			}
			dispatcher = message
		}

	}
	return nil
}
