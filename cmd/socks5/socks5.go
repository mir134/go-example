package main

import (
	"log"
	"os"
	"github.com/mir134/go-example/pkg/socks5"
	"github.com/mir134/go-example/pkg/yxydes"
	//"fmt"
	"flag"
)

func main() {
	//desStr := yxydes.Encrypt("192.168.63.85", "zzm12345")
	//fmt.Println(desStr)
	//ip := yxydes.Decrypt(desStr, "zzm12345")
	//fmt.Println(ip)
	listenAddr := flag.String("addr", "127.0.0.1:8088", "Input server listen address:")
	username := flag.String("u", "test123", "Input server proxy username:")
	password := flag.String("p", "tesQWss33.", "Input server proxy password:")
	key := flag.String("key", "zzm12345", "Input server proxy password:")
	yxydes.SetKey(*key)
	creadentials := socks5.StaticCredentials{
		//os.Getenv("USER"): os.Getenv("PASSWORD"),
		*username: *password,
	}
	authenticator := socks5.UserPassAuthenticator{Credentials: creadentials}

	// Create a SOCKS5 server
	config := &socks5.Config{
		AuthMethods: []socks5.Authenticator{authenticator},
		Logger:      log.New(os.Stdout, "", log.LstdFlags),
	}
	server, err := socks5.New(config)
	if err != nil {
		panic(err)
	}

	flag.Parse()
	log.Println("scoks5服务器正在运行中...", *listenAddr, *username, *password)
	// Create SOCKS5 proxy on localhost port 1080
	if err := server.ListenAndServe("tcp", *listenAddr); err != nil {
		panic(err)
	}
}
