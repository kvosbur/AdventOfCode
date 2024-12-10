#[allow(unused_imports)]
use crate::days::day9::part1_solution;

#[test]
fn example() {
    let inputs = "2333133121414131402"
        .split("\n")
        .map(|temp: &str| -> String { temp.to_string() })
        .collect();
    let result = part1_solution::solve(&inputs);
    assert_eq!(result, "1928")
}
