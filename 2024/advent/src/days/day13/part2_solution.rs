#[derive(Debug)]
struct Game {
    a_button: (f64, f64),
    b_button: (f64, f64),
    prize: (f64, f64),
}

fn parse_button_line(line: &String) -> (f64, f64) {
    let start_x = line.find("X").unwrap();
    let separator = line.find(", ").unwrap();
    let start_y = line.find("Y").unwrap();
    let x: f64 = line[start_x + 2..separator].parse().unwrap();
    let y: f64 = line[start_y + 2..].parse().unwrap();
    return (x, y);
}

fn parse_games(inputs: &Vec<String>) -> Vec<Game> {
    let mut games = vec![];

    for begin_index in (0..inputs.len()).step_by(4) {
        games.push(Game {
            a_button: parse_button_line(&inputs[begin_index]),
            b_button: parse_button_line(&inputs[begin_index + 1]),
            prize: parse_button_line(&inputs[begin_index + 2]),
        })
    }

    return games;
}

fn mutate_games(games: &mut Vec<Game>) {
    for game in games {
        game.prize.0 += 10000000000000.0;
        game.prize.1 += 10000000000000.0;
    }
}

fn get_game_tokens(game: &Game) -> Option<i64> {
    let determinant: f64 = game.a_button.0 * game.b_button.1 - game.b_button.0 * game.a_button.1;
    let a_button_presses =
        (game.prize.0 * game.b_button.1 - game.b_button.0 * game.prize.1) / determinant;
    let b_button_presses =
        (game.a_button.0 * game.prize.1 - game.prize.0 * game.a_button.1) / determinant;

    let a_presses_int = a_button_presses as i64;
    let b_presses_int = b_button_presses as i64;
    if b_button_presses - (b_presses_int as f64) < 0.001
        && a_button_presses - (a_presses_int as f64) < 0.001
    {
        return Some(a_presses_int * 3 + b_presses_int);
    }
    return None;
}

#[allow(dead_code)]
pub fn solve(inputs: &Vec<String>) -> String {
    let mut games = parse_games(inputs);
    mutate_games(&mut games);
    let mut tokens = 0;
    for game in &games {
        let maybe_tokens = get_game_tokens(game);
        if maybe_tokens.is_some() {
            tokens += maybe_tokens.unwrap();
        }
    }

    tokens.to_string()
}
