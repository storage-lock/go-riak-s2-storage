package riak_s2_storage

import (
	"context"

	s3_storage "github.com/storage-lock/go-s3-storage"
	"github.com/storage-lock/go-storage"
)

// RiakS2Storage Riak S2对象存储实现 (S3兼容)
type RiakS2Storage struct {
	*s3_storage.S3Storage
}

func NewRiakS2Storage(ctx context.Context, options *s3_storage.S3StorageOptions) (*RiakS2Storage, error) {
	s3Storage, err := s3_storage.NewS3Storage(ctx, options)
	if err != nil {
		return nil, err
	}
	return &RiakS2Storage{S3Storage: s3Storage}, nil
}

const StorageName = "riak-s2-storage"

func (x *RiakS2Storage) GetName() string {
	return StorageName
}

func (x *RiakS2Storage) Capabilities() []storage.StorageCapability {
	return x.S3Storage.Capabilities()
}
