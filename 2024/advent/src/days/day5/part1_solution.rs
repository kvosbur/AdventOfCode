use std::collections::HashMap;

#[allow(dead_code)]
pub fn solve(inputs: &Vec<String>) -> String {
    let mut mappings: HashMap<i32, Vec<i32>> = HashMap::new();

    let split_index = inputs.iter().position(|r| r == "").unwrap();
    for setup in &inputs[..split_index] {
        let split_line: Vec<&str> = setup.split("|").collect();
        let first_val: i32 = split_line[0].parse().unwrap();
        let second_val: i32 = split_line[1].parse().unwrap();

        mappings
            .entry(first_val)
            .and_modify(|v| v.push(second_val))
            .or_insert(vec![second_val]);
    }

    let mut cumullative_sum = 0;
    for manual_updates in &inputs[split_index + 1..] {
        let split_pages: Vec<i32> = manual_updates
            .split(",")
            .map(|page| page.parse::<i32>().unwrap())
            .collect();

        println!("pages: {:?}", split_pages);

        let mut seen_pages: Vec<i32> = vec![];
        let mut bad_update = false;
        for page in &split_pages {
            let bad_pages = mappings.entry(*page).or_default();

            let has_seen_bad_page = bad_pages.iter().any(|item| seen_pages.contains(item));
            if has_seen_bad_page {
                bad_update = true;
                break;
            }
            seen_pages.push(*page);
        }
        if !bad_update {
            let middle_index = split_pages.len() / 2;
            cumullative_sum += split_pages[middle_index];
        }
    }

    cumullative_sum.to_string()
}
