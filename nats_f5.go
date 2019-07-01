package nats_f5


import (
        "fmt"
        "log"
	//"bufio"
	//"os"
	//"strings"
	"time"
        "github.com/nats-io/go-nats"
)


func Publish_f5(subject string, msg string) {
        // [begin publish_bytes]
        nc, err := nats.Connect(nats.DefaultURL)
        if err != nil {
                log.Fatal(err)
        }
        fmt.Println("working")
        //defer nc.Close()

        if err := nc.Publish(subject, []byte(msg)); err != nil {
                log.Fatal(err)
        }
        // [end publish_bytes]
        // Make sure the message goes through before we close
        nc.Flush()
}


func Subscribe_f5(subject string)string{
	nc, err := nats.Connect(nats.DefaultURL)
        if err != nil {
                log.Fatal(err)
        }
        defer nc.Close()

        // Subscribe
        sub, err := nc.SubscribeSync(subject)
        if err != nil {
                log.Fatal(err)
        }
	
	// Wait for a message
        msg, err := sub.NextMsg(30 * time.Second)
        if err != nil {
                log.Fatal(err)
        }

        // Use the response
        //log.Printf("Reply: %s", msg.Data)

	// Close the connection
        nc.Close()

	return string(msg.Data)


}
