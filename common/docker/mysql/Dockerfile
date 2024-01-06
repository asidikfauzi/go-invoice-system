# Use MySQL 5.7 as the base image
FROM mysql:5.7

# Set environment variable to allow an empty password
ENV MYSQL_ALLOW_EMPTY_PASSWORD=yes

# Move the MySQL data directory to the desired location
VOLUME /var/lib/mysql /var/lib/mysql

# Expose port 3306
EXPOSE 3306
