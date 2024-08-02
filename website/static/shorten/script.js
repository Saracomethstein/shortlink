document.addEventListener('DOMContentLoaded', function() {
    const sessionId = getCookie('session_id');
    const authButton = document.querySelector('.auth-button');
    const profileButton = document.querySelector('.profile-button')
    const submitButton = document.querySelector('.submit-button')

    if (sessionId) {
        authButton.textContent = 'Logout';
        authButton.onclick = function() {
            document.cookie = 'session_id=; path=/; expires=Thu, 01 Jan 1970 00:00:00 UTC';
            window.location.href = '/auth/';
        };

        profileButton.onclick = function () {
            window.location.href='/profile/';
        };

        submitButton.onclick = function () {
            window.location.href='/output/'
        };
    } else {
        authButton.textContent = 'Login';
        authButton.onclick = function() {
            window.location.href = '/auth/';
        };

        profileButton.onclick = function () {
            window.location.href = '/auth/';
        };

        submitButton.onclick = function () {
            window.location.href='/auth/'
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
    const sessionId = getCookie('session_id');
    const url = document.getElementById('urlInput').value;

    if (!sessionId) {
        alert('Please authorization.')
        return
    }

    if (!url) {
        alert('Please enter a URL.');
        return;
    }

    try {
        const response = await fetch('/api/shorten', {
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
        console.log(data.shortenedUrl)
        window.location.href = '/output/';
    } catch (error) {
        console.error('There was a problem with the fetch operation:', error);
        alert('There was an error shortening the URL.');
    }
});
