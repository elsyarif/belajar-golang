package golang_goroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		// mengirim ke chanel
		channel <- "Syarif"
		fmt.Println("selesai mengirim data")
	}()

	// menerima channel
	data := <-channel
	fmt.Println(data)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Syarif Hidayatulloh"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	// menerima channel
	data := <-channel
	fmt.Println(data)
}

// channel in out
// hanya memasukan data kechannel
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Syarif Hidayatulloh"
}

// hanya menarik data/ mengambil data
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(3 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	cannel := make(chan string, 3)

	go func() {
		cannel <- "Syarif"
		cannel <- "Hidayatulloh"
	}()

	go func() {
		fmt.Println(<-cannel)
		fmt.Println(<-cannel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 11; i++ {
			channel <- "Perulangan ke:" + strconv.Itoa(i)
		}

		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data", data)
	}

	fmt.Println("===- Selesai -===")
}

func TestSelectChannel(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)
	defer close(ch1)
	defer close(ch2)

	go GiveMeResponse(ch1)
	go GiveMeResponse(ch2)

	counter := 0
	for {
		select {
		case data := <-ch1:
			fmt.Println("Data dari channel 1:", data)
			counter++
		case data := <-ch2:
			fmt.Println("Data dari channel 2:", data)
			counter++

		}

		if counter == 2 {
			break
		}

	}
}

func TestDefaultSelectChannel(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)
	defer close(ch1)
	defer close(ch2)

	go GiveMeResponse(ch1)
	go GiveMeResponse(ch2)

	counter := 0
	j := 0
	for {
		select {
		case data := <-ch1:
			fmt.Println("Data dari channel 1:", data)
			counter++
		case data := <-ch2:
			fmt.Println("Data dari channel 2:", data)
			counter++

		default:
			fmt.Println("menunggu data...", j)
			j++
		}

		if counter == 2 {
			break
		}

	}
}
