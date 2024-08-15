document.addEventListener('DOMContentLoaded', function() {
  // Define the API endpoint and parameters
  
  const city = 'Delhi';
  const start = '2024-08-15T22:55:38Z';
  const end = '2024-08-15T23:22:20Z';
  const apiUrl = `http://localhost:8080/api/v1/weather/chart?city=${city}&start=${start}&end=${end}`;
  
  // Fetch the chart image
  fetch(apiUrl)
      .then(response => {
        console.log(response);
          if (!response.ok) {
              throw new Error('Network response was not ok');
          }
          return response.blob();
      })
      .then(blob => {
          // Create a URL for the chart image
          const imageUrl = URL.createObjectURL(blob);
          
          // Set the src of the img element to the chart image URL
          const imgElement = document.getElementById('weatherChart');
          imgElement.src = imageUrl;
      })
      .catch(error => {
          console.error('Error fetching the chart:', error);
      });
});
