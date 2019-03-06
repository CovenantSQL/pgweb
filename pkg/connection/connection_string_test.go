package connection

import (
	"testing"

	"github.com/CovenantSQL/pgweb/pkg/command"
	"github.com/stretchr/testify/assert"
)

func Test_Valid_Url(t *testing.T) {
	url := "covenantsql://0000000000000000000000000000000000000000000000000000000000000000"
	str, err := BuildStringFromOptions(command.Options{Url: url})

	assert.Equal(t, nil, err)
	assert.Equal(t, url, str)
}

func Test_Url(t *testing.T) {
	str, err := BuildStringFromOptions(command.Options{
		Url: "covenantsql://0000000000000000000000000000000000000000000000000000000000000000",
	})

	assert.Equal(t, nil, err)
	assert.Equal(t, "covenantsql://0000000000000000000000000000000000000000000000000000000000000000", str)
}

func Test_Flag_Args(t *testing.T) {
	str, err := BuildStringFromOptions(command.Options{
		DbName: "0000000000000000000000000000000000000000000000000000000000000000",
	})

	assert.Equal(t, nil, err)
	assert.Equal(t, "covenantsql://0000000000000000000000000000000000000000000000000000000000000000", str)
}

func Test_Blank(t *testing.T) {
	assert.Equal(t, true, IsBlank(command.Options{}))
	assert.Equal(t, false, IsBlank(command.Options{DbName: "0000000000000000000000000000000000000000000000000000000000000000"}))
	assert.Equal(t, false, IsBlank(command.Options{Url: "url"}))
}
