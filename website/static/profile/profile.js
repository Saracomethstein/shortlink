document.addEventListener('DOMContentLoaded', async function() {
    const usernameElement = document.getElementById('username');
    const urlHistoryElement = document.getElementById('urlHistory');
    const domainChartElement = document.getElementById('domainChart').getContext('2d');

    try {
        const profileResponse = await fetch('http://localhost:8000/profile', {
            method: 'GET',
            credentials: 'include'
        });

        if (!profileResponse.ok) {
            throw new Error('Failed to fetch profile data');
        }

        const profileData = await profileResponse.json();
        usernameElement.textContent = profileData.username;

        profileData.urlHistory.forEach(url => {
            const listItem = document.createElement('li');
            listItem.textContent = `Original: ${url.originalUrl} - Shortened: ${url.shortenedUrl}`;
            urlHistoryElement.appendChild(listItem);
        });

        const domainLabels = profileData.topDomains.map(domain => domain.name);
        const domainData = profileData.topDomains.map(domain => domain.count);

        new Chart(domainChartElement, {
            type: 'bar',
            data: {
                labels: domainLabels,
                datasets: [{
                    label: 'Top 10 Most Used Domains',
                    data: domainData,
                    backgroundColor: 'rgba(75, 192, 192, 0.2)',
                    borderColor: 'rgba(75, 192, 192, 1)',
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
    } catch (error) {
        console.error('Error loading profile data:', error);
    }
});