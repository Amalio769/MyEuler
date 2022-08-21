"""
Starting in the top left corner of a 2×2 grid, and only being able to move to the right and down, there are exactly 6
routes to the bottom right corner.


How many such routes are there through a 20×20 grid?
Answer:  137846528820
Completed on Sun, 21 Aug 2022, 18:36
"""

def lattice_paths_right_down(columns: int, files: int):
    lattice_paths = [[x for x in range(0, columns+1)] for y in range(0, files+1)]
    lattice_paths[0][0] = 1

    for file in range(files+1):
        for column in range(columns+1):
            if file == 0 and column == 0:
                continue
            if file-1 < 0:
                father_file = 0
            else:
                father_file = lattice_paths[file-1][column]

            if column-1 < 0:
                father_column = 0
            else:
                father_column = lattice_paths[file][column-1]

            lattice_paths[file][column] = father_file + father_column

    return lattice_paths

print(lattice_paths_right_down(20, 20))
print(max(max(lattice_paths_right_down(20,20))))