package main

import "context"
import "fmt"

type ProcessHandler func(ctx context.Context) error

type ProcessHandlerWrapper func(ProcessHandler) ProcessHandler

func SayHello(ctx context.Context) error {
	levelNum := ctx.Value("levelNum").(int)
	var i int = 1
	for i <= levelNum {
		key := fmt.Sprintf("level-%d", i)
		val := ctx.Value(key).(int)
		fmt.Printf("%s:%d\n", key, val)
		i++
	}
	return nil
}

func NewWrapperSayHello(levelNum int) ProcessHandlerWrapper {
	return func(ph ProcessHandler) ProcessHandler {
		return func(ctx context.Context) error {
			key := fmt.Sprintf("level-%d", levelNum)
			ctx = context.WithValue(ctx, key, levelNum)
			fmt.Printf("wrapper %s:%d\n", key, levelNum)
			return ph(ctx)
		}
	}
}

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "levelNum", 3)

	var wrapperHandler []ProcessHandlerWrapper
	for i := 0; i < 3; i++ {
		wrapperHandler = append(wrapperHandler, NewWrapperSayHello(i+1))
	}

	fn := SayHello
	for i := 3; i > 0; i-- {
		fn = wrapperHandler[i-1](fn)
	}

	fn(ctx)
}
