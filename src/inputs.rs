use std::fs;

pub fn get_input(path: &str) -> Vec<String> {
    let chars: String = fs::read_to_string(format!("src/{}", path)).unwrap().parse().unwrap();
    chars.split(&['\r', '\n']).map(|s| s.to_string()).collect()
}
