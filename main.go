package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"os"

	"github.com/joho/godotenv"
	"github.com/pasannissanka/network_go/lib/client"
	"github.com/pasannissanka/network_go/lib/net"
	"github.com/pasannissanka/network_go/lib/server"
)

type Environment struct {
	Id        int
	IS_MASTER bool
	Ip        string
	TCP_PORT  int
	UDP_PORT  int
	L_PORT    int
}

var Env Environment

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Environment not provided")
		os.Exit(1)
	}

	env := os.Args[1]
	processEnv(env)

	serverData := server.Server{
		IP:    Env.Ip,
		UDP_L: uint16(Env.UDP_PORT),
		TCP:   uint16(Env.TCP_PORT),
	}

	var wg sync.WaitGroup
	wg.Add(1)

	if !Env.IS_MASTER {
		go func() {
			defer wg.Done()
			connectToMaster()
		}()
	} else {
		go func() {
			defer wg.Done()
			server.Init(serverData, Env.Id)
		}()
	}

	wg.Wait()
}

func processEnv(env string) {
	err := godotenv.Load(fmt.Sprintf(".env.%s", env))

	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
		os.Exit(1)
	}

	sId := os.Getenv("ID")
	tcpPort := os.Getenv("TCP_PORT")
	udpPort := os.Getenv("UDP_PORT")
	lPort := os.Getenv("L_PORT")
	isMaster := os.Getenv("MASTER")

	server_id, err := strconv.Atoi(sId)
	if err != nil {
		log.Fatalf("Environment variables - server id not found: %s", err)
		os.Exit(1)
	}

	tcp_port, err := strconv.Atoi(tcpPort)
	if err != nil {
		log.Fatalf("Environment variables - TCP port not found: %s", err)
		os.Exit(1)
	}

	udp_port, err := strconv.Atoi(udpPort)
	if err != nil {
		log.Fatalf("Environment variables - UDP port not found: %s", err)
		os.Exit(1)
	}

	l_port, err := strconv.Atoi(lPort)
	if err != nil {
		log.Fatalf("Environment variables - L port not found: %s", err)
		os.Exit(1)
	}

	is_master, err := strconv.ParseBool(isMaster)
	if err != nil {
		log.Printf("Environment variables - IS_MASTER not found: %s", err)
		is_master = false
	}

	if is_master {
		log.Printf("Server is master\n")
	}

	ip := getLocalIP()

	Env = Environment{
		Id:        server_id,
		Ip:        ip,
		TCP_PORT:  tcp_port,
		UDP_PORT:  udp_port,
		L_PORT:    l_port,
		IS_MASTER: is_master,
	}
}

func getLocalIP() string {
	ip, err := net.ExternalIP()

	if err != nil {
		log.Fatalf("error getting local IP: %s", err)
		os.Exit(1)
	}

	log.Printf("Local IP: %s\n", ip)
	return ip
}

func connectToMaster() {
	client.EnableTestMode(&client.TestModeOptions{
		PortStart: 8880,
		PortEnd:   8890,
	})

	for {
		client.Scan(fmt.Sprintf("%s/24", Env.Ip))
		time.Sleep(5 * time.Minute)
	}
}
