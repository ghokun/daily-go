package dailygo

/*
cons(a, b) constructs a pair, and car(pair) and cdr(pair) returns the first and
last element of that pair. For example, car(cons(3, 4)) returns 3, and
cdr(cons(3, 4)) returns 4.

Given this implementation of cons:

def cons(a, b):
    def pair(f):
        return f(a, b)
    return pair

Implement car and cdr.
*/

func _20220316(a int, b int) (int, int) {

	type f func(int, int) int
	type pair func(f) int

	cons := func(a int, b int) pair {
		p := func(fnc f) int {
			return fnc(a, b)
		}
		return p
	}
	car := func(p pair) int {
		return p(func(a int, b int) int {
			return a
		})
	}
	cdr := func(p pair) int {
		return p(func(a int, b int) int {
			return b
		})
	}

	return car(cons(a, b)), cdr(cons(a, b))
}
