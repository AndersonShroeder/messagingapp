var socket = new WebSocket("ws://localhost:8080/ws")

let connect = () => {
    console.log("Connecting...");

    socket.onopen = () => {
        console.log("Connected!");
    };

    socket.onmessage = msg => {
        console.log(msg);
    };

    socket.onclose = event => {
        console.log("Connection Closed: ", event);
    };

    socket.onerror = error => {
        console.log("Error: ", error);
    };
};

let sendMsg = msg => {
    console.log("sending message: ", msg);
    socket.send(msg);
};

export {connect, sendMsg}