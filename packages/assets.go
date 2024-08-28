/**
 * Copyright 2021 Appvia Ltd <info@appvia.io>
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

package packages

import (
	"embed"
	"fmt"
)

//go:embed *.yaml
var files embed.FS

// AssetNames lists the
func AssetNames() []string {
	var assets []string
	// read current dir
	d, err := files.ReadDir(".")
	if err != nil {
		panic(fmt.Errorf("something wrong - this should have been tested first"))
	}
	for _, f := range d {
		assets = append(assets, f.Name())
	}
	return assets
}

// Asset will get a file asset or an error if it doesn't exist
func Asset(name string) ([]byte, error) {
	return files.ReadFile(name)
}

// MustAsset will return a single file name asset or panic
func MustAsset(name string) []byte {
	// normally for cross-OS we would use filepath lib, but due to
	// embed.FS path handling: "The path separator is a forward slash,
	// even on Windows systems" (more in https://pkg.go.dev/embed)
	// We always must use path with embed! (even on Windows)
	content, err := Asset(name)
	if err != nil {
		panic(fmt.Errorf("embedded asset does not exist %s - %w", name, err))
	}
	return content
}
