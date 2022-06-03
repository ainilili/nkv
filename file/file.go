package file

import "os"

type File struct {
    path string
    file *os.File
}

func New(path string) (*File, error) {
    f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND, os.FileMode(0744))
    if err != nil {
        return nil, err
    }
    return &File{
        path: path,
        file: f,
    }, nil
}
