#[derive(Debug)]
struct DiskNode {
    used_pages: Vec<Option<u64>>,
    original_used_length: u32,
    free_pages: u32,
}

fn read_nodes(bytes: Vec<u8>) -> Vec<DiskNode> {
    let mut disk_nodes: Vec<DiskNode> = Vec::new();
    for start_index in (0..bytes.len()).step_by(2) {
        let file_id = start_index / 2;
        let used_blocks = bytes[start_index] - b'0';
        let free_blocks: u8;
        if start_index + 1 == bytes.len() {
            free_blocks = 0;
        } else {
            free_blocks = bytes[start_index + 1] - b'0';
        }
        let used_blocks_vec: Vec<Option<u64>> = (0..used_blocks)
            .map(|_r| Some(file_id.try_into().unwrap()))
            .collect();
        disk_nodes.push(DiskNode {
            original_used_length: used_blocks_vec.len().try_into().unwrap(),
            used_pages: used_blocks_vec,
            free_pages: free_blocks.try_into().unwrap(),
        });
    }
    return disk_nodes;
}

fn get_checksum(disk_nodes: &Vec<DiskNode>) -> u64 {
    let mut checksum: u64 = 0;
    let mut overall_index: u64 = 0;
    for node in disk_nodes {
        for block in &node.used_pages {
            if block.is_some() {
                checksum += block.unwrap() * (overall_index);
            }

            overall_index += 1;
        }
        let index_increase_amount: u64 = node.free_pages.try_into().unwrap();
        overall_index += index_increase_amount;
    }
    return checksum;
}

#[allow(dead_code)]
pub fn solve(inputs: &Vec<String>) -> String {
    let bytes = inputs[0].clone().into_bytes();

    let mut disk_nodes: Vec<DiskNode> = read_nodes(bytes);

    let mut right_node_index = disk_nodes.len() - 1;

    while right_node_index > 0 {
        let right_node_read = &disk_nodes[right_node_index];
        if right_node_read.used_pages.len() == 0 {
            right_node_index -= 1;
            continue;
        }
        let mut left_node_index = 0;
        while left_node_index < disk_nodes.len() {
            let left_node_read = &disk_nodes[left_node_index];
            if left_node_read.free_pages < right_node_read.original_used_length {
                left_node_index += 1;
                continue;
            }

            let mut right_moving = right_node_read.used_pages.clone();
            let right_moving_count = right_node_read.original_used_length.clone();

            let left_node_write = &mut disk_nodes[left_node_index];

            for remove_index in 0..right_moving_count {
                let i: usize = remove_index.try_into().unwrap();
                left_node_write.used_pages.push(right_moving[i]);
                right_moving[i] = None;
                left_node_write.free_pages -= 1;
            }

            let right_node_write = &mut disk_nodes[right_node_index];
            right_node_write.used_pages = right_moving;
            break;
        }
        right_node_index -= 1;
    }

    // print!("{:?}", disk_nodes);
    get_checksum(&disk_nodes).to_string()
}

// currently getting too high of a value
