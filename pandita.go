package pandita

import (
  "fmt"
  "net"
  "os"
)

func main() {
	fmt.Println("A ver young and tiiired little panda")
}




func Listen(port string) {
  l, err := net.Listen("tcp", "0.0.0.0:" + port)
  if err != nil {
    fmt.Println("Failed to bind to port " + port)
    os.Exit(1)
  }
  for {
    _, err := l.Accept()
    if err != nil {
      fmt.Println("Failed to accept incoming connection")
    }
  }
}
