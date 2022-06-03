package index

import "github.com/awesome-cap/hashmap"

type Index interface {
    Set(k string, v Value) error
    Get(k string) (*Value, error)
    Del(k string) error
}

type Value struct {
    DataID int
    Offset uint64
}

func NewCacheIndex() *CacheIndex {
    return &CacheIndex{m: hashmap.New()}
}
