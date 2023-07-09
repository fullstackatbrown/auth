# Auth

## Development Environment

### Codespaces

This project is configured to run in GitHub Codespaces. To start developing, create a codespace for this [project](https://github.com/fullstackatbrown/auth) if you haven't already, and open that codespace in your editor of choice. All secrets are already configured.

### Devcontainer

To develop in a local devcontainer, you need to have [Docker](https://www.docker.com/) installed. Then, open the project in VSCode, and click the green button in the bottom left corner to open the project in a devcontainer. You will need to add a secret `.env` file. Ask in slack for the secrets.

### Local

You need golang and nodejs installed to run the backend and frontend locally. You will need a secret `.env` file. Ask in slack for the secrets.

## Starting Development

### Backend

To start the backend, run `make backend` in the root. The bakcend will be listening on port 8000 of localhost. 

### Dashboard

To start the dashboard, run `make dashboard` in the root. The dashboard will be listening on port 3000 of localhost.

## API Documentation

The API documentation for the backend is available at [SwaggerHub](https://app.swaggerhub.com/apis-docs/tianrendong/Auth/1.0).