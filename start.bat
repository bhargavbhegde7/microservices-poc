::start all the services
call start_mongo
call data-service-picturestore\start_server
call auth-service-picturestore\start_server
call front-end-picturestore\start_server
call picture-service-picturestore\start_server