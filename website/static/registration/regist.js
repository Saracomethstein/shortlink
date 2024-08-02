document.getElementById('registrForm').addEventListener('submit', async function(event) {
    event.preventDefault();

    const login = document.getElementById('login').value;
    const password = document.getElementById('password').value;

    try {
        const response = await fetch('/api/registration', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ login, password })
        });

        const data = await response.json();

        if (response.ok) {
            document.getElementById('message').textContent = 'Registration successful!';
            window.location.href = '/auth/';
        } else {
            document.getElementById('message').textContent = 'Registration failed: ' + data.error;
        }
    } catch (error) {
        console.error('Error:', error);
        document.getElementById('message').textContent = 'An error occurred. Please try again.';
    }
});
