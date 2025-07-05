#!/usr/bin/python3

import random

# Generate a password of Size length
# allowing users to select upper-case, lower-case and numbers
# as chars in password
def gen_pass(upper:bool = True, lower:bool= True, nums:bool = True, size:int = 15) -> str:
    
    chars = "A B C D E F G H I K L M N O P Q R S T V X Y Z"

    char_pool = []
    password = ""
    if upper:
        [char_pool.append(x) for x in list(chars) if x.isalnum()]
    if lower:
        [char_pool.append(x.lower()) for x in list(chars) if x.isalnum()]
    if nums:
        [char_pool.append(str(x)) for x in range(0,10)]

    while len(password) < size:
        password += char_pool[random.randint(0, len(char_pool) -1 )]

    return password

pw = gen_pass(True, True, True, 15)
print(f" {pw} {len(pw)}")

pw_list = [gen_pass(True, True, True, 15) for _ in range(1000)]
print(f" {pw_list} {len(pw_list)}")
