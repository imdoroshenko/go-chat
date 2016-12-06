package bootstrap

import (
	"github.com/mediocregopher/radix.v2/pool"
	"log"
)

var Redis, err = pool.New("tcp", "localhost:6379", 10)

func init () {
	//Pool,err = pool.New("tcp", "localhost:6379", 10)
	if err != nil {
		log.Fatal(err)
	}

	// In another go-routine
	//conn, err := p.Get()
	//if err != nil {
	//	// handle error
	//}
	//
	//if conn.Cmd("SOME", "CMD").Err != nil {
	//	// handle error
	//}
	//
	//p.Put(conn)
}

