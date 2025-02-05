# Project Structure

This document outlines the structure of the `Backend` project.

## Root Directory

- **/Backend/**
    - Contains the main backend project files and directories.

## Directories

- **/cmd/**
    - Contains the entry point of the application.
    - **/api/**: The main application entry point (e.g., `main.go`).

- **/docs/**
    - Contains documentation files related to the project.
    - **decisions.md**: Documents the structure and decisions of the project.

- **/internal/**
    - Contains the core application code (encapsulated and not exposed to external projects).
    - **/db/**: Database-related code (e.g., connection setup).
    - **/handlers/**: HTTP handlers (e.g., API endpoints).
    - **/models/**: Data models and entities.
    - **/repositories/**: Data access layer (e.g., database queries).
    - **/services/**: Business logic layer.

- **/migrations/**
    - Contains database migration files.

- **/pkg/**
    - Contains reusable packages and utilities.
    - **/utils/**: Utility functions and helpers.

- **/scripts/**
    - Contains helper scripts for running, building, or deploying the application.
    - **migrate.sh**: Script for running database migrations.
    - **start.sh**: Script for starting the application.

- **/tests/**
    - Contains test cases and testing utilities.
    - **/integration/**: Integration tests.
    - **/e2e/**: End-to-end tests.

- **/web/**
    - Contains static files and templates for the web application.
    - **/static/**: Static files (e.g., CSS, JS, images).
    - **/templates/**: HTML templates.

## Files

- **.env**
    - Contains environment variables for the application.

- **.gitignore**
    - Specifies files and directories to be ignored by Git.

- **docker-compose.yml**
    - Docker Compose configuration for containerized development.

- **modd.conf**
    - Configuration file for Modd (used for live reloading during development).

- **go.mod**
    - Go module file for dependency management.

- **README.md**
    - Provides an overview and instructions for the project.

## Dependencies
    - Chi -> router for building Go HTTP services 
    - Godotenv -> A Go port of Ruby's dotenv library
    - lib/pq ->  Go Postgres driver for database/sql 

## Summary

This structure helps in organizing the project files and maintaining a clean architecture. Each directory and file has a specific purpose, making the project easier to navigate and manage.