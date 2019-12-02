/**
 * @Author: DollarKillerX
 * @Description: ping_test.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午2:26 2019/11/26
 */
package test

import (
	"context"
	"log"
	"net"
	"sync"
	"testing"
	"time"
)

func TestPing(t *testing.T) {
	domains := []string{
		"baidu.com",
		"dollarkiller.com",
		"ppc.c",
	}

	for _, v := range domains {
		timeout, _ := context.WithTimeout(context.TODO(), time.Millisecond*200)
		ns, err := net.DefaultResolver.LookupHost(timeout, v)
		if err != nil {
			log.Println("err == = = = = == ")
			log.Println(err)
			continue
		}

		log.Println(ns)
	}
}

func TestChannel(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(2)
	bools := make(chan bool, 0)
	go func() {
		defer wg.Done()
		for {
			select {
			case _, ok := <-bools:
				if !ok {
					log.Println("00101")
					return
				}
			}
		}
	}()

	go func() {
		defer wg.Done()
		time.Sleep(3 * time.Second)
		close(bools)
	}()

	wg.Wait()
}

func TestChannel2(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(2)
	bools := make(chan bool, 3)
	go func() {
		defer wg.Done()
		for {
			select {
			case o, ok := <-bools:
				time.Sleep(time.Second)
				if ok {
					log.Println(o)
				} else {
					log.Println("00101")
					return
				}
			}
		}
	}()

	go func() {
		defer wg.Done()
		bools <- true
		bools <- true
		bools <- true
		close(bools)
	}()

	wg.Wait()
}

func TestTime(t *testing.T) {
	ticker := time.NewTicker(time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				log.Println("cc")
			}
		}
	}()

	time.Sleep(time.Second * 10)
}
