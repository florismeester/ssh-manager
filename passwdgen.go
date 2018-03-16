package main

/* Generate a random password for ssh-manager
   @author Floris Meester floris@rawpacket.com */

import (
	"crypto/rand"
)



const encodestr = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_!@#$%^&*?~"

func genpw() string {
	
	n := configuration.Pwlen
     	pw := make([]byte, n)

    	randstr := make([]byte, n)
	
	// we use urandom, which is good enough for this purpose
    	_, err := rand.Read(randstr)
	printerr(err)

	// stuff it in  the passwd array
	for ch := range pw {
      	 	random := uint8(randstr[ch])
        	randch := random % uint8(len(encodestr))
        	pw[ch] = encodestr[randch]
    	}
    	return string(pw)
}
