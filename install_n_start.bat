::start all the services
call start_mongo
call data-service-picturestore\install_n_start
call auth-service-picturestore\install_n_start
call front-end-picturestore\install_n_start
call picture-service-picturestore\start_server