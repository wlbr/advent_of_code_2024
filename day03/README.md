
## Day 3: Mull It Over

"Our computers are having issues, so I have no idea if we have any Chief Historians in stock! You're welcome to check the warehouse, though," says the mildly flustered shopkeeper at the [North Pole Toboggan Rental Shop](https://adventofcode.com/2020/day/2). The Historians head out to take a look.

The shopkeeper turns to you. "Any chance you can see why our computers are having issues again?"

The computer appears to be trying to run a program, but its memory (your puzzle input) is __corrupted__. All of the instructions have been jumbled up!

It seems like the goal of the program is just to __multiply some numbers__. It does that with instructions like `mul(X,Y)`, where `X` and `Y` are each 1-3 digit numbers. For instance, `mul(44,46)` multiplies `44` by `46` to get a result of 2024. Similarly, `mul(123,4)` would multiply `123` by `4`.

However, because the program's memory has been corrupted, there are also many invalid characters that should be __ignored__, even if they look like part of a mul instruction. Sequences like `mul(4*`, `mul(6,9!`, `?(12,34)`, or `mul ( 2 , 4 )` do __nothing__.

For example, consider the following section of corrupted memory:

<code>x<b>mul(2,4)</b>%&amp;mul[3,7]!@^do_not_<b>mul(5,5)</b>+mul(32,64]then(<b>mul(11,8)mul(8,5)</b>)</code>

Only the four highlighted sections are real mul instructions. Adding up the result of each instruction produces __`161`__ `(2*4 + 5*5 + 11*8 + 8*5)`.

Scan the corrupted memory for uncorrupted mul instructions. __What do you get if you add up all of the results of the multiplications?__

> Your puzzle answer was 171183089.

## Part Two

As you scan through the corrupted memory, you notice that some of the conditional statements are also still intact. If you handle some of the uncorrupted conditional statements in the program, you might be able to get an even more accurate result.

There are two new instructions you'll need to handle:

   * The `do()` instruction __enables__ future mul instructions.
   * The `don't()` instruction __disables__ future mul instructions.

Only the __most recent__ `do()` or `don't()` instruction applies. At the beginning of the program, mul instructions are __enabled__.

For example:

<code>x<b>mul(2,4)</b>&amp;mul[3,7]!^<b>don't()</b>_mul(5,5)+mul(32,64](mul(11,8)un<b>do()</b>?<b>mul(8,5)</b>)</code>

This corrupted memory is similar to the example from before, but this time the `mul(5,5)` and `mul(11,8)` instructions are __disabled__ because there is a `don't()` instruction before them. The other mul instructions function normally, including the one at the end that gets re-__enabled__ by a `do()` instruction.

This time, the sum of the results is __`48`__ `(2*4 + 8*5)`.

Handle the new instructions; __what do you get if you add up all of the results of just the enabled multiplications?__

> Your puzzle answer was 63866497.

> Both parts of this puzzle are complete! They provide two gold stars: **
