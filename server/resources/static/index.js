const wsUrl = 'ws://localhost:3456'

const socket = new WebSocket(wsUrl);

socket.onopen = () => {
    console.log('Established connection');

    table.authorize();

    table.handleInput();
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
            connection.showAuthorized(ev.Data);
            table.loadCells();
            log.pushEvent(`Current user "${ev.Data?.id}" authorized`)
            break;
        case 'user/authorized': // Other user authorized.
            log.pushEvent(`User "${ev.Data?.id}" authorized`)
            break;
        case 'user/disconnected': // Other user disconnected.
            log.pushEvent(`User "${ev.Data?.id}" disconnected`)
            break;
        case 'cell/saved': // Current user saved cell.
            log.pushEvent(`Current user saved cell "${ev.Data?.name}"`)
            break;
        case 'user/cell/saved': // Other user saved cell.
            table.fillCell(ev.Data);
            log.pushEvent(`User "${ev.Data?.user_id}" saved cell "${ev.Data?.name}"`)
            break;
        case 'cell/loading': // Current user loading cell.
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

    // Loading all cells.
    loadCells: () => {
        const msg = {
            Op: 'cell/load/all', Data: {}
        }
        socket.send(JSON.stringify(msg));
    },

    // Fill cell.
    fillCell: (data) => {
        const input = document.querySelector(`.cell input[name="${data.name}"]`);

        input.value = data.value;
    },

    // Save cell.
    saveCell: (event) => {
        const msg = {
            Op: 'cell/save', Data: {
                name: event.target.getAttribute('name'), value: event.target.value,
            }
        }
        socket.send(JSON.stringify(msg));
    },

    // Handle input elements.
    handleInput: () => {
        const inputs = document.querySelectorAll('.cell input');

        inputs.forEach(function (input) {
            input.addEventListener("focusout", table.saveCell);
        });
    },
}

const connection = {

    // Show authorized connection.
    showAuthorized: (data) => {
        const elem = document.querySelector('#connection-id');

        elem.textContent = `Current user: ${data?.id}`
    },
}

const log = {

    // Push event log.
    pushEvent: (log) => {
        const list = document.querySelector('#event-log');

        const item = document.createElement("li")
        item.textContent = log

        list.append(item)
    },
}
