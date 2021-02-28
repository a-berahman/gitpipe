# Welcome to GitPipe!
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/a-berahman/gitpipe)](https://goreportcard.com/report/github.com/a-berahman/gitpipe) [![Build Status](https://travis-ci.org/a-berahman/gitpipe.svg?branch=master)](https://travis-ci.org/github/a-berahman/gitpipe) ![Build Status](https://codebuild.eu-west-3.amazonaws.com/badges?uuid=eyJlbmNyeXB0ZWREYXRhIjoiNnEyMEo3N08rV0UrNithb0Yzank0OXdoc0Z0Q1E5aDlUN0ZpTXdWY0JRRFRIVkxWTFBrYlVXQm5wQXJvS2tPTTRLL2dnN05VaU5pM3FRYUhwUTZoY0g0PSIsIml2UGFyYW1ldGVyU3BlYyI6IjhhUHlHK2xEWFAzYVowaU0iLCJtYXRlcmlhbFNldFNlcmlhbCI6MX0%3D&branch=master)

This is the code repository for my challenge with Golang and Automation pipeline(CI/CD) with Dockerized, Scalability and Resilient  approach
## information
![it is a mind map for peresenting](https://i.ibb.co/z6H3X9F/mind-map.jpg)

![automation](https://i.ibb.co/t8z0fNb/CICD.jpg)

When you push any changes in the GitHub repository, the pipeline runs and the first build makes the latest docker image version of your repository and then the first build pushes the image to DockerHub .next, the second build runs and it pushes the latest image from docker hub to Kubernetes by using docker-compose  (i used AWS ECS fargate as serverless container service), in other words, it is APIs level.

The application has an interface as a website published on Elastic Beanstalk, and, the API level is only accessible by Elastic Beanstalk and no one can send the request directly to APIs. So, when clients send the request to the website, the traffic handled by autoScaling and load balancer and after that requests are sent via a second load balancer and auto-scaling to the API level.

finally, Through an Event Bridge, we call an API for Gists update for users every three hours, and the API is not accessible by website.

*notes: in serverless services that I used in the architect, I pay only for usage, also handle HealthCheck, autoscaling and auto balancer with multi-region.

![application](https://i.ibb.co/k4K1L1J/applicaiton.jpg)
## Installation
For installing the dependencies and starting the server, please run below commands:
```sh
make start
make run
```

