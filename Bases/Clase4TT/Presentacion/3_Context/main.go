// package main

// import (
// 	"context"
// 	"fmt"
// )

// func main() {
// 	ctx := context.Background()

// 	fmt.Println(ctx)
// }

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "saludo", "hola digital house!!")
	saludoWrapper(ctx)

	ctx = context.Background()
	deadline := time.Now().Add(time.Second * 5)

	ctx, _ = context.WithDeadline(ctx, deadline)

	<-ctx.Done()
	fmt.Println(ctx.Err().Error())

	time.Sleep(time.Second * 15)
	fmt.Println("Holi")
}

func saludoWrapper(ctx context.Context) {
	saludo(ctx)
}

func saludo(ctx context.Context) {
	fmt.Println(ctx.Value("saludo"))
}
