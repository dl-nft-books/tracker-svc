package tests

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/distributed_lab/kit/kv"
	"github.com/dl-nft-books/tracker-svc/internal/config"
)

const (
	pathToFile   = "./files/test.rtf"
	configKey    = "test_config.yaml"
	fileTestName = "test"
)

func TestIpfsLoader(t *testing.T) {
	cfg := config.New(kv.NewViperFile(configKey))
	implementation := cfg.IpfsLoader()

	file, err := ioutil.ReadFile(pathToFile)
	assert.Nil(t, err, "failed to open file")
	assert.NotEqual(t, len(file), 0, "file not found")

	output, err := implementation.Upload(fileTestName, file)
	assert.Nil(t, err, "failed to upload file")

	t.Log(*output)
}
