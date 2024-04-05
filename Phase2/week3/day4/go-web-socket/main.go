/*
Reference : https://dasarpemrogramangolang.novalagung.com/D-golang-web-socket-chatting-app.html
*/
package main

import (
	"log"      // Provides a simple logging package.
	"net/http" // Provides HTTP client and server implementations.

	"github.com/gorilla/websocket" // Gorilla WebSocket package for handling websockets.
	"github.com/labstack/echo/v4"  // Echo, a high-performance web framework for Go.
)

// Global variable `upgrader` of type `websocket.Upgrader`. This is used to upgrade HTTP server connections to the WebSocket protocol.
var upgrader = websocket.Upgrader{}

// The entry point of the program.
func main() {
	e := echo.New() // Create a new instance of Echo.

	e.File("/", "public/index.html")     // Serve `index.html` file at the root URL path.
	e.Static("/static", "public/static") // Serve static files from the `public/static` directory under the `/static` URL path.

	// Define a route for WebSocket connections on `/ws`. When a request comes to `/ws`, it will be handled by `websocketHandler`.
	e.GET("/ws", func(c echo.Context) error {
		websocketHandler(c.Response().Writer, c.Request())
		return nil
	})

	// Start the server on port 8080 and log if there's an error.
	e.Logger.Fatal(e.Start(":8081"))
}

// A function to handle WebSocket requests.
func websocketHandler(w http.ResponseWriter, r *http.Request) {
	// Allow all origins for WebSocket connections (not recommended for production).
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// Attempt to upgrade the HTTP connection to a WebSocket connection.
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err) // Log if there was an error upgrading the connection.
		return
	}
	defer conn.Close() // Ensure the connection is closed when the function returns.

	for {
		// Continuously read messages from the WebSocket connection.
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err) // Log if there was an error reading the message.
			break
		}
		log.Printf("recv: %s", message) // Log the received message.

		// Echo the received message back to the client.
		err = conn.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err) // Log if there was an error writing the message.
			break
		}
	}
}
