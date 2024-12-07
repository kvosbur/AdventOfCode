use std::collections::HashSet;
use std::sync::mpsc::{self, Sender};
use std::thread;

fn get_beginning_state(characters: &Vec<Vec<u8>>) -> (usize, usize, usize) {
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

fn get_direction_key(first_index: usize, second_index: usize, direction: u32) -> String {
    let mut key = first_index.to_string();
    key.push_str(",");
    key.push_str(&second_index.to_string());
    key.push_str(",");
    key.push_str(&direction.to_string());
    key
}

fn does_movements_cause_cycle(
    characters: &Vec<Vec<u8>>,
    initial_state: &(usize, usize, usize),
) -> bool {
    let mut x = initial_state.0;
    let mut y = initial_state.1;
    let mut direction: usize = initial_state.2;
    let mut seen_locations: HashSet<usize> = HashSet::new();
    let double_len = characters.len() * 10;

    let mut moving = true;
    while moving {
        // let key = get_direction_key(x, y, direction);
        let key = (x * double_len) + (y * 10) + direction;
        if seen_locations.get(&key).is_some() {
            return true;
        }
        match direction {
            0 => {
                // up
                if x == 0 {
                    moving = false;
                    break;
                }
                let next_x = x - 1;

                if characters[next_x][y] == b'#' {
                    direction = 1;
                } else {
                    seen_locations.insert(key);
                    x = next_x;
                }
            }
            1 => {
                // right
                if y == characters[0].len() - 1 {
                    moving = false;
                    break;
                }
                let next_y = y + 1;
                if characters[x][next_y] == b'#' {
                    direction = 2;
                } else {
                    seen_locations.insert(key);
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
                if characters[next_x][y] == b'#' {
                    direction = 3;
                } else {
                    seen_locations.insert(key);
                    x = next_x;
                }
            }
            3 => {
                // left
                if y == 0 {
                    moving = false;
                    break;
                }
                let next_y = y - 1;
                if characters[x][next_y] == b'#' {
                    direction = 0;
                } else {
                    seen_locations.insert(key);
                    y = next_y;
                }
            }
            _other => panic!("unknown direction"),
        };
    }
    return false;
}

fn thread_work(
    start_x: usize,
    end_x: usize,
    mut input_chars: Vec<Vec<u8>>,
    begin_state: &(usize, usize, usize),
    tx: Sender<i32>,
) {
    let mut good_locations = 0;
    for x in start_x..end_x {
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
    tx.send(good_locations).unwrap();
}

#[allow(dead_code)]
pub fn solve(inputs: &Vec<String>) -> String {
    let input_chars: Vec<Vec<u8>> = inputs.iter().map(|s| s.clone().into_bytes()).collect();
    let thread_count = 8;
    let begin_state = get_beginning_state(&input_chars);
    let (tx, rx) = mpsc::channel();

    let x_step = input_chars.len() / thread_count;
    for t in 0..thread_count {
        let begin_x = t * x_step;
        let mut end_x = (t + 1) * x_step;
        if t == thread_count - 1 {
            end_x = input_chars.len();
        }
        println!("begin:{} end:{}", begin_x, end_x);
        let cloned_input_chars: Vec<Vec<u8>> = input_chars.iter().map(|v| v.clone()).collect();
        let cloned_tx = tx.clone();
        thread::spawn(move || {
            thread_work(begin_x, end_x, cloned_input_chars, &begin_state, cloned_tx)
        });
    }
    drop(tx);

    let mut good_locations = 0;
    for received in rx {
        good_locations += received;
        println!("received {}", received);
    }
    return good_locations.to_string();
}
