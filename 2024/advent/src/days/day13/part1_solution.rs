#[derive(Debug)]
struct Game {
    a_button: (i64, i64),
    b_button: (i64, i64),
    prize: (i64, i64),
}

fn parse_button_line(line: &String) -> (i64, i64) {
    let start_x = line.find("X").unwrap();
    let separator = line.find(", ").unwrap();
    let start_y = line.find("Y").unwrap();
    let x: i64 = line[start_x + 2..separator].parse().unwrap();
    let y: i64 = line[start_y + 2..].parse().unwrap();
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

fn get_game_tokens(game: &Game) -> Option<i64> {
    for a_button_presses in 0..(game.prize.0 / game.a_button.0 + 2) {
        for b_button_presses in 0..(game.prize.0 / game.b_button.0 + 2) {
            if a_button_presses * game.a_button.1 + b_button_presses * game.b_button.1
                == game.prize.1
                && a_button_presses * game.a_button.0 + b_button_presses * game.b_button.0
                    == game.prize.0
            {
                println!(
                    "a presses: {}  b presses: {}",
                    a_button_presses, b_button_presses
                );
                return Some(a_button_presses * 3 + b_button_presses);
            }
        }
    }
    return None;
}

#[allow(dead_code)]
pub fn solve(inputs: &Vec<String>) -> String {
    let games = parse_games(inputs);
    let mut tokens = 0;
    for game in &games {
        let maybe_tokens = get_game_tokens(game);
        println!("{:?}  {:?}", maybe_tokens, game);
        if maybe_tokens.is_some() {
            tokens += maybe_tokens.unwrap();
        }
    }

    tokens.to_string()
}
