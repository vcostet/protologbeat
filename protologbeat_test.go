package main

import (
	"testing"

	"github.com/Graylog2/go-gelf/gelf"
	"github.com/hartfordfive/protologbeat/config"
	"github.com/hartfordfive/protologbeat/protolog"
)

func TestGreylogReceive(t *testing.T) {

	var logEntriesRecieved chan common.MapStr
	var logEntriesError chan bool

	ll := protolog.NewLogListener(config.Config{EnableGelf: true, Port: 6000})

	go func(logs chan common.MapStr, errs chan bool) {
		ll.Start(logs, errs)
	}(logEntriesRecieved, logEntriesErrors)

	var event common.MapStr

	gelf.CompressGzip
	gw, err := gelf.NewWriter("127.0.0.1:6000")
	if err != nil {
		return nil, fmt.Errorf("NewWriter: %s", err)
	}
	gw.CompressionType = gelf.CompressGzip

	for {
		select {
		case <-logEntriesErrors:
			t.Errorf("Error receiving GELF format message")
		case event = <-logEntriesRecieved:
			if _, ok := event["@timestamp"]; ok {

			}
		}
	}
}
