#[allow(unused_imports)]
use crate::days::day1::part1_solution;

#[test]
fn example() {
    let inputs = "3   4
4   3
2   5
1   3
3   9
3   3"
        .split("\n")
        .map(|temp: &str| -> String { temp.to_string() })
        .collect();
    let result = part1_solution::solve(&inputs);
    assert_eq!(result, "11")
}
