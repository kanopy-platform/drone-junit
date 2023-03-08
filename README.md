# drone-junit

A Card plugin for Drone that nicely displays test suites in a web UI.

## JUnit

The JUnit specification was chosen because all popular testing frameworks have an option to output test results in this [xml spec](https://llg.cubic.org/docs/junit/). This is a common format supported on CI platforms.

## Microsoft Adaptive Cards

 To update the design of the card use [https://adaptivecards.io/](https://adaptivecards.io/)

 The `cards.json` is packaged and served from Github Pages.  This is strictly the presentation layer.


## Example Pipeline

The plugin currently the unit test file to be named `unit-tests.xml` This may be changed later to allow customization.


```yaml
---
kind: pipeline
type: kubernetes
name: default

trigger:
  branch: [main]

volumes:
  - name: unit-tests
    temp: {}

steps:
  - name: test
    image: golangci/golangci-lint:v1.42.1
    volumes:
      - name: unit-tests
        path: /tmp
    commands:
      - make ci

  - name: display-results
    image: public.ecr.aws/kanopy/drone-junit:skunkworks 
    volumes:
      - name: unit-tests
        path: /tmp
```
