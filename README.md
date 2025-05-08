# Cobalt Mine

> This is a basic HTML, JS, Golang web application to connect and display information regarding the US Geological Survey.

## 📦 Features

- Retrive tables of specified request
- Display mines information when given a feature type and the material wanted
- Displays a list of resoures that have a contained estimate and display information about them
- Using github.com/bmizerany/pat v0.0.0-20210406213842-e4b6760bdd6f, github.com/go-sql-driver/mysql v1.9.2, github.com/joho/godotenv v1.5.1

## 🚀 Getting Started

### Prerequisites

- Go 1.24.1 
- MySQL or compatible DB

### Installation

```bash
# Clone the repo
git clone https://github.com/gabeTheGrammer/cobaltMine
(Or download the files directly)

# Go into the project directory
cd project-name

# Run the tidy command
go mod tidy
(If the dependinces still don't install use 'go get required_depenincy' i.e. 'go get github.com/bmizerany/pat')

# Go to secret folder or create it
cd secret
or
mkdir secret

# Add a data.env using a editor (Or any prefrences of file editor)
vim data.env

# In this file add the following as well as your own information

DB_USERNAME=YOUR_DB_USERNAME
DB_PASSWORD=YOUR_DB_PASSWORD
DB_IP=YOUR_IP_ADDRESS
DB_NAME=YOUR_DB_NAME

(Don't include spaces)

# Run the server
go run cmd/*

```
Now visit http://localhost:8080

### Usage

You will start on a home page with no pictures at the moment as I was creating the MVP. From here if you click the 3 bars in the top you can see the three options.
User choice will allow you to explore tables using the provided drop down just select one and press submit to begin looking! Commodity Filter will find you a material you
enter as well as feature type from the drop down and when submit is clicked information will be displayed. The resource indicator tab will allow you to click a button to
display all data with a estimate contained column.
