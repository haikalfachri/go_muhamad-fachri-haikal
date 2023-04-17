### Repository for deploy:
https://github.com/hklfach/go-learn-deploy

### Public IP Address
http://13.229.134.70/

### Steps
1. Open terminal (You can use CMD if using windows)

2. Connect to instance (Different based on instance configure)

    Formula command
    ```
    ssh -i "public_key" username@Public_IPv4_DNS
    ```
    
    Example command
    ```
    ssh -i "Haikal_18.pem" ec2-user@ec2-13-229-134-70.ap-southeast-1.compute.amazonaws.com
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

4. Clone repository
    ```
    git clone https://github.com/hklfach/go-learn-deploy
    ```

5. Configure nginx for reverse proxy

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
                proxy_pass http://127.0.0.1:8080;
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

6. Use docker and nginx

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

7. Open program in public IP Adress

    For my public ip address

    http://13.229.134.70/


