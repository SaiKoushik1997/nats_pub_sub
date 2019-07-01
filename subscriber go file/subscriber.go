//SUBSCRIBER 2:

package main

import (
        //"log"
        "time"
	"fmt"
        //"github.com/nats-io/go-nats"
)

func subscribe_f5(subject string)string{
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

	return msg.Data


}

func main() {
	var trigger int
 
	
	for trigger!=1{ 
	
	fmt.Println("enter subject : ")
	reader := bufio.NewReader(os.Stdin)
	subject,_ := reader.ReadString('\n')
	subject = strings.Replace(subject,"\n","",-1)

	msg := subscribe_f5(subject)
	fmt.Println(msg)
        
	//trigger
	fmt.Println("Enter 0 for wait again\n		1 for exit\n ")
	fmt.Scan(&trigger)
	} 
	

        
// [end subscribe_sync]
