use core::str;
use std::collections::HashMap;

fn get_beginning_state(characters: &Vec<&[u8]>) -> (usize, usize, u32) {
    for row_index in 0..characters.len() {
        let row = characters[row_index];
        for value_index in 0..row.len() {
            let value = row[value_index];
            if value != b'.' && value != b'#' {
                let direction = match value {
                    b'^' => 0,
                    b'>' => 1,
                    b'v' => 2,
                    b'<' => 3,
                    _other => panic!("Incorrect starting location"),
                };
                return (row_index, value_index, direction);
            }
        }
    }
    return (0, 0, 0);
}

fn get_key(first_index: usize, second_index: usize) -> String {
    let mut key = first_index.to_string();
    key.push_str(",");
    key.push_str(&second_index.to_string());
    key
}

fn increment_count(counts: &mut HashMap<String, u32>, key: String) {
    counts
        .entry(key)
        .and_modify(|entry| {
            *entry += 1;
        })
        .or_insert(1);
}

fn do_movements(
    characters: &Vec<&[u8]>,
    counts: &mut HashMap<String, u32>,
    initial_state: (usize, usize, u32),
) {
    let mut x = initial_state.0;
    let mut y = initial_state.1;
    let mut direction = initial_state.2;

    let mut moving = true;
    while moving {
        // println!(
        //     "direction:{}  x:{}, y:{}, char: {}",
        //     direction,
        //     x,
        //     y,
        //     str::from_utf8(&[characters[x][y]]).unwrap()
        // );
        match direction {
            0 => {
                // up
                if x == 0 {
                    moving = false;
                    break;
                }
                let next_x = x - 1;
                let next_y = y;
                if characters[next_x][next_y] == b'#' {
                    direction = 1;
                } else {
                    increment_count(counts, get_key(x, y));
                    x = next_x;
                    y = next_y;
                }
            }
            1 => {
                // right
                if y == characters[0].len() - 1 {
                    moving = false;
                    break;
                }
                let next_x = x;
                let next_y = y + 1;
                if characters[next_x][next_y] == b'#' {
                    direction = 2;
                } else {
                    increment_count(counts, get_key(x, y));
                    x = next_x;
                    y = next_y;
                }
            }
            2 => {
                // down
                if x == characters.len() - 1 {
                    moving = false;
                    break;
                }
                let next_x = x + 1;
                let next_y = y;
                if characters[next_x][next_y] == b'#' {
                    direction = 3;
                } else {
                    increment_count(counts, get_key(x, y));
                    x = next_x;
                    y = next_y;
                }
            }
            3 => {
                // left
                if y == 0 {
                    moving = false;
                    break;
                }
                let next_x = x;
                let next_y = y - 1;
                if characters[next_x][next_y] == b'#' {
                    direction = 0;
                } else {
                    increment_count(counts, get_key(x, y));
                    x = next_x;
                    y = next_y;
                }
            }
            _other => panic!("unknown direction"),
        };
    }
    increment_count(counts, get_key(x, y));
}

#[allow(dead_code)]
pub fn solve(inputs: &Vec<String>) -> String {
    let mut location_counts: HashMap<String, u32> = HashMap::new();

    let input_chars: Vec<&[u8]> = inputs.iter().map(|s| s.as_bytes()).collect();

    let begin_state = get_beginning_state(&input_chars);
    // insert start location
    do_movements(&input_chars, &mut location_counts, begin_state);

    let spot_counts = location_counts.keys().count();

    spot_counts.to_string()
}
