package job

import (
	"testing"

	"github.com/SunSetPilot/cs5296-project/client/svc"
)

func TestFetchTasks(t *testing.T) {
	ctx := svc.NewMockServiceContext()
	tasks, err := fetchTasks(ctx)
	if err != nil {
		t.Fail()
	}
	if len(tasks) == 0 {
		t.Fail()
	}
	t.Logf("tasks: %v", tasks)
}
