#[derive(Debug)]
struct DiskNode {
    id: u128,
    used_pages: Vec<u128>,
    free_pages: u128,
}

#[allow(dead_code)]
pub fn solve(inputs: &Vec<String>) -> String {
    let bytes = inputs[0].clone().into_bytes();

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
        let used_blocks_vec = (0..used_blocks)
            .map(|_r| file_id.try_into().unwrap())
            .collect();
        disk_nodes.push(DiskNode {
            id: file_id.try_into().unwrap(),
            used_pages: used_blocks_vec,
            free_pages: free_blocks.try_into().unwrap(),
        });
    }

    let mut left_node_index = 0;
    let mut right_node_index = disk_nodes.len() - 1;

    while left_node_index < right_node_index {
        let right_node_read = &disk_nodes[right_node_index];
        if right_node_read.used_pages.len() == 0 {
            right_node_index -= 1;
            continue;
        }

        let left_node_read = &disk_nodes[left_node_index];
        if left_node_read.free_pages == 0 {
            left_node_index += 1;
            continue;
        }

        let moving = std::cmp::min(
            left_node_read.free_pages.try_into().unwrap(),
            right_node_read.used_pages.len(),
        );
        let right_moving = &right_node_read.used_pages.clone();

        let left_node_write = &mut disk_nodes[left_node_index];

        for remove_index in 0..moving {
            left_node_write
                .used_pages
                .push(right_moving[right_moving.len() - remove_index - 1]);
            left_node_write.free_pages -= 1;
        }

        let right_node_write = &mut disk_nodes[right_node_index];
        for _remove_index in 0..moving {
            right_node_write.used_pages.pop();
            right_node_write.free_pages += 1;
        }
    }

    let mut checksum: u128 = 0;
    let mut overall_index: u128 = 0;
    for node in &disk_nodes {
        if node.used_pages.len() == 0 {
            break;
        }
        for block in &node.used_pages {
            checksum += block * (overall_index);
            overall_index += 1;
        }
    }

    // print!("{:?}", disk_nodes);
    checksum.to_string()
}
