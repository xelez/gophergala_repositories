package goxa

import (
	"fmt"
	"testing"
	"time"
)

func Test_003_Write_Read_Commit(t *testing.T) {

	go func() {
		xa, err := Listen("tcp", ":8080")
		if err != nil {
			t.Errorf(err.Error())
		}
		buffer := make([]byte, 1024)

		id, buffer, count, err := Receive(xa)
		if err != nil {
			t.Errorf(err.Error())
		}

		if id != nil {
			fmt.Printf("Receiver received %d bytes: %s\n", count, buffer)
		}

		Commit(xa, id)

		Close(xa)
	}()

	time.Sleep(1 * time.Second)

	xa, err := Connect("tcp", "127.0.0.1:8080")
	if err != nil {
		t.Errorf(err.Error())
	}

	count, err := Send(xa, xa, []byte("hello"), true)
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Printf("Sender sent %d bytes.\n", count)

	time.Sleep(1 * time.Second)

	err = Disconnect(xa)
	if err != nil {
		t.Errorf(err.Error())
	}

	time.Sleep(1 * time.Second)

	return
}

func Test_004_Write_Read_Rollback(t *testing.T) {

	go func() {
		xa, err := Listen("tcp", ":8080")
		if err != nil {
			t.Errorf(err.Error())
		}
		buffer := make([]byte, 1024)

		id, buffer, count, err := Receive(xa)
		if err != nil {
			t.Errorf(err.Error())
		}

		if id != nil {
			fmt.Printf("Receiver received %d bytes: %s\n", count, buffer)
		}

		Rollback(xa, id)

		Close(xa)
	}()

	time.Sleep(1 * time.Second)

	xa, err := Connect("tcp", "127.0.0.1:8080")
	if err != nil {
		t.Errorf(err.Error())
	}

	count, err := Send(xa, xa, []byte("hello"), true)
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Printf("Sender sent %d bytes.\n", count)

	time.Sleep(1 * time.Second)

	err = Disconnect(xa)
	if err != nil {
		t.Errorf(err.Error())
	}

	time.Sleep(1 * time.Second)

	return
}