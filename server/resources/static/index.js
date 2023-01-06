const wsUrl = 'ws://localhost:3456'

const socket = new WebSocket(wsUrl);

socket.onopen = () => {
    console.log('Established connection');

    table.authorize()
};

socket.onclose = (event) => {
    const details = `(Code: ${event.code} Reason: ${event.reason})`

    if (event.wasClean) {
        console.log('Closed connection', details);
    } else {
        console.log('Aborted connection', details); // happens on server initiative
    }
};

socket.onmessage = (event) => {
    console.log('Event data', event.data);
};

socket.onerror = (error) => {
    console.log('Error', error.message);
};

const table = {

    // Connect to table and authorize client.
    authorize: () => {
        let msg = {Op: 'authorize', Data: {"Hello": "World"}}
        socket.send(JSON.stringify(msg));
    }
}
