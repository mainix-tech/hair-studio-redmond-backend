FROM postgres:16-alpine

WORKDIR /app

# Expose the standard PostgreSQL port
EXPOSE 5432