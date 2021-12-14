Feature: Subtraction happens with addition

  Scenario:
    Given calculator is cleared
    When i press "<press_num>"
    And i add "<add_num>"
    And i subtract "<sub_num>"
    Then the result should be "<result_num>"

    Examples: 
      | press_num    | add_num         | sub_num          | result_num |
      | 2            | 5               | 3                | 4          |
      | 10           | 4               | 8                | 6          |