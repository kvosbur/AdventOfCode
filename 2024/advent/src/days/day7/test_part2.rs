#[allow(unused_imports)]
use crate::days::day7::part2_solution;

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
    let result = part2_solution::solve(&inputs);
    assert_eq!(result, "6")
}
