fn get_start_location(input_chars: &Vec<Vec<char>>) -> (usize, usize) {
    for x in 0..input_chars.len() {
        for y in 0..input_chars[0].len() {
            if input_chars[x][y] == '@' {
                return (x, y);
            }
        }
    }
    panic!("uhoh, couldn't find start location");
}

fn get_move_directions(direction: &u8) -> (i32, i32) {
    match direction {
        b'>' => (0, 1),
        b'^' => (-1, 0),
        b'v' => (1, 0),
        b'<' => (0, -1),
        _ => panic!("got invalid direction"),
    }
}

fn do_num_conversion(input: &(i32, i32)) -> (usize, usize) {
    (input.0.try_into().unwrap(), input.1.try_into().unwrap())
}

fn move_characters(board: &mut Vec<Vec<char>>, positions: &Vec<(i32, i32)>) {
    for pos_index in (1..positions.len()).rev() {
        let prev_position = do_num_conversion(&positions[pos_index - 1]);
        let position = do_num_conversion(&positions[pos_index]);

        board[position.0][position.1] = board[prev_position.0][prev_position.1];
    }
    let first_pos = do_num_conversion(positions.first().unwrap());
    board[first_pos.0][first_pos.1] = '.';
}

fn move_bot(
    board: &mut Vec<Vec<char>>,
    direction: &u8,
    start_location: (usize, usize),
) -> (usize, usize) {
    let move_direction = get_move_directions(direction);
    let mut positions: Vec<(i32, i32)> = vec![(
        start_location.0.try_into().unwrap(),
        start_location.1.try_into().unwrap(),
    )];
    let mut last_char: char = '@';

    while last_char == '@' || last_char == 'O' {
        let prev_position = positions.last().unwrap();
        let next_position: (i32, i32) = (
            prev_position.0 + move_direction.0,
            prev_position.1 + move_direction.1,
        );
        positions.push(next_position);
        last_char = board[TryInto::<usize>::try_into(next_position.0).unwrap()]
            [TryInto::<usize>::try_into(next_position.1).unwrap()]
    }

    let last_position = positions.last().unwrap();
    if board[TryInto::<usize>::try_into(last_position.0).unwrap()]
        [TryInto::<usize>::try_into(last_position.1).unwrap()]
        == '.'
    {
        move_characters(board, &positions);
        return (
            positions[1].0.try_into().unwrap(),
            positions[1].1.try_into().unwrap(),
        );
    } else {
        return start_location;
    }
}

fn score_board(board: &Vec<Vec<char>>) -> usize {
    let mut score: usize = 0;
    for x in 0..board.len() {
        for y in 0..board[0].len() {
            if board[x][y] == 'O' {
                score += x * 100 + y;
            }
        }
    }
    score
}

#[allow(dead_code)]
fn print_board(board: &Vec<Vec<char>>) {
    for row in board {
        for c in row {
            print!("{}", c);
        }
        println!("");
    }
}

#[allow(dead_code)]
pub fn solve(inputs: &Vec<String>) -> String {
    let input_chars: Vec<&[u8]> = inputs.iter().map(|s| s.as_bytes()).collect();

    let input_split_index = input_chars.iter().position(|r| r.len() == 0).unwrap();

    let (board, movements) = input_chars.split_at(input_split_index);

    let mut mutable_board: Vec<Vec<char>> = board
        .iter()
        .map(|row| row.iter().map(|item| item.clone() as char).collect())
        .collect();

    let mut start_location = get_start_location(&mutable_board);

    for movement_line in movements {
        for movement in *movement_line {
            start_location = move_bot(&mut mutable_board, movement, start_location);
        }
    }

    let board_score = score_board(&mutable_board);
    board_score.to_string()
}
