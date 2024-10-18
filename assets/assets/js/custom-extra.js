function submitChallenge() {
    const input = document.getElementById('challengeInput').value.trim();

    if (!input) {
        alert('Please enter a challenge name.');
        return;
    }

    // Example API endpoint URL
    const apiUrl = '/setChallenge';

    fetch(apiUrl, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ name: input }),
    })
    .then(response => response.json())
    .then(data => {
        console.log('Success:', data);
        // closeModal(); // Close modal on success
    })
    .catch((error) => {
        console.error('Error:', error);
    });
}



function removeChallenge() {
    const input = document.getElementById('challengeInputRemove').value.trim();

    if (!input) {
        alert('Please enter a challenge name.');
        return;
    }

    // Example API endpoint URL
    const apiUrl = '/removeChallenge';

    fetch(apiUrl, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ id: input }),
    })
    .then(response => response.json())
    .then(data => {
        console.log('Success:', data);
        // closeModal(); // Close modal on success
    })
    .catch((error) => {
        console.error('Error:', error);
    });
}

function removeYTLink() {
    const input = document.getElementById('YTInputRemove').value.trim();

    if (!input) {
        alert('Please enter a valid ID.');
        return;
    }

    // Example API endpoint URL
    const apiUrl = '/removeYTLink';

    fetch(apiUrl, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ id: input }),
    })
    .then(response => response.json())
    .then(data => {
        console.log('Success:', data);
        // closeModal(); // Close modal on success
    })
    .catch((error) => {
        console.error('Error:', error);
    });
}



function submitYTLink() {
    const input = document.getElementById('ytTitleInput').value.trim();
    const input2 = document.getElementById('ytLinkInput').value.trim();


    if (!input) {
        alert('Please enter a challenge name.');
        return;
    }

    // Example API endpoint URL
    const apiUrl = '/setYTLink';

    fetch(apiUrl, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ 
            title: input,
            link: input2,
         }),
    })
    .then(response => response.json())
    .then(data => {
        console.log('Success:', data);
        // closeModal(); // Close modal on success
    })
    .catch((error) => {
        console.error('Error:', error);
    });
}



