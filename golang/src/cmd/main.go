package main

import (
	"fmt"
	"os"
	"log"
	"os/signal"
	"syscall"
	"strings"
	"./srv"
	"../config"
	"github.com/nic-chen/nice/micro/registry"
	"github.com/nic-chen/nice/micro/tracing"
	_ "github.com/nic-chen/nice/micro/registry/etcdv3"
	opentracing "github.com/opentracing/opentracing-go"
)

func usage() {
	fmt.Fprintf(os.Stderr, "nice examples\n")
	fmt.Fprintf(os.Stderr, "USAGE\n")
	fmt.Fprintf(os.Stderr, "  nice-test command \n")
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "The commands are\n")
	fmt.Fprintf(os.Stderr, "  all          Boots all services\n")
	fmt.Fprintf(os.Stderr, "  api          Api gateway\n")
	fmt.Fprintf(os.Stderr, "\n")
}

func main() {

	var (
		register registry.Registry
		tracer opentracing.Tracer
		err      error
	)

	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	var run func(registry.Registry, opentracing.Tracer)

	switch cmd := strings.ToLower(os.Args[1]); cmd {
	case "all":
		run = srv.RunAll
	case "api":
		run = srv.RunApi
	case "member":
		run = srv.RunMemberSrv		
	default:
		usage()
		os.Exit(1)
	}

	if config.SrvName != "" && config.SrvHost != "" && config.SrvPort != "" && config.NamingAddr != ""  {
		options := &registry.Options{
			Name: config.SrvName,
			Host: config.SrvHost,
			Port: config.SrvPort,
			TTL: config.SrvCheckTTL,
			Ssrv: config.NamingAddr,
		}
		register, err = registry.DefaultRegistry(options)
		log.Printf("NamingAddr: %s", config.NamingAddr)
		if err != nil {
			panic(err)
		}
		tracer, err = tracing.Init(config.SrvName, config.JaegerAddr)
		if err != nil {
			panic(err)
		}
		//监听退出
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)

		go func() {
			<-ch
			register.UnRegister()
			os.Exit(1)
		}()
	}

	run(register, tracer);
}
