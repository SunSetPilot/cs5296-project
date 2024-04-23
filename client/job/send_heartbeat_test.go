package job

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SunSetPilot/cs5296-project/client/svc"
	"github.com/SunSetPilot/cs5296-project/model/request"
)

func TestSendHeartbeat(t *testing.T) {
	// Create a ServiceContext object
	ctx := svc.NewMockServiceContext(false)

	// Create a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify the request method and path
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "/api/v1/internal/heartbeat", r.URL.Path)

		// Parse the request body
		var req request.HeartbeatRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		assert.NoError(t, err)

		// Verify the fields in the request body
		assert.Equal(t, ctx.PodName, req.PodName)
		assert.Equal(t, ctx.PodUID, req.PodUID)
		assert.Equal(t, ctx.PodIP, req.PodIP)
		assert.Equal(t, ctx.NodeName, req.NodeName)
		assert.Equal(t, ctx.NodeIP, req.NodeIP)
		assert.Equal(t, ctx.ClientStatus.Load(), uint32(req.ClientStatus))

		// Return a successful response
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	ctx.ServerAddr = server.URL
	// Call the sendHeartbeat function
	err := sendHeartbeat(ctx)
	assert.NoError(t, err)
}
