#[allow(unused_imports)]
use crate::days::day12::part2_solution;

#[test]
fn example1() {
    let inputs = "AAAA
BBCD
BBCC
EEEC"
        .split("\n")
        .map(|temp: &str| -> String { temp.to_string() })
        .collect();
    let result = part2_solution::solve(&inputs);
    assert_eq!(result, "80")
}

#[test]
fn example2() {
    let inputs = "OOOOO
OXOXO
OOOOO
OXOXO
OOOOO"
        .split("\n")
        .map(|temp: &str| -> String { temp.to_string() })
        .collect();
    let result = part2_solution::solve(&inputs);
    assert_eq!(result, "436")
}
