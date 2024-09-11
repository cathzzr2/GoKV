package fio

const DataFilePerm = 0644

type FileIOType = byte

const (
	// StandardFIO 
	StandardFIO FileIOType = iota

	// MemoryMap 
	MemoryMap
)

// IOManager - abstract IO interface campatible for multiple IO types，now support standard file IO
type IOManager interface {
	// Read 
	Read([]byte, int64) (int, error)

	// Write
	Write([]byte) (int, error)

	// Sync - make data long lasting
	Sync() error

	// Close 
	Close() error

	// Size 
	Size() (int64, error)
}

// NewIOManager initialize IOManager，now support standard FileIO
func NewIOManager(fileName string, ioType FileIOType) (IOManager, error) {
	switch ioType {
	case StandardFIO:
		return NewFileIOManager(fileName)
	case MemoryMap:
		return NewMMapIOManager(fileName)
	default:
		panic("unsupported io type")
	}
}
