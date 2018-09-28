package handler

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/jwuensche/autobahnausfahrt/util"
)

// Package Const
const (
	prometheusTag = "prom"
)

type metrics struct {
	RecvMessageCount      uint64            `prom:"messages_count_received_total"`
	SendMessageCount      uint64            `prom:"messages_count_send_total"`
	RecvTrafficBytesTotal uint64            `prom:"traffic_bytes_received_total"`
	SendTrafficBytesTotal uint64            `prom:"traffic_bytes_send_total"`
	Authentication        map[string]auth   `prom:"authentications"`
	AuthRolesClients      map[string]uint64 `prom:"authorization_roles"`
	SuccededAuthorization uint64            `prom:"authorizations_succeded"`
	RejectedAuthorization uint64            `prom:"authorizations_rejected"`
}

type auth struct {
	Succeded uint64 `prom:"authentications_succeded"`
	Rejected uint64 `prom:"authentications_rejected"`
}

// Render handles all incoming request for metric routes depending on the defined scraper, default prometheus
func Render(w http.ResponseWriter, r *http.Request) {
	cont, err := scrape()
	if err != nil {
		util.Log.Critical("Rendering failed: %v", err)
		w.Write([]byte("Oops"))
		w.WriteHeader(503)
		return
	}

	raw, err := cont.prometheusExport()

	if err != nil {
		util.Log.Critical("Rendering failed: %v", err)
		w.Write([]byte("Oops"))
		w.WriteHeader(503)
		return
	}

	// success
	w.Write([]byte(raw))
	return
}

func scrape() (cont metrics, err error) {
	r, err := http.Get("http://" + util.Conf.InterconnectAddress + ":" + util.Conf.InterconnectPort)
	if err != nil {
		util.Log.Critical("Failed to connect to interconnect: %v", err)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		util.Log.Critical("Scraping failed: %v", err)
		return
	}

	err = json.Unmarshal(data, &cont)
	if err != nil {
		util.Log.Critical("Scraping failed: %v", err)
		return
	}
	return
}

// export for prometheus specific endpointa
func (b *metrics) prometheusExport() (raw []byte, err error) {
	writer := bytes.Buffer{}
	content := bufio.NewWriter(&writer)
	fields := reflect.TypeOf(*b)
	for i := 0; i < fields.NumField(); i++ {
		if _, ok := (reflect.ValueOf(*b).Field(i).Interface()).(uint64); ok != true {
			// promSubExport(content, fields.Field(i).Type, reflect.ValueOf(*b).Field(i))
		} else {
			util.Log.Debugf("Tag: %s, Value: %d", fields.Field(i).Tag.Get(prometheusTag), reflect.ValueOf(*b))
			fmt.Fprintf(content, "%s %d\n", fields.Field(i).Tag.Get(prometheusTag), (reflect.ValueOf(*b).Field(i).Interface()).(uint64))
		}
	}
	// Done by now, flush into buffer
	content.Flush()
	raw = writer.Bytes()
	return
}

// semi-recursive approach for all sub categories
func promSubExport(writer io.Writer, fields reflect.Type, values reflect.Value) {
	for i := 0; i < fields.NumField(); i++ {
		if _, ok := (values.Field(i).Interface()).(uint64); ok != true {
			promSubExport(writer, fields.Field(i).Type, values.Field(i))
		}
		fmt.Fprintf(writer, "%s %d\n", fields.Field(i).Tag.Get(prometheusTag), (values.Field(i).Interface()).(uint64))
	}
}
