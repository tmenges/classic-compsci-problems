
const FIBONACCI: [u32; 16]= [0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610];

fn main() {
    let n = 8;
    println!("fib({}) = {}, {}", n, fib(n), FIBONACCI[n as usize]);
    println!("fib_iterative({}) = {}, {}", n, fib_iterative(n), FIBONACCI[n as usize]);
}

fn fib(n: u32) -> u32 {
    if n == 0 {
        return 0
    }

    if n == 1 {
        return 1
    }

    fib(n - 1) + fib(n - 2)
}

fn fib_iterative(n: u32) -> u32 {
    if n == 0 {
        return 0
    }

    if n == 1 {
        return 1
    }

    // cache[0] == n - 1
    // cache[1] == n - 2
    let mut cache: [u32; 2] = [1, 0];
    
    cache[0] = 1;
    cache[1] = 0;
    for _x in 2..n+1 {
        let next = cache[0] + cache[1];
        cache[1] = cache[0];
        cache[0] = next;
    }

    cache[0]
}

#[cfg(test)]
mod tests {
    use crate::FIBONACCI;
    use crate::fib;
    use crate::fib_iterative;

    #[test]
    fn test_zero() {
        assert_eq!(fib(0), FIBONACCI[0]);
    }

    #[test]
    fn test_fibs() {
        for n in 1..FIBONACCI.len() {
            assert_eq!(fib(n as u32), FIBONACCI[n])
        }
    }

    #[test]
    fn test_fib_iterative() {
        for n in 1..FIBONACCI.len() {
            assert_eq!(fib_iterative(n as u32), FIBONACCI[n])
        }
    }

}