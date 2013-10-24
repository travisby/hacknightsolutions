# Gets the last 10 digits from the sum of the function x^x for i = 1 and max = 1000
#!/usr/bin/env python2

# No!  The built-in math library convers to doubles, and is too small for our data
# import math


START = 1
STOP = 10
LAST_NUMS = 10


# xrange is lazy, where range is eager
print(str(reduce(lambda a, b: a + pow(b, b), range(START, STOP + 1)))[-LAST_NUMS:])
