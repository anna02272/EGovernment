# EGovernment

This project uses the Go programming language for the backend, Angular for the frontend, containerization with Docker single sign-on for user authentication, and a microservices architecture for communication between different services.

This is a project with developed 4 systems for EGovernment: MUP Vehicles, Traffic Police, Magistrates Court, and the Institute of Statistics.

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

### Description:

#### MUP Vehicles
This service under the Ministry of Internal Affairs enables the creation of driver's licenses, drivers, vehicle registration, searching, and report generation. Used by police officers, it allows for efficient information management, crucial for the operation of state institutions like the MUP.

#### Traffic Police
This application manages traffic police operations within the eGovernment system. It includes functionalities for recording violations and accidents, generating PDF documents, sending email notifications, and tracking the status of violations, facilitating communication between police, citizens, and the MUP.

#### Magistrates Court
This system handles the procedure of making judgments based on committed violations. It allows for scheduling hearings, making decisions on violations, changing case statuses, viewing notifications about hearings, and informing judges about scheduled hearings and their details.

#### Institute of Statistics
This system enables the viewing and generation of statistical data based on information from other systems. It provides users with access to statistical data and allows them to request additional information, aiding in efficient data processing and review.

### Design Specification
Class diagram and use case for each system.

#### MUP Vehicles
![MupVozila_ClassDiagram](https://github.com/user-attachments/assets/671cf367-fb9f-4d53-bf4a-74e21880d3f7)
![MupVozila_UseCaseDiagram](https://github.com/user-attachments/assets/e285317d-33ac-4297-8c46-3e25b0f967a8)

#### Traffic Police
![SaobracajnaPolicija_ClassDiagramUpdate](https://github.com/user-attachments/assets/f4a572d4-b7bc-4d50-bf85-a1359f8c9add)
![SaobracajnaPolicija_UseCaseDiagramUpdate](https://github.com/user-attachments/assets/89cddd07-eb48-4749-bbe7-19aa90ebcb3b)

#### Magistrates Court
![PrekrsajniSud_UseCase drawio (1)](https://github.com/user-attachments/assets/eb8e62c5-8b6d-49e5-8f63-dbd51683dee5)
![PrekrsajniSudClassDiagram (2)](https://github.com/user-attachments/assets/08c85b8f-5ecf-4116-af7b-9c54efbb362c)

#### Institute of Statistics
![ZavodZaStatistiku_ClassDiagram](https://github.com/user-attachments/assets/4fd24e16-c4ed-4fc6-b683-b4e404a8171f)
![ZavodZaStatistiku_UseCaseDiagram](https://github.com/user-attachments/assets/b268c514-94ab-4ba8-805e-1bb9e465d47b)

### Images of project:
![registracija](https://github.com/user-attachments/assets/78a5ad79-3b92-4e3b-a62c-c68d3a794da6)
![prijava](https://github.com/user-attachments/assets/c4e7eab9-1eba-4d60-8489-58fd68e6de96)
![pocetna](https://github.com/user-attachments/assets/bd6377f2-ca44-4b24-b9a0-3325ed2ae190)

#### MUP Vehicles
#### Traffic Police
#### Magistrates Court
#### Institute of Statistics
Report on the type of traffic accident
![izvestaj o tipu saobracajne nesrece](https://github.com/user-attachments/assets/170a6df6-c326-4b3b-92c8-a9b9ead719fd)
![izvestaj o tipu saobracajne nesrece 2 popunjen](https://github.com/user-attachments/assets/689b724d-1356-4163-8124-d671a1ee6385)

Report on the degree of traffic accident
![Izvestaj o stepenu saobracajne nesrece 2](https://github.com/user-attachments/assets/4c1faeb8-c641-4c60-953a-324f99e9d890)
![Izvestaj o stepenu saobracajne nesrece 2 popunjen](https://github.com/user-attachments/assets/5cb961ad-3730-47dd-913f-3ad4460cabcc)

Report on the type of violation
![izvestaj o tipu prekrsaja](https://github.com/user-attachments/assets/bc0440d0-0d66-4da5-8808-252d53e490e4)
![izvestaj o tipu prekrsaja 2 popunjen](https://github.com/user-attachments/assets/3679259d-c52b-4546-97a1-4db1cf4880b7)

Report on the number of registered vehicles
![izvestaj o broju registovanih](https://github.com/user-attachments/assets/2bcbf5c4-359c-41ad-a8f6-e9b270c46ca8)
![izvestaj o broju registovanih 2 popunjen](https://github.com/user-attachments/assets/1794e8e5-770c-412d-ad26-5c3547823309)

Request for additional information
![zahtev](https://github.com/user-attachments/assets/272c2ea0-bd8c-43ed-b9e5-c082e148ef47)

All requests for additional information
![zahtevi](https://github.com/user-attachments/assets/710636b5-dcf8-4ead-905f-30ce0ef5c60e)

Response to the request for additional information
![odgovor na zahtev](https://github.com/user-attachments/assets/ddd6182d-2f91-495e-87da-4de9d801a57b)

Response to the request for additional information - email
![odgovor na zahtev email](https://github.com/user-attachments/assets/b4b5ab05-10ff-457d-9e9e-e2c2a9324c47)





