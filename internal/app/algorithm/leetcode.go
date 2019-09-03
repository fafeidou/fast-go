package algorithm

import (
	"context"
	"fmt"
	"strconv"
	"sync"
)

func BubbleSort(array [] int) {
	for i := 1; i < len(array); i ++ {
		for j := 0; j < len(array)-i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	printArray(array)
}

func SelectSort(array [] int) {
	for i := 0; i < len(array)-1; i ++ {
		min := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[min] {
				min = j
			}
		}

		if i != min {
			array[min], array[i] = array[i], array[min]
		}
	}
	printArray(array)
}

func SearchTarget(array [] int, target int) (int, int) {
	record := map[int]int{}

	for index, value := range array {
		record[value] = index
		tmp := target - value
		_, ok := record[tmp]
		if ok {
			return record[tmp], record[value]
		}
	}
	return 0, 0
}

func IsPalindrome(num int) bool {
	str := strconv.Itoa(num)
	reverse := reverse(str)
	fmt.Println(str == reverse)
	return str == reverse
}

func reverse(s string) string {
	s1 := []rune(s)
	for i := 0; i < len(s1)/2; i++ {
		s1[i], s1[len(s1)-1-i] = s1[len(s1)-1-i], s1[i]

	}
	return string(s1)
}

func printArray(array []int) {
	for index, value := range array {
		fmt.Println(index, value)
	}
}

func reverseString(str string) string {
	s1 := []rune(str)
	j := len(s1) - 1
	for i := 0; i < j; {
		s1[i], s1[j] = s1[j], s1[i]
		i++
		j--
	}
	return string(s1)
}

func producer(ch chan int) {
	for i := 0; i < 10; i ++ {
		fmt.Println("sender:=>>", i)
		ch <- i
	}
}

func consumer(ch chan int) {
	for i := 0; i < 10; i ++ {
		v := <-ch
		fmt.Println("receiver: =>>", v)
	}
}

func printOdd(ch chan byte, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 9; i += 2 {
		fmt.Println(i)
		ch <- 1
		//<-ch
	}
}

func printEven(ch chan byte, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		<-ch
		fmt.Println(i)
		//ch <- 1
	}
}

var total struct {
	sync.Mutex
	value int
}

func worker(wg *sync.WaitGroup) {
	//defer wg.Done()

	for i := 0; i <= 10; i++ {
		//total.Lock()
		total.value += i
		//total.Unlock()
	}
}

func workerContext(ctx context.Context, wg *sync.WaitGroup, name int) error {
	defer wg.Done()
	for {
		select {
		default:
			fmt.Printf("%d,%s\n", name, "运行中")
		case <-ctx.Done():
			return ctx.Err()
		}
	}

}

func testDefer() int {
	i := 1
	defer func() {
		i = 2
	}()
	return i
}

func deferFuncReturn() (result int) {
	i := 1

	defer func() {
		result++
	}()

	return i
}
