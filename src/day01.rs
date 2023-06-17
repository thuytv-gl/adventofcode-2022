pub fn solution() {
    let input = "1000
2000
3000

4000

5000
6000

7000
8000
9000

10000";
    let mut that_elf = (0, 0);
    let mut counter: u16 = 0;
    let mut position = 1;
    let input: Vec<&str> = input.split("\n").collect();

    for s in input {
        let num = s.trim().parse::<u16>();
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

