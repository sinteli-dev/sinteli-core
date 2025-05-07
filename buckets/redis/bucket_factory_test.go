// Licensed under the Apache License, Version 2.0
// Details: https://raw.githubusercontent.com/square/quotaservice/master/LICENSE
//
// SPDX-FileCopyrightText: Square, Inc.
// SPDX-FileCopyrightText: HGLOW MEDIA Inc.
// SPDX-FileModified: 2025-05-07
// SPDX-License-Identifier: Apache-2.0
// Modification: Replaced github.com/pkg/errors with standard library error handling.

package redis

import (
	"errors"
	"fmt"
	"testing"
)

func TestIsRedisClientClosedError(t *testing.T) {
	tests := []struct {
		input        error
		isCloseError bool
	}{
		{
			// Test exactly the error
			input:        fmt.Errorf(redisClientClosedError),
			isCloseError: true,
		},
		{
			// Test the error wrapped
			input:        fmt.Errorf("obfuscate: %w", fmt.Errorf(redisClientClosedError)),
			isCloseError: true,
		},
		{
			// test not the error
			input:        errors.New("just another error"),
			isCloseError: false,
		},
		{
			// test not the error wrapped with the text of the error (this should never happen)
			input:        fmt.Errorf("%s: %w", redisClientClosedError, errors.New("just another error")),
			isCloseError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.input.Error(), func(t *testing.T) {
			result := isRedisClientClosedError(test.input)
			if result != test.isCloseError {
				t.Fatal("failed to detect error")
			}
		})
	}
}
