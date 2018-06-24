//
// Copyright (c) 2016-2018 The Aiicy Team
// Licensed under The MIT License (MIT)
// See: LICENSE
//

package common

import "os"

// IsFile returns true if given path is a file,
// or returns false when it's a directory or does not exist.
func IsFile(filePath string) bool {
	f, e := os.Stat(filePath)
	if e != nil {
		return false
	}
	return !f.IsDir()
}
