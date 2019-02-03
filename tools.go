package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	karma "github.com/reconquest/karma-go"
)

func parseDSN(rawDSN string) string {
	DSN := strings.TrimSpace(rawDSN)

	if !strings.HasPrefix(DSN, "http://") &&
		!strings.HasPrefix(DSN, "https://") {

		return fmt.Sprintf("http://%s", DSN)
	}

	return DSN
}

func makeHTTPClient(caPath string) (*http.Client, error) {
	destiny := karma.Describe(
		"method", "makeHTTPClient",
	)

	if caPath == noneValue {
		return &http.Client{}, nil
	}

	cert, err := ioutil.ReadFile(caPath)
	if err != nil {
		return nil, destiny.Describe(
			"CA path", caPath,
		).Describe(
			"error", err,
		).Reason(
			"can't read CA",
		)
	}

	rootCAs, err := x509.SystemCertPool()
	if err != nil {
		return nil, destiny.Describe(
			"error", err,
		).Reason(
			"can't obtain root CA pool",
		)
	}

	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}

	if hasAppended := rootCAs.AppendCertsFromPEM(cert); !hasAppended {
		return nil, destiny.Reason("cert CA has't appended")
	}

	tlsConfig := &tls.Config{
		RootCAs: rootCAs,
	}

	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}, nil
}
