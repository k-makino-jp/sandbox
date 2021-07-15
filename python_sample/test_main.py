#coding: utf-8
# coverage report:
#   pytest -v -s --cov=./
#   pytest -v -s --cov=./ --cov-report=html
import main
from unittest.mock import patch

def test_01():
    main.print_result_add_and_sub()
    with patch("main.add") as add_mock:
        add_mock.return_value = 100
        main.print_result_add_and_sub()

@patch("main.add")
@patch("main.sub")
def test_02(sub_mock, add_mock):
    add_mock.return_value = 200
    sub_mock.return_value = 300
    main.print_result_add_and_sub()
