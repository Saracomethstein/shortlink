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

            const domainLabels = Object.keys(data.domains);
            const domainCounts = Object.values(data.domains);

            const backgroundColors = domainLabels.map(() => `rgba(${Math.floor(Math.random() * 256)}, ${Math.floor(Math.random() * 256)}, ${Math.floor(Math.random() * 256)}, 0.2)`);
            const borderColors = domainLabels.map(() => `rgba(${Math.floor(Math.random() * 256)}, ${Math.floor(Math.random() * 256)}, ${Math.floor(Math.random() * 256)}, 1)`);

            const ctx = document.getElementById('domainChart').getContext('2d');
            new Chart(ctx, {
                type: 'bar',
                data: {
                    labels: domainLabels,
                    datasets: [{
                        label: 'Top 10 Most Used Domains',
                        data: domainCounts,
                        backgroundColor: backgroundColors,
                        borderColor: borderColors,
                        borderWidth: 1
                    }]
                },
                options: {
                    scales: {
                        y: {
                            beginAtZero: true
                        }
                    }
                }
            });
        })
        .catch(error => console.error('Error fetching profile data:', error));
});