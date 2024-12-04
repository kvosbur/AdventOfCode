fn searcher(
    values: &Vec<String>,
    goal: &str,
    x_start: usize,
    x_direction: i32,
    y_start: usize,
    y_direction: i32,
) -> bool {
    let max_x: i32 = values.len().try_into().unwrap();
    let max_y: i32 = values[x_start].len().try_into().unwrap();
    let mut x: i32 = x_start.try_into().unwrap();
    let mut y: i32 = y_start.try_into().unwrap();
    for c in goal.chars() {
        if x < 0 || x == max_x || y < 0 || y == max_y {
            return false;
        }
        let other_x: usize = x.try_into().unwrap();
        let other_y: usize = y.try_into().unwrap();

        if values[other_x].chars().nth(other_y).unwrap() != c {
            return false;
        }

        x += x_direction;
        y += y_direction;
    }
    return true;
}

#[allow(dead_code)]
pub fn solve(inputs: &Vec<String>) -> String {
    let goal = "XMAS";

    let mut x: usize = 0;
    let max_x: usize = inputs.len();
    let max_y: usize = inputs[0].len();

    let mut found_count = 0;
    while x < max_x {
        let mut y: usize = 0;
        while y < max_y {
            if inputs[x].chars().nth(y).unwrap() != 'X' {
                y += 1;
                continue;
            }

            let results = vec![
                searcher(inputs, goal, x, -1, y, 0),  // left
                searcher(inputs, goal, x, -1, y, 1),  // upper left
                searcher(inputs, goal, x, 0, y, 1),   // up
                searcher(inputs, goal, x, 1, y, 1),   // upper right
                searcher(inputs, goal, x, 1, y, 0),   // right
                searcher(inputs, goal, x, 1, y, -1),  // lower right
                searcher(inputs, goal, x, 0, y, -1),  // down
                searcher(inputs, goal, x, -1, y, -1), // lower left
            ];
            found_count += results
                .iter()
                .map(|res| if *res { 1 } else { 0 })
                .sum::<i32>();

            y += 1;
        }
        x += 1;
    }

    found_count.to_string()
}
