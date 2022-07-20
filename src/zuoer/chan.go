package zuoer

func Echo(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, num := range nums {
			out <- num // send num to out
		}
		close(out) // close out channel
	}()
	return out
}

func Odd(nums <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for num := range nums {
			if (num & 1) != 0 {
				out <- num
			}
		}
		close(out)
	}()
	return out
}

func Sum(nums <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		sum := 0
		for num := range nums {
			sum += num
		}
		out <- sum
		close(out)
	}()
	return out
}

func Fq(nums <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for num := range nums {
			out <- num * num
		}
		close(out)
	}()
	return out
}

type PipeFunc func(ch <-chan int) <-chan int
type EchoFunc func(nums []int) <-chan int

func Pipeline(nums []int, echo EchoFunc, pipe ...PipeFunc) <-chan int {
	ch := echo(nums)
	for _, f := range pipe {
		ch = f(ch)
	}
	return ch
}
