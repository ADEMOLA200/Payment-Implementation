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
                alert('Payment successful. Reference: ' + response.reference);
            },
            onClose: function() {
                alert('Payment closed.');
            }
        });

        handler.openIframe(); // Open Paystack inline payment modal
    });
});
