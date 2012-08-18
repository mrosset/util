package human

import (
	"fmt"
)

type ByteSize float64

const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func (b ByteSize) String() string {
	switch {
	case b >= YB:
		return fmt.Sprintf("%6.1fYB", float64(b/YB))
	case b >= ZB:
		return fmt.Sprintf("%6.1fZB", float64(b/ZB))
	case b >= EB:
		return fmt.Sprintf("%6.1fEB", float64(b/EB))
	case b >= PB:
		return fmt.Sprintf("%6.1fPB", float64(b/PB))
	case b >= TB:
		return fmt.Sprintf("%6.1fTB", float64(b/TB))
	case b >= GB:
		return fmt.Sprintf("%6.1fGB", float64(b/GB))
	case b >= MB:
		return fmt.Sprintf("%6.1fMB", float64(b/MB))
	case b >= KB:
		return fmt.Sprintf("%6.1fKB", float64(b/KB))
	}
	return fmt.Sprintf("%6.0f B", float64(b))
}
