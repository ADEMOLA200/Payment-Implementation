document.getElementById("paymentForm").addEventListener("submit", function(event) {
    event.preventDefault();

    var formData = new FormData(this);
    var jsonObject = {};
    formData.forEach(function(value, key) {
        // Convert the amount field to a number
        if (key === "amount") {
            value = parseFloat(value); // Parse string to float
        }
        jsonObject[key] = value;
    });

    fetch("http://localhost:3400/api/makePayment", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(jsonObject)
    })
    .then(response => response.text())
    .then(data => {
        document.getElementById("message").textContent = data;
    })
    .catch(error => {
        console.error("Error:", error);
    });
});
