def fibonacci (n):
	if (n == 0):
		return 0
	elif (n == 1):
		return 1
	else:
		return fibonacci(n - 1) + fibonacci(n - 2)

if __name__ == '__main__':
	i = 1
	this_fib = fibonacci(i)
	total_sum = 0
	while (this_fib < 4000000):
		i += 1
		if this_fib % 2:
			total_sum += this_fib
		this_fib = fibonacci(i)
	print 'Total sum of even terms: ' + str(total_sum)
