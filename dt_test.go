package dt

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTree(t *testing.T) {

	str, err := Tree("./")

	assert.NoError(t, err)

	fmt.Println(str)

}
