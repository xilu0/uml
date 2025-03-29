import pdb; pdb.set_trace()

def average(numbers):
    total = sum(numbers)
    count = len(numbers)
    return total / count

numbers = [1, 2, 3, 4, 5]
print(average(numbers))