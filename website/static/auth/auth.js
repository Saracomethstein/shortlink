document.getElementById('authForm').addEventListener('submit', async function(event) {
    event.preventDefault();

    const login = document.getElementById('login').value;
    const password = document.getElementById('password').value;

    if (!login || !password) {
        alert('Please enter both login and password.');
        return;
    }

    try {
        const response = await fetch('http://localhost:8000/auth', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ login, password })
        });

        if (!response.ok) {
            throw new Error('Network response was not ok');
        }

        const data = await response.json();
        if (data.success) {
            document.cookie = `session_id=${data.session_id}; path=/`;
            window.location.href = 'http://localhost:8000/shorten';
        } else {
            alert('Authentication failed.');
        }
    } catch (error) {
        console.error('There was a problem with the fetch operation:', error);
        alert('There was an error with authentication.');
    }
});