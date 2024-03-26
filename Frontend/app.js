document.addEventListener("DOMContentLoaded", function() {
    var payButton = document.getElementById('pay-button');
    payButton.addEventListener('click', function(e) {
        e.preventDefault();
        
        var amount = document.getElementById('amount').value;
        var email = document.getElementById('email').value;

        // Validate amount and email
        if (!amount || !email) {
            alert('Please enter the amount and your email address.');
            return;
        }

        // Show loading indicator
        document.getElementById("loading").style.display = "block";

        // Initialize Paystack inline
        var handler = PaystackPop.setup({
            key: 'pk_test_4c40126d35d265338ccd282209c39a306360642c', // Replace with your test public key
            email: email, // Use the email entered by the user
            amount: amount * 100, // Amount in kobo
            currency: 'NGN', // Currency code
            ref: 'TEST' + Math.floor((Math.random() * 1000000000) + 1), // Generate random reference
            metadata: {
                custom_fields: [
                    {
                        display_name: "Mobile Number",
                        variable_name: "mobile_number",
                        value: "+2348012345678" // Replace with customer's phone number
                    }
                ]
            },
            callback: function(response) {
                // Handle the response here
                try {
                    // Handle the response here
                    document.getElementById("result").style.display = "block";
                    document.getElementById("message").textContent =
                        "Payment processed successfully. Reference: " +
                        response.reference;
                } catch (error) {
                    document.getElementById("result").style.display = "block";
                    document.getElementById("message").textContent =
                        "Error processing payment: " + error.message;
                } finally {
                    // Hide loading indicator
                    document.getElementById("loading").style.display = "none";
                }
            },
            onClose: function() {
                // Handle payment modal close
                alert('Payment closed.');
            }
        });

        handler.openIframe(); // Open Paystack inline payment modal
    });
});