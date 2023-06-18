use std::fs;

pub fn solution() {
    let chars: String = fs::read_to_string("src/day01.input").unwrap().parse().unwrap();
    let input: Vec<&str> = chars.split("\n").collect();
    let mut inventory: Vec<u32> = vec![];
    let mut counter: u32 = 0;
    for s in input {
        let num = s.trim().parse::<u32>();
        match num {
            Ok(n) => counter += n,
            Err(e) => match e.kind() {
                std::num::IntErrorKind::Empty => {
                    inventory.push(counter);
                    counter = 0;
                },
                _ => {},
            }
        }
    }

    inventory.sort_by(|a, b| b.cmp(a));

    let top1: u32 = *inventory.get(0).unwrap_or(&0);
    let top3: u32 = inventory.get(0..3).unwrap_or(&[]).iter().sum();

    println!("\n| {0: <10} | {1: <10} |", "Name", "Value");
    println!("| {0: <10} | {1: <10} |", "----------",  "----------");
    println!("| {0: <10} | {1: <10} |", "Top 1", top1);
    println!("| {0: <10} | {1: <10} |", "Top 3", top3);
}

