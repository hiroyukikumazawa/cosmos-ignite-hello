package types

const (
	// ModuleName defines the module name
	ModuleName = "hello"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_hello"
)

var (
	ParamsKey = []byte("p_hello")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
