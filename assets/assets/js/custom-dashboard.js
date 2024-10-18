let timer; // Variable to store the timer
let startTime; // Variable to store the start time
let elapsedTime = 0; // Variable to store the elapsed time
let milliseconds = 0;
let seconds = 0;
let minutes = 0;
let sound1mEnabled = false;
let lastSoundTime = 0;
let currentTime = 0

    function startTimer() {
        startTime = performance.now() - elapsedTime;
        timer = setInterval(() => {
            const currentTime = performance.now();
            elapsedTime = currentTime - startTime;
            updateDisplay();
            checkPlaySound();
        }, 10);
    }

    function stopTimer() {
        clearInterval(timer);
    }

    function resetTimer() {
        clearInterval(timer);
        elapsedTime = 0;
        updateDisplay();
        lastSoundTime = 0
    }

    function updateDisplay() {
        milliseconds = Math.floor((elapsedTime % 1000)/10);
        seconds = Math.floor((elapsedTime / 1000) % 60);
        minutes = Math.floor((elapsedTime / (1000 * 60)) % 60);

        const display = document.getElementById('display');
        display.textContent = `${minutes.toString().padStart(2, '0')}:${seconds.toString().padStart(2, '0')}:${milliseconds.toString().padStart(2, '0')}`;
    }

    function checkPlaySound() {
        // if (sound1mEnabled && minutes === 1 && seconds === 0 && milliseconds === 0) {
        //     playSound();
        // }
        currentTime = elapsedTime;
        if (sound1mEnabled && currentTime - lastSoundTime >= 60000) {
            playSound()
            lastSoundTime = currentTime;
        }
        
    }
    function playSound() {
        const audio = new Audio('/assets/assets/sound/LvUp.mp3');
        audio.play();
    }

    function toggleSound1m() {
        sound1mEnabled = !sound1mEnabled;
        const button = document.getElementById('toggleSound1m');
        currentTime = 0
        button.textContent = sound1mEnabled ? 'Sound at 1m: On' : 'Sound at 1m: Off';
    }


    
document.getElementById('start').addEventListener('click', startTimer);
document.getElementById('stop').addEventListener('click', stopTimer);
document.getElementById('reset').addEventListener('click', resetTimer);
document.getElementById('toggleSound1m').addEventListener('click', toggleSound1m);




// For api weights

//For kg display------------------------------------
document.addEventListener('DOMContentLoaded', function() {
    fetch("/getAllBodyWeight")
        .then(response => response.json())
        .then(data => {
            const newestEntry = findNewestEntry(data);
            displayWeight(newestEntry.Weight);

            const ytdEntries = filterYTDEntries(data);
            const ytdDifference = calculateYTDDifference(ytdEntries);
            displayYTDDifference(ytdDifference);


        })
        .catch(error => console.error('Error fetching data:', error));
});

function findNewestEntry(entries) {
    return entries.reduce((newest, current) => {
        const newestDate = new Date(newest.Date);
        const currentDate = new Date(current.Date);
        return currentDate > newestDate ? current : newest;
    });
}

function displayWeight(weight) {
    const weightDisplay = document.getElementById('weightNumber');
    weightDisplay.textContent = weight+" KG";
}

//For ytd % difference ------------------------------------
function filterYTDEntries(entries) {
    const now = new Date();
    const currentYear = now.getFullYear();
    return entries.filter(entry => {
        const entryDate = new Date(entry.Date);
        return entryDate.getFullYear() === currentYear && entryDate <= now;
    });
}

function calculateYTDDifference(entries) {
    if (entries.length < 2) return 0; // Need at least two entries to calculate difference

    const sortedEntries = entries.sort((a, b) => new Date(a.Date) - new Date(b.Date));
    const firstWeight = parseFloat(sortedEntries[0].Weight);
    const lastWeight = parseFloat(sortedEntries[sortedEntries.length - 1].Weight);

    const difference = lastWeight - firstWeight;
    const percentageDifference = (difference / firstWeight) * 100;

    return percentageDifference.toFixed(2); // Round to two decimal places
}

function displayYTDDifference(difference) {
    const ytdDifferenceDisplay = document.getElementById('ytdDifferenceDisplay');
    ytdDifferenceDisplay.textContent = "YTD Change: "+difference+"%";
}



//------------------------------------------------------------------------

function submitWeight() {
    const input = document.getElementById('weightInput').value.trim();

    if (!input) {
        alert('Please enter a weight.');
        return;
    }

    // Example API endpoint URL
    const apiUrl = '/addBodyWeight';

    fetch(apiUrl, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ we: input }),
    })
    .then(response => response.json())
    .then(data => {
        console.log('Success:', data);
        // closeModal(); // Close modal on success
    })
    .catch((error) => {
        console.error('Error:', error);
    });
}




