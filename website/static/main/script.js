document.addEventListener('DOMContentLoaded', function() {
    const sessionId = getCookie('session_id');
    const authButton = document.querySelector('.auth-button');

    if (sessionId) {
        authButton.textContent = 'Logout';
        authButton.onclick = function() {
            document.cookie = 'session_id=; path=/; expires=Thu, 01 Jan 1970 00:00:00 UTC';
            window.location.href = '/';
        };
    } else {
        authButton.textContent = 'Login';
        authButton.onclick = function() {
            window.location.href = 'http://localhost:8000/';
        };
    }
});

function getCookie(name) {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    if (parts.length === 2) return parts.pop().split(';').shift();
}

document.getElementById('urlForm').addEventListener('submit', async function(event) {
    event.preventDefault();

    const url = document.getElementById('urlInput').value;

    if (!url) {
        alert('Please enter a URL.');
        return;
    }

    try {
        const response = await fetch('http://localhost:8000/shorten', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ url })
        });

        if (!response.ok) {
            throw new Error('Network response was not ok');
        }

        const data = await response.json();
        document.cookie = `shortenedUrl=${encodeURIComponent(data.shortenedUrl)}; path=/`;
        window.location.href = 'http://localhost:8000/output';
    } catch (error) {
        console.error('There was a problem with the fetch operation:', error);
        alert('There was an error shortening the URL.');
    }
});