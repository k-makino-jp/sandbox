#coding: utf-8

def print_result_add_and_sub():
    print("INFO print_result_add_and_sub executed.")
    print("add result: {}".format(add(1,2)))
    print("sub result: {}".format(sub(2,1)))

def sub(x, y):
    return x - y

def add(x, y):
    return x + y

def main():
    print("INFO Start Main Method.")

if __name__ == "__main__":
    main()