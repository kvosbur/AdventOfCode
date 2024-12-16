use std::collections::HashSet;

#[derive(Debug)]
struct Bot {
    pos: (i64, i64),
    vel: (i64, i64),
}

fn parse_line(line: &str) -> (i64, i64) {
    let temp: Vec<&str> = line.split(',').collect();
    return (temp[0][2..].parse().unwrap(), temp[1].parse().unwrap());
}

fn parse_bots(inputs: &Vec<String>) -> Vec<Bot> {
    let mut bots = vec![];
    for line in inputs {
        let space_index = line.find(' ').unwrap();
        let pos_line = &line[..space_index];
        let vel_line = &line[space_index + 1..];
        bots.push(Bot {
            pos: parse_line(pos_line),
            vel: parse_line(vel_line),
        })
    }

    return bots;
}

fn wrap_value(val: i64, max_value: i64) -> i64 {
    if val >= max_value {
        return val - max_value;
    } else if val < 0 {
        return val + max_value;
    }
    return val;
}

fn move_bots(bots: &mut Vec<Bot>, max_x: i64, max_y: i64) {
    for bot in bots {
        bot.pos.0 = wrap_value(bot.vel.0 + bot.pos.0, max_x);
        bot.pos.1 = wrap_value(bot.vel.1 + bot.pos.1, max_y);
    }
}

fn bots_all_at_different_pos(bots: &Vec<Bot>) -> bool {
    let mut positions: HashSet<(i64, i64)> = HashSet::new();

    for bot in bots {
        if positions.contains(&(bot.pos.0, bot.pos.1)) {
            return false;
        }
        positions.insert((bot.pos.0, bot.pos.1));
    }
    return true;
}

fn print_bots_pic(bots: &Vec<Bot>, max_x: i64, max_y: i64) {
    for x in 0..max_x {
        for y in 0..max_y {
            let mut has_bot = false;
            for bot in bots {
                if bot.pos.0 == x && bot.pos.1 == y {
                    has_bot = true;
                    break;
                }
            }
            print!("{}", if has_bot { "." } else { "x" });
        }
        println!("");
    }
}

#[allow(dead_code)]
pub fn solve(inputs: &Vec<String>) -> String {
    let mut bots = parse_bots(inputs);
    let max_x = 101;
    let max_y = 103;

    let mut seconds = 0;
    while !bots_all_at_different_pos(&bots) {
        move_bots(&mut bots, max_x, max_y);
        seconds += 1;
    }
    print_bots_pic(&bots, max_x, max_y);

    seconds.to_string()
}
