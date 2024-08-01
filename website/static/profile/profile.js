document.addEventListener('DOMContentLoaded', function () {
    fetch('http://localhost:8000/profile', {
        method: 'GET',
        credentials: 'include'
    })
        .then(response => response.json())
        .then(data => {
            const tableBody = document.getElementById('urlHistoryTable').getElementsByTagName('tbody')[0];
            tableBody.innerHTML = '';

            data.forEach(url => {
                const row = document.createElement('tr');

                const shortUrlCell = document.createElement('td');
                shortUrlCell.textContent = url.shortenedUrl;
                row.appendChild(shortUrlCell);

                const originalUrlCell = document.createElement('td');
                originalUrlCell.textContent = url.url;
                row.appendChild(originalUrlCell);

                console.log(url.url)
                console.log(url.shortenedUrl)

                tableBody.appendChild(row);
            });
        })
        .catch(error => console.error('Error fetching URL history:', error));
});
