use std::fs;

pub fn solution() {
    let chars: String = fs::read_to_string("src/day01.input").unwrap().parse().unwrap();
    let mut that_elf = (0, 0);
    let mut counter: u32 = 0;
    let mut position = 1;
    let input: Vec<&str> = chars.split("\n").collect();

    for s in input {
        let num = s.trim().parse::<u32>();
        match num {
            Ok(n) => counter += n,
            Err(e) => match e.kind() {
                std::num::IntErrorKind::Empty => {
                    if counter > that_elf.1 {
                        that_elf.0 = position;
                        that_elf.1 = counter;
                    }
                    counter = 0;
                    position += 1;
                },
                _ => {},
            }
        }
    }

    println!("The elf number {} carrying the most calories of {}",
             that_elf.0, that_elf.1);
}

