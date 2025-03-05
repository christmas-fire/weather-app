# Weather App

## Description

The `weather` application is a CLI tool that fetches and displays the current weather for a specified city using the OpenWeatherMap API. If no city is provided, it defaults to "Novosibirsk".

## Installation

### Steps to Install

1. **Clone the Repository:**
   ```sh
   git clone https://github.com/christmas-fire/weather-app.git
   cd weather-app
2. **Install Dependencies:**
    ```sh
    go mod download
3. **Build the Application:**
    ```sh
    go build -o weather cmd/main.go
4. **Move the Executable to a Directory in Your PATH:**
    ```sh
    sudo mv weather /usr/local/bin/
## Usage
- **Basic Usage**
    ```sh
    weather [global options] command [command options] [city]
- **Commands**
    ```sh
    help, h : Shows a list of commands or help for one command.
- **Global Options**
    ```sh
    --help, -h : Show help.
## Examples
1. **Fetch Weather for Novosibirsk (default city):**
    ```sh
    weather
2. **Fetch Weather for a Specific City:**
    ```sh
    weather Moscow
3. **Show Help:**
    ```sh
    weather --help
## Configuration
The application uses a configuration file to manage settings such as the API key, language and default city.
File located in `/config/config.yaml`
### Example
```yaml
api_key: "your_api_key" # required
lang: "your_lang" # required
city: "your_city" # optional (if you use command args)
```
## Output
The weather data is displayed in a colorful table format.
Here is an example of the output:
```md
|------------|-------------|---------------|-------------|------------|
|    DATE    |    CITY     |    WEATHER    | TEMPERATURE | FEELS LIKE |
|------------|-------------|---------------|-------------|------------|
| 05-03-2025 | Novosibirsk | ☀️ Clear sky  | -8°         | -15°       |
|------------|-------------|---------------|-------------|------------|
```