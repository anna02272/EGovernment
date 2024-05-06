# EGovernment

### Launch Guide:

Below are the steps to get the project up and running on your local environment.

#### Prerequisites:
1. Docker
2. Go programming language
3. Goland IDE
6. Node.js
7. Visual Studio Code

#### Step 1: Clone the Repository
Clone the repository to your local machine:

```bash
git clone https://github.com/anna02272/EGovernment
```

#### Step 2: Import the Project in Goland
Open Goland IDE and import the cloned repository.

#### Step 3: Run Go Mod Tidy
Ensure all necessary dependencies are downloaded:

```bash
go mod tidy
```

#### Step 4: Build Docker Images and Start Docker Compose

```bash
docker-compose build
docker-compose up
```

#### Step 5: Start Angular Frontend
1. Open Visual Studio Code.
2. Navigate to the frontent directory in the cloned repository.
3. Start the Angular app with npm:

```bash
npm install
npm start
```

#### Step 6: Access the Platform
Once the services are up and running, you can access the platform via the provided endpoints:

- **Frontend:** Open a web browser and go to [https://localhost:4200/](https://localhost:4200/)
