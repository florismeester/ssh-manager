package main

/* ssh keypair generation, numer of bits can be set in config file.
   Output can be set to file, JSON or both.
   @author Floris Meester floris@rawpacket.com */

import ( 
	"crypto/rsa" 
	"crypto/rand" 
	"encoding/pem" 
	"crypto/x509" 
	"golang.org/x/crypto/ssh" 
	"io/ioutil"
	"os"
	)





func createKeyPair(pubPath, privPath, passwd string) Message {
    
  	var message Message

	privKey, err := rsa.GenerateKey(rand.Reader, configuration.Bits)
	printerr(err)
   
	pw := genpw()
	message.Passwd = pw
 	// generate  private key
 	priv := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privKey)}
	priv, err = x509.EncryptPEMBlock(rand.Reader, priv.Type, priv.Bytes, []byte(passwd), x509.PEMCipherAES256)
        printerr(err)
	
	if configuration.File {
		privFile, err := os.Create( configuration.Path + "/" + privPath)
		printerr(err)

		// write out the key with the right file permission mask
 		err = os.Chmod(configuration.Path + "/" + privPath, 0400)
		printerr(err)
		defer privFile.Close()
		err = pem.Encode(privFile, priv)
		printerr(err)
	}
	if configuration.Json{
		plainpriv := string(pem.EncodeToMemory(priv))
		message.Privatekey = plainpriv
	}

	// generate  public key
	pub, err := ssh.NewPublicKey(&privKey.PublicKey)
 	printerr(err)
	if configuration.Json{
		plainpub := string(ssh.MarshalAuthorizedKey(pub))	
		message.Publickey = plainpub
	}
	// write out the key with the right file permission mask
	if configuration.File {
		err =  ioutil.WriteFile( configuration.Path + "/" + pubPath, ssh.MarshalAuthorizedKey(pub), 0644)
		printerr(err)
	}
	return message
}
