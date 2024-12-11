#[allow(unused_imports)]
use crate::days::day11::part1_solution;

#[test]
fn example1() {
    let inputs = "125 17"
        .split("\n")
        .map(|temp: &str| -> String { temp.to_string() })
        .collect();
    let result = part1_solution::solve(&inputs);
    assert_eq!(result, "55312")
}
