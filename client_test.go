package main

import "testing"
import "os"
import "paradise/server"
import "paradise/client"
import "math/rand"
import "time"

var file *os.File
var fileBytes []byte

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestSimple(t *testing.T) {
	go server.Start()
	time.Sleep(1 * (time.Second * 1))
	testConnect(t)
	testLots(t)
}

func testConnect(t *testing.T) {
	c := client.NewClient()
	c.Connect()
	c.List()
	c.Stor(1024)
}

func testLots(t *testing.T) {
	s1 := rand.NewSource(time.Now().UnixNano())

	for {
		rb := int64(byte(s1.Int63() * 400))
		go randClient(rb)
		time.Sleep(time.Duration(rb) * (time.Millisecond * 1))
	}
}

func randClient(size int64) {
	c := client.NewClient()
	c.Connect()
	c.List()
	c.Stor(1024 * size)
}