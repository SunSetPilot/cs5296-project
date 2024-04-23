package job

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SunSetPilot/cs5296-project/client/svc"
	"github.com/SunSetPilot/cs5296-project/model/request"
	"github.com/stretchr/testify/assert"
)

func TestSendHeartbeat(t *testing.T) {
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
		assert.Equal(t, "pod1", req.PodName)
		assert.Equal(t, "uid1", req.PodUID)
		assert.Equal(t, "10.0.0.1", req.PodIP)
		assert.Equal(t, "node1", req.NodeName)
		assert.Equal(t, "192.168.0.1", req.NodeIP)
		assert.Equal(t, uint8(1), req.ClientStatus)

		// Return a successful response
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// Create a ServiceContext object
	ctx := &svc.ServiceContext{
		PodName:    "pod1",
		PodUID:     "uid1",
		PodIP:      "10.0.0.1",
		NodeName:   "node1",
		NodeIP:     "192.168.0.1",
		ServerAddr: server.URL,
	}
	ctx.ClientStatus.Store(1)

	// Call the sendHeartbeat function
	sendHeartbeat(ctx)
}
