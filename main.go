package main

import (
	"coin"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

type experiment struct {
	first, rand, min float64
}

func main() {
	f, err := os.Create("log.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	n := 100000
	c := make(chan experiment)
	for i := 0; i < n; i++ {
		go func() {
			v := run()
			c <- v
		}()
	}
	for i := 0; i < n; i++ {
		v := <-c
		fmt.Fprintf(f, "v_1 = %f, v_rand = %f, v_min = %f\n", v.first, v.rand, v.min)
	}
}

func run() experiment {
	n_coins := 1000
	n_flips := 10
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	c_rand := r.Intn(n_coins)

	var v_1, v_rand, v_min float64
	v_min = float64(n_flips)

	for i := 0; i < n_coins; i++ {
		c := coin.New()

		heads := 0
		for j := 0; j < n_flips; j++ {
			f := c.Flip()
			if f == coin.Heads {
				heads++
			}
		}
		heads_f := float64(heads) / float64(n_flips)
		if i == 0 {
			v_1 = heads_f
		}
		if i == c_rand {
			v_rand = heads_f
		}
		if heads_f < v_min {
			v_min = heads_f
		}
	}

	return experiment{first: v_1, rand: v_rand, min: v_min}
}
