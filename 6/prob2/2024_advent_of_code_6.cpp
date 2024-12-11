#include <fstream>
#include <iostream>
#include <string>
#include <vector>

using namespace std;

// 0 = 00000000
// 1 = 00000001
// (1 << 1) = 00000010 = 2
// (1 << 2) = 00000100 = 4
// 13 = 00001101
// (13 << 1) = 00011010 = 26

// Directions:
// 0 = up       (1 << 0) = 1
// 1 = right    (1 << 1) = 2
// 2 = down     (1 << 2) = 4
// 3 = left     (1 << 3) = 8

// 13       = 1101
// (1 << 0) = 0001
// 13 & (1 << 0) = 0001 != 0

// 13       = 1101
// (1 << 1) = 0010
// 13 & (1 << 1) = 0000 == 0

// 13       = 1101
// (1 << 2) = 0100
// 13 & (1 << 2) = 0100 != 0

bool has_visited(char visited, int direction) {
  return visited & (1 << direction);
}

void set_visited(char &visited, int direction) { visited |= (1 << direction); }

bool process(const vector<vector<char>> &data, int start_i, int start_j,
             int obstacle_i, int obstacle_j) {
  int N_i = data.size();
  int N_j = data[0].size();
  vector<vector<char>> visited(N_i, vector<char>(N_j, 0));

  int pos_i = start_i;
  int pos_j = start_j;
  int direction = 0;
  while (true) {
    set_visited(visited[pos_i][pos_j], direction);
    int next_pos_i = pos_i;
    int next_pos_j = pos_j;
    switch (direction) {
    case 0:
      next_pos_i--;
      break;
    case 1:
      next_pos_j++;
      break;
    case 2:
      next_pos_i++;
      break;
    case 3:
      next_pos_j--;
      break;
    }
    if (next_pos_i == -1 || next_pos_i == N_i || next_pos_j == -1 ||
        next_pos_j == N_j) {
      return false;
    }
    if (data[next_pos_i][next_pos_j] == '#' ||
        (next_pos_i == obstacle_i && next_pos_j == obstacle_j)) {
      direction = (direction + 1) % 4;
      continue;
    }
    if (has_visited(visited[next_pos_i][next_pos_j], direction)) {
      return true;
    }
    pos_i = next_pos_i;
    pos_j = next_pos_j;
  }
}

int main() {
  ifstream file("input.txt");
  if (!file.is_open()) {
    cerr << "Error: Could not open file." << endl;
    return 1;
  }

  vector<vector<char>> grid;
  string line;

  while (getline(file, line)) {
    vector<char> row;
    for (char ch : line) {
      row.push_back(ch);
    }
    grid.push_back(row);
  }

  file.close();

  int start_i;
  int start_j;
  for (int i = 0; i < grid.size(); i++) {
    for (int j = 0; j < grid[i].size(); j++) {
      if (grid[i][j] == '^') {
        start_i = i;
        start_j = j;
      }
    }
  }

  int answer = 0;
  for (int i = 0; i < grid.size(); i++) {
    for (int j = 0; j < grid[i].size(); j++) {
      if (i == start_i && j == start_j) {
        continue;
      }

      if (process(grid, start_i, start_j, i, j)) {
        answer++;
      }
    }
  }

  std::cout << "The answer is " << answer << std::endl;
  return 0;
}
