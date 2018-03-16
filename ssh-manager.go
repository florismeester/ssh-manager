package main

/* ssh-manager: generate passwd protected RSA ssh keypair to files and/or json with an HTTP/HTTPS request
   @author Floris Meester floris@rawpacket.com */


    import (
	"fmt"
	"os"
	"log"
	"encoding/json"
	"flag"
	"sync"
	"log/syslog"
    )

type Configuration struct {
	Debug bool
	File bool
	Json bool
	Bits int
	Listen string
	Path string
	Tls bool
	Keypath string
	Certpath string
	Pwlen int
} 

type Message struct {
	Status string
	Username string
	Passwd string
	Publickey string
	Privatekey string
}


var configuration Configuration

func main(){

	// Create a sync waitgroup to prevent exit from our go routines
	var wg sync.WaitGroup

        // Open and parse configuration file
        conf := flag.String("config","ssh-manager.conf", "Path to configuration file")
        flag.Parse()
        file, err := os.Open(*conf)
        if err != nil {
                log.Fatal("Can't find configuration file ", *conf)
        }
        decoder := json.NewDecoder(file)
        // configuration:= Configuration{}
        err = decoder.Decode(&configuration)
        if err != nil {
                fmt.Println("error opening configuration:", err)
        }
	// Create a syslog writer for debug logging
	logwriter, e := syslog.New(syslog.LOG_NOTICE, "ssh-manager")
    		if e == nil {
        	log.SetOutput(logwriter)
    	}
	
	wg.Add(1)
	go  restapi()
	wg.Wait()

	

}

