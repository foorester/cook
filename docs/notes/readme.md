# Notes

This file serves as a dump for important notes regarding the project.

## Database Migrations
Currently, there is no implemented database migrator to handle migrations. This means that executing migrations must be done manually, which can be a time-consuming process. It is important to prioritize the implementation of a database migrator to streamline this task and ensure database changes are properly managed.

## Temporary Duplication in Database Scripts
There is temporary duplication between the database generation and initial population scripts and the scripts used by Docker Compose for testing. This repetition is acknowledged and will be eliminated in future updates. Removing this duplication will help maintain a more efficient and organized codebase.

## Testing
Unit testing, integration testing, and end-to-end testing are recognized as important factors in the development process. It is evident that there is a debt in this project in terms of testing, but it is only temporary. The exploratory process aimed at defining the best possible structure involves frequent changes. Once a stable implementation of the service is achieved, dedicated time will be allocated to thoroughly test every corner of the application.
