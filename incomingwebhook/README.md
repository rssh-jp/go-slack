# incoming webhook

# usage
``` main.go
package main

import(
    "log"

    "github.com/rssh-jp/go-slack/incomingwebhook"
)

func main(){
    s := incomingwebhook.New("your-webhook-url")
    err := s.SendSimple("simple text")
    if err != nil{
        log.Fatal(err)
    }
}
```
