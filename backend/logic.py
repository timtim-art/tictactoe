import math
import time
import random

# Create the single objects and a list of these objects

canvas_y = 30
canvas_x = 10
num_of_participants = 9
types = ['scissor', 'stone', 'paper']
num = int(num_of_participants / len(types))


class Warrior:
    def __init__(self, option):
        self.position_x = random.randint(0, canvas_x)
        self.position_y = random.randint(0, canvas_y)
        self.type = option

    def __repr__(self):
        return f"Warrior({self.type}, x={self.position_x}, y={self.position_y})"


warrior_list = []

for type in range(len(types)):
    for index in range(num):
        warrior = Warrior(types[type])
        warrior_list.append(warrior)


# Methods to move the warriors depending on their next victims or threats

def chase(opponent_to_find, warrior):
    # If warriors meet, change opponent's type to the warrior's type
    if opponent_to_find.position_y == warrior.position_y and opponent_to_find.position_x == warrior.position_x:
        opponent_to_find.type = warrior.type
        return

    # Move warrior towards the opponent
    if opponent_to_find.position_y < warrior.position_y and warrior.position_y > 0:
        warrior.position_y -= 1
    elif opponent_to_find.position_y > warrior.position_y and warrior.position_y < canvas_y:
        warrior.position_y += 1
    if opponent_to_find.position_x < warrior.position_x and warrior.position_x > 0:
        warrior.position_x -= 1
    elif opponent_to_find.position_x > warrior.position_x and warrior.position_x < canvas_x:
        warrior.position_x += 1


def run(opponent_to_find, warrior):
    # If warriors meet, change warrior's type to the opponent's type
    if opponent_to_find.position_y == warrior.position_y and opponent_to_find.position_x == warrior.position_x:
        warrior.type = opponent_to_find.type
        return

    # Move warrior away from the opponent
    if opponent_to_find.position_y < warrior.position_y and warrior.position_y < canvas_y:
        warrior.position_y += 1
    elif opponent_to_find.position_y > warrior.position_y and warrior.position_y > 0:
        warrior.position_y -= 1
    if opponent_to_find.position_x < warrior.position_x and warrior.position_x < canvas_x:
        warrior.position_x += 1
    elif opponent_to_find.position_x > warrior.position_x and warrior.position_x > 0:
        warrior.position_x -= 1


# Iterate through the list as long as there is more than one type

while len(set(warrior.type for warrior in warrior_list)) > 1:  # Continue until only one type is left
    for warrior in warrior_list:
        distance = None
        opponent_to_find = None
        for opponent in warrior_list:
            # Find distance to opponents
            if opponent.type != warrior.type:
                new_distance = int(math.sqrt(
                    (opponent.position_y - warrior.position_y) ** 2 + (opponent.position_x - warrior.position_x) ** 2))
                # Find the opponent with the smallest distance
                if distance is None or new_distance < distance:
                    distance = new_distance
                    opponent_to_find = opponent

        # Ensure opponent_to_find is valid
        if opponent_to_find:
            if warrior.type == 'scissor' and opponent_to_find.type == 'stone':
                run(opponent_to_find, warrior)
            elif warrior.type == 'scissor' and opponent_to_find.type == 'paper':
                chase(opponent_to_find, warrior)
            elif warrior.type == 'stone' and opponent_to_find.type == 'scissor':
                chase(opponent_to_find, warrior)
            elif warrior.type == 'stone' and opponent_to_find.type == 'paper':
                run(opponent_to_find, warrior)
            elif warrior.type == 'paper' and opponent_to_find.type == 'stone':
                chase(opponent_to_find, warrior)
            elif warrior.type == 'paper' and opponent_to_find.type == 'scissor':
                run(opponent_to_find, warrior)

    # Adding a delay to slow down the loop for observation
    print(warrior_list)
    time.sleep(0.1)

# Print the final state
final_type = warrior_list[0].type
print(f"The game ends with all warriors being of type: {final_type}")
