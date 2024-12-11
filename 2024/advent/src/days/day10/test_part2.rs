#[allow(unused_imports)]
use crate::days::day10::part2_solution;

#[test]
fn example() {
    let inputs = "..90..9
...1.98
...2..7
6543456
765.987
876....
987...."
        .split("\n")
        .map(|temp: &str| -> String { temp.to_string() })
        .collect();
    let result = part2_solution::solve(&inputs);
    assert_eq!(result, "13")
}

#[test]
fn example2() {
    let inputs = "012345
123456
234567
345678
4.6789
56789."
        .split("\n")
        .map(|temp: &str| -> String { temp.to_string() })
        .collect();
    let result = part2_solution::solve(&inputs);
    assert_eq!(result, "227")
}

#[test]
fn example3() {
    let inputs = "89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732"
        .split("\n")
        .map(|temp: &str| -> String { temp.to_string() })
        .collect();
    let result = part2_solution::solve(&inputs);
    assert_eq!(result, "81")
}
