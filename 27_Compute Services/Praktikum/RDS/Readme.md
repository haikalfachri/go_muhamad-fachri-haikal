### Repository for deploy:
https://github.com/hklfach/go-learn-deploy/tree/rds-connection

### Public IP Address
http://18.136.126.223/

### RDS Endpoint and Port
mysql-1.cmaggkp0ghea.ap-southeast-1.rds.amazonaws.com:3306

### Steps
1. Open terminal (You can use CMD if using windows)

2. Connect to instance (Different based on instance configure)

    Formula command
    ```
    ssh -i "public_key" username@Public_IPv4_DNS
    ```

    Example command
    ```
    ssh -i "Haikal_18.pem" ec2-user@ec2-18-136-126-223.ap-southeast-1.compute.amazonaws.com
    ```

3. Install docker, docker-compose, git, and nginx

    Get into super user mode
    ```
    sudo su
    ```

    Update package
    ```
    yum update -y
    ```

    Install git
    ```
    yum install git -y
    ```
    ```
    git -v
    ```

    Install docker
    ```
    yum install -y docker
    ```
    ```
    docker -v
    ```

    Install docker-compose
    ```
    curl -L https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m) -o /usr/bin/docker-compose && chmod +x /usr/bin/docker-compose
    ```
    ```
    docker-compose -v
    ```

    Install nginx
    ```
    amazon-linux-extras install nginx1.12
    ```
    ```
    nginx -v
    ```

    Install mysql server
    ```
    amazon-linux-extras install epel -y 
    ```
    ```
    yum install https://dev.mysql.com/get/mysql80-community-release-el7-5.noarch.rpm 
    ```
    ```
    yum install mysql-community-server
    ```

4. Clone repository
    ```
    git clone --single-branch --branch rds-connection https://github.com/hklfach/go-learn-deploy
    ```

5. Configure db in ./drivers/mysql/mysql.go
    Change the config based on RDS DB config

    Formula
    ```
    ...
    config := Config{
	  DB_Username: "{USERNAME}",
	  DB_Password: "{PASSWORD}",
	  DB_Port:     "{PORT}",
	  DB_Host:     "{RDS_ENDPOINT}",
	  DB_Name:     "{DATABASE_NAME}",
	}
    ...
    ```
    example
    ```
    ...
    config := Config{
	  DB_Username: "admin",
	  DB_Password: "Haikal*18",
	  DB_Port:     "3306",
	  DB_Host:     "mysql-1.cmaggkp0ghea.ap-southeast-1.rds.amazonaws.com",
	  DB_Name:     "crud_go",
	}
    ...
    ```

6. Configure nginx for reverse proxy

    Go to nginx directory
    ```
    cd /etc/nginx/
    ```

    Edit nginx.conf using nano
    ```
    nano nginx.conf
    ```

    Comment root variable in line 42 and add proxy_pass inside location {}
    ```
    ...
     server {
        listen       80 default_server;
        listen       [::]:80 default_server;
        server_name  _;
        # root         /usr/share/nginx/html;

        # Load configuration files for the default server block.
        include /etc/nginx/default.d/*.conf;

        location / {
                proxy_pass http://127.0.0.1:8000;
        }

        error_page 404 /404.html;
            location = /40x.html {
        }

        error_page 500 502 503 504 /50x.html;
            location = /50x.html {
        }
    }
    ...
    ```

7. Use docker, nginx and mysql server

    Start docker
    ```
    service docker start
    ```

    Get into go-learn-deploy directory
    ```
    cd go-learn-deploy
    ```

    Use docker compose
    ```
    docker-compose up -d
    ```

    Start nginx
    ```
    systemctl start nginx
    ```

    Mysql server is used for check the db through console (optional)

    Start mysql server
    ```
    systemctl start mysqld
    ```

8. Open program in public IP Adress and Test API using Postman

    Example
    ```
    http://18.136.126.223/
    ```
### Backup mysqldump for local migrate

    docker exec {CONTAINERID} sh -c 'exec mysqldump --all-databases -uroot -p"{DB_PASSWORD}"' > migrate.sql

### Connect EC2 instance to RDS trough console
    
    mysql -h mysql-1.cmaggkp0ghea.ap-southeast-1.rds.amazonaws.com -P 3306 -u admin -p
