package integration_test

import (
	"net/http"
	"os"
	"testing"

	"github.com/gorilla/websocket"
)

func TestFrontHealth(t *testing.T) {
	url := os.Getenv("FRONT_URL") + "/api/health"
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil {
			t.Fatal(cerr)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("front not healthy, status: %d", resp.StatusCode)
	}
}

func TestAPI(t *testing.T) {
	url := os.Getenv("API_URL") + "/api/hello"
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil {
			t.Fatal(cerr)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("api failed, status: %d", resp.StatusCode)
	}
}

func TestWebSocket(t *testing.T) {
	wsURL := os.Getenv("WS_URL")
	c, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if cerr := c.Close(); cerr != nil {
			t.Fatal(cerr)
		}
	}()

	msg := "ping"
	if err := c.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
		t.Fatal(err)
	}

	// メッセージ受信
	_, reply, err := c.ReadMessage()
	if err != nil {
		t.Fatal(err)
	}

	if string(reply) != msg {
		t.Fatalf("expected %q, got %q", msg, reply)
	}
}
