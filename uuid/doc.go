// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
The uuid command generates a UUID and prints it to the standard output.
It supports an option to generate UUIDs without hyphens.

Usage: uuid [flags]

Flags:

	-no-hyphen  Generate UUID without hyphens

Examples:

Generate a UUID with hyphens:

	$ uuid

Generate a UUID without hyphens:

	$ uuid -no-hyphen

The UUIDs are generated using the github.com/google/uuid package.
*/
package main
