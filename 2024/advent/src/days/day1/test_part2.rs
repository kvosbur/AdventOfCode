use crate::days::day1::part2_solution;

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
    let result = part2_solution::solve(&inputs);
    assert_eq!(result, "31")
}
