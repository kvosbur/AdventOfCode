#[allow(unused_imports)]
use crate::days::day3::part2_solution;

#[test]
fn example() {
    let inputs = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))
    xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
        .split("\n")
        .map(|temp: &str| -> String { temp.to_string() })
        .collect();
    let result = part2_solution::solve(&inputs);
    assert_eq!(result, "96")
}

#[test]
fn break_the_thing() {
    let inputs =
        "xmul(2,4)& do() mul(2,4)mul(2,4)mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(2,4)mul(2,4)don't()mul(2,4)do()mul(2,4)do()mul(2,4)"
            .split("\n")
            .map(|temp: &str| -> String { temp.to_string() })
            .collect();
    let result = part2_solution::solve(&inputs);
    assert_eq!(result, "56")
}
