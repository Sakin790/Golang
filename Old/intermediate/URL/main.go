package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Learning URL")
	Myurl := "https://www.google.com/search?q=kim+kardashian&sca_esv=c9bd33c212e8a1a7&sxsrf=ADLYWIKDV6Zu5CbxddT7-n8u6sf_aX4gbA%3A1733373699421&source=hp&ei=Ay9RZ6reF4aXvr0P78mdoQU&iflsig=AL9hbdgAAAAAZ1E9E3T8hbHBlDvyWNF5HnzTQk9tpvxo&gs_ssp=eJzj4tLP1TcwMjOssCgxYPTiy87MVchOLEpJLM7ITMwDAHIaCM0&oq=kim+k&gs_lp=Egdnd3Mtd2l6IgVraW0gayoCCAAyCBAuGIAEGLEDMgsQLhiABBixAxiDATIFEAAYgAQyBRAAGIAEMgUQABiABDIFEAAYgAQyBBAAGAMyBRAAGIAEMgUQABiABDIFEAAYgARIrUhQ2yxYojxwAXgAkAEAmAGkAqABmgmqAQUwLjEuNLgBAcgBAPgBAZgCBqAC_wmoAgrCAgcQIxgnGOoCwgIKECMYgAQYJxiKBcICCxAAGIAEGJECGIoFwgILEAAYgAQYsQMYgwHCAg4QLhiABBixAxiDARiKBcICDhAAGIAEGLEDGIMBGIoFwgIIEAAYgAQYsQPCAgQQIxgnwgIQEAAYgAQYsQMYgwEYFBiHAsICCBAuGIAEGNQCwgILEC4YgAQYkQIYigWYAxjxBUZnNexc721-kgcHMS4xLjMuMaAHy0c&sclient=gws-wiz"
	data, err := url.Parse(Myurl)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Type of Url is %T\n", data)
	fmt.Println("Host is", data.Host)
	fmt.Println("Schemas", data.Scheme)
	fmt.Println("Raw Q",data.RawQuery)
	fmt.Println(data.OmitHost)

}
