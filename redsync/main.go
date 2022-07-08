package main

import (
	"fmt"
	"time"

	"log"

	goredislib "github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
)

type UnLock interface {
	UnLock() error
}

type lock struct {
	mutex *redsync.Mutex
}

func newLock(client *goredislib.Client, ns, k string) *lock {
	pool := goredis.NewPool(client)
	rs := redsync.New(pool)
	expiryOption := redsync.WithExpiry(10 * time.Second)
	return &lock{
		mutex: rs.NewMutex(ns+":"+k, expiryOption),
	}
}

func (l *lock) UnLock() error {
	if _, err := l.mutex.Unlock(); err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func newKeyLock(client *goredislib.Client, k, ns string) UnLock {
	l := newLock(client, "lock:"+ns, k)
	err := l.mutex.Lock()
	if err != nil {
		log.Println(err.Error())
	}

	return l
}

func CreateObj(client *goredislib.Client, key string) UnLock {
	return newKeyLock(client, key, "create_obj")
}

func fn(i int) {
	client := goredislib.NewClient(&goredislib.Options{
		Addr: "192.168.1.234:6379",
	})

	defer CreateObj(client, "house").UnLock()
}

func main() {
	go fn(1)
	go fn(2)
	go fn(3)
	go fn(4)

	time.Sleep(5 * time.Second)
	fmt.Println("ok")
}
