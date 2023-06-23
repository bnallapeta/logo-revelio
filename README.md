# Logo Revelio

Logo Revelio is a web application that allows users to play a logo quiz game. Users are presented with various logos, and they have to guess the correct name of each logo. The application keeps track of the user's scores and provides a leaderboard to showcase the top players.

## Features

- Interactive logo quiz game with a wide range of logos to guess.
- Score tracking for each user.
- Leaderboard displaying the top players with their scores.
- User-friendly interface with a responsive design.

## Installation
- Simple run the below command from the root of the project:
  ```
  kubectl apply -f argocd/logo-revelio-app.yaml
  ```

## Technologies Used

- **Backend**: Go (Golang) programming language with Gin framework and GORM for the database.
- **Frontend**: HTML, CSS, JavaScript.
- **Database**: SQLite.
- **API**: RESTful API design.
- **Deployment**: The application can be deployed on a web server.

## Continuous Integration (CI) with Tekton Pipelines

Tekton is setup for the project to take care of the Continuous Integration needs. Various tasks and pipelines are defined to automate the process of building and deploying the application.

### Tasks
1. clone-build-push-task: Clones the source code, builds the Docker image and pushes it to the registry.

2. update-image-tag: Clones the source code, updates the Docker image tag in the Kubernetes deployment file, and creates a pull request with these changes.

### Pipelines
1. clone-build-push-pipeline: Executes the clone-build-push-task and update-image-tag tasks.

### PipelineRun
1. img-pipeline-run: Runs the clone-build-push-pipeline.


Pipelines as Code (PaC) is setup in order to watch for PR creation on the repository and trigger the PipelineRun based on the event.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvement, please feel free to open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).

## Authors

- [Bharath Nallapeta](https://github.com/bnallapeta)
