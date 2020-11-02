Feature: Say helloWorld
  In order to exchange pleasantries with server
  As an API user
  I need to be able invoke SayHelloworld

  Scenario: Retrieve response from server
    Given client is configured to contact server
    When I say hello to server with "<message>"
    Then server should respond with helloWorld "<message>"
    Examples:
    | message   |
    | Test      |
    | Apollo    |