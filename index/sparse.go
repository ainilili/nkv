package index

import (
    "errors"
    "github.com/ainilili/nkv/file"
)

type SparseIndex struct {
    file *file.File
}

func (s *SparseIndex) Set(k, v string) error {
    return errors.New("not implements")
}

func (s *SparseIndex) Get(k string) error {
    return errors.New("not implements")
}

func (s *SparseIndex) Del(k string) error {
    return errors.New("not implements")
}
