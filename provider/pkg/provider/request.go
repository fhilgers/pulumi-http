package provider

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	rehttp "github.com/hashicorp/go-retryablehttp"

	"github.com/pulumi/pulumi/pkg/v3/resource/provider"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

type request struct {
	Method             string               `pulumi:"method"`
	URL                string               `pulumi:"url"`
	Header             *map[string][]string `pulumi:"header,optional"`
	Body               *string              `pulumi:"body,optional"`
	ExpectedStatusCode *int                 `pulumi:"expectedStatusCode,optional"`
	MaxRetries         int                  `pulumi:"maxRetries,optional"`
	RetryWaitMin       *int                 `pulumi:"retryWaitMin,optional"`
	RetryWaitMax       *int                 `pulumi:"retryWaitMax,optional"`
	Certificates       *[]string            `pulumi:"certificates,optional"`
	RootCAs            *[]string            `pulumi:"rootCAs,optional"`
	ServerName         *string              `pulumi:"serverName,optional"`
	InsecureSkipVerify *bool                `pulumi:"insecureSkipVerify,optional"`
}

type response struct {
	Status     string              `pulumi:"status"`
	StatusCode int                 `pulumi:"statusCode"`
	Header     map[string][]string `pulumi:"header"`
	Body       string              `pulumi:"body"`
}

type requestResource struct {
	Create *request `pulumi:"create,optional"`
	Delete *request `pulumi:"delete,optional"`

	// Output
	Response response `pulumi:"response"`
}

func (r *requestResource) RunCreate(ctx context.Context, host *provider.HostClient, urn resource.URN) (string, error) {
	if r.Create == nil {
		id, _ := resource.NewUniqueHex("", 0, 0)
		return id, nil
	}

	response, id, err := r.run(ctx, r.Create, host, urn)
	if err != nil {
		return "", err
	}

	r.Response = *response

	return id, nil
}

func (r *requestResource) RunDelete(ctx context.Context, host *provider.HostClient, urn resource.URN) error {
	if r.Delete == nil {
		return nil
	}

	_, _, err := r.run(ctx, r.Delete, host, urn)
	return err
}

func (rr *requestResource) run(ctx context.Context, r *request, host *provider.HostClient, urn resource.URN) (*response, string, error) {

	url, err := url.Parse(r.URL)

	if err != nil {
		return nil, "", err
	}

	header := http.Header{}
	if r.Header != nil {
		header = *r.Header
	}

	var body io.ReadCloser
	if r.Body != nil {
		body = io.NopCloser(strings.NewReader(*r.Body))
	}

	request := &http.Request{
		Method: r.Method,
		URL:    url,
		Header: header,
		Body:   body,
	}

	var certificates []tls.Certificate

	if r.Certificates != nil {
		for _, cert := range *r.Certificates {
			pair, err := tls.X509KeyPair([]byte(cert), []byte(cert))
			if err != nil {
				return nil, "", err
			}
			certificates = append(certificates, pair)
		}
	}

	caCertPool := x509.NewCertPool()
	if r.RootCAs != nil {
		for _, ca := range *r.RootCAs {
			caCertPool.AppendCertsFromPEM([]byte(ca))
		}
	}

	var serverName string
	if r.ServerName != nil {
		serverName = *r.ServerName
	}

	var skipVerify bool
	if r.InsecureSkipVerify != nil {
		skipVerify = *r.InsecureSkipVerify
	}

	tlsConfig := &tls.Config{
		Certificates:       certificates,
		RootCAs:            caCertPool,
		ServerName:         serverName,
		InsecureSkipVerify: skipVerify,
	}

	tlsConfig.BuildNameToCertificate()

	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	client := &http.Client{
		Transport: transport,
	}

	minDelay := time.Second
	maxDelay := time.Second

	if r.RetryWaitMin != nil {
		minDelay = time.Duration(*r.RetryWaitMin) * time.Second
	}
	if r.RetryWaitMax != nil {
		maxDelay = time.Duration(*r.RetryWaitMax) * time.Second
	}

	expectedStatus := 200
	if r.ExpectedStatusCode != nil {
		expectedStatus = *r.ExpectedStatusCode
	}

	checkRetry := func(ctx context.Context, resp *http.Response, err error) (bool, error) {
		if ctx.Err() != nil {
			return false, ctx.Err()
		}

		if err != nil {
			return true, nil
		}

		if resp.StatusCode != expectedStatus {
			return true, nil
		}

		return false, nil
	}

	retryClient := rehttp.NewClient()

	retryClient.RetryMax = r.MaxRetries
	retryClient.RetryWaitMin = minDelay
	retryClient.RetryWaitMax = maxDelay
	retryClient.HTTPClient = client
	retryClient.CheckRetry = checkRetry

	retryRequest, err := rehttp.FromRequest(request)
	if err != nil {
		return nil, "", err
	}
	retryRequest.WithContext(ctx)

	resp, err := retryClient.Do(retryRequest)
	if err != nil {
		return nil, "", err
	}

	wrappedResponse, err := wrapResponse(resp)
	if err != nil {
		return nil, "", err
	}

	id, err := resource.NewUniqueHex("", 0, 0)

	return wrappedResponse, id, nil
}

func wrapResponse(resp *http.Response) (*response, error) {
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &response{
		Status:     resp.Status,
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
		Body:       string(respBody),
	}, nil
}
