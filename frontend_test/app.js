document.getElementById("payment-form").addEventListener("submit", async (e) => {
    e.preventDefault();
  
    document.getElementById("loading").style.display = "block";
  
    const formData = new FormData(e.target);
    const payload = {
      amount: parseFloat(formData.get("amount")),
      email: formData.get("email"),
    };
  
    try {
      const response = await fetch("http://localhost:3400/api/makePayment", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(payload),
      });
  
      if (!response.ok) {
        throw new Error("Payment processing failed");
      }
      
  
      const resultData = await response.json();
      document.getElementById("result").style.display = "block";
      document.getElementById("message").textContent =
        "Payment processed successfully. Reference: " +
        resultData.reference;
    } catch (error) {
      document.getElementById("result").style.display = "block";
      document.getElementById("message").textContent =
        "Error processing payment: " + error.message;
    }
  
    document.getElementById("loading").style.display = "none";
  });