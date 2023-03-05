# Auth

## Local Development Guide

A `.env` secret file is needed to run the backend server. Request a copy of this file through our slack channel.

To start the backend server, make sure you are in the `server` folder by running `cd server` from the project root, then run `make run` to start the backend server.

To start the frontend client, make sure you are in the `client` folder by running `cd client` from the project root, then run `npm install` to install necessary packages, and run `npm start` to start the frontend.

## Backend API

### Users

| Description         | Route                                     | Body                                 | Auth  |
|---------------------|-------------------------------------------|--------------------------------------|-------|
| Cretae user         | `POST /users`                             | `name`, Optional: `pronouns`         | Admin |
| Get user by id      | `GET /users/{userId}`                     |                                      | Admin |
| Delete user by id   | `DELETE /users/{userId}`                  |                                      | Admin |

### Roles

| Description         | Route                                     | Body                                 | Auth  |
|---------------------|-------------------------------------------|--------------------------------------|-------|
| Assign role         | `POST /users/{userId}/roles`              | `role`                               | Admin |
| Unassign role       | `DELETE /users/{userId}/roles/{roleId}`   |                                      | Admin |
| List roles          | `GET /users/{userId}/roles`               |                                      | Admin |
