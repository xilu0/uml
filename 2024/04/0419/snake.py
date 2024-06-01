import pygame
import sys
import random

# Initialize Pygame
pygame.init()

# Set the screen dimensions
screen_width = 640
screen_height = 480
screen = pygame.display.set_mode((screen_width, screen_height))

# Set the title of the window
pygame.display.set_caption("Snake Game")

# Game constants
SPEED = 10
BLOCK_SIZE = 20
GRID_SIZE = (screen_width // BLOCK_SIZE, screen_height // BLOCK_SIZE)

# Game variables
snake = [(0, 0)]  # initial snake position
direction = (1, 0)  # initial direction (right)
food = None  # initial food position
score = 0  # initial score

def draw_snake(snake):
    for x, y in snake:
        pygame.draw.rect(screen, (0, 255, 0), (x * BLOCK_SIZE, y * BLOCK_SIZE, BLOCK_SIZE, BLOCK_SIZE))

def draw_food(food):
    pygame.draw.rect(screen, (255, 0, 0), (food[0] * BLOCK_SIZE, food[1] * BLOCK_SIZE, BLOCK_SIZE, BLOCK_SIZE))

def move_snake(snake, direction):
    new_head = (snake[0][0] + direction[0], snake[0][1] + direction[1])
    snake.insert(0, new_head)
    return snake

def check_collision(snake, food):
    if snake[0] == food:
        return True
    return False

def check_boundaries(snake):
    if (snake[0][0] < 0 or snake[0][0] >= GRID_SIZE[0] or
        snake[0][1] < 0 or snake[0][1] >= GRID_SIZE[1]):
        return True
    return False


while True:
    # Handle events
    for event in pygame.event.get():
        if event.type == pygame.QUIT:
            pygame.quit()
            sys.exit()
        elif event.type == pygame.KEYDOWN:
            if event.key == pygame.K_UP and direction!= (0, 1):
                direction = (0, -1)
            elif event.key == pygame.K_DOWN and direction!= (0, -1):
                direction = (0, 1)
            elif event.key == pygame.K_LEFT and direction!= (1, 0):
                direction = (-1, 0)
            elif event.key == pygame.K_RIGHT and direction!= (-1, 0):
                direction = (1, 0)

    # Move snake
    snake = move_snake(snake, direction)

    # Check collision with food
    if food is None or check_collision(snake, food):
        food = (random.randint(0, GRID_SIZE[0] - 1), random.randint(0, GRID_SIZE[1] - 1))
        score += 1

    # Check collision with boundaries
    if check_boundaries(snake):
        print("Game Over! Score:", score)
        pygame.quit()
        sys.exit()

    # Draw everything
    screen.fill((0, 0, 0))  # clear screen
    draw_snake(snake)
    draw_food(food)
    pygame.display.flip()  # update screen
    pygame.time.Clock().tick(SPEED)  # control frame rate