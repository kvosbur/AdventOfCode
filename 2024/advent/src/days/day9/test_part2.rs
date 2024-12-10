#[allow(unused_imports)]
use crate::days::day9::part2_solution;

#[test]
fn example() {
    let inputs = "2333133121414131402"
        .split("\n")
        .map(|temp: &str| -> String { temp.to_string() })
        .collect();
    let result = part2_solution::solve(&inputs);
    assert_eq!(result, "2858")
}

#[test]
fn example2() {
    let inputs = "12345"
        .split("\n")
        .map(|temp: &str| -> String { temp.to_string() })
        .collect();
    let result = part2_solution::solve(&inputs);
    assert_eq!(result, "132")
}
