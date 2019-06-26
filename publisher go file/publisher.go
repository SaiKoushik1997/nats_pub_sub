//PUBLISHER 2

package main

import (
        "fmt"
        "log"
	      "bufio"
	      "os"
	      "strings"
        "github.com/nats-io/go-nats"
)

func publish_f5(subject string, msg string) {
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
func main(){
	var trigger int	
	//msg := []byte("All is Well")
	//publish_f5("updates",msg)
	
	//trigger
	for trigger!=1{
	fmt.Println("enter subject : ")
	reader := bufio.NewReader(os.Stdin)
	subject,_ := reader.ReadString('\n')
	subject = strings.Replace(subject,"\n","",-1)

	fmt.Println("enter msg : ")
	msg,_ := reader.ReadString('\n')
	msg = strings.Replace(msg,"\n","",-1)
		publish_f5(subject, msg)

	fmt.Println("Enter 0 for wait again\n		1 for exit")
	fmt.Scan(trigger)
	
}
}
