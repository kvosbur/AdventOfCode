#[allow(unused_imports)]
use crate::days::day2::part2_solution;

#[test]
fn example() {
    let inputs = "7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9"
        .split("\n")
        .map(|temp: &str| -> String { temp.to_string() })
        .collect();
    let result = part2_solution::solve(&inputs);
    assert_eq!(result, "4")
}

#[test]
fn special_case_1() {
    let inputs = "8 2 3 5 6"
        .split("\n")
        .map(|temp: &str| -> String { temp.to_string() })
        .collect();
    let result = part2_solution::solve(&inputs);
    assert_eq!(result, "1")
}
