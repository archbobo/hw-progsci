#!/usr/bin/env python
import random 

def primes(n):
    "Add one to limit in order to generate the range with upper bound just equal to limit"
    act_limit = n + 1
    not_prime = set()
    primes = []
    
    for i in range(2, act_limit):
        if i in not_prime:
            continue
        for f in range(i * 2, act_limit, i):
            not_prime.add(f)
        primes.append(i)
    return primes

print primes(48)

def sumOfPrimes(k):
    "Use a random strategy to generate multiple possibilies of a b pairs"
    primes_list = primes(k)
    random.shuffle(primes_list)
    for a in primes_list:
        for b in primes_list:
            if (a <= b) and ((a + b) == k):
                return a, b
    return 0

print sumOfPrimes(48)

def allSumOfPrimes(k):
    "Append all of the a b pairs to one result list and return"
    primes_list = primes(k)
    # print primes_list
    rslt = []
    for a in primes_list:
        for b in primes_list:
            if (a <= b) and ((a + b) == k):
               rslt.append((a, b))
    return rslt

print allSumOfPrimes(48)
    
def goldbach(k):
    "Just make use of previous functions"
    triples = []
    temp_z = 0
    for z in range(4, k + 1):
        if z % 2 == 0:
            a, b = sumOfPrimes(z)
            triples.append((z, a, b))
            temp_z = z
    if z <= k:
        return triples, True
    else:
        return triples, False

print goldbach(10)

def goldbachWidth(k):
    "Simple function"
    rslt = {}
    for z in range(3, k + 1):
        if z % 2 == 0:
            rslt[z] = len(allSumOfPrimes(z))
    return rslt

print goldbachWidth(25)
