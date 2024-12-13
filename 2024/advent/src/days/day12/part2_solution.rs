fn score_trailhead(start_x: usize, start_y: usize, characters: &Vec<&[u8]>) -> usize {
    let mut locations = vec![(start_x, start_y, b'0')];
    let mut trailhead_score = 0;
    while locations.len() > 0 {
        let (x_loc, y_loc, value) = locations.pop().unwrap();
        if value == b'9' {
            trailhead_score += 1;
            continue;
        }
        let target_value = value + 1;

        // check left
        if x_loc > 0 && characters[x_loc - 1][y_loc] == target_value {
            locations.push((x_loc - 1, y_loc, target_value));
        }
        // check down
        if y_loc > 0 && characters[x_loc][y_loc - 1] == target_value {
            locations.push((x_loc, y_loc - 1, target_value));
        }

        // check right
        if x_loc < characters.len() - 1 && characters[x_loc + 1][y_loc] == target_value {
            locations.push((x_loc + 1, y_loc, target_value));
        }

        // check up
        if y_loc < characters[0].len() - 1 && characters[x_loc][y_loc + 1] == target_value {
            locations.push((x_loc, y_loc + 1, target_value));
        }
    }

    trailhead_score
}

#[allow(dead_code)]
pub fn solve(inputs: &Vec<String>) -> String {
    let input_chars: Vec<&[u8]> = inputs.iter().map(|s| s.as_bytes()).collect();

    let mut trailhead_score_sum = 0;
    for x in 0..input_chars.len() {
        for y in 0..input_chars[0].len() {
            if input_chars[x][y] == b'0' {
                trailhead_score_sum += score_trailhead(x, y, &input_chars);
            }
        }
    }

    trailhead_score_sum.to_string()
}
