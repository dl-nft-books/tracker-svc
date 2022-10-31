package ipfs_loader

type LoaderImplementation interface {
	Load(name string, file []byte) (response *string, err error)
}
