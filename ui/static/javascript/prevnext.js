function Previous() {
    var currentUrl = window.location.href;
    var parts = currentUrl.split('/');
    var lastPart = parseInt(parts[parts.length - 1], 10);

    if (lastPart >= 2) {
        lastPart = lastPart - 1
        parts[parts.length - 1] = lastPart;
        var newUrl = parts.join('/')
        window.history.pushState({}, '', newUrl);
        window.location.href = newUrl;
    }
}

function Next() {
    var currentUrl = window.location.href;
    var parts = currentUrl.split('/');
    var lastPart = parseInt(parts[parts.length - 1], 10);

    if (lastPart <= 51) {
        lastPart = lastPart + 1
        parts[parts.length - 1] = lastPart;
        var newUrl = parts.join('/')
        window.history.pushState({}, '', newUrl);
        window.location.href = newUrl;
    }
}