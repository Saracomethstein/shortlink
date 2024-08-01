document.addEventListener('DOMContentLoaded', function () {
    fetch('http://localhost:8000/profile', {
        method: 'GET',
        credentials: 'include'
    })
        .then(response => response.json())
        .then(data => {
            const usernameElement = document.getElementById('username');
            usernameElement.textContent = data.username;

            const tableBody = document.getElementById('urlHistoryTable').getElementsByTagName('tbody')[0];
            tableBody.innerHTML = '';

            data.urlHistory.forEach(url => {
                const row = document.createElement('tr');

                const shortUrlCell = document.createElement('td');
                shortUrlCell.textContent = url.shortenedUrl;
                row.appendChild(shortUrlCell);

                const originalUrlCell = document.createElement('td');
                originalUrlCell.textContent = url.url;
                row.appendChild(originalUrlCell);

                tableBody.appendChild(row);
            });
        })
        .catch(error => console.error('Error fetching profile data:', error));
});