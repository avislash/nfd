# NFD 

NFD is a Golang Client library for accessing the [Non-Fungible Domains (NFD) API](https://editor.swagger.io/?url=https://api.nf.domains/info/openapi3.yaml).

The library currently implements all of the GET methods listed in the [NFD API documentation](https://editor.swagger.io/?url=https://api.nf.domains/info/openapi3.yaml).

# Installation
Standard `go get:`
```
$ go get github.com/avislash/NFD
```


# Usage & Examples
Instantiate a NFD Client object to being using the library

```
package main

import (
    "fmt"
    "github.com/avislash/nfd/client"
)

func main() {
    nfdClient := client.NewClient()
}
```


All of the methods on the NFD Client mirror the Endpoints listed in the NFD API documentation. 

*Example*: The `/nfd/totals/` endpoint maps directly to the `Totals()` method.

Any field marked required in a particular  endpoint is a required argument to the method. Any non-required/optional field listed for an endpoint can be set by employing the `Options` struct defined for that method. 

All implemented methods are listed in `nfd/clients.go` and thier corresponding options are defined in `nfd/api_options.go`

**Example:** 

```
package main

import (
    "fmt"
    "github.com/avislash/nfd/client"
)

func main() {
    nfdClient := client.NewClient()

    //Lookup an address
    addresses := []string{"4F5OA5OQC5TBHMCUDJWGKMUZAQE7BGWCKSJJSJEMJO5PURIFT5RW3VHNZU", 
                           "MQ2QJHZSZ6A7ZXPFE2EPIWLYUMRRDO3DQBEO6NIQ2B5A5OJ4VMWOOI2AX4"
                        }
    records, err := client.Address(addresses, nil)
    if err != nil {
        fmt.Println("Error while looking up addresses: ", err)
        return
    }
    
    for _, record := range records {
        fmt.Printf("%+v\n", record)
    }

    //Lookup an address with options
    options := client.AddressOption{
              View: client.FULL_VIEW,
              Limit: 10,
             
    }
    records, err = client.Address(addresses, &options)
    })
    
    if err != nil {
        fmt.Println("Error while looking up addresses with options: ", err)
        return
    }
    
    for _, record := range records {
        fmt.Printf("%+v\n", record)
    }
    
}
```


# Contributing
Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request


# Contact
Avi Misra - | Email: avi@algowatcher.net | NFD: [avislash.algo](https://avislash.algo.xyz) 


Project Link: [https://github.com/avislash/nfd](https://github.com/avislash/nfd)

# License
See `LICENSE` file for details


# TODOs
TODOs (in no particular order)
- Rate limiter
- POST Methods
- Go Docs


