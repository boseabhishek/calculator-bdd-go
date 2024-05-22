# building a calculator using godog BDD

## TIL

- **Feature** files are just a way of organising scenarios, gives you file/proj org capability. It doesn't any way dictate the run or sequence etc.

- **Background** is invoked first. They are same as scenarios.
  
- **Scenarios** matters the most!
  - For example:
    ```go
        Scenario: Add two numbers
            Given I have the first number 1
            And I have the second number 2
            When I add the numbers
            Then the result should be 3
    ```
    Everytime the above scenario is invoked, all the steps are pattern matched inside the `InitializeScenario` func which is already been at play and subsequent func is called.
    `I have the first number 1` calls step `ctx.Step(`^I have the first number (\d+)$`, s.iHaveTheFirstNumber)`