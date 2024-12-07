use std::collections::HashMap;

fn get_beginning_state(characters: &Vec<Vec<u8>>) -> (usize, usize, u32) {
    for row_index in 0..characters.len() {
        let row = &characters[row_index];
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

fn add_seen_location(counts: &mut HashMap<String, Vec<u32>>, key: String, direction: u32) {
    counts
        .entry(key)
        .and_modify(|entry| entry.push(direction))
        .or_insert(vec![direction]);
}

fn does_movements_cause_cycle(
    characters: &Vec<Vec<u8>>,
    initial_state: (usize, usize, u32),
) -> bool {
    let mut x = initial_state.0;
    let mut y = initial_state.1;
    let mut direction = initial_state.2;
    let mut seen_locations: HashMap<String, Vec<u32>> = HashMap::new();

    let mut moving = true;
    while moving {
        let key = get_key(x, y);
        let has_seen_location = seen_locations.get(&key);
        if has_seen_location.is_some() {
            if has_seen_location
                .unwrap()
                .iter()
                .any(|dir| *dir == direction)
            {
                return true;
            }
        }
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
                    add_seen_location(&mut seen_locations, get_key(x, y), direction);
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
                    add_seen_location(&mut seen_locations, get_key(x, y), direction);
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
                    add_seen_location(&mut seen_locations, get_key(x, y), direction);
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
                    add_seen_location(&mut seen_locations, get_key(x, y), direction);
                    x = next_x;
                    y = next_y;
                }
            }
            _other => panic!("unknown direction"),
        };
    }
    return false;
}

#[allow(dead_code)]
pub fn solve(inputs: &Vec<String>) -> String {
    let mut input_chars: Vec<Vec<u8>> = inputs.iter().map(|s| s.clone().into_bytes()).collect();

    let begin_state = get_beginning_state(&input_chars);
    // insert start location
    let mut good_locations = 0;
    for x in 0..input_chars.len() {
        println!("x:{} total: {}", x, input_chars.len());
        for y in 0..input_chars[0].len() {
            if input_chars[x][y] == b'.' {
                input_chars[x][y] = b'#';
                if does_movements_cause_cycle(&input_chars, begin_state) {
                    good_locations += 1;
                }
                input_chars[x][y] = b'.';
            }
        }
    }

    good_locations.to_string()
}
