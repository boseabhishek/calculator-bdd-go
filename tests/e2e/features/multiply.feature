Feature: Multiply
  As a calculator user
  I want to use a calculator to perform basic multiply operations

  Background: Background name
    Given I set up the calculator

# Each scenario finds it's steps inside the InitializeScenario steps definition
  Scenario: Multiply two numbers
    Given I have the first number 4
    And I have the second number 3
    When I multiply the numbers
    Then the result should be 12