// Copyright 2021 CloudWeGo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package backend

import "github.com/cloudwego/thriftgo/plugin"

// LogFunc .
type LogFunc struct {
	Info      func(v ...interface{})
	Warn      func(v ...interface{})
	MultiWarn func(warns []string)
}

// Backend handles the code generation for a language.
type Backend interface {
	// The name of this backend.
	Name() string

	// Lang returns the target langauge that this backend supports.
	Lang() string

	// Generate generates codes.
	Generate(req *plugin.Request, log LogFunc) (res *plugin.Response)

	// Options returns the available options with their descriptions.
	Options() []plugin.Option

	// BuiltinPlugins returns a list of built-in plugins of the backend.
	BuiltinPlugins() []*plugin.Desc

	// GetPlugin returns a plugin that match the description or nil.
	GetPlugin(desc *plugin.Desc) plugin.Plugin
}
