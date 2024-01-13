package server

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/dev-warrior777/go-electrum-server.git/electrumx/lib"
)

const defaultDaemonTimeout = 10 * time.Second

type positional []any

// DaemonClient is an Electrum wallet HTTP JSON-RPC client.
type DaemonClient struct {
	reqID      uint64
	url        string
	auth       string
	walletFile string

	// HTTPClient may be set by the user to a custom http.Client. The
	// constructor sets a vanilla client.
	HTTPClient *http.Client
	// Timeout is the timeout on http requests. A 10 second default is set by
	// the constructor.
	Timeout time.Duration
}

// NewDaemonClient constructs a new Daemon RPC client with the given
// authorization information and endpoint. The endpoint should include the
// protocol, e.g. http://127.0.0.1:4567. To specify a custom http.Client or
// request timeout, the fields may be set after construction.
func NewDaemonClient(user, pass, endpoint string) *DaemonClient {
	// Prepare the HTTP Basic Authorization request header. This avoids
	// re-encoding it for every request with (*http.Request).SetBasicAuth.
	auth := "Basic " + base64.StdEncoding.EncodeToString([]byte(user+":"+pass))
	return &DaemonClient{
		url:        endpoint,
		auth:       auth,
		HTTPClient: &http.Client{},
		Timeout:    defaultDaemonTimeout,
	}
}

func (dc *DaemonClient) nextID() uint64 {
	return atomic.AddUint64(&dc.reqID, 1)
}

// Call makes a JSON-RPC request for the given method with the provided
// arguments. args may be a struct or slice that marshalls to JSON. If it is a
// slice, it represents positional arguments. If it is a struct or pointer to a
// struct, it represents "named" parameters in key-value format. Any arguments
// should have their fields appropriately tagged for JSON marshalling. The
// result is marshaled into result if it is non-nil, otherwise the result is
// discarded.
func (dc *DaemonClient) Call(ctx context.Context, method string, args any, result any) error {
	reqMsg, err := lib.PrepareRequest(dc.nextID(), method, args)
	if err != nil {
		return err
	}

	bodyReader := bytes.NewReader(reqMsg)
	ctx, cancel := context.WithTimeout(ctx, dc.Timeout)
	defer cancel()
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, dc.url, bodyReader)
	if err != nil {
		return err
	}
	httpReq.Close = true
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", dc.auth) // httpReq.SetBasicAuth(ec.user, ec.pass)

	resp, err := dc.HTTPClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("%d: %s", resp.StatusCode, string(b))
	}

	jsonResp := &lib.JsonResponse{}
	err = json.NewDecoder(resp.Body).Decode(jsonResp)
	if err != nil {
		return err
	}
	if jsonResp.Error != nil {
		return jsonResp.Error
	}

	if result != nil {
		return json.Unmarshal(jsonResp.Result, result)
	}
	return nil
}

func (dc *DaemonClient) GetBlockHash(ctx context.Context, height int64) (string, error) {
	var res string
	err := dc.Call(ctx, "getblockhash", positional{height}, &res)
	if err != nil {
		return "", err
	}
	return res, nil
}
