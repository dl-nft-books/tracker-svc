package test

import (
	"github.com/stretchr/testify/assert"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/tokend/nft-books/contract-tracker/internal/config"
	"io/ioutil"
	"testing"
)

const (
	pathToFile = "./files/test_2.rtf"
)

func TestIpfsLoader(t *testing.T) {
	cfg := config.New(kv.NewViperFile("test_config.yaml"))
	implementation := cfg.IpfsLoader()
	file, err := ioutil.ReadFile(pathToFile)
	assert.Nil(t, err, "failed to open file")
	if len(file) == 0 {
		t.Error("File not found")
	}

	output, err := implementation.Load("test", file)
	assert.Nil(t, err, "failed to upload file")

	t.Log(*output)
}
