package config

type FileStore struct {
	path string
}

func NewFileStore(path string) (*FileStore, error) {
	fs := &FileStore{
		path: path,
	}
	return fs, nil
}

func (fs *FileStore) Close() error {
	return nil
}
