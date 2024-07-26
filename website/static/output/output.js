function getCookie(name) {
    let cookieArr = document.cookie.split(";");
    for(let i = 0; i < cookieArr.length; i++) {
        let cookiePair = cookieArr[i].split("=");
        if(name == cookiePair[0].trim()) {
            return decodeURIComponent(cookiePair[1]);
        }
    }
    return null;
}

document.addEventListener('DOMContentLoaded', () => {
    const shortenedUrl = getCookie('shortenedUrl');
    const shortUrlField = document.getElementById('shortUrl');

    if (shortenedUrl) {
        shortUrlField.value = shortenedUrl;
    } else {
        shortUrlField.value = 'No URL provided.';
    }

    document.getElementById('copyButton').addEventListener('click', function () {
        shortUrlField.select();
        shortUrlField.setSelectionRange(0, 99999);

        navigator.clipboard.writeText(shortUrlField.value).then(function () {
            console.log('URL copied to clipboard');
        }).catch(function (error) {
            console.error('Failed to copy URL: ', error);
        });
    });
});
