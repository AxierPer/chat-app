# WebSocket Echo Server 🎉

This is a simple WebSocket server built with Go that allows clients to connect and send messages to each other. It is an ideal project to practice implementing WebSockets in Go using the `gorilla/websocket` library.

## 🚀 Description 

This WebSocket server allows users to connect and send text messages. When a client sends a message, the server forwards it to all other connected clients (“echo” function). 

## 📦 Requirements

- Go 1.18 or higher.
- Dependency: [gorilla/websocket](https://github.com/gorilla/websocket)

## 🛠️ Installation

1. Clone this repository:

   ```bash
   git clone https://github.com/tu_usuario/mi-websocket-app.git
   ```
2. Enter the project directory:

    ```bash
    cd mi-websocket-app
    ```
3. Install dependencies:

    ```bash
    go get github.com/gorilla/websocket
    ```
## 🎮 How Use
1. Runs the WebSocket server:

    ```bash
    go run main.go
    ```
2. The server should be running on `http://localhost:8080`.

3. To test WebSocket, open a browser and access the URL `http://localhost:8080`. You will be able to connect multiple clients and see how messages sent by one client are forwarded to all other clients.

## ⚙️ Operation
This WebSocket server makes use of the `gorilla/websocket` library to manage WebSocket connections and send messages between clients. Here is a breakdown:
###  👩‍💻 WebSocket Handler
The server listens on the `/ws` path for incoming WebSocket connections. Once a client connects, it is mapped to a `clients` map and the connection is kept open. Each time a client sends a message, the server forwards it to all other connected clients.

### 🖥️ Code Components
- `handleConnections`: This function handles WebSocket connections. When it receives a message from a client, it forwards it to all connected clients.

- `broadcast`: Channel where messages are sent to be forwarded to all clients. If a message is sent on this channel, all clients will receive the message.

- `main`: The entry point of the application that configures the server and starts the WebSocket server. A goroutine is also started here to listen for messages to be sent to all clients.

## 📨 How it works
1. WebSocket Connection: The client connects to the server on the /ws path.
2. Message Reading: The server reads the messages sent by the clients and relays them to all other connected clients.
3. Disconnection: If a client disconnects or if there is an error, the connection is closed and removed from the client map.

# 🌍 Project Structure
```go
my-websocket-app/
│
├─── main.go // Main file to start the WebSocket server.
├─── go.mod // Go module file to manage dependencies.
├─── static/
│ └── index.html // Simple frontend to interact with WebSocket server.
```

# 💬 Testing
To test the server, simply open multiple browser tabs or use tools such as Postman or wscat to connect to the WebSocket server.

1. Open multiple browser tabs: Open multiple browser tabs at `http://localhost:8080` and you will see how messages are propagated between all connected tabs.

2. Using wscat: Install `wscat` with `npm install -g wscat` and connect using:

```bash
wscat -c ws://localhost:8080/ws
```
# 🤝 Contribute
Contributions are welcome. If you have an idea or improvements for this project, please open a pull request.

# 📜 License
This project is licensed under the MIT License. See the LICENSE file for details.