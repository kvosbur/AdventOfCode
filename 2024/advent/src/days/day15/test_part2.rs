#[allow(unused_imports)]
use crate::days::day15::part2_solution;

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

#[test]
fn example3() {
    let inputs = "EEEEE
EXXXX
EEEEE
EXXXX
EEEEE"
        .split("\n")
        .map(|temp: &str| -> String { temp.to_string() })
        .collect();
    let result = part2_solution::solve(&inputs);
    assert_eq!(result, "236")
}

#[test]
fn example4() {
    let inputs = "AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA"
        .split("\n")
        .map(|temp: &str| -> String { temp.to_string() })
        .collect();
    let result = part2_solution::solve(&inputs);
    assert_eq!(result, "368")
}

#[test]
fn example5() {
    let inputs = "RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE"
        .split("\n")
        .map(|temp: &str| -> String { temp.to_string() })
        .collect();
    let result = part2_solution::solve(&inputs);
    assert_eq!(result, "1206")
}
