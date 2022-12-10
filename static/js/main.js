(function() {
    const unlockAnimationDurationMs = 3000
    const classUnlocking = 'unlocking'

    let httpRequest;
    let unlockBtn = document.getElementById('unlockBtn')
    unlockBtn.addEventListener('click', makeUnlockRequest);

    function makeUnlockRequest() {
        if (unlockBtn.classList.contains(classUnlocking)) {
            return false;
        }
        unlockBtn.classList.add(classUnlocking);

        httpRequest = new XMLHttpRequest();
        if (!httpRequest) {
            console.warn('Giving up :( Cannot create an XMLHttp instance');
            return false;
        }
        httpRequest.onreadystatechange = processResponse;
        httpRequest.open('GET', '/unlock');
        httpRequest.send();
    }

    function processResponse() {
        if (httpRequest.readyState === XMLHttpRequest.DONE) {
            if (httpRequest.status !== 200) {
                console.warn('There was a problem with the unlocking.');
            } else {
                setTimeout(() => {
                    unlockBtn.classList.remove(classUnlocking);
                }, unlockAnimationDurationMs)
            }
        }
    }
})();
