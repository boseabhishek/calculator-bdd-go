Feature: Divide
  As a calculator user
  I want to use a calculator to perform basic division

  Background: Background name
    Given I set up the calculator

# Each scenario finds it's steps inside the InitializeScenario steps definition
  Scenario: Divide two numbers
    Given I have the first number 8
    And I have the second number 2
    When I divide the numbers
    Then the result should be 4
