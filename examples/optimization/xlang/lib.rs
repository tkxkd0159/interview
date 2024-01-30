fn double_num(a: i32) -> i32 {
    2 * a
}

fn square(a: i32) -> i32 {
    a * a
}

#[no_mangle]
pub extern "C" fn calculation(a: i32, b: i32) -> i32 {
    square(a) * double_num(b)
}
