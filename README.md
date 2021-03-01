# Welcome to GitPipe!
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/a-berahman/gitpipe)](https://goreportcard.com/report/github.com/a-berahman/gitpipe) ![Build Status](https://codebuild.eu-west-3.amazonaws.com/badges?uuid=eyJlbmNyeXB0ZWREYXRhIjoiNnEyMEo3N08rV0UrNithb0Yzank0OXdoc0Z0Q1E5aDlUN0ZpTXdWY0JRRFRIVkxWTFBrYlVXQm5wQXJvS2tPTTRLL2dnN05VaU5pM3FRYUhwUTZoY0g0PSIsIml2UGFyYW1ldGVyU3BlYyI6IjhhUHlHK2xEWFAzYVowaU0iLCJtYXRlcmlhbFNldFNlcmlhbCI6MX0%3D&branch=master)

This is the code repository for my challenge with Golang and Automation pipeline(CI/CD) with Dockerized, Scalability and Resilient  approach
## information

![automation](https://i.ibb.co/crwDYS1/CICD.jpg)

When any changes push in the GitHub repository, the pipeline runs and the first build makes the latest docker image version of your repository and then the first build pushes the image to DockerHub (ahmadberahman/gitpipe:latest). Next, the second build runs and it pushes the latest image from docker hub to orchestration by using docker-compose  (i used AWS ECS Fargate as serverless container service), in other words, this is APIs level.

The application has an interface as a website published by Docker image (ahmadberahman/gitpipeui:latest) on ECS Fargate and presents by Cloud Front, and, the API level isn't accessible by the user directly. So, when clients send the request to the website, the traffics handle by autoScaling and load balancer and after that requests are sent via a second load balancer and auto-scaling to the API level.

Finally, Through an Event Bridge, we call the API for users Gists to update every three hours, and the API is not accessible by website.

*notes: in serverless services that I used in the architect, I pay only for usage, also I handle health-check, autoscaling and auto balancer in multi-region.

![application](https://i.ibb.co/k4K1L1J/applicaiton.jpg)

The application contains four levels:
Router: presents APIs expose routs and connects them to handler functions
Handler: implements behaviour of each API rout
Repository: implements database interface(i used mongo cloud DB)
Service: implements GitHub and Pipedrive APIs interface

I used two design patterns, Factory for making instance from interfaces, and, future/promise to handle concurrency in the action that works as a gist list updater which runs every three hours.
Each file has a unit-test, and whenever you run tests (go test ./...)  you can monitor the behaviour of the application through logs. Also, each function provides dedicate documentation.

Requirements:

![it is a mind map for peresenting](https://i.ibb.co/z6H3X9F/mind-map.jpg)
## Installation
To install the dependencies and start the server, please run below commands:
```sh
make start
make run
```

