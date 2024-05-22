Feature: Subtract
  As a calculator user
  I want to use a calculator to perform basic subtraction operations

  Background: Background name
    Given I set up the calculator
  
# Each scenario finds it's steps inside the InitializeScenario steps definition
  Scenario: Subtract two numbers
    Given I have the first number 5
    And I have the second number 3
    When I subtract the numbers
    Then the result should be 2
