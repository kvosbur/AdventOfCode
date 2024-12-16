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

fn score_bots(bots: &Vec<Bot>, mid_x: i64, mid_y: i64) -> i64 {
    let mut quadrant_values = [0, 0, 0, 0];
    for bot in bots {
        // top left quadrant
        if bot.pos.0 < mid_x && bot.pos.1 < mid_y {
            quadrant_values[0] += 1;
        }
        // top right
        if bot.pos.0 > mid_x && bot.pos.1 < mid_y {
            quadrant_values[1] += 1;
        }
        // bottom right
        if bot.pos.0 > mid_x && bot.pos.1 > mid_y {
            quadrant_values[2] += 1;
        }
        // bottom left
        if bot.pos.0 < mid_x && bot.pos.1 > mid_y {
            quadrant_values[3] += 1;
        }
    }
    return quadrant_values[0] * quadrant_values[1] * quadrant_values[2] * quadrant_values[3];
}

#[allow(dead_code)]
pub fn solve(inputs: &Vec<String>) -> String {
    let mut bots = parse_bots(inputs);
    let max_x = 101;
    let max_y = 103;
    for _ in 0..100 {
        move_bots(&mut bots, max_x, max_y);
    }

    let score = score_bots(&bots, 50, 51);

    score.to_string()
}
