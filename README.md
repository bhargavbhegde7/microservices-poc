# microservices-poc

This is a proof of concept application for microservices architecture.
There are 3 services each of which act as a separate web service, which can be hosted anywhere independant of each other.

1. Front end service:
  This service serves the front end static pages. The data is loaded into these pages via external api call to other two services.
  
2. Auth service:
  This uses JWT (JSON Web Token) pattern for user authentication. When the user signs up, her details are hashed and stored as a temporary token in the database.
  When she tries to login, this token is sent to her. 
  when she tries to access the home page the front end app sends this token to the data service which in turn verifies the truthfulness of the client using the auth service.

3. Data service
  This service is just a dummy service to show how the JWT verification works. The front end service sends the token to this service, asking for some particular data.
  This service sends the token to the auth service. Auth service verifies this token with the DB and replies with a "OK" (http 200) or "auth fail"(http 401)

TO RUN:

1. clone the repo
2. run "docker-compose up"
