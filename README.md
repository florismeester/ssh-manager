

<h1>SSH encrypted key generator</h1>    
 

HTTP/HTTPS service for generating public/private RSA keypairs. Public key is in authorized_keys file format.  
It can save to local files and/or output to JSON format for use in automation tools like Puppet, Salt, Chef, Ansible  
or whatever you prefer.    
I needed something like this for PCI compliance at a customer site (PCI DSS 3.2 requirement 3.6.4), regular key rollover/rotation. Run it in a cronjob  
with curl/wget.  Compiled Linux AMD64 version is in the bin directory for demo purposes.  

Configuration:  
	

        "file": true		--> output to local files (will output <username>-id_rsa.pub and <username>-id_rsa)  
        "json": true		--> output to JSON through HTTP/HTTPS, obviously you should use the latter.When false only the generated passphrase is returned.  
        "bits": 2048		--> number of bits  
        "listen": ":443"	--> ip:port to listen on  
        "path": "keyfiles"	--> path to save keypairs  
        "tls": true		--> this *should* be true  
        "keypath": "key.pem"	--> TLS Key file  
        "certpath": "cert.pem"	--> TLS Cert file  
        "pwlen": 14		--> passwphrase length  



Usage:  
	 Start the service with: ./ssh-manager & or create a systemd service unit  
	 curl  https://localhost:443/keygen/alice  

Output:   

{"Status":"succes","Username":"alice","Passwd":"oMgr6d7ElWCKK?","Publickey":"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCtXnDlOyc06UnDGnr8vsDfeUG8Jk9MzOOQhAgRUnUyyytgzNv76dP7kGQ6blswds3qeMkGVleF47w1f0cahhlBNZHGWNVgFS8jm0qvSqdfR/8OlUrdI60sVwDIISqoKWCl8uS0pIvBgXGCB1/HtsyImOfL+lxHN1h/2Xej+HGNDtDR1+UY7pR00TR+Envvw0Uelm62Ia5eXFadCxmiZ6k33/TNodgCaIsIT2K5EYyqOgWLwZb2u1OeY+P9zH0F2nSbiklFkIZJQGfw+XUtTg2LRVxpJcYsxYIETT+44qWC/s0l8ry9mzeF6xz/VnCokt6Ri6ITK/IaxlxtRFz086Eb\n","Privatekey":"-----BEGIN RSA PRIVATE KEY-----\nProc-Type: 4,ENCRYPTED\nDEK-Info: AES-256-CBC,4f0feb1a5229f27d8a0f445670b51d9f\n\nlaqzn2Fc13BY+oVs/BFlGol59KGuO9VVAmv8DeO7SuCKIxAVAe/VR9B0R5FrjOKZ\nJY21EfnVjg01fbb0sO9QQwF3e2yT+sZFfCh1RIBfhNdwwMMSG2NWIkgQylR+Mby8\nGVpX7sO7kYE5OovModvZcjOMx19x30RBomkpptw4jKImvH1LP8d+Dp4n5gvkvjFk\nlCLwT2XnYiV4Tv+fQ+MQBe23NHQ03BlMbzQgJ7tO+GTGlXUGtqrRrJBaZskal/WG\nz1XkkthaW7w1y4OEiN7/4aOo35vE4KI4DT4S4JOjwZq63tHwlela5w2F/mKnnMZa\n3rVuujF/7ir1vyiyih4WH3cF7CUR2NnkG2HG3ewJwrlKYTE1b3bywKxkGt6RkbxE\nQoIAuR02mbifIBHzCjXG29S5LjgcNYhoI6LyOQHJD1mZc3ohfZce6sI1LB5Z8Y86\ns7sKMMx66G9jXs7ms+wSmvGybaaODy+hN7OmqQz6nl2k5IzVZShIv3eNdWMl9fE/\nQ7wJsu3DhG/3mMb/Tisz81RkQsSqsJwpQJSXYOiRf8CptIFoIj//z0NDL2IScOQA\nmOvDS6tUZ5mCLiysYHynPi0zpvb/0El6PIJ6XTtai+MtbE5ctD93ZI6gVED0lrbF\n76TsUlJPdAW6C+3PFnQv0RTiEfM0lfSiFZUjZxAFm/4LK6IMIg+VkaizFMTfCF+F\ng4DL3NXDyKPULcOB15TnrDvK2k/zKCp6MVhTsFHVbHW+tuBgPBUFQm5korEeKk66\noJ6dy2ymgH6/sLJq6IaUM7VHIHZpe9N/CeACi0r7nYulGQeeZCjA5bKXgGC/eZGn\nBJj2IpcPaTWvsSi89MygnYyZ+ho0h4wWrRQYNMHyqvtbXdlvxu62s+jGP3qOLwQf\nQYWEyCaDfaQKrOERzYtNK6BJSqK/WBJAjsr2xsw3UU7Bxrn8J9zDMw9i1NiJfN0h\nQ6oAOzs3glUIm3W09qIZFRUs1CrwrIqrg0ybCfYEBPnjwT6q5b3ItMTNOk5/51IL\n0wctudp8cyqh5/9GDIlw3Bk8QefDpJVq+Kj0bkBvaYRj8GmaRMpcbanC6UpjXlUv\nWNv8QaK6HKbUqQoltx7ioNW9H+81tUlCetUy8XnJvy6AFHUnkI/ZJsMRtvUXXcwF\nTbpbTH/L3MbzxD2lt1ZScHuNMZVdghAbDygFmnuYpFm88HpzvdVcFoN9VaGIPwkG\nM3Rf6sxRQgPD8z7vGrVZhjE5Chvw++lDdn8bRXI4zhV5Awl9Dplh2qiyTcecMYT8\nKtyWQnGTR79LRMbajkg3YGSJ/ZqCEE/PxmjkpQeccdXY+tLZ+hLonZ4pUUC739R8\nQDzd6IeBfc/kMQqwJZbCwVHLH9pjnNw97iWzsRLZ0UaeGGj7BXrWqjBTf5m9GxIf\nNlUooK0Xhr6K8ODa6DK9IG1nxMmnZYfEAZznubX2+JADdNiySObAqVix6hKdM8rn\n2dlVu8QGRoIwOhn/Gkm7NNbNdBHRFGmasSS+/Tci3nlck47A96mHQx1ytEi08YFa\no9fzRS/zGV+8PyZuihl6pyK8lXG7papmvB3dPwXmtwz6Mw8oQuTqQw3KsN7dza6S\n-----END RSA PRIVATE KEY-----\n"}  


