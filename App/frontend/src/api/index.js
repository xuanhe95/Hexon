var socket = new WebSocket('ws://localhost:8080/ws');

let connect = cb => {
    console.log("Attempting Connection...");
    socket.onopen = () => {
        console.log("Successfully Connected");
        sendMsg("Hello From The Client!");
    };

    socket.onmessage = msg => {
        console.log("Received message: " + msg.data);
        cb(msg);
    };

    socket.onclose = event => {
        console.log("Socket Closed Connection: ", event);
    };

    socket.onerror = error => {
        console.log("Socket Error: ", error);
    };
};

let sendMsg = msg => {
    console.log("Sending message: ", msg);
    socket.send(msg);
};

export { connect, sendMsg };