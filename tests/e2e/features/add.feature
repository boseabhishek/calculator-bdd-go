Feature: Add
  As a calculator user
  I want to use a calculator to perform basic addition operations

  Background: Background name
    Given I set up the calculator

# Each scenario finds it's steps inside the InitializeScenario steps definition
  Scenario: Add two numbers
    Given I have the first number 1
    And I have the second number 2
    When I add the numbers
    Then the result should be 3
