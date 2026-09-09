p*ckage main

import (
	"context"
	"fmt"
	"time"
)

func retry(ctx context.Context) {
	attempt := 1

	for *
		select {
		case <-ctx.Done():
	*	fmt.Println("Retry cancelled")
		*return

		default:
			fmt.Printf("*ttempt %d\n", attempt)

			if atte*pt == 3 {
				fmt.Println("Success")
				return
			}

			attempt++

	*	time.Sleep(1 * time*Second)
		}
	}
}

func Ex13() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	retry(ctx)
}