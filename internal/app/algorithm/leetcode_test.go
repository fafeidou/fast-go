package algorithm

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/go-gomail/gomail"
	"runtime"
	"sync"
	"testing"
	"time"
)

var i = 0

var wg sync.WaitGroup

func TestName(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	//BubbleSort([]int{4,3,2,1})
	//SelectSort([]int{4,3,2,1})

	//i, i2 := SearchTarget([]int{1, 2, 7, 15}, 9)
	//fmt.Println(i,i2)
	//s := reverseString("我是中国人")
	//fmt.Println(s)
	//wg.Add(2)
	ch := make(chan int)

	go consumer(ch)
	go consumer(ch)
	go producer(ch)
	//go counter(ch)
	//go counter(ch)
	//go counter2()
	//go counter2()

	//go incCounter(1)
	//go incCounter(2)

	//ch <- 1
	//go incCounterChannel(1, ch)
	//go incCounterChannel(2, ch)
	//wg.Wait()
	time.Sleep(time.Second * 1)
	//ch <- 1
	fmt.Println(i)

}

func TestPrintOddAndEven(t *testing.T) {
	ch := make(chan byte)
	var wg sync.WaitGroup
	wg.Add(2)
	go printOdd(ch, &wg)
	go printEven(ch, &wg)
	wg.Wait()
}

func TestPrint(t *testing.T) {

	//var wg sync.WaitGroup
	//wg.Add(2)
	go worker(&wg)
	go worker(&wg)
	//wg.Wait()
	time.Sleep(time.Second * 3)
	fmt.Println(total.value)
}

func TestWorkerContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go workerContext(ctx, &wg, i)
	}

	time.Sleep(time.Second)
	cancel()
	wg.Wait()
}

func TestMail(t *testing.T) {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", "ckc_it@163.com", "ckc_it")
	m.SetAddressHeader("To", "943104990@qq.com", "账号2")
	m.SetAddressHeader("To", "476688386@qq.com", "账号3")
	m.SetHeader("Subject", "gomail-邮件测试")
	m.SetBody("text/html", "<h1>hello world</h1>")

	d := gomail.NewDialer("smtp.163.com", 465, "ckc_it@163.com", "liangge666")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("***%s\n", err.Error())
	}

}

func TestDefer(t *testing.T) {
	fmt.Println(deferFuncReturn())
}
