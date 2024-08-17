<!-- Weather Data Project -->
<!-- Overview -->
This project is a Go application for fetching, storing, and querying weather data using the OpenWeatherMap API and a PostgreSQL database. The application includes functionality for:

1.Fetching weather data from OpenWeatherMap.
2.Storing weather data in a PostgreSQL database.
3.Generating daily weather summaries.
4.Serving weather data via a REST API.

Features
1.Weather Data Fetching: Fetches current weather data from the OpenWeatherMap API.
2.Data Storage: Stores weather data in a PostgreSQL database.
3.Daily Weather Summaries: Computes and stores daily summaries of weather data.
4.REST API: Provides endpoints to access weather data and summaries.
5.Retry Logic: Handles database connection retries to ensure reliability.


<!-- Prerequisites -->
Go 1.18 or higher
Docker and Docker Compose
PostgreSQL
OpenWeatherMap API key
Setup

1. Clone the Repository
bash
Copy code
git clone https://github.com/dushyant1435/WeatherWatch.git
cd WeatherWatch

2. Configure Environment Variables
Create a .env file in the root directory and add your environment variables:

env
Copy code
POSTGRES_USER=postgres
POSTGRES_PASSWORD=mysecretpassword
POSTGRES_DB=codedb
DATABASE_URL=postgres://postgres:mysecretpassword@db:5432/codedb?sslmode=disable

3. Build and Run the Application
Using Docker Compose:

Build and start the services:

cd server
docker-compose up --build
This command will set up the PostgreSQL database and the Go server.
 
<!-- to start client  -->
1. open client folder
2. click on index.html file
3. go live on this file