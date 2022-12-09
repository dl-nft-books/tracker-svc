package ipfs

// Uploader is the interface that every other service integration
// struct must implement. Format of response is arbitrary since
// it will be simply just debugged later
type Uploader interface {
	Upload(name string, file []byte) (response *string, err error)
}
