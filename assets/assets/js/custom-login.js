document.getElementById('loginData').addEventListener('submit', function(event) {
    event.preventDefault(); // Prevent the form from submitting normally

    const password = document.getElementById('password').value;

    // Create the request payload
    const payload = {
        password: password
    };

    // Send the login request to the server
    fetch('/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(payload)
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        return response.json();
    })
    .then(data => {
        // Assuming the server returns a JWT in the 'token' field
        const token = data.token;

        // Store the JWT in localStorage or sessionStorage
        // localStorage.setItem('jwt', token);
        document.cookie = `token=${token}; path=/`;

        // Redirect to the secured page
        window.location.href = '/dashboard';
    })
    .catch(error => {
        console.error('There was a problem with the fetch operation:', error);
        alert('Login failed. Please try again.');
    });
});