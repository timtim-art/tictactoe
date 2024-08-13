# Example file showing a circle moving on screen
import pygame
from network import Network

def read_position(str):
    str = str.split(',')
    return int(str[0]), int(str[1])

def make_position(tup):
    return str(tup[0]) + ',' + str(tup[1])

# pygame setup
pygame.init()
screen = pygame.display.set_mode((1280, 720))
clock = pygame.time.Clock()
running = True
# Here we are connecting to the server
n = Network()
# Returns starting position from server for the circle the client is playing with
start_position = read_position(n.get_position())
dt = 0





player_pos = pygame.Vector2(start_position[0], start_position[1])

player2_pos = pygame.Vector2(500, 500)


while running:

    # fill the screen with a color to wipe away anything from last frame
    screen.fill("purple")

    p = pygame.draw.circle(screen, (255, 255, 255), player_pos, 40)
    p2 = pygame.draw.circle(screen, (0, 255, 255), player2_pos, 40)

    p2_position = read_position(n.send(make_position((int(player_pos[0]), int(player_pos[1])))))
    player2_pos = pygame.Vector2(p2_position[0], p2_position[1])

    # poll for events
    # pygame.QUIT event means the user clicked X to close your window
    for event in pygame.event.get():
        if event.type == pygame.QUIT:
            running = False





    keys = pygame.key.get_pressed()
    if keys[pygame.K_UP]:
        player_pos.y -= 300 * dt
    if keys[pygame.K_DOWN]:
        player_pos.y += 300 * dt
    if keys[pygame.K_LEFT]:
        player_pos.x -= 300 * dt
    if keys[pygame.K_RIGHT]:
        player_pos.x += 300 * dt

    # flip() the display to put your work on screen
    pygame.display.flip()

    # limits FPS to 60
    # dt is delta time in seconds since last frame, used for framerate-
    # independent physics.
    dt = clock.tick(60) / 1000

pygame.quit()