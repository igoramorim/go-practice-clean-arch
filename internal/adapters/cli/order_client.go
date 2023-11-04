package cli

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/igoramorim/go-practice-clean-arch/internal/domain/dorder"
	"io"
	"net/http"
	"strconv"
)

func newOrderClient() *orderClient {
	return &orderClient{
		baseURL: "http://localhost:8080",
	}
}

type orderClient struct {
	baseURL string
}

func (oc *orderClient) createOrder(ctx context.Context, in dorder.CreateOrderUseCaseInput) (string, error) {
	url := fmt.Sprintf("%s/%s", oc.baseURL, "orders")
	b, err := json.Marshal(in)
	if err != nil {
		return "", err
	}
	body := bytes.NewBuffer(b)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return "", err
	}

	httpRes, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer httpRes.Body.Close()

	resBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return "", err
	}

	return string(resBody), nil
}

func (oc *orderClient) listOrders(ctx context.Context, in dorder.FindAllOrdersByPageUseCaseInput) (string, error) {
	url := fmt.Sprintf("%s/%s", oc.baseURL, "orders")
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	params := req.URL.Query()
	params.Add("page", strconv.Itoa(in.Page))
	params.Add("limit", strconv.Itoa(in.Limit))
	params.Add("sort", in.Sort)
	req.URL.RawQuery = params.Encode()

	httpRes, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer httpRes.Body.Close()

	resBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return "", err
	}

	return string(resBody), nil
}
