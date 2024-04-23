package job

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SunSetPilot/cs5296-project/client/svc"
)

func TestFetchTasks(t *testing.T) {
	ctx := svc.NewMockServiceContext(true)
	tasks, err := fetchTasks(ctx)
	assert.NoError(t, err)
	if len(tasks) == 0 {
		t.Fail()
	}
	t.Logf("tasks: %v", tasks)
}
