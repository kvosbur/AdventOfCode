use std::{collections::HashMap, ops::Sub};

fn add_to_group(
    x: usize,
    y: usize,
    input_chars: &Vec<&[u8]>,
    groups: &mut HashMap<u8, Vec<Vec<(usize, usize)>>>,
) {
    let value = input_chars[x][y];

    if !groups.contains_key(&value) {
        groups.insert(value, vec![]);
    }

    let value_sets = groups.get_mut(&value).unwrap();

    // check top
    let mut top_set_index: Option<usize> = None;
    if y > 0 && input_chars[x][y - 1] == value {
        top_set_index = value_sets.iter().position(|set| set.contains(&(x, y - 1)));
    }

    // check left
    let mut left_set_index: Option<usize> = None;
    if x > 0 && input_chars[x - 1][y] == value {
        left_set_index = value_sets.iter().position(|set| set.contains(&(x - 1, y)));
    }

    if top_set_index.is_some() && left_set_index.is_some() {
        if top_set_index == left_set_index {
            value_sets
                .get_mut(left_set_index.unwrap())
                .unwrap()
                .push((x, y));
        } else {
            // combine sets
            if top_set_index < left_set_index {
                left_set_index = Some(left_set_index.unwrap().sub(1));
            }
            let mut top = value_sets.remove(top_set_index.unwrap());
            let left = value_sets.get_mut(left_set_index.unwrap()).unwrap();
            left.append(&mut top);
            left.push((x, y));
        }
    } else if top_set_index.is_some() {
        value_sets
            .get_mut(top_set_index.unwrap())
            .unwrap()
            .push((x, y));
    } else if left_set_index.is_some() {
        value_sets
            .get_mut(left_set_index.unwrap())
            .unwrap()
            .push((x, y));
    } else {
        value_sets.push(vec![(x, y)]);
    }
}

fn make_groups(input_chars: &Vec<&[u8]>) -> HashMap<u8, Vec<Vec<(usize, usize)>>> {
    let mut groups: HashMap<u8, Vec<Vec<(usize, usize)>>> = HashMap::new();

    for x in 0..input_chars.len() {
        for y in 0..input_chars[0].len() {
            add_to_group(x, y, input_chars, &mut groups);
        }
    }

    return groups;
}

fn score_group(set: &Vec<(usize, usize)>, input_chars: &Vec<&[u8]>) -> u64 {
    let mut perimeter: u64 = 0;
    let value = input_chars[set[0].0][set[0].1];
    for (x, y) in set {
        // check top
        if *y == 0 || input_chars[*x][y - 1] != value {
            perimeter += 1;
        }
        //check left
        if *x == 0 || input_chars[*x - 1][*y] != value {
            perimeter += 1;
        }
        // check bottom
        if *y == input_chars[0].len() - 1 || input_chars[*x][y + 1] != value {
            perimeter += 1;
        }
        // check right
        if *x == input_chars.len() - 1 || input_chars[*x + 1][*y] != value {
            perimeter += 1;
        }
    }
    let area: u64 = set.len().try_into().unwrap();
    return perimeter * area;
}

#[allow(dead_code)]
pub fn solve(inputs: &Vec<String>) -> String {
    let input_chars: Vec<&[u8]> = inputs.iter().map(|s| s.as_bytes()).collect();

    let groups = make_groups(&input_chars);

    let mut score = 0;
    for (_, sets) in groups {
        for set in sets {
            score += score_group(&set, &input_chars);
        }
    }

    score.to_string()
}
