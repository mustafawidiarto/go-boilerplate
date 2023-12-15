# Boilerplate (Golang)

## Overview

This boilerplate is used to standardize the directory structure for projects of medium to large complexity or with the potential for it. However, for other cases that only handle one or a few processes, it is not necessary to implement this boilerplate in order to avoid over-abstraction. For example, you can use a single file like main.go or a flat architecture instead.

## Directories

### `/command`

Main applications for this project. The directory name for each application should match the name of the executable you want to have. Don't put a lot of code in the directory.

### `/common`

The directory is used to hold code that is shared across different parts of the application. The common directory may contain utility functions, constants, error types, database connection and other modules that are used by multiple packages within the application. The purpose of the common directory is to avoid code duplication and to keep the shared code organized in one place.

### `/handler`

This directory will act as the presenter layer. Decide how the data will presented. Could be as REST API, or HTML File, or gRPC whatever the delivery type. This layer also will accept the input from user. Sanitize the input and sent it to Usecase layer.

### `/model`

The model layer is a representation of the application's business logic in the code. It contains the entities and value objects that the application uses to model its business concepts and rules. The domain layer is independent of any specific technology or implementation details and is designed to be reusable and independent.

### `/repository`

This directory containing adapters to different storage implementations. A data source might be an adapter to a SQL database, an elastic search adapter, or REST API. A data source implements methods defined on the repository and stores the implementation of fetching and pushing the data.

### `/usecase`

This directory will act as the business process layer, any process will handled here. This layer will decide, which repository layer will use. And have responsibility to provide data to serve into delivery. Process the data doing calculation or anything will done here. Usecase layer will accept any input from Delivery layer, that already sanitized, then process the input could be storing into DB , or Fetching from DB ,etc.

## Sample Use Case

Create tagging and storing room data in the database when a user initiates a chat on Qiscus Omnichannel by utilizing a new session webhook, besides that there is a cronjob that does an auto resolved room when the on going room has reached 10 minutes. And also create an API to get all rooms that are stored in the database with an API key authentication.

### Run Locally

You can use two method, Docker or Makefile:

To run the project locally using docker, follow these step:
- Install docker if not installed
- Navigate to the directory
- Copy docker-compose.local.yaml to compose.yaml
- Run docker-compose build
- Run docker-compose up -d
- Every change you make, you need to re-run docker-compose build & docker-compose up -d

To run the project locally, follow these steps:
- Install Make if not installed
- Navigate to the directory
- Format code and tidy modfile: **make tidy**
- Run the server: **make run app=server**, or run the application with reloading on file changes with: **make run/live app=server**
- The backend server will be accessible at **http://localhost:8080**
- Run the cronjob: **make run app=cron**, or run the application with reloading on file changes with: **make run/live app=cron**
