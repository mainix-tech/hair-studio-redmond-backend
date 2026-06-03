package main

import (
	"encoding/json"
	"net/http"
)

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	/*
		Even without WebSockets, a standard HTTP request/response still uses a persistent TCP network connection while the data is actively being sent.
		HTTP relies on TCP "Pipes"
		HTTP doesn't just float through the air as a single completed package. Under the hood, every HTTP request opens a raw TCP network connection (a digital pipe) between the client and your server.


		2. Streaming vs. Buffering
			When you send a response, your server has to transmit those bytes across that network pipe.
			The Buffering Way (json.Marshal): Your server converts the entire data structure into a full JSON string in its RAM first, then drops that massive block into the pipe all at once.
			The Streaming Way (json.NewEncoder): Your server reads the data structure and sends it down the pipe piece-by-piece (byte-by-byte).


		Imagine you are sending a large JSON list of 1,000 drivers.
		While json.NewEncoder is busy converting and pushing those drivers down the network pipe one-by-one, the user suddenly enters a subway tunnel and loses reception.
		The TCP pipe instantly breaks.
		Because json.NewEncoder is actively streaming, Go realizes on that exact byte that the pipe is broken.
		It immediately stops wasting CPU and memory trying to format the rest of the 900 drivers, cancels the operation, and returns an error to your code.
	*/

	// We use "Encode" to handle case if something happen during response to the user, so we can handle it.
	//Instead of converting the data to JSON in memory first and then writing it to the network, .Encode(data) converts and pushes the data piece by piece straight into the network pipe.
	//If the user's connection snaps on byte 10, Go immediately stops processing and gives you an error, saving your server memory and CPU time.
	return json.NewEncoder(w).Encode(data)
}
