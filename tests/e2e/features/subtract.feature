Feature: Subtract
  As a calculator user
  I want to use a calculator to perform basic subtraction operations

# Each scenario calls InitializeScenario everytime
  Scenario: Subtract two numbers
    Given I have the first number 5
    And I have the second number 3
    When I subtract the numbers
    Then the result should be 2
