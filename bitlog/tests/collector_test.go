package tests

import (
	"context"
	"log"
	"strings"
	"testing"
	"time"

	"github.com/1uvu/bitlog/collector"
	"github.com/1uvu/bitlog/collector/event"
	"github.com/1uvu/bitlog/collector/rpcclient"
	// "github.com/1uvu/bitlog/collector/handler"
	"github.com/1uvu/bitlog/pkg/config"
	"github.com/1uvu/bitlog/pkg/utils"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/fsnotify/fsnotify"
)

func TestCollector(t *testing.T) {
	confPath := "../config/_example/collector_config.yaml"

	conf, err := config.NewCollectorConfig(confPath, "yaml")
	if err != nil {
		t.Error(err.Error())
		return
	}
	_collector, err := collector.NewCollector(context.Background(), conf)
	if err != nil {
		t.Error(err.Error())
		return
	}
	// logHandler, err := _collector.logHandlerLog(0, nil, "")
	logHandler, err := _collector.DefaultHandlerLog()
	if err != nil {
		t.Error(err.Error())
		return
	}
	// loop and listen
	logHandler.Loop(func(ce event.CollectorEvent) {
		e := ce.EventFS
		if e.Op&fsnotify.Write == fsnotify.Write {
			line, err := utils.ReadLastLine(e.Name)
			if err != nil {
				t.Errorf("read line got error: %s.\n", err.Error())
			}
			// base the line parse result to deal with other logHandler
			// wg.Add() should be written in parent goroutine
			t.Logf("LoggingEvent. modified file: %s. last line: %s", e.Name, strings.TrimSpace(line))
		}
	})

	// test rpc logHandler
	rpclogHandler, _ := _collector.DefaultHandlerRPC()

	go func() {
		timer := time.NewTimer(5 * time.Second)
		select {
		case <-timer.C:
			reply := new(rpcclient.ConnCallReply)
			rpclogHandler.Call("getbestblock", nil, reply)
			// TODO test the rpc logHandler
			ha := reply.Reply[0].(*chainhash.Hash)
			d := reply.Reply[1].(int32)
			err = reply.Err
			if err != nil {
				t.Error(err.Error())
				// return
			}
			t.Log(ha, d)
		}
	}()

	reply := new(rpcclient.ConnCallReply)
	rpclogHandler.Call("getbestblock", nil, reply)
	// TODO test the rpc logHandler
	ha := reply.Reply[0].(*chainhash.Hash)
	d := reply.Reply[1].(int32)
	err = reply.Err
	if err != nil {
		t.Error(err.Error())
		// return
	}
	t.Log(ha, d)

	// watch file
	logClient := _collector.ClientMgr().ClientLog()
	logClient.Watch("/home/rise1007/Projects/Bitcoin/btcd/output.log")
	// For this example gracefully shutdown the client after 10 seconds.
	// Ordinarily when to shutdown the client is highly application
	// specific.
	log.Println("Log Client shutdown in 5 seconds...")
	time.AfterFunc(time.Second*5, func() {
		t.Log("Log Client shutting down...")
		logHandler.Shutdown()
		t.Log("Log Client shutdown complete.")
	})

	// run and wait
	logHandler.RunAndWait()
}
