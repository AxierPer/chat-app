package main

import (
  "fmt"
  "log"
  "net/http"
  "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
  CheckOrigin: func (r *http.Request) bool{
    return true
  },
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan string)

func handleConnections(w http.ResponseWriter, r *http.Request)  {
  conn, err := upgrader.Upgrade(w, r, nil)
  
  if err != nil {
    log.Println(err)
    return
  }

  defer conn.Close()

  clients[conn] = true
  fmt.Println("New conection")

  for {
    messageType, p, err := conn.ReadMessage()

    if err != nil{
      log.Println(err)
      delete(clients, conn)
      break
    }

    for client := range clients {
      if err := client.WriteMessage(messageType, p); err != nil {
        log.Println(err)
        client.Close()
        delete(clients, client)
      }
    }
  }
}

func main() {
  http.HandleFunc("/ws", handleConnections)

  go func ()  {
    for {
      select {
      case message := <- broadcast:
        for client := range clients{
          if err := client.WriteMessage(websocket.TextMessage, []byte(message)); err != nil{
            log.Println(err)
            client.Close()
            delete(clients, client)
          }
        }
      }
    }
  }()

  http.Handle("/", http.FileServer(http.Dir("./static")))
  fmt.Println("Run server in http://localhost:8080")
  log.Fatal(http.ListenAndServe(":8080", nil))
}
