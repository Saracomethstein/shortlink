document.getElementById('urlForm').addEventListener('submit', async function(event) {
    event.preventDefault();

    const url = document.getElementById('urlInput').value;

    if (!url) {
        alert('Please enter a URL.');
        return;
    }

    try {
        const response = await fetch('http://your-go-server-endpoint/shorten', {
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
        alert('Shortened URL: ' + data.shortenedUrl);
    } catch (error) {
        console.error('There was a problem with the fetch operation:', error);
        alert('There was an error shortening the URL.');
    }
});
