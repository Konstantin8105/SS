package nginx

import (
	"fmt"

	"github.com/Konstantin8105/ss/starter"
)

func init() {
	var n Nginx
	n.ProgramName = "nginx"
	starter.Register(n.ProgramName, &n)
}

// Nginx - program `nginx`
type Nginx struct {
	starter.SimpleInstall
}

// Run - running preparing the program
func (n Nginx) Run() (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("Cannot run %v. err = %v", n.ProgramName, err)
		}
	}()

	// install
	n.SimpleInstall.Run()
	if err != nil {
		return err
	}

	// TODO: add configuration
	// TODO: sudo /etc/init.d/nginx restart
	// TODO: sudo /etc/init.d/mysql status
	// TODO: add logging
	// log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
	//                   '$status $body_bytes_sent "$http_referer" '
	//                   '"$http_user_agent" "$http_x_forwarded_for"';

	// TODO:
	/*

		## nginx

		## Basic configuration

		```
		systemctl status nginx
		sudo systemctl restart nginx
		sudo systemctl stop nginx
		```

		## Typical configuration:

		```conf
		############################
		## Defines which Linux system user will own and run the nginx server.
		user www-data;
		############################
		## Defines how many threads, or simultaneous instances, of nginx to run.
		worker_processes auto;
		############################
		## Defines where nginx will write its master process ID, or PID. The PID is used by the operating system to keep track of and send signals to the nginx process.
		pid /run/nginx.pid;
		############################

		events {
			###############################
			# minimal allowable connection
			# for avoid DDoS attack, we
			# have to have ~65000
			###############################
			worker_connections 65000;
			###############################
			# multi_accept on;
		}

		http {

		########################################
		########################################
		########################################
		########################################
		########################################
		########################################
		########################################
		########################################
		########################################
		########################################

			## Example of service
		#	server {
		#		listen 80 default_server;
		#		listen [::]:80 default_server ipv6only=on;
		#		index index.html index.htm;
				####################
				# Edit for each service
				#
		#		server_name localhost;
				#
				####################
		#		location / {
					####################
					# Edit for each service
					#
		#			proxy_pass http://127.0.0.1:8000;
					#
					####################
					# First attempt to serve request as file, then
					# as directory, then fall back to displaying a 404.
		#			try_files $uri $uri/ /index.html;
		#		}
				# Media: images, icons, video, audio, HTC
		#		location ~* \.(?:css|js|jpg|jpeg|gif|png|ico|cur|gz|svg|svgz|mp4|ogg|ogv|webm|htc)$ {
		#			expires 1M;
		#			access_log off;
		#			add_header Cache-Control "public";
		#		}
		#	}

		########################################
		########################################
		########################################
		########################################
		########################################
		########################################
		########################################
		########################################
		########################################
		########################################
		########################################
			##
			# Basic Settings
			##

			sendfile on;
			tcp_nopush on;
			tcp_nodelay on;
			keepalive_timeout 65;
			types_hash_max_size 2048;
			# server_tokens off;
			# client_max_body_size 32M;

			# server_names_hash_bucket_size 64;
			# server_name_in_redirect off;

			include /etc/nginx/mime.types;
			default_type application/octet-stream;

			##
			# SSL Settings
			##

			ssl_protocols TLSv1 TLSv1.1 TLSv1.2; # Dropping SSLv3, ref: POODLE
			ssl_prefer_server_ciphers on;

			##
			# Logging Settings
			##
			log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
		                      '$status $body_bytes_sent "$http_referer" '
		                      '"$http_user_agent" "$http_x_forwarded_for"';

			# Every request to your web server is recorded in this log file unless Nginx is configured to do otherwise.
			# access_log /var/log/nginx/access.log main;
			access_log off;

			# Any Nginx errors will be recorded in this log.
			# error_log /var/log/nginx/error.log;
			error_log /dev/null crit;

			##
			# Gzip Settings
			##

			gzip on;
			gzip_disable "msie6";
		}

		```








		```conf
		############################
		## Defines which Linux system user will own and run the nginx server.
		user www-data;
		############################
		## Defines how many threads, or simultaneous instances, of nginx to run.
		worker_processes auto;
		############################
		## Defines where nginx will write its master process ID, or PID. The PID is used by the operating system to keep track of and send signals to the nginx process.
		pid /run/nginx.pid;
		############################

		events {
			###############################
			# minimal allowable connection
			# for avoid DDoS attack, we
			# have to have ~65000
			###############################
			worker_connections 65000;
			###############################
			# multi_accept on;
		}

		http {

			## Example
			server {
				listen 80;
				server_name  www.ve-server1.com;
				location / {
					proxy_pass http://127.0.0.1:8000
				}
			}

			## Example
			server {
		        listen 80 default_server;
		        listen [::]:80 default_server ipv6only=on;

		        root /usr/share/nginx/html;
		        index index.html index.htm;

		        # Make site accessible from http://localhost/
		        server_name localhost;

		        location / {
		                # First attempt to serve request as file, then
		                # as directory, then fall back to displaying a 404.
		                try_files $uri $uri/ /index.html;
		                # Uncomment to enable naxsi on this location
		                # include /etc/nginx/naxsi.rules
		        }

				# Media: images, icons, video, audio, HTC
				location ~* \.(?:jpg|jpeg|gif|png|ico|cur|gz|svg|svgz|mp4|ogg|ogv|webm|htc)$ {
					expires 1M;
					access_log off;
					add_header Cache-Control "public";
				}
				location ~*  \.(jpg|jpeg|png|gif|ico|css|js)$ {
					expires 365d;
				}
			}

		# CSS and Javascript
		# location ~* \.(?:css|js)$ {
		#  expires 1y;
		#  access_log off;
		#  add_header Cache-Control "public";
		#}
		#	}

			##
			# Basic Settings
			##

			sendfile on;
			tcp_nopush on;
			tcp_nodelay on;
			keepalive_timeout 65;
			types_hash_max_size 2048;
			# server_tokens off;
			# client_max_body_size 32M;

			# server_names_hash_bucket_size 64;
			# server_name_in_redirect off;

			include /etc/nginx/mime.types;
			default_type application/octet-stream;

			##
			# SSL Settings
			##

			ssl_protocols TLSv1 TLSv1.1 TLSv1.2; # Dropping SSLv3, ref: POODLE
			ssl_prefer_server_ciphers on;

			##
			# Logging Settings
			##
			log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
		                      '$status $body_bytes_sent "$http_referer" '
		                      '"$http_user_agent" "$http_x_forwarded_for"';

			# Every request to your web server is recorded in this log file unless Nginx is configured to do otherwise.
			access_log /var/log/nginx/access.log main;
			# access_log off;

			# Any Nginx errors will be recorded in this log.
			error_log /var/log/nginx/error.log;

			##
			# Gzip Settings
			##

			gzip on;
			gzip_disable "msie6";

			# gzip_vary on;
			# gzip_proxied any;
			# gzip_comp_level 6;
			# gzip_buffers 16 8k;
			# gzip_http_version 1.1;
			# gzip_types text/plain text/css application/json application/javascript text/xml application/xml application/xml+rss text/javascript;

			##
			# Virtual Host Configs
			##

			#include /etc/nginx/conf.d/*.conf;
			#include /etc/nginx/sites-enabled/*;
		}
		```
	*/
	return nil
}
