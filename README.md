# go-jenkins-sdk

go-jenkins-sdk is a SDK interact with your Jenkins server using API.

## Installation

Use [go module](https://go.dev/blog/using-go-modules) to install jenkins-sdk

```
go get github.com/wizk3y/go-jenkins-sdk
```

## Usage

```
package main

import (
    "github.com/wizk3y/go-jenkins-sdk"
    restclient "github.com/wizk3y/go-jenkins-sdk/rest"
)

func main() {
    config := restclient.NewConfig("", username, password)

    c, err := jenkins.NewForConfig(config)
    if err != nil {
        log.Println("error when create client", err)
        return
    }

    computers, err := c.Computer().GetComputers()
}
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)