const wsUrl = 'ws://localhost:3456'

const socket = new WebSocket(wsUrl);

socket.onopen = () => {
    console.log('Established connection');

    table.authorize();

    table.handleInput()
};

socket.onclose = (event) => {
    const details = `(Code: ${event.code})`

    if (event.wasClean) {
        console.log('Closed connection', details);
    } else {
        console.log('Aborted connection', details); // happens on server initiative
    }
};

socket.onmessage = (event) => {
    console.log('Event data', event.data);

    const ev = JSON.parse(event.data);

    // handle event.
    switch (ev.Op) {
        case 'authorized': // Current user authorized.
            table.showAuthorized(ev.Data);
            break;
        case 'user/authorized': // Other user authorized.
            break;
        case 'cell/edited': // Current user edited cell.
            break;
        case 'user/cell/edited': // Other user edited cell.
            table.fillCell(ev.Data);
            break;
    }
};

socket.onerror = (error) => {
    console.log('Error', error.message);
};

const table = {

    // Connect to table and authorize client.
    authorize: () => {
        const msg = {Op: 'authorize', Data: {}}
        socket.send(JSON.stringify(msg));
    },

    // Fill cell.
    fillCell: (data) => {
        const input = document.querySelector(`.cell input[name="${data.name}"]`);

        input.value = data.value;
    },

    // Edit cell.
    editCell: (event) => {
        const msg = {
            Op: 'cell/edit', Data: {
                name: event.target.getAttribute('name'), value: event.target.value,
            }
        }
        socket.send(JSON.stringify(msg));
    },

    // Handle input elements.
    handleInput: () => {
        const inputs = document.querySelectorAll('.cell input');

        inputs.forEach(function (input) {
            input.addEventListener("focusout", table.editCell);
        });
    },

    // Show authorized connection.
    showAuthorized: (data) => {
        const elem = document.querySelector('#connection-id');

        elem.textContent = `Current connection: ${data?.id}`
    },
}
