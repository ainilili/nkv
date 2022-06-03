package index

import (
    "errors"
    "github.com/awesome-cap/hashmap"
)

type CacheIndex struct {
    m *hashmap.HashMap
}

func (c *CacheIndex) Set(k string, v Value) error {
    c.m.Set(k, v)
    return nil
}

func (c *CacheIndex) Get(k string) (*Value, error) {
    if v, ok := c.m.Get(k); ok {
        return v.(*Value), nil
    }
    return nil, errors.New("index not exists")
}

func (c *CacheIndex) Del(k string) error {
    c.m.Del(k)
    return nil
}
