--My solution to the triangle number words.
--I'd like to point out the whole map floor take sum thing. Basically
--what is going on is that triangle numbers is returning a list of doubles
--so in order to use a sexy list comprenension I actually have to waste some time
--maybe I'll find a better way to do it at some point.
--Everything else is rather straight fordward. Basically it's just a list
--comprehension that I grab triangle numbers from and the compair it to
--the sum of the letters in the word.

triangleNumbers = [ x | x <- [1..], y <- [1..x], x == (y / 2) * (y+1) ]

isTriangle :: String -> Bool
isTriangle word = elem sum (map floor (take sum triangleNumbers))
  where
    sum = (foldl (\acc l -> acc + (fromEnum l - 64)) 0 word)