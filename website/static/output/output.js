function getQueryParam(param) {
    const urlParams = new URLSearchParams(window.location.search);
    console.log(urlParams)
    console.log(urlParams.get(param))
    return urlParams.get(param);
}

document.addEventListener('DOMContentLoaded', () => {
    const shortenedUrl = getQueryParam('shortenedUrl');
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
