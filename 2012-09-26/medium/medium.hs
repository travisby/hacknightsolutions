--My solution to the sexy primes. The prime generation is actually kinda slow but in this case I doubt it matters
primes :: [Int]
primes = sieve [2..]
  where
    sieve (p:xs) = p : sieve [x|x <- xs, x `mod` p > 0]

--check to see if a number is contained in the set of primes
isPrime :: Int -> Bool
isPrime x = elem x (take x primes)

--a function that passes a limit into a list comprehension which generates
--infinite sexy primes
sexyPrimes n = [ (x, x+6) | x <- [3..n-6], isPrime x, isPrime (x+6)]