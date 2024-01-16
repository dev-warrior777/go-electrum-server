package server

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync/atomic"
	"time"

	"github.com/dev-warrior777/go-electrum-server.git/electrumx/lib"
	"go.uber.org/zap"
)

const defaultDaemonTimeout = 10 * time.Second

type positional []any

// DaemonClient is an Bitcoin & clones HTTP JSON-RPC client.
// Some methods are used by some coins but the majority are
// common.
type DaemonClient struct {
	reqID      uint64
	url        string
	auth       string
	HTTPClient *http.Client
	Timeout    time.Duration
}

// NewDaemonClient constructs a new Daemon RPC client with the given
// authorization information and endpoint. The endpoint should include the
// protocol, e.g. http://127.0.0.1:4567. To specify a custom http.Client or
// request timeout, the fields may be set after construction.
func NewDaemonClient(endpoint, user, pass string) *DaemonClient {
	// Prepare the HTTP Basic Authorization request header.
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
	// fmt.Println(string(reqMsg))

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
		zap.S().Errorf("HttpClient.do %v", err)
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

func (dc *DaemonClient) GetBlockCount(ctx context.Context) (int, error) {
	var res int
	err := dc.Call(ctx, "getblockcount", nil, &res)
	if err != nil {
		return -1, err
	}
	return res, nil
}

func (dc *DaemonClient) GetBlockHash(ctx context.Context, height int64) (string, error) {
	var res string
	err := dc.Call(ctx, "getblockhash", positional{height}, &res)
	if err != nil {
		return "", err
	}
	return res, nil
}

// GetBlockHashes returns 'count' block hashes starting at height 'height'. If
// If there is less blocks than 'count' whatever blocks left will be returned.
// The caller should check the returned blockcountcount 'i'
func (dc *DaemonClient) GetBlockHashes(ctx context.Context, height int64, count int) ([]string, int, error) {
	var hashes = make([]string, 0, 8)
	var i int
	for i = 0; i < count; i++ {
		var res string
		err := dc.Call(ctx, "getblockhash", positional{height + int64(i)}, &res)
		if err != nil {
			rpcErr, parseErr := parseDaemonError(err)
			if parseErr != nil {
				return nil, 0, parseErr
			}
			if rpcErr.Code != RPC_BLK_OUT_OF_RANGE {
				newErr := fmt.Sprintf("invalid error: code %d - %s",
					rpcErr.Code, rpcErr.Message)
				return nil, 0, errors.New(newErr)
			}
			// return what we can
			return hashes, i, nil
		}
		hashes = append(hashes, res)
	}
	return hashes, i, nil
}

type BlockHeader struct {
}
type Block struct {
	Header *BlockHeader
	//
}

// GetDeserializedBlock returns a basic block struct.
func (dc *DaemonClient) GetDeserializedBlock(ctx context.Context, blockHash string) (*Block, error) {

	return nil, errors.New("method not implemented")
}

func (dc *DaemonClient) GetRawBlocks(ctx context.Context, blockHashes []string) ([]string, error) {

	return nil, errors.New("method not implemented")
}

func (dc *DaemonClient) GetMempoolHashes(ctx context.Context) ([]string, error) {
	var hashes = make([]string, 0, 8)
	err := dc.Call(ctx, "getrawmempool", nil, &hashes)
	if err != nil {
		return nil, err
	}
	return hashes, nil
}

const (
	RPC_BLK_OUT_OF_RANGE = -8
	RPC_IN_WARMUP        = -28
	RPC_PARSE_ERROR      = -32700
)

func parseDaemonError(e error) (*lib.RPCError, error) {
	type ErrRes struct {
		Result json.RawMessage `json:"result"`
		Error  *lib.RPCError   `json:"error"`
		ID     uint64          `json:"id"`
	}
	s := e.Error()
	startIdx := strings.Index(s, "{")
	js := s[startIdx:]
	var res ErrRes
	err := json.Unmarshal([]byte(js), &res)
	if err != nil {
		return nil, err
	}
	return res.Error, nil
}
