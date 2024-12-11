#[allow(unused_imports)]
use crate::days::day10::part1_solution;

#[test]
fn example1() {
    let inputs = "...0...
...1...
...2...
6543456
7.....7
8.....8
9.....9"
        .split("\n")
        .map(|temp: &str| -> String { temp.to_string() })
        .collect();
    let result = part1_solution::solve(&inputs);
    assert_eq!(result, "2")
}

#[test]
fn example2() {
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
    let result = part1_solution::solve(&inputs);
    assert_eq!(result, "4")
}

#[test]
fn example3() {
    let inputs = "10..9..
2...8..
3...7..
4567654
...8..3
...9..2
.....01"
        .split("\n")
        .map(|temp: &str| -> String { temp.to_string() })
        .collect();
    let result = part1_solution::solve(&inputs);
    assert_eq!(result, "3")
}

#[test]
fn example4() {
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
    let result = part1_solution::solve(&inputs);
    assert_eq!(result, "36")
}
