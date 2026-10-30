package io

import (
	"lab2/internal/core"
)

// ParseModeC handles the input pipeline for static mode C
func ParseModeC(st *core.State, path string) error {
	return st.LoadFromFile_C(path)
}

// You can add ParseModeB, ParseModeA, etc., here as you implement them later!