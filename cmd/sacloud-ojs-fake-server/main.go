// Copyright 2022-2023 The sacloud/object-storage-api-go authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	objectstorage "github.com/sacloud/object-storage-api-go"
	"github.com/sacloud/object-storage-api-go/fake"
	"github.com/sacloud/object-storage-api-go/fake/server"
	"github.com/spf13/cobra"
)

var (
	listenAddr    string
	dataFile      string
	outputExample bool
)

//go:embed example-data.json
var defaultData []byte

var cmd = &cobra.Command{
	Use:          "sacloud-ojs-fake-server",
	Short:        "Start the fake API server",
	RunE:         run,
	Version:      objectstorage.Version,
	SilenceUsage: true,
}

func init() {
	cmd.Flags().StringVarP(&listenAddr, "addr", "", ":8080", "the address for the server to listen on")
	cmd.Flags().StringVarP(&dataFile, "data", "", "", "the file path to the fake data JSON file")
	cmd.Flags().BoolVarP(&outputExample, "output-example", "", false, "the flag to output a fake data JSON example")
}

func main() {
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	if err := cmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, _ []string) error {
	if outputExample {
		fmt.Println(string(defaultData))
		return nil
	}

	ctx := cmd.Context()
	errCh := make(chan error)

	fmt.Printf("starting fake server with %s\n", listenAddr)
	go func() {
		errCh <- startServer(listenAddr, dataFile)
	}()

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		fmt.Println("shutting down")
	}
	return ctx.Err()
}

func startServer(addr, dataFile string) error {
	var engine fake.Engine
	fakeData := defaultData

	if dataFile != "" {
		data, err := os.ReadFile(dataFile)
		if err != nil {
			return err
		}
		fakeData = data
	}
	if err := json.Unmarshal(fakeData, &engine); err != nil {
		return err
	}

	fakeServer := server.Server{
		Engine: &engine,
	}
	httpServer := &http.Server{
		Handler:           fakeServer.Handler(),
		ReadHeaderTimeout: time.Second,
	}

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer listener.Close()

	return httpServer.Serve(listener)
}
