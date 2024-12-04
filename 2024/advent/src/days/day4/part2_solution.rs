use std::collections::HashMap;

fn searcher(
    values: &Vec<String>,
    goal: &str,
    x_start: usize,
    x_direction: i32,
    y_start: usize,
    y_direction: i32,
) -> (bool, Option<(usize, usize)>) {
    let max_x: i32 = values.len().try_into().unwrap();
    let max_y: i32 = values[x_start].len().try_into().unwrap();
    let mut x: i32 = x_start.try_into().unwrap();
    let mut y: i32 = y_start.try_into().unwrap();
    let mut a_location: (usize, usize) = (0, 0);
    for c in goal.chars() {
        if x < 0 || x == max_x || y < 0 || y == max_y {
            return (false, None);
        }
        let other_x: usize = x.try_into().unwrap();
        let other_y: usize = y.try_into().unwrap();

        if values[other_x].chars().nth(other_y).unwrap() != c {
            return (false, None);
        }

        if c == 'A' {
            a_location = (other_x, other_y);
        }

        x += x_direction;
        y += y_direction;
    }
    return (true, Some(a_location));
}

#[allow(dead_code)]
pub fn solve(inputs: &Vec<String>) -> String {
    let goal = "MAS";

    let mut x: usize = 0;
    let max_x: usize = inputs.len();
    let max_y: usize = inputs[0].len();

    let mut cross_locations: HashMap<String, i32> = HashMap::new();
    while x < max_x {
        let mut y: usize = 0;
        while y < max_y {
            if inputs[x].chars().nth(y).unwrap() != 'M' {
                y += 1;
                continue;
            }

            let results = vec![
                searcher(inputs, goal, x, -1, y, 1),  // upper left
                searcher(inputs, goal, x, 1, y, 1),   // upper right
                searcher(inputs, goal, x, 1, y, -1),  // lower right
                searcher(inputs, goal, x, -1, y, -1), // lower left
            ];
            results.iter().for_each(|res| {
                if res.0 {
                    let a_location = res.1.unwrap();
                    let mut key = a_location.0.to_string();
                    key.push_str(",");
                    key.push_str(&a_location.1.to_string());

                    cross_locations
                        .entry(key)
                        .and_modify(|count| *count += 1)
                        .or_insert(1);
                }
            });

            y += 1;
        }
        x += 1;
    }
    let mut found_count = 0;

    for (_, value) in &cross_locations {
        if *value > 1 {
            found_count += 1;
        }
    }

    found_count.to_string()
}
