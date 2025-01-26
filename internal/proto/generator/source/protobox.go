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

package source

import (
	"bytes"
	"embed"
	_ "embed"
	"strings"
	"text/template"

	"github.com/notjustmoney/protobox/internal/proto/generator/source/formatter"
)

var (
	//go:embed templates/*
	protoboxTemplatesFS embed.FS
)

type ProtoboxParams struct {
	Package string
	Imports []struct {
		Alias string
		Path  string
	}
	Events []struct {
		Name    string
		Topic   string
		Version string
	}
	Dispatchers []struct {
		Name       string
		Messages   []string
		Sequential bool
		Parallel   bool
	}
}

func Protobox(params ProtoboxParams) (string, error) {
	tmpl := template.Must(
		template.New("protobox").Funcs(template.FuncMap{
			"lower": strings.ToLower,
		}).ParseFS(protoboxTemplatesFS, "templates/*.tmpl"),
	)

	var buf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buf, "protobox", params); err != nil {
		return "", err
	}
	return formatter.Format(buf)
}
