package functions

import (
	"io/ioutil"
	"log"
	"os"
)

func FileExists(f string) bool {
	info, err := os.Stat(f)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func SetDNSDir() error {
        dir, err := os.Getwd()
        if err != nil {
                return err
        }
        _, err = os.Stat(dir + "/config/dnsconfig")
        if os.IsNotExist(err) {
                os.Mkdir(dir+"/config/dnsconfig", 744)
        } else if err != nil {
                log.Println("couldnt find or create /config/dnsconfig")
                return err
        }
	return nil
}

func SetCorefile(domains string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	_, err = os.Stat(dir + "/config/dnsconfig")
	if os.IsNotExist(err) {
		os.Mkdir(dir+"/config/dnsconfig", 744)
	} else if err != nil {
		log.Println("couldnt find or create /config/dnsconfig")
		return err
	}

	corefile := domains + ` {
    reload 15s
    hosts /root/dnsconfig/netmaker.hosts {
	fallthrough	
    }
    forward . 8.8.8.8 8.8.4.4
    log
}
`
	corebytes := []byte(corefile)

	err = ioutil.WriteFile(dir+"/config/dnsconfig/Corefile", corebytes, 0644)
	if err != nil {
		log.Println(err)
		log.Println("")
		return err
	}
	return err
}
