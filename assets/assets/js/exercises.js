let imageFile = null;

// Function to open the modal
function openModal() {
    document.getElementById("modal").classList.remove("hidden");
}

// Function to handle image upload and show preview
function handleImageUpload(event) {
    const file = event.target.files[0];
    if (file) {
        imageFile = file;
        const reader = new FileReader();
        reader.onload = function(e) {
            const previewImage = document.getElementById("previewImage");
            previewImage.src = e.target.result;
            document.getElementById("imgPreview").classList.remove("hidden");
        };
        reader.readAsDataURL(file);
    }
}

// Function to submit form data
async function submitForm() {
    const name = document.getElementById("name").value;
    const currentWeight = document.getElementById("currentWeight").value;
    const bWeight = document.getElementById("bWeight").value;
    const iWeight = document.getElementById("iWeight").value;
    const aWeight = document.getElementById("aWeight").value;

    if (!name || !currentWeight || !bWeight || !iWeight || !aWeight || !imageFile) {
        alert("All fields must be filled in.");
        return;
    }

    const formData = new FormData();
    formData.append("name", name);
    formData.append("current_weight", currentWeight);
    formData.append("b_weight", bWeight);
    formData.append("i_weight", iWeight);
    formData.append("a_weight", aWeight);
    formData.append("img", imageFile);

    try {
        const response = await fetch('/addExercise', {
            method: 'POST',
            body: formData
        });

        if (response.ok) {
            alert("Data successfully uploaded!");
            closeModal();
        } else {
            alert("Failed to upload data.");
        }
    } catch (error) {
        console.error("Error uploading data:", error);
        alert("Error uploading data.");
    }
}

// Function to cancel the action and close the modal
function cancel() {
    closeModal();
}

// Function to close the modal
function closeModal() {
    document.getElementById("modal").classList.add("hidden");
    location.reload()

}

// Add event listener to open modal on click
document.getElementById("openModal").addEventListener("click", openModal);








function removeExercise() {
    const input = document.getElementById('textboxExerciseRemove').value.trim();
    if (!input) {
        alert('Please enter a valid name.');
        return;
    }

    // Example API endpoint URL
    const apiUrl = '/removeExercise';

    fetch(apiUrl, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ id: input }),
    })
    .then(response => response.json())
    .then(data => {
        console.log('Success:', data);
        document.getElementById('textbox').value = 'Refresh Page';    })
        location.reload()


    .catch((error) => {
        console.error('Error:', error);
    });

}



let selectedExercise = ""; // Variable to store the selected exercise name

function openModall(exerciseName) {
    selectedExercise = exerciseName; // Store the clicked exercise name

    document.getElementById('weightModal').classList.remove('hidden');
}

function closeModall() {
    document.getElementById('weightModal').classList.add('hidden');
}

function updateWeight() {
    const newWeight = document.getElementById('weightInput').value;
    if (!newWeight) {
        alert('Please enter a valid number.');
        return;
    }
    const apiUrl = '/updateExercise';

    const payload = {
        name: selectedExercise,
        weight: newWeight
    };

    fetch(apiUrl, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(payload),
    })
    .then(response => response.json())
    .then(data => {
        console.log('Success:', data);
        document.getElementById('textbox').value = 'Refresh Page';})
        location.reload()


    .catch((error) => {
        console.error('Error:', error);
    });
    
    console.log("New Weight:", newWeight);
    closeModal(); // Close the modal after updating
}