#[allow(unused_imports)]
use crate::days::day6::part1_solution;

#[test]
fn example() {
    let inputs = "....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#..."
        .split("\n")
        .map(|temp: &str| -> String { temp.to_string() })
        .collect();
    let result = part1_solution::solve(&inputs);
    assert_eq!(result, "41")
}
