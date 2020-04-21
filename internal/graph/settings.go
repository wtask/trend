package graph

import "sync"
// ResolverSettings aggregates settings of all resolvers.
type ResolverSettings struct {
	*ProgressionSettings
}

// ProgressionSettings contains options and operational values for progression resolvers.
type ProgressionSettings struct {
	LengthRange [2]int
	operational struct {
		// progression current length
		length struct {
			sync.RWMutex //protects concurrent changes
			value        int
		}
	}
}

// SetLength changes the length of computed progressions, is safe for concurrent calls.
func (s *ProgressionSettings) SetLength(l int) *ProgressionSettings {
	s.operational.length.Lock()
	defer s.operational.length.Unlock()
	if l >= s.LengthRange[0] && l <= s.LengthRange[1] {
		s.operational.length.value = l
	}
	return s
}

// Length returns current length value, is save for concurrent calls.
func (s *ProgressionSettings) Length() int {
	s.operational.length.RLock()
	defer s.operational.length.RUnlock()
	return s.operational.length.value
}
