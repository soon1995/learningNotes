# JavaScript - The Definitive Guide -- Seventh Edition

Master the World's Most-Used Programming Language

By David Flanagan

| ES5  | ES6  | ES2016...ES2020 |
| ---- | ---- | --------------- |
| 2009 | 2015 | 2016-2020       |

> Table of Contents

[toc]

## 1 - Preface

### Exploring JavaScript

To try out a few lines of JavaScript:

1. Open up the web developer tools in web browser

2. Download and install node.

  - type `node` to begin an interactive JavaScript session

  - run js file with `node <yourfile>`

3. Load script in html

  - ```html
    <script src="hello.js"></script>
    ```

## 2 - Lexical Structure

### 2.1 The Text of a JavaScript Program

- Case-sensitive language

### 2.2 Comments

- `//`
- `/* ... */`

### 2.3 Literals

The following are all literals:

```js
12; // The number twelve
1.2; // The number one point two
("hello world"); // A string of text
("Hi"); // Another string
true; // A Boolean value
null; // Absence of an object
```

### 2.4 Identifiers and Reserved Words

Valid Identifiers:

- Begin with a letter, an underscore (\_), or a dollarsign
- Subsequent characters can be letters, digits, underscore (\_)
  or a dollar sign

These are all legal identifiers:

```js
i;
my_variable_name;
v13;
_dummy;
$str;
```

#### 2.4.1 Reserved Words

`let` can't be fully reserved in order to retain backward compatibility with
older programs. (`let` can be used with a variable name if declared with var outside of a class). **The simplest course is to avoid using
any of these words as identifiers, except for `from`, `set`, and `target`, which are safe to use and are already in common use.**

```
as    const     export    get        null    target   void
async continue  extends   if         of      this     while
await debugger  false     import     return  throw    with
break default   finally   in         set     true     yield
case  delete    for       instanceof static  try
catch do        from      let        super   typeof
class else      function  new        switch  var

/* Reserve or restricts certain keywords that might be used in future versions */
enum  implements    interface     package   private   protected   public
```

For historical reasons, `arguments` and `eval` are not allowed as
identifiers in certain circumstances and are best avoided.

### 2.5 Unicode

Programming convention says only ASCII letters and digits
in identifiers. JavaScript allows Unicode letters, digits, and
ideographs in identifiers.

#### 2.5.1 Unicode Escape Sequences

Some computer hardware and software cannot display, input, or
correctly process the full set of Unicode characters. To support
programmers and systems using older technology, JavaScript defines
escape sequences that allow us to write Unicode characters using
only ASCII characters (which begin with the characters \u and are
either followed by exactly four hexadecimal digits (using upper or lowercase A-F) or by one to six hexadecimal digits enclosed within curly
braces. The latter was introduced in ES6 to support Unicode codepoints
that require more than 16 bits. (such as emoji)

```js
let café = 1; // Define a variable using a Unicode character
café;
café;
```

### 2.6 Optional Semicolons

In JavaScript, you can usually omit the semicolon between two statements
if those statements are written on separate lines.

However, you should understand about optional semicolons in JavaScript.

JavaScript usually treats line breaks as semicolons if **it can't parse
the code without adding an implicit semicolon**. Consider the following code:

```
let a
a
=
3
console.log(a);
```

JavaScript interprets this code like this:

```js
let a;
a = 3;
console.log(a);
```

These statement termination rules lead to some surprising cases:

```
let y = x + f
(a + b).toString()
```

And JavaScript interprets the code like this:

```js
let y = x + f(a + b).toString();
```

In general, **if a statement begins with (, [, /, + or -,
there is a chance that it could be interpreted as a continuation of the statement before.** ( and [ is not uncommon at all. Some programmers
like to put a defensive semicolon at the beginning of
any such statement:

```
let x = 0;
;[x, x + 1, x + 2].forEach(console.log);
```

There are three exception to the general rule that JavaScript interprets
line breaks as semicolons when it cannot parse the second line as a
continuation of the statement on the first line.

1. `return`, `throw`, `yield`, `break`, `continue`. JavaScript will
   always interpret that line break as a semicolon.

- ```bash
  return
  true;

  // JavaScript assumes you meant:
  return; true;
  ```

2. The `++` and `--` operators. If you want to use either of these operators
   as postfix operators, they must appear on the same line as the expression
   they apply to.

3. `=>` must appear on the same line **as the parameter list**

## 3 - Types, Value, and Variables

### 3.1 Overview and Definitions

JavaScript types can be divided into two categories:

- primitive types (**immutable**)
  - numbers
  - strings
  - booleans
  - undefined
  - null
  - Symbol (introduced in ES6)
- object types, is a collection of properties. (**mutable**)
  - Arrays
  - Object
  - Set
  - Map
  - Typed Array
  - RegExp type
  - Date type
  - Error and its subtypes
  - Functions
  - Classes
  - Global Object

**Tip**: In JavaScript, `null` and `undefined` are the only
values that methods cannot be invoked on. **Primitive types** do
have methods.

**Tip**: Strings are immutable: you can access the text at any
index of a string, but JavaScript provides no way to alter
the text of an existing string.

### 3.2 Numbers

Real numbers. A JavaScript number format allows you to exactly represent
all integers between -9,007,199,254,740,992 (-2<sup>53</sup>) to 9,007,199,254,740,922 (2<sup>53</sup>)
You may lose precision in the trailing digits if you use integer
values larger than this (9,007,199,254,740,923 became ...922). If you
need exactly represent larger integers, see [3.2.5 Arbitrary Precision Integers with BigInt](#3.2.5-arbitrary-precision-integers-with-bigint)

#### 3.2.1 Integer Literals

```js
0;
3;
10000000;
0xff; // => 255 (15*16 + 15)
0b10101; // => 21, In ES6 and later
0o377; // => 255, In ES6 and later (Octal)
```

#### 3.2.2 Floating-Point Literals

```js
[digits][.digits][(E|e)[(+|-)]digits]
// Eponential notation represents the real number multiplied by 10 to
// the power of the exponent
```

```js
3.14;
2345.6789;
0.3333333333333;
6.02e23; // 6.02 × 10²³
1.4738e-32; // 1.4738 × 10⁻³²
```

**Tip** Separators in numeric literals for easier reading.

```js
let billion = 1_000_000_000;
let bytes = 0x89_ab_cd_ef;
let btis = 0b0001_1101_0111;
let fraction = 0.123_456_789;
```

#### 3.2.3 Arithmetic in JavaScript

Arithmetic operators includes: +, -, \*, /, %, \*\*(ES2016, exponentiation)

Arithmetic in JavaScript does not raise errors in cases of overflow,
underflow (closer to zero than the smallest representable number),
or division by zero.

Possible values: `Infinity`, `-Infinity`, `NaN`, `-0`

**Tip**: NaN if 0/0, else Infinity for 1/0, -Infinity for -1/0

**Tip**: Infinity and NaN is a global constants. They also available as properties of the Number object -- Number.POSITIVE_INFINITY.

**Tip**: Following causes NaN

- Divide Infinity by Infinity
- Take the square root of a negative number
- Use Arithmetic operators with non-numeric operands that cannot be converted to numbers.

```js
Infinity;
Number.POSITIVE_INFINITY;
1 / 0; // => Infinity
Number.MAX_VALUE * 2 - // => Infinity; overflow

=======

-Infinity;
Number.NEGATIVE_INFINITY;
-1 / 0;
-Number.MAX_VALUE * 2;

NaN;
Number.NaN;
0 / 0; // => NaN
Infinity / Infinity; // => NaN

Number.MIN_VALUE / 2 - // => 0: underflow
Number.MIN_VALUE / 2 - // => -0: negative zero
1 / Infinity - // => -0
0;
```

Following Number properties are defined in ES6

```js
Number.parseInt();
Number.parseFloat();
Number.isNaN();
Number.isFinite(); // true if argument is number other than NaN, Infinity or -Infinity
Number.isInteger();
Number.isSafeInteger(x); // -(2**53) < x < 2**53 ?
Number.MIN_SAFE_INTEGER;
Number.MAX_SAFE_INTEGER; // 2**53 - 1
Number.EPSILON; // 2**-52: smallest difference between number
```

**Caution**: You cannot compare NaN with ===, use Number.isNaN function, or x != x

#### 3.2.4 Binary Floating-Point and Rounding Errors

The IEEE-754 floating-point representation used by JavaScript
is a binary representation, which can exactly represent fractions
like 1/2, 1/8, 1/1024. However, the fractions we use most commonly
are decimal fractions: 1/10, 1/100, and so on. **Binary floating-point
representation cannot exactly represent numbers as
simple as 0.1.**

Consider this code:

```js
let x = 0.3 - 0.2;
let y = 0.2 - 0.1;
x === y; // false!!
x === 0.1; // false
y === 0.1; // true
```

If these floating-point approximations are problematic for
your programs, consider using scaled integers. For example,
you might manipulate monetary values as integer cents rather than
fractional dollars.

#### 3.2.5 Arbitrary Precision Integers with BigInt

A new numeric type known as BigInt was defined in ES2020. (Note,
that BigInt implementations are not suitable for cryptography because
they do not attempt to prevent timing attacks.)

BitInt literals are written as a string of digits followed by a
lowercase `n`. By default, they are in base 10, but you can use
the `0b`, `0o` and `0x`.

```js
1234n;
0b111111n;
0o7777n;
0x8000000000000000n;
```

You can use `BigInt()` to convert numbers or strings to BigInt values

```js
BigInt(Number.MAX_SAFE_INTEGER);
let string = "1" + "0".repeat(100);
BigInt(string);
```

Arithmetic works like with regular JavaScript numbers, **except
that division drops any remainder and rounds down (toward zero):**

```js
1000n + 1000n;
3000n - 2000n;
2000n * 3000n;
3000n / 997n; // => 3n
3000n % 997n; // => 9n
(2n ** 131071n) - 1n // A Mersenne prime with 39457 decimal digits
```

You may not mix operands of type BigInt with regular number operands.
Comparison operators, by contrast, do work with mixed numeric types

```js
1 < 2n; // => true
2 > 1n; // => true
0 == 0n; // => true
0 === 0n; // => false
```

#### 3.2.6 Dates and Times

JavaScript Dates are objects, but they also have
a numeric representation as a timestamp that
**specifies the number of elapsed milliseconds since
January 1, 1970**:

```js
let ts = Date.now(); // timestamp (a number)
let now = new Date(); // a Date Object
let ms = now.getTime(); // Convert to a milliseconds timestamp
let iso = now.toISOString(); // Convert to a string in standard format '2023-11-23T15:58:18.752Z'
```

### 3.3 Text

JavaScript use the UTF-16 encoding of the Unicode character set,
and JavaScript strings are sequences of unsigned 16-bit values.
Unicode characters whose codepoints do not fit in 16 bits are encoded
using the rules of UTF-16 as a sequence (known as a "surrogate pair") of
two 16-bit values.

```js
let euro = "€";
let love = "👀";
euro.length; // => 1: this character has one 16-bit element
love.length; // => 2: UTF-16 encoding of ❤ is "\ud83d\udc99" (NOPE!)
```

**In ES6, however, strings are iterable, if you use the `for/of` loop
of ...operator with a string, it will iterate the actual characters
of the string, not the 16-bit values.**

```js
for (const l of love) {
  console.log(l)
}
```

#### 3.3.1 String Literals

`"` or `'` or backtick (backticks a re a feature of ES6)

```js
""; // empty string
"testing";
"3.14";
'name="myform"';
"Wouldn't you prefer O'Reilly's book?"`"She said 'hi'", he said.`;
```

To concatenating single-line strings, use `+` operator.
A of ES5, you can break string literal across multiple lines
by ending each line but last with a backslash (\).

```js
let a =
  "I \
Love \
Books.";
// I Love Books.
// Single quote do the same thing.
```

Backtick syntax allows strings to be broken across multiple
lines

```js
let ihatebook = `I
Hate
Books.`;
// I
// Hate
// Books.
```

You can escape a new line

```js
let a = `\
asdf`;

// asdf
```

#### 3.3.2 Escape Sequences in String Literals

The backslash character (\) has a special purpose in JavaScript strings. For example,
\n is an escape sequence that represents a newline character.

**Tip**: \unnnn or \u{n} is the Unicode character

**Tip**: If the \ character precedes any character that is not a common escape character.
**The backslash is simply ignored**. Finally, as noted earlier, ES5 allows a
backslash before a line break to break a string literal across multiple lines.

#### 3.3.3 Working with Strings

```bash
let msg = "Hello " + "world"; // Hello world
```

| Compare         | Description                                                                                                 |
| --------------- | ----------------------------------------------------------------------------------------------------------- |
| === or !==      | If they consist of exactly the same sequence of 16-bit values                                               |
| <, <=, > and >= | comparing the 16-bit values, [see here for more robust locale aware string comparison and sorting](#11.7.3) |

Rich API for working with strings:

Example: `s = "Hello, world"`

| API                                                        | description                                         | Example                                |
| ---------------------------------------------------------- | --------------------------------------------------- | -------------------------------------- |
| **Obtaining portions of a string**                         |                                                     |                                        |
| s.substring(1,4)                                           |                                                     | ell                                    |
| s.substring(1)                                             |                                                     | ello, world                            |
| s.slice(1,4)                                               |                                                     | ell (same thing)                       |
| s.slice(-3)                                                |                                                     | rld                                    |
| s.split(", ")                                              | split at **delimiter string(can be not character)** | ["Hello", "world"]                     |
|                                                            |                                                     |                                        |
| **Searching a string**                                     |                                                     |                                        |
| s.indexOf("l")                                             |                                                     | 2                                      |
| s.indexOf("l", 3)                                          |                                                     | 10                                     |
| s.indexOf("zz")                                            |                                                     | -1                                     |
| s.lastIndexOf("l")                                         |                                                     | 10                                     |
|                                                            |                                                     |                                        |
| **Boolean searching functions in ES6 and later**           |                                                     |                                        |
| s.startsWith("Hell")                                       |                                                     | true                                   |
| s.endWith("!")                                             |                                                     | false                                  |
| s.includes("or")                                           |                                                     | true, it includes **substring**        |
|                                                            |                                                     |                                        |
| **Creating modified versions of a string**                 |                                                     |                                        |
| s.replace("llo", "ya")                                     |                                                     | Heya, world                            |
| s.toLowerCase()                                            |                                                     | hello, world                           |
| s.toUpperCase()                                            |                                                     | HELLO, WORLD                           |
| s.normalize()                                              |                                                     | Unicode NFC normalization: ES6         |
| s.normalize("NFD")                                         |                                                     | NFD normalization. Also "NFNC", "NFKD" |
|                                                            |                                                     |                                        |
| **Inspecting individual (16-bit) characteres of a string** |                                                     |                                        |
| s.charAt(0)                                                | s[0]                                                | "H"                                    |
| s.charAt(s.length - 1)                                     |                                                     | "d"                                    |
| s.charCodeAt(0)                                            |                                                     | 72                                     |
| s.codePointAt(0)                                           |                                                     | 72                                     |
|                                                            |                                                     |                                        |
| **String padding functions in ES2017**                     |                                                     |                                        |
| "x".padStart(3)                                            |                                                     | "\_\_x"                                |
| "x".padEnd(3)                                              |                                                     | "x\_\_"                                |
| "x".padStart(3, "\*")                                      |                                                     | "\*\*x"                                |
| "x".padEnd(3, "-")                                         |                                                     | "x--"                                  |
|                                                            |                                                     |                                        |
| **Space triming functions. trim() is ES5; others 2019**    |                                                     |                                        |
| " test ".trim()                                            |                                                     | "test"                                 |
| " test ".trimStart()                                       |                                                     | "test "                                |
| " test ".trimEnd()                                         |                                                     | " test"                                |
|                                                            |                                                     |                                        |
| **Miscellaneous string methods**                           |                                                     |                                        |
| s.concat("!")                                              | just use + operator instead                         | "Hello, world!"                        |
| "<>".repeat(5)                                             |                                                     | "<><><><><>"                           |

**Important**: Strings are immutable in JavaScript.
Methods like replace() and toUppercase() return new strings instead of modifying the string.

**Tip**: Strings can also be treated like read-only arrays: you can use `s[0]` instead of s.charAt(0)

<a id="3.3.4"></a>

#### 3.3.4 Template Literals

In ES6 and later, string literals can be delimited with backticks:

```js
let s = `hello world`;
```

These **template literals** can include arbitrary JavaScript expression.

```js
let name = "Bill";
let greeting = `Hello ${name}.`; // Hello Bill., ${name} without space is also acceptable
let two = `${1 + 1}`;
```

##### Tagged Template Literals

The value of this "tagged template literal" is the return value of the function.

ES6 has one built-in tag function: String.raw(). It returns the text
within backticks without any processing of backslash escapes:

```js
`\n`.length; // => 1
String.raw`\n`.length; // => 2
```

Note that even though the tag portion of a tagged template literal is a function,
**there are no parentheses used in its invocation**. In this very
specific case, the backtick characters replace the open and close parentheses.

See an example of usage [here](#14.5)

#### 3.3.5 Pattern Matching

[More in 11.3](#11.3)

```js
/^HTML/; // Match the letters H T M L at the start of a string
/[1-9][0-9]*/; // Match a nonzero digit, followed by any # of digits
/\bjavascript\b/i; // Match "javascript" as a word, case-insensitive, for example javsascriptiasdfasdf is a false
```

```js
let text = "testing: 1, 2, 3";
let pattern = /\d+/g; // without g, it only match the first
pattern.test(text); // => true
text.search(pattern); // => 9, position of first match
text.match(pattern); // => ["1", "2", "3"]: array of all matches
text.replace(pattern, "#"); // => "testing, #, #, #"
text.split(/\D/); // => split on non digit
// [
//   '', '',  '', '',  '', '',
//   '', '',  '', '1', '', '2',
//   '', '3'
// ]
```

### 3.4 Boolean Value

Following are all false:

```js
undefined;
null;
0;
- 0;
NaN;
("");
```

```js
// Do this
if (o)

// instead of
if (o === null)
```

**Tip**: if you need to distinguish `null` from `0` and `""`, then you should use
an explicit comparison.

### 3.5 null and undefined

`null` is usually used to indicate the absence of a value. Using the `typeof`
operator on `null` returns string "object" which indicating "no object".

`undefined` represents a deeper kind of absence. It is the value of variables
that **have not been initialized**. It is the return value of function that return nothing. It is also the value in a function when no argument is passed.

`undefined` is a predefined global constant. `typeof undefined` is `undefined`.
The `==` operator considers them to be equal. But use the strict equality `===` to
distinguish them.

None of them have any properties or methods. Using . or [] on them causes a TypeError.

The author consider `undefined` to represent a system-level, unexpected, or
error-like absence of value, and `null` to represent a program-level,
normal, or expected absence of value. He avoid using `null` and `undefined` when
he can, but if he needs to assign one of these values, he usually uses `null`. Some
programmers strive to avoid `null` entirely and use `undefined` in its place
whenever they can.

### 3.6 Symbols

```js
let strname = "string name";
let symname = Symbol("propname");
let symname2 = Symbol("propname");
typeof strname; // => "string"
typeof symname; // => "symbol"
let o = {};
o[strname] = 1;
o[symname] = 2;
o[symname2] = 3;
o[strname]; // => 1
o[symname]; // => 2
o[symname2]; // => 3

symname.toString(); // => Symbol(propname)
Symbol().toString(); // => Symbol()
```

Symbol type does not have literal syntax. The Symbol() function
never returns the same value twice, **even when called with the same argument**.
If you use symbolic property names and do not share those symbols, you
can be confident that other modules of code in your program will
not accidentally overwrite your properties.

JavaScript defines a global Symbol registry. The **Symbol.for()** function
takes a string argument and **returns a Symbol value that is
associated with the string you pass**.

```js
let s = Symbol.for("shared");
let t = Symbol.for("shared");
t === s; // => true
s.toString(); // => Symbol(shared)
Symbol.keyFor(t); // => shared
```

**Tip**: Symbol.keyFor() return string of the symbol defined with Symbol.For().
**String defined with Symbol() returns undefined with the use of Symbol.keyFor()**

```js
Symbol.keyFor(symname); // => undefined
```

### 3.7 The Global Object

The properties of this object are globally defined identifiers that
**are available to a JavaScript program**. When the JavaScript
interpreter starts (or whenever a web browser loads a **new page**), it
creates a new global object and gives it an initial set of properties
that define:

- Global constants like `undefined`, `Infinity`, and `NaN`
- Global functions like `isNaN()`, `parseInt()` [more](#3.9.2), and
  `eval()` [more](#4.12)
- Constructor functions like `Date()`, `RegExp()`, `String()`, `Object()`,
  and `Array()`(#3.9.2)
- Global objects like Math and JSON [more](#6.8)

In Node, the global object has a property named `global`.

In web browser, the Window object serves as the global object for
all JavaScript code contained in the browser window it represents.
The `window` not only defines the core global properties, it also
defines quite a few other globals that are specific to web browsers and
client-side JavaScript. [Web worker threads](#15.13) have
a different global object than the Window with which they can associated.
Code in a worker can refer to its global object as `self`

**Tip**: ES2020 finally defines `globalThis` as the standard way to
refer to the global object in any context -- implemented by all
modern browser and by Node.

### 3.8 Immutable Primitive Values and Mutable Object References

> Primitive values (`undefined`, `null`, booleans, numbers, and strings) are immutable.

```js
let s = "hello";
s[0] = "H";
s; // => hello
```

Primitive values are also compared by _value_, there is no other way that they
could be compared.

> Objects (including arrays, and functions) are mutable

```js
let o = { x: 1 };
o.x = 2;
o.x; // => 2

let a = [1, 2, 3];
a[3] = 4;
a; // => [1,2,3,4]
```

Object are **not** compared by value, two distinct objects are not equal
even if they have the same properties and values. Two object values
are the same if and only if they refer to the **same underlying object**.

```js
let a = [];
let b = a;
b[0] = 1;
a[0]; // => 1
// when an object is in a function argument, it is referenced too, thus changing value inside
// a function does modifying the object.
```

If you want to make a new copy of an object or array, you must explicitly
copy the properties of the object or the elements of the array:

```js
for (let i = 0; i < a.length; i++) {
  b[i] = a[i];
}
// In ES6, copy arrays with Array.from()
let c = Array.from(a);
```

If we want to compare two distinct objects or arrays, we must compare their
properties or elements:

```js
function equalArrays(a, b) {
  if (a === b) return true;
  if (a.length !== b.length) return false;
  for (let i = 0; i < a.length; i++) {
    if (a[i] !== b[i]) return false;
  }
  return true;
}
```

### 3.9 Type Conversions

If JavaScript wants a number, it will try to convert the value you
give it to a number (or to NaN if it cannot perform a meaningful conversion)

```js
10 + "objects"; // => "10 objects"
"7" * "4"; // => 28
"7" + "4"; // => 74
let n = 1 - "x"; // n == NaN
n + " object"; // => "NaN object"
```

| Value                     | to String           | to Number           | to Boolean |
| ------------------------- | ------------------- | ------------------- | ---------- |
| undefined                 | "undefined"         | NaN                 | false      |
| null                      | "null"              | 0                   | false      |
| true                      | "true"              | 1                   |            |
| false                     | "false"             | 0                   |            |
| ""                        |                     | 0                   | false      |
| "1.2"                     |                     | 1.2                 | true       |
| "one"                     |                     | NaN                 | true       |
| 0                         | "0"                 |                     | false      |
| -0                        | "0"                 |                     | false      |
| 1                         | "1"                 |                     | true       |
| Infinity                  | "Infinity"          |                     | true       |
| -Infinity                 | "-Infinity"         |                     | true       |
| NaN                       | "NaN"               |                     | false      |
| {} (any object)           | see [3.9.3](#3.9.3) | see [3.9.3](#3.9.3) | true       |
| [] (empty array)          | ""                  | 0                   | true       |
| [9] (one numeric element) | "9"                 | 9                   | true       |
| ["9"]                     | "9"                 | 9                   | true       |
| ['a'] (any other array)   | use join() method   | NaN                 | true       |
| function(){}              | see [here](#3.9.3)  | NaN                 | true       |

**Tip**: Leading and trailing spaces are allow.

#### 3.9.1 Conversions and Equality

`===` does not consider its operands to be equal if they are not of the same type.

```js
null == undefined; // => true
"0" == 0; // => true: string converts to a number before comparing
0 == false; // => true: boolean converts to a number before comparing
"0" == false; // => true: both operands convert to 0 before comparing
```

<a id="3.9.2"></a>

#### 3.9.2 Explicit Conversions

```js
Number("3"); // => 3
String(false); // => "false"
Boolean([]); // => true
```

Any value other than `null` and `undefined` has a `toString()` method,
and the result of this method is usually the same as that returned
by the String() function.

**Tip**: They can be invoked with `new` -- as constrctor, you'll get a "wrapper" object that behaves just like a
primitive values. But they are a historical leftover from the earliest days
of JavaScript, and there is never really any good reason to use them.

Implicit type conversion:

- If one operand of the `+` is a string, it converts the other one to a string

- The unary `+` operator converts its operand to a number.

- The unary `!` operator converts its operand to a boolean and negates it.

```js
x +
  "" + // => String(x)
  x; // => Number(x)
x - 0; // => Number(x)
!!x; // => Boolean(x)
```

The toString() method defined by the Number class accepts an optional
argument that specifies a radix, or base, for the conversion. If
you do not specify the argument, the conversion is done in base 10.

```js
let n = 17;
let binary = "0b" + n.toString(2); // => "0b10001"
let octal = "0o" + n.toString(8); // => "0o21"
let hex = "0x" + n.toString(16); // => "0x11"
```

The Number class defines three methods of number-to-string conversions:

- toFixed() converts a number to a string with a specified number of digits after the decimal point.
- toExponential() converts a number to string using exponential notation
- toPrecision() converts a number to a string with the number of significant digits you specify.

```js
let n = 123456.789;
n.toFixed(5); // => 123456.78900
n.toPrecision(2); // => 1.2e+5
n.toPrecision(10); // => 123456.7890
```

Number() conversion function parse that string as an integer or floating-point literal,
it works only for base 10 integer and does not allow trailing characters (e.g "1a" => NaN). The **parseInt()** and **parseFloat()** global functions are more flexible. They
ignore anything that after the first occurrence of non numeric characters. It interpret
a number as a hexadecimal number if a string begins with "0x" or "0X".

```js
parseInt("3 blind mice"); // => 3
parseFloat("3.14 meters"); // => 3.14
parseInt("-12.34"); // => -12
parseInt("0xFF"); // => 255
parseFloat(".1"); // => 0.1
parseFloat / Int("."); // => NaN
parseFloat / Int("$ 12.34"); // => NaN
```

parseInt() accepts an optional second argument specifying the radix (base)
of the number to be parsed:

```js
parseInt("11", 2); // => 3
parseInt("ff", 16); // => 255
```

<a id="3.9.3"></a>

#### 3.9.3 Object to Primitive Conversions

The JavaScript specification defines three fundamental algorithms
for converting objects to primitive values:

- **prefer-string**

  - This algorithm returns a primitive value, preferring a string value,
    if a conversion to string is possible.

- **prefer-number**

  - This algorithm returns a primitive value, preferring a number,
    if such a conversion is possible.

- **no-preference**
  - Classes can define their own conversions.
  - Of the built-in JavaScript types, all except Date implement this algorithm
    as **prefer-number**. The Date class implements this algorithm as **prefer-string**.

##### OBJECT-TO-BOOLEAN Conversions

All objects convert to `true`, including empty arrays and the wrapper object
`new Boolean(false)`.

##### OBJECT-TO-STRING Conversions

When an object needs to be converted to a string, JavaScript first
converts it to a primitive using the prefer-string algorithm,
then converts the resulting primitive value to a string.

It happens when:

```js
String(<your object>)
// interpolate objects into template literals
let a = `hello ${<your object>}`
```

##### OBJECT-TO-NUMBER Conversions

Same sequence with OBJECT-TO-STRING Conversions.

It happens when:

```js
Number(<your object>)
1 * <your object>
```

##### SPECIAL CASE OPERATOR Conversions

With `+` operator with string or number operand, object operand is converted to the type of other operand, either performing concatenation or summing.

With "==" and "!=" operator. If one operand is an object and the other
is a primitive value, these operators convert the object to primitive using
the no-preference algorithm and comparing them.

`<`, `<=`, `>`, `>=`, if either operand is an object, it is
converted to a primitive value using _prefer-number algorithm_.

##### The toString() and valueOf() Methods

All objects inherit two conversion methods that are used
by object-to primitive conversions.

`toString()` return a string representation of the object.

```js
{{x: 1, y: 2}}.toString() // => "[object Object]"
(function abc(x){f(x);}).toString() // => function abc(x){f(x);}
(new Date(2020,0,1)).toString(); // 'Wed Jan 01 2020 00:00:00 GMT+0800 (Malaysia Time)'
/^abc/.toString() // => /^abc/
```

`valueOf()` convert an object to a primitive value that represents the object,
if any such primitive value exists. Objects are compound values, and
most objects cannot really be represented by a single primitive values, so
the default `valueOf()` method **simply returns the object itself rather than
returning a primitive**.

Wrapper classes such as String, Number, and Boolean define valueOf() methods
that simply return the wrapped primitive value.

```js
let a = new Number(5);
a.valueOf(); // => 5
```

The Date class defines a valueOf() method that returns the date
in its internal representation: the number of milliseconds
since January 1, 1970;

```js
let e = new Date();
e.valueOf(); // => 1702132817168
```

##### OBJECT-TO-PRIMITIVE Conversions Algorithms

- The `prefer-string` algorithm first tries the `toString()` method. If
  the method is defined and returns a primitive value, then JavaScript uses
  that primitive value (even if it is not a string).

  - If `toString()` does not exist or if it returns an object, then JavaScript
    tries the `valueOf()` method. If that method exists and returns primitive value,
    then JavaScript uses that value.
  - Otherwise, the conversion fails with a TypeError

- The `prefer-number` algorithm works like the prefer-string algorithm, except
  it tries `valueOf()` first then `toString()`.

- The `no-preference algorithm` depends on the class of the object being
  converted. If the object is a Date object, then JavaScript uses the prefer-string
  algorithm. Other object: `prefer-number` algorithm.

The rules described here are true for all built-in JavaScript types and
are the default rules for any classes you define yourself.
[Symbol.toPrimitive](#14.4.7) explains how you can define your own object-to primitive
conversion algorithms for the classes you define.

Explanation:

```js
Number([]); // => 0
Number([99]); // => 99
```

When using `object to number` conversion, a `prefer-number` algorithm was used.
Array's `valueOf()` method returns an object, thus it falls back to `toString()`.
The `toString()` of [] is '', while [99].toString() is '99'. And that is why
Number([]) is 0 and Number([99]) is 99.

### 3.10 Variable Declaration and Assignment

**Tip**: An attempt to evaluate a nonexistent variable throws a ReferenceError

```js
console.log(xx);
Uncaught ReferenceError: xx is not defined
```

#### 3.10.1 Declarations with `let` and `const`

It is a good programming practice to assign an initial value to your
variables when you declare them, when this is possible:

```js
let message = "hello";
let i = 0,
  j = 0;
let x = 2,
  y = x * x;

const H0 = 74;
```

**Tip**: It is a common convention to declare constants using
names with all capital letters such as H0 or HTTP_NOT_FOUND.

We may use `let` to declare the loop variable

```js
for (let i = 0, len = date.length; i < len; i++) console.log(data[i]);
for (let datum of data) console.log(datum);
for (let property in object) console.log(property);
```

You may use `const` to declare the loop variable for `for/in` and `for/of` loops,
as long as **the body of the loop does not reassign a new value**.

```js
for (const datum of data) console.log(datum);
```

##### Variable and Constant Scope

Variable and constants declared with `let` and `const` are block scoped.
This means that they only defined within the block of code in which
the `let` or `const` statement appears. JavaScript class
and function definitions are blocks, and so are the bodies of `if/else`,
`while`, `for`.

When a declaration appears at the top level, it is a `global variable` or constant.
In traditional client-side JavaScript, the scope of global variable is the
HTML document in which it is defined. That is: if one \<script\> declares a
global variable or constant, that variable or constant is defined in
all of the \<script\> elements in that document. In Node and in client-side
JavaScript modules ([Chapter 10](#10)), **the scope of a global variable
is the file that it is defined in**: the use of `var` will not be visible in another
script if the declaration was in a module script.

```html
<script>let a = 1; var b = 2</script>
<script>console.log(a); console.log(b)</script>
<script>console.log(globalThis.a); console.log(globalThis.b)</script>
```

Result:

```
1
2
undefined
2
```

##### Repeated Declarations

A syntax error:

- Same name with more than one `let` or `const`

##### Declarations And Types

A JavaScript variable can hold a value of any type.

```js
let i = 10;
i = "ten";
```

#### 3.10.2 Variable Declarations with var

In versions of JavaScript before ES6, the only way to declare a
variable is with the `var` keyword and there is no way to declare constants.

Although `var` and `let` have the same syntax, there are important
differences in the way they work:

- `var` do not have block scope. They are scoped to the body of the containing function
  **no matter how deeply nested they are inside that function**.

  ```js
  var v = 1;
  if (true) var v = 2;
  v; // => 2
  ```

- `var` outside of a function body declares a global variable, and it is implemented
  as properties of the global object. It cannot be deleted by [`delete` operator](#4.13.4). Global variables and constants declared with
  `let` and `const` are not properties of the global object.

  ```js
  var x = 2;
  globalThis.x; // => 2
  ```

- It is legal to declare the same variable multiple times with `var`

- When a variable is declared with `var`, the declaration is lifted up (or "hoisted")
  to the top of the enclosing function. So the variables declared with `var`
  can be used, without error, anywhere in the enclosing function. **If the initialization
  code has not run yet, then the value of the variable may be `undefined`**. `let` corrects it, you will get an actual error.

  ```js
  console.log(hello);
  var hello = 1; // => undefined
  ```

---

> **_Note_**: Using Undeclared Variable

In [strict mode](#5.6.3), attempting to use an undeclared variable cause
a reference error. However outside of strict mode, you'll end up creating
a new **global variable**, with `var`, which is the property of globalThis. However
it can be deleted with the `delete` operator.

**This is one of the best reasons for using strict mode!**

---

#### 3.10.3 Destructuring Assignment

ES6.

In a destruturing assignment, the value on the righthand side of the
equals sign is an array or object (a "structured" value), and the
lefthand side specifies one or more variable names using a syntax
that mimics array and object literal syntax.

```js
let [x, y] = [1, 2, 3]; // => x = 1, y = 2
[x, y] = [x + 1, y + 1];
[x, y] = [y, x];
[x, y]; // => [3,2]
[, z] = [1, 99];
z; // => 99

let [a, b, c, d] = [1, 2];
[a, b, c, d]; // => [ 1, 2, undefined, undefined ]

// Three dots to collect all unused or remaining values into single variable
let [a, ...b] = [1, 2, 3, 4]; // b == [2,3,4]

// Nested arrays
let [a, [b, c]] = [1, [2, 2.5], 3]; // a == 1, b == 2, c == 2.5
```

```js
function toPolar(x, y) {
  return [Math.sqrt(x * x + y * y), Math.atan2(y, x)];
}
let [r, theta] = toPolar(1.0, 1.0);
```

```js
let o = { x: 1, y: 2 };
for (const [name, value] of Object.entries(o)) {
  console.log(name, value); // Prints "x 1" and "y 2"
}
```

It does not actually required an array! You can use **any [iterable object](#12)
on the righthand side of the assignment; any object that can be used with a
`for/of` loop can also be destructed**:

```js
let [first, ...rest] = "Hello"; // first == "H", rest == ["e", "l", "l", "o"]
```

The righthand side is an object value:

```js
let transparent = { r: 0.0, g: 0.0, b: 0.0, a: 1.0 };
let { r, g, b } = transparent; // r == 0.0; g == 0.0; b == 0.0
// you may make own identifiers:
let { r: red, g: green, b: blue } = transparent; // red, green, blue == 0.0
```

Nested object

```js
let points = [
  { x: 1, y: 2 },
  { x: 3, y: 4 },
];
let [{ x: x1, y: y1 }, { x: x2, y: y2 }] = points;
```

Array of objects:

```js
let points = {p1: [1,2], p2: [3,4]};
let { p1: [x1, y1], p2: [x2, y2]};
```

**Complex destructing syntax like this can be hard to write and
hard to read**, you may better off just writing your assignments:
`let x1 = points.p1[0]`

## 4 - Expressions And Operators

### 4.2 Object and Array Initializers

Undefined elements can be included in an array literal by
simply omitting a value between commas.

```js
let sparseArray = [1, , , , 5];
```

**Tip**: A single trailing comma does not create an undefined element.

Object Initializers:

```js
let rectangle = {
  upperLeft: {x: 2, y:2};
  lowerRight: {x:4, y:5};
}
```

### 4.3 Function Definition Expressions

```js
let square = function (x) {
  return x * x;
};
```

A function definition expression can also include a name for the function.

Functions can also be defined using a function statement rather than a
function expression.

And in ES6 and later, function expressions can use a compact new
"arrow function" syntax. [Complete details on function definition](#8)

### 4.4 Property Access Expressions

JavaScript defines two syntaxes for property access:

```js
expression.identifier;
expression[expression];
```

With either type of property access expression, the expression
before the . or [ is first evaluated. If the value is `null` or `undefined`,
the expression throws a TypeError.

If the named property does not exist, then the value of the
property access expression in `undefined`.

If the property name includes spaces or punctuation characters,
or when it is a number (for arrays), you must use the square bracket notation.

#### 4.4.1 Conditional Property Access

ES2020

```js
expression?.identifier;
expression?.[expression];
```

This guards against TypeError thrown when the first value is `null` or `undefined`.

```js
let a = { b: null };
a.b?.c.d; // => undefined
```

Note that the side effect will not occur in `a?.[b][c]` if `a` is not defined.

```js
let a;
let index = 0;
try {
  a[index++]; // Throws TypeError
} catch (e) {
  index; // => 1
}
a?.[index++];
index; // => 1
```

### 4.5 Invocation Expressions

An invocation expression is JavaScript's syntax for calling a function or method.

```js
f(0);
```

If the value of the function expression is not a function, a TypeError is thrown.

#### 4.5.1 Conditional Invocation

In ES2020, you can also invoke a function using `?.()`

Before ES2020, if you want to write a method like `sort()` that
takes an optional function argument, you would:

```js
function square(x, log) {
  if (log) {
    log(x);
  }
  ...
}
```

With conditional invocation syntax:

```js
function square(x, log) {
  log?.(x);
  ...
}
```

Same with conditional property access, none of the argument expressions within the
parentheses are evaluated if the value to the left of ? is `null` or `undefined`.

### 4.6 Object Creation Expressions

An `object creation expression` creates a new object and invokes
a function (called a constructor) to initialize the properties of
that object.

```js
new Object();
new Point(2, 3);
```

If no arguments are passed to the constructor function in an object
creation expression, the empty pair of parentheses can be omitted:

```
new Object;
new Date;
```

### 4.7 Operator Overview

The table below summarizes the operators and serves as a convenient
reference.

- `A` - operator associativity.

  - `L` (left-to-right)
  - `R` (right-to-left)

- `N` - the number of operands

- `Types`
  - expected types of the operands
  - (after the -> symbol) the result type for the operator.

| Operator            | Operation                        | A   | N   | Types               |
| ------------------- | -------------------------------- | --- | --- | ------------------- |
| ++                  | Pre- or post-increment           | R   | 1   | lval -> num         |
| --                  | Pre- or post-decrement           | R   | 1   | lval -> num         |
| -                   | Negate number                    | R   | 1   | num -> num          |
| +                   | Convert to number                | R   | 1   | any -> num          |
| ~                   | Invert bits                      | R   | 1   | int -> int          |
| !                   | Invert boolean value             | R   | 1   | bool -> bool        |
| delete              | Remove a property                | R   | 1   | lval -> bool        |
| typeof              | Determine type of operand        | R   | 1   | any -> str          |
| void                | Return undefined value           | R   | 1   | any -> undef        |
| \*\*                | Exponentiate                     | R   | 2   | num,num -> num      |
| \*,/,%              | Multiply, divided, remainder     | L   | 2   | num,num -> num      |
| +,-                 | Add, substract                   | L   | 2   | num,num -> num      |
| +                   | Concatenate strings              | L   | 2   | str,str -> str      |
| <<                  | Shift left                       | L   | 2   | int,int -> int      |
| >>                  | Shift right with sign extension  | L   | 2   | int,int -> int      |
| >>>                 | Shift right with zero extension  | L   | 2   | int,int -> int      |
| <, <=, >, >=        | Compare in numeric order         | L   | 2   | num,num -> bool     |
| <, <=, >, >=        | Compare in alphabetical order    | L   | 2   | str,str -> bool     |
| instanceof          | Test object class                | L   | 2   | obj,func -> bool    |
| in                  | Test whether property exists     | L   | 2   | any,obj -> bool     |
| ==                  | Test for non-strict equality     | L   | 2   | any,any -> bool     |
| !=                  | Test for non-strict inequality   | L   | 2   | any,any -> bool     |
| ===                 | Test for strict equality         | L   | 2   | any,any -> bool     |
| !==                 | Test for strict inequality       | L   | 2   | any,any -> bool     |
| &                   | Compute bitwise AND              | L   | 2   | int,int -> int      |
| ^                   | Compute bitwise XOR              | L   | 2   | int,int -> int      |
| \|                  | Compute bitwise OR               | L   | 2   | int,int -> int      |
| &&                  | Compute logical AND              | L   | 2   | any,any -> any      |
| \|\|                | Compute logical OR               | L   | 2   | any,any -> any      |
| &&                  | Compute logical AND              | L   | 2   | any,any -> any      |
| ??                  | Choose 1st defined operand       | L   | 2   | any,any -> any      |
| ?:                  | Choose 2nd or 3rd operand        | R   | 3   | bool,any,any -> any |
| =                   | Assign to a variable or property | R   | 2   | lval,any -> any     |
| \*_=, _=, /=, %=,   | Operate and assign               | R   | 2   | lval,any -> any     |
| +=, -=, &=, ^=,\|=, |                                  |     |     |                     |
| <<=, >>=, >>>=      |                                  |     |     |                     |
| ,                   | Discard 1st operand, return 2nd  | L   | 2   | any, any -> any     |

#### 4.7.2 Operand and Result Type

JavaScript operators usually convert the type of their operands as needed. The
expression "3" \* "5" is legal because JavaScript can convert the
operands to numbers.

Some operators expect an operand of type `lval`. lvalue is a
historical term that means "an expression that can legally appear on the
left side of an assignment expression". In JavaScript, variables,
properties of objects, and elements of arrays are lvalues.

#### 4.7.3 Operator Side Effects

Side effects mean evaluation of expressions may affect the result of future
evaluations. For example, if you assign a value to a variable or property,
that changes the value of any expression that uses that variable or property.

#### 4.7.4 Operator Precedence

It controls the order in which operations are performed. In the table above,
are arranged in order from high precedence to low precedence.

```js
let x = 1;
x = ++x + x++;
x; // => 4
```

The property access and invocation expressions have higher precedence than
any of the operators listed in the table.

**In practice, if you are at all unsure about the precedence of your
operators, the simplest thing to do is to use parentheses to make
the evaluation order explicit**.

#### 4.7.5 Operator Associativity

Left-to-right associativity:

```js
w = x - y - z;
// is
// w = ((x - y) - z);
```

Other example (right-to-left):

```js
y = a ** (b ** c); // => y = (a ** (b**c))
x = ~-y; // => x = ~(-y);
w = x = y = z; // => w = (x = (y = z))
q = a ? b : c ? d : e ? f : g; // => q = a?b:(c?d:(e?f:g))
```

#### 4.7.6 Order of Evaluation

JavaScript always evaluates expressions in strictly left-to-right order.
For example, in the `w = x + y * z`, the subexpression `w` is
evaluated first, followed by x, y and z. Then the values of y and z are
multiplied, added to the value of x, and assigned to the variable or property
specified by expression w.

Adding the parentheses to the expressions can change the relative
order of the multiplication, addition, and assignments,
but not the left-to-right order or evaluation.

### 4.8 Arithmetic Expressions

#### 4.8.1 The + Operator

The conversion rules for + give priority to string concatenation:
if either operand is a string or an object that converts to a string,
the other operand is converted to a string and concatenation is performed.

```js
"1" + 2; // => "12"
true + true; // => 2
2 + null; // => 2
2 + undefined; // => NaN
```

#### 4.8.3 Bitwise Operators

NaN, Infinity, and -Infinity all convert to 0 when using as
operands of these bitwise operators.

```js
7 << 2; // => 28
7 >> 1; // => 3
-7 >> 1; // => -4
-1 >> 4; // => -1
-1 >>> 4; // => 0x0FFFFFFF
```

### 4.9 Relational Expressions

#### 4.9.1 Equality and Inequality Operators

> **Strict Equality**

- If the two values have different types, they are not equal.

- If both values are `null` or both values are `undefined`, they are equal.

- If both values are the boolean value `true` or both are `false`, they are equal.

- If one or both values is NaN, they are not equal (use isNaN() instead.)

- If one value is 0 and the other is -0. They are equal.

- If both values refer to same object, array or function, they are equal.

> **Equality with Type Conversion**

- If the two values have the same type, test them for strict equality as described
  previously.

- If the two values do not have same type:

  - If one value is `null` and the other is `undefined`, they are equal.

  - If one value is number and the other is string, convert the string to number
    and try the comparison again.

  - if either value is `true`, convert it to 1 and try the comparison again. If either
    value is `false`, convert it to 0 and try the comparison again.

  - If one value is an object and the other is a number or string,
    convert the object to a primitive using the algorithm and try the comparison.
    The built-in classes of core JavaScript attempt valueOf() conversion before toString()
    conversion; except for the Date class.

  - any other combinations of values are not equal. e.g `"true" == true`, `{} == {}`

#### 4.9.2 Comparison Operators

**Tip**: Comparison favours number over string

Comparison and conversion occur as follows:

- If either operand evaluates to an object, that object is converted to a primitive value.

  ```js
  {} >= {a:1} // => true, "[object Object]" == "[object Object]"
  ```

- If, after object-to-primitive conversion, at least one operand is
  not a string, both operands are converted to numbers and compared numerically.
  `0` and `-0` are considered equal.

  ```js
  "-0" >= "0"; // => false
  "-0" >= 0; // => true
  ```

- If either operand is NaN, the comparison operator always returns false.

  ```js
  "one" < 3; // => false, NaN < 3
  ```

- Although the arithmetic operators do not allow BigInt values to be mixed
  with regular numbers, the comparison operators do allow comparisons between
  numbers and BigInts.

For a more robust string-comparison algorithm, try the
`String.localeCompare()` method, which also takes locale-specific definitions
of alphabetical order into account.

```js
"a" > "A"; // => true
"a".localeCompare("A"); // => -1
```

For case-insensitive comparisons, you can convert the strings to all lowercase
or all uppercase using `String.toLowerCase()` or `String.toUpperCase()`.

#### 4.9.3 The `in` Operator

The `in` operator expects a left-side operand that is a string, symbol, or
value that can be converted to a string. It expects a
right-side operand that is an object.

```js
let point = { x: 1, y: 1, w: undefined, w1: null };
"x" in point; // => true
"z" in point; // => false
"w" in point; // => true
"w1" in point; // => true
"toString" in point; // => true
```

However, when a array as a right-side operand, it expect left-side operand a index.

```js
let data = [7, , 9];
"0" in data; // => true;
0 in data; // => true, it is converted to string
3 in data; // => false, no element 3
1 in data; // => false
```

<a id="4.9.4"></a>

#### 4.9.4 The `instanceof` Operator

It expects a left-side operand that is an object and a right-side operand
that identifies a class of objects.

```js
let d = new Date();
d instanceOf Date // => true
d instanceOf Object // => true
d instanceOf Number // => false

let a = [1,2,3];
a instanceof Array // => true
a instanceof Object // => true

let r = /hello/;
r instanceof RegExp // => true
```

To evaluate expression o instanceof f, JavaScript evaluates f.prototype,
and then looks for that value in the prototype chain of o.
If it finds it, then o is an instance of f.

### 4.10 Logical Expressions

#### 4.10.1 Logical AND (&&)

**Tip**: The below is not true and false.

```js
let o = { x: 1 };
let p = null;
o && o.x; // => 1
p && p.x; // +> null
```

The following two lines of JavaScript code have equivalent effects:

```js
if (a === b) stop();
// (a === b) && stop();
```

#### 4.10.2 Logical OR (||)

An idiomatic usage of this operator is to **select the first truthy value** in
a set of alternative:

```js
let max = maxWidth || preferences.maxWidth || 500;
```

**Note that if 0 is the legal value for `maxWidth`, then this code will not work
correctly**. (Use ?? instead)

Prior to ES6, this idiom is often used in functions to supply default values
for parameters:

```js
function copy(o, p) {
  p = p || {};
}
```

In ES6 and later, this trick is no longer needed because the default parameter
value could simply be written in the function definition itself:

`function copy(o, p = {}) {...}`

#### 4.10.3 Logical NOT (!)

You can convert any value x to its equivalent boolean value by applying
this operator twice: `!!x`

### 4.11 Assignment Expressions

#### 4.11.1 Assignment with Operation

```js
a op= b
a = a op b
```

In the first line, the expression `a` is evaluated once. In the second,
it is evaluated twice. The following two assignments, for example,
are not the same:

```js
let data = [1,2,3,4,5]
let i = 0
data[i++] *= 2; // data[1] *= 2
data[i++] = data[i++] * 2; //  data[0] = data[1] * 2
```

<a id="4.12"></a>

### 4.12 Evaluation Expressions

Dynamic evaluation of strings of source code is a powerful language
feature that is almost never necessary in practice. In particular,
`eval()` can be a security hole, and you should never pass any
string derived from user input to `eval()`. Because of these
security issues, **some web servers use the HTTP "Content-Security-Policy" header
to disable `eval()` for an entire website**.

Modern JavaScript interpreters perform a lot of code analysis and
optimization. If a function calls eval(), the interpreter cannot optimize
that function.

The problem with defining `eval()` as a function is that it can be given other names:

```js
let f = eval;
let g = f;
```

If this is allowed, then the interpreter can't know for sure
which functions call `eval()`, so it cannot optimize aggressively.

We'll learn about [restrictions placed on `eval()` to make it more operator-like](#4.12.3)

#### 4.12.1 eval()

`eval()` expects one argument. If you pass any value other than a string,
it simply returns that value. If you pass a string, **it attempts to parse the string as
JavaScript code**, throwing a SyntaxError if it fails. If it
successfully parses the string, then it evaluates the code and returns
the value of the last expression or statement in the string or `undefined`
if the last expression or statement had no value.

**If the evaluated string use `let` or `const`, the variable or constant declared
will be local to the evaluation and will not be defined in the calling environment**.

```js
let x = 11;
eval("function f() {return x+1;}");
console.log(f()); // => 12

eval("var v = 11")
console.log(v) // => 11

eval("let v = 11")
console.log(v) // ReferenceError: v is not defined
```

#### 4.12.2 Global eval()

```js
const geval = eval;

let x = "global",
  y = "global";
function f() {
  let x = "local";
  eval("x += 'changed'");
  return x;
}

function g() {
  let y = "local";
  geval("y += 'changed';");
  return y;
}

console.log(f(), x); // Local variable changed: prints "localchanged global"
console.log(g(), y); // Global variable changed: prints "local globalchanged"
```

If a script defines an alias for `eval()` and then calls that function by
another name, it should evaluate the string as if it were top-level global code.
The evaluated code may define new global variables or global functions.

Direct calls to `eval()` use the variable environment of the calling context.
Any other call uses the global object as its variable environment and cannot read,
write, or define local variables or functions.

Notice that the ability to do a global eval is not just an accommodation
to the needs of the optimizer; it is actually a tremendously useful
feature that allows you to execute strings of code as if they
were independent, top-level scripts.

**As noted at the beginning of this section, it is rare to truly need
to evaluate a string of code. But if you do find it necessary,
you are more likely to want to do a global eval than a local eval**.

<a id="4.12.3"></a>

#### 4.12.3 Strict eval()

In strict mode ("use strict" directive), evaluated code can query and set local variables, **but it
cannot define new variables or functions in the local scope**.

In strict mode, you are not allowed to overwrite
the eval() function with a new value. And you are
not allowed to declare a variable, function parameter,
or catch block parameter with the name eval.

### 4.13 Miscellaneous Operators

#### 4.13.1 The Conditional Operator (?: )

It is sometimes actually called the ternary operator.

```js
x > 0 ? x : -x;

greeting = "hello " + (username ? username : "there");
```

#### 4.13.2 First-Defined (??)

If its left operand is not `null` and not `undefined`, it returns that value.
Otherwise, it returns the value of the right operand. It is equivalent to:

```js
a !== null && a !== undefined ? a : b;
```

?? is a useful alternative to || when you want to select the first
defined operand rather than the first truthy operand.

```js
let max = maxWidth ?? preferences.maxWidth ?? 500;
```

If you use it in an expression with either && or ||, you must use
explicit parentheses to specify which operation you want to perform first,
else it has SyntaxError

#### 4.13.3 The typeof Operator

| x                      | typeof x    |
| ---------------------- | ----------- |
| undefined              | "undefined" |
| null                   | "object"    |
| true or false          | "boolean"   |
| any number or NaN      | "number"    |
| any BigInt             | "bigint"    |
| any string             | "string"    |
| any symbol             | "symbol"    |
| any function           | "function"  |
| any nonfunction object | "object"    |

```js
typeof value === "string" ? "'" + value + "'" : value.toString();
```

To distinguish one class of object from another, you must use other
techniques, such as [`instanceof`](#4.9.4-the-instanceof-operator), [`class` attribute](#14.4.3), or the [`constructor` property](#9.2.2)

<a id="4.13.4"></a>

#### 4.13.4 The `delete` Operator

It is typically used for its property deletion side effect and not for the
value it returns.

```js
let o { x: 1, y: 2};
delete o.x;
"x" in o; // => false

let a = [1,2,3];
delete a[2];
2 in a; // => false
a.length // => 3: the array length doesn't change
```

**Note that a deleted property or array element is not merely set to the
`undefined` value.**

`delete` expects its operand to be an lvalue. If it is not an lvalue,
the operator takes no action and return true. It also returns true if delete
a non existent property.

```js
delete 1; // => true
```

Otherwise, `delete` returns true if it successfully deletes the specified lvalue.
Not all properties can be deleted: [non-configurable properties](#14.1) are
immune from deletion.

```js
delete Object.prototype;
```

In strict mode, a SyntaxError if its operand is an unqualified identifier such as
variable, function, or function parameter: it only works when the operand is a property access expression. Strict mode also
specifies that `delete` raises a TypeError if asked to delete any non-configurable
property. Outside of strict mode, no exception occurs, `delete` simply returns false.

#### 4.13.5 The `await` Operator

`ES2017`

It expects a Promise object. The value of the `await` operator is the fulfillment value of the Promise object.
Importantly `await` is only legal within functions that have been declared
asynchronous with the `async` keyword.

[See here for full details](#13)

#### 4.13.6 The `void` Operator

This operator is unusual and infrequently used; it evaluates its operand,
then discards the value and returns `undefined`.

```js
let counter = 0;
const increment = () => void counter++;
increment(); // => undefined
counter; // => 1
```

#### 4.13.7 The `comma` Operator (,)

```js
(i = 0), (j = 1), (k = 2);
```

Equivalent to:

```js
i = 0;
j = 1;
k = 2;
```

The comma operator is commonly used with a `for` loop.

```js
for (let i = 0, j = 10; i < j; i++, j--) {
  console.log(i + j);
}
```

## 5 - Statements

### 5.4 Loops

#### 5.4.4 `for/of`

`ES6`

The for/of loop works with **_iterable objects_**, such as
arrays, strings, sets, and maps.

```js
// array
let data = [1, 2, 3, 4],
  sum = 0;
for (let e of data) {
  sum += element;
}

// string
for (let letter of "missisippi") {
  ...
}

// set
let text = "Na na na na na na na na Batman!"
let wordSet = new Set(text.split(" "));
let unique = [];
for (let word of wordSet) {
  unique.push(word);
}
unique // => [ 'Na', 'na', 'Batman!' ]

// map
let m = new Map([[1, "one"]]);
for (let [k, v] of m) {
  k // => 1
  v // => 'one'
}
```

**Important**: The strings are iterated by Unicode codepoint, not by UTF-16.

**Arrays are iterated "live", if we modify the preceding code
by adding the line data.push(sum); inside the loop body. Then
it will be infinite loop**

> Iterate the properties of an object, Object.keys() returns array of property names.
> Object.values() returns array of values. Object.entries() returns array of [k, v]

```js
let o = { x: 1, y: 2, z: 3 };
let keys = "";
for (let k of Object.keys(o)) {
  keys += k;
}
keys; // => "xyz"
```

##### Asynchronous Iteration with For/await

`ES2018`

```js
async function printStream(stream) {
  for await (let chunk of stream) {
    console.log(chunk);
  }
}
```

#### 5.4.5 `for/in`

A `for/in` loop works with object after `in`.

```
for (variable in object)
  statement
```

If the JavaScript interpreter evaluates to `null` or `undefined`, the
interpreter skips the loop and moves on to the next statement.
The interpreter now executes the body of the loop once for **each
enumerable property** of the object. Before each iteration,
however, the interpreter evaluates the variable expression and
assigns the name of the property (a string value) to it.

You can use code like the following to copy the names of all object properties
into an array (for/of loops can be applied this too):

```js
let o = { x: 1, y: 2, z: 3 };
let a = [],
  i = 0;
for (a[i++] in o /*empty*/);
a; // => [ 'x', 'y', 'z' ]
```

Array indexes are object properties that can be enumerated with a for/in
loop.

```js
for (let i in a) console.log(i); // 0 1 2
```

**When working with arrays, you almost always want to use for/of instead of for/in.**

The `for/in` loop does not actually enumerate all properties of an object. It:

- does not enumerate properties whose name are symbols.

- of the properties whose name are strings, it only loops over the [`enumerate` properties](#14.1). The various built-in methods defined by core JavaScript are not enumerable. e.g toString() method is not enumerable. All properties and methods defined
  by your code are enumerable, by default.

- enumerable inherited properties are also enumerated by the `for/in` loop. For this reason,
  many programmers prefer to use a `for/of` loop with Object.keys() instead of `for/in` loop

If the body of a `for/in` loop deletes a property that has not yet been
enumerated, the property will not be enumerated. If the body of the
loop defines new properties on the project, those properties may or may not be
enumerated. See [here](#6.6.1) for more information on the order in which
`for/in` enumerates the properties of an object.

### 5.5 Jumps

#### 5.5.1 Labeled Statements

```js
identifier: statement;
```

By giving a loop a name, you can use `break` and `continue` statements inside
the body of the loop to exit the loop or to jump directly to the top of the loop
to begin then next iteration.

```js
mainloop: while (token !== null) {
  // ...
  continue mainloop;
  // ...
}
```

#### 5.5.5 `yield`

The `yield` statement is much like the `return` statement but is used
only in `ES6` generator functions to produce the next value in the generated
sequence of values without actually returning:

```js
function* range(from, to) {
  for (let i = from; i <= to; i++) {
    yield i;
  }
}
```

### 5.6 Miscellaneous Statements

#### 5.6.1 `with`

Runs a block of code as if the properties of a specified object
were variables in scope for that code:

```
with (object)
  statement
```

This statement creates a temporary scope with the properties of
object as variables and then executes statement within that scope.

The `with` statement is forbidden in strict mode and should be
considered deprecated in non-strict mode. JavaScript code
that uses `with` is difficult to optimize and is likely to run
significantly more slowly than the equivalent code written without `with` statement.

The common use of the `with` statement is to make it easier to work
with deeply nested object hierarchies:

```js
// from
document.forms[0].name.value = 0;
document.forms[0].address.value = 0;
document.forms[0].email.value = 0;

// to
with (document.forms[0]) {
  name.value = "";
  address.value = "";
  email.value = "";
}
```

Of course, to avoid the `with` statement, write this:

```js
let f = document.forms[0];
f.name.value = "";
f.address.value = "";
f.email.value = "";
```

**Tip**: Note that if you use `const` or `let` or `var` to declare a variable
or constant within the body of a `with` statement, it creates an ordinary
variable and does not define a new property within the specified object.

#### 5.6.2 `debugger`

In practice, this statement acts like a breakpoint: execution of JavaScript code
stops, and you can use the debugger to print variables' values, examine the call
stack, and so on.

```js
function f(o) {
  if (o === undefined) debugger;
  // ...
}
```

<a id="5.6.3"></a>

#### 5.6.3 `use strict`

"use strict" is a directive introduced in `ES5`. Directive are not
statements.

- It does not include any language keywords: the directive is just an expression statement
  that consists of a special string literal (in single or double quote).

- It can appear only at the start of a script of at the start
  of a function body, before any real statements have appeared

The top-level (non-function) code of a script is strict code if the script
has a "use strict" directive.

A function is strict code if it is defined within strict code or if it
has use a "use strict" directive. Code passed to eval() method
is strict code if eval() is called from strict code or
if the string of code includes a "use strict" directive.

```js
"use strict";

a = 1; // => Reference Error
```

In addition to code explicitly declared to be strict,
**any code in a class body or in an ES6 module is automatically
strict code**. Which you will never need to use an explicit "use strict"
directive.

Strict mode is a restricted subset of the language that fixes important
language deficiencies and provides stronger error checking and
increased security.

The differences between strict mode and non-strict mode:

- `with` statement is not allowed in strict mode.

- all variables must be declared in strict mode.

- in strict mode, function invoked as functions have `this` value of undefined. In non-strict mode, `this` is global object.

- in strict mode, when a function is invoked with call() or apply(), the `this`
  value is exactly the value passed as the first argument to `call()` or `apply()`

- in strict mode, assignments to non writable properties
  and attempts to create new properties on non-extensible objects
  throw a TypeError. (fail silently in non-strict mode.)

- in strict mode, code passed to eval() cannot declare variables or define
  functions in the caller's scope as it can in non-strict mode (var). Instead,
  variable and function definitions live in a new scope created for the eval().

- in strict mode, the Arguments object in a function holds a static copy of the
  values passed to the function. In non-strict mode, elements of array and named
  function parameters both refer to the same value.

- In strict mode, a SyntaxError is thrown if the `delete` operator is followed
  by an unqualified identifier such as a variable, function, or function parameter.

- In strict mode, an attempt to delete a non-configurable property
  throws a TypeError

- In strict mode, it is a syntax error for an object literal to define
  two or more properties by the same name.

- In strict mode, it is a syntax error for a function declaration to have two
  or more parameters with the same name.

- In strict mode, octal integer literals (beginning with a 0 that is not followed
  by an x) are not allowed.

- In strict mode, the identifiers `eval` and `arguments` are treated like keywords,
  you are not allowed to change their value.

- In strict mode, the ability to examine the call stack is restricted. `arguments.caller`
  and `arguments.callee` both throw a TypeError within a strict mode function.

### 5.7 Declarations

#### 5.7.2 function

Function declarations are "hoisted", they had all been moved up
to the top of whatever scope they are defined within. Below code is
valid.

```js
hello();

function hello() {
  console.log("hello");
}
```

#### 5.7.3 Class

Unlike functions, class declarations are not hoisted, and
you cannot use a class declared this way in code that appears before the
declaration.

#### 5.7.4 import and export

They are used together to make values defined in one
module of JavaScript code available in another module.
A module is a file of JavaScript code with its own global namespace,
completely independent of all other modules.

## 6 - Objects

Each property has three property attributes:

- The `writable` attribute specifies whether the value of the property
  can be set

- The `enumerable` attribute specifies whether the property name
  is returned by a for/in loop.

- The `configurable` attribute specifies whether the property can be
  deleted and whether its attributes can be altered.

Many of JavaScript's built-in objects have properties that are read-only,
non-enumerable, or non-configurable. By default, all properties of the
objects you create are writable, enumerable, and configurable.

### 6.2 Creating Objects

Objects can be created

- with object literals

- with the `new` keyword,

- and with the `Object.create()` function.

#### 6.2.1 Object Literals

```js
let empty = {};
let point = { x: 0, y: 0 };
let book = {
  "main title": "JavaScript",
  author: {
    firstname: "David",
  },
};
```

The trailing comma following the last property in an object literal is legal,
and some programming styles encourage the use of these trailing
commas so you're less likely to cause a syntax error if you add a
new property at the end of the object literal at some later time.

#### 6.2.2 Creating Objects with `new`

The `new` operator creates and initializes a new object. The `new`
keyword must be followed by a function invocation. **A function used
in this way is called a _constructor_ and serves to initialize a
newly created object**. JavaScript includes constructors
for its built-in types:

```js
let o = new Object(); // {}
let a = new Array(); // []
let d = new Date(); // Date object representing the current time
let r = new Map();
```

#### 6.2.3 Prototypes

Almost every JavaScript object has a second JavaScript object associated with it.
This second object is known as _prototype_, and the first object **inherits** properties
from the prototype.

All objects created by object literals have the same prototype object,
**we can refer to this prototype object in JavaScript code as `Object.prototype`**

Objects created using the `new` keyword and a constructor invocation
use the value of the `prototype` property of the constructor function as their
prototype. So the object created by `new Object()` inherits from
Object.prototype. Similarly, `new Array()` uses Array.prototype as its prototype,
and so Date.

**Remember**: almost all objects have a prototype, **but only a relatively
small number objects have a `prototype` property**. It is these objects
with `prototype` properties that define the `prototypes` for the other objects.

Most built-in constructors (and most user-defined constructors) have
a prototype that inherits from `Object.prototype`. For example, `Date.prototype`
inherits properties from `Object.property`, so a Date object created by
`new Date()` inherits properties from both Date.prototype and Object.prototype.
The linked series of prototype objects is known as a **prototype chain**.

#### 6.2.4 Object.create()

It create a new object, using its first argument **as the prototype of that object**.
It also takes an optional second argument that describes the properties of the new
object.

```js
let o1 = Object.create({ x: 1, y: 2 });
o1.x + o1.y;
```

**Tip**: passing a null as argument create an object that inherit nothing, not even
basic methods like toString()

`Object.create(Object.prototype)` has same effect as {} or new Object{}.

One use for `Object.create()` is when you want to guard against
unintended modification of an object by a library function that you
don't have control over. If it sets properties, however, those writes
will not affect the original object.

```js
let o = { x: "dont' change this value" };
library.function(Object.create(o)); // Guard against accidental modification
```

### 6.3 Querying and Setting Properties

#### 6.3.1 Objects As Associative Arrays

You can write the following code in JavaScript:

```js
let addr = "";
for (let i = 0; i < 4; i++) {
  addr += customer[`address${i}`] + "\n";
}
```

In ES6 and later, the Map class is often a better choice than
using a plain object.

#### 6.3.2 Inheritance

Suppose you query the property `x` in the object `o`. If `o` does not have
an own property with that name, the prototype object of `o` is queried
for the property `x`. This continues until the property `x` is found
or until object with a `null` prototype is searched.

```js
let o = { x: 1 };
let p = Object.create(o);
let q = Object.create(p);
let f = q.toString(); // toString is inherited from Object.prototype
q.x; // 1
```

Property assignment examines the prototype chain only to determine
whether the assignment is allowed. If `o` inherits a read-only property
named x, for example, **then the assignment is not allowed**.
If the assignment is allowed, however, **it always creates or
sets a property in the original object and never modifies
objects in the prototype chain**.

There is one exception: if `o` inherits the property `x`,
and that property is an accessor property with a setter method,
then the setter method is called rather than creating a new property `x` in `o`.
However, that the setter method is called on the object `o`, which
again leave the prototype chain unmodified.

#### 6.3.3 Property Access Errors

Attempting to set a property on `null` or `undefined` causes a TypeError.
Some properties are read-only and cannot be set, and some objects do not allow
the addition of new properties. In strict mode, a TypeError is thrown whenever
an attempt to set a property fails, outside of strict mode these failures are usually
silent.

An attempt to set a property `p` of an object `o` fails in these circumstances:

- `p` is read-only in `o`.

- `o` has an inherited property `p` that is read-only

- `o` does not have an own property `p`, and `o`'s extensible attribute is `false`

### 6.4 Deleting Properties

The `delete` operator only deletes own properties, not inherited ones.
To delete an inherited property, you must delete it from the prototype
object in which it is defined. Doing this affects every object that
inherits from the prototype.

```js
let baba = {a:1}
const child = Object.create(baba)
console.log(child.a) // => 1
delete(baba.a)
console.log(child.a) // => undefined
```

A `delete` expression evaluates to `true` if the delete succeeded or if the
`delete` had no effect (such as deleting a nonexistent property)

```js
let o = { x: 1 };
delete o.x; // => true: deletes property x
delete o.x; // => true: does nothing
delete o.toString; // => true: does nothing (toString isn't an own property)
delete 1; // => true: nonsense
```

In strict mode, delete a non configurable property causes a TypeError, as are
properties of the global object created by variable declaration.
In non-strict mode, `delete` simply evaluates to `false` in this case:

```js
delete Object.prototype; // false: non configurable
var x = 2;
delete globalThis.x; // false
function f() {} // a global function
delete globalThis.f; // false
```

In non strict mode, when deleting configurable properties of global object,
you can omit the reference and simply delete property:

```js
globalThis.y = 1; // create a configurable global
delete y; // => true
```

However in strict mode, `delete` raises a SyntaxError if its operand is an
unqualified identifier like x:

```js
delete x; // SyntaxError
delete globalThis.x; // ok
```

### 6.5 Testing Properties

Test for whether an object has a property with a given name.
You can do this with `in` operator
, `hasOwnProperty()` and `propertyIsEnumerable()` method, or
simply by querying the property.

**Tip**: They works fine with Symbols property too.

> `in`: It returns `true` if the object has an own property or an inherited property
> by that name

```js
let o = { x: 1 };
"x" in o; // true
"y" in o; // false
"toString" in o; // true: inherits
```

> `hasOwnProperty()`: Returns `false` for inherited properties:

```js
let o = { x: 1 };
o.hasOwnProperty("x"); // true
o.hasOwnProperty("y"); // false
o.hasOwnProperty("toString"); // false
```

> `propertyIsEnumerable()`: Returns `true` only if the named
> property is an own property and its **enumerable attribute is true**

```js
let o = { x: 1 };
let p = Object.create(o);
o.propertyIsEnumerable("x"); // true
o.propertyIsEnumerable("y"); // false
p.propertyIsEnumerable("x") // false
Object.prototype.propertyIsEnumerable("toString"); // false: not enumerable
```

> Querying and use `!==` to make sure it is not undefined. However this cannot
> distinguish the `undefined` value

```js
let o = { x: 1 };
o.x !== undefined; // true
o.y !== undefined; // false
o.toString !== undefined; // true: o inherits
```

### 6.6 Enumerating Properties

The `for/in` loop runs the body of the loop once for each **enumerable property**.
The inherited properties are enumerable as well.

Below code guard against enumerating inherited or function properties:

```js
for (let p in o) {
  if (!o.hasOwnProperty(p)) continue;
}

for (let p in o) {
  if (typeof o[p] === "function") continue;
}
```

Another alternative is `for/of`: There are four functions you can use to get an array of property names:

- Object.keys(): does not include non-enumerable properties, inherited properties, or properties whose
  name is a Symbol.

- Object.getOwnPropertyNames(): works like Object.keys() including non-enumerable own properties as well, as long as their names are strings

- Object.getOwnPropertySymbols(): like Object.getOwnPropertyNames() but as long as
  their names are Symbols.

- Reflect.ownKeys() returns all own property names including enumerable and non-enumerable,
  and both string and Symbol.

<a id="6.6.1"></a>

#### 6.6.1 Property Enumeration Order

`for/of` of above mentioned Object methods (and related methods such as JSON.stringify) list properties in the following order:

- String properties whose names are non-negative integers are listed first from
  smallest to largest.

- String names are listed (including look like negative numbers or float-point number).
  In the order in which they were added to the object.

- Finally, the Symbols, in the order in which they were added to the object.

For `for/in`, it is not as tightly specified as it is for these enumeration functions,
but implementations typically enumerate **own properties** in the order just described,
then travel up the prototype chain enumerating properties in the same order for
each prototype object. However, same names occur only once, even non-enumerable property has already been considered.

#### 6.7 Extending Objects

`Object.assign()` copy **enumerable own** properties (including Symbols) of one object to another object.
It expects two or more objects as its arguments. It **modifies and
returns the first argument**, but **does not alter the second or any
subsequent arguments** (source objects).

Source object override properties by the same name in the target object.
And the properties in the second source object override properties with
the same name in the first source object.

If the source object has a **getter method or the target object has a
setter method, they will be invoked during the copy, but they will not themselves
be copied.**

```js
let o = {
  _x: 1,
  get x() {
    console.log("called get x() in o");
    return this._x;
  },
  set x(x) {
    console.log(`called set x() in p with ${x}`);
    this._x = x;
  },
  get abc() {
    return this._x++;
  }
};
let p = {
  _x: 2,
  get x() {
    console.log("called get x() in p");
    return this._x;
  },
  set x(x) {
    console.log(`called set x() in p with ${x}`);
    this._x = x;
  },
};
let q = {
  _y: 3,
  set x(x) {
    console.log(`called set x() in q with ${x}`);
    this._x = x;
  },
  get y() {
    console.log("called get y() in q");
    return this._x;
  },
};
let r = Object.assign(q, o, p);
r; // => {y:3, x:2}
q; // => {y:3, x:2}
q.abc // => 1
q.abc // => 1
q.abc // => 1
q.abc // => 1


Result:
called get x() in o
called set x() in q with 1
called get x() in p
called set x() in q with 2
```

Follow example shows assigning a object with default value:

```js
o = Object.assign({}, defaults, o);
// or another way using ...
o = { ...defaults, ...o };
```

We could also avoid the overhead of the extra object creation by writing
a version of Object.assign() that copies properties only if they are missing:

```js
function merge(target, ...sources) {
  for (let source of sources) {
    for (let key of Object.keys(source)) {
      if (!(key in target)) { // this is different with Object.assign
        target[key] = source[key];
      }
    }
  }
  return target;
}

let a = {a:0}
let b = {a:1}
console.log(merge(a,b)) // => {a:0}
console.log(Object.assign(a,b)) // => {a:1}
```

It is straightforward to write other property manipulation utilities like
this merge() function. A restrict() function could delete
properties of an object they do not appear in other template object.
Or a subtract() function could remove all of the properties of one object
from another object.

```js
function substract(target, ...sources) {
  for (let source of sources) {
    for (let key of Object.keys(source)) {
      if (key in target) {
        delete target[key];
      }
    }
  }
  return target;
}
```

<a id="6.8"></a>

### 6.8 Serializing Objects

JSON.stringify() and JSON.parse() serialize and restore JavaScript objects
(own enumerable properties).

- `NaN`, `Infinity` and `-Infinity` are serialized to `null`.

- `undefined`, RegExp, and Error objects cannot be serialized (omitted from the
  stringified output)

- Date object are serialized to ISO-formatted date string
  but JSON.parse() leaves these in string form and does not restore the original
  Date object.

JSON.stringify() and JSON.parse() accept optional second arguments that
can be used to customize the serialization and/or restoration process. [See here](#11.6)

### 6.9 Object Methods

#### 6.9.1 The `toString()` Method

The default toString() method is not very informative. For example: `[object Object]`.
Because of this, many classes define their own versions of toString().

```js
let point = {
  x: 1,
  y: 2,
  toString: function () {
    return `(${this.x}, ${this.y})`;
  },
};
String(point); // => "(1, 2)"
```

#### 6.9.2 The `toLocaleString()` Method

The default `toLocaleString()` method defined by Object doesn't do any
localization itself: it simply calls `toString()` and returns that value.
The Date and Number classes defined customized versions of toLocaleString()

```js
a = 12312232131231231.123123;
a.toLocaleString(); // => '12,312,232,131,231,232'

let b = new Date();
b.toLocaleString(); // => '12/16/2023, 9:52:00 AM'
b.toString(); // => 'Sat Dec 16 2023 09:52:00 GMT+0800 (Malaysia Time)'
```

#### 6.9.3 The `valueOf()` Method

#### 6.9.4 The `toJSON()` Method

Object.prototype does not actually define a toJSON() method,
but the JSON.stringify() method looks for a toJSON()
method on any object it is asked to serialize. If this
method exists on the object to be serialized, it is invoked, and
the return value is serialized instead of the original object.

```js
let point = {
  x: 1,
  y: 2,
  toString: function () {
    return `(${this.x}, ${this.y})`;
  },
  toJSON: function () {
    return this.toString();
  },
};
JSON.stringify([point]); // => '["(1,2)"]' instead of '[{"x":1,"y":2}]'
```

### 6.10 Extended Object Literal Syntax

#### 6.10.1 Shorthand Properties

```js
let x = 1,
  y = 2;
let o = {
  x: x,
  y: y,
};
```

In ES6 and later, you can drop the colon and one copy of the identifier
and end up with much simpler code:

```js
let x = 1,
  y = 2;
let o = { x, y };
o.x + o.y; // => 3
```

#### 6.10.2 Computed Property Names

With an `ES6` feature known as _computed properties_ that
lets you take the square brackets from the preceding code
and move them directly into the object literal:

```js
const PROPERTY_NAME = "p1";
function computePropertyName() {
  return "p" + 2;
}

let p = {
  [PROPERTY_NAME]: 1,
  [computePropertyName()]: 2,
};

p.p1 + p.p2; //=> 3
```

#### 6.10.3 Symbols as Property Names

In `ES6` and later, property names can be strings or symbols. You can
use that symbols as a property name using the computed property syntax.

```js
const extension = Symbol("my extension symbol");
let o = {
  [extension]: {},
};
o[extension].x = 0;
```

#### 6.10.4 Spread Operator

In `ES2018` and later, you can copy the properties of an existing object
into a new object using the "spread operator" ... inside an object
literal:

```js
let position = { x: 0, y: 0 };
let dimensions = { w: 100, h: 75 };
let ret = { ...position, ...dimensions };
```

**Important**: It spreads the own properties of an object, not any inherited ones.

```js
let father = {a: 3}
let child = Object.create(father)

let child1 = {...child}
console.log("a" in child1) // => false

let child2 = Object.create(father)
console.log("a" in child2) // => true
```

It is a special-case syntax available only within object literals. (Three
dots are used for other purposes in other JavaScript contexts, but
object literals are the only context where the three dots cause
this kind of interpolation of one object into another one.)


#### 6.10.5 Shorthand Methods

Prior to `ES6`, you would define a method in an object literal using
function definition expression just as you would define any other property of an
object:

```js
let square = {
  area: function () {
    return this.side * this.side;
  },
  side: 10,
};
square.area(); // => 100
```

In `ES6`, the object literal syntax has been extended to allow a shortcut
where the `function` keyword and the colon are omitted, this shorthand
syntax makes it clearer that area() is a method and not a data property:

```js
let square = {
  area() {
    return this.side * this.side;
  },
  side: 10,
};
square.area(); // => 100
```

You can also use string literals and computed property names, which can
include Symbol property names:

```js
const METHOD_NAME = "m";
const symbol = Symbol();
let weirdMethods = {
  "method With Spaces"(x) {
    return x + 1;
  },
  [METHOD_NAME](x) {
    return x + 2;
  },
  [symbol](x) {
    return x + 3;
  },
};
```

In order to make an object iterable for Symbols, you must
define a method with the symbolic name `Symbol.iterator`

#### 6.10.6 Property Getters and Setters

`ES5`

JavaScript also supports _accessor_ properties, which
do not have a single value but instead have one
or two accessor methods: a getter and/or a setter

When a program queries the value of an accessor property, JavaScript
invokes the getter method; When a program sets the value of an accessor
property, JavaScript invokes the setter method. If a property has only
a setter method, attempts to read it always evaluate to `undefined`

```js
let o = {
  dataProp: 0,

  get accessorProp() {
    return this.dataProp;
  },
  set accessorProp(value) {
    this.dataProp = value;
  },
};
```

In ES6, you can also use computed property names when defining getters
and setters. Simply replace the property name after `get` or `set` with
an expression in square brackets.

```js
let o = {
  dataProp: 0,

  set ["space space"](value) {
    this.dataProp = value;
  },
};
o["space space"] = 123;
o.dataProp; // => 123
```

Accessor properties are inherited, just as data properties are.

Other reasons to use accessor properties include sanity checking
of property writes and returning different values on each property read:

```js
const serialnum = {
  _n: 0,

  get next() {
    return this._n++;
  },
  set next(n) {
    if (n > this._n) this._n = n;
    else throw new Error("serial number can only be set to a larger value");
  },
};
```

**Important**: `_` in the property name hints that it is for internal use only.

## 7 - Arrays

### 7.1 Creating Arrays

#### 7.1.1 Array Literals

```js
let empty = [];
let primes = [2, 3, 5, 7, 11];
let misc = [1.1, true, "a"];
```

If an array literals contains multiple commas in a row, with no value
between, the array is sparse.

```js
let count = [1, ,];
count[1]; //=> undefined
count.length; // => 2
```

#### 7.1.2 The Spread Operator

```js
let a = [1, 2, 3];
let b = [0, ...a, 4]; // b == [0,1,2,3,4]
```

The spread operator is a convenient way to create a (shallow) copy of
an array

```js
let o = {};
let original = [1, 2, 3, o];
let copy = [...original];
copy[0] = 0;
copy[3].x = 4;
original[0]; // => 1
original[3]; // => {x:4}
```

The spread operator works on any iterable object (Iterable objects are
what the for/of loop iterates over). Strings are iterable:

```js
let digits = [..."0123456789ABCDEF"];
digits; // => ["0","1","2","3","4","5","6","7","8","9","A","B","C","D","E", "F"]
```

Set objects are iterable, below example remove duplicate elements from an array,
and convert back to an array:

```js
let letters = [..."hello world"];
[...new Set(letters)]; // => ["h","e","l","o"," ","w","r","d"]
```

#### 7.1.3 The Array() Constructor

**Using an array literal is almost always simpler than this usage of
Array() constructor.**

> With no arguments

```js
let a = new Array();
```

> Call it with a single numeric argument, which specifies a length:

```js
let a = new Array(10);
a.length; // => 10
```

> Explicitly specify two or more array elements or a single non-numeric
> element for the array

```js
let a = new Array(5, 4, 3, 2, 1);
```

#### 7.1.4 Array.of()

`ES6`

```js
Array.of(); // => []
Array.of(10); // => [10]
Array.of(1, 2, 3); // => [1,2,3]
```

In `ES6`, the Array.of() function addresses the problem of Array constructor's
creates a array with length specified with single numeric value.

#### 7.1.5 Array.from()

`ES6`

It expects an iterable or array-like object as it first argument and returns
a new array that contains the elements of that object. Array.from(iterable)
works like [...iterator] does. It is also a simple way to make a copy of
an array:

```js
let copy = Array.from(original);
```

Array.from() is also important because it defines a way to make a
**true-array copy of an array-like object**.

```js
let truearray = Array.from(arraylike);
```

It also accepts an optional second argument. If you pass a function
as the second argument, then as the new array is being
built, each element from the source object will be passed to the
function you specify (like the array map() method, but it is more efficient
to perform the mapping while the array is being built than it is to build
the array and then map it to another new array).

```js
let from = [1, 2, 3, 4, 5];
let to = Array.from(from, (v) => v * 2);
to; // => [2,4,6,8,10]
```

### 7.2 Reading and Writing Array Elements

Remember that arrays are a specialized kind of object. JavaScript
converts numeric array index you specify to a string -- the index 1
becomes the string "1".

```js
let o = [1];
o["0"] = 123;
o[0]; // => 123
```

When you try to query a nonexistent property of any
object, you don’t get an error; you simply get undefined

### 7.3 Sparse Array

Note that `undefined` element and sparse is not the same case:

```js
let a1 = [,];
let a2 = [undefined];
0 in a1; // => false
0 in a2; // => true
a1[0]; // undefined
a2[0]; // undefined
```

### 7.4 Array Length

Array have two special behaviors

- If you assign a value to an array element whose index
  i is greater than or equal to array's current length,
  the value of the length property is set to i+1

  ```js
  let a = [];
  a[10] = 1;
  a.length; // => 11
  ```

- If you set the length property to a non negative integer n smaller
  than its current value, any array elements whose index is greater than
  or equal to n **are deleted from the array**:

  ```js
  let a = [1, 2, 3, 4, 5];
  a.length = 0;
  a.length = 5;
  a[0]; // => undefined
  ```

### 7.5 Adding and Deleting Array Elements

> Assignments

```js
let a = [];
a[0] = "zero";
a[1] = "one";
```

> push()

```js
let a = [];
a.push(0);
a.push(1);
a; // => [0,1]
```

> unshift()

```js
let a = [];
a.unshift(0);
a.unshift(1);
a; // => [1,0]
```

> shift()

```js
let a = [1, 2, 3];
a.shift();
a; // => [2,3]
```

> pop()

```js
let a = [1, 2, 3];
a.pop();
a; // => [1,2]
```

> delele operator

```js
let a = [1, 2, 3];
delete a[2];
a.length; // => 3
a[2]; // => undefined
```

### 7.6 Iterating Arrays

`ES6`

> for/of loop

```js
for (let v of arr) {
  ...
}
```

If there is a need to know the index, use the entries() method of the
array along with destructuring assignment:

```js
for (let [index, v] of arr.entries()) {
  ...
}
```

> forEach() method

```js
arr.forEach(v => {
  ...
})
```

The second argument is the index:

```js
arr.forEach((v,i) => {
  ...
})
```

**Tip**: Unlike the for/of loop, the forEach() is aware
of sparse arrays and does not invoke your function for elements that are not here.

```js
let letters = [..."asdasdasd"]
delete letters[3]
letters.forEach((v,i) => {console.log(i + v)})
0a
1s
2d
4s
5d
6a
7s
8d
```

> Good old fashioned for loop

If you want to skip undefined and nonexistent elements, you might write:

```js
for (let i = 0; i < a.length; i++) {
  if (a[i] === undefined) continue;
}
```

### 7.7 Multidimensional Arrays

### 7.8 Array Methods

#### 7.8.1 Array Iterator Methods

All of these methods accept a function as their first qrgument
and invoke that function once for each element (or some elements) of
the array. **If the array is sparse, the function you pass
is not invoked for nonexistent elements**. In most cases,
the function you supply is invoked with **three arguments:
the value of the array element, the index of the array element,
and the array itself**.

Most of these iterator methods described in the following subsections
accept an optional second argument. If specified, the function
**is invoked as if it is a method of this second argument**. That is,
the second argument you pass becomes the value of the `this` keyword
inside of the function you pass as the first argument.

```js
const numbers = [1, 2, 3, 4, 5];
// notice here, the arrow () => will not work!
numbers.forEach(
  function () {
    console.log(this.message);
  },
  { message: "Hello, world!" },
);
```

##### FOREACH()

```js
let data = [1, 2, 3, 4, 5],
  sum = 0;
data.forEach((v) => {
  sum += v;
});
sum; // => sum == 15

data.forEach((v, i, a) => {
  a[i] = v + 1;
});
data; // => [2,3,4,5,6]
```

**Important**: There is no way to terminate iteration before all elements have been
passed to the function. No equivalent of `break` statement you can use.

##### MAP()

Invoke the function you specify and returns an array containing the
values returned by your function.

How it react with sparse array was shown below, the returned array
will be sparse in the same way as the original array:

```js
let a = [1, , 3];
a.map((x) => x * x); // => [ 1, <1 empty item>, 9 ]
```

##### FILTER()

It returns an array containing a subset of the elements of the array
on which it is invoked.

```js
let a = [5, 4, 3, 2, 1];
a.filter((x) => x < 3); // => [2,1]
a.filter((x, i) => i % 2 === 0); // => [5,3,1]
```

Note that `filter()` skips missing elements in sparse arrays, to
close the gaps in a sparse array, you can do this:

```js
let dense = sparse.filter(() => true);
```

To close gaps and remove undefined and null elements:

```js
let dense = sparse.filter((x) => x !== undefined && x !== null);
```

##### FIND() And FINDINDEX()

They stopped iterating the first time the predicate finds an element.
When that happens:

- `find()` returns the matching element, returns undefined if not found

- `findIndex()` returns the index of the matching element, returns -1 if not found

```js
let a = [1, 2, 3, 4, 5];
a.findIndex((x) => x === 3); // => 2
a.find((x) => x === 3); // => 3
```

##### EVERY() And SOME()

They apply a predicate function you specify to the elements of the array,
then return `true` or `false`

The every() returns true if and only if your predicate function returns `true`
for all elements in the array:

```js
let a = [1, 2, 3, 4, 5];
a.every((x) => x < 10); // => true
```

The some() method like returns true if there exist at least one element
in the array for which the predicate returns true.

```js
let count = 0;
let a = [1, 2, 3, 4, 5];
a.some((x) => {
  count++;
  return x === 3;
}); // => true
count; // => 3

a.some(isNaN); // => false
```

##### REDUCE() And REDUCERIGHT()

They combine the elements of an array, using the function you specify,
to produce a single value.

```js
let a = [1, 2, 3, 4, 5];
a.reduce((x, y) => x + y, 0); // => 15
a.reduce((x, y) => x * y, 1); // => 120
a.reduce((x, y) => (x > y ? x : y)); // => 5
```

`reduce()` takes two arguments. The first is the function that performs
the reduction operation. The second (optional) argument is an initial value
to pass to the function.

- If no second argument was given, it uses the first element of the
  array as the initial value. (In the sum and product examples, we could
  have omitted the initial value argument)

Calling reduce() on an empty array with no initial value argument causes a
TypedError. If you call it with only one value
(either one element array without initial value or empty array with initial value)
, it just returns that one value without ever calling the reduction function.

`reduceRight()` works just like `reduce()`, except that it processes from
highest index to lowest.

```js
let a = [2, 3];
a.reduceRight((acc, val) => Math.pow(val, acc)); // => 8
```

**Neither `reduce()` nor `reduceRight()` accepts an optional argument
that specifies the `this` value on which the reduction function is
to be invoked**. See the [Function.bind() method](#functionbind) if you need that.

#### 7.8.2 Flattening arrays with `flat()` and `flatMap()`

`ES2019`

The `flat()` method creates and returns a new array
that contains the same elements as the array it is called on,
except that any elements that are themselves array are "flattened" into
the returned array. For example:

```
[1, [1, 2]].flat() // => [1,1,2]
[1, [1, [2]]].flat(); // => [1,1,[2]]
```

If you want to flatten more levels, pass a number to flat()

```js
[1, [2, [3, [4]]]].flat(1); // => [1,2,[3,[4]]]
[1, [2, [3, [4]]]].flat(2); // => [1,2,3,[4]]
```

The `flatMap()` works just like `map()` except the returned array
is automatically flattened as if passed to `flat()`, it is same
as (but more efficient due to no extra array created) a.map(f).flat()

```js
let phrases = ["hello world", "the definitive guide"];
let words = phrases.flatMap((phrase) => phrase.split(" "));
words; // => ["hello", "world", "the", "definitive", "guide"];
```

<a id="7.8.3"></a>

#### 7.8.3 Adding arrays with `concat()`

The `concat()` method **creates and returns new array** that contains
the elements of the original array on which `concat()` was invoked,
followed by each of the arguments to `concat()`.

```js
let a = [1, 2, 3];
a.concat(4, 5); // => [1,2,3,4,5]
a.concat([4, 5], [6, 7]); // => [1,2,3,4,5,6,7]
a.concat([4, [5]]); // => [1,2,3,4,[5]]
```

**If the argument is not an array, it simply is appended to the new array.**

**Note that concat() makes a new copy of the array, if you find yourself
writing `a = a.concat(x)`, you should use `push()` or `splice()`**

#### 7.8.4 Stacks and Queues with `push()`, `pop()`, `shift()`, and `unshift()`

> Stack

`push()` + `pop()`. You could use `unshift()` and `shift()` to implement a stack,
**but it would be less efficient than push() + pop() because the array
elements need to be shifted up or down every time the element is added or removed
at the start of the array.**

> Queue

`push()` + `shift()`

#### 7.8.5 Subarrays with `slice()`, `splice()`, `fill()`, and `copyWithin()`

##### SLICE()

Its two arguments specify the start and end (not including) of the slice to be returned.
If either argument is negative, ie specifies an array element relative to the length of
the array (e.g -1 is the last element of the array)

```js
let a = [1, 2, 3, 4, 5];
a.slice(0, 3); // => Returns [1,2,3]
a.slice(-2, -1); // => Returns [4]
a.slice(-1, -2); // => Returns []
```

##### SPLICE()

Is a general-purpose method for inserting or removing elements from an array.

Unlike `slice()` and `concat()`, it modifies the array on which it is invoked.

The first argument to `splice()` specifies the array position at which
the insertion and/or deletion is to begin.

The second argument specifies the number of elements that should be deleted from
(spliced out of) the array. If this second argument is omitted, all
array elements from the start element to the end of the array are removed.
`splice()` returns an array of the deleted elements, or an empty array if
no elements were deleted

```js
let a = [1, 2, 3, 4, 5, 6, 7, 8];
a.splice(4); // => [5,6,7,8], a is now [1,2,3,4]
a.splice(1, 2); // => [2,3], a is now [1,4]
a.splice(5); // => [], a is now [1,4]
```

The first two arguments to splice() specify which array elements
are to be deleted. Additional arguments can be followed which
are the elements to be inserted into the array:

```js
let a = [1, 2, 3, 4, 5];
a.splice(2, 0, "a", "b"); // => []; a is now [1,2,"a","b",3,4,5]
a.splice(2, 2, [1, 2], 3); // => ["a","b"]; a is now [1,2,[1,2],3,3,4,5]
```

##### FILL()

It mutates the array it is called on, also returns the modified array

```js
let a = new Array(5);
a.fill(0); // => [0,0,0,0,0]
a.fill(9, 1); // => [0,9,9,9,9]
a.fill(8, 2, -1); // => [0,9,8,8,9]
```

##### COPYWITHIN()

`copyWithin()` copies a slice of an array to a new position within the array.
It modifies the array in place and returns the modify array,
**but it will not change the length of the array**.

- first argument: the destination index to which the first element will be copied
- second argument: the index of the first element to be copied.
  - if it is omitted, 0 is used
- third argument: the end of the slice of elements to be copied.
  - if omitted, length of array is used

```js
let a = [1, 2, 3, 4, 5];
a.copyWithin(1); // => [1,1,2,3,4]
a.copyWithin(2, 3, 5); // => [1,1,3,4,4]
a.copyWithin(0, -2); // => [4,4,3,4,4]
```

#### 7.8.6 Array Searching and Sorting Methods

##### INDEXOF() And LASTINDEXOF()

**Tip**: They use `===`, thus, it cannot detect NaN

The second argument to indexOf() can be used to find
matches beyond the first.

```js
let a = [1, 2, 1, 2, 1];
a.indexOf(2, 2); // => 3
a.lastIndexOf(2, 2); // => 1
```

##### INCLUDES()

`ES2016`

If the array contains the value.

The `includes()` method of an array also performs membership testing, **but the time
it takes is proportional to the size of the array, and using an array as
a set can be much, much slower than using a real Set object**

```js
let a = [1, true, 3, NaN];
a.includes(true); // => true
a.includes(2); // => false
a.includes(NaN); // => true
a.indexOf(NaN); // => -1
```

##### SORT()

It mutates the array it is called on, also returns the modified array.

**It sorts the array elements in alphabetical order (temporarily converting
them to strings to perform the comparison if necessary)**

If an array contains undefined elements, they are sorted to the end of
the array. To sort an array into some order than alphabetical,
you must pass a comparison function as an argument to sort().
**If the first argument should appear before the second, the
comparison function should return a number lesser than zero.**

To sort by numerical rather than alphabetical order:

```js
let a = [33, 4, 1111, 222];
a.sort(); // => [1111,222,33,4]
a.sort(function (a, b) {
  return a - b;
}); // [4,33,222,1111]
```

Case insensitive alphabetical sort:

```js
let a = ["ant", "Bug", "cat", "Dog"];
a.sort(); // a == ["Bug","Dog","ant","cat"]; casesensitive
sort;
a.sort(function (s, t) {
  let a = s.toLowerCase();
  let b = t.toLowerCase();
  if (a < b) return -1;
  if (a > b) return 1;
  return 0;
}); // a == ["ant","Bug","cat","Dog"]; case-insensitive sort
```

##### REVERSE()

It mutates the array it is called on, also returns the modified array.

```js
let a = [1, 2, 3];
a.reverse(); // => [3,2,1]]
```

#### 7.8.7 Array to String Conversions

Three methods that can convert arrays to strings
which is generally something you might do when creating log
and error messages.

##### Join()

```js
let a = [1, 2, 3];
a.join(); // => "1,2,3"
a.join(" "); // => "1 2 3"
a.join(""); // => "123"
```

Use it to repeat character:

```js
let b = new Array(10);
b.join("-"); // => "----------"
```

##### Split()

##### toString()

Works like join() method with no argument

```
[1,2,3].toString() // => "1,2,3"

[1,2,[3]].toString() // => "1,2,3"
[1,2,[3,[4]]].toString() // => "1,2,3,4"
```

```
[1,2,3].toLocaleString() // => "1,2,3"

[1,2,[3]].toLocaleString() // => "1,2,3"
[1,2,[3,[4]]].toLocaleString() // => "1,2,3,4"
```

#### 7.8.8 Static Array Functions

```js
Array.of();
Array.from();
Array.isArray([]); // => true
Array.isArray({}); // => false
```

### 7.9 Array-Like Objects

In client-side JavaScript, a number of methods for working
with HTML documents (such as `document.querySelectorAll()`),
return array-like objects. Here's function you might use to
test for objects that work like arrays.

```js
function isArrayLike(o) {
  if (o && type o === "object" && Number.isFinite(o.length) &&
  o.length >= 0 && Number.IsInteger(o.length) && o.length < 4294967295) {
  return true
  } else {
  return false
  }
}
```

Since array-like objects do not inherit from `Array.prototype`,
you cannot invoke array methods on them directly. You can invoke
them indirectly using Function.call method, however:

```js
let a = { 0: "a", 1: "b", 2: "c", length: 3 };
Array.prototype.join.call(a, "+"); // => "a+b+c"
Array.prototype.map.call(a, (x) => x.toUpperCase()); // => ["A","B","C"]
Array.from(a); // => ["a", "b", "c"]
```

### 7.10 Strings As Arrays

```js
Array.prototype.join.call("JavaScript", " "); // => 'J a v a S c r i p t'
```

Keep in mind that strings are immutable values, so when they are
treated as arrays, they are read-only arrays. Arrays methods like
`push()`, `sort()`, `reverse()`, and `splice()` modify an array in
place and do not work on strings. It fails silently.

<a id="8"></a>

## 8 - Functions

**Tip**: `f.name` is the name of the function

<a id="8.1"></a>

### 8.1 Defining Functions

`ES6` defines another important new way to define functions
without `function` keyword: "arrow functions".

Note that functions can also be defined with the [`Function()` constructor](#8.7.7). Also, JavaScript defines some
specialized kinds of functions. `function*` defines [generator functions](#12) and `async function` defines
[asynchronous functions](#13)

#### 8.1.1 Function Declarations

Function declarations consist of the `function` keyword:

```js
function identifier(parameter) {
  ...
}
```

Function declaration statements are "hoisted" to the top
of the enclosing script, function, or block so that
functions defined in this way may be invoked from code that
appears before the definition.

Prior to `ES6`, function declarations were only allowed at the
top level within a JavaScript file or within another function.
In the strict mode of `ES6`, however, function declarations are
allowed within blocks. A function defined within a block only exists
within that block.

#### 8.1.2 Function Expressions

It looks a lot like function declaration, but they appear within the
context of a larger expression or statement, and the
name is **optional** (name is useful for recursion):

```js
const square = function (x) {
  return x * x;
};
const f = function fact(x) {
  if (x <= 1) return 1;
  else return x * fact(x - 1);
};

// It can also be used as arguments
[3, 2, 1].sort(function (a, b) {
  return a - b;
});

// Are sometimes defined and immediately invoked
let tensquared = (function (x) {
  return x * x;
})(10);
```

It is a good practice to use `const` with function
expressions so you don't accidentally overwrite your
functions by assigning new values.

#### 8.1.3 Arrow Functions

`ES6`

```js
const sum = (x, y) => {
  return x + y;
};

const sum = (x, y) => x + y; // if the body of the function is a single return statement

const polynomial = x => x * x + 2 * x + 3; // exactly one parameter

const constantFunc = () => 42; // no arguments at all
```

If the expression to be returned is an object literal:

```js
const f = (x) => { return { value: x }; };
const g = x =>  ({value:x});

const h = x => {value:x}; // Bad: h() returns nothing
```

The concise syntax of arrow functions makes them ideal
when you need to pass one function to another function:

```js
[1, null, 2, 3].filter((x) => x !== null);
```

Arrow function differ from functions defined in other ways in one critical
way: they inherit the value of `this` keyword from the
environment in which they are defined rather than defining their
own invocation context as functions defined in other ways do.

```js
function hello() {
  this.a = "1"
  // From the environment
  const fn = () => {console.log(this.a)}
  fn()
}
hello() // => 1

// Not defining their own invocation context
let a = {
  a: 1,
  fn: function () {
    console.log(this.a);
  },
  fn2: () => {
    console.log(this.a);
  },
};
a.fn(); // 1
a.fn2(); // undefined
```

Arrow functions also differ from other functions in that they do not
have a `property` property, **which means that cannot be used as
constructor functions for new classes**

#### 8.1.4 Nested Functions

### 8.2 Invoking Functions

JavaScript functions can be invoked in five ways:

#### 8.2.1 Function Invocation

For function invocation in non-strict mode, the invocation context (the
`this` value) is the global object. In strict mode, however, the
invocation context is `undefined`. Note that functions defined using
the arrow syntax behave differently: they always inherits the `this`
value that is in effect where they are defined.

**Tip**: to determine whether we are in strict mode

```js
const strict = (function () {
  return !this;
})();
```

#### 8.2.2 Method Invocation

Method invocations differ from function invocations in one important way:
the invocation context.

> **Method Chaining**

When methods return objects, you can use the return value of one
method invocation as part of a subsequent invocation:

```js
doStepOne().then(doStepTwo).then(doStepThree);
new Square().x(100).y(100).size(50).fill("blue").draw();
```

Except for arrow functions, nested functions do not inherit the `this` value
of the containing function.

```js
let o = {
  m: function () {
    let self = this;
    this === o; // => true
    f();

    function f() {
      this === o; // false
      self === o; // true
    }
  },
};
```

Another workaround instead of assigning to another variable is by using
arrow function, **arrow functions inherit this**:

```js
const f = () => {
  this === o; // true
};
```

Another workaround is to invoke the `bind()` method of nested
function to define a new function that is implicitly invoked on a specified
object:

```js
const f = function () {
  this === o;
}.bind(this);
```

#### 8.2.3 Constructor Invocation

If a function or method invocation is preceded by the
keyword `new`, then it is a constructor invocation.

This is not common practice, but you can omit a pair of empty
parentheses in a constructor invocation:

```
o = new Object();
// is same as
o = new Object;
```

If a constructor explicitly uses the `return` statement to return an
object, then that object becomes the value of the invocation expression.
If the constructor use `return` with no value, or if it returns a
primitive value, that return value is **ignored** and the new object is
used as the value of the invocation.

#### 8.2.4 Indirect Invocation

`call()` and `apply()` invoke the function indirectly. They
allow you to explicitly specify the `this` value for the invocation,
which means you can invoke any function as a method of any object,
**even if it is not actually a method of that object**.

See [8.7.4 for more](#8.7.4)

#### 8.2.5 Implicit Function Invocation

The language features that can cause implicit function invocation include:

- If an object has getters or setters defined, then querying or setting the value
  of its properties may invoke those methods.

- When an object is used in a string context (eg concatenated), its `toString()` method
  is called; `valueOf()` is called when an object is used in a numeric context.

- When you loop over the elements of an iterable object: a number of method calls
  that occur (e.g next())

- A tagged template literal. [Here for more](#14.5)

- Proxy objects have their behavior completely controlled by functions. Just
  about any operation on one of these objects will cause a function to be invoked.
  [Describe here](#14.7)

### 8.3 Function Arguments and Parameters

#### 8.3.1 Optional Parameters and Defaults

When a function is invoked with fewer arguments than declared parameters,
the additional parameters are set to their default value, which is
normally `undefined`.

```js
function getPropertyName(o, a) {
  a = a || [];
  for (let property in o) a.push(property);
  return a;
}
```

In `ES6`, you can define a default value for each of your function
parameters directly in the parameter list of your function.

```js
function getPropertyName(o, a = []) {
  for (let property in o) a.push(property);
  return a;
}
```

You can use a variable too for your default value.

You can use the value of a previous parameter:

```
const rectangle = (width, height = width*2) => ({width, height});
rectangle(1) // => {width:1, height:2}
```

#### 8.3.2 Rest Parameters and Variable-Length Argument Lists

`Rest` parameters allows us to write functions that can be invoked
with arbitrarily more arguments than parameters. The value of
a rest parameter will always be an array. **The array may be
empty, but a rest parameter will never be `undefined`**. A
rest parameter being defined a parameter default.

```js
function max(first = -Infinity, ...rest) {
  let maxValue = first;
  for (let n of rest) {
    if (n > maxValue) {
      maxValue = n;
    }
  }
  return maxValue;
}

max(1, 10, 100, 2, 3, 1000, 4, 5, 6); // => 1000
```

Above function called _variadic function_, _variable arity function_
or _vararg_ functions.

#### 8.3.3 The Arguments Object

Rest parameters were introduced into JavaScript in ES6.
Before that version, varargs functions were written using
the Arguments object.

```js
function max(x) {
  let maxValue = -Infinity;
  for (let i = 0; i < arguments.length; i++) {
    if (arguments[i] > maxValue) maxValue = arguments[i];
  }
  return maxValue;
}
max(1, 10, 100, 2, 3, 1000, 4, 5, 6); // => 1000
```

**You should avoid using it in any new code you write**.

**When refactoring old code, use rest parameters**.

#### 8.3.4 The Spread Operator for Function Calls

```js
Math.min(...numbers);
```

Rest parameters and the spread operator are
often useful together, following function returns
an instrumented version of the function for testing:

<a id="timedfunction"></a>

```js
function timed(f) {
  return function (...args) {
    console.log(`Entering function ${f.name}`);
    let startTime = Date.now();
    try {
      return f(...args);
    } finally {
      console.log(`Exiting ${f.name} after ${Date.now() - startTime}ms`);
    }
  };
}
function benchmark(n) {
  let sum = 0;
  for (let i = 1; i <= n; i++) sum += i;
  return sum;
}
timed(benchmark)(1000000); // => 500000500000
```

#### 8.3.5 Destructuring Function Arguments into Parameters

```js
function vectorAdd([x1, y1], [x2, y2]) {
  return [x1 + x2, y1 + y2];
}
vectorAdd([1, 2], [3, 4]); // => [4,6]
```

This example of destructing a single object argument into two
parameters:

```js
function vectorMultiply({ x, y }, scalar) {
  return { x: x * scalar, y: y * scaler };
}
vectorMultiply({ x: 1, y: 2 }, 2); // => {x:2, y:4}
```

The syntax is more verbose and more confusing when you need
to destructure properties with one name into parameters with different names

```js
function vectorAdd({ x: x1, y: y1 }, { x: x2, y: y2 }) {
  return { x: x1 + x2, y: y1 + y2 };
}
vectorAdd({ x: 1, y: 2 }, { x: 3, y: 4 }); // => {x: 4, y: 6}
```

**The rule to keep in mind for destructuring assignment and destructuring
function calls is that the variables or parameters being declared
go in the spots where you'd expect values to go in an object literal**.
The property names are always on the **lefthand side** of the colon.

You can define parameter defaults with destructured parameters:

```js
function vectorMultiply({ x, y, z = 0 }, scalar) {
  return { x: x * scalar, y: y * scalar, z: z * scalar };
}
vectorMultiply({ x: 1, y: 2 }, 2); // => {x: 2, y: 4, z: 0}
```

Other example:

```js
// Define a rest parameter for extra values within the array that is being
// unpacked.
function f([x, y, ...coords], ...rest) {
  return [x + y, ...rest, ...coords];
}
f([1, 2, 3, 4], 5, 6); // => [3,5,6,3,4]

// Destruture an object
function vectorMultiply({ x, y, z = 0, ...props }, scalar) {
  return { x: x * scalar, y: y * scalar, z: z * scalar, ...props };
}
vectorMultiply({ x: 1, y: 2, w: -1 }, 2); // => {x: 2, y: 4, z: 0, w: -1}
```

#### 8.3.6 Argument Types

It is better for a function to fail immediately and predictably
when passed bad values than to begin executing and
fail later with an error message that is likely to be unclear.

```js
function sum(a) {
  let total = 0;
  for (let element of a) {
    // Throws TypeError if a is not iterable
    if (typeof element !== "number") {
      throw new TypeError("sum(): elements must be numbers");
    }
    total += element;
  }
  return total;
}
sum([1, 2, 3]); // => 6
sum(1, 2, 3); // !TypeError: 1 is not iterable
sum([1, 2, "3"]); // !TypeError: element 2 is not a number
```

### 8.4 Functions as Values

#### 8.4.1 Defining Your Own Function Properties

Functions are not primitive values in JavaScript, but a specialized
kind of object, which means that functions can have properties.

<a id="uniqueInteger1"></a>

```js
// Initialize the counter property of the function object.
// Function declarations are hoisted so we really can
// do this assignment before the function declaration.
uniqueInteger.counter = 0;

function uniqueInteger() {
  return uniqueInteger.counter++; // Return and increment counter property
}
uniqueInteger(); // => 0
uniqueInteger(); // => 1
```

Another example to cache previously computed results:

<a id="factorial"></a>

```js
function factorial(n) {
  if (Number.isInteger(n) && n > 0) {
    if (!(n in factorial)) {
      factorial[n] = n * factorial(n - 1);
    }
    return factorial[n];
  } else {
    return NaN;
  }
}
factorial[1] = 1; // init
factorial(6); // => 720
factorial[5]; // => 120
```

### 8.5 Functions as Namespaces

Variable declared within a function are not visible outside of the
function. For this reason, it is sometimes useful to define a function
simply to act as a temporary namespace in which you can define
variables without cluttering the global namespace.

```js
function chunkNamespace() {
  // chunk of code goes here
  // Any variables defined in the chunk are local to this function
  // instead of cluttering up the global namespace.
}
chunkNamespace(); // don't forget to invoke the function!
```

Or a immediately invoked function expression. **Parentheses is required!**

```js
(function () {
  // chunk of code goes here
})();
```

This use of functions as namespaces becomes really useful when
**we define one or more functions inside the namespace function
suing variables within that namespace, but then
pass them back out as the return value of the namespace function**.
Functions like this are known as `closures`.

### 8.6 Closures

```js
let scope = "global scope";
function checkScope() {
  let scope = "local scope";
  function f() {
    return scope;
  }
  return f();
}
checkScope(); // => "local scope"
```

Like most modern programming languages, JavaScript uses lexical
scoping. This means that functions are executed using the variable
scope that was in effect **when they were defined**, not the variable
scope that is in effect when they are invoked.

How about this?

```js
let scope = "global scope";
function checkScope() {
  let scope = "local scope";
  function f() {
    return scope;
  }
  return f;
}
checkScope()(); // => ?
```

It return `local scope`, remember that the nested function f()
was defined in a scope where the variable `scope` was bound to the
value "local scope".

We could [rewrite the `uniqueInteger()`](#uniqueInteger1), the shortcoming of the previous approach
is that buggy or malicious code could reset the counter or set it to a non-integer.

```js
let uniqueInteger = function () {
  let counter = 0;
  return function () {
    return counter++;
  };
};
uniqueInteger(); // => 0
uniqueInteger(); // => 1
```

No other code can see the `counter` variable: the inner function
has exclusive access to it.

It is perfectly possible for two or more nested functions to be defined
within the same outer function and share the same scope:

```js
function counter() {
  let n = 0;
  return {
    count: function () {
      return n++;
    },
    reset: function () {
      n = 0;
    },
  };
}
```

It is worth noting here that you can combine this closure technique
with property getters and setters.

```js
function counter(n) {
  return {
    get count() {
      return n++;
    },
    set count(m) {
      if (m > n) n = m;
      else throw Error("count can only be set to a larger value");
    },
  };
}
```

Think about the following code:

```js
function constfuncs() {
  let funcs = [];
  for (var i = 0; i < 10; i++) {
    funcs[i] = () => i;
  }
  return funcs;
}

let funcs = constfuncs();
funcs[5](); // => 10; Why doesn't this return 5?
```

The fundamental problem is that the variables declared with `var`
are defined throughout the function. If we just replace the `var`
with a `let` or `const`, then the problem goes away. Because `let`
and `const` are block scoped, each iteration of the loop defines a scope
that is independent of the scopes for all other iterations.

### 8.7 Function Properties, Methods, and Constructor

#### 8.7.1 The `length` Property

Read-only property. The number of parameters it declares in its parameter list

#### 8.7.2 The `name` Property

Read-only property. The name that was used when the function was defined.
This property is primarily useful when writing debugging or error messages.
Either, it was defined with name (e.g. function abc() {}, const what = function abc(){}, both cases names are abc),
or the name of the variable or property that an unnamed function expression
was assigned to when it was first created (e.g let unnamed = function(){}).

#### 8.7.3 The `prototype` Property

All functions, except arrow functions, have a `prototype` property
that refers to an object known as the prototype object. Every function
has a different prototype object.

<a id="8.7.4"></a>

#### 8.7.4 The `call()` and `apply()` Methods

Allows you to indirectly invoke a function as if it were a method of some
other object.

To invoke the function f() as a method of the object o:

```js
f.call(o);
f.apply(o);
```

Either of these lines of code are similar to the following:

```js
o.m = f;
o.m();
delete o.m;
```

**Remember that arrow functions inherit the `this` value of the context
where they are defined**. This cannot be overridden with the `call()` and
`apply()` methods. If you call either of those methods on an arrow function,
**the first argument is effectively ignore**.

```js
let a = {
  fn2: () => {
    this.a = "1"
    console.log(this.a);
  },
};
a.fn2.call({}) // 1
```

The subsequent arguments of call() are passed to the function that is invoked.

```js
f.call(o, 1, 2);
```

The apply() method is like the call() method, except that the arguments to be passed
to the function are specified as an array:

```js
f.apply(o, [1, 2]);
```

The following example is similar to the [timed()](#timedfunction), but it works
for methods instead of functions.

```js
function trace(o, m) {
  let original = o[m];
  o[m] = function (...args) {
    console.log(new Date(), "Entering:", m);
    let result = original.apply(this, args);
    console.log(new Date(), "Exiting:", m);
    return result;
  };
}
```

<a id="functionbind"></a>

#### 8.7.5 The `bind()` Method

The `bind()` bind a function to an object, and returns a new function.

```js
function f(y) {
  return this.x + y;
}
let o = { x: 1 };
let g = f.bind(o);
g(2); // => 3
let p = { x: 10, g };
p.g(3); // => 3: g is still bound to o, not p
```

Arrow functions inherits their `this` value from the environment in which they
are defined, and that value cannot be overridden with bind().

```js
let a = () => {
  console.log(this.a)
}

let b = a.bind({a: "hoho"})
b() // => undefined
```

The most common use case for calling `bind()` is to make non-arrow functions
behave like arrow functions, however, so this limitation on binding arrow
functions is not a problem in practice.

It can also do more than bind a function to an object:
Partial application is a common technique in functional programming and
is sometimes called currying. It works well with arrow functions.

```js
let sum = (x, y) => x + y;
let succ = sum.bind(null, 1);
succ(2); // => 3: x is bound to 1, and we pass 2 for the y argument
succ(2, 3); // => 3

function f(y, z) {
  return this.x + y + z;
}
let g = f.bind({ x: 1 }, 2);
g(3); // => 6
g.name // => `bound g`
```

**Tip**: The name of function returned by bind() is the name of.
the function that bind() was called on, prefixed with the word "bound". e.g. `bound f`

#### 8.7.6 The `toString()` Method

In practice, most implementations of this toString() method
return the complete source code for the function.
Built-in functions typically return a string that includes something
like "[native code]" as the function body.

<a id="8.7.7"></a>

#### 8.7.7 The Function() Constructor

```js
const f = new Function("x", "y", "return x*y;");
// equivalent to
const f = function (x, y) {
  return x * y;
};
```

- The Function() constructor allows JavaScript functions to be dynamically created
  and compiled at runtime.

- The Function() constructor parses the function body and create new function
  object each time it is called. If the call to the constructor appear
  within a loop or within a frequently called function, this process can be
  inefficient. By contrast, **nested functions and function expression that
  appear within loops are not recompiled each time they are encountered**.

- The functions it creates **do not use lexical scoping;** instead, they are always
  compiled as if they were top-level functions:

  - ```js
    let scope = "global";
    function constructFunction() {
      let scope = "local";
      return new Function("return scope");
    }
    constructFunction()() // => "global"
    ```

  - The Function() constructor is best thought of as a globally scoped
    version of eval() that defines new variables and functions in its own private scope.

You will probably never need to use this constructor in your code.

#### 8.8 Functional Programming

JavaScript is not a functional programming language like Lisp or Haskell,
but the fact that JavaScript can manipulate functions as objects
means that we can use functional programming languages in
JavaScript.

The follow sections are intended as a mind-expanding exploration of the
power of JavaScript's functions, not as a prescription for
good programming style.

##### 8.8.1 Processing Arrays with Functions

> Non functional style:

```js
let data = [1, 1, 3, 5, 5];
let total = 0;
for (let i = 0; i < data.length; i++) total += data[i];
let mean = total / data.length;
total = 0;
for (let i = 0; i < data.length; i++) {
  let deviation = data[i] - mean;
  total += deviation * deviation;
}
let stddev = Math.sqrt(total / (data.length - 1));
```

```js
const sum = (x, y) => x + y;
const square = (x) => x * x;
let data = [1, 1, 3, 5, 5];
let mean = data.reduce(sum) / data.length;
let deviations = data.map((x) => x - mean);
let stddev = Math.sqrt(deviations.map(square).reduce(sum) / (data.length - 1));
stddev; // => 2
```

> Functional style:
>
> Above is still invoking methods on objects, so it has some object-oriented conventions
> remaining.

```js
const map = function (a, ...args) {
  return a.map(...args);
};
const reduce = function (a, ...args) {
  return a.reduce(...args);
};

const sum = (x, y) => x + y;
const square = (x) => x * x;
let data = [1, 1, 3, 5, 5];
let mean = reduce(data, sum) / data.length;
let deviations = map(data, (x) => x - mean);
let stddev = Math.sqrt(
  reduce(map(deviations, square), sum) / (data.length - 1),
);
stddev; // => 2
```

#### 8.8.2 Higher-Order Functions

A _higher-order_ function is a function that operates on functions,
taking one or more functions as arguments and returning a new function.

```js
function not(f) {
  return function (...args) {
    let result = f.apply(this.args);
    return !result;
  };
}

const even = (x) => x % 2 == 0;
const odd = not(even);
[1, 1, 3, 5, 5].every(odd); // true
```

```js
function mapper(f) {
  return (a) => map(a, f);
}
const increment = (x) => x + 1;
const incrementAll = mapper(increment);
incrementAll([1, 2, 3]); // => [2,3,4]
```

```js
function compose(f, g) {
  return function (...args) {
    return f.call(this, g.apply(this, args));
  };
}
const sum = (x, y) => x + y;
const square = (x) => x * x;
compose(square, sum)(2, 3); // => 25; the square of the sum
```

#### 8.8.3 Partial Application of Functions

```js
function partialLeft(f, ...outerArgs) {
  return function (...innerArgs) {
    let args = [...outerArgs, ...innerArgs];
    return f.apply(this, args);
  };
}

function partialRight(f, ...outerArgs) {
  return function (...innerArgs) {
    let args = [...innerArgs, ...outerArgs];
    return f.apply(this, args);
  };
}

function partial(f, ...outerArgs) {
  return function (...innerArgs) {
    let args = [...outerArgs];
    let innerIndex = 0;
    for (let i = 0; i < args.length; i++) {
      if (args[i] === undefined) args[i] = innerArgs[innerIndex++];
    }
    args.push(...innerArgs.slice(innerIndex));
    return f.apply(this, args);
  };
}

const f = function (x, y, z) {
  return x * (y - z);
};
partialLeft(f, 2)(3, 4); // => -2
partialRight(f, 2)(3, 4); // => 6
partial(f, undefined, 2)(3, 4); // => -6
```

These partial application functions allows us to easily define
interesting functions:

```js
const increment = partialLeft(sum, 1);
const cuberoot = partialRight(Math.pow, 1 / 3);
cuberoot(increment(26)); // => 3
```

More interesting when we combine it with other higher-order functions

```js
const not = partialLeft(compose, (x) => !x);
const even = (x) => x % 2 == 0;
const odd = not(even);
const isNumber = not(isNaN);
odd(3) && isNumber(2); // => true
```

We can also use composition and partial application to redo our
mean and standard deviation calculations in extreme functional style:

```js
const product = (x, y) => x * y;
const neg = partial(product, -1);
const sqrt = partial(Math.pow, undefined, 0.5);
const reciprocal = partial(Math.pow, undefined, neg(1));
// Now compute the mean and standard deviation.
let data = [1, 1, 3, 5, 5]; // Our data
let mean = product(reduce(data, sum), reciprocal(data.length));
let stddev = sqrt(
  product(
    reduce(map(data, compose(square, partial(sum, neg(mean)))), sum),
    reciprocal(sum(data.length, neg(1))),
  ),
);
[mean, stddev]; // => [3, 2]
```

**Important:** This is not a style the author advocate for JavaScript programming,
but it is an interesting exercise to see how deeply functional JavaScript
code can be.

#### 8.8.4 Memoization

In functional programming, [this kind](#factorial) of caching is called memoization.
Below shows a higher-order function that accepts a function as its argument
and return a memozied version of the function

```js
function memoize(f) {
  const cache = new Map();
  return function (...args) {
    let key = args.length + args.join("+");
    if (cache.has(key)) {
      return cache.get(key);
    } else {
      let result = f.apply(this, args);
      cache.set(key, result);
      return result;
    }
  };
}

const factorial = memoize(function (n) {
  return n <= 1 ? 1 : n * factorial(n - 1);
});
factorial(5);
```

## 9 - Classes

In JavaScript, classes use prototype-based inheritance:
if two objects inherit properties from the same prototype, then we
say that those objects are instances of the same class.

`ES6` introduced a brand-new syntax (including a class keyword) that
makes it even easier to create classes.

### 9.1 Classes and Prototypes

<a id="factoryfuncton"></a>

> A factory function that returns a new range object

```js
function range(from, to) {
  let r = Object.create(range.methods);
  r.from = from;
  r.to = to;
  return r
}

range.methods = {
  includes(x) { return this.from <= x && x <= this.to},
  *[Symbol.iterator]() {
    for (let x = Math.ceil(this.from); x <= this.to; x++) {
      yield x;
      }
  },
  toString() { return "(" +this.from + "..." + this.to + ")"},
}

let r = range(1,3);
r.includes(2) // => true
r.toString() // => "{1...3}"
[...r] // => [1,2,3]
```

**Tip**: [Read here](#12) to understand what is Symbol.iterator

### 9.2 Classes and Constructors

A constructor is a function designed for the initialization
of newly created objects. Constructors are invoked using the `new` keyword.

The critical feature of constructor invocations is that the
`prototype` property of the constructor is used as the prototype of the new
object. An invocation of the Range() constructor automatically uses
Range.prototype as the prototype of the new Range object.

Finally, we can clarify this statement: "while most all objects have a prototype,
only a few objects have a `prototype` property": It is function objects that
have a `prototype` property. This means that all objects created with the
same constructor function inherit from the same object and are therefore
members of the same class.

Below is the idiomatic way to create a class in versions of JavaScript that do
not support the ES6 class keyword

<a id="override_prototype"></a>

```js
function Range(from, to) {
  this.from = from;
  this.to = to;
}

Range.prototype = {
  includes(x) {
    return this.from <= x && x <= this.to;
  },
  *[Symbol.iterator]() {
    for (let x = Math.ceil(this.from); x <= this.to; x++) {
      yield x;
    }
  },
  toString() {
    return "(" + this.from + "..." + this.to + ")";
  },
};

let r = new Range(1,3);
r.includes(2) // => true
r.toString() // => "{1...3}"
[...r] // => [1,2,3]
```

Because Range() constructor is invoked with new, it does not have to call Object.create().
Constructors do not even have to return the newly created object. Constructor
invocation automatically creates a new object, invokes the constructor as
a method of that object. **They usually won't work properly if they are
invoked as regular functions**.

**Tip**: This is a very common coding convention:
constructor functions define, in a sense, classes and classes have names that begin
with capital letters. Regular functions and methods have names that begin
with lowercase letters.

**Tip**: Arrow functions do not have a `prototype` property, which
means that they cannot be used as constructor functions for new classes.
Also, arrow functions inherit the `this` keyword from the context in
which they are defined rather than setting it based on the object through which
they are invoked.

> **Constructors and NEW.TARGET**

Within a function body, you can tell whether the function has been
invoked as a constructor with the special expression new.target.
If the value of that expression is defined, then you know that the function
was invoked as a constructor, with the `new` keyword, else it is undefined.

```js
function C() {
  if (!new.target) return new c();
  // initialization code goes here
}
```

**Tip**: JavaScript's various error constructors can be invoked without new.

**Important**: This technique only works for constructor defined in this old-fashioned
way. Classes created with the `class` keyword do not allow their constructors to be
invoked without new.

#### 9.2.1 Constructors, Class Identity, and instanceof

The expression `o instanceof C` evaluates to `true` if `o` inherits from
C.prototype. The inheritance need not be direct.

```js
function Strange() {}
Strange.prototype = Range.prototype;
new Strange() instanceof Range; // => true
```

If you want to test the prototype chain of an object for a specific
prototype and do not want to use the constructor function as an intermediary,
you can use `isPrototypeOf()` method. For example, there is no way to use
`instanceof` with the example in [function constructor without prototype](#factoryfuncton), we can test whether an object r
was a member of that constructor-less class with this code:

```js
range.methods.isPrototypeOf(r);
```

<a id="9.2.2"></a>

#### 9.2.2 The `constructor` Property

Any regular JavaScript function (excluding arrow functions, generator functions,
and async functions) can be used as a constructor. The value of the `constructor`
property is the function object:

```js
let F = function () {};
let p = F.prototype;
let c = p.constructor;
c === F; // => true

let o = new F();
o.constructor === F;
```

![Constructor And Prototype](images/constructor-and-prototype.png)

However, in [Example 9.2](#override_prototype), we overwrites the
predefined Range.prototype object with an object of its
own. So the new prototype object does not have a
`constructor` property. We can remedy this problem
by explicitly adding a constructor to the prototype:

```js
Range.prototype = {
  constructor: Range,
  // ...
};
```

Another common technique that you are likely to see
in older JavaScript code is adding methods to
predefined prototype object:

```js
Range.prototype.includes = function (x) {
  return this.from <= x && x <= this.to;
};
Range.prototype.toString = function () {
  return "(" + this.from + "..." + this.to + ")";
};
```

### 9.3 Classes with the class Keyword

`ES6`

```js
class Range {
  constructor(from, to) {
    this.from = from
    this.to = to
  }

  includes(x) {
    return this.from <= x && x <= this.to;
  }

  *[Symbol.iterator]() {
    for (let x = Math.ceil(this.from); x <= this.to; x++) {
      yield x;
    }
  }

  toString() {
    return "(" + this.from + "..." + this.to + ")";
  }
}

let r = new Range(1,3)
r.includes(2)
r.toString()
[...r]
```

Although it use the `class` keyword, the resulting object is a
constructor function.

Note the following things about the class syntax:

- Unlike object literals, no commas are used to separate the methods from each other.

- The keyword `constructor` is used to define the constructor function for the class.

- If your class does not need to do any initialization, you can
  omit the `constructor` keyword and its body, and an empty constructor function
  will be implicitly created for you.

If you want to define a class that subclasses -- or inherits from --
another class, you can use the `extends` keyword

```js
class Span extends Range {
  constructor(start, length) {
    if (length >= 0) {
      super(start, start + length);
    } else {
      super(start + length, start);
    }
  }
}
```

More about `class` syntax:

- All code within the body of a `class` declaration is implicitly in **strict mode**,
  even if no "use strict" directive appear.

- Unlike function declarations, class declarations are not "hoisted". You
  cannot instantiate a class before you declare it.

#### 9.3.1 Static Method

Static methods are defined as properties of the constructor function
rather than properties of the prototype object.

```js
class Range {
  static parse(s) {
    let matches = s.match(/^\((\d+)\.\.\.(\d+)\)$/);
    if (!matches) {
      throw new TypeError(`Cannot parse Range from "${s}`);
    }
    return new Range(parseInt(matches[1]), parseInt(matches[2]));
  }
  // ...
}

let r = Range.parse("(1...10)"); // OK
r.parse("(1...10)"); // TypeError
```

You'll sometimes see static methods called _class methods_ because they
are invoked using the name of the class/constructor.

**Tip**: It almost never makes sense to use the `this` keyword in a static
method.

#### 9.3.2 Getters, Setters, and other Method Forms

The only difference is that in class bodies, you don't put a comma after
the getter or setter.

#### 9.3.3 Public, Private, and Static Fields

The `ES6` standard only allows the creation of methods and static methods;
it does not include syntax for defining fields. If you wan to define a
field on a class instance, you must do that in the constructor function or
in one of the methods.

However, for extended class syntax that allows the definition of instance
and static fields, in both public and private forms. The syntax for
public instance fields is in common use by JavaScript programmers
using the React framework and the Babel transpiler.

```js
// From
class Buffer {
  constructor() {
    this.size = 0;
    this.capacity = 4096;
    this.buffer = new Uint8Array(this.capacity);
  }
}

// To
class Buffer {
  size = 0;
  capacity = 4096;
  buffer = new Uint8Array(this.capacity);
}
```

This syntax makes it clear to readers exactly what fields will hold the
state of each instance. You can declare fields without an initializer
by just writing the name of the field followed by a semicolon (initial value
will be `undefined`). **It is better style to always make the initial
value explicit for all your class fields**

To define private instance fields, prefix with '#', the field will be
usable within the class body but will be invisible and inaccessible
to any code outside of the class body:

```js
class Buffer {
  #size = 0;
  get size() {
    return this.#size;
  }
}
```

Note that the private fields must be declared using this new field syntax before
they can be used. You cannot write `this.#size = 0` in constructor.

We can add `static` before a public or private file declaration,
those fields will be created as properties of the constructor function
instead of properties of instance

```js
class Range {
  static integerRangePattern = /^\((\d+)\.\.\.(\d+)\)$/;
  static parse(s) {
    let matches = s.match(Range.integerRangePattern);
    if (!matches) {
      throw new TypeError(`Cannot parse Range from "${s}`);
    }
    return new Range(parseInt(matches[0]), parseInt(matches[2]));
  }
}
```

If we wanted this static field to be accessible only within the class,
write `static #name ...`

#### 9.3.4 Example: A Complex Number Class

```js
class Complex {
  #r = 0;
  #i = 0;
  constructor(real, imaginary) {
    this.#r = real;
    this.#i = imaginary;
  }

  plus(that) {
    return new Complex(this.#r + that.#r, this.#i + that.#i);
  }

  times(that) {
    return new Complex(
      this.#r * that.#r - this.#i * that.#i,
      this.#r * that.#i + this.#i * that.#r,
    );
  }

  static sum(c, d) {
    return c.plus(d);
  }
  static product(c, d) {
    return c.times(d);
  }

  get real() {
    return this.#r;
  }
  get imaginary() {
    return this.#i;
  }
  get magnitude() {
    return Math.hypot(this.#r, this.#i);
  }

  toString() {
    return `{${this.#r},${this.#i}}`;
  }

  equals(that) {
    return (
      that instanceof Complex && this.#r === that.#r && this.#i === that.#i
    );
  }
}

Complex.ZERO = new Complex(0, 0);
Complex.ONE = new Complex(1, 0);
Complex.I = new Complex(0, 1);
```

### 9.4 Adding Methods to Existing Classes

```js
Complex.prototype.conj = function () {
  return new Complex(this.r, -this.i);
};
```

The prototype object of built-in JavaScript classes is
also open like this. This is useful for implementing
new language features in older versions of the language.

```js
if (!String.prototype.startsWith) {
  String.prototype.startsWith = function (s) {
    return this.indexOf(s) === 0;
  };
}

// Another example
Number.prototype.times = function (f, context) {
  let n = this.valueOf();
  for (let i = 0; i < n; i++) f.call(context, i);
};
```

**Adding methods to the prototypes of built-in types like this is
generally considered to be a bad idea because it will cause confusion
and compatibility problems in the future**.

It is even possible to add methods to Object.prototype. **But it is
never a good thing to do because properties added to Object.prototype
are visible to `for/in` loops** (Unless you set enumerable to `false`)

```js
Object.prototype.hello = function() {
  console.log("hoho")
}
for (let item in Object.keys({})) {
  console.log(item)
} // hello

Object.defineProperty(Object.prototype, "hello", {enumerable: false});
for (let item in Object.keys({})) {
  console.log(item)
} // Nothing
```

### 9.5 Subclasses

The subclass constructor B() **must** typically invoke superclass constructor
A() in order to ensure that instances are completely initialized.

#### 9.5.1 Subclasses and Prototypes

```js
function Span(start, span) {
  if (span >= 0) {
    this.from = start;
    this.to = start + span;
  } else {
    this.to = start;
    this.from = start + span;
  }
}

Span.prototype = Object.create(Range.prototype);
Span.prototype.constructor = Span;
Span.prototype.toString = function () {
  return `(${this.from}...+${this.to - thisfrom})`;
};
```

In order to make Span a subclass of Range, we need to arrange for Span.prototype
to inherit from Range.prototype. The key line of code is:

```js
Span.prototype = Object.create(Range.prototype);
```

```js
function Range() {
}
Range.prototype = {}
Range.prototype.constructor = Range

function Span() {
}

Span.prototype = Object.create(Range.prototype)
Span.prototype.constructor = Span

let a = new Span
console.log(a instanceof Span) // => true
console.log(a instanceof Range) // => true
```

A robust subclassing mechanism needs to allow classes to invoke the
methods and constructor of their superclass, but prior to
ES6, JavaScript did not have a simple way to do these things.

Fortunately, `ES6` solves these problems with the `super` keyword as
part of the `class` syntax.

#### 9.5.2 Subclasses with `extends` and `super`

`ES6`

```js
class EZArray extends Array {
  get first() {
    return this[0];
  }
  get last() {
    return this[this.length - 1];
  }
}

let a = new EZArray();
a.push(1, 2, 3, 4);
a.first; // => 1
a.last; // => 4
```

```js
Array.prototype.isPrototypeOf(EZArray.prototype); // => true

// Static methods are inherited too
Array.isPrototypeOf(EZArray);
```

Below is a more fully fleshed-out example -- TypedMap

```js
class TypedMap extends Map {
  #keyType;
  #valueType;
  constructor(keyType, valueType, entries) {
    if (entries) {
      for (let [k, v] of entries) {
        if (typeof k !== keyType || typeof v !== valueType) {
          throw new TypeError(`Wrong type for entry [${k}, ${v}]`);
        }
      }
    }
    super(entries);
    this.#keyType = keyType;
    this.#valueType = valueType;
  }

  set(key, value) {
    if (this.#keyType && typeof key !== this.#keyType) {
      throw new TypeError(`${key} is not of type ${this.#keyType}`);
    }
    if (this.#valueType && typeof value !== this.#valueType) {
      throw new TypeError(`${value} is not of type ${this.#valueType}`);
    }
    return super.set(key, value);
  }
}
```

There are a few important rules of `super()`:

- If you define a class with the extends keyword, then the constructor
  for your class must use `super()` to invoke the superclass constructor

- If no constructor in your subclass, one with `super(...args)` will be defined
  automatically for you. Whatever values are passed to it will be passed to super()

- You may not use the `this` keyword in your constructor until **after
  you have invoked the superclass constructor with super()**. Else, it will throw
  ReferenceError when `new`.

- `new.target` is a reference to the constructor that was invoked. When
  a subclass constructor is invoked and use `super()` to invoke the superclass constructor,
  the superclass constructor **will see the subclass constructor as the value of new.target**. A well designed superclass should not need to know whether it has been subclassed,
  but it might be useful to be able to use new.target.name in logging messages, for example

  - ```js
    class ABC {
      constructor() {
        console.log(new.target);
        console.log(new.target.name);
      }
    }
    class DEF extends ABC {
      constructor() {
        super();
      }
    }
    new DEF();
    ```

  - Result :

    ```
    [class DEF extends ABC]
    DEF
    ```

#### 9.5.3 Delegation Instead of Inheritance

The `extends` keyword makes it easy to create subclasses. **But
that does not mean that you should create lots of subclasses**. It is
often easier and more flexible to get that desired behaviour into your class
by having your class create an instance of the other class and simply delegating
to that instance as needed -- composition.

```js
class Histogram {
  constructor() {
    this.map = new Map();
  }

  count(key) {
    return this.map.get(key) || 0;
  }

  has(key) {
    return this.count(key) > 0;
  }

  get size() {
    return this.map.size;
  }

  add(key) {
    this.map.set(key, this.count(key) + 1);
  }

  delete(key) {
    let count = this.count(key);
    if (count === 1) {
      this.map.delete(key);
    } else if (count > 1) {
      this.map.set(key, count - 1);
    }
  }

  [Symbol.iterator]() {
    return this.map.keys();
  }

  keys() {
    return this.map.keys();
  }

  values() {
    return this.map.values();
  }

  entries() {
    return this.map.entries();
  }
}
```

Because we used delegation rather than inheritance, a Histogram
object is not an instance of Set or Map. But Histogram
implements a number of commonly used Set methods, and in an untyped language
like JavaScript, that is often good enough: a formal inheritance relationship
is sometimes nice, but often optional.

#### 9.5.4 Class Hierarchies and Abstract Classes

Example below defines lots of subclasses, it also demonstrates how you can
define abstract class -- classes that do not include a complete implementation --
to serve as a common superclass for a group of related subclasses.

Note that JavaScript does not have any formal definition of abstract methods
or abstract classes.

```js
class AbstractSet {
  has(x) {
    throw new Error("Abstract method");
  }
}

class NotSet extends AbstractSet {
  constructor(set) {
    super();
    this.set = set;
    j;
  }

  has(x) {
    return !this.set.has(x);
  }

  toString() {
    return `{ x| x ∉ ${this.set.toString()} }`;
  }
}

class RangeSet extends AbstractSet {
  constructor(from, to) {
    super();
    this.from = from;
    this.to = to;
  }

  has(x) {
    return x >= this.from && x <= this.to;
  }

  toString() {
    return `{ x| ${this.from} ≤ x ≤ ${this.to} }`;
  }
}

class AbstractEnumerableSet extends AbstractSet {
  get size() {
    throw new Error("Abstract Method");
  }
  [Symbol.iterator]() {
    throw new Error("Abstract Method");
  }

  isEmpty() {
    return this.size === 0;
  }
  toString() {
    return `{${Array.from(this).join(", ")}}`;
  }
  equals(set) {
    if (!(set instanceof AbstractEnumerableSet)) return false;
    if (this.size != set.size) return false;

    for (let element of this) {
      if (!set.has(element)) return false;
    }

    return true;
  }
}

class SingletonSet extends AbstractEnumerableSet {
  constructor(member) {
    super();
    this.member = member;
  }

  has(x) {
    return x === this.member;
  }
  get size() {
    return 1;
  }
  *[Symbol.iterator]() {
    yield this.member;
  }
}

class AbstractWritableSet extends AbstractEnumerableSet {
  insert(x) {
    throw new Error("Abstract Method");
  }
  remove(x) {
    throw new Error("Abstract Method");
  }

  add(set) {
    for (let element of set) {
      this.insert(element);
    }
  }

  subtract(set) {
    for (let element of set) {
      this.remove(element);
    }
  }

  intersect(set) {
    for (let element of set) {
      if (!set.has(element)) {
        this.remove(element);
      }
    }
  }
}

class BitSet extends AbstractWritableSet {
  constructor(max) {
    super();
    this.max = max;
    this.n = 0;
    this.numBytes = Math.floor(max / 8) + 1;
    this.data = new Uint8Array(this.numBytes);
  }

  _valid(x) {
    return Number.isInteger(x) && x >= 0 && x <= this.max;
  }
  _has(byte, bit) {
    return (this.data[byte] & BitSet.bits[bit]) !== 0;
  }

  has(x) {
    if (this._valid(x)) {
      let byte = Math.floor(x / 8);
      let bit = x % 8;
      return this._has(byte, bit);
    } else {
      return false;
    }
  }

  insert(x) {
    if (this._valid(x)) {
      let byte = Math.floor(x / 8);
      let bit = x % 8;
      if (!this._has(byte, bit)) {
        this.data[byte] |= BitSet.bits[bit];
        this.n++;
      }
    } else {
      throw new TypeError("Invalid set element: " + x);
    }
  }

  remove(x) {
    if (this._valid(x)) {
      let byte = Math.floor(x / 8);
      let bit = x % 8;
      if (this._has(byte, bit)) {
        this.data[byte] &= BitSet.masks[bit];
        this.n--;
      }
    } else {
      throw new TypeError("Invalid set element: " + x);
    }
  }

  get size() {
    return this.n;
  }

  *[Symbol.iterator]() {
    for (let i = 0; i <= this.max; i++) {
      if (this.has(i)) {
        yield i;
      }
    }
  }
}

BitSet.bits = new Uint8Array([1, 2, 4, 8, 16, 32, 64, 128]);
BitSet.masks = new Uint8Array([~1, ~2, ~4, ~8, ~16, ~32, ~64]);
```

**Important**: `_` in the property name hints that it is for internal use only.

<a id="10"></a>

## 10 - Modules

As a practical matter, modularity is mostly about encapsulating or hiding private
implementation details and keeping the global namespace tidy so that
modules cannot accidentally modify the variables, functions and classes
defined by other modules.

`require()`-based modules are a fundamental part of the Node programming environment
but were never adopted as an official part of the JavaScript language.
Instead, `ES6` defines modules using `import` and `export` keywords.
Although `import` and `export` have been part of the languages for years, they
were only implemented by web browsers and Node relatively recently.
And, as a practical matter, JavaScript modularity still depends on code-bundling tools.

### 10.1 Modules with Classes, Objects, and Closures

JavaScript defines quite a few mathematical functions and constants, but instead of defining
them all globally, they are grouped as properties of a single global Math object.
e.g. Instead of defining global classes with names like SingletonSet and BitSet,
that example could have been written to Sets.Singleton and Sets.Bit

From the example of 9.5 Subclass, if we were writing that example as a module, we would
have wanted to:

- make isValid and has methods internal only accessible
- hide BitSet.bits and BitSet.masks

As we saw in [8.6 Closures](#8.6-closures), we can use immediately invoked funciton
expressions to achieve a kind of modularity by leaving the implementation details
and utility functions hidden within the enclosing function but making the public
API of the module the return value of the function.

```js
// example from 9.5 Subclass
const BitSet = (function() {
  function isValid(set, n) {...}
  function has(set,byte,bit){...}
  const BITS = new Uint8Array([1,2,4,8,16,32,64,128])
  const MASKS = new Uint8Array([~1,~2,~4,~8,~16,~32,~64,~128])
  return class BitSet extends AbstractWritableSet {
    // ... implementation omitted...
  }
}())
```

More interesting example when the module has more than one item in it:

```js
const stats = (function () {
  const sum = (x, y) => x + y;
  const square = (x) => x * x;
  function mean(data) {
    return data.reduce(sum) / data.length;
  }
  function stddev(data) {
    let m = mean(data);
    return Math.sqrt(
      data
        .map((x) => x - m)
        .map(square)
        .reduce(sum) /
        (data.length - 1),
    );
  }
  return { mean, stddev };
})();

stats.mean([1, 3, 5, 7, 9]); // => 5
stats.stddev([1, 3, 5, 7, 9]); // => Math.sqrt(10)
```

#### 10.1.1 Automating Closure-Based Modularity

Note that it is a fairly mechanical process to transform a file of
JavaScript code into this kind of module by inserting some text at the
beginning and end of the file.

A tool example:

```js
const modules = {};
function require(moduleName) { return modules[moduleName]}

modules["set.js"] = (function() {
  const exports = {};
  exports.BitSet = class BitSet {...}
  return exports
})()

modules["stats.js"] = (function() {
  const exports = {};
  const sum = (x, y) => x + y;
  const square = x => x*x
  exports.mean = function(data) {...}
  exports.stddev = function(data) {...}
  return exports
})()
```

With modules bundled up into a single file like this, you can imagine
writing code like the following to make use of those modules:

```js
const stats = require("stats.js")
const BitSet = require("sets.js").BitSet;

let s = new BitSet(100);
...
```

This code is a rough sketch of how code-bundling tools (such as
webpack and Parcel) for web browser work, and it's also a simple introduction to
the `require()` function like the one used in Node programs.

### 10.2 Modules in Node

In Node programming, it is normal to split programs into as many
files as seems natural. These files of JavaScript code are assumed to
all live on a fast filesystem. Unlike web browser, which have to read
files of JavaScript over a relatively slow network connection, there
is no need or benefit to bundling a Node program into a single JavaScript file.

**In Node, each file is an independent module with a private namespace.**
Constants, variables, functions, and classes defined in one file **are
private to that file unless the file exports them**.

Node modules import other modules with the `require()` function
and export their public API by setting properties of the **Exports object**
or by replacing the `module.export` object entirely.

#### 10.2.1 Node Exports

Node defines a global `exports` object that is **always defined**.

```js
const sum = (x, y) => x * y;
const square = (x) => x * x;

exports.mean = (data) => data.reduce(sum) / data.length;
exports.stddev = function (d) {
  let m = exports.mean(d);
  return Math.sqrt(
    d
      .map((x) => x - m)
      .map(square)
      .reduce(sum) /
      (d.length - 1),
  );
};
```

We can simply assign the single value:

```js
module.exports = class BitSet extends AbstractWritableSet {
  // implementation omitted
};
```

**Tip**: `module.exports` and `exports` refer to same object.

Another approach:

```js
const sum = (x, y) => x * y;
const square = (x) => x * x;
const mean = (data) => data.reduce(sum) / data.length;
const stddev = function (d) {
  let m = mean(d);
  return Math.sqrt(
    d
      .map((x) => x - m)
      .map(square)
      .reduce(sum) /
      (d.length - 1),
  );
};
module.exports = { mean, stddev };
```

#### 10.2.1 Node Imports

By calling the `require()` function.

The argument to this function is the name of the module to be imported,
and the return value is whatever value that modules exports.

If you want to import a system module **built in** to Node or a module that
you have installed on your system via a package manager, then you simply
use the unqualified name of the module, without any "/" characters (that would
turn it into a filesystem path):

```js
const fs = require("fs");
const http = require("http");

// A third-party module which is not part of Node but has been installed
// locally.
const express = require("express");
```

**When you want to import a module of your own code, the module
name should be the path to the file that contains that code, relative
to the current module's file starts with "./" or sometimes "../"**. It is
legal to use absolute paths that begin with a / character but it is unusual.

```js
const stats = require("./stats.js");
const BitSet = require("./utils/bitset.js");
```

**Tip**: You can omit the `.js` suffix, but it is common to see these file
extensions explicitly included

You can also import specific properties using destructuring assignment:

```js
const { stddev } = require("./stats.js");
```

#### 10.2.3 Node-Style Modules on the Web

Modules with an Exports object and require() function are built
in to Node. **But if you're willing to process your code with a bundling
tool like webpack**, then it is also possible to use this style of modules
for code that is intended to run in web browser. Until recently, this
was a very common thing to do, and you may see lots of web-based code
that still does it.

**But**

Now that JavaScript has its own standard module syntax `import` and `export`.
Developers who use bundlers are more likely to use the official JavaScript modules.

### 10.3 Modules in ES6

In `ES6`, JavaScript finally supports real modularity as a core language feature.
ES6 modularity is conceptually the same as Node modularity: **each file is
its own module**, constants, variables, functions and classes defined within
a file are private to that module unless they are explicitly exported.

The most obvious difference between ES6 module and regular JavaScript "scripts" is
in regular scripts, top-level declarations of variables, functions, and classes go into a single global
context shared by all scripts; with modules, each file has its own private context.
Code inside an ES6 module is **automatically in strict mode**. In modules
`this` is `undefined` even in top-level code. (By contrast, scripts in web browsers
and Node set `this` to the global object)

> **ES6 Modules On The Web And In Node**

ES6 modules have been in use on the web for years **with the help of code
bundlers** like webpack, which **combine independent modules of JavaScript
code into large, non-modular bundles suitable for inclusion into web pages**.

At the time of this writing, however, **ES6 modules are finally supported natively by
all web browser other than IE**. When use natively, ES6 modules are added
into HTML pages with a special `<script type="module">` tag.

And meanwhile, Node finds itself in the awkward position of having to support
two not entirely compatible module systems. Node 13 supports ES6 modules.

#### 10.3.1 ES6 Exports

Adding the keyword `export` before declaration

```js
export const PI = Math.PI;

export function degreesToRadians(d) {
  return (d * PI) / 1000;
}

export class Circle {
  constructor(r) {
    this.r = r;
  }
  area() {
    return PI * this.r * this.r;
  }
}
```

We could have equivalently written a single line at the end:

```js
export { PI, degreesToRadians, Circle };
```

**Important**: In this case, the curly braces **do not actually define
an object literal**. This export syntax simply requires a comma-separated list
of identifiers within curly braces.

It is common to write modules that export only one value (typically a
function or class), in this case, we usually use `export default` instead
of `export`:

```js
export default class BitSet {
  // implementation omitted
}
```

**Tip**: Default exports do not need to have name. It can export
any expression including anonymous function expressions and anonymous class
expressions.

It is legal, but somewhat uncommon, for modules to have a set of regular
exports and also a default export. If a module has a default export,
**it can only have one**.

**export keyword can only appear at the top level of
your JavaScript code.**

#### 10.3.2 ES6 Imports

The simplest form of import is used for modules
that **defined a default export**:

```js
import BitSet from "./bitset.js";
```

The identifier to which the imported value is assigned
**is a constant**. Like export, `imports` can only appear
at the top level of a module. By near-universal
convention, the imports needed by a module
are placed at the start of the module. Interestingly, however,
this is not required: like function declaration, **imports are hoisted to the top**.

A module specifier string must be an absolute path starting with "/",
or a relative path starting with "./" or "../", or a complete
URL with protocol and hostname. **ES6 spec** does not allow unqualified
module specifier strings like "util.js", because it is **ambiguous whether
this is intended to name a module in the same directory or some kind of system
module**.

To import values from a module that exports multiple values:

```js
import { mean, stddev } from = "./stats.js"
```

Recall that default exports **do not need to have a
name in the module that defines them**. Instead,
we provide a local name when we import those value.
But non-default exports of a module do have names
in the exporting module, so when we import those
values, we refer to them by those names.

Style guides sometimes recommend that you explicitly
import every symbol that your module will use. When
importing from a module that defines many exports,
however, you can easily import everything:

```js
import * as stats from "./stats.js";
```

This creates an object and assign it to a constant named stats.
Each of the non-default exports of the module being
imported becomes a property of this `stats` object. The functions
can be invoked by `stats.xxx()`

Again, this is legal, but uncommon, for a module to use both `export`
and `export default`. You can import both of them:

```js
import Histogram, { mean, stddev } from "./histogram-stats.js";
import { default as Histogram, mean, stddev } from "./histogram-stats.js";
```

We can `import` the modules that have no exports at all.

```js
import "./analytics.js";
```

A module like this runs the first time it is imported.
(**And subsequent imports do nothing**). The usage example
is an analytics module for a web application that
run code to register various event handlers and then
use those event handlers to send telemetry data
back to the server at appropriate times.

**Tip**: You can use import-nothing `import` syntax even with modules
that do have exports.

**Important**: If you import one from multiple JavaScript files,
that imported one only run **one** time. **The same goes to require() method**

#### 10.3.3 Imports and Exports with Renaming

If two modules export two different values using the same name
and you want to import both of those values, you will need to
rename one or both of the values when you import it.

```js
import { render as renderImage } from "./imageutils.js";
import { render as renderUI } from "./ui.js";
```

It is also possible to rename values as you export them:

```js
export { layout as calculateLayout, render as renderLayout };
```

Again, the `export` key word expects a single identifier before `as`, not
an expression. This means, unfortunately, that you cannot use export renaming like this:

```js
export { Math.sin as sin, Math.cos as cos}
```

#### 10.3.4 Re-Export

If we were writing such a module and we thought that many users of the module
would want only one function or the other, then we might want to
define `mean()` in a "./stats/mean.js" module and define `stddev()` in
"./stats/stddev.js". That way, programs only need to import exactly the functions
they need and are not bloated by importing code they do not need.

We might expect that there would be plenty of programs that want both functions and would
appreciate a convenient './stats.js' module which they could import both on
one line.

```js
import { mean } from "./stats/mean.js";
import { stddev } from "./stats/stddev.js";
export { mean, stddev };
```

Instead of importing a symbol simply to export it again, we can:

```js
export { mean } from "./stats/mean.js";
export { stddev } from "./stats/stddev.js";
```

If we are not being selective:

```js
export * from "./stats/mean.js";
export * from "./stats/stddev.js";
```

We can also define imported item as another name:

```js
export { mean, mean as average } from "./stats/mean.js";
export { stddev } from "./stats/stddev.js";
```

If we have done `export default` in individual files, we can:

```js
export { default as mean } from "./stats/mean.js";
export { default as stddev } from "./stats/stddev.js";
```

If you want to re-export a named symbol from another module as the default export:

```js
export { mean as default } from "./stats.js";
```

And re-export the default to default (funny huh)

```js
export { default } from "./stats/mean.js";
```

#### 10.3.5 JavaScript Modules on the Web

We'll be discussing how modules actually work in web browsers.

As of early 2020, production code using `ES6` modules is still generally bundled
with a tool like webpack. There are trade-offs to doing this: web apps
that have frequent incremental updates and users who make frequent return
visits may find that using small modules instead of large bundles can result
in better average load times because of better utilization of the user's browser
cache. But on the whole, code bundling tends to give a better performance.
**That may well change in the future as network speeds grow and browser
vendors continue to optimize their ES6 module implementations.**

Even though bundling tools may still be desirable in production, they
are no longer required in development since all current browsers
provide native support for JavaScript modules. If you want to natively use
`import` directives in a web browser, you must tell the web browser that your code
is a module by using a `<script type="module">` tag.

Using a `<script type="module">` tag to define the
main entry point for a modular JavaScript program can be as simple
as this:

```js
<script type="module">import "./main.js";</script>
```

Scripts with the `type="module"` attribute are loaded and
executed like scripts with the defer attribute. Loading of the code begins as
soon as the HTML parser encounters the `<script>` tag. But code
execution does not begin until HTML parsing is complete. And once HTML
is complete, scripts (both modular and non) are executed in the order in
which they appear in the HTML document.

You can modify the execution time of modules with the
`async` attribute. An `async` module will execute as soon as the
code is loaded, even if HTML parsing is not complete
and even if this changes the relative ordering of the scripts.
(By default, with `type="module"`, the behaviour is `defer` which
execute only after HTML parsing is completed)

Web browser that support `<script type="module>` must
also support `<script nomodule>`. Browser that are
module-aware ignore any script with the `nomodule`
attribute and will not execute it. Browser that
do not support modules will not recognize the
`nomodule` attribute, so they will ignore it and
run the script. This provides a powerful technique
for dealing with browser compatibility issues.
Browser that support `ES6` modules also support other
modern JavaScript features like classes, arrow functions,
and the `for/of` loop. If you write modern JavaScript
and load it with `<script type="module">`, you
know that it will only be loaded by browsers that can
support it. And as a fallback for IE11 (which in 2020 is
effectively the only remaining browser that does not
support ES6), you can use tools like Babel and webpack
to transform your code into non-modular code, then load
that less-efficient transformed code via `<script nomodule>`

Another important difference between regular scripts
and module scripts has to do with cross-origin loading. Unlike regular scripts,
`<script type="module">` allows modules only be
loaded from the same origin as the containing
HTML document or when proper `CORS` headers are in
place to securely allow cross-origin loads.
An unfortunate side effect of this new security restriction
is that it makes it difficult to test `ES6` modules
in browser development mode using `file: URLS`.

Some programmers like to use the filename extension `.mjs` to
distinguish their modular JavaScript files from their regular, non-modular
JavaScript files with traditional `.js` extension. For web browsers, the file extension is actually irrelevant. Node's support for `ES6`
does use the filename extension as a hint to distinguish which module system
is used by each file it loads. **So if you are writing `ES6` modules and want
them to be usable with Node, then it may be helpful to adopt the `.mjs` naming convention.**

<a id="10.3.6"></a>

#### 10.3.6 Dynamic Import with import()

On the web, code has to be transferred over a network instead of being
read from the filesystem. And once transferred, that code is often
executed on mobile devices with relative slow CPUs. This is not the
kind of environment where static module imports -- which require an entire
program to be loaded before any of it runs -- make a lot of sense.

It is common for web applications to initially load only enough of their code
to render the first page displayed to the users. Then, once the user
has some preliminary content to interact with, they can begin to load
the often much larger amount of code needed for the rest of the web app.
Web browsers make it easy to dynamically load code by using the DOM API to inject
a new `<script>` tag into the current HTML document, and web apps have
been doing this for many years.

Introduction of `import()` in `ES2020` -- you pass a module specifier to `import()`
and it returns a Promise object that represents the asynchronous process
of loading and running the specified module. Once the dynamic import is complete,
the Promise is "fulfilled" and produces an object like the one you would get
with the `import * as` form of the static import statement.

```js
// from
import * as stats from "./stats.js";

// to
import("./stats.js").then((stats) => {
  let average = stats.mean(data);
});
```

Or, in an `async` function:

```js
async analyzeData(data) {
  let stats = await import("./stats.js")
  return {
    average: stats.mean(data)
    stddev: stats.stddev(data)
  }
}
```

With `import()`, you are not constrained to use a constant string literal:
any expression that evaluates to a string in the proper form will do.

Dynamic import() looks like a function invocation, but it actually is not.
You'll notice that if you try writing code like `console.log(import)` or
`let require = import`

Finally, note that dynamic `import()` is not just
for web browsers. Code-packaging tools like
webpack can also make good use of it. The
most straightforward way to use code bundler **is to
tell it the main entry point for your program** and
let it find all the static `import` directives and
assemble everything into one large file. By
strategically using dynamic `import()` calls,
however, you can break that one monolithic bundle
up into a set of smaller bundles that can be loaded
on demand.

#### 10.3.7 import.meta.url

Within an `ES6` module (not regular script or a Node module), the
special syntax `import.meta` refers to an object that contains
**metadata about the currently executing module**.
The `url` property of this object is the URL from which the module
was loaded.

The primary use case of `import.meta.url` is to be able to refer
to images, data files, or other resources that are stored in the same
directory as (or relative) to module. The `URL()` constructor makes
it easy to resolve relative URL against an absolute URL (import.meta.url).

For example, your module that includes strings that need to be localized
and that the localization files are stored in `l10n/directory`, which is
in the same directory as the module itself:

```js
function localStringURL(locale) {
  return new URL(`l10/${locale}.json`, import.meta.url);
}
// Example =>
// URL {
//   href: 'http://abc.o/l10n',
//   origin: 'http://abc.o',
//   protocol: 'http:', // For Node, it is file://
//   username: '',
//   password: '',
//   host: 'abc.o',
//   hostname: 'abc.o',
//   port: '',
//   pathname: '/l10n',
//   search: '',
//   searchParams: URLSearchParams {},
//   hash: ''
// }
```

## 11 - The JavaScript Standard Library

### 11.1 Sets and Maps

`ES6`

#### 11.1.1 The Set Class

A set is a collection of values, like an array is. Unlike arrays,
however, sets are not ordered or indexed, and they do not allow
duplicates.

```js
let s = new Set();
let t = new Set([1, s]);
```

The argument to the `Set()` constructor need not be array:
any iterable object (including other Set objects) is allowed

```js
let t = new Set(s);
let unique = new Set("Mississippi"); // 4 elements: "M", "i", "s", "p"

unique.size; // => 4
```

```js
s.add(1); // add a number, {1}
s.add("1"); // add a string, {1,"1"}
s.add([2, 3, 4]); // {1, [2,3,4]}
s.add(true); // add a boolean
s.delete(1); // => true
s.delete("test"); // => false
s.has(true); // => true
s.clear();
```

**Tip**: If you want to add multiple values to set, you can use chained method add

**Important**: set membership is based on **strict equality checks**, `===`.
A set can contain both number 1 and the string "1", because it considers
them to be distinct values.

The `includes()` method of an array also performs membership testing, **but the time
it takes is proportional to the size of the array, and using an array as
a set can be much, much slower than using a real Set object**

The Set class is **iterable**:

```js
// you can use a for/of loop
let sum = 0;
for (let p of oneDigitPrimes) {
  sum += p
}
sum // => 17: 2 + 3 + 5 + 7

// you can use spread operator
[...oneDigitPrimes] // [2,3,5,7]
```

**JavaScript Set class always remember the order that elements were inserted in**,
**and it always uses this order when you iterate a set**.

Set class also implements a `forEach()` method

```js
let product = 1;
OneDigitPrimes.forEach((n) => {
  product *= n;
});
product; // => 210: 1 * 2 * 3 * 5 * 7
```

Any JavaScript value can be used as a key in a Set. This includes `null`, `undefined`, and `NaN`.

#### 11.1.2 The Map Class

```js
let m = new Map();
let n = new Map([
  ["one", 1],
  ["two", 2],
]);
```

You can also use the Map() constructor to copy other map or o copy
property names and values from an existing object:

```js
let copy = new Map(n);
let o = { x: 1, y: 2 };
let p = new Map(Object.entries(o));
```

```js
let m = new Map();
m.size;
m.set("one", 1);
m.get("two"); // => undefined
m.has("one"); // => true
m.delete("one"); // => true
m.delete("two"); // => false
m.clear();
```

Like Set, the `set()` method of Map can be chained, which allows
maps to be initialized without using arrays of arrays.

As with Set, any JavaScript value can be used as a key or a value in
a Map. This includes `null`, `undefined`, and `NaN`.

When we use spread operator with a Map object, we'll get an array of arrays.

When iterating a map with a `for/of` loop:

```js
let m = new Map([["x", 1], ["y", 2]])
[...m] // => [["x", 1], ["y", 2]]

for (let [key, value] of m) {
  ...
}
```

Like the Set class, the Map class iterates in insertion order.
**If you want to iterate just the keys or just the associated values
of a map, use the `key()` and `values()`.** `entries()` method
returns an iterable object and iterates key/value pairs, exactly
the same as iterating the map directly.

Map object can also be iterated using the `forEach()`, but please remember the
first parameter is value, and second is key (try to think of forEach in array,
the first is value, the second parameter is index):

```js
m.forEach((value, key) => {});
```

#### 11.1.3 WeakMap and WeakSet

The WeakMap class is a variant (not an actual subclass) of the Map class **that
does not prevent its key values from being garbage collected**. Garbage
collection is the process by which the JavaScript interpreter reclaims the memory
of objects that are no longer "reachable" and cannot be used by the program.
A regular map hold "strong" references to its key values, and they remain reachable
through the map, even if all other references to them are gone. The WeakMap, by
contrast, keeps "weak" references to its key values so that they are not
reachable through the WeakMap, and their presence in the map does no prevent
memory from being reclaimed.

Some significant differences between WeakMap and Map:

- WeakMap key must be objects (including arrays); primitive values are not
  subject to garbage collection and cannot be used as keys.

- WeakMap implements only the `get()`, `set()`, `has()` and `delete()` methods.
  It is not iterable and does not define `keys()`, `values()` or `forEach()`.
  If WeakMap was iterable, then its keys would be reachable and it wouldn't be weak.

- WeakMap does not implement the `size` property, because size of a WeakMap could
  change at any time objects are garbage collection.

The intended use of WeakMap is to allow you to associate values with objects
without causing memory leaks. For example, that you are writing a function
that takes an object argument and needs to perform some time-consuming computation
on that object. For efficiency, you'd like to cache the computed value for later reuse.
If you use Map object to implement the cache, you will prevent any of the objects
from ever being reclaimed, but by using a WeakMap, you avoid this problem. (You
can often achieve a similar result using a private Symbol property
to cache the computed value directly on the object [6.10.3 Symbols as Property Names](#6.10.3-symbols-as-property-names))

Same logics go for WeakSet. It is not frequently used: it use case are like
those for WeakMap. If you want to mark an object as having some special
property or type, for example, you can add it to a WeakSet.

<a id="11.2"></a>

### 11.2 Typed Arrays and Binary Data

#### 11.2.1 Typed Array Types

`ES6`

Typed arrays are not technically arrays (Array.isArray() returns `false` for them),
but they implement all of the array methods plus a few more of their own.

They differ from regular arrays:

- Typed arrays allow you to specify the type (signed and unsigned integers and
  IEEE-754 floating point) and size (8 bits to 64 bits) of the numbers to be stored in
  the array.

- You must specify the length of a typed array when you create it, **the length can never change**

- The elements of a typed array are always initialized to 0 when the array is created

| Constructor         | Numeric type                                             |
| ------------------- | -------------------------------------------------------- |
| Int8Array()         | signed bytes                                             |
| Uint8Array()        | unsigned bytes                                           |
| Uint8ClampedArray() | unsigned bytes without rollover                          |
| Int16Array()        | signed 16-bit short integers                             |
| Uint16Array()       | unsigned 16-bit short integers                           |
| Int32Array()        | signed 32-bit short integers                             |
| Uint32Array()       | unsigned 32-bit short integers                           |
| BitInt64Array()     | signed 64-bit BigInt values (ES2020)                     |
| BigUint64Array()    | unsigned 64-bit BigInt values (ES2020)                   |
| Float32Array()      | 32-bit floating-point value                              |
| Float64Array()      | 64-bit floating-point value: a regular JavaScript number |

Each of the typed array constructor has a BYTES_PER_ELEMENT property with the value
1, 2, 4, or 8, depending on the type.

**Tip**: ClampedArray functions example: when you defined a number that more than 255, that element become 255, when you defined a number that lesser than 0: the element became 0.

Remember that typed arrays have fixed lengths, so the methods that can change
the length (such as `push()`, `pop()`, `unshift()`, `shift()`, and `splice()`)
are not implemented for typed arrays.

Every typed array has three properties that relate to the underlying buffer:

```js
typedarr.buffer; // The ArrayBuffer object for a typed array
typedarr.byteOffset; // Where the view starts at
typedarr.byteLength;
typedarr.length * typedarr.BYTES_PER_ELEMENT === typedarr.byteLength; // => true
```

#### 12.2.2 Creating Typed Arrays

Call the appropriate constructor

```js
let bytes = new Uint8Array(1024); // 1024 bytes
let matrix = new Float64Array(9); // A 3x3 matrix
let point = new Int16Array(3); // A point in 3D space
let rgba = new Uint8ClampedArray(4); // A 4-byte RGBA pixel value
let sudoku = new Int8Array(81); // A 9x9 sudoku board
```

You can specify those values when you create the array by using `of()` or `from()`,
as mentioned in section Array:

```js
let white = Uint8ClampedArray.of(255, 255, 255, 0);
let ints = Uint32Array.from(white);
```

The values may be truncated in order to fit the type constraints of you array,
without warnings and error:

```js
Uint8Array.of(1.23, 2.99, 45000); // => new Uint8Array([1,2,200])
```

There is one more way to create typed arrays that involves the
ArrayBuffer type. An ArrayBuffer is an opaque reference to a chunk of memory.
You can create one with constructor; just pass in the number
of bytes of memory you'd like to allocate:

```js
let buffer = new ArrayBuffer(1024 * 1024);
buffer.byteLength; // => 1024*1024
```

The ArrayBuffer class does not allow you to read or write any of the
bytes that you have allocated. But you can create typed arrays that use
the buffer's memory and that do allow you to read and write that memory.

```js
// 1st argument: ArrayBuffer
// 2nd argument: a byte offset within the array buffer (optional)
// 3rd argument: array length in elements (not bytes) (optional, if non specified, until end of array)
let ints = new Uint32Array(buffer, 4, 100);
ints.length // => 100
```

One more thing to bear in mind about this form of the typed array constructor:
array **must be memory aligned**, if you specify a byte offset, the value should
be a multiple of the size of your type. e.g Int32Array() constructor requires
a multiple of four.

**Tip**: All typed arrays have an underlying ArrayBuffer, even if you do not
explicitly specify one.

**Doing numeric array indexing with ArrayBuffer does not give you access
to the bytes in the buffer**.

You can use buffer of another typed array to create other views:

```js
let bytes = new Uint8Array(1024); // 1024 bytes
let ints = new Uint32Array(bytes.buffer); // or 256 integers;
let floats = new Float64Array(bytes.buffer); // or 128 doubles;
```

#### 11.2.3 Using Typed Arrays

```js
// Return the largest prime smaller than n, using the sieve
// of Eratosthenes
function sieve(n) {
  let a = new Uint8Array(n + 1);
  let max = Math.floor(Math.sqrt(n));
  let p = 2;
  while (p <= max) {
    for (let i = 2 * p; i <= n; i += p) a[i] = 1;
    while (a[++p]); // The next unmarked index is prime
  }
  while (a[n]) n--;
  return n;
}
```

The code is exactly the same as it would be with a regular JavaScript array,
but using `Uint8Array()` instead of `Array()` makes the code run more than
four times faster and use eight times memory in my testing.

Accessing and manipulating data in a `Uint8Array` can be faster
compared to a regular JavaScript array because of the typed nature of the array

#### 11.2.4 Typed Array Methods and Properties

Typed arrays also implement a few methods of their own.

`set()`, `subarray()`

The `set()` method sets multiple elements of a typed array at once by copying
the elements of a regular or typed array into a typed array:

```js
let bytes = new Uint8Array(1024);
let pattern = new Uint8Array([0, 1, 2, 3]);
bytes.set(pattern); // Copy them to the start of another byte array
bytes.set(pattern, 4); // Copy them again at a different offset
bytes.set([0, 1, 2, 3], 8); // // Or just copy values direct from a regular array
bytes.slice(0, 12); // => new Uint8Array([0,1,2,3,0,1,2,3,0,1,2,3])
```

The `subarray()` method returns a portion of the array on which it is called (**shared memory**):

```js
let ints = new Int16Array([0, 1, 2, 3, 4, 5, 6, 7, 8, 9]);
let last3 = ints.subarray(ints.length - 3, ints.length);
last3; // =>  Int16Array(3) [7,8,9]
ints[ints.length - 1] = 10;
last3[2]; // => 10
```

<a id="11.2.5"></a>

#### 11.2.5 DataView and Endianness

For efficiency, typed arrays use the native endianness of the underlying
hardware. On little-endian systems, the bytes of a number are arranged in
an ArrayBuffer from least significant to most significant. On a big-endian
platforms, the bytes are arranged from most significant to least significant.

```js
// If the integer 0x00000001 is arranged in memory as 01 00 00 00, then
// we're on the little-endian platform. On a big-endian
// platform, we'd get bytes 00 00 00 01 instead
let littleEndian = new Int8Array(new Int32Array([1]).buffer)[0] === 1;
// => true
```

Today, the most common CPU architectures are little-endian. Many network
protocols, and some binary file formats, require big-endian byte ordering,
however. If you're using typed arrays with data that came from the network
or from a file, you can't just assume that the platform endianness matches
the byte order of the data. In general, when working with the external data,
you can use Int8Array and Uint8Array to view the data as an array
of individual bytes, but you should not use the other typed arrays with
multibyte word sizes. Instead you can use the DataView class, which defines
methods for reading and writing values from an ArrayBuffer with explicitly
specified byte ordering.

```js
let bytes = new Int32Array([1,2,3,4,5])
let view = new DataView(bytes.buffer, bytes.byteOffset, bytes.byteLength)

view.getInt32(0) // => 16777216; Read big-endian signed int from byte 0
view.getInt32(4,false) // => 33554432; Read big-endian signed int from byte 0
view.getUint32(8, true) // => 3: Read little-endian unsigned
view.setUint32(8, 123123, false); // Write it back in big-endian format
```

DataView defines 10 get methods for each of the 10 typed array (excluding Uint8ClampedArray). e.g getBigInt64(). If the second argument is omitted or `false`, big-endian
byte ordering is used. Other than `getInt8()` and `getUint8()`, accept an optional boolean
value as their second argument. If the second argument is `true`, little-endian ordering is
used.

DataView also defines 10 corresponding Set methods. Except `setInt8()` and `setUint8()`,
methods accepts an optional third argument. If the argument is omitted or `false`, the
value is written in big-endian format with the most significant byte first. If the
argument is `true`, the value is written in little-endian format with the least
significant byte first.

Typed arrays and the DataView class give you all the tools you need to process binary
data and enable you to write JavaScript programs that do things like decompressing
ZIP files or extracting metadata from JPEG files.

<a id="11.3"></a>

### 11.3 Pattern Matching with Regular Expressions

#### 11.3.1 Defining Regular Expressions

RegExp object maybe created with the RegExp(), but they are more often
created using a special literal syntax:

```js
let pattern = /s$/;
let pattern = new RegExp("s$");
```

##### LITERAL CHARACTERS

JavaScript regular expression syntax also supports certain nonalphabetic
characters through escape sequences that begin with a backslash(\).

| Character              | Matches                                                                                                                                                                                                 |
| ---------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Alphanumeric character | itself                                                                                                                                                                                                  |
| \0                     | The NUL character (\u0000)                                                                                                                                                                              |
| \t                     | Tab(\u0009)                                                                                                                                                                                             |
| \n                     | Newline (\u000A)                                                                                                                                                                                        |
| \v                     | Vertical tab (\u000B)                                                                                                                                                                                   |
| \f                     | Form feed (\u000C)                                                                                                                                                                                      |
| \r                     | Carriage return (\u000D)                                                                                                                                                                                |
| \xnn                   | The Latin character specified by the hexadecimal number nn, e.g. \x0A is the same as \n                                                                                                                 |
| \uxxxx                 | Unicode character specified by hexadecimal number xxxx; for example, \u0009 is the same as \t                                                                                                           |
| \u{n}                  | The Unicode character specified by the codepoint n, where n is one to six hexadecimal digits between 0 and 10FFFF. Note that this syntax is only supported in regular expressions that use the `u` flag |
| \cX                    | The control character ^X; for example, \cJ is equivalent to the newline character \n                                                                                                                    |

A number of punctuation characters have special meanings in regular expression:

```js
^ $ . * + ? = ! : | \ / ( ) [ ] { }
```

As a general rule, however, if you want to include any of these punctuation chatacters
literally in a regular expression, you must precede them with a `\`. Other
punctuation characters, such as quotation marks and @, do not have special meaning
and simply match themselves literally in a regular expression. **If you can't
remember exactly which punctuation characters need to be escaped, you can safely
place a backslash before any puntuation character**.

##### CHARACTER CLASSES

```js
/[abc]/ match any one of the letters a, b, or c
/[^abc]/ match any one of the letters NOT a, b, and c
/[a-zA-Z0-9]/ match any letter or digit from the Latin alphabet
/[\s]/ match the space character, tab character, and any other Unicode whitespace character
/[\S]/ match any character that is NOT Unicode whitespace
/[\u0400-\u04FF] match Unicode character range
/[\w]/ match any ASCII word character. Equivalent to [a-zA-Z0-9_]
/[\W]/ match any character that is not ASCII word character
/[\d]/ match ASCII digit. Equivalent to [0-9]
/[\D]/ match character other than an ASCII digit.
/[\b]/ A literal backspace
```

> **UNICODE CHARACTER CLASSES**

In `ES2018`, if a regular expression uses the `u` flag, then character class
`\p{...}` and its negation `\P{...}` are supported. (As of early 2020, this
is implemented by Node, Chrome, Edge, and Safari, but not Firefox) (Update: as of
30/12/2023, the day of writing note, Firefox has implemented it). These
character classes are based on properties defined by the Unicode standard,
and the set of characters they represent may change as Unicode evolves.

The \d character class matches only ASCII digit. If you want to match one decimal
digit from any of the world's writing systems, you can use /\p{Decimal_Number}/u.
If you want to match any number-like character, including fractions and roman numerals,
you can use \p{Number}.

The \w character class only works for ASCII, but with \p, we can approximate
an internationalized version like this:

```js
/[\p{Alphabetic}\p{Decimal_Number}\p{Mark}]/u;
```

As a final example, the `\p` syntax also allows us to define regular expressions
that match characters from a particular alphabet or script:

```js
let greekLetter = /\p{Script=Greek}/u;
let cyrillicLetter = /\p{Script=Cyrillic}/u;
```

##### REPETITION

| Character | Meaning                                                                                                          |
| --------- | ---------------------------------------------------------------------------------------------------------------- |
| {n,m}     | Match the previous item at least n times but no more than m times.                                               |
| {n,}      | Match the previous item n or more times.                                                                         |
| {n}       | Match exactly n occurrences of the previous item.                                                                |
| ?         | Match zero or one occurrences of the previous item. That is, the previous item is optional. Equivalent to {0,1}. |
| +         | Match one or more occurrences of the previous item. Equivalent to {1,}.                                          |
| \*        | Match zero or more occurrences of the previous item. Equivalent to {0,}.                                         |

##### NON-GREEDY REPETITION

The repetition above match as many times as possible while still allowing
parts of the regular expression to match. It is also possible to specify
that repetition should be done in a non-greedy way. Simply follow
the repetition character or characters with a question mark:
`??`, `+?`, `*?` or even `{1,5}?`.

For example, `/a+/.test("aaa")` matches all three letters. But
`/a+?/` matches one or more occurrences of the letter a, matching as
few characters as necessary. When applied to the same string, this pattern
matches only the first letter a.

```js
"baaaa".replace(/ba+?/, "ad"); // => 'adaaa'
"baaaa".replace(/ba+/, "ad"); // => 'ad'
```

**Tip**: In vim, use `a.\{-}` as non greedy

However, regular expression pattern matching is done by finding the first position
in the string at which a match is possible, the below case demonstrates:

```js
"aaab".replace(/a+?b/, "ad"); // => 'ad'
"aaab".replace(/a+b/, "ad"); // => 'ad'
"baaab".replace(/ba+?b/, "ad"); // => 'ad'
```

##### ALTERNATION, GROUPING, AND REFERENCES

> `|` - Alternation

```js
/ab|cd|ef/.test("ab") // =>  true
/ab|cd|ef/.test("cd") // =>  true
/ab|cd|ef/.test("de") // =>  false
```

> `()` - Grouping

One of the parentheses in regular expressions is grouping.

```js
/java(script)?/.test("java") // => true
/java(script)?/.test("javascript") // => true
/java(script)?/.test("jav") // => false
```

> `()` - References

For example, you can extract the digits from `/[a-z]+(\d+)/`. A related use
is to allow you to **refer back to a subexpression later in the
same regular expression**. `\1` refers back
to the first subexpression. **Because subexpressions can be nested within
others, it is the position of the left parenthesis that is counted**. For example,
`/([Jj]ava([Ss]Script)?)\sis\s(fun\w*)`, has ([Ss]cript) as `\2`.

A reference to a previous subexpression of a regular expression
does **NOT** refer to the pattern for that pattern for that subexpression but
rather to the text that matched the pattern:

```
/(['"])[^'"]*\1/.test(`'hello'`) // => true
/(['"])[^'"]*\1/.test(`'hello"`) // => false
```

Note, you cannot use a reference **within a character class**:

```js
/(['"'][^\1]*\1/)
Uncaught SyntaxError: Invalid regular expression: /(['"'][^\1]*\1/: Unterminated group
```

It is also possible to group items in a regular expression without
creating a numbered reference to those items. We can begin the group
with `(?:` and end with `)`:

```js
/([Jj]ava(?:[Ss]cript)?)\sis\s(fun\w*)/;
```

`(?:[Ss]cript)` is **used simply for grouping. These modified parentheses do not produce
a reference**, so in above regular expression, `\2` refers to the text matched by `(fun\w*)`

> **NAMED CAPTURE GROUPS**

`ES2018` standardizes a new feature that can make regular expressions more self-documenting
and easier to understand. This new feature is known as "named capture groups" and it
allows us to associate a name with each left parenthesis in a regular expression
so that we can refer to matching text by name rather than by number. As of early 2020, this
is implemented by Node, Chrome, Edge, and Safari, but not Firefox. (Firefox implemented it as of 31/12/2023)

```js
/(?<city>\w+) (?<state>[A-Z]{2}) (?<zipcode>\d{5})(?<zip9>-\d{4})?/;
```

We can do backreference `\k<reference>`:

```js
/(?<quote>['"])[^'"]*\k<quote>/;
```

##### SPECIFYING MATCH POSITION

The most commonly used anchor elements are `^` and `$`.

If you wan to search for "Java" as a word by itself, you can try the pattern
`/\sJava\s/` but this does not match "Java" at the beginning or the end of a string (because of no space).
The second problem with this pattern is when this pattern does find a match,
the matched string it returns has leading and trailing space. So instead of
matching actual space characters with `\s`, match (or anchor to) word
boundaries with `\b`.

**Tip**: `\B` anchors the match to a location that is not a word boundary.
Thus, the patter `/\B[Ss]cript/` matches "JavaScript" but not "Script"

You can also use arbitrary regular expressions as anchor conditions.
If you include an expression within `(?=` and `)` characters, it is a
lookahead assertion, require that the following characters match the
pattern, but do not include those characters in the match:

```js
/[Jj]ava([Ss]cript)?(?=\:)/.test("JavaScript: The Definitive Guide") // => true
/[Jj]ava([Ss]cript)?(?=\:)/.test("Java: The Definitive Guide") // => true
/[Jj]ava([Ss]cript)?(?=\:)/.test("Java in a Nutshell") // => false
```

The third expression is false because "Java" is not followed by a colon.

Instead, `(?!` and `)` is a negative lookahead assertion: the following characters
must not match

```js
/Java(?!Script)([A-Z]\w*)/.test("JavaHello") // => true
/Java(?!Script)([A-Z]\w*)/.test("JavaScript") // => false
```

> **LOOKBEHIND ASSERTIONS**

`ES2018` extends regular expression syntax to allow "lookbehind" assetion.
As of early 2020, these are implemented in Node, Chrome and Edge, but not Firefox
or Safari. (Update: as of 30/12/2023, Firefox have implemented it)

```js
/(?<=[\d]{2})\d{5}/.test("322222") // => false
/(?<=[\d]{2})\d{5}/.test("3322222") // => true
```

You could match a string of digits that is not preceded by a Unicode currency symbol
with a negative lookbehind assertion like this:

```js
/(?<![\p{Currency_Symbol}\d.])\d+(\.\d+)?/u.test("USD1f33.33"); // => true
/(?<![\p{Currency_Symbol}\d.])\d+(\.\d+)?/u.test("$2.33"); // => false
```

##### FLAGS

Six possible flags

- `g`

  - "global"
  - Find all matches within a string rather than just find the first match
  - alter the behavior of the String `match()` and the RegExp `exec()` method

- `i`

  - "case-insensitive"

- `m`

  - "multiline"
  - The RegExp will be used with multiline strings
  - changes the behavior of the "^" and "$" anchors to match the start and end of each line within a multiline string, rather than just the start and end of the entire string.

- `s`

  - Normally, a "." in a regular expression matches any character except a line terminator.
    When the `s` flag is used, "." will match any character, including the terminators.
  - `ES2018`, as of early 2020, is supported in Node, Chrome, Edge, and Safari, but
    not Firefox. (Update: as of 30/12/2023, Firefox supports it)
  - ```js
    /./.test("\n") // => false
    /./s.test("\n") // => true
    ```

- `u`

  - "Unicode"
  - `ES6`
  - Makes the regular expression match full Unicode codepoints rather than
    matching 16-bit values.
  - **without u, "." character matches any 1 UTF-16 16-bit value. With the flag, "."
    matches one Unicode codepoint, including those that have more than 16bits**
  - ```js
    /\u{1F601}/u.test("\u{1F601}") // => true
    /\u{1F601}/.test("\u{1F601}") // => false
    ```

- `y`

  - "sticky"
  - indicates that the regular expression is "sticky" and should match at
    the beginning of a string or at the first character following the previous match.
  - When used with a regular expression that is designed to find a single match,
    it effectively treats that regular expression as if it begins with ^ anchor it
    to the beginning of string.
  - It is more useful with regular expressions that are used repeatedly to find
    all matches within a string.
  - It causes special behaviors of the String `match()` method and the RegExp
    `exec()` method to enforce that each subsequent match is anchored to the string position
    at which the last one ended.

#### 11.3.2 String Methods for Pattern Matching

##### `search()`

This method takes a regular expression argument and
returns either the character position of the start of the first matching
substring or −1 if there is no match

```js
"JavaScript".search(/script/iu); // => 4
"Python".search(/script/iu); // => -1
```

If the argument to `search()` is not a regular expression, it is first
converted to one by passing it to the RegExp constructor.
**search() does not support global searches; it ignores the `g` flag of
its regular expression argument.**

##### `replace()`

Perform search and replace operation.

If the pattern has the `g` flag set, the `replace()` replaces all matches; otherwise
it replaces only the first match it finds.

```js
"javascRIPT".replace(/javascript/gi, "JavaScript") => JavaScript
'He said "stop" and "run"'.replace(/"([^'"]*)"/g, "<<$1>>") // => He said <<stop>> and <<run>>
```

If your pattern uses named capture group:

```js
'He said "stop"'.replace(/"(?<quotedText>[^'"]*)"/g, "<<$<quotedText>>"); // => He said <<stop>>
```

Instead of passing a replacement string as the second argument to `replace()`,
you can also pass a function that will be invoked to compute the replacement value.

The function includes:

- 1st argument: entire matched text

- Followed by arguments:

  1. capturing groups if any
  2. position within the string at which the match was found
  3. the entire string that `replace()` was called on

    - ```js
      "hello world hello world".replace(/(hello) (world)/g, function(v, ...args){
        console.log(v);
        console.log(args);
        return args[1] + " " + args[0]
      }) // => 'world hello world hello'

      // log:
      hello world // v
      [ 'hello', 'world', 0, 'hello world hello world' ] // args
      hello world // v
      [ 'hello', 'world', 12, 'hello world hello world' ] // args
      ```

  4. If the RegExp contained any named capture groups, then object whose property names match the capture group names and whose values are the matching text.

    - ```js
      `'SingleQuoted' is not "DoubleQuoted"`.replace(/'(?<single>.*)'/g, (v, ...args) => {
        console.log(v)
        console.log(args)
        return v
      })

      // log
      'SingleQuoted' // v
      [
        'SingleQuoted',
        0,
        `'SingleQuoted' is not "DoubleQuoted"`,
        [Object: null prototype] { single: 'SingleQuoted' }
      ]
      ```

##### `MATCH()`

Returns an array that contains the results of the match, or `null` if no match is found

```js
"7 plus 8 equals 15".match(/\d+/g); // => ["7", "8", "15"]
```

If the `match()` is used without `g` flag, then the output's element are:

- matched string
- any remaining elements are the substrings matching the parenthesized capturing groups of the regular expression.
- If the regular expression contains named capture groups, then the returned
  array also has a `groups` property whose value is an object

```js
"7 plus 8 equals 15".match(/\d+/); // => [ '7', index: 0, input: '7 plus 8 equals 15', groups: undefined ]

"7 'plus' 8 equals 15".match(/\d+ '(.*)'/); // =>
[
  "7 'plus'",
  'plus',

  index: 0,
  input: "7 'plus' 8 equals 15",
  groups: undefined
]
length: 2

"7 'plus' 8 equals 15".match(/\d+ '(?<quoted>.*)'/); // =>
[
  "7 'plus'",
  'plus',

  index: 0,
  input: "7 'plus' 8 equals 15",
  groups: [Object: null prototype] { quoted: 'plus' }
]
length: 2
```

An Example (**Parsing URLs with regular expressions is not a good idea. See [§11.9 -- URL API](#11.9) for a more robust
URL parser**):

```js
// A very simple URL parsing RegExp
let url = /(\w+):\/\/([\w.]+)\/(\S*)/;
let text = "Visit my blog at http://www.example.com/~david";
let match = text.match(url);
let fullurl, protocol, host, path;
if (match !== null) {
  fullurl = match[0]; // fullurl ==
  ("http://www.example.com/~david");
  protocol = match[1]; // protocol == "http"
  host = match[2]; // host == "www.example.com"
  path = match[3]; // path == "~david"
}

// rewrite into named captured group
let url = /(?<protocol>\w+):\/\/(?<host>[\w.]+)\/(?<path>\S*)/;
let text = "Visit my blog at http://www.example.com/~david";
let match = text.match(url);
let fullurl, protocol, host, path;
if (match !== null) {
  fullurl = match[0];
  protocol = match.groups.protocol;
  host = match.groups.host;
  path = match.groups.path;
}
```

If a RegExp has both the `g` and `y`. It return in array with
the first match begin at the start of the string.

```js
let pattern = /\d+ /gy;
"7 plus 8 equals 15".match(pattern); // => [ '7' ]
"7 8 plus 8 equals 15".match(pattern); // => [ '7', '8' ]
"a7 plus 8 equals 15".match(pattern); // => null
```

If the `y` flag is set without `g`, then `match()` tries to find a single match, and by default,
this match is constrained to the start of the string. You can change this
default match start position by setting the `lastIndex` property of the RegExp object.
If a match is found, then this `lastIndex` will be automatically updated to the
first character after match, **so if you call `match()` again, it will look for
a subsequent match)**.

```js
let vowel = /[aeiou]/y;
"test".match(vowel); // => null: "test" does not begin with a vowel
vowel.lastIndex = 1;
"test".match(vowel); // => "e"
vowel.lastIndex; // => 2
"test".match(vowel); // => null
vowel.lastIndex; // => 0: lastIndex gets reset after failed match
```

**Tip**: Passing a non-global regular expression to the `match()` method of string is same as
passing the string to `exec()` method of RegExp: the returned array and its properties
are the same in both cases

##### `MATCHALL()`

`ES2020`

`matchAll()` expects a RegExp with the `g` flag set. Instead of returning an array
of matching substrings, however, it returns and iterator that yields the kind
of match objects that `match()` returns when used with a non-global RegExp.

```js
const words = /\b\p{Alphabetic}+\b/gu;
const text = "This is a naive test of the matchAll() method"
for (let word of text.matchAll(words)) {
  console.log(`Found '${word[0]}' at index ${word.index}`)
}

Found 'This' at index 0
Found 'is' at index 5
Found 'a' at index 8
Found 'naive' at index 10
Found 'test' at index 16
Found 'of' at index 21
Found 'the' at index 24
Found 'matchAll' at index 28
Found 'method' at index 39
```

**Important**: The other properties are **exactly same as `match()` without g**

You can set the `lastIndex` property of a RegExp object to tell `matchAll()`
what index in the string to begin matching at. However, `matchAll()` never
modifies the `lastIndex` property of the RegExp you call it on, makes it
much less likely to cause bugs in you code.

##### `SPLIT()`

`split()` can be used with a string argument, it can also take a regular expression

```js
"1, 2, 3,\n4, 5".split(/\s*,\s*/); // => ["1","2","3","4","5"]
```

If you can `split()` with a RegExp delimiter ant the regular expression includes
capturing groups, then the text matches the capturing groups will be included in the
returned array.

```js
const htmlTag = /<([^>]+)>/;
"Testing<br/>1,2,3".split(htmlTag); // => ["Testing", "br/", "1,2,3"]
```

#### 11.3.3 The RegExp Class

The RegExp() constructor accepts one or two string arguments. The first
rgument is a string that contains the body of the regular expression. The
second is optional, it indicates the regular expression flags

```js
let zipcode = new RegExp("\\d{5}", "g");
```

**Warning**: Note that you must replace each `\` character with `\\`

The `RegExp()` constructor is useful when a regular expression is
being dynamically created and thus cannot be represented with the
regular expression literal syntax.

##### REGEXP PROPERTIES

| Property   | Description                                                                                                        |
| ---------- | ------------------------------------------------------------------------------------------------------------------ |
| source     | Read-only property: characters that appear between the slashes                                                     |
| flags      | Read-only property: the set of letters that represent the flags for the RegExp                                     |
| global     | Read-only property: `g` is set                                                                                     |
| ignoreCase | Read-only property: `i` is set                                                                                     |
| multiline  | Read-only property: `m` is set                                                                                     |
| dotAll     | Read-only property: `s` is set                                                                                     |
| unicode    | Read-only property: `u` is set                                                                                     |
| sticky     | Read-only property: `y` is set                                                                                     |
| lastIndex  | R/W integer, for patterns with the `g` or `y`, specifies the character position at which the next search is begin. |

##### `TEST()`

Returns `true` if the string matches the pattern or `false` if it does not match

```js
/a/.test("a"); // => true
```

If you use `test()` with a RegExp that use the `g` or `y` flags, then its
behavior depends on the value of the `lastIndex` property of the RegExp object

##### `EXEC()`

It returns an array just like the array returned by the `match()` method
for non-global searches.

Unlike the String `match()` method, `exec()` returns the same kind of
array whether or not the regular expression has the global `g` flag.
It always returns a single match and provides complete information about that match.
When `exec()` is called on a regular expression that has either the global `g` flag
or the sticky `y` flag set, it consults the `lastIndex` property of the RegExp
object to determine where to start looking for a match. Each time `exec()` successfully
finds a match, it updates the `lastIndex` property to the index of the character
immediately after the matched text. If `exec()` fails to find a match, it reset
`lastIndex` to 0. This special behavior allows you to call `exec()` repeatedly
in order to loop through all the regular expression matches in a string (Although,
as we've described, in `ES2020`, the `matchAll()` method of string is an
easier way to loop through all matches)

```js
let pattern = /Java/g;
let text = "JavaScript > Java";
let match;
while ((match = pattern.exec(text)) !== null) {
  console.log(`Matched ${match[0]} at ${match.index}`);
  console.log(`Next search begins at ${pattern.lastIndex}`);
}
```

> **The LastIndex Property and RegExp Reuse**

You need to be particularly careful when calling the match(), exec(), and test() methods. The `lastIndex` makes it easy to write buggy code.

```js
let match,
  position = [];
while ((match = /<p>/g.exec(html)) !== null) {
  position.push(match.index);
}
```

The above problem is we use a RegExp literal in the while loop condition as it always
start from `lastIndex` 0.

On the other hand, sometimes reusing a `RegExp` object is the wrong thing to do.

```js
let dictionary = ["apple", "book", "coffee"];
let doubleLetterWords = [];
let doubleLetter = /(\w)\1/g;

for (let word of dictionary) {
  if (doubleLetter.test(word)) {
    doubleLetterWords.push(word);
  }
}
doubleLetterWords; // => ["apple", "coffee"]: "book" is missing!
```

We could fix this problem by removing the `g` flag, or by moving the
RegExp literal into the body of the loop.

**The moral here is that `lastIndex` makes the RegExp API error prone. As in
`ES2020` and later, use the String `matchAll()` method instead of `exec()` to
sidestep this problem since `matchAll()` does not modify `lastIndex`.**

### 11.4 Dates and Times

```js
let now = new Data(); // The current time
```

With one numeric argument, the `Date()` constructor interprets
that argument as the number of milliseconds since the 1970 epoch:

```js
let epoch = new Date(0); // Midnight, January 1st, 1970, GMT
```

If you specify two or more integer arguments, they are interpreted as
the year, month, day-of-month, hour, minute, second, and millisecond in your
**local time zone**. If you want to specify a date and time
in UTC (Universal Coordinated Time, aka GMT), then you can use the
Date.UTC() static method, which takes the same arguments as the Date() constructor,
and return the millisecond timestamp.

```js
let century = new Date(2100, 0, 1, 2, 3, 4, 5);
let century = new Date(Date.UTC(2100, 0, 1, 2, 3, 4, 5));
```

**One quirk of the Date API is that first month of a year is number 0.**

| Method               | Example                                           |
| -------------------- | ------------------------------------------------- |
| console.log(date)    | 2099-12-31T18:03:04.005Z                          |
| toString()           | Fri Jan 01 2100 02:03:04 GMT+0800 (Malaysia Time) |
| valueOf()            | 4102423384005                                     |
| toUTCString()        | Thu, 31 Dec 2099 18:03:04 GMT                     |
| toISOString()        | 2099-12-31T18:03:04.005Z                          |
| toLocaleString()     | 1/1/2100, 2:03:04 AM                              |
| toLocaleDateString() | 1/1/2100                                          |
| toLocaleTimeString() | 2:03:04 AM                                        |

If you pass a string to the `Date()` constructor, it will attempt to parse that
string
as a date and time specification. The constructor can
parse dates specified in the formats produced by the `toString()`,
`toUTCString()` and `toISOString()` method.

Following methods allow you to query and modify the year, month, day-of-month,
hour, minute, second, and millisecond fields of the Date. **Each of these
methods has two forms: one that gets or sets using local time and one that
gets or sets using UTC time**.

| Method             |
| ------------------ |
| `getFullYear()`    |
| `getUTCFullYear()` |
| `setFullYear()`    |
| `setUTCFullYear()` |

Options: `FullYear`, `Month`, `Date` (day-of-month), `Hours`, `Minutes`, `Seconds`, `Milliseconds`, `Day` (day-of-week, read-only, 0 for Sunday, 6 for Saturday)

```js
let d = new Date();
d.setFullYear(d.getFullYear() + 1);
```

Some of the date set methods allow you to set more than one field at a time.
`setFullYear()` and `setUTCFullYear()` also optionally allow you to set
the month and day-of-month as well. And `setHours()` and `setUTCHours()`
allow you to specify the minutes, seconds, and milliseconds fields in addition
to the hours field.

```js
century.setFullYear(2333, 5, 3);
century.setHours(3, 5, 6, 7);
```

#### 11.4.1 Timestamps

Integers are as large as 8,640,000,000,000,000, so JavaScript won't be
running out of milliseconds for more than 270,000 years.

For any Date object, the `getTime()` returns this internal value, and
the `setTime()` method sets it. So you can add 30 seconds to a Date:

```js
d.setTime(d.getTime() + 30000);
```

The static Date.now() returns the current time as a timestamp and is helpful
when you want to measure how long your code takes to run:

```js
let startTime = Date.now();
reticulateSpline();
let endTime = Date.now();
console.log(`Spline reticulation took ${endTime - startTime}ms`);
```

> **High-Resolution TimeStamps**

The `performance.now()` returns a millisecond-based timestamp which includes fractions
of a millisecond. It indicates how much time has elapse since a web page was
loaded or since the Node process started.

The `performance` object is part of a larger Performance API that is not
defined by the ECMAScript standard but is implemented by web browsers and by Node. In
order to use the performance object in Node, you must import it with

```js
const { performance } = require("perf_hook");
```

Allowing high-precision timing on the web may allow unscrupulous websites
to fingerprint visitors, so browser may reduce the precision of performance.now() by default. As a web-developer, you should be able to re-enable high-precision timing
somehow (such as by setting `privacy.reduceTimerPrecision` to false in Firefox)

#### 11.4.2 Date Arithmetic

Date objects can be compared with JavaScript's standard `<`, `<=`, `>`, and `>=` comparison
operators. (It works because of Date class defines a `valueOf()` method that
returns a timestamp)

If you want to add or subtract a specified number of **second, minutes, or hours**
from a Date, it is often easiest to simply modify the timestamp. This technique
becomes more cumbersome if we want to add days, months, and years since they have
varying numbers of days. To do date arithmetic involving days, months, and years,
use `setDate()`, `setMonth()` and `setYear()` instead.

```js
let d = new Date();
d.setMonth(d.getMonth() + 3, d.getDate() + 14);
```

**Date setting methods work correctly even when they overflow.**

#### 11.4.3 Formatting and Parsing Date Strings

| Method               | Example                                           |
| -------------------- | ------------------------------------------------- |
| console.log(date)    | 2099-12-31T18:03:04.005Z                          |
| valueOf()            | 4102423384005                                     |
| toUTCString()        | Thu, 31 Dec 2099 18:03:04 GMT                     |
| toISOString()        | 2099-12-31T18:03:04.005Z (**ISO-8601**)           |
| toString()           | Fri Jan 01 2100 02:03:04 GMT+0800 (Malaysia Time) |
| toDateString()       | Fri Jan 01 2100                                   |
| toTimeString()       | 02:03:04 GMT+0800 (Malaysia Time)                 |
| toLocaleString()     | 1/1/2100, 2:03:04 AM                              |
| toLocaleDateString() | 1/1/2100                                          |
| toLocaleTimeString() | 2:03:04 AM                                        |

None of these date-to-string methods is ideal when formatting date and times
to be displayed to end users. See [here for more general-purpose and local-aware date-and time- formatting technique](#11.7.2)

Date.parse() method takes a string as its argument, attempts to parse it as a
date and time, and returns a **timestamp** representing that date. It is
able to parse the output of `toString()`, `toUTCString()` and `toISOString()` method.

### 11.5 Error Classes

There is no exception type that must be used to signal errors. However,
one good reason to use an Error object is that, **when you create an Error,
it captures the state of the JavaScript stack, and if the exception is uncaught,
the stack trace will be displayed with the error message, which help you in debugging
the issue**.

Error objects have two properties: `message` and `name`, and a `toString()` method.
`message` is the value you passed to the `Error()` constructor. For "Error"
objects created, then the name is "Error". toString() method return `name: message`.

Node and all modern browser also define a `stack` property on Error objects.
This can be useful information to log when an unexpected error is caught.

```js
err.stack;
"Error: test\n" +
  "    at REPL83:1:11\n" +
  "    at Script.runInThisContext (node:vm:129:12)\n" +
  "    at REPLServer.defaultEval (node:repl:572:29)\n" +
  "    at bound (node:domain:433:15)\n" +
  "    at REPLServer.runBound [as eval] (node:domain:444:12)\n" +
  "    at REPLServer.onLine (node:repl:902:10)\n" +
  "    at REPLServer.emit (node:events:525:35)\n" +
  "    at REPLServer.emit (node:domain:489:12)\n" +
  "    at [_onLine] [as _onLine] (node:internal/readline/interface:422:12)\n" +
  "    at [_line] [as _line] (node:internal/readline/interface:893:18)";
```

**You should feel free to define your own Error subclasses that best
encapsulate the error conditions of your own program.** Note that
you are not limited to the `name` and `message` properties: for example,
you might want to define an HTTPError class that has a `status` property
that holds the HTTP status code of the failed request

```js
class HTTPError extends Error {
  constructor(status, statusText, url) {
    super(`${status} ${statusText}: ${url}`);
    this.status = status;
    this.statusText = statusText;
    this.url = url;
  }

  get name() {
    return "HTTPError";
  }
}

let error = new HTTPError(404, "Not Found", "http://example.com");
error.status; // => 404
error.message; // => "404 Not Found: http://example.com"
error.name; // => "HTTPError"
```

<a id="11.6"></a>

### 11.6 JSON Serialization and Parsing

When a program needs to save data or needs to transmit data across a
network connection to another program, it must to convert its in
memory data structures into a string of bytes or characters than
can be saved or transmitted and then later be parsed to restore the original
in-memory data structures. This process of converting data structures into streams
of bytes or characters is known as serialization (or marshaling or even pickling)

The easiest way to serialize data in JavaScript uses a serialization format
known as JSON. JSON supports primitive numbers and strings, booleans, and null,
as well as arrays and object built up from those primitive values.
JSON does not support other JavaScript types like Map, Set, RegExp, Date
or typed arrays.

```js
JSON.stringify();
JSON.parse();

let o = { s: "", n: 0, a: [true, false, null] };
let s = JSON.stringify(o);
let copy = JSON.parse(s);
```

**We can use this pair of functions as a somewhat inefficient way of creating
a deep copy of an object**:

```js
function deepcopy(o) {
  return JSON.parse(JSON.stringify(o));
}
```

Both `JSON.stringify()` and `JSON.parse()` accept an optional second argument
that allows us to extend the **JSON format**.

If you would like your JSON formatted string to be human-readable, then
you should pass `null` as the second argument (it is replacer,
see [here for more detail](#json_replacer)), and pass a number or string as the
third argument (if it is number, it is number of space, if it is a string of whitespace '\t', then it will use that string for each level of indent):

```js
let o = { s: "test", n: 0, m:{a:1}};
JSON.stringify(o, null, 2); // => '{\n  "s": "test",\n  "n": 0,\n  "m": {\n    "a": 1\n  }\n}'
{
  "s": "test",
  "n": 0,
  "m": {
    "a": 1
  }
}
```

JSON.parse() ignores whitespace, so passing a third argument to JSON.stringify() has
no impact on our ability to convert the string back into a data structure.

If `JSON.stringify()` is asked to serialize a value that is not natively supported
but the JSON format, it **looks to see if that value has a `toJSON()` method**,
and if so, it calls that method and then stringifies the return value in place
of the original value. **Date objects implement `toJSON()` by returning `toISOString()` method**. **When you parse the serialized string, the re-created data structure
will not be exactly the same as the one you started with because it will
have string instead of a Date object**

```js
let aa = {
  o: 1,
  toJSON: function () {
    return "hoho";
  },
};
JSON.stringify(aa); // => '"hoho"'

let d = JSON.stringify({ d: new Date() });
JSON.parse(d); // => { d: '2023-12-31T08:10:15.032Z' }
```

If you need to re-create Date objects (or modify the parsed object in any other way),
you can pass a "reviver" function as the second argument to JSON.parse():
The first argument is a property name (or an array index), second argument is
the primitive value of that object property or array element. If the reviver
returns `undefined`, then the named property **will be deleted from object
or array before JSON.parse() returns to the user**

**Tip**: function is invoked as a method of the object or array that
contains the primitive value, so you can refer to that containing object
with the `this` keyword

```js
let text = { d: new Date(), hello: "hello" };
let date = JSON.parse(JSON.stringify(text), function (key, value) {
  console.log(this.hello); // => hello
  if (key[0] === "_") return undefined;

  if (
    typeof value === "string" &&
    /^\d\d\d\d-\d\d-\d\dT\d\d:\d\d:\d\d.\d\d\dZ$/.test(value)
  ) {
    return new Date(value);
  }

  return value;
});
```

In addition to its use of `toJSON()` described earlier, `JSON.stringify()` also
allows its output to be customized by passing an array or a function as the
optional argument -- array of strings. Any property whose name is not in
the array **will be omitted from stringification**. Furthermore, the returned
string will include properties in the **same order** they appear in the array.

```js
let o1 = {
  a: "123",
  b: "asdfasdf",
};

JSON.stringify(o1, ["b", "a", "c"]); // => {"b":"asdfasdf","a":"123"}
```

<a id="json_replacer"></a>

If the second argument is a function, it is a replacer function. If
specified, the replacer function is invoked for each value to be stringified.
The first argument to the replacer function is the object property name,
second argument is value itself. **If the replacer returns `undefined` or
return nothing, then that value is omitted from the stringification**.

```js
let o1 = {
  a : "123",
  b : "asdfasdf",
  c : /asdfasdf/,
}
JSON.stringify(o1, (k, v) => v instanceof RegExp ? undefined : v, 2)

{
  "a": "123",
  "b": "asdfasdf"
}
```

In general, though, if you define a toJSON() method for a type, of if
you use a replacer function, then you will typically need to use a custom reviver function
with JSON.parse() to get your original structure back. **You should understand that
you are defining a custom data format and sacrificing portability and compatibility
with a large ecosystem of JSON-compatible tools and languages.**

### 11.7 The Internationalization API

`Intl.NumberFormat`, `Intl.DateTimeFormat`, `Intl.Collator`.

These classes are not part of the ECMAScript standard but are
defined as part of the ECMA402 standard and are well-supported
by web browsers.

#### 11.7.1 Formatting Numbers

`Intl.NumberFormat` class takes two arguments: the first argument
specifies the **locale** that the number should be formatted for and the second
is an object that specifies more **details** about how the number should be formatted.
If the first argument is omitted or `undefined`, then the system locale (which
we assume to be the user's preferred locale) will be used. If the first argument is
a string, it specifies a desired locale, e.g. `en-US` (English as used in US), `fr`,
or `zh-Hans-CN` (Simplified Chinese). The first argument can also be
array of locale string --- choose the specific one that is well-supported.

The second argument if optional -- object that defines one or more of the following:

| Field                                                  | Description                                                                                                                                                                                                                  |
| ------------------------------------------------------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `style`                                                | `decimal` (Default), `percent`, `currency`                                                                                                                                                                                   |
| `currency`                                             | If style is `currency`, then this property is **required** to specify ISO3 currency code e.g. `USD`, `MYR`, `GBP`                                                                                                            |
| `currencyDisplay`                                      | If style is `currency`, then this field specifies how the currency is displayed: `symbol` (Default), `code` (ISO-3) , `name` (Spells out the name)                                                                           |
| `useGrouping`                                          | Set this property to `false` if you do not want numbers to have thousand separators                                                                                                                                          |
| `minimumIntegerDigits`                                 | The minimum number of digits to use to display integer part of the number. If the number has fewer digits than this, it will be padded on the left with zeros. The default value is 1, but you can use values as high as 21. |
| `minimumFractionDigits`, `maximumFractionDigits`       | Default min is `0` and max is `3` (round if longer), except when formatting monetary amounts, when the length of the fractional part varies depending on the specified currency.                                             |
| `minimumSignificantDigits`, `maximumSignificantDigits` | Making number suitable when formatting scientific data.                                                                                                                                                                      |

`format()`

```js
let euros = Intl.NumberFormat("es", { style: "currency", currency: "EUR" });
euros.format(10); // => "10,00 €"

let pounds = Intl.NumberFormat("en", { style: "currency", currency: "GBP" });
pounds.format(1000); // => "£1,000.00"
```

**Tip**: `format()` method is bound to the NumberFormat object which it belongs to

```js
let data = [0.05, 0.75, 1];
let formatData = Intl.NumberFormat(undefined, {
  style: "percent",
  minimumFractionDigits: 1,
  maximumFractionDigits: 1,
}).format;

data.map(formatData); // => [ '5.0%', '75.0%', '100.0%' ]
```

Some languages, such as Hindi, use a script that its own set of digits, but
tend to use the ASCII digits 0-9 by default. If you want to override
the default script, add `-u-nu-` to the locale and follow it with
an abbreviated script name:

```js
let hindi = Intl.NumberFormat("hi-IN-u-nu-deva").format;
hindi(1234567890); // => "१,२३,४५,६७,८९०"
```

`-u-`: what comes next is a Unicode extension. `nu` is the extension name for the
numbering system, and `deva` is short for Devanagari.

<a id="11.7.2"></a>

#### 11.7.2 Formatting Dates and Times

The `Intl.DateTimeFormat()` constructor takes the same two arguments that `Intl.NumberFormat()` does.

The `Intl.DateTimeFormat` class provides fine-grained control over what is output based
on the properties in the options object that is passed as the second argument to
the constructor.

| Field                | Description                                                                          |
| -------------------- | ------------------------------------------------------------------------------------ |
| year                 | `numeric` for a 4 digit year, `2-digit` for a two digit abbr                         |
| month                | `numeric` for "1"                                                                    |
|                      | `2-digit` for "01"                                                                   |
|                      | `long` for "January"                                                                 |
|                      | `short` for "Jan"                                                                    |
|                      | `narrow` for "J"                                                                     |
| day                  | `numeric` for "1"                                                                    |
|                      | `2-digit` for a "01"                                                                 |
| weekday              | `long` for "Monday"                                                                  |
|                      | `short` for "Mon"                                                                    |
|                      | `narrow` for "M"                                                                     |
| era                  | e.g. CE or BCE. Available: `long`, `short`, `narrow`                                 |
| hour, minute, second | `numeric` for "1"                                                                    |
|                      | `2-digit` for "01"                                                                   |
| timeZone             | If omitted, the local time zone is used                                              |
|                      | Implementations always recognize "UTC" and may also recognize                        |
|                      | Internet Assigned Numbers Authority (IANA) timezone name: e.g. "America/Los_Angeles" |
| timeZoneName         | `long` for a fully spelled-out timezone name                                         |
|                      | `short` for an abbreviated or numeric time zone.                                     |
| hour12               | Boolean property specifies whether or not to use 12-hour time                        |
|                      | The default is locale dependent.                                                     |
| hourCycle            | Specify whether midnight is written as 0 hours, 12 hours or 24 hours                 |
|                      | The default is locale dependent                                                      |
|                      | `h11` specify that midnight is 0 and the hour before midnight is 11pm                |
|                      | `h12` specify that midnight is 12                                                    |
|                      | `h23` specify that midnight is 0 and the hour before midnight is 23.                 |
|                      | `h24` specify that midnight is 24                                                    |

```js
let d = new Date("2020-01-02T13:14:15Z");

Intl.DateTimeFormat("en-US").format(d); // => '1/2/2020'
Intl.DateTimeFormat("fr-FR").format(d); // => '02/01/2020'

let opts = { weekday: "long", month: "long", year: "numeric", day: "numeric" };
Intl.DateTimeFormat("en-US", opts).format(d); // => Thursday, January 2, 2020
Intl.DateTimeFormat("es-ES", opts).format(d); // => 'jueves, 2 de enero de 2020'

opts = { hour: "numeric", minute: "2-digit", timeZone: "America/New_York" };
Intl.DateTimeFormat("fr-CA", opts).format(d); // => // '8 h 14'
```

`Intl.DateTimeFormat` can display dates using calendars other than the default
Julian calendar based on the Christian era. To specify non-Christian calendar,
add `-u-ca-` to the locale and follow with the name of the calender:
`buddhist`, `chinese`, `coptic`, `ethiopic`, `gregory`, `hebrew`, `indian`,
`islamic`, `iso8601`, `japanese`, `persian`.

```js
opts = { year: "numeric", era: "short" };
Intl.DateTimeFormat("en", opts).format(d); // 2024 AD
Intl.DateTimeFormat("en-u-ca-iso8601", opts).format(d); // 2024 AD
Intl.DateTimeFormat("en-u-ca-hebrew", opts).format(d); // 5784 AM
Intl.DateTimeFormat("en-u-ca-buddhist", opts).format(d); // 2567 BE
Intl.DateTimeFormat("en-u-ca-islamic", opts).format(d); // 1445 AH
Intl.DateTimeFormat("en-u-ca-persian", opts).format(d); // 1403 AP
Intl.DateTimeFormat("en-u-ca-indian", opts).format(d); // 1946 Saka
Intl.DateTimeFormat("en-u-ca-chinese", opts).format(d); // 2024(jia-chen)
Intl.DateTimeFormat("en-u-ca-japanese", opts).format(d); // 6 Reiwa
```

<a id="11.7.3"></a>

#### 11.7.3 Comparing Strings

Sorting are not simple in other languages other than English: Spanish, for
example treats ñ as a distinct letter that comes after n and before o.
Lithuanian alphabetizes Y before J, and Welsh treats digraphs like CH
and DD as single letters with CH coming after C and DD sorting after
D.

If you want to display strings to a use in an order that they will find natural,
it is not enough use the `sort()` method on an array of strings.
But if you create an `Intl.Collator` object, you can pass the `compare()`
method of that object to the `sort()` method to perform locale appropriate sorting
of the strings.

The arguments are same as above two utilities: The first
specifies a locale or an array of locales, and the second is an optional
object whose properties specify exactly what kind of string comparison
is to be done. The supported properties are these:

| Property          | Description                                                                                                     |
| ----------------- | --------------------------------------------------------------------------------------------------------------- |
| usage             | Specifies how the collator object is to be used                                                                 |
|                   | `sort` (default)                                                                                                |
|                   | `search` -- less strict comparison that ignore accents                                                          |
| sensitivity       | Specifies whether the collator is sensitive to lettercase and accents when comparing strings                    |
|                   | If usage is `search`, the default sensitivity depends on the locale                                             |
|                   | `base` ignore case and accents, considering only the base letter for each character                             |
|                   | `accent` considers accents in comparisons but ignores case                                                      |
|                   | `case` considers case in comparisons but ignores accent                                                         |
|                   | `variant` (default when usage is `sort`) considers both accents and case                                        |
| ignorePunctuation | Boolean value, specifies whether to ignore spaces and puntuation. e.g. "anyone" === "any one" if it is set true |
| numeric           | Boolean value, set it to `true` if the strings you are comparing are integers or contain                        |
|                   | integers and you want them to be sorted into numerical order instead of alphabetical order.                     |
| caseFirst         | `upper` -- "A" will sort before "a"                                                                             |
|                   | `lower` -- "a" will sort before "A"                                                                             |
|                   | Note, it is `AaBbCcDdEeFf` instead of Unicode lexicographic ordering                                            |
|                   | The default for this property is locale dependent                                                               |

```js
const collator = new Intl.Collator().compare;
["a", "z", "A", "Z"].sort(collator); // => ["a", "A", "z", "Z"]

const filenameOrder = new Intl.Collator(undefined, { numeric: true }).compare;
["page10", "page9"].sort(filenameOrder); // => ["page9", "page10"]

const fuzzMatcher = new Intl.Collator(undefined, {
  sensitivity: "base",
  ignorePunctuation: true,
}).compare;
let strings = ["food", "fool", "Foo Bar"];
strings.findIndex((s) => fuzzMatcher(s, "foobar") === 0); // => 2
```

Some locales have more than one possible collation order. They can be
selected by adding `-u-co-` too the locale string and adding the name of the
desired variant.

```bash
const modernSpanish = Intl.Collator("es-ES").compare;
const traditionalSpanish = Intl.Collator("es-ES-u-co-trad").compare;
let palabras = ["luz", "llama", "como", "chico"];
palabras.sort(modernSpanish) // => [ 'chico', 'como', 'llama', 'luz' ]
palabras.sort(traditionalSpanish) // => [ 'como', 'chico', 'luz', 'llama' ]
```

### 11.8 The Console API

The Console API defines a number of useful functions in addition to console.log().
The API is no part of any ECMAScript standard, but it is supported by browsers
and by Node and has been formally written and standardized at https://console.spec.whatwg.org

| Function                                                         | Description                                                                                                                                                                                                                                    |
| ---------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| console.log()                                                    | Converts its arguments to strings and outputs them to the console.                                                                                                                                                                             |
|                                                                  | It includes spaces between the arguments and starts a new line after outputting all arguments                                                                                                                                                  |
| console.debug(), console.info(), console.warn(), console.error() | In Node, console.error() sends its output to the stderr stream rather than stdout stream, the other functions                                                                                                                                  |
|                                                                  | are aliases of console.log(). In browsers, the outputs may be prefixed by an icon that indicates its                                                                                                                                           |
|                                                                  | level or severity                                                                                                                                                                                                                              |
| console.assert()                                                 | If the first argument is truthy, then this function does nothing; else the remaining arguments are printed as if they had been passed to console.error() with                                                                                  |
|                                                                  | an "Assertion failed" prefix                                                                                                                                                                                                                   |
| console.clear()                                                  | Clears the console when that is possible. If Node's output has been                                                                                                                                                                            |
|                                                                  | redirected to a file or a pipe, however, then calling this function has no effect                                                                                                                                                              |
| console.table()                                                  |                                                                                                                                                                                                                                                |
| console.trace()                                                  | Like `console.log()`, in addition, follows its output with a stack trace.                                                                                                                                                                      |
|                                                                  | In Node, the output goes to stderr instead of stdout                                                                                                                                                                                           |
| console.count()                                                  | Takes a string argument and logs that string followed by the number of times it has been called with that string                                                                                                                               |
|                                                                  | e.g. you can keep track of how many times the event handler has been triggered                                                                                                                                                                 |
| console.countReset()                                             | Takes a string argument and resets the counter for that string                                                                                                                                                                                 |
| console.group()                                                  |                                                                                                                                                                                                                                                |
| console.groupCollapsed()                                         | This function works like console.group() except that in web browsers, the group will be “collapsed” by default and the messages it contains will be hidden unless the user clicks to expand the group. In Node, this function is a synonym for  console.group().
| console.groupEnd()                                               | Ends the indentation and grouping caused by the most recent call to `console.group()` or `console.groupCollapsed()`                                                                                                                            |
| console.time()                                                   | Takes a single argument, make a note of the time it was called with that string, and produces no output                                                                                                                                        |
| console.timeLog()                                                | Takes a single argument, prints that string followed by the elapsed time                                                                                                                                                                       |
|                                                                  | since the `console.time()` call. Additional arguments as if they had been passed                                                                                                                                                               |
|                                                                  | to `console.log()`                                                                                                                                                                                                                             |
| console.timeEnd()                                                | Takes a single argument.prints that string followed by the elapsed time                                                                                                                                                                        |
|                                                                  | since the `console.time()` call. After this function, it is no longer legal to call                                                                                                                                                            |
|                                                                  | `console.timeLog()` without first calling `console.time()`                                                                                                                                                                                     |

#### 11.8.1 Formatted Output With Console

Little-known feature: if the first arguments is a string that includes %s, %i, %d, %f,
%o, %O, or %c, then this first argument is treated as format string.

| Format    | Description                                                                    |
| --------- | ------------------------------------------------------------------------------ |
| %s        | string                                                                         |
| %i and %d | number and then truncated to an integer                                        |
| %f        | number                                                                         |
| %o and %O | The property names and values are displayed.                                   |
|           | The uppercase variant uses an implementation-dependent output format that is   |
|           | judged to be most useful for software developers                               |
| %c        | A string of CSS styles and used to style any text that follows (until the next |
|           | %c sequence or end of the string). In Node, %c are simply ignored              |

**Note**: If you pass an Error object to `console.log()`, it is automatically printed along with its stack trace.

```js
console.log("%o", {hello:{hello: "bye"}})
{ hello: { hello: 'bye' } }
```

#### Example of `console.table()`

```js
console.table([{a:1,b:3},{a:1,b:4}])
┌─────────┬───┬───┐
│ (index) │ a │ b │
├─────────┼───┼───┤
│    0    │ 1 │ 3 │
│    1    │ 1 │ 4 │
└─────────┴───┴───┘
```

#### Example of `console.trace()`

```js
console.trace("asdas")
Trace: asdas
    at REPL53:1:9
    at Script.runInThisContext (node:vm:129:12)
    at REPLServer.defaultEval (node:repl:572:29)
    at bound (node:domain:433:15)
    at REPLServer.runBound [as eval] (node:domain:444:12)
    at REPLServer.onLine (node:repl:902:10)
    at REPLServer.emit (node:events:525:35)
    at REPLServer.emit (node:domain:489:12)
    at [_onLine] [as _onLine] (node:internal/readline/interface:422:12)
    at [_line] [as _line] (node:internal/readline/interface:893:18)
```

#### Example of `console.count` and `console.countReset()`

```js
> console.count("hello")
hello: 1
> console.count("hello")
hello: 2
> console.count("hello1")
hello1: 1
> console.countReset("hello")
> console.count("hello")
hello: 1
```

#### Example of `console.group() / groupCollapsed()` and `console.groupEnd()`

```js
> console.group("hello")
> console.group("hello")
> console.group("hello")
> console.group("hello")
> console.groupEnd()
> console.groupEnd()
> console.group("hello")

=output=
hello
  hello
    hello
      hello
    hello
```

#### Example of `console.time()`, `console.timeLog()` and `console.timeEnd()`

```js
> console.time("hello")
> console.timeLog("hello")
hello: 9.241s
> console.timeLog("hello1")
>       (node:3578) Warning: No such label 'hello1' for console.timeLog()
      (Use `node --trace-warnings ...` to show where the warning was created)
> console.timeEnd("hello")
hello: 25.979s
> console.timeLog("hello1")
>       (node:3578) Warning: No such label 'hello1' for console.timeLog()
> console.timeLog("hello")
>       (node:3578) Warning: No such label 'hello' for console.timeLog()
```

<a id="11.9"></a>

### 11.9 URL APIs

Since JavaScript is so commonly used in web browsers and web servers,
it is common for JavaScript code to need to manipulate URLs.

The URL class is not part of any ECMAScript standard, but it works in
Node and all internet browsers other than Internet Explorer.

Create a URL object with `URL()` constructor:

- Passing an absolute URL string as the argument
- Passing a relative URL as the first argument and the absolute URL that it is
  relative to as a second argument

```js
let url = new URL("https://example.com:8000/path/name?q=term#fragment");
url.href; // => "https://example.com:8000/path/name?q=term#fragment"
url.origin; // => "https://example.com:8000"
url.protocol; // => "https:"
url.host; // => "example.com:8000"
url.hostname; // => "example.com"
url.port; // => "8000"
url.pathname; // => "/path/name"
url.search; // => "?q=term"
url.hash; // => "#fragment"
```

Although it is not commonly used, URLs can include a username or a username and
password:

```js
let url = new URL("ftp://admin:1337!@ftp.example.com/");
url.href; // => "ftp://admin:1337!@ftp.example.com/"
url.origin; // => "ftp://ftp.example.com"
url.username; // => "admin"
url.password; // => "1337!"
```

`origin` is read only property, it is just a simple combination of protocol and host.

```js
let url = new URL("https://example.com");
url.pathname = "api/search";
url.search = "q=test";
url.href; // => "https://example.com/api/search?q=test"
url.toString(); // => "https://example.com/api/search?q=test"
url.host = "example2.com";
url.port = "8888";
url.origin; // => "https://example2.com:8888"
```

**One of the important features of the URL class is that it correctly adds
punctuation and escapes special characters in URLs when that is needed**:

```js
let url = new URL("https://example.com");
url.pathname = "path with spaces";
url.search = "q=foo#bar";
url.pathname; // => "/path%20with%20spaces"
url.search; // => '?q=foo%23bar'
url.href; // => 'https://example.com/path%20with%20spaces?q=foo%23bar'
```

**Tip**: Setting `href` to a new string reruns the URL parser.

`searchParams` will be more useful than the `search` property. It is a
read-only reference to a URLSearchParam object, which has an API for getting,
setting, adding, deleting, and sorting the parameters encoded into the query
portion of the URL.

```js
let url = new URL("https://example.com/search");
url.search; // => ""
url.searchParams.append("q", "term");
url.search; // => "?q=term"
url.searchParams.set("q", "x");
url.search; // => "?q=x"
url.searchParams.get("q"); // => "x"
url.searchParams.has("q"); // => true
url.searchParams.has("p"); // => false
url.searchParams.append("opts", "1");
url.searchParams.append("opts", "&");
url.search; // => '?q=x&opts=1&opts=%26'
url.searchParams.get("opts"); // => "1" : The first value
url.searchParams.getAll("opts"); // => ["1", "&"]: all values
url.searchParams.sort(); // => Put params in alphabetical order
url.search; // => "?opts=1&opts=%26&q=x"
url.searchParams.set("opts", "y") // change the opts params
[...url.searchParams]; // It is iterable! => [["opts", "y"], ["q", "x"]]
url.searchParams.delete("opts")
url.search // => '?q=x'
url.href // => 'https://example.com/search?q=x'
```

If you wan to encode URL parameters into a query string, you can
create a `URLSeacrhParams` object, append parameters, then convert it to a string
and set it on the `search` property of a URL:

```js
let url = new URL("http://example.com");
let params = new URLSearchParams();
params.append("q", "term");
params.append("opts", "exact");
params.toString(); // => "q=term&opts=exact"
url.search = params;
url.href; // => 'http://example.com/?q=term&opts=exact'
```

#### 11.9.1 Legacy URL Functions:

`escape()` and `unescape()` functions are now deprecated (but still widely implemented).
They should not be used.

When there are deprecated, ECMAScript introduced two pairs of alternative global functions. The fundamental problem with all of these legacy functions is that **they seek
to apply a single encoding scheme to all parts of a URL** when the
fact is that different portions of a URL use different encodings. The solution
is simply to use the URL class for all URL manipulation.

> `encodeURI` and `decodeURI()`

`encodeURI` takes a string as its argument and returns a new
string in which non-ASCII characters plus certain ASCII characters (such as space)
are escaped; while `decodeURI()` reverses the process

Because `encodeURI()` is intended for encoding entire URLs, it cannot work
correctly for URLs that has "/", "?" and "#" (and etc.) within their various components.

```js
let href = "https://example.com/?q=hello world&q=&";
encodeURI(href); // => 'https://example.com/?q=hello%20world&q=&'
```

> `encodeURIComponent()` and `decodeURIComponent()`

They work just like `encodeURI()` and `decodeURI()` except that they are intended
to **escape individual components of a URI**. **But be aware that `encodeURIComponent()`
will escape '/' character in a path name that you probably do not want escaped**.
And it will convert spaces in a query parameter to `%20`, even though spaces are
supposed to be escaped with a `+` in that portion of a URL.

### 11.10 Timers

`setTimeout()` and `setInterval()`

```js
setTimeout(() => {
  console.log("Ready...");
}, 1000);
setTimeout(() => {
  console.log("set...");
}, 2000);
setTimeout(() => {
  console.log("go...");
}, 3000);
```

Bot `setTimeout()` and `setInterval()` return a value. If you save
this value in a variable, **you can then use it later to cancel the
execution of the function by passing it to `clearTimeout()` or `clearInterval()`**.

```js
let t = setTimeout(() => {
  console.log("Ready...");
}, 1000);
clearTimeout(t);
```

```js
let clock = setInterval(() => {
  console.clear();
  console.log(new Date().toLocaleTimeString());
}, 1000);
setTimeout(() => clearInterval(clock), 10000);
```

<a id="12"></a>

## 12 - Iterators and Generators

Iterators means

- the contents of these data structures can be iterated: `for/of` loop.

- be used with the `...` operator to expand or "spread" and iterable
  object into an array initializer or function invocation.

- be used with destructing assignment :
  - ```js
    let purpleHaze = Uint8Array.of(255, 0, 255, 128);
    let [r, g, b, a] = purpleHaze;
    ```

A number of built-in functions and constructors hat are commonly used with
Array objects are actually written to accept arbitrary iterators instead:

```js
// Strings are iterable, so the two sets are the same
new Set("abc"); // => new Set(["a", "b", "c"])
```

### 12.1 How Iterators Work

There are three separate types that you need to understand to understand
iteration in JavaScript:

- `iterable object`: types like Array, Set, and Map

- `iterator object`: performs the iteration

- `iteration result object`: holds the result of each step of iteration.

**An iterable object is any object with a special iterator method that returns an
iterator object.** An _iterator_ is any object with a `next()` method that returns
an iteration result object. And an _iteration result_ object is an object
with properties named `value` and `done`.

The iterator method of an iterable object does not have a conventional name
but uses the Symbol `Symbol.iterator` as its name.

```js
let iterable = [99];
let iterator = iterable[Symbol.iterator]();
for (let result = iterator.next(); !result.done; result = iterator.next()) {
  console.log(result.value); // => 99
}
```

This is occasionally useful in code like the following when you want
to iterate though a "partially used" iterator:

```js
let list = [1, 2, 3, 4, 5];
let iter = list[Symbol.iterator]();
let head = iter.next().value; // => 1
let tail = [...iter]; // => [2,3,4,5]
```

<a id="12"></a>

### 12.2 Implementing Iterable Objects

Let's try making a class without relying on a generator:

```js
class Range {
  constructor(from, to) {
    this.from = from;
    this.to = to;
  }

  has(x) {
    return typeof x === "number" && this.from <= x && x <= this.to;
  }

  toString() {
    return `{ x | ${this.from} ≤ x ≤ ${this.to} }`;
  }

  [Symbol.iterator]() {
    let next = Math.ceil(this.from);
    let last = this.to;
    return {
      next() {
        return next <= last ? { value: next++ } : { done: true };
      },
      // As a convenience, we make the iterator itself iterable
      [Symbol.iterator]() {return this;}
    };
  }
}

for (let x of new Range(1,10)) console.log(x) // => Logs numbers 1 to 10
[...new Range(-2,2)] // => [ -2, -1, 0, 1, 2 ]
```

In addition to making your classes iterable, it can be quite useful to **define
functions that return iterable values**. Consider these iterable-based alternatives
to the `map()` and `filter()` methods of JavaScript arrays:

```js
function map(iterable, f) {
  let iterator = iterable[Symbol.iterator]();
  return {
    [Symbol.iterator]() {
      return this;
    },
    next() {
      let v = iterator.next();
      if (v.done) {
        return v;
      } else {
        return { value: f(v.value) };
      }
    },
  };
}

[...map(new Range(1, 4), (x) => x * x)]; // => [1,4,9,16]

function filter(iterable, predicate) {
  let iterator = iterable[Symbol.iterator]();
  return {
    [Symbol.iterator]() {
      return this;
    },
    next() {
      for (;;) {
        let v = iterator.next();
        if (v.done || predicate(v.value)) {
          return v;
        }
      }
    },
  };
}

[...filter(new Range(1, 10), (x) => x % 2 === 0)]; // => [2,4,6,8,10]
```

One key feature of iterable objects and iterators is that they are inherently lazy;
when computation is required to compute the next value, the computation can
be deferred until the value is actually needed. For example, you could
use `split()` method to tokenize a very long string of text into space-separated words.
But if you do this, then the entire string has to be processed before you can use
even the first word, and you end up allocating lots of memory for the returned
array and all of the strings within it. Instead, we can lazily iterate the
words of a string without keeping them all in memory at once:

```js
function words(s) {
  let r = /\s+|$/g; // Match one or more spaces or end
  r.lastIndex = s.match(/[^ ]/).index; // Start matching at first non space
  let matches = s.matchAll(r);
  let start = r.lastIndex;
  return {
    [Symbol.iterator]() {
      return this;
    },
    next() {
      let match = matches.next();
      if (match.done) {
        return { done: true };
      }
      let v = { value: s.substring(start, match.value.index) };
      start = match.value.index + 1;
      return v;
    },
  };
}
[...words(" abc def ghi!")]; // => ["abc", "def", "ghi!"]
```

<a id="12.2.1"></a>

#### 12.2.1 "Closing" and Iterator: The Return Method

Iterator objects may implement a `return()` method to go along with the
`next()` method. If iteration stops before `next()` has returned an
iteration result with the `done` property set to `true` (most commonly
because you left a `for/of` loop early via a `break` statement), then
the interpreter will check to see if the iterator object has a `return()`
method. If this method exists, the interpreter will invoke it with
no arguments, giving the iterator the chance to do something (e.g. close file),
release memory, and otherwise clean up after itself. The `return()` method
must return an iterator result object. The properties of the object are ignored,
but it is an error to return a non-object value.

```js
class Range {
  constructor(from, to) {
    this.from = from;
    this.to = to;
  }

  has(x) {
    return typeof x === "number" && this.from <= x && x <= this.to;
  }

  toString() {
    return `{ x | ${this.from} ≤ x ≤ ${this.to} }`;
  }

  [Symbol.iterator]() {
    let next = Math.ceil(this.from);
    let last = this.to;
    return {
      next() {
        return next <= last ? { value: next++ } : { done: true };
      },
      // As a convenience, we make the iterator itself iterable
      [Symbol.iterator]() {
        return this;
      },
      return() {
        console.log("bye");
        return {};
      },
    };
  }
}
for (let r of new Range(1, 5)) {
  if (r === 3) {
    break; // => "bye"
  }
}
```

### 12.3 Generators

`ES6`

It is complicated to implement iterable object. Fortunately, generators
can dramatically **simplify the creation of custom iterators**

To create a generator, you must first define a _generator function_ -- `function*`.
When you invoke a generator function, it does not actually execute the function body,
but instead returns a generator object, instead, it returns a generator object.
**The generator object is an iterator**. Calling its `next()` method causes
the body of the generator function to run from the start (or whatever its
current position is) until it reaches a `yield` statement. The value
of `yield` statement becomes the value returned by the `next()` call on the iterator.

```js
function* oneDigitPrimes() {
  yield 2;
  yield 3;
  yield 5;
  yield 7;
}
let primes = oneDigitPrimes();

primes.next().value // => 2
primes.next().value // => 3
primes.next().value // => 5
primes.next().value // => 7
primes.next().done // => true

// Generators have a Symbol.iterator method to make them iterable
primes[Symbol.iterator]() // => primes

[...oneDigitPrimes()] // => [2,3,5,7]
let sum = 0;
for (let prime of oneDigitPrimes()) sum += prime;
sum // => 17
```

```js
const seq = function* (from, to) {
  for (let i = from; i <= to; i++) yield i;
};
[...seq(3, 5)]; // => [ 3, 4, 5 ]
```

In classes and object literals, we can use shorthand notation to omit the `function` keyword:

```js
let o = {
  x:1, y:2, z:3,
  *g() {
    for (let key of Object.keys(this)) {
      yield key
    }
  }
}
[...o.g()] // => [ 'x', 'y', 'z', 'g' ]
```

**Important**: There is no way to write a generator function using arrow function syntax.

Generators often make it particularly easy to define iterable classes:
We can replace the [Symbol.iterator]() to `*[Symbol.iterator]()`

```js
*[Symbol.iterator]() {
  for(let x = Math.ceil(this.from); x <= this.to; x++) yield x;
}
```

And words() function:

```js
function* words(s) {
  let r = /\s+|$/g; // Match one or more spaces or end
  let start = (r.lastIndex = s.match(/[^ ]/).index); // Start matching at first non space
  let matches = s.matchAll(r);
  for (let match = matches.next(); !match.done; match = matches.next()) {
    yield s.substring(start, match.value.index);
    start = match.value.index + 1;
  }
}
[...words(" abc def ghi!")]; // => ["abc", "def", "ghi!"]
```

#### 12.3.1 Generator Examples

Fibonacci numbers:

```js
function* fibonacciSequence() {
  let x = 0,
    y = 1;
  for (;;) {
    yield y;
    [x, y] = [y, x + y];
  }
}
```

Note that above is infinitive loop, it is more useful with a `take()` generator like:

```js
function* take(n, iterable) {
  let it = iterable[Symbol.iterator]();
  while (n-- > 0) {
    let next = it.next();
    if (next.done) return;
    else yield next.value;
  }
}
[...take(5, fibonacciSequence())]; // => [ 1, 1, 2, 3, 5 ]
```

Another generator function that interleaves the element of multiple iterable objects:

```js
function* zip(...iterables) {
  let iterators = iterables.map((i) => i[Symbol.iterator]());
  let index = 0;
  while (iterators.length > 0) {
    if (index >= iterators.length) {
      index = 0;
    }
    let item = iterators[index].next();
    if (item.done) {
      iterators.splice(index, 1); // remove it from the array
    } else {
      yield item.value;
      index++;
    }
  }
}
[...zip(oneDigitPrimes(), "ab", [0])]; // => [ 2, 'a', 0, 3, 'b', 5, 7 ]

function* oneDigitPrimes() {
  yield 2;
  yield 3;
  yield 5;
  yield 7;
}
```

#### 12.3.2 yield\* and Recursive Generators

Yields the elements of multiple iterable objects sequentially rather than
interleaving them:

```js
function* sequence(...iterables) {
  for (let iterable of iterables) {
    for (let item of iterable) {
      yield item;
    }
  }
}
[...sequence("abc", oneDigitPrimes())]; // => [ 'a', 'b', 'c', 2, 3, 5, 7 ]
```

`ES6` has special syntax for it. The `yield*` keyword is like `yield` except that,
rather than yielding a single value, it iterates and iterable object and yields each
of the resulting values:

```js
function* sequence(...iterables) {
  for (let iterable of iterables) {
    yield* iterable;
  }
}
[...sequence("abc", oneDigitPrimes())]; // => [ 'a', 'b', 'c', 2, 3, 5, 7 ]
```

The array `forEach()` method is often an elegant way to loop over the elements
of an array:

```js
function* sequence(...iterables) {
  iterables.forEach(iterable =>
    yield* iterable;
  )
}
[...sequence("abc", oneDigitPrimes())]; // => [ 'a', 'b', 'c', 2, 3, 5, 7 ]
```

**Above does not work**. `yield` and `yield*` can only be used within generator functions, but the
nested arrow function in this code is a regular function, not a `function*` generator
function, so `yield` is not allowed.

### 12.4 Advanced Generator Features

The most common use of generator functions is to create iterators, **but the
fundamental feature of generators is that they allow us to pause a
computation, yield intermediate results, and then resume the computation later**

#### 12.4.1 The Return Value of a Generator Function

If a generator function has `return` statements, they have been used to cause an
early return, not to return value. If in the case of a generator `return`s a value,
the final call to `next` returns an object that has both `value` (the return value) and
`done` (true) defined.

```js
function* oneAndDone() {
  yield 1;
  return "done";
}
[...oneAndDone()]; // => [1]
let generator = oneAndDone();
generator.next(); // => { value: 1, done: false }
generator.next(); // => { value: 'done', done: true }
```

#### 12.4.2 The Value of a yield Expression

`yield` is an expression, and it can have a value by passing argument to `next` method
. Note that the first argument
is discarded:

```js
function* smallNumbers() {
  console.log("next() invoked the first time; argument discarded");
  let y1 = yield 1; // y1 == "b"
  console.log("next() invoked a second time with argument", y1);
  let y2 = yield 2; // y2 == "c"
  console.log("next() invoked a third time with argument", y2);
  let y3 = yield 3; // y3 == "d"
  console.log("next() invoked a fourth time with argument", y3);
  return 4;
}

let g = smallNumbers();
console.log("generator created; no code runs yet");
let n1 = g.next("a"); // n1.value == 1
console.log("generator yielded", n1.value);
let n2 = g.next("b"); // n2.value == 2
console.log("generator yielded", n2.value);
let n3 = g.next("c"); // n3.value == 3
console.log("generator yielded", n3.value);
let n4 = g.next("d"); // n4 == { value: 4, done: true }
console.log("generator returned", n4.value);
```

```js
// result:
generator created; no code runs yet
next() invoked the first time; argument discarded
generator yielded 1
next() invoked a second time with argument b
generator yielded 2
next() invoked a third time with argument c
generator yielded 3
next() invoked a fourth time with argument d
generator returned 4
```

#### 12.4.3 The `return()` and `throw()` Methods of a Generator

You can also alter the flow of control inside the generator by calling its
`return()` and `throw()` methods.

[Recall from earlier in the chapter that](#12.2.1) , if an iterator defines a `return()`
method and iteration stops early, then the interpreter automatically calls the
`return()` method. In the case of generators, you can't define a custom `return()`
method to handle cleanup, but you can structure the generator code to use a
`try/finally` statement that ensures the necessary cleanup is done in the `finally` block
when the generator returns. By forcing the generator to return, the generator's built-in
`return()` method ensures that the cleanup code is run when the generator will no longer
be used.

Calling `throw()` method always causes an exception inside the generator. We can
alter the behavior of the generator with appropriate exception-handling code. Imagine,
for example, a counter generator that yields and ever increasing sequence of integers.
This could be written so that an exception sent with `throw()` would reset the counter
to zero.

<a id="13"></a>

## 13 - Asynchronous JavaScript

Asynchronous programming is commonplace in JavaScript, and this chapter
documents three important language features that help make it easier to work
with asynchronous code.

- Promises (`ES6`), are objects that represent the not-yet-available result of
  an asynchronous operation

- `async` and `await` (`ES2017`) simplifies asynchronous programming by allowing you
  to structure your Promise-based code as if it was synchronous.

- Asynchronous iterators and the `for/await` loop (`ES2018`) allow you to
  work with streams of asynchronous events using simple loops that appear synchronous.

### 13.1 - Asynchronous Programming and Callbacks

At its most fundamental level, asynchronous programming in JavaScript is done
with _callbacks_.

#### 13.1.1 Timers

`setTimeout()` function.

`setInterval()` function

#### 13.1.2 Events

Client-side JavaScript programs are almost universally event driven:
they typically wait for the user to do something and then respond to the
user's actions.

```js
let okay = document.querySelector("#confirmUpdateDialog button.okay");
// register a callback function to be invoked when the user
// clicks on that button.
okay.addEventListener("click", applyUpdate);
```

#### 13.1.3 Network Events

**Client-side** JavaScript code can use the XMLHttpRequest class
plus callback functions to make HTTP requests and asynchronously
handle the server's response when it arrives.

**Tip**: The XMLHttpRequest has nothing to do with XML. In modern
client side JavaScript, it has largely been replaced by the `fetch()` API.

<a id="XMLHttpRequest"></a>

```js
function getCurrentVersionNumber(versionCallback) {
  let request = new XMLHttpRequest();
  request.open("GET", "http://localhost:8080");
  request.send();
  request.onload = function () {
    if (request.status === 200) {
      let currentVersion = parseFloat(request.responseText);
      versionCallback(null, currentVersion);
    } else {
      versionCallback(request.statusText, null);
    }
  };
  request.onerror = request.ontimeout = function (e) {
    versionCallback(e.type, null);
  };
}
getCurrentVersionNumber(function (err, response) {
  console.log(err, response);
});
```

By convention, event listener properties like these always have names that begin
with _on_. `addEventListener()` is more flexible technique **because it allows
for multiple event handlers**.

#### 13.1.4 Callbacks and Events in Node

The Node.js server-side JavaScript environment is deeply asynchronous and defines
many APIs that use callbacks and events.

The default API for reading the contents of a file is asynchronous and invokes a callback
function:

```js
const fs = require("fs");
let options = {
  // default options would go here
};

fs.readFile("config.json", "utf-8", (err, text) => {
  if (err) {
    console.warn("Could not read config file:", err);
  } else {
    Object.assign(options, JSON.parse(text));
  }
  startProgram(options);
});
function startProgram(o) {
  console.log("%o", options);
}
```

Making HTTP request:

```js
const https = require("https");

function getText(url, callback) {
  request = https.get(url);
  request.on("response", (response) => {
    let httpStatus = response.statusCode;
    response.setEncoding("utf-8");
    let body = "";
    response.on("data", (chunk) => {
      body += chunk;
    });
    response.on("end", () => {
      if (httpStatus === 200) {
        callback(null, body);
      } else {
        callback(httpStatus, null);
      }
    });
  });
  request.on("error", (err) => {
    callback(err, null);
  });
}
```

### 13.2 Promises

_Promises_, a core language feature designed to simplify asynchronous programming.

A Promise is an object that represents the result of an asynchronous computation.
**That result may or may not be ready yet**, and the Promise API is intentionally
vague about this: **There is no way to synchronously get the value of Promise;
you can only ask the Promise to call a callback function when the value is ready**.

If you are defining an asynchronous API like the `getText()` function in the previous
section, but want to make it Promise-based, omit the callback argument,
and instead return a Promise object. The caller can then register one or
more callbacks on this Promise object, and they will be invoked when the asynchronous
computation is done.

So, at the simplest level, **Promises are just a different way of working with callbacks**.

One real problem with callback-based asynchronous programming is that it is
common to end up with callbacks inside callbacks inside callbacks. Promises allow
this kind of nested callback to be re-expressed as a more linear Promise chain that
tends to be easier to read and easier to reason about.

Other than this, Promises help error handling by standardizing a way to handle errors
and providing a way for errors to propagate correctly through a chain of promises.

#### 13.2.1 Using Promises

```js
getJSON(url).then((jsonData) => {
  // This is a callback function that will be asynchronously
  // invoked with the parsed JSON value when it becomes available.
});
```

`getJSON()` starts an asynchronous HTTP request for the URL you specify and then,
while the request is pending, it **returns a Promise object**.
The Promise object defines a `then()` instance method. Instead
of passing our callback function directly to getJSON(), we instead
pass it to the `then()` method.

As a simple syntactical lever, the `then()` method is the distinctive feature
of Promises, and it is idiomatic to append `.then()` directly to the
function invocation that returns the Promise, without the
intermediate step of assigning the Promise object to a variable.

It is also idiomatic to name functions that return Promises and
functions that use the results of Promises with **verbs**.

```js
function displayUserProfile(profile) {
  /* implementation omitted */
}

getJSON("/api/user/profile").then(displayUserProfile);
```

##### Handling Errors with Promises

For Promises, we can do this by passing a second function to the `then()` method:

**Important**: The second function does not catch error thrown in first argument.

```js
getJSON("api/user/profile").then(displayUserProfile, handleProfileError);
```

When a synchronous computation completes normally, it simply returns its result to
its caller. When a Promise-based asynchronous computation completes normally, it passes
its result to the function that is first argument to `then()`

When something goes wrong in a synchronous computation, it throws an exception that propagates up the
callstack until there is a `catch` clause to handle it. When an asynchronous computation
runs, its caller is no longer on the stack, so if something goes wrong, **it is simply not
possible to throw an exception back to the caller**.

Instead, Promised-based asynchronous computations pass the exception (typically
as an Error object of some kind, though this is not required) to the
second function passed to `then()`.

In practice, it is rare to see two functions passed to `then()`. Instead,
the more idiomatic way to handle errors:

```js
getJSON("api/user/profile").then(displayUserProfile).catch(handleProfileError);
```

With this code, any error in `getJSON()` or `displayUserProfile()` get passed to
`handleProfileError()`.
The `catch()` method is just a shorthand for calling `then()` with a `null` first
argument and the specified error handler function as the second argument.

---

> **Promise Terminology**

We say that the promise has been fulfilled if and when the first callback is called. And we say that the Promise
has been rejected if and when the second callback is called. If a Promise is neither fulfilled nor
rejected, then it is pending. And once a promise is fulfilled or rejected, we say that it is settled. Note
that a Promise can never be both fulfilled and rejected. Once a Promise settles, it will never change
from fulfilled to rejected or vice versa.

---

#### 13.2.2 Chaining Promises

One of the most important benefits of Promises is that they provide a
natural way to express a sequence of asynchronous operations as a
linear chain of then() method invocations, without having to nest
each operation within the callback of the previous one.

```js
fetch(documentURL)
  .then((response) => response.json())
  .then((document) => {
    return render(document);
  })
  .then((rendered) => {
    cacheInDatabase(rendered);
  })
  .catch((error) => handle(error));
```

[XMLHttpRequest object](#XMLHttpRequest) used to make an HTTP request in JavaScript has
an old and awkward API, and it has largely been replaced by the newer,
Promised-based Fetch API. That promise is fulfilled when the HTTP response
begins to arrive and the HTTP status and headers are available:

```js
fetch("/api/user/profile").then((response) => {
  if (
    response.ok &&
    response.headers.get("Content-Type") === "application/json"
  ) {
    // What can we do here? We don't actually have the response body yet.
  }
});
```

The Response object gives you access to request status and headers,
and it also defines methods like `text()` and `json()`.

**Important**: `json()` method returns a SyntaxError if the response body cannot be
parsed as JSON. However, it doesn't check Content-Type

Although the initial Promise is fulfilled, **the body of the response may not yet have
arrived**. **So these `text()` and `json()` methods for accessing the body
of the response themselves return Promises**.

```js
fetch("/api/user/profile")
  .then((response) => {
    return response.json();
  })
  .then((profile) => {
    displayUserProfile(profile);
  });
```

#### 13.2.3 Resolving Promises

When you pass a callback `c` to the `then()` method, `then()` returns
a Promise `p` and arranges to asynchronously invoke `c` at some later time.
The callback perform some computation and returns a value `v`. When the
callback returns, `p` is resolved with the value `v`.

When a Promise is resolved with a value that is not itself a Promise,
it is immediately fulfilled with that value. **So if `c` returns a non-Promise,
that return value becomes the value of `p`, `p` is fulfilled and we are done**.

If the return value `v` is itself a Promise, then `p` is resolved **but not yet fulfilled**. At this stage, `p` cannot settle until the Promise `v` settles. If `v` is fulfilled,
then `p` will be fulfilled to the same value. If `v` is rejected, then `p` will be
rejected for the same reason.

#### 13.2.4 More on Promises and Errors

With synchronous code, if you leave out error-handling code, you'll
at least get an exception and a stack trace that you can use to figure
out what is going wrong.
**With asynchronous code, unhandled exceptions will often go unreported,
and errors can occur silently, making them much harder to debug**.

##### The Catch and Finally Methods

In `ES2018`, Promise objects also define a `.finally()` method
whose purpose is similar to the `finally` clause in a `try/catch/finally` statement.
If you add a `.finally()` invocation to your Promise chain, then the callback you pass to
`.finally()` will be invoked **when the Promise you called it on settles**.

Let's correct our code with a more realistic version:

```js
fetch("http://localhost:8080/api/user/profile")
  .then((response) => {
    if (!response.ok) {
      return null; // Maybe user is logged out, return null profile
    }
    let type = response.headers.get("content-type");
    if (type !== "application/json") {
      throw new TypeError("Expected JSON, got ${type}");
    }
    return response.json();
  })
  .then((profile) => {
    if (profile) {
      displayUserProfile(profile);
    } else {
      displayLoggedOutProfile();
    }
  })
  .catch((e) => {
    if (e instanceof NetworkError) {
      displayErrorMessage("Check your internet connection.");
    } else if (e instanceof TypeError) {
      displayErrorMessage("Something is wrong with your server!");
    } else {
      console.error(e);
    }
  });
```

Although it is idiomatic to end every Promise chain with a `.catch()` to
clean up any errors that occurred in the chain, it is also perfectly valid
to use `.catch()` elsewhere in a Promise chain. If one of the
stages in your Promise chain can fail with an error, and if the error
is some kind of recoverable error that should not stop the rest of the
chain from running, then you can insert a `.catch()` call in the chain:

```js
startAsyncOperation()
  .then(doStageTwo)
  .catch(recoverFromStageTwoError)
  .then(doStageThree)
  .then(doStageFour)
  .catch(logStageThreeAndFourErrors);
```

Remember that the callback you pass to `.catch()` will only be
invoked if the callback at a previous stage throws an error. If the
callback returns normally, then the `.catch()` callback will be
skipped, and the return value of the previous callback will become the
input to the next `.then()` callback.

Sometimes, in complex network environments, errors can occur more or less
random, and it can be appropriate to handle those errors by simply retrying
the asynchronous request.

```js
queryDatabase()
  .catch((e) => wait(500).then(queryDatabase))
  .then(displayTable)
  .catch(displayDatabaseError);
```

Adding this one line of code should reduce your error rate from 1% to 0.1%

#### 13.2.5 Promises in Parallel

`Promise.all()` execute a number of asynchronous operations in parallel.
It takes an array of Promise objects as its input and returns a Promise.
**The returned Promise will be rejected if any of the input Promises are
rejected**, this happens immediately upon the first rejection and can
happen while other input Promises are still pending.
Otherwise, it will be fulfilled with an array of
the fulfillment values of each of the input Promises.

```js
const urls = [
  /*zero or more URL here */
];
promises = urls.map((url) => fetch(url).then((r) => r.text()));
Promise.all(promises)
  .then((bodies) => {
    /* do something with the array of strings */
  })
  .catch((e) => console.error(e));
```

In `ES2020`, `Promise.allSettled()` takes an array of input Promises
and returns a Promise, just like `Promise.all()` does. But it never
rejects the returned Promise, and it does not fulfill that Promise until
all of the input Promises have settled. The Promise resolves to an array of
objects, with one object for each input Promise. Each of these returned objects
has a `status` property set to "fulfilled" or "rejected". If the status
is "fulfilled", then the object will also have a `value` property that gives
the fulfillment value; else the object will have a `reason` property that
gives the error or rejection value of the corresponding Promise:

```js
Promise.allSettled([Promise.resolve(1), Promise.reject(2), 3]).then(
  (results) => {
    results[0]; // => { status: 'fulfilled', value: 1 }
    results[1]; // => { status: 'rejected', reason: 2 }
    results[2]; // => { status: 'fulfilled', value: 3 }
  },
);
```

Occasionally, you may want to run a number of Promises at once but may only care
about the value of the first one to fulfill **or rejected**. Use `Promise.race()`. It returns
a Promise that is fulfilled or rejected when the first of the Promises in the
input array is fulfilled or rejected.

```js
Promise.race([Promise.resolve(1), Promise.reject(2), 3]).then((result) => {
  console.log(result); // => 1
});
```

#### 13.2.6 Making Promises

##### Promises Based On Other Promises

Returns a Promise if you have some other Promise-returning function to start with.

```js
function getJSON(url) {
  return fetch(url).then((response) => response.json());
}

function getHighScore() {
  return getJSON("/api/user/profile").then((profile) => profile.highScore);
}
```

##### Promises Based On Synchronous Values

`Promise.resolve()` and `Promise.reject()` returns a Promise.

It is unusual to write a Promise-based function where the value is computed
synchronously and returned asynchronously with `Promise.resolve()`. It is fairly common,
however, to have synchronous special cases within an asynchronous function,
and you can handle these special cases with `Promise.resolve()` and `Promise.reject()`.

In particular, if you detect error conditions (such as bad argument values) before
beginning an asynchronous operation, you can report that error by returning a
Promise created with `Promise.reject()`. **You could also just throw an error
synchronously in that case, but that is considered poor form because then
the caller of your function needs to write both a synchronous `catch` clause
and use an asynchronous `.catch()` method to handle errors**.

##### Promises From Scratch

`Promise()` constructor. A function as the only argument:
The function you pass should be written to expect two parameters, which
, by convention, should be named `resolve` and `reject`. The constructor
synchronously calls your function with function arguments for the `resolve` and `reject`
parameters. After calling your function, the `Promise()` constructor returns the
newly created Promise.

That function should perform some asynchronous operation and then **call the `resolve`
function to resolve or fulfill the returned Promise** or **call the `reject` function
to reject the returned Promise**.

```js
function a() {
  return new Promise((resolve, reject)=> {
    resolve(1)
    console.log("hello")
    resolve(1)
    console.log("hello")
    resolve(1)
    console.log("hello")
  })
}

a().then(console.log)

Result:
hello
hello
hello
1
```

Here's how to write the Promise-based `wait()` function:

```js
function wait(duration) {
  return new Promose((resolve, reject) => {
    if (duration < 0) {
      reject(new Error("Time travel not yet implemented"));
    }
    setTimeout(resolve, duration);
  });
}
```

`getJSON()` function:

```js
const http = require("http");

function getJSON(url) {
  return new Promise((resolve, reject) => {
    request = http.get(url, (response) => {
      if (response.statusCode !== 200) {
        reject(new Error(`HTTP status ${response.statusCode}`));
        response.resume(); // so we don't leak memory
      } else if (response.headers["content-type"] !== "application/json") {
        reject(new Error("Invalid content-type"));
        response.resume(); // so we don't leak memory
      } else {
        let body = "";
        response.setEncoding("utf-8");
        response.on("data", (chunk) => {
          body += chunk;
        });
        response.on("end", () => {
          try {
            let parsed = JSON.parse(body);
            resolve(parsed);
          } catch (e) {
            reject(e);
          }
        });
      }
    });
    // We also reject the Promise if the request fails before we
    // even get a response (such as when the network is down)
    request.on("error", (error) => {
      reject(error);
    });
  });
}
```

#### 13.2.7 Promises in Sequence

Promise.all() makes it easy to run an arbitrary number of
Promises in parallel. And Promise chains make it easy to express a
sequence of a fixed number of Promises. Running an arbitrary number
of Promises in sequence is trickier, however. Suppose, for example,
that you have an array of URLs to fetch, but that to avoid overloading
your network, you want to fetch them one at a time.

You need to build one dynamically:

```js
function fetchSequentially(urls) {
  const bodies = [];
  function fetchOne(url) {
    return fetch(url)
      .then((response) => response.text())
      .then((body) => {
        bodies.push(body);
      });
  }

  let p = Promise.resolve(undefined);

  for (url of urls) {
    p = p.then(() => fetchOne(url));
  }

  return p.then(() => bodies);
}

fetchSequentially([
  /* urls */
])
  .then((bodies) => {
    /* do something with the array of strings */
  })
  .catch((e) => console.error(e));
```

There is another (possibly more elegant) approach that we can take.
Rather than creating the Promises in advance, we can have the callback
for each Promise create and return the next Promise.

```js
function promiseSequence(inputs, promiseMaker) {
  inputs = [...inputs];
  function handleNextInput(outputs) {
    if (inputs.length === 0) {
      return outputs;
    } else {
      let nextInput = inputs.shift();
      return promiseMaker(nextInput)
        .then((output) => outputs.concat(output))
        .then(handleNextInput);
    }
  }
  return Promise.resolve([]).then(handleNextInput);
}

function fetchBody(url) {
  return fetch(url).then((r) => r.text());
}

promiseSequence(urls, fetchBody)
  .then((bodies) => {
    /* do something with the array of strings */
  })
  .catch(console.error);
```

### 13.3 `async` and `await`

`ES2017`

These new keywords dramatically simplify the use of Promises and allows us to write
Promise-based, asynchronous code that looks like synchronous code that blocks
while waiting for network responses or other asynchronous events.

#### 13.3.1 `await` Expressions

The `await` keyword takes a Promise and turns it back into a return value or
a thrown exception. Given a Promise object `p`, the expression `await p` waits
until `p` settles. We don't usually use `await` with a variable that holds a Promise;
instead, we use it before the invocation of a function that returns a Promise:

```js
let response = await fetch("/api/user/profile");
let profile = await response.json();
```

#### 13.3.2 `async` Functions

Because any code that use `await` is asynchronous, **there is one critical rule:
you can only use the `await` keyword within functions that have been
declared with the `async` keyword**.

```js
async function getHighScore() {
  let response = await fetch("/api/user/profile");
  let profile = await response.json();
  return profile.highScore;
}

getHighScore().then(displayHighScore).catch(console.error);
```

Declaring a function `async` means **that the return value of the function
will be a Promise even if no Promise-related code appears in the body of the function**.

You can use the `async` keyword with [any kind of function](#8.1)

#### 13.3.3 Awaiting Multiple Promises

```js
let [value1, value2] = await Promise.all([getJSON(url1), getJSON(url2)]);
```

#### 13.3.4 Implementation Details

```js
async function f(x) {
  /* body */
}
```

You can think about this as a Promise-returning function wrapped
around the body of your original function:

```js
function f(x) {
  return new Promise((resolve, reject) => {
    try {
      resolve(
        (function (x) {
          /* body */
        })(x),
      );
    } catch (e) {
      reject(e);
    }
  });
}
```

### 13.4 Asynchronous Iteration

Asynchronous iterators are like the iterators described in Chapter 12,
but they are Promise-based and are meant to be used with a new form of the
`for/of` loop: `for/await`

#### 13.4.1 The `for/await` Loop

`Node 12` makes its readable streams asynchronously iterable.
This means you can read successive chunks of data from a stream:

```js
const fs = require("fs");

async function parseFile(filename) {
  let stream = fs.createReadStream(filename, { encoding: "utf-8" });
  for await (let chunk of stream) {
    parseChunk(chunk);
  }
}
```

The below code is equivalent:

```js
for (const promise of promises) {
  response = await promise;
  handle(response);
}

for await (const promise of promises) {
  handle(response);
}
```

#### 13.4.2 Asynchronous Iterators

Asynchronous iterators are quite similar to regular iterators, but there are
two important differences:

- implements a method with symbolic name `Symbol.asyncIterator` instead of `Symbol.iterator`

  - for/await is compatible with regular iterable
    objects but it prefers asynchronously iterable objects, and tries the
    Symbol.asyncIterator method before it tries the
    Symbol.iterator method.

  - ```js
    async function awaitForIterableObject() {
      for await (const a of [1, 2, 3, 4, 5]) {
        console.log(a);
      }
    }
    ```

- the `next()` method of an asynchronous iterator returns a Promise that resolves
  to an iterator result object instead of returning an iterator result object directly.

#### 13.4.3 Asynchronous Generators

The easiest way to implement an asynchronous iterator is often to use a generator.

```js
function elapsedTime(ms) {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

async function* clock(interval, max = Infinity) {
  for (let count = 1; count <= max; count++) {
    await elapsedTime(interval);
    yield count;
  }
}

async function test() {
  for await (let tick of clock(300, 100)) {
    console.log(tick);
  }
}
```

#### 13.4.4 Implementing Asynchronous Iterators

Instead of using async generators to implement asynchronous iterators,
it is also possible to implement them directly by defining an object
with a `Symbol.asyncIterator` method that returns an object with a
`next()` method that returns a Promise that resolves to an iterator result object.

**Notice that the `next()` method in this example does not explicitly return
a Promise; instead, we just declare `next()` to be async**

```js
function clock(interval, max = Infinity) {
  function until(time) {
    return new Promise((resolve) => setTimeout(resolve, time - Date.now()));
  }
  return {
    startTime: Date.now(),
    count: 1,
    async next() {
      if (this.count > max) {
        return { done: true };
      }
      let targetTime = this.startTime + this.count * interval;
      await until(targetTime);
      return { value: this.count++ };
    },
    [Symbol.asyncIterator]() {
      return this;
    },
  };
}

async function test() {
  for await (const c of clock(300, 100)) {
    console.log(c);
  }
}

test();
```

This iterator-based version of the clock() function fixes a flaw in
the generator-based version.

- Time accuracy: If we use clock() with a for/await loop, this
  version will run loop iterations more precisely at the specified interval
  because it accounts for the time required to actually run the body of the
  loop.

- If we use `next()` instead of `for/await` loop, instant call of `next()` returns
  three Promises that will all fulfill at almost exactly the same time if we use generator-based
  version.

The benefit of asynchronous iterators is that they allow us to represent
streams of asynchronous events or data. When we are trying to work with other asynchronous
sources, such as the triggering of event handlers, it becomes substantially harder to
implement asynchronous iterators -- we typically have a single event handler function
that responds to events, but each call to the iterator's `next()` method
must return a distinct Promise object, and multiple calls to `next()` may occur
before the first Promise resolves. **This means that any asynchronous iterator
method must be able to maintain an internal queue of Promises that
it resolves in order as asynchronous events are occurring** -- `AsyncQueue` class.

The AsyncQueue class has `enqueue()` and `dequeue()` method. `dequeue()` method
returns a Promise rather than an actual value, however, **which means that it is
OK to call `dequeue()` before `enqueue()` has ever been called**. The
AsyncQueue class is also an asynchronous iterator, and is intended to be used with
a `for/await` loop whose body runs once each time a new value is asynchronously enqueued.

AsyncQueue has a `close()` method. **Once called, no more values can be enqueued**.
When a closed queue is empty, the `for/await` loop will stop looping.

Note that the implementation of AsyncQueue does not use async or
await and instead works directly with Promises. The code is
somewhat complicated, and you can use it to test your understanding of
the material we’ve covered in this long chapter. Even if you don’t fully
understand the AsyncQueue implementation, do take a look at the
shorter example that follows it: it implements a simple but very
interesting asynchronous iterator on top of AsyncQueue.

```js
class AsyncQueue {
  constructor() {
    this.values = [];
    this.resolvers = [];
    this.closed = false;
  }

  enqueue(value) {
    if (this.closed) {
      throw new Error("AsyncQueue closed");
    }
    if (this.resolvers.length > 0) {
      const resolve = this.resolvers.shift();
      resolve(value);
    } else {
      this.values.push(value);
    }
  }

  dequeue() {
    if (this.values.length > 0) {
      const value = this.values.shift();
      return Promise.resolve(value);
    } else if (this.closed) {
      // Promise for the "end-of-stream" marker
      return Promise.resolve(AsyncQueue.EOS);
    } else {
      return new Promise((resolve) => {
        this.resolvers.push(resolve);
      });
    }
  }

  close() {
    while (this.resolvers.length > 0) {
      this.resolvers.shift()(AsyncQueue.EOS);
    }
    this.closed = true;
  }

  [Symbol.asyncIterator]() {
    return this;
  }

  next() {
    return this.dequeue().then((value) =>
      value === AsyncQueue.EOS
        ? { value: undefined, done: true }
        : { value: value, done: false },
    );
  }
}

AsyncQueue.EOS = Symbol("end-of-stream");
```

```js
function eventStream(elt, type) {
  const q = new AsyncQueue();
  elt.addEventListener(type, (e) => q.enqueue(e));
  return q;
}

async function handleKeys() {
  for await (const event of eventStream(document, "keypress")) {
    console.log(event.key);
  }
}
```

## 14 - Metaprogramming

This chapter covers a number of advanced JavaScript features
that are not commonly used in day-to-day programming but that may be
valuable to programmers writing reusable libraries and of interest to
anyone who wants to tinker with the details about how JavaScript objects
behave.

- [Controlling the enumerability, deletability, and configurability of object properties](#14.1)

- [Controlling the extensibility of objects, and creating "sealed" and "frozen" objects](#14.2)

- [Querying and setting the prototypes of objects](#14.3)

- [Fine-Tuning the behavior of your types with well-known Symbols](#14.4)

- [Creating DSLs (domain-specific languages) with template tag functions](#14.5)

- [Probing objects with `reflect` methods](#14.6)

- [Controlling object behavior with Proxy](#14.7)

<a id="14.1"></a>

### 14.1 Property Attributes -- Control

| Attribute    | Description                                                                                 |
| ------------ | ------------------------------------------------------------------------------------------- |
| writable     | whether or not the value of property can change                                             |
| enumerable   | whether the property is enumerated by the `for/in` loop and the `Object.keys()` method      |
| configurable | whether a property can be deleted and also whether the property's attributes can be changed |

Property defined in object literals or by ordinary assignment to an
object are **writable, enumerable, and configurable**. But many
of the properties defined by the JavaScript standard library are not.

The four attributes of a data property are:

- `value`
- `writable`
- `enumerable`
- `configurable`

The four attributes of a accessor (getter and setter) are:

- `get`
- `set`
- `enumerable`
- `configurable`

The JavaScript methods for querying and setting the attributes
of a property use an object called a **property descriptor** to represent
the set of four attributes. The `writable`, `enumerable`, and `configurable`
properties are boolean values, and the `get` and `set` properties are function values.

To obtain the property descriptor for a named property of a specified object,
call `Object.getOwnPropertyDescriptor()`:

```js
Object.getOwnPropertyDescriptor({ x: 1 }, "x");
// => { value: 1, writable: true, enumerable: true, configurable: true }

const random = {
  get octet() {
    return Math.floor(Math.random() * 256);
  },
};

Object.getOwnPropertyDescriptor(random, "octet");
// {
//   get: [Function: get octet],
//   set: undefined,
//   enumerable: true,
//   configurable: true
// }

Object.getOwnPropertyDescriptor({}, "x"); // => undefined
Object.getOwnPropertyDescriptor({}, "toString"); // => undefined; inherited
```

**As its name implies, `Object.getOwnPropertyDescriptor` works only for own properties**
To query attributes of inherited properties, see [here to traverse the prototype chain (Object.getPrototypeOf())](#14.3).

To set the attributes of a property or to create a new property with the specified
attributes, call `Object.defineProperty()`:

```js
let o = {};
Object.defineProperty(o, "x", {
  value: 1,
  writable: true,
  enumerable: false,
  configurable: true,
});

o.x; // => 1
Object.keys(o); // => []

// Modify the property x to read-only, other properties remains unchanged
Object.defineProperty(o, "x", { writable: false });
o.x = 2; // Fails
o.x; // => 1

// The property is still configurable, so we can change its value (attribute)
Object.defineProperty(o, "x", { value: 2 });
o.x; // => 2

// Now change x from a data property to an accessor property
Object.defineProperty(o, "x", {
  get: function () {
    return 0;
  },
});
o.x; // => 0
```

**Tip**: If you're creating a new property, then omitted attributes are taken to be `false`
or `undefined` If you’re modifying an existing property, then the
attributes you omit are simply left unchanged.

If you want to create or modify more than one property at a time, use
`Object.defineProperties()`

```js
let p = Object.defineProperties(
  {},
  {
    x: { value: 1, writable: true, enumerable: true, configurable: true },
    y: { value: 1, writable: true, enumerable: true, configurable: true },
    r: {
      get() {
        return Math.sqrt(this.x * this.x + this.y * this.y);
      },
      enumerable: true,
      configurable: true,
    },
  },
);
p.r; // => 1.4142135623730951
```

The `Object.create()` method accepts a second optional argument, which
is the same as the second argument to `Object.defineProperties()`, they
are used to add properties to the newly created object.

`Object.defineProperty()` and `Object.defineProperties` throw TypeError if
the attempt to create or modify a property is not allowed.
This happens if you a attempt to add a new property to `non-extensible` object.
Another reason is the `configurable` attribute governs attempts to change
the other attributes.

Below listed the complete rules for attribute settings, violating them throw a TypeError:

- If an object is not extensible, you can edit its existing own properties, but
  you cannot add new properties to it.

- If a property is not configurable, you cannot change its configurable or enumerable
  attributes.

- If an accessor property is not configurable, you cannot change its getter or setter
  method, and you cannot change it to a data property.

- If a data property is not configurable, you cannot change it to an accessor property.

- If a data property is not configurable, you cannot change its _writable_ attribute
  from `false` to `true`, but you can change it from `true` to `false`

- If a data property is not configurable and not writable, you cannot change its value.

**`Object.assign()` only copies enumerable properties**, **and property value, not property attributes**

Note that the getter function below is copied to the target object,
not the getter function.

```js
let o = {
  x: 1,
  get hello() {
    return this.x;
  },
};
Object.defineProperty(o, "x", { configurable: false, writable: false });
Object.getOwnPropertyDescriptor(o, "hello");
// {
//   get: [Function: get hello],
//   set: undefined,
//   enumerable: true,
//   configurable: true
// }

let o1 = Object.assign({}, o);
Object.getOwnPropertyDescriptor(o1, "hello");
// => { value: 1, writable: true, enumerable: true, configurable: true }
Object.getOwnPropertyDescriptor(o1, "x");
// => { value: 1, writable: true, enumerable: true, configurable: true }
```

Below demonstrates how we can use
`Object.getOwnPropertyDescriptor()` and `Object.defineProperty()` to
create a variant of `Object.assign()` that copies entire property descriptors
rather than just copy property values.

```js
Object.defineProperty(Object, "assignDescriptors", {
  writable: true,
  enumerable: false,
  configurable: true,
  value: function (target, ...sources) {
    for (let source of sources) {
      for (let name of Object.getOwnPropertyNames(source)) {
        let desc = Object.getOwnPropertyDescriptor(source, name);
        Object.defineProperty(target, name, desc);
      }
      for (let symbol of Object.getOwnPropertySymbols(source)) {
        let desc = Object.getOwnPropertyDescriptor(source, symbol);
        Object.defineProperty(target, symbol, desc);
      }
    }
    return target;
  },
});

let o = {
  c: 1,
  get count() {
    return this.c++;
  },
};
let p = Object.assign({}, o);
let q = Object.assignDescriptors({}, o);
p.count; // => 1
p.count; // => 1
q.count; // => 2: increment once when we copied it the first time in Object.assign
q.count; // => 3
```

<a id="14.2"></a>

### 14.2 Object Extensibility

The extensible attribute of an object specifies whether new properties
can be added to the object or not. Ordinary JavaScript objects are
extensible by default, but you can change that with the functions
described in this section.

To determine whether an object is extensible:

```js
Object.isExtensible(o);
```

To make an object non-extensible:

```js
Object.preventExtensions(o);
o.g = 3; // => silently in non strict mode, but TypeError in strict mode
o.g; // => undefined in non strict mode
```

Attempting to change the [prototype](#14.3) of a non-extensible object will
always throw a TypeError.

```js
let o = { a: 1 };
let p = { b: 2 };
Object.preventExtensions(o);
Object.setPrototypeOf(o, p); // TypeError: #<Object> is not extensible
```

**There is no way to make an object extensible again once you have made it non-extensible**.

The purpose of the extensible attribute is to able to lock down the objects
into a known state and prevent outside tampering. The _extensible_
attribute of objects is often used in conjunction with the
_configurable_ and _writable_ attributes of properties, and JavaScript
defines functions that make it easy to set these attributes together:

- `Object.seal()`

  - `Object.preventExtensions()`

  - makes all of the own properties nonconfigurable

  - There is no way to unseal a sealed object

  - Use `Object.isSealed()` to determine whether an object is sealed.

- `Object.freeze()`

  - `Object.preventExtensions()`

  - makes all of the own properties nonconfigurable

  - makes all of the own data properties read only (If the object has accessor properties with
    setter methods, these are not affected and can still be invoked by assignment to the property.)

  - Use `Object. isFrozen()` to determine if an object is frozen
  - If it is frozen, `Object.isSealed()` returns true.

`Object.preventExtensions()`, `Object.seal()`, and `Object.freeze()` all return
the object that they are passed, which means that you can use them in nested function
invocations:

```js
let o = Object.seal(
  Object.create(Object.freeze({ x: 1 }), { y: { value: 2, writable: true } }),
);
```

If you are writing a JavaScript library that passes objects to callback functions
written by the users of your library, you might use
`Object.freeze()` on those objects to prevent the user's code from
modifying them. This is easy and convenient to do, but there are trade-offs:
frozen objects can interfere with common JavaScript testing strategies, for example.

<a id="14.3"></a>

### 14.3 The prototype Attribute

To determine whether one object is the prototype of
another object:

```js
let p = { x: 1 };
let o = Object.create(p);
p.isPrototypeOf(o); // => true
Object.prototype.isPrototypeOf(p); // =>true
```

**Note that `isPrototypeOf()` performs a function similar to the
`instanceof` operator.**

You can change the prototype of an object with `Object.setPrototypeOf()`

```js
let o = { x: 1 };
let p = { y: 2 };
Object.setPrototypeOf(o, p);
o.y; // =>  2
let a = [1, 2, 3];
Object.setPrototypeOf(a, p);
a.join; // => undefined; a no longer has a join() method
```

**There is generally no need to ever use `Object.setPrototypeOf()`.
JavaScript implementations may make aggressive optimizations based
on the assumption that the prototype of an object is fixed and unchanging. This means that if you ever call `Object.setPrototypeOf()`, any
code that uses the altered objects may run much slower than it would normally.**

Some early browser implementations of JavaScript exposed the
`prototype` attribute of an object through the `__proto__` property.
**This has long since been deprecated.** In modern JavaScript,
`__proto__` is readable and writable, and you can **(though you
shouldn't)** use it as an alternative to `Object.getPrototypeOf()`
and `Object.setPrototypeOf()`. One interesting use of
`__proto__`, however, is to define the prototype **of an object literal**:

```js
let p = { z: 3 };
let o = {
  x: 1,
  y: 2,
  __proto__: p,
};
o.z; // => 3
p.isPrototypeOf(o); // => true
```

<a id="14.4"></a>

### 14.4 Well-Known Symbols

The Symbol type was added to JavaScript in `ES6`, and one of the main
reasons for doing so was to safely add extensions to the language without
breaking compatibility with code already deployed on the web.

#### 14.4.1 Symbol.iterator and Symbol.asyncIterator

Make objects or classes iterable or asynchronously iterable.

#### 14.4.2 Symbol.hasInstance

When we describe [instanceof](#4.9.4):

To evaluate expression o instanceof f, JavaScript evaluates f.prototype,
and then looks for that value in the prototype chain of o.
If it finds it, then o is an instance of f.

In `ES6` and beyond, `Symbol.hasInstance` provides an alternative.
In `ES6`, if the righthand side of `instanceof` is any object with a
[Symbol.hasInstance] method, then that method is invoked with the lefthand
side value as its argument.

```js
let uint8 = {
  [Symbol.hasInstance](x) {
    return Number.isInteger(x) && x >= 0 && x <= 255;
  },
};
128 instanceof uint8; // => true
256 instanceof uint8; // => false
```

Note that this example is clever but confusing: it uses a
non class object where a class would normally be expected. You should
write a `isUint8()` function instead of relying on this `Symbol.hasInstance`

<a id="14.4.3"></a>

#### 14.4.3 Symbol.toStringTag

If you invoke the `toString()` method of a basic JavaScript object. You get the
`[object Object]`

If you invoke this same `Object.prototype.toString`, you get some interesting
results:

```js
Object.prototype.toString.call([]); // => "[object Array]"
Object.prototype.toString.call(/./); // => "[object RegExp]"
Object.prototype.toString.call(() => {}); // => "[object Function]"
Object.prototype.toString.call(""); // => "[object String]"
Object.prototype.toString.call(0); // => "[object Number]"
Object.prototype.toString.call(false); // => "[object Boolean]"
```

It turns out that you can use this `Object.prototype.toString.call()`
technique with any JavaScript value to obtain the "class attribute" of an
object that contains type information that is not otherwise available.
The following `classof()` function is arguably more useful than the `typeof` operator:

```js
function classof(o) {
  return Object.prototype.toString.call(o).slice(8, -1);
}

classof(null); // => "Null"
classof(undefined); // => "Undefined"
classof(1); // => "Number"
classof(10n ** 100n); // => "BigInt"
classof(""); // => "String"
classof(false); // => "Boolean"
classof(Symbol()); // => "Symbol"
classof({}); // => "Object"
classof([]); // => "Array"
classof(/./); // => "RegExp"
classof(() => {}); // => "Function"
classof(new Map()); // => "Map"
classof(new Set()); // => "Set"
classof(new Date()); // => "Date"
```

Prior to `ES6`, the behavior of the `Object.prototype.toString` method return [object Object] for non built-in types. However, in `ES6`, it looks for a property with
the symbolic name `Symbol.toStringTag` on its argument, and if such
a property exists, it uses the property value in its output.

```js
function classof(o) {
  return Object.prototype.toString.call(o).slice(8, -1);
}

class Range {
  get [Symbol.toStringTag]() {
    return "Range";
  }
}
Object.prototype.toString.call(new Range()); // => [object Range]
classof(new Range()); // => "Range"

class Range1 {}
Object.prototype.toString.call(new Range1()); // => [object Object]
classof(new Range1()); // => "Object"
```

#### 14.4.4 Symbol.species

```js
class EZArray extends Array {
  get first() {
    return this[0];
  }
  get last() {
    return this[this.length - 1];
  }
}
let e = new EZArray(1, 2, 3);
let f = e.map((x) => x * x);
e.last; // => 3
f.last; // => 9
```

Array defined methods `concat()`, `filter()`, `map()`, `slice()`, and `splice()`
which return arrays. In case EZArray, `ES6` specification by default will return instance of subclass instead
of Array.

Here's how it works:

1. In `ES6` and later, the `Array()` constructor has a property `Symbol.species`

2. When we create a subclass with `extends` the resulting subclass inherits `Symbol.species`.

3. Methods like `map()` and `slice()` that create and return new arrays are tweaked slightly in `ES6` and later. Instead of just creating a regular Array, they invoke `new this.constructor[Symbol.species]()` to create the new array.

Sometimes this default behavior is not what you want. If you wanted the
array-returning methods of EZArray to return regular Array objects:

```js
// This does not work... read-only property
EZArray[Symbol.species] = Array;

Object.defineProperty(EZArray, Symbol.species, { value: Array });
```

The simplest option:

```js
class EZArray extends Array {
  static get [Symbol.species]() {
    return Array;
  }
  get first() {
    return this[0];
  }
  get last() {
    return this[this.length - 1];
  }
}
let e = new EZArray(1, 2, 3);
let f = e.map((x) => x * x);
e.last; // => 3
f.last; // => undefined
```

Other than Array, Typed array, ArrayBuffer (slice method), Promise methods (like then() that return new Promise objects)
create those objects via this species protocol as well. Finally, if you
find yourself subclassing Map and defining methods that return new Map objects,
you might want to use `Symbol.species` yourself for the benefit of subclasses of your subclass.

#### 14.4.5 Symbol.isConcatSpreadable

Before `ES6`, `concat()` just use Array.isArray()
to determine whether to treat a value as an array or not. In `ES6`
the algorithm is changed slightly: If the argument (or the `this` value)
to `concat()` is an object and has a property with the symbolic name
`Symbol.isConcatSpreadable`, then the boolean value of that property is used
to determine whether the argument should be spread.

**Recall of `concat()`**: **If the argument is not an array, it simply is appended to the new array.**

```js
let a = [1, 2, 3];
a.concat(4, 5); // => [1,2,3,4,5]
a.concat([4, 5], [6, 7]); // => [1,2,3,4,5,6,7]
a.concat([4, [5]]); // => [1,2,3,4,[5]]
```

There are two cases when you might want to use this Symbol:

- If you create an Array-like object and want it to behave like a real
  array when passed to `concat()`:

  - ```js
    let arraylike = {
      length: 1,
      0: 1,
      [Symbol.isConcatSpreadable]: false
    }
    [].concat(arraylike) // => [1]: would be [{ '0': 1, length: 1 }] if not spread
    ```

- Array subclasses are spreadable by default, if you are defining an array
  subclass that you do not want to act like an array when used with `concat()`:

  - ```js
    class NonSpreadableArray extends Array {
      get [Symbol.isConcatSpreadable]() { return false; }
    }
    let a = new NonSpreadableArray(1,2,3)
    [].concat(a).length // => 1 ([[1,2,3]]) (would be 3 elements long if a was spead)
    ```

#### 14.4.6 Pattern-Matching Symbols

For each of the string method `match()`, `matchAll()`, `search()`,
`replace()`, and `split()`, there is a corresponding well-known
Symbol: `Symbol.match`, `Symbol.search`, and so on.

You can define your own pattern classes using the well-known Symbol methods
to provide custom matching.

In general, when you invoke one of these five String methods on a pattern object
like this:

```js
string.method(pattern, arg);
```

That invocation turns into an invocation of a symbolically named
method on your pattern object:

```js
pattern[symbol](string, arg);
```

Example:

```js
class Glob {
  constructor(glob) {
    this.glob = glob;

    // ? match any one character except /, and * matches zero or more
    // of those characters.
    let regexpText = glob.replace("?", "([^/])").replace("*", "([^/]*)");

    this.regexp = new RegExp(`^${regexpText}$`, "u");
  }

  toString() {
    return this.glob;
  }
  [Symbol.search](s) {
    return s.search(this.regexp);
  }
  [Symbol.match](s) {
    return s.match(this.regexp);
  }
  [Symbol.replace](s, replacement) {
    return s.replace(this.regexp, replacement);
  }
}

let pattern = new Glob("docs/*.txt");
"docs/js.txt".search(pattern); // => 0
"docs/js.htm".search(pattern); // => -1
let match = "docs/js.txt".match(pattern);
match[0]; // => "docs/js.txt"
match[1]; // => "js"
match.index; // => 0
"docs/js.txt".replace(pattern, "web/$1.htm"); // => "web/js.htm"
```

<a id="14.4.7"></a>

#### 14.4.7 Symbol.toPrimitive

[Chapter 3.9.3 - Object to Primitive Conversions](#3.9.3) explained that JavaScript has three slightly different
algorithms for converting objects to primitive values.

In `ES6`, a well-known Symbol `Symbol.toPrimitive` allows you
to override this default object-to-primitive behavior and gives you
complete control over how instances of your own classes will be
converted to primitive values.

The argument tells you what kind of conversion JavaScript
is trying to do on your object:

- "string": eg. interpolate the object into template literal

- "number": used with `<` or `>` or `-` or `*` etc

- "default": it means tat JavaScript is converting your object
  in a context where either a numeric or string value could work. This
  happens with the `+`, `==`, `!=` operators.

If you want instances of your class to be comparable and sortable with
`<` and `>`, then that is a good reason to define [`Symbol.toPrimitive`] method.

```js
class Hello {
  [Symbol.toPrimitive](kind) {
    console.log(kind);
  }
}

let hello = new Hello();

1 + hello; // "default"
1 * hello; // "number"
`hoho ${hello}`; // "string"
```

#### 14.4.8 Symbol.unscopables

`ES6`

As a workaround for compatibility issues caused by the deprecated `with` statement.
Recall that the `with` statement takes an object and executes its statement body as if it were in a scope where
the properties of that object were variables. This caused compatibility
problems when new methods were added to the Array class, and it
broke some existing code.
When used with an object `o`, a `with` statement computes
`Object.keys(o[Symbol.unscopables]||{})` and ignores properties
whose names are in the resulting array when creating the simulated scope in which to
execute its body.

```js
class Hello {
  a = 1;
  b = 3;
  [Symbol.unscopables] = { a: true, b: true };
}
with (new Hello()) {
  a;
  b;
} // => Uncaught ReferenceError: a is not defined
```

<a id="14.5"></a>

### 14.5 Template Tags

When an expression whose value is a function
is followed by a template literal, it turns into a function invocation, and
we call it a **"tagged template literal"**. Tagged template literal are often
used to define DSLs -- domain-specific languages -- and defining a new tag function
is
like adding new syntax to JavaScript.

The first argument is an array of strings, and this is followed by zero
or more additional arguments (depends on the number of values that are
interpolated), which can have values of any type.

The following `html\`\`` template is useful when you want to safely
interpolate values into a string of HTML:

```js
function html(strings, ...values) {
  let escaped = values.map((v) =>
    String(v)
      .replace("&", "&amp;")
      .replace("<", "&lt;")
      .replace(">", "&gt;")
      .replace('"', "&quot;")
      .replace("'", "&#39;"),
  );
  let result = strings[0];
  for (let i = 0; i < escaped.length; i++) {
    result += escaped[i] + strings[i + 1];
  }
  return result;
}
let operator = "<";
html`<b>x ${operator} y</b>`; // '<b>x &lt; y</b>'

let kind = "game",
  name = "D&D";
html`<div>class="${kind}>${name}</div>`; // => '<div> class="game>D&amp;D</div>'
```

An example of a tag function that does not return a string but instead a parsed
representation of a string:

```js
function glob(strings, ...values) {
  let s = strings[0];
  for (let i = 0; i < values.length; i++) {
    s += values[i] + strings[i + 1];
  }
  return new Glob(s);
}

let root = "/tmp";
let filePattern = glob`${root}/*.html`;
"/tmp/test.html".match(filePattern)[1]; // => test
```

One of the features mentioned in passing in [3.3.4](#3.3.4) is the
`String.raw\`\`` tag function that returns a string in its "raw" form.

The first argument, the array, has a property named `raw`, and the
value of that property is another array of strings, with the
same number of elements, includes strings in which escape sequences are not
interpreted. This obscure feature is important if you want to define a DSL with
a grammar that uses backslashes.

For example, if we wanted our `glob` tag function to support pattern matching
on Windows-style path (which use backslashes instead of forward slashes)
and we did not want users to have to double every backslash, we could
rewrite that function to use `strings.raw[]` instead of `strings[]`
The downside is that we could no longer use escape like `\u` in our glob literals.

<a id="14.6"></a>

### 14.6 The Reflect API

`ES6`

Although the `Reflect` functions do not provide any new features,
they do group the features together in one convenient API.
And importantly, the set of `Reflect` functions maps one-to-one with
the set of Proxy handler methods that we'll learn about in [14.7](#14.7)

#### `Reflect.apply(f, o, args)`

Invokes the function `f` as a method of `o`. It is equivalent to f.apply(o, args)

```js
class A {
  f() {
    console.log(this.a);
  }
}

let f = new A().f;
let o = { a: 1 };
Reflect.apply(f, o, []);
```

#### `Reflect.construct(c, args, newTarget)`

Invokes the constructor `c` is if the `new` keyword had been used.

If the optional `newTarget` argument is specified, it is used as the value of `new.target`
within the constructor invocation. If not specified, then the `new.target` value will be c.

```js
class A {
  constructor(a) {
    console.log(a);
  }
}
Reflect.construct(A, ["hello"]); // => "hello"
```

#### `Reflect.defineProperty(o, name, descriptor)`

It is very similar to `Object.defineProperty()` but returns `true` on success
and `false` on failures. `Object.defineProperty()` returns `o` on success and
throws TypeError on failure.

```js
let o = {};
Reflect.defineProperty(o, "a", { value: 123 });
```

#### `Reflect.deleteProperty(o, name)`

Returning `true` if successful (or no such property existed), `false` if failed.
Similar to `delete o[name]`

```js
Reflect.deleteProperty(o, "a");
```

#### `Reflect.get(o, name, receiver)`

Returns the value of the property of `o`.

If the property is an accessor method with a getter, and if the optional `receiver`
argument is specified, then the getter function is called as a method of
`receiver` instead of as a method of `o`. Calling this function is similar to
evaluating `o[name]`

```js
class B {
  name = "hello";
}

class C {
  get a() {
    return this.name;
  }
}

console.log(Reflect.get(new C(), "a", new B())); // => "hello"
```

#### `Reflect.getOwnPropertyDescriptor(o, name)`

This function is nearly identical to `Object.getOwnPropertyDescriptor()`,
except that the Reflect API version of the function requires the first
argument be an object and throws TypeError if it is not.

#### `Reflect.getPrototypeOf(o)`

This function returns the prototype of object o or null if the
object has no prototype. It throws a TypeError if o is a primitive
value instead of an object.

This function is almost identical to `Object.getPrototypeOf()` except that
`Object.getPrototypeOf()` only throws a TypeError for `null` and `undefined`
arguments and coerces other primitive values to their wrapped objects.

#### `Reflect.has(o, name)`

Returns true if the object o has a property with the
specified name (which must be a string or a symbol). Calling this
function is similar to evaluating name in o.

#### `Reflect.ownKeys(o)`

Calling this
function is similar to calling
Object.getOwnPropertyNames() and
Object.getOwnPropertySymbols() and combining their
results.

#### `Reflect.preventExtensions(o)`

This function sets the extensible attribute (§14.2) of the object o to
false and returns true to indicate success. It throws a
TypeError if o is not an object.

Object.preventExtensions() has the same effect but
returns o instead of true and does not throw TypeError for
non-object arguments.

#### `Reflect.set(o, name, value, receiver)`

This function sets the property with the specified name of the
object o to the specified value. It returns true on success and
false on failure (which can happen if the property is read-only).
It throws TypeError if o is not an object. If the specified property is
an accessor property with a setter function, and if the optional
receiver argument is passed, then the setter will be invoked as a
method of receiver instead of being invoked as a method of o.
Calling this function is usually the same as evaluating o[name] =
value.

#### `Reflect.setPrototypeOf(o, p)`

This function sets the prototype of the object o to p, returning
true on success and false on failure (which can occur if o is
not extensible or if the operation would cause a circular prototype
chain). It throws a TypeError if o is not an object or if p is neither
an object nor null. Object.setPrototypeOf() is similar,
but returns o on success and throws TypeError on failure.
Remember that calling either of these functions is likely to make
your code slower by disrupting JavaScript interpreter
optimizations.

<a id="14.7"></a>

### 14.7 Proxy Objects

`ES6`

JavaScript's most powerful metaprogramming feature.

It allows us to write code that alters the fundamental behavior of JavaScript objects.

When we create a Proxy object, we specify two other objects, the **target object**
and the **handlers object**:

```js
let proxy = new Proxy(target, handlers);
```

The resulting Proxy object has no state or behavior of its own. **Whenever
you perform an operation on it, it dispatches those operations to the handlers object
or to the target object**. The operations supported by Proxy objects
are the same as those defined by Reflect API. e.g. deleteProperty.

Proxy work this way for all of the fundamental operations: if an
appropriate method exists on the handlers object, it invokes that method.
If that method does not exist on the handlers object, then the Proxy performs
the fundamental operation on the target object. If the handlers is empty,
it is a transparent wrapper around the target object:

```js
let t = { x: 1, y: 2 };
let p = new Proxy(t, {});
p.x; // => 1
delete p.y; // => true
t.y; // => undefined
p.z = 3;
t.z; // => 3
```

Transparent wrappers can be useful, however, when created as "revocable proxies".
Instead of creating a Proxy with the `Proxy()` constructor, you can use the
`Proxy.revocable()` factory function. This function returns an
object that includes a Proxy object and also a `revoke()` function.
Once you call the `revoke()` function, the proxy immediately stops working:

```js
function accessTheDatabase() {
  return 42;
}

let { proxy, revoke } = Proxy.revocable(accessTheDatabase, {});

proxy(); // => 42
revoke();
proxy(); // => !TypeError
```

The main point here is that revocable proxies are a building block for a kind
of code isolation, and you might use them when dealing with untrusted third-party
libraries, for example. If you have to pass a function to a library that you don't
control, you can pass a revocable proxy instead and then revoke the proxy
when you are finished with the library. This prevents the library from keeping
a reference to your function and calling it at unexpected times.
This kind of defensive programming is not typical in JavaScript programs, but
the Proxy class at least makes it possible.

The following code, uses Proxy to create a read-only wrapper for a target object

```js
function readOnlyProxy(o) {
  function readonly() {
    throw new TypeError("Readonly");
  }
  return new Proxy(o, {
    set: readonly,
    defineProperty: readonly,
    deleteProperty: readonly,
    setPrototypeOf: readonly,
  });
}

let o = { x: 1, y: 2 };
let p = readOnlyProxy(o);
p.x = 2; // => TypeError : Readonly --- can't change
delete p.x; // => TypeError : Readonly
p.z = 3; // => TypeError : Readonly --- can't add
p.__proto__ = {}; // => TypeError: Readonly --- cant change the prototype
```

```js
function loggingProxy(o, objname) {
  const handlers = {
    get(target, property, receiver) {
      console.log(`Handler get(${objname}, ${property.toString()})`);

      let value = Reflect.get(target, property, receiver);

      if (
        Reflect.ownKeys(target).includes(property) &&
        (typeof value === "object" || typeof value === "function")
      ) {
        return loggingProxy(value, `${objname}.${property.toString()}`);
      }

      return value;
    },
    set(target, prop, value, receiver) {
      console.log(`Handler set(${objname},${prop.toString()},${value})`);
      Reflect.set(target, prop, value, receiver);
    },
    apply(target, receiver, args) {
      console.log(`Handler ${objname}(${args})`);
      return Reflect.apply(target, receiver, args);
    },
    construct(target, args, receiver) {
      console.log(`Handler ${objname}($args)`);
      return Reflect.construct(target, args, receiver);
    },
  };
  Reflect.ownKeys(Reflect).forEach((handlerName) => {
    if (!(handleName in handlers)) {
      handlers[handlerName] = function (target, ...args) {
        console.log(`Handler ${handlerName}(${objname}, ${args})`);
        return Reflect[handlerName](target, ...args);
      };
    }
  });
  return new Proxy(o, handlers);
}
```

The `loggingProxy()` defined earlier creates a proxies that log all of the ways
they are used. If you are trying to understand how an undocumented function uses the
objects you pass it, using a logging proxy can help.

```js
let data = [10, 20];
let methods = { square: (x) => x * x };

let proxyData = loggingProxy(data, "data");
let proxyMethods = loggingProxy(methods, "methods");

data.map(methods.square); // => [ 100, 400 ]
proxyData.map(methods.square); // => [ 100, 400 ]
// It produces:
Handler get(data, map)
Handler get(data, length)
Handler get(data, constructor)
Handler has(data, 0)
Handler get(data, 0)
Handler has(data, 1)
Handler get(data, 1)

data.map(proxyMethods.square); // => [100,400]
// It produces:
Handler get(methods, square)
Handler methods.square(10,0,10,20)
Handler methods.square(20,1,10,20)

for (let x of proxyData) console.log("Datum, x");
// Log output:
Handler get(data, Symbol(Symbol.iterator))
Handler get(data, length)
Handler get(data, 0)
Datum 10
Handler get(data, length)
Handler get(data, 1)
Datum 20
Handler get(data, length)
```

#### 14.7.1 Proxy Invariants

At the start of this section, we described proxies as objects with
no behavior of their own because they simply forward all operations
to he handlers object and the target object. But this is not entirely true:
after forwarding an operation, the Proxy class performs some sanity
checks on the result to ensure important JavaScript invariants are not
being violated. If it detects a violated, the proxy will throw a
TypeError instead of letting the operation proceed.

```js
let target = Object.preventExtensions({});
let proxy = new Proxy(target, {
  isExtensible() {
    return true;
  },
});
Reflect.isExtensible(proxy); // => TypeError: invariant violation
```

If the target object has nonwritable, non configurable properties, then
the Proxy class will throw a TypeError if the get() handler returns anything
other than the actual value

```js
let target = Object.freeze({ x: 1 });
let proxy = new Proxy(target, {
  get() {
    return 99;
  },
});
proxy.x; // !TypeError: value returned by get() doesn't match target
```

Proxy enforces a number of additional invariants, almost all of them
having to do with non-extensible target objects and nonconfigurable
properties on the target object.

## 15 - JavaScript In Web Browser

### 15.1 Web Programming Basics

#### 15.1.1 JavaScript in HTML `<script>` Tags

Web browser display HTML documents. If you want a web browser to execute
JavaScript code, you must include (or reference) that code from an HTML document,
and this is what the HTML `<script>` tag does.

Below makes a digital clock

```html
<!doctype html>
<html>
  <head>
    <title>Digital clock</title>
    <style>
      #clock {
        font: bold 24px san-serif;
        background #ddf;
        padding: 15px;
        border: solid black 2px;
        border-radius: 10px;
      }
    </style>
  </head>
  <body>
    <h1>Digital Clock</h1>
    <span id="clock"></span>
    <script>
      function displayTime() {
        let clock = document.querySelector("#clock");
        let now = new Date();
        clock.textContent = now.toLocaleTimeString();
      }
      displayTime();
      setInterval(displayTime, 1000);
    </script>
  </body>
</html>
```

![Digital Clock](images/digital-clock.png)

If we took the JavaScript code out of this HTML file and
stored it in its own `scripts/digital_clock.js` file, then the `<script>` tag
might reference that file of code like this:

```js
<script src="scripts/digital_clock.js"></script>
```

**Important**: The closing `</script>` is required, HTML does not support a `<script/>` self closing tag

There of a number of advantages to using the `src` attribute:

- It simplifies your HTML files -- keep content and behavior separate
- Using the `src` allows you to maintain only a single copy of that code when multiple
  web pages share the same JavaScript code.
- It only needs to be downloaded once if a file of JavaScript code is shared by more
  than one page -- subsequent pages can retrieve it from the browser cache.
- A JavaScript program or web page from one web server can employ code exported by other web servers.

##### MODULES

If you have written your JavaScript program using modules (and have
not used a code-bundling tool to combine all your modules into
a single non modular file of JavaScript), then you must
load the top-level module of your program with a `<script>` tag that
has a `type="module"` attribute.

##### SPECIFYING SCRIPT TYPE

In the early days of the web, it was thought that browsers might some
day implement languages other than JavaScript, and programmers
added attributes like `language="javascript"` and
`type="application/javascript"` to their `<script>` tags.

JavaScript is the default (and only) language of the web. The `language` attribute
is deprecated, and there are only two reasons to use a `type` attribute on a `<script>`
tag:

- To specify that the script is a module
- To embed data into a web page without displaying it [See 15.3.4](#15.3.4)

##### WHEN SCRIPTS RUN: ASYNC AND DEFFERRED

When JavaScript was first added to web browsers, **there was no API for
traversing and manipulating the structure and content of an already rendered document**.
The only way was to generate the content on the fly while the document was in the
process of loading by using `document.write()` method to inject HTML text into the document at the location of script.

**Tip**: `document.write()` has severe performance impact and their use is no longer considered acceptable.

When the HTML parser encounters a `<script>` element, it must, **by default**,
run the script. Fortunately, this default _synchronous_ or _blocking_script_
execution mode is not the only option. The `<script>` tag can have `defer` and
`async` attribute, which cause the scripts to be executed differently.
**These are boolean attribute (they don't have a value)**. These attributes are
only meaningful when used in conjunction with the `src` attribute:

```js
<script defer src="deferred.js"></script>
<script async src="async.js"></script>
```

Both the `defer` and `async` attributes are ways of telling the
browser that the linked script does not use `document.write()` to generate
HTML output, and that the browser, therefore, can continue to parse
and render the document while downloading the script.

The `defer` attribute causes the browser to defer execution of the script
until after the document has been fully loaded and parsed and is ready to be
manipulated.

The `async` attribute cause the browser to run the script as soon as possible
but does not block document parsing while the script is being
downloaded. If a `<script>` tag has both attributes, the `async` attribute takes
precedence.

Note that `defer` scripts run in the order in which whey appear in the document.
`async` scripts run as they load, which means that they may execute out of order.

Scripts with the `type="module"` attribute are, by default, executed after
the document has load, as if they had a `defer` attribute. You can override
this default with the `async` attribute, which will cause the code to be executed
as soon as the module and all of its dependencies have loaded.

A simple alternative to the `async` and `defer` attribute -- especially
for code that is included directly in the HTML -- is to simply put your
scripts at the end of the HTML file.

##### LOADING SCRIPTS ON DEMAND

Sometimes, you may have JavaScript code that is not used when a document
first loads and is only needed if the user takes some action like
clicking on a button or opening a menu. If you are developing your code
using modules, you can load a module on demand with `import()` [10.3.6](#10.3.6)

If you are not using modules, you can load a file of JavaScript
on demand simply by adding a `<script>` tag to your document when you want the
script to load:

```js
function importScript(url) {
  return new Promise((resolve, reject) => {
    let s = document.createElement("script");
    s.onload = () => {
      resolve();
    };
    s.onerror = (e) => {
      reject(e);
    };
    s.src = url;
    document.head.append(s);
  });
}
```

#### 15.1.2 The Document Object Model (DOM)

One of the most important objects in client-side JavaScript
programming is the Document object -- which represent the HTML
document that is displayed in a browser window or tab.

The API for working with HTML documents is known as the **Document
Object Model, or DOM**.

The DOM API mirrors the tree structure of an HTML document. **For each
HTML tag in the document, there is a corresponding JavaScript Element object**.
And **for each run of text in document, there is a corresponding Text object**.
The Element and Text classes, as well as the Document class itself,
are **all subclasses of the more general Node class**, and Node objects
are organized into a tree structure that JavaScript can query and traverse
using DOM API.

![Tree Html](images/tree-html.png)

There is a JavaScript class corresponding to each HTML tag type,
and each occurrence of the tag in a document is represented by an
instance of a class. For example, `<body>` tag is represented by
an instance of HTMLBodyElement, and a `<table>` tag is represented
by an instance of HTMLTableElement.

The JavaScript element objects have properties that correspond to the
HTML attributes of the tags. For example, instances of HTMLImageElement
which represent `<img>` tags, have a `src` property. Setting this property
causes the browser to load and display a new image.

#### 15.1.3 The Global Object in Web Browsers

There is **one global object per browser window or tab**. All of the
JavaScript code (except code running in worker thread [15.13](#15.13))
running in that window shares this single global object, **regardless of
how many scripts or modules are in the document**.

The global object is where JavaScript's standard library is defined. e.g. `fetch()`,
`parseInt()`, and the Math object, the Set class...

In web browsers, the global object does double duty:

- Defining built-in types and functions

- Represents the current web browser window and defines properties like `history`,
  `innerWidth` (window's width in pixels) (NOT EXCLUDING SCROLLBAR). One of the properties of this
  global object is named `window`, and its value is the global object itself.

**Tip**: When using window-specific features, it is often a good idea to include a
`window.` prefix: `window.innerWidth` is clearer than `innerWidth`

#### 15.1.4 Scripts Share a Namespace

With modules, the constants, variables, functions, and classes defined
at the top level of the module are private to the module unless they are
explicitly exported, in which case, they can be selectively imported by other modules.

With non-module scripts, however, the situation is completely different.
If the top-level code in a script defines a constant, variable, function, or class, that
declaration will be **visible to all other scripts in the same document**.

There are some historical quirks with how this non-module shared namespace works.
`var` and `function` declarations at the top level create
properties in the shared global object. On the other hand, the `ES6` declarations
`const`, `let`, `class`, when used at the top level, do not create properties
in the global object. They are still
defined in a shared namespace, however: **if one script defines a class C, other
scripts will be able to create instances of that class with `new C()`, but not with
`new window.C()`**

To summarize: in modules, top-level declarations are scoped to the
module and can be explicitly exported. In nonmodule scripts, however,
top-level declarations are scoped to the containing document, and the
declarations are shared by all scripts in the document. Older `var` and
`function` declarations are shared via properties of the global object.
Newer `const`, `let`, and `class` declarations are also shared and have
the same document scope, but they do not exist as properties of any
object that JavaScript code has access to.

#### 15.1.5 Execution of JavaScript Programs

If a webpage includes an embedded frame (using the `<iframe>` element), the
JavaScript code in the embedded document has a **different global object
and Document object than the code in the embedding document, and it can
be considered a separate JavaScript program.** However, they can interact with
each other [15.13.6 -- Cross-Origin Messaging with postMessage()](#15.13.6)

You can think of JavaScript program execution as occurring in two phases,

In the first phase, the document content is loaded, and the code from `<script>` elements
(both inline scripts and external scripts) is run.

Second phase is asynchronous and event-driven. If a script is going to participate
in this second phase, then one of the things it must have done during the first
phase is to register at least one event handler or other callback function that will
be invoked asynchronously. Some of the first events to occur during
event-driven phase are the "DOMContentLoaded" and "load" events. "DOMContentLoaded" is
triggered when the HTML document has been completely loaded and
parsed. The "load" event is triggered when all of the document's external
resources -- such as images -- are also fully loaded. It is
common to see programs whose scripts define functions but take no action other
than registering an event handler function to be triggered by the "load" event
at the beginning of the event-driven phase of execution. It is this "load"
event handler that then manipulates the document and does whatever it is that the program
is supposed to do.

##### CLIENT-SIDE JAVASCRIPT THREADING MODEL

JavaScript is a single-thread language: you can write code with the assurance that
two event handlers will never run at the same time. You can manipulate document
content knowing that no other thread is attempting to modify it at the same time,
and you never need to worry about locks, deadlocks, or race conditions when writing JavaScript
code.

Single-threaded execution means that web browsers stop responding to user input
while scripts and event handlers are executing. This places a burden
on JavaScript programmers: it means that JavaScript scripts and event handlers
must not run for too long.

The web platform defines a controlled form of concurrency called a "web worker".
A web worker is a background thread for performing computationally intensive tasks
without freezing the user interface. The code that runs in a web worker thread does not
have access to document content, does not share any state with the main thread or with
other workers, and can only communicate with the main thread and other workers
through asynchronous message events, so the concurrency is not detectable to the
main thread, and web workers do not alter the basic single-threaded execution model of
JavaScript programs. See [15.13](#15.13)

##### CLIENT-SIDE JAVASCRIPT TIMELINE

The two phase (script-execution phase and event-handling phase) can further broken
down into the following steps:

1. The web browser creates a Document object and begins parsing the web page,
   adding Element objects and Text nodes to the document as it parses HTML elements
   and their textual content. The `document.readyState` property has the value "loading"
   at this stage.

2. When the HTML parser encounters a `<script>` tag that **does not** have any
   of `async`, `defer`, and `type="module"` attributes, it adds the script tag to the
   document and then executes the script, **synchronously**, and the **HTML parser pauses
   while the script downloads and runs**

3. When the parser encounters a `<script>` element that has the `async`attribute
   set, it begins downloading the script text (and if the script is a module,
   it also recursively downloads all of the script's dependencies) and continues
   parsing the document. The script will be executed as soon as possible after it
   has downloaded, but the parser does not stop and wait for it to download.

4. When the document is completely parsed, the `document.readyState` property
   changes to "interactive"

5. Any scripts that had the `defer` attribute set (along with any
   module scripts that do not have an `async` attribute) are
   executed in the order in which they appeared in the document. Async scripts
   may also be executed at this time. **Deferred scripts have access to the
   complete document and they must not use the `document.write()` method**

6. The browser fires a "DOMContentLoaded" event on the Document object. This
   marks the transition from synchronous script-execution phase to the asynchronous,
   event-driven phase of program execution. Note, however, that there may
   still be `async` scripts that have not yet executed at this point.

7. The document is completely parsed at this point, but the browser may still be
   waiting for additional content, such as images, to load. When all such content
   finishes loading, and **when all `async` scripts have loaded and executed**, the
   `document.readyState` property changes to "complete", and the web browser fires a "load"
   event on the Window object.

8. From this point on, event handlers are invoked asynchronously in response to
   user input event, network events, time expirations, and so on.

#### 15.1.6 Program Input and Output

There are a variety of inputs available:

- The content of the document itself, which JavaScript code can access with the DOM API [15.3 -- Scripting Documents](#15.3)

- User input, in the form of events, such as mouse clicks. [15.2 -- Events](#15.2)

- The URL of the document -- `document.URL`

- The content of the HTTP "Cookie" request header -- `document.cookie`. [15.12.2 -- Cookies](#15.12.2)

- The global `navigator` property to access information about the web browser,
  and OS. For example, `navigator.userAgent` is a string that identifies the web browser,
  `navigator.language` is the user's preferred language, and `navigator.hardwareConcurrency` returns the number of logical CPUs available to the web browser. `screen` property
  provides access to the user's display size via the `screen.width` and `screen.height`.
  In a sense, these `navigator` and `screen` objects are to web browser what environment
  variables are to Node program.

Client-side JavaScript typically produces output, when it needs to by
manipulating the HTML document with the DOM API [15.3 -- Scripting Documents](#15.3) or by
using a higher-level framework such as React or Angular to manipulate
the document. `console.log()` also being used to produce non user-visible output.

#### 15.1.7 Program Errors

If an exception occurs while your JavaScript program is running, and if you do not
have a `catch` statement to handle it, an error message will be displayed in the developer
console, but any event handlers that have been registered keep running and responding
to events. JavaScript programs in a web browser can't really "crash".

If you would like to define an error handler of last resort to be invoked
when this kind of uncaught exception occurs, set the `onerror` property of the
Window object to an handler function. The first argument is a message describing
the error. The second argument is a string that contains the URL of
the JavaScript code that caused the error. The third argument is the
line number within the document where the error occurred. If the
`onerror` handler returns true, it tells the browser that the handler
has handled the error and that no further action is necessary -- in other words,
**the browser should not display its own error message.**

```js
window.onerror = (msg, url, line) => {
  console.log(msg);
  console.log(url);
  console.log(line);
  return true;
};
```

![Window Onerror](images/window-onerror.png)

For rejected Promise, use `window.onunhandledrejection` (or register with `window.addEventListener()`). The argument is a promise property whose value is the Promise object that
rejected, and a `reason` property whose value is what would have been passed to a `.catch()` function. If you call `preventDefault()` on this object, it will be
considered handled and won't cause an error message in the developer console.

```js
window.onunhandledrejection = (promise) => {
  console.log(promise);
  console.log(promise.reason);
  promise.preventDefault();
};
```

![Window Onunhandledrejection](images/window-onunhandledrejection.png)

It is not often necessary to define `onerror` or `onunhandledrejection` handlers,
but it can be quite useful as a telemetry mechanism if you want to report
client-side errors to the server (using the `fetch()` function to make an
HTTP POST request, for example) so that you can get information about unexpected
errors that happen in your users' browsers.

#### 15.1.8 The Web Security Model

Browser vendors have worked hard to balance two competing goals:

- Defining powerful client-side APIs to enable useful web applications

- Prevent malicious code from reading or altering your data, compromising
  your privacy, scamming you, or wasting your time

The subsections that follow give a quick overview of the security restrictions
and issues that you, as a JavaScript programmer, should to be aware of.

##### What JavaScript Can't Do

- does not provide any way to write or delete arbitrary files or list arbitrary directories on the client computer.

- client-side JavaScript does not have general-purpose networking capabilities.

##### The Same-Origin Policy

The _same-origin_ policy is a sweeping security restriction on the web content
JavaScript code can interact with. It typically comes into play when a web page
includes `<iframe>` elements. In this case, the same-origin policy governs the interactions
of JavaScript code in one frame with the content of other frames.
**Specifically, a script can read only the properties of windows and documents that
have the same origin as the document that contains the script**.

The origin of document is defined as the **protocol, host, and port** of the URL
from which the document was loaded. Documents loaded from different web servers
have different origins. Documents loaded through different ports of the
same host have different origins. And a document loaded with the `http:` protocol
has a different origin than one loaded with the `https:` protocol.

Browsers typically treat **every** `file: URL` as a separate origin, which mean you will
have to run a static web server during development to displays
more than one document from the same server.

It is important to understand that the origin of the script itself is not
relevant to the same-origin policy: **what matters is the origin of the document in
which the script is embedded**. For example, a webpage served by host B. And that
webpage has script src hosted by host A. The origin of this script is host B, and
the script has full access to the content of the document that contains it.

If the document contains an `<iframe>` that contains a second document from host B,
then the script has full access to this second document as well. But if the top-level
document contains another `<iframe>` that display a document from host C (**even one from host A)**, then the same-origin policy comes into effect and prevents the script from
accessing this nested document.

The same-origin policy also applies to scripted HTTP requests. JavaScript code
can make arbitrary HTTP requests to the web server from which the containing document
was loaded, but it does not allow scripts to communicate with other web servers
(unless those web servers opt in with CORS)

**Scripts** with origin _orders.example.com_ might need to read properties from documents
on _example.com_.

To support multidomain websites of this sort, scripts can
alter their origin setting `document.domain` to a domain suffix (https://example.com).
But **that script cannot set `document.domain` to "orders.example", "ample.com", or "com"**

The second technique for relaxing the same-origin policy is Cross-Origin Resource
Sharing, or CORS, which allows servers to decide which origins
they are willing to serve. CORS extends HTTP with a
new `Origin` request header and a new `Access-Control-Allow-Origin` response header.
It allows servers to use a header to explicitly list origins that may request
a file to use a wildcard and allow a file to be requested by any site. Browsers honor these
CORS headers and do not relax same-origin restrictions unless they are present.

**Tip**: `Access-Control-Allow-Origin` allows `*`, and full origin (include protocol, host, and optionally port). It does not allow
`https://*.abc.com`

##### Cross-site Scripting

Cross-site scripting, XSS, is a term for a category of security issues
in which an attacker injects HTML tags or scripts into a target website.
Client-side JavaScript programmers must be aware of, and defend
against, cross-site scripting.

As a trivial example, consider the following web page uses JavaScript to greet
the user by name:

```js
<h1></h1>
<script>
  let name = new URL(document.URL).searchParams.get("name")
  document.querySelector('h1').innerHTML = "Hello " + name;
</script>
```

If we invoke the page with below URL, the malicious code will be executed after the image
loads.

`http://www.example.com/greet.html?name=\<img src="x.png" onload="alert('hacked')"/>`

Cross-site scripting attacks are so called because more than one site is
involved. Site B includes a specially crafted link to site A (like the URL above). If site B
can convince users to click the link, they will be taken to site A, but that
site will now be running code from site B. That code might deface the page or cause
it to malfunction. More dangerously, the malicious code could read cookies stored
by site A (perhaps account numbers or other personally identifying information) and
send that data back to site B. The injected code could even track the user's keystrokes
and send that data back to site B.

In general, the way to prevent XSS attacks is to remove HTML tags from any untrusted
data before using it to create dynamic document
content:

```js
name = name
  .replace(/&/g, "&amp;")
  .replace(/</g, "&lt;")
  .replace(/>/g, "&gt;")
  .replace(/"/g, "&quot;")
  .replace(/'/g, "&#x27;")
  .replace(/\//g, "&#x2F;");
```

Another approach to the problem of XSS is to structure your web applications
so that untrusted content is always displayed in an `<iframe>` with the `sandbox`
attribute set to disable scripting and other capabilities

Cross-site scripting is a pernicious vulnerability whose roots go deep
into the architecture of the web. It is worth understanding this
vulnerability in-depth.

<a id="15.2"></a>

### 15.2 Events

Client-side JavaScript programs use an asynchronous event-driven programming model.

In client-side JavaScript, events can occur on any element within an HTML document,
and this fact makes the event model of web browsers significantly more complex
than Node's event model. Below are some definitions that help to explain that
event model:

- _event type_ or _event name_

  - What kind of event occurred. e.g "mousemove", "keydown", "load"

- _event target_

  - This is the object on which the event occurred or with which the event is associated

  - e.g. a `<button>` Element, Window, Document, Worker object

- _event handler_, or _event listener_

  - This function handles or responds to an event.

  - When event handlers are invoked for an object, we say that the browser has "fired",
    "triggered" or "dispatched" the event

- _event object_

  - This object is associated with a particular event and contains details about that event.

  - Are passed as an argument to the event handler function. They have a `type` property
    that specifies the **event type** and a `target` property that specifies the **event target**.

  - Each event type defines a set of properties for its associated event object. e.g. mouse event includes the coordinates of the mouse pointer.

- _event propagation_

  - When certain kinds of event occur on elements within the HTML document, they propagate
    or "bubble" up the document tree.

  - If the user moves the mouse over a hyperlink, the mousemove event is first fired
    on the `<a>` element that defines that link. Then it is fired on the containing element (e.g. a `<p>`, `<section>`, then Document object).

  - It is sometimes more convenient to register a single event handler on a Document
    or other container element than to register handlers on each individual element you're
    interested in.

  - An **event handler** can stop the propagation of an event so that it will not continue to bubble.

Some events have _default_ actions associated with them. When a click
event occurs on a hyperlink, for example, the default action is for the browser
to follow the link and load a new page. To prevent this default action, **event handlers** invoke a method of the event object.

#### 15.2.1 Event Categories

##### Device-dependent input events

Directly tied to a specific input device such as mouse or keyboard.

- `mousedown`, `mousemove`, `mouseup`

- `touchmove`, `touchend`

- `keydown`, `keyup`

- `scroll`

##### Device-independent input events

- `click` -- link or button has been activated.

  - could be done via a mouse click, keyboard, or a tap.

- `input` -- alternative to "keydown" event as well as alternatives such as cut-and-paste
  input methods used for ideographic scripts

- `pointerdown`, `pointermove`, `pointerup` are alternatives to mouse and touch events.

  - work for mouse-type pointers, touch screens, and for pen- or stylus-style input as well.

##### User interface events

- `focus` -- when a text input field gains keyboard focus

- `blur`

- `change` -- when the user changes the value displayed by a form element

- `submit` -- when the user clicks a submit button in a form

##### State-change events

Some events are not triggered directly by user activity, but by network or
browser activity, and indicate some kind of life-cycle or state-related change.

- `load` and `DOMContentLoaded` -- fired on the Window and Document object respectively.

- `online` and `offline` events on Window object when network connectivity changes

- `popstate` event in response to the browser's Back button

##### API-specific events

- `<video>` and `<audio>` elements

  - `waiting`, `playing`, `seeking`, `volumechange`, and so on

- IndexedDB API

  - `success`, `error` when database requests succeed or fail

#### 15.2.2 Registering Event Handlers

The first: set a property on the object e.g. `x.onclick`.

The second (newer and more general) technique: `addEventListener()`

##### Setting Event Handler Properties

By convention, event handler properties have names that consist of the word
"on" followed by the event name. Note that these property names are **case sensitive
and are written in ALL lowercase**, even when the event type such as mousedown consists
of multiple words.

**Tip**: If you have used the React framework to create client-side user interfaces, this may
surprise you. React makes a number of minor changes to the client-side event model, and
one of them is that in React, event handler property names are written in camelCase:
`onClick`, `onMouseOver`, and so on. When working with the web platform natively,
however, the event handler properties are written entirely in lowercase.

```js
window.onload = function () {
  let form = document.querySelector("form#shipping");
  form.onsubmit = function (event) {
    if (!isFormValid(this)) {
      event.preventDefault();
    }
  };
};
```

The shortcoming of event handler properties is that they are designed
around the assumption that event targets will have at most one handler
for each type of event. It is often better to use `addEventListener()` because that
technique **does not overwrite any previously registered handlers**

##### Setting Event Handler Attributes

The event handler properties of document elements can also be defined
directly in the HTML file as attributes on the corresponding HTML tag.
(Window element handler can be defined with attribute on the `<body>` tag in HTML).
This technique **is generally frowned upon in modern web development**.

```html
<button onclick="console.log('thank you')">Please Click</button>
```

When you specify a string of JavaScript code as the value of an HTML
event handler attribute, the browser converts your string into a function
that works something like this one:

```js
function(event) {
  with(document) {
    with(this.form || {}) {
      with(this) {
        /* your code here */
      }
    }
  }
}
```

The `event` argument means that your handler code can refer to the current event object as `event`.

Event handlers defined in this way are
executed in an environment in which unexpected variables are defined.
This can be a source of confusing bugs and is a good reason to avoid
writing event handlers in HTML.

##### `addEventListener()` and `removeEventListener()`

Any object that can be an event target defines a method `addEventListener()` that
you can use to register an event handler for that target.

It takes three arguments

- The first is the event type. **The event type does not include "on"**
- The second argument is the handler function
- The third (optional) argument is a boolean value or object.

  - if `true`, then your handler function is registered as a _capturing_ event handler
  - if you pass `true` here, you **must** pass `true` to `removeEventListener()`'s third
    argument
  - you can pass an object:

    - ```js
      document.addEventListener("click", handleClick, {
        capture: true,
        once: true, // automatically removed after it is triggered once if true
        passive: true, // never call `preventDefault()` to cancel the default action if true
      });
      ```

  - `passive` is particularly important for touch events on mobile devices,
    if event handlers for "touchmove" events can prevent the browser's default scrolling action, then the browser cannot implement smooth scrolling.

  - Smooth scrolling is so important for a good user experience that Firefox and
    Chrome make “touchmove” and “mousewheel” events passive by
    default. So if you actually want to register a handler that calls
    `preventDefault()` for one of these events, you should explicitly
    set the passive property to `false`

```html
<button id="mybutton">Click me</button>
<script>
  let b = document.querySelector("#mybutton");
  b.onclick = function () {
    console.log("Thanks for clicking me");
  };
  b.addEventListener("click", () => {
    console.log("Thanks again!");
  });
</script>
```

**Tip**: More importantly, you can call `addEventListener()` multiple times to
register more than one handler function for the same event type on the
same object.

`removeEventListener()` expects the same two arguments (plus one optional third).
It is often useful to temporarily register an event handler and then remove it soon
afterward. For example, when you get a "mousedown" event, you might register temporary
event handlers for "mousemove" and "mouseup" events so that you can see if
user drags the mouse. You'd then deregister these handlers when the "mouseup"
event arrives. In such a situation, your event handler removal code:

```js
document.removeEventListener("mousemove", handleMouseMove);
document.removeEventListener("mouseup", handlerMouseUp);
```

**Tip**: You can also pass an Options object to `removeEventListener()`,
but the `capture` property is the only one that is relevant. There is no
need to specify `once` or `passive` when removing a listener, and
these properties are ignored.

#### 15.2.3 Event Handler Invocation

##### Event Handler Argument

An Event object with properties:

- `type`

  - The type of the event that occurred

- `target`

  - The object on which the event occurred

- `currentTarget`

  - For events that propagate, this property is the object on which the
    current event handler was registered. (which is `this`)

- `timeStamp`

  - A timestamp (in milliseconds) that represents when the event occurred
    but that does not represent an absolute time. You can determine the elapsed
    time between two events by subtracting the timestamp of the first event from
    timestamp of the second.

- `isTrusted`

  - This property will be `true` if the event was dispatched by the web browser itself,
    and `false` if the event was dispatched by JavaScript code.

Specific kinds of events have additional properties e.g. `clientX` and `clientY`

##### Event Handler Context

Within the body of an event handler, the `this` keyword refers to the object on
which the event handler was registered (when use `function` expression).

##### Handler Return Value

In modern JavaScript, **event handlers should not return anything**.
You may see event handlers that return values in older code, the return
value is typically a signal to browser that it should not perform the default action
associated with the event.

**The standard and preferred way to prevent default action is to call
the `preventDefault()` method on the Event object.**

##### Invocation Order

If there are more than one event handler registered for a particular type of event,
the browser invokes all the handlers in the order in which they were registered.

#### 15.2.4 Event Propagation

After the event handlers registered on the target element are invoked,
most events "bubble" up the DOM tree. The event handlers of the target's parent are invoked. This continues up to the Document object, and then beyond to the Window object.
You can register a single handler on a common ancestor element and handle events there e.g. _change_ event in a `form` element.

Most events that occur on document bubble. **Notable exception are the
"focus", "blur", "scroll" events**. The "load" event on document elements bubbles,
but it stops bubbling at Document object and does not propagate on to the
Window object. (The "load" event handlers of the Window object are triggered
only when the entire document has loaded)

Event bubbling is the **third phase** of event propagation. The second
face is invocation of the event handlers. **The first phase, which before
the target handlers are invoked, is called the "capturing" phase**.
The capturing phase of event propagation is like the bubbling phase **in reverse**.

Event capturing provides **an opportunity to peek at events before they are
delivered to their target**. A capturing event handler can be used for
debugging, or it can be used along with the event cancellation technique (in next section)
to filter events so that the target event handlers are never actually invoked.

#### 15.2.5 Event Cancellation

You can prevent the browser from performing its
default action by invoking the `preventDefault()` method of the event object.

Browsers responds to many user events:

- `click` on a hyperlink, the browser follows the link

- `touchmove` on touch-screen device, the browser scrolls

We can also cancel the propagation of events by calling the `stopPropagation()` method
of the event object. If there are other handlers defined on the same object,
the rest of those handlers will still be invoked, but no event handlers on
any other object will be invoked after `stopPropagation()` is called.

```html
<body>
  <div id="bbb">
    <button id="bbb">JAJA</button>
  </div>
  <script>
    let d = document.querySelector("div#bbb");
    let btn = document.querySelector("button#bbb");
    btn.addEventListener("click", function (e) {
      e.stopPropagation();
    });
    btn.addEventListener("click", function (e) {
      // here will be invoked
    });
    d.addEventListener("click", function (e) {
      // here won't be invoked when btn was clicked
    });
  </script>
</body>
```

`stopPropagation()` works during the capturing phase.

```html
<body>
  <div id="bbb">
    <button id="bbb">JAJA</button>
  </div>
  <script>
    let d = document.querySelector("div#bbb");
    let btn = document.querySelector("button#bbb");
    btn.addEventListener("click", function (e) {
      // here won't be invoked when btn was clicked
    });
    d.addEventListener(
      "click",
      function (e) {
        e.stopPropagation();
      },
      true,
    ); // capturing handler
    d.addEventListener("click", function (e) {
      // here won't be invoked when btn was clicked
    });
  </script>
</body>
```

`stopImmediatePropagation()` works like `stopPropagation()`,
but it also prevents the invocation of any subsequent event handlers registered
on the same object:

```js
<body>
  <div id="bbb">
    <button id="bbb">JAJA</button>
  </div>
  <script>
    let d = document.querySelector("div#bbb");
    let btn = document.querySelector("button#bbb");
    btn.addEventListener("click", function (e) {
      e.stopImmediatePropagation();
    });
    btn.addEventListener("click", function (e) {
      // here won't be invoked when btn was clicked
    });
    d.addEventListener("click", function (e) {
      // here won't be invoked when btn was clicked
    });
  </script>
</body>
```

#### 15.2.6 Dispatching Custom Events

If a JavaScript object has an `addEventListener()` method, then it
is an "event target", and this means it also has a `dispatchEvent()`
method. **You can create your own event object with the `CustomEvent()`
constructor and pass it to `dispatchEvent()`**. The first argument
to `CustomEvent()` is a string that specifies the type of your event,
and the second argument is an object that specifies the
properties of the event object. Set the `detail` property of this object to a string,
object, or other value that represents the content of your event.
If you plan to dispatch your event on a document element and want it
to bubble up the document tree, add `bubbles: true` to the second argument.

```js
document.addEventListener("busy", (e) => {
  if (e.detail) {
    showSpinner();
  } else {
    hideSpinner();
  }
});

document.dispatchEvent(
  new CustomEvent("busy", { detail: true, bubbles: true }),
);

fetch(url)
  .then(handleNetworkResponse)
  .catch(handleNetworkError)
  .finally(() => {
    document.dispatchEvent(
      new CustomEvent("busy", { detail: false, bubbles: true }),
    );
  });
```

<a id="15.3"></a>

### 15.3 Scripting Documents

Every window object has a `document` property that refers to a Document object.
The Document object represents the content of the window. The Document object
does not stand alone, however, It is the central object in the DOM for representing
and manipulating document content.

#### 15.3.1 Selecting Document Elements

##### Selecting Elements With CSS Selectors

CSS stylesheets have a very powerful syntax, known as _selectors_,
for describing elements or sets of elements within a document.
The DOM methods `querySelector()` and `querySelectorAll()` allow
us to find the element or elements within a document that match
a specified CSS selector. You can use invoke it with any HTML Element.

**The `querySelector()` method takes a CSS selector string as its argument
and returns the FIRST matching element in the document that it finds,
or return `null` if none match**

`querySelectorAll()` is similar, but it returns **all matching elements in
the document**. It is an array-like object known as NodeList. It has a `length`
property and can be indexed like arrays. It is also iterable. If you want
to convert a NodeList into a true array, use `Array.from()`

CSS selector syntax:

```js
div             // Any <div> element
#nav            // The element with id="nav"
.warning        // Any element with "warning" in its class attribute
p[lang="fr"]    // A paragraph written in Frech: <p lang="fr">
*[name="x"]     // Any element with a name="x" attribute
a[href]         // <a> element with an attribute href
*[name]         // all element with an attribute name

// Examples:
// button[hello=bye][today=tmr]#a.warning
// span.fatal.error
// span[lang="fr"].warning
```

Selectors can also specify document structure:

```js
#log span           // Any <span> descendant (grand childrens...) of the element with id="log"
#log>span           // Any <span> child of the element id="log"
body>h1:first-child // The first <h1> child of the <body>
img + p.caption     // A <p> with class "caption" immediately **after** an <img> (!not inside!), (immediately! There must not have in between elements.)
h2 ~ p              // Any <p> that follows an <h2> and is a sibling of it, the difference with '+' is it is not immediate after, in between elements are allowed.
```

If two selectors are separated by a comma, it means that we've selected elements
that match either one of the selectors:

```js
button, input[(type = "button")]; // All <button> and <input type="button> elements"
```

Pseudo elements match portions of text nodes rather than actual elements (e.g. `::firstline::`, and `::first-letter::`).
**They will not match if used with `querySelectorAll()` or `querySelector()`.**
Also, many browser will refuse to return matches for the `:link` and `:visited`
pseudoclasses, as this could expose information about the user's browsing history.

Another CSS-based element selection method is `closest()`. This
method is defined by the Element class and takes a selector as its only
argument. **If the selector matches the element it is invoked on, it returns that element. Otherwise, it returns the closest ANCESTOR element that the selector matches**.

In a sense, `closest()` is the opposite of `querySelector()`: `closest()` looks for
a match above it in the tree. `closest()` can be useful when you have
registered an event handler at a high level in the document tree. If you
are handling a "click" event, for example, you might want to know whether it is a
click a hyperlink. The event object will tell you what the target was, but that
target might be the text inside a link:

```js
let hyperlink = event.target.closest("a[href]");
```

Another way that you might use `closest()`

```js
function insideList(e) {
  return e.closest("ul,ol,dl") !== null;
}
```

`matches()` does not return ancestors or descendants:
it simply tests whether an element is matched by a CSS
selector and returns `true` if so:

```js
function isHeading(e) {
  return e.matches("h1,h2,h3,h4,h5,h6");
}
```

##### Other Element Selection Methods

In addition to `querySelector()` and `querySelectorAll()`,
the DOM also defines a number of older element selection methods
that are more or less obsolete now. You may still see some of these methods
(especially `getElementById()`) in use, however.

```js
let sect1 = document.getElementById("sect1");
let colors = document.getElementsByName("color"); // equal to querySelectorAll('*[name=color]')
let headings = document.getElementsByTagName("h1"); // equal to querySelectorAll('h1')
let subheads = sect1.getElementsByTagName("h2"); // get all <h2> elements within the sect1 element, sect1.querySelectorAll('h2')
let tooltips = document.getElementsByClassName("tooltip"); // equal to querySelectorAll('.tooltip')
let sidebars = sect1.getElementsByClassName("sidebar"); // equal to sect1.querySelectorAll(".sidebar"), look up all descendants of sect1 that have class "sidebar"
```

Unlike `querySelectorAll()`, however, the NodeLists returned by these older selection
methods are "live", which means that the length and content of the list can change
if the document content or structure changes.

##### Preselected Elements

For historical reasons, the Document class defines shortcut properties
to access certain kinds of nodes. The `images`, `forms`, and `links`
properties, for example, provide easy access to the `<img>`, `<form>`,
and `<a>` elements (but only `<a>` tags that have an href attribute) of a document.
These properties refer to HTMLCollection objects, which are much like NodeList objects,
but they can additionally be indexed by element ID or name. With the `document.forms`
property, for example, you can access the `<form id="address">` tag as:

```js
document.forms.address;
```

An even more outdated API for selecting elements is the `document.all` property, which
is like an HTMLCollection for all elements in the document. `document.all` is deprecated, and you
should no longer use it.

#### 15.3.2 Document Structure and Traversal

There is a traversal API that allows us to treat a document as a tree
of Element objects, ignoring Text nodes that are also part of the document.
This traversal API does not involve any methods; it is simply a set of properties
on Element objects that allow us to refer to the parent, children, and siblings
of a given element:

**_parentNode_**

**_children_**

- excludes non-Element children like Text nodes and Comment nodes

**_childElementCount_**

- same value as `children.length`

**_firstElementChild,lastElementChild_**

- They are `null` if the Element has no Element children

**_nextElementSibling,previousElementSibling_**

- They are `null` if there is no such sibling

```js
document.children[0].children[1];
document.firstElementChild.firstElementChild.nextElementSibling;
```

Depth-first traversal:

```js
function traverse(e, f) {
  f(e);
  for (let child of e.children) {
    traverse(child, f);
  }
}

function traverse2(e, f) {
  f(e);
  let child = e.firstElementChild;
  while (child != null) {
    traverse2(child, f);
    child = child.nextElementSibling;
  }
}
```

##### Documents as Trees of Nodes

If you want to traverse a document or some portion of a document and **do not
want to ignore the Text nodes**, you can use a different set of properties:

**_parentNode_**

**_childNodes_**

**_firstChild,lastChild_**

**_nextSibling,previousSibling_**

**_nodeType_**

- Element nodes has value '1', Text nodes has value '3', Comment nodes have value '8'.

**_nodeValue_**

- The textual content of a **Text or Comment node**.

**_nodeName_**

- The HTML tag name of an Element, converted to uppercase.
- The Text Node show '#text', and Comment node shows '#comment'

**Note, however, that this API is extremely sensitive to variations in the document
text**. e.g the document is modified by inserting a single newline between
the `<html>` and the `<head>` tag.

Below function returns all of the text within an element or document. However,
**This function is demonstration only -- in practice, you would simply write
`e.textContent` to obtain the textual content of the element e**

```js
function textContent(e) {
  let s = "";
  for (let child = e.childNodes; child !== null; child = child.nextSibling) {
    let type = child.nodeType;
    if (type === 3) {
      s += child.nodeValue;
    } else if (type === 1) {
      s += textContent(child);
    }
  }
  return s;
}
```

#### 15.3.3 Attributes

The Element class defines general `getAttribute()`, `setAttribute()`,
`hasAttribute()`, and **`removeAttribute()`** methods for querying, setting,
testing, and removing the attributes of an element. **But the attribute
values of HTML elements are available as properties of the HTMLElement objects
that represent those elements**, and it is usually much easier to work with them.

##### HTML Attributes as Element Properties

```js
let image = document.querySelector("#main_image");
let url = image.src;
image.id === "main_image"; // => true
```

To set:

```js
let f = document.querySelector("form");
f.action = "https://www.example.com/submit";
f.method = "POST";
```

For some elements, such as `<input>` element, some HTML attribute names map to differently named properties. The HTML `value` attribute of an `<input>`, for example,
is mirrored by the **JavaScript** `defaultValue` property. The JavaScript `value` property
of the `<input>` element contains the user's current input, but changes to the
`value` property do not affect the `defaultValue` property nor the `value` attribute.

```js
input.defaultValue = "123";
input.value; // => '123'
input.value = "456";
input.defaultValue = "123";
input.value; // => '456'
```

**HTML attributes are not case sensitive, but JavaScript property names are.** To
convert an attribute name to the JavaScript property, write it in lowercase.
camelCase if there in more than one word long: `defaultChecked` and `tabIndex`.
Event handlers like `onclick` are an exception.

Some HTML attribute names are reserved words: `class`, `for` (`<label>`), etc.
The general rule is to prefix them with `html`: `htmlFor`. **However, one exception,
which is `className` for `class`.**

```js
label.htmlFor = "aforattribute";
input.className = "aclassname";
```

Some attribute is a boolean (`defaultChecked`) or numeric value (`maxLength`).
Event handler attributes always have functions (or `null`) as their values.

```js
typeof input.defaultChecked; // => "boolean"
typeof input.maxLength; // => "number"
```

In particular, the `delete` operator cannot be used for removing an attribute
from an element. Use `removeAttribute()` instead.

```js
input.removeAttribute("maxLength");
```

##### The `class` Attribute

`class` attribute is poorly named: its value is a list of CSS classes,
not a single class, and it is common in client-side JavaScript programming to want
to add and remove **individual class names** from the list instead of working with
the list as a single string.

For this reason, Element objects define a `classList` property.
The value of the `classList` property is an iterable Array-like object.
Although the name of the property is `classList`, it behaves more like a **set**
of classes, and defines `add()`, `remove()`, `contains()`, and `toggle()` methods

```js
let spinner = document.querySelector("#spinner");
spinner.classList.remove("hidden");
spinner.classList.add("animated");
```

##### Dataset Attribute: `data-`

It is sometimes useful to attach additional information to HTML. In HTML,
any attribute whose name is lowercase and begins with the prefix "data-" is considered
valid. These "dataset" attribute will not affect the presentation of the elements on
which they appear. **They will still be displayed in HTML**.

In the DOM, Element objects **have a `dataset` property that refers to an
object that has properties that correspond to the `data-` attributes**
with their prefix removed. Hyphenated attributes map to camelCase property name:

```html
<h2 id="title" data-section-number="16.1">Attribute</h2>
```

```js
let number = document.querySelector("#title").dataset.sectionNumber;
```

<a id="15.3.4"></a>

#### 15.3.4 Element Content

##### Element Content As HTML

---

> **WARNING**

Never insert user input (e.g. it can be malicious script) into document. If you do this, you allow malicious
users to inject their own script into your application. See Cross-site-scripting
for details.

---

Reading the `innerHTML` property of an Element returns the content of the
element as a string of markup. Setting this property on an
element invokes the web browser's parser and replaces the element's current
content with a parsed representation of the new string:

```js
document.body.innerHTML = "<h1>Oops</h1>";
```

**Important**: Appending text to the `innerHTML` property with the `+=`
operator **is not efficient** because it requires both a serialization step to
convert element content to a string and then parsing step to convert the new
string back into content.

The `outerHTML` property of an Element is like `innerHTML` except
that its value includes the element itself. Thus the new content replaces the
element itself:

```js
let h1 = document.querySelector("h1");
h1.outerHTML = "<h2>I became h2</h2>";
```

A related Element method is `insertAdjacentHTML()`, which
allows you to insert a string of arbitrary HTML markup "adjacent" (`beforebegin`, `afterbegin`, `beforeend`, or `afterend`) to the specified element:

![InsertAdjacentHTML](images/insertAdjacentHTML.png)

```
let h1 = document.querySelector("h1")
h1.insertAdjacentHTML('beforebegin','<span style="color:red">before begin</span>')
h1.insertAdjacentHTML('afterbegin','<span style="color:blue">after begin</span>')
h1.insertAdjacentHTML('beforeend','<span style="color:green">before end</span>')
h1.insertAdjacentHTML('afterend','<span style="color:yellow">after end</span>')
```

##### Element Content As Plain Text

Sometimes you want to query the content of an element:

```html
<h1 id="title">Begin h1 <span>a span</span> h1 End</h1>
<script>
  let h1 = document.querySelector("h1");
  console.log(h1.textContent); // Begin h1 a span h1 End
  h1.textContent = "The span element disappered! "; // now HTML became: <h1>The span element disappered! </h1>
  // Setting textContent won't parse containing tags into element
  h1.textContent = '<span>This is a text, not element</span>'; // now textContent became: <span>This is a text, not element</span>
  h1.textContent // => <span>This is a text, not element</span>
</script>
```

The Element class defines an `innerText` property that is similar to
`textContent`. However, it has some unusual and complex behaviors such
as attempting to preserve table formatting. It is not well specified nor
implemented compatibly between browsers, however, and should no longer be used.

> `text` in `<script>` Elements

Inline `<script>` elements (i.e. **those that do not have `src` attribute**) have
a `text` property that you can use to retrieve their text. The content of a
`<script>` element **is never displayed by the browser, and the HTML parser ignores
angle brackets and ampersands within the script**. This makes a
`<script>` element an ideal place to embed arbitrary textual data for use
by your application. Simply set the `type` attribute of the element to some
value (such as "text/x-custom-data") that makes it clear that the script is not executable
JavaScript code. **If you do this, the JavaScript interpreter will ignore the
script**, but the element will exist in the document tree, and its `text` property will
return the data to you.

```js
<body>
  <script id="ascript" type="text/x-custom-data">
    Just for demo. console.log("this will not be logged because I changed the
    type")
  </script>
  <script>
    let s = document.querySelector("#ascript"); console.log(s.text);
  </script>
</body>
```

Result:

```

    Just for demo.
    console.log("this will not be logged because I changed the type")

```

#### 15.3.5 Creating, Inserting and Deleting Nodes

Create a new element with the `createElement()` method, and
append strings of text or other elements to it with `append()` and `prepend()` method.
These methods add the text inside (contained by) the elements. **These methods defined
on only Element objects, not Text nodes.**

```js
let paragraph = document.createElement("p");
let emphasis = document.createElement("em");
emphasis.append("World"); // Add text to the <em> element
paragraph.append("Hello", emphasis, "!"); // Add text and <em> to <p>
paragraph.prepend("i"); // Add more text at start of <p>
paragraph.outerHTML; // => '<p>iHello<em>World</em>!</p>'
```

`append()` and `prepend()` take any number of arguments, which can be Node objects or
strings. String arguments are automatically converted to Text nodes (Though you can
create it explicitly with `document.createTextNode()`, but there is rarely any
reason to do so).

To insert an Element or Text node into the middle of containing element's child list,
obtain a reference to a sibling node and use `before()` or `after()` method.
These methods defined on both Element and Text nodes.

```js
let greetings = document.querySelector("h2.greetings")
greetings.after(paragraph, document.createElement"hr")) // => <h2 class="greetings">...</h2><<<INSERTED ELEMENTS HERE>>>
```

**Note that elements can only be inserted at one spot in the document.** If
an element is already in the document and you insert it somewhere
else, **it will be moved to the new location, not copied**.

```js
let firstLine = document.querySelector("h1");
let secondLine = document.querySelector("h2");
secondLine.after(firstLine); // firstLine became after secondLine
```

To make a copy of an element, use the `cloneNode()` method, passing `true`
to copy all of its content.

```js
let firstLine = document.querySelector("h1");
let secondLine = document.querySelector("h2");
let thirdLine = firstLine.cloneNode(true);
secondLine.after(thirdLine);
```

You can remove an Element or Text node from the document by calling
its `remove()` method, or you can replace it by calling `replaceWith()` instead.

```js
let firstLine = document.querySelector("h1");
let secondLine = document.querySelector("h2");
let thirdLine = firstLine.cloneNode(true);
secondLine.after(thirdLine);

firstLine.remove();
secondLine.replaceWith(thirdLine);
// Now left only one line
```

The DOM API also defines an older generation of methods for inserting
and removing content. `appendChild()`, `insertBefore()`, `replaceChild()`,
and `removeChild()` are harder to use than the methods shown here and should
never be needed.

#### 15.3.6 Example: Generating a Table of Contents

Example below shows how to dynamically create a table of contents for
a document.

```css
#TOC {
  border: solid black 1px;
  margin: 10px;
  padding: 10px;
}
.TOCEntry { margin: 5px 0px;}
.TOCEntry a { text:decoration: none; }
.TOCLevel1 { font-size: 16pt; font-weight: bold; }
.TOCLevel2 { font-size: 14pt; margin-left: .25in; }
.TOCLevel3 { font-size: 12pt; margin-left: .5in; }
.TOCSectNum:after { content: ": "; }
```

```js
document.addEventListener("DOMContentLoaded", () => {
  let toc = document.querySelector("#TOC");
  if (!toc) {
    toc = document.createElement("div");
    toc.id = "TOC";
    document.body.prepend(toc);
  }

  let headings = document.querySelectorAll("h2,h3,h4,h5,h6");

  let sectionNumbers = [0, 0, 0, 0, 0];

  for (let heading of headings) {
    if (heading.parentNode === toc) {
      continue;
    }

    let level = parseInt(heading.tagName.charAt(1)) - 1;
    sectionNumbers[level - 1]++;
    for (let i = level; i < sectionNumbers.length; i++) {
      sectionNumbers[i] = 0;
    }

    let sectionNumber = sectionNumbers.slice(0, level).join(".");

    let span = document.createElement("span");
    span.className = "TOCSectNum";
    span.textContent = sectionNumber;
    heading.prepend(span);

    let anchor = document.createElement("a");
    let fragmentName = `TOC${sectionNumber}`;
    anchor.name = fragmentName;

    heading.before(anchor); // Insert anchor before heading
    anchor.append(heading); // Move heading inside anchor

    let link = document.createElement("a");
    link.href = `#${fragmentName}`;

    link.innerHTML = heading.innerHTML;
    let entry = document.createElement("div");
    entry.classList.add("TOCEntry", `TOCLevel${level}`);
    entry.append(link);
    toc.append(entry);
  }
});
```

### 15.4 Scripting CSS

JavaScript can also control the visual appearance and layout of those documents
by scripting CSS.

CSS styles that are commonly scripted from JavaScript:

- Setting the `display` style to "none" hides an element. You can
  later show the element by setting `display` to some other value.

  - ```js
    e.style.display = "none";
    ```

  - ```css
    .hidden {
      display: none;
    }
    ```

    ```js
    document.querySelector("#tooltip").classList.remove("hidden");
    document.querySelector("#tooltip").classList.add("hidden");
    ```

- You can dynamically position elements by setting the `position` style to
  "absolute", "relative", or "fixed", and then setting the `top` and `left` styles
  to the desired coordinates. This is important when
  using JavaScript to display dynamic content like model dialogues and tool tips.

- You can shift, scale, and rotate elements with the `transform` style.

  ```js
  let circle = document.querySelector(".circle");
  circle.style.transform = "translateY(10%)";
  ```

- You can animate changes to other CSS styles with the `transition` style.
  These animations are handled automatically by the web browser and do not
  require JavaScript, but you can use JavaScript to initiate the animations.

  ```js
  let box = document.querySelector(".box");
  box.classList.add(".transparent");
  ```

#### 15.4.1 CSS Classes

The simplest way to use JavaScript to affect the styling of document is modifying
class names, from the `class` attribute or `classList` property of Element object explained previously.

```css
.hidden {
  display: none;
}
```

```js
document.querySelector("#tooltip").classList.remove("hidden");
document.querySelector("#tooltip").classList.add("hidden");
```

#### 15.4.2 Inline Styles

DOM defines a `style` property on all Element objects that correspond to
the `style` attribute. **(ONLY INLINE PROPERTIES)** It is a CSSStyleDeclaration object:
a parsed representation of the CSS styles that appear in textual form
in the `style` attribute.

```html
<style>
.black {
  color: black;
}

</style>
<a id="element1" class="black">hohohoho</a>
<script>
  console.log(document.querySelector(".black").style.color)
</script>
```

Result: empty string because it is not inline style

To display and set the position of our hypothetical tooltip with JavaScript:

```js
function displayAt(tooltip, x, y) {
  tooltip.style.display = "block";
  tooltip.style.position = "absolute";
  tooltip.style.left = `${x}px`;
  tooltip.style.top = `${y}px`;
}
```

**Important**: Note that the semicolons go outside the strings.

---

> **Naming Conventions: CSS Properties In JavaScript**

Many CSS style properties such as `font-size`, contains hyphen in their names.
Some names of the properties of the CSSStyleDeclaration Object are slightly
different from the names of actual CSS properties. The names with hyphens to be
converted to camelCase. e.g. `border-left-width` to `borderLeftWidth`, `font-family`
to `fontFamily`

---

Furthermore, remember that many CSS properties require units such as "px"
for pixels or "pt" for points. Thus it is not correct to set the marginLeft
without unit.

```js
e.style.marginLeft = 300; // Incorrect!
e.style.marginLeft = "300"; // Incorrect!
e.style.marginLeft = "300px"; // ok
```

You can make some computation:

```js
e.style.left = `${x0 + left_border + left_padding}px`;
e.style.margin = `${top}px ${right}px ${botton}px ${left}px`;
```

You can query whole inline style by using `getAttribute()` or access `cssText` property
of the CSSStyleDeclaration object:

```js
f.setAttribute("style", e.getAttribute("style"));
// or
f.style.cssText = e.style.cssText;
```

Keep in mind that querying the `style` property of an element represents only the
inline styles of an element. In general, you want to query the styles of an element, you probably
want the _computed style_, which is discussed next.

#### 15.4.3 Computed Styles

The computed style for an element is the set of property values that **the browser
derives from the element's inline style plus all applicable style rules in all
stylesheets**. Like inline styles, computed styles are represented with a CSSStyleDeclaration object.
Unlike inline styles, however, **there are read-only**.

Obtain the computed style for an element with the `getComputedStyle()` method
of the Window object. The second argument is used to specify a [CSS pseudoelement](https://www.w3schools.com/css/css_pseudo_elements.asp) such
as "::before" or "::after"

```js
let title = document.querySelector("#section1title");
let styles = window.getComputedStyle(title);
let beforeStyles = window.getComputedStyle(title, "::before");
```

The return value of `getComputedStyle()` is a CSSStyleDeclaration object. The
important differences between CSSStyleDeclaration object that
represents inline styles and one that represents computed styles are:

- Computed style properties are read-only.

- Computed style properties are _absolute_: relative units like percentages and points
  are converted to absolute values. **This value will be a string with a "px" suffix**, you won't have to worry about parsing or converting other units. Properties whose
  values are colors will be returned in `rgb()` or `rgba()` format instead of something like `black` even though they were set like that.

- Shortcut properties are not computed e.g. `margin`. Use `marginTop`, `marginLeft` and so on instead. (they can still be query, but the result might be incorrect, for example, if we set it to duplicate value : `margin: 50px 50px`, we got only `50px` by querying `margin`)

- The `cssText` property of the computed style is undefined.

A CSSStyleDeclaration object returned by `getComputedStyle()` generally contains
much more information about an element than the CSSStyleDeclaration obtained from the
inline `style` property of that element. But they can be tricky, and querying them does
not always provide the information you might expect. Consider the `font-family` attribute, when we query it, we're simply getting the value of the most specific `font-family`
style that applies to the element, it does not tell you which typeface is **actually** in use. Similarly, if an element is not absolutely positioned, attempting to query its
position and size through the `top` and `left` properties of its computed style often
returns the value `auto`.

Querying the computed style of an element is not the preferred way to determine
the element's **size and position**. See [15.5.2](#15.5.2) for a simpler, portable
alternative.

#### 15.4.4 Scripting Stylesheets

Stylesheets are associated with an HTML document with a `<style>` tag
or with a `<link rel="stylesheet>"` tag, you can give them both `id` attributes
and then look them up with `document.querySelector()`.

There is a `disabled` property to disable the entire stylesheet.

```js
function toggleTheme() {
  let lightTheme = document.querySelector("#light-theme");
  let darkTheme = document.querySelector("#dark-theme");
  if (darkTheme.disabled) {
    lightTheme.disabled = true;
    darkTheme.disabled = false;
  } else {
    lightTheme.disabled = false;
    darkTheme.disabled = true;
  }
}
```

Another simple way to script stylesheets is to replace it.

```js
function setTheme(name) {
  let link = document.createElement("link");
  link.id = "theme";
  link.rel = "stylesheet";
  link.href = `themes/${name}.css`;

  let currentTheme = document.querySelector("#theme");
  if (currentTheme) {
    currentTheme.replaceWith(link);
  } else {
    document.head.append(link);
  }
}
```

Less subtly, you can also just insert a string of HTML containing a `<style>` tag
into your document. This is a fun trick:

```js
document.head.insertAdjacentHTML(
  "beforeend",
  "<style>body{transform: rotate(180deg)}</style>",
);
```

Browsers define an API that allows JavaScript to look inside
stylesheets to query, modify, insert, and delete style rules in that stylesheet.
This API is so specialized that it is not documented here. You can read about
it on MDN by searching for ["CSSStyleSheet"](https://developer.mozilla.org/en-US/docs/Web/API/CSSStyleSheet) and "CSS Object Model"

#### 15.4.5 CSS Animations And Events

We can play some animations:

```css
.transparent {
  opacity: 0;
}

.fadeable {
  transition: opacity 0.5s ease-in;
}
```

```js
<div id="subscribe" class="fadeable notification">it is notification</div>
<script>
document.querySelector("#subscribe").classList.add("transparent")
</script>
```

JavaScript can be used to monitor the progress of a CSS transition
because the web browser fire events at the start and end of a transition.
The "transitionrun" event is dispatched when the transition is first triggered.
This may happen before any visual changes begin, when the `transition-delay` style
has been specified. Once the visual changes begin a "transitionstart" event is dispatched,
and when the animation is complete, a "transitionend" event is dispatched.
The event object passed to handlers for these events is a TransitionEvent, which
has a `propertyName` property that specifies the CSS property being animated, and
an `elapsedTime` property that specifies how many seconds have passed since the "transitionstart" event.

```js
let sub = document.querySelector("#subscribe");
sub.addEventListener("transitionrun", (e) => {
  console.log("transitionrun");
  console.log(e.propertyName);
  console.log(e.elapsedTime);
});
sub.addEventListener("transitionstart", (e) => {
  console.log("transitionstart");
  console.log(e.propertyName);
  console.log(e.elapsedTime);
});
sub.addEventListener("transitionend", (e) => {
  console.log("transitionend");
  console.log(e.propertyName);
  console.log(e.elapsedTime);
});
```

In addition to transitions, CSS also supports a more complex form of
animation known simply as "CSS Animations". These use CSS properties such
as `animation-name` and `animation-duration` and a special `@keyframes` rule
to define animation details. Details of how CSS animations work are beyond the
scope of this book. Once again, you can use JavaScript to trigger the animation
simply by adding the class to the element that is to be animated.

HTML:

```html
<div class="box"></div>
```

CSS:

```css
.box {
  width: 100px;
  height: 100px;
  background-color: red;
}

.slide {
  animation-name: slide;
  animation-duration: 2s;
  animation-iteration-count: infinite;
}

@keyframes slide {
  0% {
    transform: translateX(0);
  }
  50% {
    transform: translateX(200px);
    background-color: blue;
  }
  100% {
    transform: translateX(0);
  }
}
```

JavaScript:

```js
setTimeout(() => {
  let box = document.querySelector("box");
  box.classList.add("slide");
}, 2000);
```

![Animation1](images/animation1.png)

![ANimation2](images/Animation2.png)

![Animation1](images/animation1.png)

And like CSS transitions, CSS animations also trigger events that your JavaScript
code can listen form. `animationstart`, `animationend`. If the animation
repeats more than once, then an `animationiteration` event is dispatched
after each repetition except the last. The event object passed to handler
functions is an AnimationEvent object, which include an `animationName` property,
and `elapsedTime` property.

```js
box.addEventListener("animationstart", (e) => {
  console.log(e.type);
  console.log(e.animationName);
});
box.addEventListener("animationend", (e) => {
  console.log(e.type);
  console.log(e.animationName);
});
box.addEventListener("animationiteration", (e) => {
  console.log(e.type);
  console.log(e.animationName);
});
```

### 15.5 Document Geometry and Scrolling

The following subsections explain how you can go back and forth
between the abstract, tree-based **_model_** of a document, and the
geometrical, coordinate-based **_view_** of the document as it is laid out
in a browser window.

#### 15.5.1 Document Coordinates and Viewport Coordinates

The position of a document element is measured in CSS pixels, with
the `x` coordinate increasing to the right and the `y` coordinate increasing
as we go down. There are two different points we can use as the coordinate system
origin, however: **the x and y coordinates of an element can be relative to the
top-left corner of the document** or **relative to the top-left corner of the
_viewport_ in which the document is displayed**. In top-level windows and
tabs, the "viewport" is the portion of the browser that **actually displays document
content**: it excludes browser "chrome" such as menus, toolbars, and tabs.
For documents displayed in `<iframe>` tags, it is the iframe element in the DOM that
defines the viewport for the nested document.

**We must be clear whether we are using document coordinates or viewport coordinates when
we talk about the position of an element** (Viewport coordinate are sometimes called "window coordinates")

**Important**: Note that window.innerWidth (viewport width) includes also the scrollbar. That is, the width of viewport has no dependence of scrollbar.

In general, to convert between the two coordinate systems, we must add
or subtract the _scroll offsets_. For example, if an element has a _y_ coordinate
of 200pixels in document coordinates, if the user has scrolled down by 75 pixels,
then that element has a _y_ coordinate of 125 pixels in viewport coordinates.

It is logical to assume that every element in a document must have a unique position
in document coordinates.
However in general, document coordinates don't really work on the web. The problem is that
the CSS `overflow` property allows elements within a document to contain more
content than it can display. Elements can have their own scrollbars and serve
as viewports for the content they contain. The fact that the web allows scrolling
elements within a scrolling document means that it is simply not possible to describe the
position of an element within the document using a single (x,y) point.

![Double Scroll Bar](images/double-scroll-bar.png)

Because document coordinates don't really work, client-side JavaScript tends to use
viewport coordinates. The `getBoundingClientRect()` and `elementFromPoint()` methods
described next use view coordinates, for example, and the `clientX` and `clientY` properties
of mouse and pointer event object also use this coordinate system.

```js
document.addEventListener("mousemove", (e) => {
  console.log(e.clientY);
});
```

---

**Please understand the below paragraph well!**

When you explicitly **position** an element using CSS `position:fixed`,
the `top` and `left` properties are interpreted in viewport coordinates. If
you use `position:relative`, the element is positioned relative to where it
would have been if it didn't have the `position` property set. If you use
`position:absolute`, then `top` and `left` are relative to the document or
to the nearest containing positioned element. **This means, for example, that
an absolutely positioned element inside a relatively positioned element is positioned
relative to the container element, not relative to the overall document.**

It is sometimes very useful to create a relatively positioned container with
`top` and `left` set to 0 (so the container is laid out normally) in order to
establish a new coordinate system origin for the absolutely positioned elements it contains.
We might refer to this new coordinate system as "container coordinates" to distinguish
it from document and viewpoint coordinates.

---

---

> **CSS Pixels**

Today's 4K monitors and "retina" displays have such high resolution that software
pixels have been **decoupled from hardware pixels**. A CSS pixel -- and
therefore a client-side JavaScript pixel -- may in fact consist of multiple device
pixels. The **`devicePixelRatio` property of the Window object** specifies how many device
pixels are used for each software pixel. A "dpr" of 2, for example, means that each
software pixel is actually a 2 x 2 grid of hardware pixels. The `devicePixelRatio` value
depends on the physical resolution of your hardware, on settings in your operating system,
and on the zoom level in your browser.

`devicePixelRatio` does not have to be an integer. If you are using a CSS font size
of "12px" and the device pixel ratio is 2.5, then the actual font size, in device
pixels, is 30.

Because the pixel values we use in CSS no longer correspond directly to individual pixels
on the screen, pixel coordinates no longer need to be integer. If the `devicePixelRatio`
is 3, then a coordinate of 3.33 make perfect sense (3 * 3.33 = 10). And if the ratio
is actually 2, then a coordinate of 3.33 will just be rounded up to 3.5 (2*3.5 =7).
This because the coordinate needs
to be adjusted to align with the available pixels on the screen.

---

<a id="15.5.2"></a>

#### 15.5.2 Querying the Geometry of an Element

You can determine the size (including CSS border and padding but not margin) and
postition (in view coordinates) of an element by calling its `getBoundingClientRect()`
method. It returns an object with properties `left`, `right`, `top`, `bottom`, `width`,
and `height`.

Block elements, such as images, paragraphs, and `<div>` elements are always
rectangular when laid out by the browser. Inline elements, such as `<span>`, `<code>`
, `<em>` and `<b>` elements, however, may span multiple lines and may therefore consists
of multiple rectangles. If you want to query the individual rectangles of inline elements,
call the `getClientRects()` method to obtain a read-only, array-like object
whose elements are rectangles objects like those returned by `getBoundingClientRect()`
If you call `getBoundingRect()` on this element, the bounding rectangle would include
the entire width of all lines.

#### 15.5.3 Determining the Element at a Point

The `getBoundingClientRect()` method allow us to determine the current position of an element in a **viewport**. Sometime we want to go into the other direction and
determine which element is at given location in the viewport. You can determine this with
the `elementFromPoint()` method of the Document object. Call this
method with the `x` and `y` coordinates of a point (**viewport coordinates**: the
`clientX` and `clientY` coordinates of a mouse event work, for example).

```js
let box = document.querySelector(".box");
box.hello = function () {
  alert("HELLO!");
};
document.addEventListener("mousemove", (e) => {
  let el = document.elementFromPoint(e.clientX, e.clientY);
  console.log(el?.hello?.());
});
```

The hit detection algorithm for selecting the element is not precisely specified,
but the intent of this method is that it returns the
innermost (most deeply nested) and the uppermost (highest CSS z-index attribute) element
at that point.

#### 15.5.4 Scrolling

The `scrollTo()` method of the Window object takes the _x_ and _y_ coordinates
of a point (**document coordinates**) and sets these as the scrollbar offsets.
That is, it scrolls the window so that the specified point **is in the upper-left corner
of a viewport**. The following code scrolls the browser so that bottom-most page of the document
is visible:

```js
let documentHeight = document.documentElement.offsetHeight;
let viewportHeight = window.innerHeight;
window.scrollTo(0, documentHeight - viewportHeight);
```

The `scrollBy()` method of the Window is similar to `scrollTo()`, but its arguments
are relative and are added to the current scroll position:

```js
setInterval(() => {
  window.scrollBy(0, 50);
}, 500);
```

**If you want to scroll smoothly with `scrollTo()` or `scrollBy()`, pass
a single object argument instead of two numbers, like this:**

```js
window.scrollTo({
  left: 0,
  top: documentHeight - viewportHeight,
  behavior: "smooth",
});

window.scrollBy({
  left: 0,
  top: 50,
  behavior: "smooth",
});
```

Often, instead of scrolling to a numeric location in a document, we can scroll
to certain element in the document with `scrollIntoView()` method on the desired
HTML element. By default, it tries to put the top edge of the element at or near the
top of the viewport. If `false` is passed as the only argument, **it tries to put the
bottom edge of the element at the bottom of the viewport**. You can set the behaviour
to smooth as well.

```js
let el = document.querySelector(".box");
el.scrollIntoView(false);
el.scrollIntoView({
  behavior: "smooth",
  block: "end", // put the bottom edge of the element at the bottom of the viewport
  inline: "end", // how it should be positioned horizontally if horizontal scrolling is needed.
});
```

The values for `block`, and `inline` are: `start`, `end`, `nearest`, and `center`

#### 15.5.5 Viewport Size, Content Size, and Scroll Position

For browser windows, the viewport size is given by the `window.innerWidth` and `window.innerHeight` properties. (Web pages optimized for mobile devices often use a `<meta name="viewport">` tag in their
`<head>` **to set the desired viewport width for the page**). The total size of the
document is the same as the size of `<html>` element,
`document.documentElement`. You can call `getBoundingClientRect()` on `document.documentElement` to get the
width and height of the document, or you can use the `offsetWidth` and `offsetHeight`
properties of `document.documentElement` (The width is already excluded the scrollbar).
The scroll offsets of the document within its viewport are available as `window.scrollX`
and `window.scrollY`. These are **read-only properties**.

Things are a little more complicated for elements. Every element
object defines the following three groups of properties:

```
offsetWidth     clientWidth      scrollWidth
offsetHeight    clientHeight     scrollHeight
offsetLeft      clientLeft       scrollLeft
offsetTop       clientTop        scrollTop
offsetParent
```

The `offsetWidth` and `offsetHeight` properties of an element return its on-screen
size in CSS pixels. The returned sizes include the element border and padding but not
margins. The `offsetLeft` and `offsetTop` return the _x_ and _y_ coordinates of the
element. For many elements, these values are **document coordinates**. But for
descendants of positioned elements and for some other elements, such as
table cells, these properties return coordinates that are **relative to an ancestor
element**. The `offsetParent` property specifies which element the properties
are relative to. These offset properties are all read-only.

`clientWidth` and `clientHeight` are like `offsetWidth` and `offsetHeight` except
that do not include the border size, and inline element just returns 0 for these properties. `clientLeft` and `clientTop` are not very useful:
they return the horizontal and vertical distance between the outside of an element's
padding and the outside of its border (the width of the border). These
client properties are all read-only. Inline element just returns
0 for these properties.

`scrollWidth` and `scrollHeight` return the size of an element's content area
plus its padding plus any overflowing content. When the content fits within
the content area without overflow, these properties are same as `clientWidth`
and `clientHeight`. But when there is overflow, they include the overflowing. `scrollTop`
and `scrollLeft` give the scroll offset of the element content **within** the element's
viewport. Unlike all the other, `scrollTop` and `scrollLeft` is writable properties,
and you can set them to **scroll the content within an element.** (e.g. it is applicable, for example, when you have a container that set overflow to "scroll")

### 15.6 Web Components

There is a lot of work to do every time you want to display a search box
in a web application, and most web applications today are not written
using "raw" HTML. Instead, many web developers use frameworks like React
and Angular that support the creation of reusable user interface components
like the search box shown here. **Web components is a browser-native alternative
to those frameworks** based on three relatively recent additions to web standards that
allow JavaScript to extend HTML with **new tags** that work as self-contained, reusable
UI components.

#### 15.6.1 Using Web Components

Web components are defined in JavaScript, so in order to use a web
component in your HTML file, you need to include the JavaScript file
that defines the component. Because web components are a relatively new technology,
they are often written as JavaScript modules, so you might include one in your HTML
like this:

```html
<script type="module" src="components/search-box.js"></script>
```

Web components define their own HTML tag names, with the
important restriction that those tag names **must include a hyphen**. (This
means that future version of HTML can introduce new tags without hyphens, and
there is no chance that the tags will conflict with anyone's web component.)

```html
<search-box placeholder="Search..."></search-box>
```

Web components can have attributes just like regular HTML tags can;
the documentation for the component you are using should tell you which
attributes are supported. Web components cannot be defined with self-closing tags.

Some web components are written to expect children. Some web components are written
so they can optionally
accept **specially labeled children** that will appear in named "slot".

```html
<search-box>
  <img src="images/search-icon.png" slot="left" />
  <img src="images/cancel-icon.png" slot="right" />
</search-box>
```

The `slot` attribute is an extension to HTML that it is used to specify which
children should go where. The slot name -- "left" and "right" in above example
are defined by the web component. -- It should be included in its documentation
if your web component supports slot.

Note that web components are often implemented as JavaScript modules and can be
loaded into HTML files with a `<script type="module>` tag (default `deferred`).
If a web component has children, then those children will probably be displayed
incorrectly before component is defined. You can use this CSS to keep web
components hidden until they are defined:

```css
search-box:not(:defined) {
  opacity: 0;
  display: inline-block;
  width: 300px;
  height: 50px;
}
```

---

> **DocumentFragment Nodes**

The DOM API organizes a document into a tree of Node objects, where a
Node can be a Document, an Element, a Text node, or even a Comment.
None of these node types allows you to represent a fragment of a document
that consists of a set of sibling node without their parent. This
is where DocumentFragment comes in: it is another type of Node that **serves
as a temporary parent when you want to manipulate a group of sibling nodes
as a single unit**. Create this node with `document.createDocumentFragment()`.

Once you have a DocumentFragment, you can use it like an Element and `append()`
content to it. **A DocumentFragment does not have a parent, when you insert
it into the document, it itself is not inserted. Instead, all of its children
are inserted**

---

#### 15.6.2 HTML Templates

The HTML `<template>` tag is only loosely related to web components, but it does enable
a useful optimization for components that appear frequently in web pages. `<template>`
tags and their children **are never rendered by a web browser** and are only useful
on web pages that use JavaScript.
The idea behind this tag is that when a web pages contains multiple repetitions
of the same basic HTML structure (such as rows in a table or the internal
implementation of a web component), then we can use a `<template>` to define
that element structure once, then **use JavaScript to duplicate the structure
as many times as needed**.

In JavaScript, a `<template>` tag is represented by an HTMLTemplateElement object.
This object defines a single `content` property, and the value of this property is a
**DocumentFragment of all the child nodes**. You can clone this
DocumentFragment and then insert the cloned copy into your document as needed.

Suppose you're working with a document that includes a `<table>` and `<template id="row">`
tag and that the template defines the structure of rows for that table. You might use the
template like this:

```js
let tableBody = document.querySelector("tbody");
let template = document.querySelector("#row");
let clone = template.content.cloneNode(true);
tableBody.append(clone);
```

Template elements do not have to appear literally in an HTML document in order to
be useful. You can create a template in your JavaScript code, create its children
with `innerHTML`, and then make as many clones as needed without the parsing overhead
of `innerHTML`. This is how HTML templates are typically used in web components,
and [SearchBox Example](#15.6.5) demonstrates this technique.

#### 15.6.3 Custom Elements

Custom elements **enable a JavaScript class with an HTML tag name so that any
such tags in the document are automatically turned into instances of the
class in the DOM tree**.

**Important**: The name **must** include a hyphen.

The `customElements.define()` method takes tag name as its first argument, and a
subclass of HTMLElement as its second argument. Any existing elements
in the document with that tag name are "upgraded" to newly
created instances of the class. And if the browser parses any HTML in the future,
it will automatically create an instance of the class for each of the
tags it encounters.

```js
customElements.define(
  "search-box",
  class SearchBox extends HTMLElement {
    connectedCallback() {
      // initialization
    }

    static get observedAttributes() {
      return []; // returns attributes names
    }

    attributeChangedCallback(name, oldValue, newValue) {
      switch (name) {
        case "":
      }
    }
  },
);
```

The second argument should extend HTMLElement ant not a more specific type like
HTMLButtonElement (It is not supported in Safari and a different syntax is required
to use a custom element that extends anything other than HTMLElement).

The browser will automatically invoke certain "lifecycle methods" of a custom
element class. The `connectedCallback()` method is invoked when an instance
of the custom element is inserted into the document, and many elements use this
method to **perform initialization**. `disconnectedCallback()` method is invoked when
(and if) the element is removed from the document though this is less often used.

If a custom element class defines a static `observedAttributes` property
whose value is an array of attribute names, and if any of the
named attributes are set (or changed) on an instance of the custom element,
the browser will invoke the `attributeChangedCallback()` method, passing
the attribute name, its old value, and its new value. **This callback can
take whatever steps are necessary to update the component based on its attribute
values**. **You must use `setAttribute()` to trigger it, directly setting the property
will not trigger this callback.**

As an example of a custom element, suppose we want to be able to display
circles within paragraphs of regular text. We'd like to write HTML like this:

```html
<p>
  the document has one marble: <inline-circle></inline-circle>
  The HTML parser instatiates two more marbles:
  <inline-circle diameter="1.2em" color="blue"></inline-circle>
  <inline-circle diameter=".6em" color="gold"></inline-circle>
  How many marbles does the document contain now?
</p>
```

```js
customElements.define(
  "inline-cicle",
  class InlineCircle extends HTMLElement {
    connectedCallback() {
      console.log("connected");
      this.style.display = "inline-block";
      this.style.borderRadius = "50%";
      this.style.border = "solid black 1px";
      this.style.transform = "translateY(10%)";
      // set a default size
      if (!this.style.width) {
        this.style.width = "0.8em";
        this.style.height = "0.8em";
      }
    }

    disconnectedCallback() {
      console.log("removed");
    }

    static get observedAttributes() {
      return ["diameter", "color"];
    }

    attributeChangedCallback(name, oldValue, newValue) {
      switch (name) {
        case "diameter":
          this.style.width = newValue;
          this.style.height = newValue;
          break;
        case "color":
          this.style.backgroundColor = newValue;
          break;
      }
    }

    get diameter() {
      return this.getAttribute("diameter");
    }
    set diameter(diameter) {
      return this.setAttribute("diameter", diameter);
    }
    get color() {
      return this.getAttribute("color");
    }
    set color(color) {
      return this.setAttribute("color", color);
    }
  },
);
```

![Inline Circle](images/inline-circle.png)

Note that above example is not well encapsulated, continue reading next section
for more detail.

#### 15.6.4 Shadow DOM

The custom element demonstrated above is not well encapsulated. When
you set its `diameter` or `color` attributes, it responds by altering its own `style`
attribute, which is not behavior we would ever expect from a real HTML Element.
**To turn a custom element into a true web component, it should use the powerful encapsulation
mechanism known as shadow DOM**.

**Tip**: We can encapsulates the circle in the shadow root, thus not accessible by
DOM API.

Shadow DOM allows a "shadow root" to be attached to a custom element (and also
to a `<div>`, `<span>`, `<body>`, `<article>`, `<main>`, `<nav>`, `<header>`,
`<footer>`, `<section>`, `<p>`, `<blockquote>`, `<aside>`, or `<h1>` through `<h6>` element) known as a "shadow host". A shadow root is the root of another (shadow host), more
private, tree of descendant elements that sprouts from the shadow host and can
be though of as a distinct minidocument.

The word "shadow" refers to the fact that elements that descend from a shadow root
**are hiding in the shadow**: not part of the normal DOM tree, do not appear in the
`children` array of their host element, and are not visited by normal DOM traversal mehods
such as `querySelector()`. For contrast, the normal, regular DOM children of a shadow host
are sometimes referred to as the "light DOM."

To understand the purpose of the shadow DOM, picture the HTML `<audio>` and `<video>`
elements: they display a nontrivial user interface for controlling media playback, but the
play and pause buttons and other UI elements are not part of the DOM tree and
cannot be manipulated by JavaScript.

##### Shadow DOM Encapsulation

**The key feature of shadow DOM is the encapsulation it provides**.

- hidden from regular DOM methods like `querySelectorAll()`

  - When a shadow root is created and attached to its shadow host, it can be
    created in "open" or "closed" mode.

  - More commonly, though, shadow roots are created in "open" mode, means that the
    shadow host has a `shadowRoot` property that JavaScript can use to gain access to the
    elements of the shadow root.

- Styles defined beneath a shadow root are private to that tree and will never affect
  the light DOM elements on the outside. (A shadow root can define default styles
  for its host element, but these will be overridden by light DOM styles)

  - Similarly, the light DOM styles that apply to the shadow host element
    have no effect on the descendants of the shadow root.

  - Element in the shadow DOM will inherit things like font size and background color
    from the light DOM, and styles in the shadow DOM can choose to use **CSS variables
    defined in the light DOM**.

  - For the most part, however, the styles of the light DOM and the styles of the shadow DOM
    are completely independent: the author of web component and the user of a
    web component do not have to worry about collisions or conflicts between their stylesheets.

- Some events (like "load") that occur within the shadow DOM are confined to the
  shadow DOM. Others, including focus, mouse, and keyboard events bubble up and out.

  - When it begins to propagate in the light DOM, its `target` property is changed
    to the shadow host element, so it appears to have originated directly on that element.

##### Shadow DOM Slots and Light DOM Children

You may be wondering how two distinct content trees can be displayed within
the same host element. Here's how it works:

- The descendants of the shadow root are always displayed within the shadow host.

- **If those descendants include a `<slot>` element, then the regular light DOM children
  of the host element are displayed as if they were children of that `<slot>`, replacing
  any shadow DOM content in the slot**. If the shadow DOM does not include a `<slot>`,
  then any light DOM content of the host **is never displayed**. If the shadow DOM
  has a `<slot>`, but the shadow host has no light DOM children, then the shadow
  DOM content of the slot is displayed as a default.

- When light DOM content is displayed within a shadow DOM slot, we say that those elements
  have been "distributed", but it is important to understand that the elements do not actually
  become part of the shadow DOM. They can still be queried with `querySelector()`.

- If the shadow DOM defines more than one `<slot>` and names those slots with a `name`
  attribute then the shadow host can specify which slot they would like to appear in
  by specifying a `slot="slotname"` attribute.

##### Shadow DOM API

The Shadow DOM doesn't have much of a JavaScript API. To turn a light DOM
element into a shadow host, just call its `attachShadow()` method, passing
`{mode: "open"}` as the only argument. This returns a shadow root object. The shadow root object is a DocumentFragment, and you can use DOM methods to add content to it
or just set its `innerHTML` property to a string of HTML.

If your web component needs to know when the light DOM content of a shadow DOM
`<slot>` has changed, it can register a listener for "slotchange" events
directly on the `<slot>` element. **The change of attributes does not fire the event,
this is triggered when element is remove or added (e.g. append)**

<a id="15.6.5"></a>

#### 15.6.5 Example: a `<search-box>` Web Component

![Search Box Element](images/search-box-element.png)

This example shows how to use the low-level web component APIs directly.
In practice, many web components developed today create them using higher-level
libraries such as "lit-element". One of the reasons to use a library is that
creating reusable and customizable components is actually quite hard to do well,
and there are many details to get right.

```html
<search-box>
  <img src="images/search-icon.png" slot="left" />
  <img src="images/cancel-icon.png" slot="right" />
</search-box>
<script>
  let search = document.querySelector("search-box");
  search.addEventListener("search", () => {
    console.log("do something");
  });
</script>
```

```js
class SearchBox extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({ mode: "open" });
    this.shadowRoot.append(SearchBox.template.content.cloneNode(true));

    this.input = this.shadowRoot.querySelector("#input");
    let leftSlot = this.shadowRoot.querySelector("slot[name='left']");
    let rightSlot = this.shadowRoot.querySelector("slot[name='right']");

    this.input.addEventListener("focus", () => {
      this.setAttribute("focused", "");
    });
    this.input.addEventListener("blur", () => {
      this.removeAttribute("focused", "");
    });
    this.input.addEventListener("search", () => {
      console.log("hohohoho");
    });

    leftSlot.addEventListener("click", (event) => {
      event.stopPropagation();
      if (this.disabled) return;
      this.dispatchEvent(new CustomEvent("search"), {
        detail: this.input.value,
      });
    });

    rightSlot.addEventListener("click", (event) => {
      event.stopPropagation();
      if (this.disabled) return;
      let e = new CustomEvent("clear", { cancelable: true });
      this.dispatchEvent(e);
      if (!e.defaultPrevented) {
        // If the event was not cancelled
        this.input.value = "";
      }
    });
  }

  attributeChangedCallback(name, oldValue, newValue) {
    if (name === "disabled") {
      this.input.disabled = newValue !== null;
    } else if (name === "placeholder") {
      this.input.placeholder = newValue;
    } else if (name === "size") {
      this.input.size = newValue;
    } else if (name === "value") {
      this.input.value = newValue;
    }
  }

  get placeholder() {
    return this.getAttribute("placeholder");
  }
  set placeholder(value) {
    this.setAttribute("placeholder", value);
  }
  get size() {
    return this.getAttribute("size");
  }
  set size(value) {
    this.setAttribute("size", value);
  }
  get value() {
    return this.getAttribute("value");
  }
  set value(value) {
    this.setAttribute("value", value);
  }
  get disabled() {
    return this.getAttribute("disabled");
  }
  set disabled(value) {
    if (value) this.setAttribute("disabled", "");
    else this.removeAttribute("disabled");
  }
  get hidden() {
    return this.getAttribute("hidden");
  }
  set hidden(value) {
    if (value) this.setAttribute("hidden", "");
    else this.removeAttribute("hidden");
  }
}

SearchBox.observedAttributes = ["disabled", "placeholder", "size", "value"];

SearchBox.template = document.createElement("template");
SearchBox.template.innerHTML = `
<style>
/*
 * The :host selector refers to the <search-box> element in the light DOM.
 * These styles are defaults and can be overridden by the user of the
 * <search-box> with styles in the light DOM.
 */
:host {
  display: inline-block;
  border: solid black 1px;
  border-radius: 5px;
  padding: 4px 6px;
}
:host([hidden]) { /* Note the parentheses: when host has hidden... */
  display:none;
}
:host([disabled]) { /* When host has the disabled attribute */
  opacity: 0.5;
}
:host([focused]) { /* When host has the focused attribute */
  box-shadow: 0 0 2px 2px #6AE;
}

/* The rest of the stylesheet only applies to elements in the Shadow DOM */
input {
  border-width: 0;
  outline: none; /* hide the focus ring */
  font: inherit; /* input elements don't inherit font by default */
  background: inherit: /* ditto */
}

slot {
  cursor: default; /* arrow pointer cursor */
  user-select: none; /* don't let user select the emoji text */
}
</style>
<div>
  <slot name="left">\u{1f50d}</slot>
  <input type="text" id="input" />
  <slot name="right">\u{2573}</slot>
</div>
`;

customElements.define("search-box", SearchBox);
```

Inline Circle Enhancement:

```js
class InlineCircle extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({ mode: "open" });
    this.shadowRoot.append(InlineCircle.template.content.cloneNode(true));
  }

  static get observedAttributes() {
    return ["diameter", "color", "hoho"];
  }

  attributeChangedCallback(name, oldValue, newValue) {
    switch (name) {
      case "diameter":
        this.style.width = newValue;
        this.style.height = newValue;
        break;
      case "color":
        this.style.backgroundColor = newValue;
        break;
    }
  }

  get diameter() {
    return this.getAttribute("diameter");
  }
  set diameter(diameter) {
    return this.setAttribute("diameter", diameter);
  }
  get color() {
    return this.getAttribute("color");
  }
  set color(color) {
    return this.setAttribute("color", color);
  }
}
InlineCircle.template = document.createElement("template");
InlineCircle.template.innerHTML = `
<style>
  :host {
    display : inline-block;
    border-radius : 50%;
    border : solid black 1px;
    transform : translateY(10%);
    width : 0.8em;
    height : 0.8em;
  }
</style>
`;

customElements.define("inline-circle", InlineCircle);
```

### 15.7 SVG: Scalable Vector Graphics

There are three ways you can use SVG in web browsers:

1. Use _.svg_ image files with regular HTML `<img>` tags.

2. Embed SVG tags directly into your HTML document. (Because the XML-based
   SVG format is so similar to HTML)

3. You can use the DOM API to dynamically create SVG elements to generate images
   on demand.

In addition to simple shape-drawing primitives, it includes support
for arbitrary curves, text and animation. However, the goal of this section
is just to show you how you can use SVG in your HTML documents and script it
with JavaScript.

#### 15.7.1 SVG in HTML

If you embed SVG directly in HTML, you can even use CSS stylesheets to specify
things like fonts, colors, and line widths:

```css
#clock {
  stroke: black;
  stroke-linecap: round; /* rounded ends */
  fill: #ffe; /* on an off-white background */
}
#clock .face {
  stroke-width: 3;
}
#clock .ticks {
  stroke-width: 2;
}
#clock .hands {
  stroke-width: 3;
}
#clock .numbers {
  font-family: sans-serif;
  font-size: 10px;
  font-weight: bold;
  text-anchor: middle;
  stroke: none;
  fill: black;
}
```

HTML:

```html
<body>
  <svg id="clock" viewBox="0 0 100 100" width="250" height="250">
    <circle class="face" cx="50" cy="50" r="45" />
    <g class="ticks">
      <line x1="50.00" y1="5.000" x2="50.00" y2="10.00" />
      <line x1="72.50" y1="11.03" x2="70.00" y2="15.36" />
      <line x1="88.97" y1="27.50" x2="84.64" y2="30.00" />
      <line x1="95.00" y1="50.00" x2="90.00" y2="50.00" />
      <line x1="88.97" y1="72.50" x2="84.64" y2="70.00" />
      <line x1="72.50" y1="88.97" x2="70.00" y2="84.64" />
      <line x1="50.00" y1="95.00" x2="50.00" y2="90.00" />
      <line x1="27.50" y1="88.97" x2="30.00" y2="84.64" />
      <line x1="11.03" y1="72.50" x2="15.36" y2="70.00" />
      <line x1="5.000" y1="50.00" x2="10.00" y2="50.00" />
      <line x1="11.03" y1="27.50" x2="15.36" y2="30.00" />
      <line x1="27.50" y1="11.03" x2="30.00" y2="15.36" />
    </g>
    <g class="numbers">
      <text x="50" y="18">12</text>
      <text x="85" y="53">3</text>
      <text x="50" y="88">6</text>
      <text x="15" y="53">9</text>
    </g>
    <g class="hands">
      <line class="hourhand" x1="50" y1="50" x2="50" y2="25" />
      <line class="minutehand" x1="50" y1="50" x2="50" y2="20" />
    </g>
  </svg>
  <script src="clock.js"></script>
</body>
```

![Clock](images/clock.png)

There are many other SVG tags, however, you'll need to consult an SVG reference
to learn more.

You may also notice that the stylesheet is odd. Styles like `fill`, `stroke-width`,
and `text-anchor` are not normal CSS style properties. In this case,
CSS is essentially being used to set attributes of SVG tags that appear in the
document. **Note that the CSS `font` shorthand property does not work
for SVG tags**, you must explicitly set `font-family`, `font-size`, and `font-weight`
as separate style properties

#### 15.7.2 Scripting SVG

One reason to embed SVG directly into your HTML file instead of `<img>` tag
is that we can use DOM API to manipulate the SVG image.

Below is the _clock.js_ to compute the current time and reflect it on the clock.

```js
(function updateClock() {
  let now = new Date();
  let sec = now.getSeconds();
  let min = now.getMinutes() + sec / 60;
  let hour = (now.getHours() % 12) + min / 60;
  let minangle = min * 6; // 6 degrees per minutes;
  let hourangle = hour * 30; // 30 degrees per hours;
  let minhand = document.querySelector("#clock.minutehand");
  let hourhand = document.querySelector("#clock.hourhand");

  minhand.setAttribute("transform", `rotate(${minangle},50,50)`);
  hourhand.setAttribute("transform", `rotate(${hourangle},50,50)`);

  setTimeout(updateClock, 10000);
})();
```

![Clock With Js](images/clock-with-js.png)

#### 15.7.3 Creating SVG Images with JavaScript

If you want to create SVG element with the JavaScript DOM API, you can't use
the normal `createElement()` function. Instead you must use `createElementNS()`,
which takes an XML namespace string as its first argument. For SVG,
that namespace is the literal string "http://www.w3.org/2000/svg".

![Slice](images/slice.png)

The element used to display each slice is `<path>`. **This SVG element
describes arbitrary shapes comprised of lines and curves.** The shape description
is specified by the `d` attribute. The value of this attribute uses a compact
grammar of letter codes and numbers that specify coordinates, angles, and other
values.

`M` -- "move to", and is followed by _x_ and _y_ coordinates.

`L` -- "line to", draws a line from the current point to the coordinates that follows it

`A` -- "arc", draw an arc, and is followed by seven number describing the arc.

`Z` -- Close path back

**Note**: The computation is in radian, not degree

```html
<div id="chart"></div>
<script src="./scripts/piechart.js"></script>
<script>
  document.querySelector("#chart").append(
    pieChart({
      width: 640,
      height: 400,
      cx: 200,
      cy: 200,
      r: 180,
      lx: 400,
      ly: 10,
      data: {
        JavaScript: 71.5,
        Java: 45.4,
        "Bash/Shell": 40.4,
        Python: 37.9,
        "C#": 35.3,
        PHP: 31.4,
        "C++": 24.6,
        C: 22.1,
        TypeScript: 18.3,
        Ruby: 10.3,
        Swift: 8.3,
        "Objective-C": 7.3,
        Go: 7.2,
      },
    }),
  );
</script>
```

```js
function pieChart(options) {
  let { width, height, cx, cy, r, lx, ly, data } = options;
  let svg = "http://www.w3.org/2000/svg";

  let chart = document.createElementNS(svg, "svg");
  chart.setAttribute("width", width);
  chart.setAttribute("height", height);
  chart.setAttribute("viewBox", `0 0 ${width} ${height}`);

  chart.setAttribute("font-family", "sans-serif");
  chart.setAttribute("font-size", "18px");

  let labels = Object.keys(data);
  let values = Object.values(data);
  let total = values.reduce((x, y) => x + y);

  let angles = [0];
  values.forEach((x, i) => angles.push(angles[i] + (x / total) * 2 * Math.PI));
  values.forEach((value, i) => {
    let x1 = cx + r * Math.sin(angles[i]);
    let y1 = cy - r * Math.cos(angles[i]);
    let x2 = cx + r * Math.sin(angles[i + 1]);
    let y2 = cy - r * Math.cos(angles[i + 1]);

    // This is a flag for angles larger than a half circle
    // It is required by the SVG arc drawing component
    let big = angles[i + 1] - angles[i] > Math.PI ? 1 : 0;

    let path =
      `M${cx},${cy}` + `L${x1},${y1}` + `A${r},${r} 0 ${big} 1 ${x2},${y2} Z`;

    // Compute the CSS color for this slice. This formula works for only
    // about 15 colors. So don't include more than 15 slices in a chart.
    let color = `hsl(${(i * 40) % 360}, ${90 - 3 * i}%, ${50 + 2 * i}%)`;

    let slice = document.createElementNS(svg, "path");
    slice.setAttribute("d", path);
    slice.setAttribute("fill", color);
    slice.setAttribute("stroke-width", "1");
    slice.setAttribute("stroke", "black");
    chart.append(slice);

    let icon = document.createElementNS(svg, "rect");
    icon.setAttribute("x", lx);
    icon.setAttribute("y", ly + 30 * i);
    icon.setAttribute("width", "20");
    icon.setAttribute("height", "20");
    icon.setAttribute("fill", color);
    icon.setAttribute("stroke", "black");
    icon.setAttribute("stroke-width", "1");
    chart.append(icon);

    let label = document.createElementNS(svg, "text");
    label.setAttribute("x", lx + 30);
    label.setAttribute("y", ly + 30 * i + 16);
    label.append(`${labels[i]} ${value}`);
    chart.append(label);
  });

  return chart;
}
```

### 15.8 Graphics in a `<canvas>`

The `<canvas>` element has no appearance of its own but creates a
drawing surface within the document and exposes a powerful drawing API
to client-side JavaScript.

The main difference between the `<canvas>` API and SVG is that with
the canvas you create drawings by calling methods, and with SVG you create
drawings by building a tree of XML elements. Each has its strengths and weakness.
An SVG drawing, for example, is easily edited by removing elements from its description.
To remove an element from the same graphic in a `<canvas>`, it is often necessary to
erase the drawing and redraw it from scratch.

---

> **3D Graphics in a Canvas**

You can also call `getContext()` with the string `webgl` to obtain a context
object that allows you to draw 3D graphics using WebGL API. WebGL
is a large, complicated, and low-level API that allows
JavaScript Programmers to access the GPU, write custom shaders, and perform
other very powerful graphics operations. WebGL is not documented in this book,
however, web developers are more likely to use utility libraries built on top
of WebGL than to use the WebGL API directly.

---

Most of the Canvas drawing API is defined not on the `<canvas>` element itself,
but instead on a **"drawing context"** object obtained with the `getContext()` method
of the canvas. Call `getContext()` with the argument "2d" to obtain a CanvasRenderingContext2D object
that you can use to draw two-dimensional graphics into the canvas.

A simple example to display two simple shapes:

```html
<p>
  This is a red square: <canvas id="square" width="10" height="10"></canvas>.
</p>
<p>
  This is a blue circle: <canvas id="circle" width="10" height="10"></canvas>.
</p>
<script>
  let canvas = document.querySelector("#square");
  let context = canvas.getContext("2d");
  context.fillStyle = "#f00"; // red
  context.fillRect(0, 0, 10, 10);

  canvas = document.querySelector("#circle");
  context = canvas.getContext("2d");
  context.beginPath();
  context.arc(5, 5, 5, 0, 2 * Math.PI, true);
  context.fillStyle = "#00f"; // blue
  context.fill();
</script>
```

![Canvas1](images/canvas1.png)

#### 15.8.1 Paths and Polygons

To draw lines on a canvas and to fill the areas enclosed by those lines,
you **begin by defining a _path_**. A path is a **sequence** of one or more subpaths.
A subpath is a **sequence of two or more points connected by line (or curve) segments**.

Begin a new path -- `beginPath()`

Begin a new subpath -- `moveTo()`

Connect that point to new point -- `lineTo()`

```js
let canvas = document.querySelector("#demo");
let c = canvas.getContext("2d");
c.beginPath();
c.moveTo(100, 100);
c.lineTo(200, 200);
c.lineTo(100, 200);
c.fillStyle = "#ccc";
c.strokeStyle = "#008";
c.lineWidth = 5;
c.fill();
c.stroke();
```

![Canvas2](images/canvas2.png)

To connect back to starting point, use `closePath()` (You could also call `lineTo(100,100)`,
but then you end up with three line segments that share a start and end point but are not
truly closed. When drawing with the wide lines, the visual results are better if you use
`closePath()`)

Case with `lineTo(100,100)`

![Canvas Open](images/canvas-open.png)

Case `closePath()`

![Canvas Closed](images/canvas-closed.png)

There are two important points to notice about `stroke()` and `fill()`.
First, **both methods operate on all subpaths in the current path**. Suppose
we had added another subpath in the preceding code:

```js
let canvas = document.querySelector("#demo");
let c = canvas.getContext("2d");
c.beginPath();
c.moveTo(100, 100);
c.lineTo(200, 200);
c.lineTo(100, 200);

c.moveTo(300, 100); // new lines
c.lineTo(200, 200); // new lines

c.fillStyle = "#ccc";
c.strokeStyle = "#008";
c.lineWidth = 5;
c.fill();
c.stroke();
```

![2 Subpaths](images/2-subpaths.png)

The second point is that **neither one alters the current path**:
you can call `fill()` and the path will sill be there when you call `stroke()`.
**When you are done with a path and want to begin another, you must remember
to call `beginPath()`**. If you don't, you'll end up adding new subpaths to the existing
path, and you may end up drawing those old subpaths over and over again.

Polygons example:

```js
function polygon(c, n, x, y, r, angle = 0, counterclockwise = false) {
  c.moveTo(x + r * Math.sin(angle), y - r * Math.cos(angle));
  let delta = (2 * Math.PI) / n;
  for (let i = 1; i < n; i++) {
    angle += counterclockwise ? -delta : delta;
    c.lineTo(x + r * Math.sin(angle), y - r * Math.cos(angle));
  }
  c.closePath();
}

let c = document.querySelector("canvas").getContext("2d");

c.beginPath();
polygon(c, 3, 50, 70, 50); // Triagle
polygon(c, 4, 150, 60, 50, Math.PI / 4); // Square
polygon(c, 5, 255, 55, 50); // Pentagon
polygon(c, 6, 365, 53, 50, Math.PI / 6); // Hexagon
polygon(c, 4, 365, 53, 20, Math.PI / 4, true); // Small square inside the hexagon

c.fillStyle = "#ccc";
c.strokeStyle = "#008";
c.lineWidth = 5;

c.fill();
c.stroke();
```

![Polygons](images/polygons.png)

Note that this example draws a hexagon with a square inside it. The square and the hexagon are separate subpaths,
but they overlap. When this happens (or when a single
subpath intersects itself), the canvas needs to be
able to determine which regions are inside the path and which are outside.
The canvas uses a test known as the "nonzero winding rule" to achieve this. **In
this case, the interior of the square is not filled** because the square and the hexagon were
drawn in the opposite directions: The vertices of the hexagon were connected with
line segments moving clockwise around the circle; The vertices of the square were connected counter clockwise. Had the square been drawn clockwise as well, the call to `fill()`
would have filled the interior of the square as well.

![Both Clockwise](images/both-clockwise.png)

Another example of single subpath intersects itself:

```js
c.beginPath();
c.moveTo(100, 100);
c.lineTo(25, 25);
c.lineTo(25, 100);
c.lineTo(100, 25);
c.fillStyle = "#ccc";
c.strokeStyle = "#800";
c.lineWidth = 5;
c.fill();
c.stroke();
```

![Single Subpath Intersect](images/single-subpath-intersect.png)

#### 15.8.2 Canvas Dimensions and Coordinates

The default canvas coordinate system places the origin(0,0) at the upper-left
corner of the canvas. Points on the canvas can be specified using floating-point values.

The dimensions of a canvas cannot be altered without completely resetting the canvas.
**Setting either `width` or `height` properties of a Canvas (even setting them to their
current value) clears the canvas.**

The `width` and `height` attribute also specify the default size (**in CSS pixels**)
at which the canvas will be displayed on the screen. If `window.devicePixelRatio` is 2,
then 100x100 CSS pixels is actually 40000 hardware pixels. When the contents
of the canvas are drawn onto the screen, the 10,000 pixels in memory will need to be
enlarged to cover 40,000 physical pixels on the screen, and this means that your
graphics will not be as crisp as they could be.

For optimum image quality, **you should not use the `width` and `height` attributes
to set the on-screen size of the canvas. Instead, set the desired on-screen
size CSS pixel size of the canvas with CSS `width` and `height` style attributes.**
Then, before you begin drawing in your JavaScript code, set the `width` and `height`
properties of the canvas object to the number of CSS pixels times
`window.devicePixelRatio`.

HTML:

```html
<canvas id="myCanvas"></canvas>
```

CSS:

```css
#myCanvas {
  width: 300px;
  height: 200px;
}
```

JavaScript:

```javascript
const canvas = document.getElementById("myCanvas");
const context = canvas.getContext("2d");
// Set the canvas width and height properties
canvas.width = canvas.clientWidth * window.devicePixelRatio;
canvas.height = canvas.clientHeight * window.devicePixelRatio;
// Begin drawing on the canvas
context.fillStyle = "red";
context.fillRect(0, 0, canvas.width, canvas.height);
```

#### 15.8.3 Graphics Attributes

`fillStyle`, `strokeStyle`, and `lineWidth` properties are graphics
attributes that specify the color to be used by `fill()` and by `stroke()`.

If you define a method that draws a shape and do not set these properties
yourself, the caller of your method can define the color of the shape by setting
the `strokeStyle` and `fillStyle` properties before calling your method.
This separation of graphics state from drawing commands is fundamental to the
Canvas API and is akin to the separation of presentation from context achived
by applying CSS stylesheets to HTML documents.

##### Line Styles

`lineWidth` -- how wide (in CSS pixels) the lines drawn by `stroke()` will be. The
default is 1. The line are centered over the path. If you're stroking a closed path
and only want the line to appear outside the path. `stroke()` first then only `fill()`
If you only want the line to appear inside a closed path, call the `save()` and
`clip()` methods first, then call `stroke()` and `restore()`.

Below shows difference: LHS is `stroke()` first then `fill()`; middle is `fill()`, followed by `stroke()`; RHS is `save()`, `clip()`, `stroke()`, then `restore()`

![Sequences Of Line Styles](images/sequences-of-line-styles.png)

The `lineCap` and `lineJoin` properties can have a significant impact on the
visual appearance of the ends of a path and the vertices at which two
path segments meet.

`lineCap` -- `butt` (default), `square`, `round`

![Linecap](images/linecap.png)

`lineJoin` -- `miter` (defalt), `round`, `bevel`

![Linejoin](images/linejoing.png)

**Tip**: If two lines meet at a very narrow angle, then the resulting miter can
become quit long and visually distracting. If the miter at a given vertex would be
longer than half of the line width times the `miterLimit` property, that vertex
will be drawn with a beveled join instead of a mitered join. The default value for `miterLimit` is 10.

To draw dashed and dotted lines, unlike other line-drawing properties, it is
set and queried with the methods `setLineDash()` and `getLineDash()`
instead of with a property. `lineDashOffset` property specifies how fat into
the dash pattern drawing should begin (default is 0).

```js
// First
c.setLineDash([18, 3, 3, 3]); // 18px dash, 3px space, 3px dot, 3px space

// Second
c.lineDashOffset = 21;
c.setLineDash([18, 3, 3, 3]); // 18px dash, 3px space, 3px dot, 3px space
```

![Dash Line](images/dash-line.png)

##### Colors, Patterns, and Gradients

The `fillStyle` and `strokeStyle` properties specify how paths
are filled and stroked. **The word "style" often means color, but these
properties can also be used to specify a color gradient or an image to be used
for filling and stroking**.

To set a solid color, simply set these properties to a valid CSS color string.

To fill (or stroke) with a color gradient, set `fillStyle` (or `strokeStyle`)
to a CanvasGradient object returned by the `createLinearGradient()` or `createRadialGradient()` methods of the context.

The arguments of `createLinearGradient()` method: coordinates of two points that
define a line along which the colors will vary.

The arguments of `createRadialGradient()` method: centers and radii of two circles.
Area inside smaller circle or outside the larger will be filled with solid colors.

After creating these object, you must define the gradient colors by calling
the `addColorStop()` method of the CanvasGradient. The first argument to this
method is a number between 0.0 and 1.0. The secon argument is a CSS color specification.
You must call this method at least twice to define a simple color gradient. The color
at 0.0 will appear at the start of the gradient, the color at 1.0 will appear at the end.

```js
let canvas = document.querySelector("canvas");
canvas.width = canvas.clientWidth * window.devicePixelRatio;
canvas.height = canvas.clientHeight * window.devicePixelRatio;
let w = canvas.width;
let h = canvas.height;

let c = canvas.getContext("2d");
let bgfade = c.createLinearGradient(0, 0, w, h);
bgfade.addColorStop(0, "#88f"); // light blue
bgfade.addColorStop(1, "#fff"); // white
c.fillStyle = bgfade;
c.fillRect(0, 0, w, h);

c.beginPath();
c.moveTo(w / 2, h / 2);
c.arc(w / 2, h / 2, w / 2, 0, 2 * Math.PI, true);
let donut = c.createRadialGradient(w / 2, h / 2, 0, w / 2, h / 2, w / 2);
donut.addColorStop(0.0, "transparent");
donut.addColorStop(0.7, "rgba(100,100,100,.9)"); // gray
donut.addColorStop(1.0, "rgba(0,0,0,0)"); // transparent again
c.fillStyle = donut;
c.fill();
```

![Gradient Canvas](images/gradient-canvas.png)

You can also fill and stroke using images. To do this, set `fillStyle` or
`strokeStyle` to a CanvasPattern returned by the createPattern() method of the
context object. The first argument to this method should be an `<img>`
or `<canvas>` element that contains the image you wan to fill or stroke with.
The second argument is either `repeat`, `repeat-x`, `repeat-y`, `no-repeat`.

```js
let img = document.createElement("canvas");
img.width = 300;
img.height = 300;
let imgc = img.getContext("2d");
let bgimgc = imgc.createLinearGradient(0, 0, 300, 300);
bgimgc.addColorStop(0, "rgba(0,0,0,0)");
bgimgc.addColorStop(1.0, "rgba(100,0,0,1)");
imgc.fillStyle = bgimgc;
imgc.fillRect(0, 0, 300, 300);

let canvas = document.querySelector("canvas");
canvas.width = 1200;
canvas.height = 1200;
let c = canvas.getContext("2d");
let pat = c.createPattern(img, "repeat");
c.fillStyle = pat;
c.fillRect(0, 0, 1200, 1200);
```

![Canvas Image](images/canvas-image.png)

##### Text Styles

The `font` property specifies the font to be used by the text-drawing method
`fillText()` and `strokeText()`. The value
of the `font` property should be a string in the same syntax as the CSS
`font` attribute.

The `textAlign` property specifies how the text should be horizontally aligned with
respect to the X coordinate passed to `fillText()` or `strokeText()`.
Legal values: `start` (default), `left`, `center`, `right`, `end`, for left-to-right text,
`start` has the same meaning as `left`

```js
c.textAlign = align;
c.fillText("hello", x, y);
```

![Textalign](images/textalign.png)

The `textBaseline` property specifies how the text should be vertically
aligned with respect to the y coordinate. The default value is `alphabetic`, and
it is appropriate for Latin and similar scripts. The value "ideographic" is
intended for use with scripts such as Chinese and Japanese. The
value `hanging` is intended for use with Devanagari and similar scripts.
The "top", "middle", "bottom" baselines are purely geometric baselines, based
on the "em square" of the font.

```js
c.textBaseline = baseline;
c.fillText("hello", x, y);
```

![Text Baseline](images/text-baseline.png)

![Text Baseline Chinese](images/text-baseline-chinese.png)

##### Shadows

Four properties of the context object control the drawing of drop
shadows.

`shadowColor` -- color of the shadow. The default is fully transparent black. Patterns
and gradients are not allowed.

`shadowOffsetX` and `shadowOffsetY` -- X and Y offsets of the shadow. The default is 0.
These values are not affected by coordinate transformations (15.8.5)(#15.8.5):
shadow direction and "height" remain consistent even when shapes are rotated and scaled.

`shadowBlur` -- how blurred the edges of the shadow are. The default is 0.

```js
let canvas = document.querySelector("canvas");
canvas.width = 1200;
canvas.height = 1200;
let c = canvas.getContext("2d");
c.fillStyle = "#008";
c.shadowBlur = 50;
c.shadowColor = "black";
c.shadowOffsetX = 50;
c.shadowOffsetY = 50;
c.fillRect(200, 200, 800, 800);
```

![Shadow](images/shadow.png)

##### Translucency and Compositing

A way to use a translucent color is by using `rgba()`. But the
CanvasAPI provides another way to work with translucent colors.
If you do not want to explicitly specify an alpha channel for each color, or
if you want to add translucency to opaque images or patterns, you can
set the `globalAlpha` pattern. **Every pixel you draw will have its alpha value
multiplied by `globalAlpha`, the default is 1, which adds no transparency.**

The process of combining new (possibly translucent) source pixels with
existing (possibly translucent) destination pixels is called compositing.

You can set the `globalCompositeOperation` property to specify other way of
combining pixels. The default value is "source-over": source pixels
are drawn over the destination pixels and are combined with them if
the source is translucent. "destination-over" will cause combining pixels
as if the new source pixels were drawn beneath the existing destination pixels.

[See here for excellent demonstration](https://www.w3schools.com/jsref/canvas_globalcompositeoperation.asp)

##### Saving and Restoring Graphics State

Each `<canvas>` element has only **a single context object**, and every
call to `getContext()` returns the same CanvasRenderingContext2D object. So you
can't call `getContext()` multiple times to obtain multiple context object and define
different attributes on each context.

However, Canvas API allow you to save the current graphics state so that you can
alter it and then easily restore it later. The `save()` method pushes
the current graphics state onto a **stack** of saved states. The `restore()` method
pops the stack.

All of the properties that have been described in this section are part of
the saved state.
Importantly, the currently defined path and the current point are not part
of the graphics state and cannot be saved and restored.

#### 15.8.4 Canvas Drawing Operations

Other drawing methods:

##### Rectangles

CanvasRenderingContext2D defines four methods for drawing rectangles.

`fillRect()` fills the specified rectangle with the `strokeRect()`
strokes the outline of the specified rectangle using the current `strokeStyle` and
other line attributes.

`clearRect()` is like `fillRect()`, but it ignores the current fill styles
and fills the rectangle with transparent black pixels (the default color of
all blank canvases).

These three methods **do not affect the current path or the current point within that path.**

`rect()` does affect the current path: it adds the specified rectangle, in
a subpath of its own, to the path. Like other path-definition methods, it does not fill
or stroke anything itself until explicitly call `fill()` or `stroke()`

##### Curves

The CanvasRenderingContext2D object defines a number of methods that add
a new point to the subpath and connect the current point point to that new point
with a curve:

`arc()` -- adds a circle, or a portion of a circle (an arc), to the path.
The arc to be drawn is specified with six parameters: x, y coordinates of the center
of the arc, the radius of the circle, the start and end angles of the arc, and the direction (clockwise or counterclockwise) of the arc between those two angles. If there is
a current point in the path, then this method connects the current point to the beginning
of arc with a straight line, then connects the beginning of the arc to the end
of the arc with a portion of a circle, leaving the end of the arc as the new current point.
If there is no current point when this method is called,then it only adds the
circular arc to the path.

```js
c.arc(200, 200, 200, 0, Math.PI);
c.stroke();

c.beginPath();
c.moveTo(0, 600);
c.arc(300, 750, 150, (Math.PI / 2) * 3, Math.PI / 2);
c.stroke();
```

![Arc](images/arc.png)

`ellipse()` -- Much like `arc()` except that it adds an ellipse or a portion
of an ellipse to the path. Instead of one radius, it has two: an x-axis radius
and a y-axis radius. It also takes another argument to specifies the number of radians
by which the ellipse is rotated clockwise about its center.

```js
c.ellipse(200, 200, 200, 100, Math.PI / 4, 0, Math.PI, true);
c.stroke();
```

![Ellipse](images/ellipse.png)

`arcTo()` -- Draws a straight line and a circular arc just like `arc()` does, but
it specifies the arc to be drawn using different parameters. The arguments to `arcTo()`
specify points P1 and P2 and a radius.

```js
c.moveTo(200, 20);
c.arcTo(200, 100, 100, 20, 40);
c.stroke();

c.beginPath();
c.fillStyle = "blue";
c.arc(200, 20, 3, 0, Math.PI * 2);
c.arc(200, 100, 3, 0, Math.PI * 2);
c.fill();
c.beginPath();
c.arc(100, 20, 3, 0, Math.PI * 2);
c.fill();
```

![ArcTo](images/arcTo.png)

`bezierCurveTo()` -- Adds a new point P to the subpath and connects it to the
current point with a cubic Bezier curve. The shape of the curve is specified
by two "control points", C1 and C2. A the start of the curve
(at the curren point), the curve heads in the direction of C1.
At the end of the curve (at point P), the curve arrives from the direction
of C2. In between these points, the direction of the curve varies smoothly.
The point P becomes the new current point for the subpath.

```js
c.moveTo(200, 200); // red
// green blue yellow
c.bezierCurveTo(250, 100, 350, 220, 200, 300);
c.stroke();
```

![Baziercurve](images/baziercurve.png)

`quadraticCurveTo()` -- This method is like `bezierCurveTo()`, but it uses
a quadratic Bezier curve instead of a cubic Bezier curve and has only a single
control point.

```js
c.moveTo(100, 100);
c.quadraticCurveTo(200, 50, 300, 100);
c.quadraticCurveTo(400, 150, 500, 100);
c.stroke();
```

![Quadraticcurve](images/quadraticcurve.png)

Example:

```js
function rads(x) {
  return (Math.PI * x) / 180;
}

let c = document.querySelector("canvas").getContext("2d");

c.fillStyle = "#aaa";
c.lineWidth = 2;

c.beginPath();
c.arc(75, 100, 50, 0, rads(360), false);
c.fill();
c.stroke();

c.beginPath();
c.ellipse(200, 100, 50, 35, rads(15), 0, rads(360));

c.moveTo(325, 100);
c.arc(325, 100, 50, rads(-60), rads(0), true);
c.closePath();

c.moveTo(340, 92);
c.arc(340, 92, 42, rads(-60), rads(0), false);
c.closePath();

c.moveTo(450, 50);
c.arcTo(500, 50, 500, 150, 30);
c.arcTo(500, 150, 400, 150, 20);
c.arcTo(400, 150, 400, 50, 10);
c.arcTo(400, 50, 500, 50, 0);
c.closePath();

c.moveTo(525, 125);
c.quadraticCurveTo(550, 75, 625, 125);
c.fillRect(550 - 3, 75 - 3, 6, 6);

c.moveTo(625, 100);
c.bezierCurveTo(645, 70, 705, 130, 725, 100);
c.fillRect(645 - 3, 70 - 3, 6, 6);
c.fillRect(705 - 3, 130 - 3, 6, 6);

c.fill();
c.stroke();
```

![Curves](images/curves.png)

##### Text

To draw text in canvas, you normally use the `fillText()` method,
which draw text using the color (or gradient or pattern) specified by
the `fillStyle` property. For special effects at large text sizes, you can
use `strokeText()` to draw the outline of the individual font glyphs.
Both methods take the text to be drawn as their first argument, and take
the _x_ and _y_ coordinates of the text as the second and third arguments.
Neither method affects the current path or the current point.

These methods takes an optional fourth argument. If given, this argunment
specifies the **maximum width of the text to be displayed**. If the text would be
wider than the specified value, the canvas will make it fit by scaling it or by
using a narrower or smaller font.

```js
c.font = "bold 100px serif";
c.fillText("hello", 200, 200, 100);
```

![Narrowed Font](images/narrowed-font.png)

If you need to measure text yourself before drawing it, pass it to the
`measureText()` method. This method returns a TextMetrics object that specifies
the measurements of the text when drawn with the current font. (At the time the author
of this writing, the only "metric" contained in the `TextMetrics` object
is the width. Query the on-screen width of a string like this)

```js
let width = c.measureText("hello world").width;
```

This is useful if you want to center a string of text within a canvas, for example.
We can calculate it and get the coordinate we need to place the text on.

##### Images

The `drawImage()` method **copies the pixels** of a source image (or of a rectangle
within the source image) onto the canvas, scaling the rotating the pixels of the image as necessary.

This method can be invoked with three, five or nine arguments.
In all cases, the first argument is the source image. This image argument is often
an `<img>` element, but it can be also another `<canvas>` element or even a `<video>`
element (from which a single frame will be copied). If you specify an `<img>` or `<video>` element is still loading its data, the `drawImage()` call will do nothing.

In the three-argument, the **entire source image is copied to the canvas**, the _x_ and
_y_ coordinates are interpreted in the current coordinate system:

```js
img.onload = () => {
  c.drawImage(img, 0, 0);
};
```

The five-arguments version of `drawImage()` adds `width` and `height` arguments
to the _x_ and _y_ arguments described earlier.

```js
img.onload = () => {
  c.drawImage(img, 0, 0, 100, 100);
};
```

The nine-argument version of `drawImage()` specifies both a source rectangle and
a destination rectangle and copies only the pixels **within the source rectangle**.
2-5 arguments specify the source rectangle. They are measured in CSS pixels. 6-9 arguments
specify the destination rectangle into which the image is drawn and are in the current
coordinate system of the canvas, not in the default coordinate system.

```js
// drawImage(image: CanvasImageSource, sx: number, sy: number, sw: number, sh: number, dx: number, dy: number, dw: number, dh: number): void;
img.onload = () => {
  c.drawImage(
    img,
    img.width / 2 - 50,
    img.height / 2 - 50,
    100,
    100,
    0,
    0,
    100,
    100,
  );
};
```

In addition to drawing images into a canvas, we can also extract the content
of a canvas as an image using the **canvas's** `toDataURL()` method. It returns
the content of the canvas as a PNG image, encoded as a string using a `data: URL`.
You can make a static snapshot of a canvas with code like this:

```js
let img = document.createElement("img");
img.src = canvas.toDataURL();
document.body.appendChild(img);
```

Result: `<img src="data:image/png;base64,iVBOR…A+xhKkCoMq7cAAAAASUVORK5CYII=">`

<a id="15.8.5"></a>

#### 15.8.5 Coordinate System Transforms

In the default coordinate system of canvas, the coordinates of a point map directly
to a CSS pixel (which then maps directly to one or more device pixels). In addition
to the default coordinate system, however, every canvas has a "current transformation
matrix" as part of its graphics state. This matrix defines the current coordinate system
of the canvas. In most canvas operations, it is taken to be point of
the specified point instead of the default coordinate system. The current transformation matrix is used to **convert
the coordinates you specified** to the equivalent coordinates in the default system.

The `setTransform()` method allows you to set a canvas's transformation matrix directly,
but coordinate system transformations are usually easier to specify as **a sequence
of translations, rotations, and scaling** operations.

Below are the examples:

![Canvas Transform](images/canvas-transform.png)

```js
function shear(c, kx, ky) {
  c.transform(1, ky, kx, 1, 0, 0);
}

c.save();
c.translate(100, 100);
c.fillText("translate(100,100)", 10, 16);
c.drawImage(graph, 0, 0);
c.restore();

c.save();
c.translate(400, 100);
c.rotate(Math.PI / 6);
c.fillText("translate(400,100)", 10, 16);
c.fillText("rotate(Math.PI/6)", 10, 32);
c.drawImage(graph, 0, 0);
c.restore();

c.save();
c.translate(100, 400);
c.scale(1.5, 0.8); // x, y
c.fillText("translate(100,400)", 10, 16);
c.fillText("c.scale(1.5, 0.8)", 10, 32);
c.drawImage(graph, 0, 0);
c.restore();

c.save();
c.translate(400, 400);
shear(c, 0.5, 0);
c.fillText("translate(400,400)", 10, 16);
c.fillText("shear(c, 0.5, 0)", 10, 32);
c.drawImage(graph, 0, 0);
c.restore();

c.save();
c.translate(100, 700);
shear(c, 0, -0.25);
c.fillText("translate(100,700)", 10, 16);
c.fillText("shear(c, 0, -0.25)", 10, 32);
c.drawImage(graph, 0, 0);
c.restore();

c.save();
c.translate(100, 1000);
c.fillText("translate(100,1000)", 10, 16);
c.fillText("c.scale(1, -1)", 10, 32);
c.scale(1, -1);
c.drawImage(graph, 0, 0);
c.restore();

c.save();
c.translate(400, 1000);
c.fillText("translate(400,1000)", 10, 16);
c.fillText("c.rotate(-Math.PI)", 10, 32);
c.rotate(-Math.PI);
c.drawImage(graph, 0, 0);
c.restore();
```

Note that a negative scale factor to the `scale()` method **flips** that axis across
the origin.

##### Understanding Transformations Mathematically

The transformations supported by the canvas are known as _affine transforms_. Affine
transforms may modify the distances between points and the angles between them,
but parallel lines always remain parallel after an affine transformation.
An arbitrary affine transform can be described by the six parameters:

```bash
x' = ax + cy + e
y' = bx + dy + f
```

You can apply an arbitrary transformation to the current coordinate
system by passing those six parameters to the `transform()` method.

```js
function shear(c, kx, ky) {
  c.transform(1, ky, kx, 1, 0, 0);
}
// transform(a: number, b: number, c: number, d: number, e: number, f: number): void;
```

The `setTransform()` method takes the same arguments as `transform()`, but
instead of transforming the current coordinate system, it ignores the current system,
**transform the default coordinate system**. **`setTransform()` is useful to
temporarily reset the canvas to its default coordinate system**:

```js
c.save();
c.setTransform(1, 0, 0, 1, 0, 0);
c.restore();
```

##### Transformation Example

Koch snowflakes with 0, 1, 2, 3, 4 levels of recursion

![Koch Snowflakes](images/koch-snowflakes.png)

```js
let deg = Math.PI / 180;
function snowflake(c, n, x, y, len) {
  c.save();
  c.translate(x, y);
  c.moveTo(0, 0);
  leg(n);
  c.rotate(-120 * deg);
  leg(n);
  c.rotate(-120 * deg);
  leg(n);
  c.closePath();
  c.restore();

  function leg(n) {
    c.save();
    if (n === 0) {
      c.lineTo(len, 0);
    } else {
      c.scale(1 / 3, 1 / 3);
      leg(n - 1);
      c.rotate(60 * deg);
      leg(n - 1);
      c.rotate(-120 * deg);
      leg(n - 1);
      c.rotate(60 * deg);
      leg(n - 1);
    }
    c.restore();
    c.translate(len, 0);
  }
}

let canvas = document.querySelector("canvas");
canvas.width = 2000;
canvas.height = 2000;
let c = canvas.getContext("2d");

snowflake(c, 0, 25, 125, 125);
snowflake(c, 1, 175, 125, 125);
snowflake(c, 2, 325, 125, 125);
snowflake(c, 3, 475, 125, 125);
snowflake(c, 4, 625, 125, 125);
c.stroke();
```

#### 15.8.6 Clipping

After defining a path, you usually call `stroke()` or `fill()`.
You can also call the `clip()` method to define a clipping region.
Once a clipping region is defined, nothing will be drawn outside of it.

![Clipping](images/clipping.png)

```js
c.font = "bold 60pt sans-serif";
c.lineWidth = 2;
c.strokeStyle = "#000";

polygon(c, 3, 200, 225, 200); // Large triangle
polygon(c, 3, 200, 225, 100, 0, true); // Smaller triangle

c.clip(); // Make that path the clipping region

c.lineWidth = 10;
c.stroke();

c.fillStyle = "#aaa";
c.fillRect(175, 25, 50, 325);
c.fillStyle = "#888";
c.fillText("<canvas>", 15, 330);

function polygon(c, n, x, y, r, angle = 0, counterclockwise = false) {
  c.moveTo(x + r * Math.sin(angle), y - r * Math.cos(angle));
  let delta = (2 * Math.PI) / n;
  for (let i = 1; i < n; i++) {
    angle += counterclockwise ? -delta : delta;
    c.lineTo(x + r * Math.sin(angle), y - r * Math.cos(angle));
  }
  c.closePath();
}
```

When you call `clip()`, the current path is itself clipped to the current clipping
region, then that clipped path becomes the new clipping region.
This means that the `clip()` method can shrink the clipping region but can never
enlarge it. **There is no method to reset the clipping region, so before calling
`clip()`, you should typically call `save()` so that you can latter `restore()` the
unclipped region.**

#### 15.8.7 Pixel Manipulation

The `getImageData()` method returns an ImageData object that represent
the raw pixels (as R,G,B, and A components) from a rectangular region of your
canvas. You can create empty `ImageData` objects with `createImageData()`.
The pixels in an ImageData object are writable, so you can set them any way you
want, then copy those pixels back onto the canvas with `putImageData()`.

Pixel manipulation methods are useful for implementing image processing.

![Pixel Manipulation](images/pixel-manipulation.png)

```js
// Smear the pixels of the rectangle to the right, producing a
// sort of motion blur as if objects are moving from right to left.
// n must be 2 or larger. Larger values produce bigger smears.
// The rectangle is specified in the default coordinate system.
function smear(c, n, x, y, w, h) {
  let pixels = c.getImageData(x, y, w, h);
  let width = pixels.width,
    height = pixels.height;
  let data = pixels.data;
  let m = n - 1;
  for (let row = 0; row < height; row++) {
    let i = row * width * 4 + 4; // The offset of the second pixel of the row
    for (let col = 1; col < width; col++, i += 4) {
      data[i] = (data[i] + data[i - 4] * m) / n; // Red pixel component
      data[i + 1] = (data[i + 1] + data[i - 3] * m) / n; // Red pixel component
      data[i + 2] = (data[i + 2] + data[i - 2] * m) / n; // Red pixel component
      data[i + 3] = (data[i + 3] + data[i - 1] * m) / n; // Red pixel component
    }
  }

  c.putImageData(pixels, x, y);
}

smear(c, 10, 0, 0, 1200, 1200);
```

### 15.9 Audio APIs

The HTML `<audio>` and `<video>` tags allow you to easily include sound
and videos in your web pages, These are complex elements with
significant APIs and nontrivial user interfaces. You can control media
playback with the `play()` and `pause()` method.
You can set the `volume` and `playbackRate` properties to control the audio volume
and speed of playback. And you can skip to a particular time within
the media by setting the `currentTime` property.

However, we will not cover these tags in any further detail. Below subsections
demonstrate two ways to add scripted sound effects to your web pages.

#### 15.9.1 The Audio() Constructor

You can dynamically create `<audio>` elements with the normal DOM
`document.createElement()` method, or, as a shortcut, you can simply
use the `Audio()` constructor. **You do not have to add the created
element to your document in order to play it**. You can simply
call its `play()` method:

```js
// let soundeffect = document.createElement("audio")
// soundeffect.src = "./audios/t-rex-roar.mp3"
let soundeffect = new Audio("soundeffect.mp3");
document.addEventListener("click", () => {
  soundeffect.cloneNode().play();
});
```

**Because the Audio elements are not added to the document, they will be
garbage collected when they are done playing**.

#### 15.9.2 The WebAudio API

Web browser also allow the generation and playback of synthesized sounds
with the WebAudio API. Using the WebAudio API is like hooking up an
old-style electronic synthesizer with patch cords. With
WebAudio, you can create a set of AudioNode objects, which represents sources,
transformations, or destinations of waveforms, and then connect these nodes
together into a network to produce sounds.
The API is not particularly complex, but a full explanation requires an understanding
of electronic music and signal processing concepts that are beyond the scope
of this book.

The following code below uses the WebAudio API to synthesize a short chord that fades
out over about a second. This example demonstrates the basics of the WebAudio API.
If this is interesting to you, you can find much more about this API online:

```js
// Begin by creating an audioContext object. Safari still requires;
// us to use webkitAudioContext instead of AudioContext.
let audioContext = new (this.AudioContext || this.webkitAudioContext)();

// Define the base sound as a combination of three pure sine waves;
// D major chord: D, F#, and A
let notes = [293.7, 370.0, 440.0];

// Create oscillator nodes for each of the notes we want to play
let oscillators = notes.map((note) => {
  let o = audioContext.createOscillator();
  o.frequency.value = note;
  return o;
});

// Shape the sound by controlling its volume over time.
// Starting at time 0 quickly ramp up to full volume.
// Then starting at time 0.1 slowly ramp down to 0.
let volumeControl = audioContext.createGain();
volumeControl.gain.setTargetAtTime(1, 0.0, 0.02);
volumeControl.gain.setTargetAtTime(0, 0.1, 0.2);

let speakers = audioContext.destination;
// Connect each of the source notes to the volume control
oscillators.forEach((o) => o.connect(volumeControl));

// And connect the output of the volume control to the speakers.
volumeControl.connect(speakers);

// Now start playing the sounds and let them run for 1.25 seconds.
let startTime = audioContext.currentTime;
let stopTime = startTime + 1.25;
oscillators.forEach((o) => {
  o.start(startTime);
  o.stop(stopTime);
});

// If we want to create a sequence of sounds we can use event handlers;
oscillators[0].addEventListener("ended", () => {
  // This event handler is invoked when the note stops playing;
});
```

### 15.10 Location, Navigation, and History

The `location` property of both the Window and Document objects refers to the Location
object, **which represents the current URL of the document displayed in the window**,
and which also provides an API for loading new documents into the window.

The Location object is very much like a [URL object](#11.9)

```js
window.location.href;
window.location.toString();

window.location.protocol;
window.location.hostname;
window.location.port;
window.location.search;
window.location.hash;
```

The Location object does not have a `searchParams` property, but if you want to
parse `window.location.search`, you can simply create a URL object from
the Location object and then use the URL's searchParams:

```js
let url = new URL(window.location);
let query = url.searchParams.get("q");
```

Browsers also define a `document.URL` property, this is a string that holds
the URL of the current document

#### 15.10.1 Loading New Documents

Assigning string to `window.location`

```js
window.location = "http://www.oreilly.com";
```

You can also assign **relative URLs to the location**:

```js
window.location = "page2.html";
```

Above assignments cause the browser loads the URL and **replacing the current document with a new one**. A bare fragment identifier is a special kind of relative URL that **does not
cause the browser to load a new document but simply to scroll** so that the document
element with `id` or ~~`name`~~ matches the fragment is visible at the top of the browser window.

**Tip**: As a special case, the fragment identifier **#top** makes the browser jump
to the start of the document (**assuming no element has an `id="top"` attribute, else this element will be located**)

```js
window.location = "#top";
```

The individual properties of the Location object are **writable**, and causing browser to load
a new document or navigate within the current document.

```js
document.location.pathname = "pages/3.html"; // load a new document
document.location.hash = "#TOC"; // navigate within the current document
location.search = "?page=" + (page + 1); // load a new document
```

`assign()` method of Location object has the same effect with assigning the
string to the `location` property, however, so it's not particularly interesting.

The `replace()` method of the Location object, on the other hand, is quite useful.
**When you pass a string to `replace()`, it is interpreted as a URL and causes
the browser to load a new page**, just as `assign()` does. **The difference is that
`replace()` replaces the current document in the browser's history**. If
a script in document A set the `location` property or calls `assign()` to load
document B and then user clicks the Back button, the browser will go back to document A.
**If you use `replace()` instead, then document A is erased from the browser's history,
when the user clicks the Back button, the browser returns to whatever document was displayed before document A**.

Imagine if the user's browser does not support the web platform APIs that you want to use,
you could use `location.replace()` to load the static version.

```js
if (!isBrowserSupported()) location.replace("staticpage.html");
```

In addition, the Location object also defines `reload()`, which simply makes the browser
reload the document.

**Tip**: Assigning `location` a same url with current page reloads the page.

#### 15.10.2 Browsing History

The `history` property of the Window object refers to the History object for the window.
The `length` property specifies the number of elements in the browsing history list.
For security reasons, scripts are not allowed to access the stored URLs.

The History object has `back()` and `forward()` methods. A third method
`go()` takes an integer argument and can skip any number of pages forward (+ve argument)
or backward (-ve argument)

```js
window.history.length;
window.history.back();
window.history.forward();
window.history.go(1);
window.history.go(-1);
window.history.go(0); // another way to reload the current page
```

If a window contains child windows (such as `<iframe>` elements),
the browsing histories of the child windows are chronologically
interleaved with the history of the main window. **This means that
calling `history.back()` on the main window may cause on of the child
windows to navigate back to a previously displayed document but leaves
the main window in its current state**.

The History object described here dates back to the early days of the web
when documents were passive and all computation was performed on the server.
Today, **web applications often generate or load content
dynamically and display new application states without actually loading new documents**.
Applications like these must perform their own history management if they
want the user to be able to use the Back and Forward buttons (or the equivalent gestures) to navigate from one application state to another in an intuitive way. There
are two ways to accomplish this, described in the next two sections:

**Tip**: Setting the `location.hash` **adds an entry to the browser's history**.

#### 15.10.3 History Management with hashchange Events

Key facts you need to know to understand this technique:

- The `location.hash` does not have to be an element ID: you can set it to any string.
  As long as no element happens to have that string as its ID, the browser won't scroll when
  you set the `hash` property like this.

- Setting the `location.hash` **adds an entry to the browser's history**.

- Whenever the fragment identifier of the document changes, the browser fires
  a "hashchange" event on the Window object.

- If the user now clicks the Back button, the browser will return to its previous URL before you
  set `location.hash`. But this means that the fragment identifier has changed again,
  so another "hashchange" event is fired in this case. This means that as long
  as you can create a **unique fragment identifier for each possible state for your application.**

To use this history management mechanism, you'll need to be able to encode
the state information necessary to render a "page" of your application
into a relatively short string of text that is suitable for use as a fragment
identifier. And you'll need to write a function to convert page state
into a string and another function to parse the string and recreate the page
state it represents.

#### 15.10.4 History Management with `pushState()`

This more robust history-management technique is based on the
`history.pushState()` method and the `popstate` event.
When a web app enters a new state, it calls `history.pushState()` to add
an object representing the state to the browser's history. If the user
then clicks the Back button, the browser fires a "popstate" event with
a **copy of that saved state object**, and the app uses this object to
re-create its previous state. In addition to the saved state object,
applications can also save a URL with each state, which is important
if you want users to be able to bookmark and share links to the internal
states of the app.

The first argument to `pushState()` is an object that contains all the
state information necessary to restore the current state of the document.
The second argument was intended to be a title string for the state,
but most browsers do not support it, and you should just pass an empty string.
The third argument is an optional URL that will be **displayed in the location bar
immediately and also if the user returns to this state via Back and Forward buttons**.

**Tip**: The object is saved using HTML's _structured clone algorithm_, it
can support Map, Set, RegExp and Date objects as well as typed arrays and ArrayBuffers.

**Tip**: Associating a URL with each state allows the user to bookmark internal states
of your application (see the code example below). If the user saves a bookmark and
then visits it a day later, you won't get a "popstate"event about that visit:
you'll have to restore your application state by parsing the URL.

---

> **The Structured Clone Algorithm**

It is defined by the HTML standard. It enables serialization of most other
JavaScript types, including Map, Set, Date, RegExp, and typed arrays, and it can
handle data structures that include circular references. The structured clone algorithm
cannot serialize functions or classes, however. When cloning objects it does not
copy the prototype object, getters and setters, or non-enumerable properties.

---

In addition to the `pushState()` method, the History object also
defines `replaceState()`, which takes the same arguments but **replace the
current history state** instead of adding adding a new state to the browsing history.
When an application that uses `pushState()` is first loaded, it is often a good
idea to call `replaceState()` to define a state object for this initial state
of the application.

When the user navigates to saved history states using the Back or Forward buttons,
the browser fires a "popstate" event on the Window object. The event object associated
with the event has a property named `state`, which is the copy of the state object
you passed to `pushState()`

Below is a simple web application which use `pushState()` to save its history,
allowing the user to "go back" to review or redo their guess

![Guess Game](images/guess-game.png)

HTML:

```html
<!doctype html>
<html lang="en">
  <head>
    <title>I'm thinking of a number...</title>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <style>
      body {
        height: 250px;
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: space-evenly;
      }
      #heading {
        font: bold 36px sans-serif;
        margin: 0;
      }
      #container {
        border: solid black 1px;
        height: 1em;
        width: 80%;
      }
      #range {
        background-color: green;
        margin-left: 0%;
        height: 1em;
        width: 100%;
      }
      #input {
        display: block;
        font-size: 24px;
        width: 60%;
        padding: 5px;
      }
      #playagain {
        font-size: 24px;
        padding: 10px;
        border-radius: 5px;
      }
    </style>
  </head>
  <body>
    <h1 id="heading">I'm thinking of a number...</h1>
    <div id="container">
      <div id="range"></div>
    </div>
    <input type="text" id="input" />
    <button id="playagain" onclick="location.search='';" hidden>
      Play Again
    </button>
    <script src="./scripts/guessgame.js"></script>
  </body>
</html>
```

JS:

```js
class GameState {
  static newGame() {
    let s = new GameState();
    s.secret = s.randomInt(0, 100);
    s.low = 0;
    s.high = 100;
    s.numGuesses = 0;
    s.guess = null;
    return s;
  }

  static fromStateObject(stateObject) {
    let s = new GameState();
    for (let key of Object.keys(stateObject)) {
      s[key] = stateObject[key];
    }
    return s;
  }

  toURL() {
    let url = new URL(window.location);
    url.searchParams.set("l", this.low);
    url.searchParams.set("h", this.high);
    url.searchParams.set("n", this.numGuesses);
    url.searchParams.set("g", this.guess);
    return url.href;
  }

  static fromURL(url) {
    let s = new GameState();
    let params = new URL(url).searchParams;
    s.low = parseInt(params.get("l"));
    s.high = parseInt(params.get("h"));
    s.numGuesses = parseInt(params.get("n"));
    s.guess = parseInt(params.get("g"));

    if (
      isNaN(s.low) ||
      isNaN(s.high) ||
      isNaN(s.numGuesses) ||
      isNaN(s.guess)
    ) {
      return null;
    }

    s.secret = s.randomInt(s.low, s.high);
    return s;
  }

  randomInt(min, max) {
    return min + Math.ceil(Math.random() * (max - min - 1));
  }

  render() {
    let heading = document.querySelector("#heading");
    let range = document.querySelector("#range");
    let input = document.querySelector("#input");
    let playagain = document.querySelector("#playagain");

    heading.textContent =
      document.title = `I'm thinking of a number between ${this.low} and ${this.high}`;
    range.style.marginLeft = `${this.low}%`;
    range.style.width = `${this.high - this.low}%`;

    input.value = "";
    input.focus();

    if (this.guess === null) {
      input.placeholder = "Type your guess and hit Enter";
    } else if (this.guess < this.secret) {
      input.placeholder = `${this.guess} is too low. Guess again`;
    } else if (this.guess > this.secret) {
      input.placeholder = `${this.guess} is too high. Guess again`;
    } else {
      input.placeholder = document.title = `${this.guess} is correct!`;
      heading.textContext = `You win in ${this.numGuesses} guesses`;
      playagain.hidden = false;
    }
  }

  updateForGuess(guess) {
    if (guess > this.low && guess < this.high) {
      if (guess < this.secret) this.low = guess;
      else if (guess > this.secret) this.high = guess;
      this.guess = guess;
      this.numGuesses++;
      return true;
    } else {
      alert(
        `Please enter a number greater than ${this.low} and less than ${this.high}`,
      );
      return false;
    }
  }
}

let gamestate = GameState.fromURL(window.location) || GameState.newGame();

// Save this initial state of the game into the browser history, but use
// replaceState instead of pushState() for this initial page
history.replaceState(gamestate, "", gamestate.toURL());

gamestate.render();

document.querySelector("#input").onchange = (event) => {
  if (gamestate.updateForGuess(parseInt(event.target.value))) {
    history.pushState(gamestate, "", gamestate.toURL());
  }
  gamestate.render();
};

window.onpopstate = (event) => {
  gamestate = GameState.fromStateObject(event.state);
  gamestate.render();
};
```

### 15.11 Networking

This sections covers three network APIs:

- `fetch()` method defines a Promised-based API

- Server-Sent Events (or SSE) API is a convenient, event-based interface to
  HTTP "long polling" techniques where the web server holds the network connection
  open so that it can send data to client whenever it wants.

- WebSocket is a networking protocol that is not HTTP but is designed to interoperate
  with HTTP. It defines an asynchronous message passing API where clients and servers
  can send and receive messages from each other in a way that is similar to TCP network
  sockets.

#### 15.11.1 `fetch()`

For basic HTTP requests, using `fetch()` is a three-step process:

1. Call `fetch()`

2. Get the response object that is asynchronously returned by step 1, and call
   a method of this response object to ask for the body of the response

3. Get the body object that is asynchronously returned by step 2 and process it however
   you want.

The `fetch()` API is completely Promised-based.

```js
fetch("http://localhost:8080")
  .then((response) => response.text())
  .then(console.log);
```

Here's a example made using the `async` and `await` keywords

```js
async function isServiceREady() {
  let response = await fetch("...");
  let body = await response.text();
  return body === ready;
}
```

---

> **Goodbye XMLHTTPRequest**

It has nothing to do with XML. There is no reason today to use XMLHttpRequest.
It has been completely replaced by `fetch()`

---

##### Http Status Codes, Response Headers, and Network Errors

Here's a more realistic version:

```js
fetch("http://localhost:8081")
  .then((response) => {
    console.log(response.ok);
    console.log(response.headers.get("content-type"));
    if (
      response.ok &&
      response.headers.get("content-type") === "application/json"
    ) {
      return response.json();
    } else {
      throw new Error(
        `Unexpected response status ${response.status} or content type`,
      );
    }
  })
  .then((currentUser) => {
    displayUserInfo(currentUser);
  })
  .catch((error) => {
    console.log("Error while fetching current user:", error);
  });
```

**Tip**: Headers are **case insensitive**, according to RFC 2616. If you call get with
all lowercase characters, it will still obtain the Content Type successfully.

**Tip**: response.ok property is `true` if status is between 200 and 299

**Tip**: You can use `has()` method of response object to test the presence of
a header

The Headers object is also iterable:

```js
for (let [name, value] of response.headers) {
  console.log(name, ":", value);
}
```

It is always a good idea to include a `.catch()` clause any time you make
a `fetch()` call.

##### Setting Request Parameters

Instead of placing "?" in url for `fetch()` method,
the `fetch()` function also accepts URL objects as its first argument. So you
can include request parameters like this:

```js
async function search(term) {
  let url = new URL("/api/search");
  url.searchParams.set("q", term);
  let response = await fetch(url);
  if (!response.ok) throw new Error(response.statusText);
  let resultsArray = await response.json();
  return resultsArray;
}
```

##### Setting Request Headers

In order to set request headers, you need a second argument for `fetch()` method.
The second argument is an object that can provide additional options, including
request headers.

```js
let headers = new Headers();
headers.set("Authorization", `Basic ${btoa(`${username}:${password}`)}`);
fetch("http://localhost:8081", { headers })
  .then((response) => response.json())
  .then(console.log);
```

You can init them in constructor:

```js
let headers = new Headers({
  Authorization: `Basic ${btoa(`${username}:${password}`)}`,
});
```

**Tip**: Same case as content type. All headers: includes "Authorization" can be
mixed-case.

An alternative to passing two arguments to `fetch()` is to instead pass the same
two arguments to the Request() constructor and then pass the resulting Request object
to `fetch()`

```js
let request = new Request("http://localhost:8081", { headers })
fetch(request).then(response => ...)
```

##### Parsing Response Bodies

The `response.json()` or `response.text()` are probably the two most common scenarios,
but they are not the only ways to obtain the body of a web server's response.

The Response object also has these methods:

_arrayBuffer()_

This method returns a promise that resolves to an ArrayBuffer. This
is useful when the response contains binary data. You can use ArrayBuffer
to create a typed array [11.2](#11.2) or a DataView object [11.2.5](#11.2.5)
from which you can read the binary data.

```js
..
.then(response => response.arrayBuffer())
.then(buf => {
  let textDecoder = new TextDecoder("utf-8")
  decoded = textDecoder.decode(buf)
  console.log(decoded)
})
```

_blob()_

This method returns a promise that resolves to a Blob object. Blob
stands for "Binary Large Object", and they are useful when you expect large
amounts of binary data such as pdf. If you ask for the body of the response
as a Blob, **the browser implementation may stream the response data to a temporary
file and then return a Blob object that represents the temporary file**, therefore,
do not allow random access to the response body the way that an ArrayBuffer does.

Once you have a Blob, you can create a URL that refers to it with `URL.createObjectURL()`, or you can use the event-based FileReader API to asynchronously obtain the
content of the Blob as a string or an ArrayBuffer. At the time of writing,
some browsers also define Promise-based `text()` and `arrayBuffer()` to obtain the
content of Blob.

Below is the URL (including blob: prefix) created with `URL.createObjectURL()`

![Url Create Object Url](images/url-create-object-url.png)

_formData()_

You should uses this method if you expect the body of the Response
to be encoded in "multipart/form-data" format. This format is common
in POST requests made to server, but uncommon in server responses,
so this method is not frequently used.

##### Streaming Response Bodies

In addition to the five response methods that asynchronously return
some form of the complete response body to you, there is also an option to
stream the response body, which is useful if there is some kind of processing
you can do on the chunks of the response body as they arrive over the network.
But streaming the response is also useful if you want to display a
progress bar so that the user can see the progress of the download.

The `body` project of a Response object is a **ReadableStream** object.
If you have already called a response method like `json()`, that reads,
parses, and returns the body, the `bodyUsed` will be `true` to indicate
the `body` stream have already been read. In the case of stream has not yet been
read, you can call `getReader()` to obtain a stream reader object, then use
the `read()` method of this reader to **asynchronously read chunks of text from the stream**. The `read()` return a Promise that resolves to an object with `done` and `value`
properties. `value` will either be the next chunk, **as a Uint8Array, or `undefined` if
there is no more chunk**.

This streaming API is relatively straightforward if you use `async` and `await`
but is surprisingly complex if you attempt to use it with raw promises.

Below example demonstrates downloading a large JSON file and report download progress to the user.

**Tip**: I faced a difficulty to get decompressed content length. The user agent
forbids the modification of "Accept-Encoding" header, the below JavaScript
compares the bytes read (decompressed) with compressed content length, causing
the incorrect result. Instead, we can add a response header that indicate the decompressed length
the resource has.

HTML:

```html
<div id="progress-bar">
  <div id="progress"></div>
</div>
```

CSS:

```css
#progress-bar {
  border: solid black 1px;
  height: 1em;
}
#progress {
  width: 0%;
  height: 1em;
  background-color: blue;
}
```

JavaScript:

```js
fetch(
  "https://raw.githubusercontent.com/json-iterator/test-data/master/large-file.json",
)
  .then((response) => streamBody(response, updateProgress))
  .then((bodyText) => JSON.parse(bodyText))
  .then(handleBigJSONObject);

async function streamBody(response, reportProgress, processChunk) {
  let expectedBytes = parseInt(response.headers.get("Content-Length"));
  let bytesRead = 0;
  let reader = response.body.getReader();
  let decoder = new TextDecoder("utf-8");
  let body = "";
  while (true) {
    let { done, value } = await reader.read();
    if (value) {
      if (processChunk) {
        let processed = processChunk(value);
        if (processed) {
          body += processed;
        }
      } else {
        body += decoder.decode(value, { stream: true });
      }
      if (reportProgress) {
        bytesRead += value.length;
        reportProgress(bytesRead, bytesRead / expectedBytes);
      }
    }
    if (done) {
      break;
    }
  }

  return body;
}

function updateProgress(read, percentage) {
  let progressBar = document.querySelector("#progress");
  progressBar.style.width = `${percentage * 100}%`;
}

function handleBigJSONObject(text) {
  console.log(text);
}
```

![Progress Bar](images/progress-bar.png)

##### Specifying The Request Method And Request Body

To use a different request method (POST, DELETE, PATCH...), simply
use the two argument version of `fetch()`

```js
fetch(url, { method: "POST" })
  .then((r) => r.json())
  .then(console.log);
```

As long as the `method` property is not set to "GET" or "HEAD" (which do not support
request bodies, they throw TypeError if you attempt to do so), you can specify a request body by setting the `body` property
of the Options object:

```js
fetch(url, {
  method: "POST",
  body: "hello, world.",
});
```

When you specify a request body, the browser automatically adds an appropriate
"Content-Length" header to the request. When the body is a string,
**the browser defaults the Content-Type header to "text/plain;charset=UTF-8"**.
You may override it:

```js
fetch(url, {
  method: "POST",
  body: JSON.stringify(requestBody),
  headers: new Headers({ "Content-Type": "application/json" }),
});
```

If you have binary data in Blob form, you can simply set `body` to the Blob.
Blobs have a `type` property that specifies their content type, and the
value of this property is used as the **default value** of the "Content-Type" header.

```js
const obj = { hello: "world" };
const blob = new Blob([JSON.stringify(obj, null, 2)], {
  type: "application/json",
});
fetch("...", {
  method: "POST",
  body: blob,
});
```

With POST requests, it is somewhat common to pass a set of name/value
parameters in the request body (instead of encoding them). There are two ways to do this:

- Specifying your parameters names and values with URLSearchParams and then pass
  URLSearchParams as the value of `body` property. If you do this, the "Content-Type"
  header will be **automatically** set to "application/x-www-form-urlencoded;charset=UTF8"

- If instead you specify your parameter names and values with a FormData object,
  the body will use a more verbose multipart encoding and "Content-Type" will be
  set to "multipart/form-data;boundary=..." with a unique boundary string that matches
  the body.

  - Using a FormData object is particularly useful when the values you want to upload
    are long or are File or Blob objects that may each have its own "Content-Type".
    FormData objects can be created and initialized with values by passing a `<form>`
    element to the `FormData()` constructor. But you can also create "multipart/form-data"
    request bodies by invoking the `FormData()` constructor with no arguments
    and initializing the name/value pairs it represents with the `set()` and `append()`
    methods.

    ![Contenttypeformdata](images/contenttypeformdata.png)

    ![Formdata](images/formdata.png)

##### File Upload With `fetch()`

It can be accomplished using a FormData object as the request body.
A common way to obtain a File object is to display an `<input type="file">` element
on your web page and listen for "change" events on that element. When
a "change" event occurs, the `files` **array** of the input element **should contain
at least one File object**.

```js
let files = document.querySelector("input[name=files]");
files.onchange = (event) => {
  console.log(event.target.files);
};
```

File objects are also available through the HTML drag-and-drop API -- not covered in this book,
but you can get files from the `dataTransfer.files` array of the event object passed
to an event listener for "drop" events.

```js
files.ondrop = (e) => {
  console.log(e.dataTransfer.files);
};
```

Remember also the File objects are a kind of Blob, and sometimes it can be useful
to upload Blobs. Suppose you've written a web application
that allows the user to create drawings in a `<canvas>` element. You
can upload the user's drawings as PNG files with code like the following:

```js
async function getCanvasBlob(canvas) {
  return new Promise((resolve, reject) => {
    canvas.toBlob(resolve);
  });
}

async function uploadCanvasImage(canvas) {
  let pngblob = await getCanvasBlob(canvas);
  let formdata = new FormData();
  formdata.set("canvasimage", pngblob);
  let response = await fetch("/upload", { method: "POST", body: formdata });
  let body = await response.json();
}
```

![Png](images/png.png)

##### Cross-Origin Requests

Same origin: protocol plus hostname plus port

For security reasons, web browsers generally disallow **(though there are
exceptions for images and scripts)** cross-origin network requests.
However, Cross-Origin Resource Sharing, or CORS, enables safe cross-origin requests.
When `fetch()` is used with a cross-origin URL, the browser adds an "Origin"
header to the request (and does not allow it to be overridden via the `headers` property-- Forbidden header name)
to **notify the web server** that the request is coming from a document with an
appropriate "Access-Control-Allow-Origin" header, then the request proceeds.
Otherwise, if the server does not explicitly allow the request, then the Promise returned
by `fetch()` is rejected.

##### Aborting A Request

The fetch API allows requests to be aborted using the AbortController and AbortSignal
classes. (These classes defined a generic abort mechanism suitable for use
by other APIs as well).

If you want to have the option of aborting a `fetch()` request, then
create an AbortController object before starting the request. The `signal` property
of AbortController object is an AbortSignal object. Pass this `signal`
object to the options object that you pass to `fetch()`

Having done that, you can call the `abort()` method of the controller object to abort
the request, **which will cause any Promise objects related to the fetch request
to reject with an exception**.

---

**Important**: Web server should be able to detect if the context is aborted, and the connections to database should know this status of context as well to cancel the operations.

For example in Golang, we can use `request.Context().Err()` to know if
the context is cancelled; for db connection to know the context, for example with GORM, we use
`db.WithContext(ctx).[Operation:eg Create]`.

---

```js
fetchWithTimeout("http://localhost:8081", { timeout: 2000 });

function fetchWithTimeout(url, options = {}) {
  if (options.timeout) {
    let controller = new AbortController();
    options.signal = controller.signal;
    setTimeout(() => {
      controller.abort();
    }, options.timeout);
  }
  return fetch(url, options);
}
```

![Abort Network Request](images/abort-network-request.png)

![Abort Error](images/abort-error.png)

##### Miscellaneous Request Options

Options object supports a number of other options as well, including these:

`cache`

Use this property to override the browser's default caching behavior.
See [Mozilla Explanations for more](https://developer.mozilla.org/en-US/docs/Web/API/Request/cache)

- "default"

  - This value specifies the default caching behavior. Fresh responses
    in the cache are served directly from the cache, and stale (no longer fresh) responses are
    revalidated before being served.

- "no-store"

  - This value makes the browser ignore its cache. The cache is not checked
    for matches when the request is made and is not updated when the response arrives.

- "reload"

  - This value tells the browser to always make a normal network
    request, ignoring the cache. When the response arrives, however, it is stored
    in the cache.

- "no-cache"

  - Misleading named.. It tells the browser to not serve fresh values from the cache.
    Fresh or stale cached values are revalidated before being returned.

- "force-cache"

  - This value tells the browser to serve responses from the cache even if they
    are stale.

`redirect`

This property controls how the browser handles redirect responses (300-309 and with Location response header) from the server:

**Important**: only simple response headers are exposed when using CORS. If you want to expose non-simple headers,
set the `Access-Control-Expose-Headers` headers:

```js
Access-Control-Expose-Headers: Location
```

- "follow"

  - It is the default value, makes the browser follow redirect automatically.

- "error"

  - This value makes `fetch()` reject its returned Promise if the server returns
    a redirect response.

    ![Failed To Fetch](images/failed-to-fetch.png)

- "manual"

  - This value means that you want to manually handle redirect responses, and the
    Promise returned by `fetch()` may resolve to a Response object with a `status`
    in the 300 to 399 range. In this case, you will have to use the "Location" header
    of the Response to manually follow the redirection.

  - **Question**: I tried with this option to get `response.headers.get("location")` but
    it returned null. Why causes it and how to solve it?

`referrer`

You can set this property to a string that contains a relative URL to specify the
value of the HTTP "Referer" header (which is historically misspelled with three Rs instead of four). If you set this property to the empty string,
then the "Referer" header will be omitted from the request.

**Tip**: This header contains where user access the API. It is useful for things such as
analytics (knowing what journey a visitor took on a site).

#### 15.11.2 Server-Sent Events

Some web apps find it useful, to have their server send them notifications when
events occur. This technique has been devised is for the client to make a
request to the server, and then neither the client nor the server close the connection.
When the server has something to tell the client about, it writes data to the connection
but keeps it open. Network connections like this don't usually stay open forever,
but if the client detects that the connection has closed, it can simply make
another request to reopen the connection.

This technique for allowing servers to send messages to clients is surprisingly
effective (though it can be expensive on the server side because the server must
maintain active connection to all its clients).

Because it is a useful programming pattern, client-side JavaScript supports
it with the EventSource API. When the server writes data to the connection,
the EventSource object translates those into events that you can listen for:

```js
let ticker = new EventSource("http://localhost:8081/sse");
ticker.addEventListener("bid", (event) => {
  console.log(event.data);
});
```

The server determines the type of the events that are generated. **If the
server omits an event name in the data it writes, then the event type defaults
to "message": not "bid"**

Server:

```go
http.HandleFunc("/sse", func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	for i := 0; i < 100; i++ {
		fmt.Fprintf(w, "event: bid\ndata: streaming %d\n\n", i)
		time.Sleep(2 * time.Second)
		w.(http.Flusher).Flush()
	}

	// Simulate closing the connection
	<-r.Context().Done() // Client's connection has gone
})
```

![Event Stream Network](images/event-stream-network.png)

![Event Stream Console](images/event-stream-console.png)

There are some additional details to the protocol that allow events
to be given IDs and allow a reconnecting client to tell the server
what the ID of the last event it received was, so that a server can resend
any events is missed. Those details are invisible on the client side,
however, and are not discussed here.

One obvious application for SSE is for multiuser
collaborations like online chat. A chat client might use `fetch()` to
post messages to the chatroom and subscribe to the stream of chatter
with an EventSource object.

![Chatroom](images/chatroom.png)

Front End:

```html
<input id="input" style="width: 100%; padding: 10px; border: solid black 2px" />
<script>
  let nick = prompt("Enter your nickname");
  let escapednick = encodeURIComponent(nick);
  let input = document.querySelector("#input");
  input.focus();

  let chat = new EventSource(`http://localhost:8081/chat?name=${escapednick}`);
  chat.addEventListener("chat", (event) => {
    let div = document.createElement("div");
    div.append(event.data);
    input.before(div);
    input.scrollIntoView();
  });

  input.addEventListener("change", () => {
    fetch("http://localhost:8081/chat", {
      method: "POST",
      body: nick + ": " + input.value,
    }).catch(console.error);

    input.value = "";
  });
</script>
```

Back End:

```go
func main() {
  ...
	chatRoom := NewChatRoom()
	http.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Content-Type", "text/event-stream")
			w.Header().Set("Cache-Control", "no-cache")
			w.Header().Set("Connection", "keep-alive")

			name := r.URL.Query().Get("name")

			client := chatRoom.addClient()
			chatRoom.sendToAll(name + " joined the chat")

			for {
				select {
				case msg := <-client:
					msg = strings.ReplaceAll(msg, "\n", "\ndata: ")
					fmt.Fprintf(w, "event: chat\ndata: %s\n\n", msg)
					w.(http.Flusher).Flush()
				case <-r.Context().Done(): // Client's connection has gone
					fmt.Println("bye?")
					chatRoom.removeClient(client)
					chatRoom.sendToAll(name + " left the chat")
					close(client)
					return
				}
			}
		case http.MethodPost:
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.WriteHeader(200)
			b, _ := io.ReadAll(r.Body)
			msg := string(b)
			chatRoom.sendToAll(msg)
		}
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}

type ChatRoom struct {
	addch    chan chan string
	removech chan chan string
	clientch map[chan string]struct{}
}

func NewChatRoom() *ChatRoom {
	chatroom := &ChatRoom{
		addch:    make(chan chan string, 10),
		removech: make(chan chan string, 10),
		clientch: make(map[chan string]struct{}),
	}
	go func() {
		for {
			select {
			case client := <-chatroom.addch:
				chatroom.clientch[client] = struct{}{}
			case client := <-chatroom.removech:
				fmt.Println("here2?")
				delete(chatroom.clientch, client)
				fmt.Println(len(chatroom.clientch))
			}
		}
	}()

	return chatroom
}

func (c *ChatRoom) addClient() chan string {
	ch := make(chan string, 10)
	c.addch <- ch
	return ch
}

func (c *ChatRoom) removeClient(client chan string) {
	fmt.Println("here1?")
	c.removech <- client
}

func (c *ChatRoom) sendToAll(msg string) {
	for ch := range c.clientch {
		ch <- msg
	}
}
```

NodeJS version server:

```js
const http = require("http");
const fs = require("fs");
const url = require("url");

const clientHTML = fs.readFileSync("../exp/web/chatroom.html");

let clients = [];

let server = new http.Server();
server.listen(8082);

server.on("request", (request, response) => {
  let pathname = url.parse(request.url).pathname;

  if (pathname === "/") {
    response.writeHead(200, { "Content-Type": "text/html" }).end(clientHTML);
  } else if (pathname === "/chat" && request.method === "GET") {
    acceptNewClient(request, response);
  } else if (pathname === "/chat" && request.method === "POST") {
    broadcastNewMessage(request, response);
  } else {
    response.writeHead(404).end();
  }
});

function acceptNewClient(request, response) {
  clients.push(response);

  request.connection.on("end", () => {
    clients.splice(clients.indexOf(response), 1);
    response.end();
  });

  response.writeHead(200, {
    "Content-Type": "text/event-stream",
    Connection: "keep-alive",
    "Cache-Control": "no-cache",
  });
  response.write("event: chat\ndata: Connected\n\n");
}

async function broadcastNewMessage(request, response) {
  request.setEncoding("utf8");
  let body = "";
  for await (let chunk of request) {
    body += chunk;
  }

  response.writeHead(200).end();

  let message = "data: " + body.replace("\n", "\ndata: ");
  let event = `event: chat\n${message}\n\n`;
  clients.forEach((client) => client.write(event));
}
```

#### 15.11.3 WebSocket

WebSockets allow JavaScript code in the browser to easily exchange text
and binary messages with a server. As with SSE, the client
must establish the connection, but once the connection is establish,
the server can asynchronously send messages to the client. Unlike SSE,
binary messages are supported, and messages can be sent in both directions,
not just from server to client.

Connect to a service using WebSocket protocol: `wss://`. (Browser typically
restrict WebSockets to only work in pages loaded over secure `https://` connections)

To establish a WebSocket connection, the browser first establishes an HTTP
connection and sends the server an `Upgrade: websocket` header requesting
the connection be switched from HTTP protocol to the WebSocket protocol.
You will need to be working with a web server that also speaks the WebSocket protocol.

![Websocket Headers](images/websocket-headers.png)

##### Creating, Connecting, and Disconnecting WebSockets

Communicate with a WebSocket-enabled server:

```js
let socket = new WebSocket("wss://example.com/stockticker");
```

When you create a WebSocket, the connection **process** begins automatically.
But a newly created WebSocket will not be connected when it is first returned.

There are few state of the connections is in, the `readyState` property of the
socket specifies them. This property can have the following values:

| state                  | Description                                                          |
| ---------------------- | -------------------------------------------------------------------- |
| `WebSocket.CONNECTING` |                                                                      |
| `WebSocket.OPEN`       | This WebSocket is connected and ready for communication              |
| `WebSocket.CLOSING`    |                                                                      |
| `WebSocket.CLOSED`     | This WebSocket has been closed; no further communication is possible |
|                        | This state can also occur when the initial connection attempt fails. |

When a WebSocket transitions from CONNECTING to OPEN, it fires an "open" event.
If a protocol or other error occurs for a WebSocket connection, the WebSocket object
fires an "error" event

When you are done with a WebSocket, you can close the connection by
calling the `close()` method. When
a WebSocket changes to the CLOSED state, it fires a "close" event.

##### Sending Messages Over A WebSocket

By invoking `send()` method of the WebSocket object, it send a message to the server.
This method expects **a** single message argument, which can be a
string, Blob, ArrayBuffer, typed array, or DataView object.

The `send()` method **buffers** the specified message to be transmitted and returns
before the message is actually sent. The `bufferedAmount` property of the
WebSocket object specifies the number of bytes that are buffered but not yet sent. (
surprisingly, **WebSockets** do not fire any event when this value reaches 0.)

##### Receiving Messages From A WebSocket

To receive messages from a server over a WebSocket, register an event
handler for "message" events. The object associated with this event is a MessageEvent
instance with a `data` property that contains the server's message.

If the server sent UTF-8 encoded text, then `event.data` will be a string
containing that text. If the server sends a message that consists of
binary data instead of text, then the `data` property will (by default)
be a Blob object representing that data. If you prefer to receive binary messages
as ArrayBuffers instead of Blobs, set the `binaryType` property of the
WebSocket object to the string "arraybuffer".

##### Protocol Negotiation

Protocol tend to evolve, however. If a hypothetical stock quotation protocol is updated,
you can define a new URL: `wss://example.com/stockticker/v2`.
URL-based versioning is not always sufficient, however. With complex protocols
that have evolved over time, you may end up with deployed servers that support multiple
versions of the protocol and deployed clients that support a different set of protocol
versions.

Anticipating this situation, the WebSocket protocol and API include an application-level
protocol negotiation feature. When you call the `WebSocket()` constructor,
the `wss://` URL is the first argument, but you can also pass an array of strings as the
second argument. If you do this, you are specifying a list of application protocols that you
know how to handle and asking the server to pick one.
**During the connection process, the server will choose one of the protocols (or will
fail with an error if it does not support any of the client's options).** Once
the connection has been established, the `protocol` property of the WebSocket object specifies
which protocol version the server chose.

##### Exercise

[Followed this step by step](https://programmingpercy.tech/blog/mastering-websockets-with-go/)

FrontEnd:

```html
<!doctype html>
<html lang="en">
  <head>
    <title></title>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
  </head>
  <body>
    <h1>Amazing Chat Application</h1>
    <h3 id="connection-header">Connected to Websocket: false</h3>

    <h3 id="chat-header">Currently in chat: general</h3>

    <form id="chatroom-selection">
      <label for="chatroom">Chatroom:</label>
      <input type="text" id="chatroom" name="chatroom" /><br /><br />
      <input type="submit" value="Change chatroom" />
    </form>

    <textarea
      class="messagearea"
      id="chatmessages"
      readonly
      name="chatmessages"
      rows="4"
      cols="50"
      placeholder="Welcome to the general chatroom, here messages from others will appear"
    ></textarea>
    <br />

    <form id="chatroom-message">
      <label for="message">Message:</label>
      <input type="text" name="message" id="message" /><br /><br />
      <input type="submit" value="Send message" />
    </form>

    <div>
      <form id="login-form">
        <label for="username">Username:</label>
        <input type="text" id="username" name="username" /><br /><br />
        <label for="password">Password:</label>
        <input type="password" id="password" name="password" /><br /><br />
        <input type="submit" value="Login" />
      </form>
    </div>
    <script src="./websocket.js"></script>
  </body>
</html>
```

```js
var selectedchat = "general";

class Event {
  // Each Event needs a Type
  // The payload is not required
  constructor(type, payload) {
    this.type = type;
    this.payload = payload;
  }
}

/**
 * SendMessageEvent is used to send messages to other clients
 * */
class SendMessageEvent {
  constructor(message, from) {
    this.message = message;
    this.from = from;
  }
}
/**
 * NewMessageEvent is messages comming from clients
 * */
class NewMessageEvent {
  constructor(message, from, sent) {
    this.message = message;
    this.from = from;
    this.sent = sent;
  }
}

class ChangeChatRoomEvent {
  constructor(name) {
    this.name = name;
  }
}
function changeChatRoom(event) {
  event.preventDefault();
  let newchat = document.querySelector("input#chatroom");
  if (newchat?.value != selectedchat) {
    selectedchat = newchat.value;
    header = document.querySelector("#chat-header").innerHTML =
      "Currently in chat: " + selectedchat;

    let changeEvent = new ChangeChatRoomEvent(selectedchat);
    sendEvent("change_room", changeEvent);
    textarea = document.querySelector("textarea#chatmessages");
    textarea.innerHTML = `You changed room into: ${selectedchat}`;
  }
}

function sendMessage(event) {
  event.preventDefault();
  let newmessage = document.querySelector("input#message");
  if (newmessage) {
    let outgoingEvent = new SendMessageEvent(newmessage.value, "percy");
    sendEvent("send_message", outgoingEvent);
    newmessage.value = "";
  }
}

function login(event) {
  event.preventDefault();
  let formData = {
    username: document.querySelector("input#username").value,
    password: document.querySelector("input#password").value,
  };

  fetch("/login", {
    method: "post",
    body: JSON.stringify(formData),
  })
    .then((response) => {
      if (response.ok) {
        return response.json();
      } else {
        throw new Error("unauthorized");
      }
    })
    .then((data) => {
      connectWebsocket(data.otp);
    })
    .catch(alert);
}

function sendEvent(eventName, payload) {
  const event = new Event(eventName, payload);
  conn.send(JSON.stringify(event));
}

function routeEvent(event) {
  if (!event.type) {
    alert("no 'type' is in event");
  }
  switch (event.type) {
    case "new_message":
      const messageEvent = Object.assign(new NewMessageEvent(), event.payload);
      appendChatMessage(messageEvent);
      break;
    default:
      alert("unsupported message type");
      break;
  }
}

function appendChatMessage(messageEvent) {
  var date = new Date(messageEvent.sent);
  const formattedMsg = `${date.toLocaleString()}: ${messageEvent.message}`;

  const textarea = document.querySelector("textarea#chatmessages");
  textarea.innerHTML = textarea.innerHTML + "\n" + formattedMsg;
  textarea.scrollTop = textarea.scrollHeight;
}

function connectWebsocket(otp) {
  if (window["WebSocket"]) {
    conn = new WebSocket("wss://" + document.location.host + "/ws?otp=" + otp);

    conn.onopen = function (evt) {
      document.querySelector("#connection-header").innerHTML =
        "Connected to Websocket: true";
    };

    conn.onclose = function (evt) {
      document.querySelector("#connection-header").innerHTML =
        "Connected to Websocket: false";
    };

    conn.onmessage = function (evt) {
      const eventData = JSON.parse(evt.data);
      const event = Object.assign(new Event(), eventData);
      routeEvent(event);
    };
  } else {
    alert("not supporting websocket");
  }
}

window.onload = function () {
  document.querySelector("form#chatroom-selection").onsubmit = changeChatRoom;
  document.querySelector("form#chatroom-message").onsubmit = sendMessage;
  document.querySelector("form#login-form").onsubmit = login;
};
```

Backend:

```go
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"golang.org/x/net/context"
)

var (
	upgrader = &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     checkOrigin,
	}
	// pongWait is how long we will await a pong response from client
	pongWait = 10 * time.Second
	// pingInterval has to be less than pongWait, We cant multiply by 0.9 to get 90% of time
	// Because that can make decimals, so instead *9 / 10 to get 90%
	// The reason why it has to be less than PingFrequency is because otherwise it will send a new Ping before getting response
	pingInterval = (pongWait * 9) / 10
)

func main() {
	rootCtx := context.Background()
	ctx, cancel := context.WithCancel(rootCtx)

	defer cancel()

	setupAPI(ctx)

	err := http.ListenAndServeTLS(":8082", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func setupAPI(ctx context.Context) {
	manager := NewManager(ctx)
	http.Handle("/", http.FileServer(http.Dir("./frontend")))
	http.HandleFunc("/ws", manager.serveWS)
	http.HandleFunc("/login", manager.loginHandler)
}

type Manager struct {
	clients ClientList

	sync.Mutex

	handler map[string]EventHandler

	otps RetentionMap
}

func NewManager(ctx context.Context) *Manager {
	m := &Manager{
		clients: make(ClientList),
		handler: make(map[string]EventHandler),
		otps:    NewRententionMap(ctx, 5*time.Second),
	}
	m.setupEventHandler()
	return m
}

func (m *Manager) setupEventHandler() {
	m.handler[EventSendMessage] = SendMessageHandler
	m.handler[EventChangeRoom] = ChangeRoomHandler
}

func (m *Manager) routeEvent(event Event, c *Client) error {
	if handler, ok := m.handler[event.Type]; ok {
		if err := handler(event, c); err != nil {
			return err
		}
		return nil
	} else {
		return ErrEventNotSupported
	}
}

func (m *Manager) serveWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := NewClient(conn, m)
	m.AddClient(client)

	go client.readMessages()
	go client.writeMessage()
}

func (m *Manager) loginHandler(w http.ResponseWriter, r *http.Request) {
	type userLoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var req userLoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.Username == "qwer" && req.Password == "qwer" {
		type response struct {
			OTP string `json:"otp"`
		}

		otp := m.otps.NewOTP()

		resp := response{
			OTP: otp.Key,
		}

		data, err := json.Marshal(resp)
		if err != nil {
			log.Println(err)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(data)
		return
	}

	w.WriteHeader(http.StatusUnauthorized)
}

func (m *Manager) AddClient(client *Client) {
	m.Lock()
	defer m.Unlock()
	m.clients[client] = true
}

func (m *Manager) RemoveClient(client *Client) {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.clients[client]; ok {
		client.connection.Close()
		delete(m.clients, client)
	}
}

type Client struct {
	connection *websocket.Conn
	manager    *Manager
	egress     chan Event
	chatroom   string
}

func NewClient(conn *websocket.Conn, manager *Manager) *Client {
	return &Client{
		connection: conn,
		manager:    manager,
		egress:     make(chan Event),
		chatroom:   "",
	}
}

func (c *Client) readMessages() {
	defer c.manager.RemoveClient(c)

	c.connection.SetReadLimit(512) //If you restart and try sending a long message, the connection will close.

	if err := c.connection.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		log.Println(err)
		return
	}
	c.connection.SetPongHandler(c.pongHandler)

	for {
		_, payload, err := c.connection.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error reading message: %e", err)
			}
			break
		}

		var request Event
		if err := json.Unmarshal(payload, &request); err != nil {
			log.Printf("error marshalling message: %v", err)
			break // Breaking the connection here might be harsh xD
		}

		if err := c.manager.routeEvent(request, c); err != nil {
			log.Println("Error handleing Message: ", err)
		}
	}
}

func (c *Client) writeMessage() {
	ticker := time.NewTicker(pingInterval)
	defer func() {
		ticker.Stop()
		c.manager.RemoveClient(c)
	}()

	for {
		select {
		case message, ok := <-c.egress:
			if !ok {
				if err := c.connection.WriteMessage(websocket.CloseMessage, nil); err != nil {
					log.Println("Connection closed: ", err)
				}
				return
			}

			data, err := json.Marshal(message)
			if err != nil {
				log.Println(err)
				return // closes the connection, should we really
			}

			if err := c.connection.WriteMessage(websocket.TextMessage, data); err != nil {
				log.Println(err)
			}
			log.Println("Sent message")
		case <-ticker.C:
			log.Println("tick")
			if err := c.connection.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				log.Println("writemsg: ", err)
				return
			}
		}
	}
}

func (c *Client) pongHandler(pongMsg string) error {
	log.Println("pong")
	return c.connection.SetReadDeadline(time.Now().Add(pongWait))
}

type ClientList map[*Client]bool

type Event struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type EventHandler func(event Event, c *Client) error

const (
	EventSendMessage = "send_message"
	EventNewMessage  = "new_message"
	EventChangeRoom  = "change_room"
)

var (
	ErrEventNotSupported = errors.New("this event type is not supported")
)

type OTP struct {
	Key     string
	Created time.Time
}

type RetentionMap map[string]OTP

func (rm RetentionMap) NewOTP() OTP {
	o := OTP{
		Key:     uuid.NewString(),
		Created: time.Now(),
	}
	return o
}

func NewRententionMap(ctx context.Context, retentionPeriod time.Duration) RetentionMap {
	rm := make(RetentionMap)
	go rm.Retention(ctx, retentionPeriod)
	return rm
}

func (rm RetentionMap) VerifyOTP(otp string) bool {
	if _, ok := rm[otp]; !ok {
		return false
	}
	delete(rm, otp)
	return true
}

func (rm RetentionMap) Retention(ctx context.Context, retentionPeriod time.Duration) {
	ticker := time.NewTicker(400 * time.Second)
	for {
		select {
		case <-ticker.C:
			for _, otp := range rm {
				if otp.Created.Add(retentionPeriod).Before(time.Now()) {
					delete(rm, otp.Key)
				}
			}
		case <-ctx.Done():
			return
		}
	}
}

func checkOrigin(r *http.Request) bool {
	origin := r.Header.Get("origin")
	switch origin {
	case "https://localhost:8082":
		return true
	default:
		return false
	}
}

type SendMessageEvent struct {
	Message string `json:"message"`
	From    string `json:"from"`
}

type NewMessageEvent struct {
	Sent time.Time `json:"sent"`
	SendMessageEvent
}

type ChangeRoomEvent struct {
	Name string `json:"name"`
}

func ChangeRoomHandler(event Event, client *Client) error {
	var changeRoomEvent ChangeRoomEvent
	if err := json.Unmarshal(event.Payload, &changeRoomEvent); err != nil {
		return fmt.Errorf("bad payload in request: %v", err)
	}

	client.chatroom = changeRoomEvent.Name

	return nil
}

func SendMessageHandler(event Event, client *Client) error {
	var chatEvent SendMessageEvent
	if err := json.Unmarshal(event.Payload, &chatEvent); err != nil {
		return fmt.Errorf("bad payload in request: %v", err)
	}

	var broadMessage NewMessageEvent

	broadMessage.Sent = time.Now()
	broadMessage.Message = chatEvent.Message
	broadMessage.From = chatEvent.From

	data, err := json.Marshal(broadMessage)
	if err != nil {
		return fmt.Errorf("failed to marshal broadcast message: %v", err)
	}

	var outgoingEvent Event
	outgoingEvent.Type = EventNewMessage
	outgoingEvent.Payload = data

	for wsclient := range client.manager.clients {
		if wsclient.chatroom == client.chatroom {
			wsclient.egress <- outgoingEvent
		}
	}

	return nil
}
```

### 15.12 Storage

Web applications can use browser API **to store data locally on the user's computer**.
This client-side storage serves to give the web browser a memory. Web apps
can store user preferences, for example, or even store their complete state, so that
they can resume exactly where you left off at the end of your last visit.
**Client-side storage is segregated by origin**, **so pages from one site can't read
the data stored by pages from another site**. But two pages from the same site can
share storage and use it as a communication mechanism. Web applications can choose
the lifetime of the data they store.

There are a number of forms of client-side storage:

_Web Storage_

The Web Storage API consists of `localStorage` and `sessionStorage` objects,
which are essentially persistent objects that map string keys to string values.
Web Storage is very easy to use and is suitable for storing large (but not huge)
amounts of data

_Cookies_

Cookies are an old cliend-side storage mechanism that was designed for use
by server-side scripts. An awkward JavaScript API makes cookies scriptable on the
client side, but they're hard to use. It suitable only for storing small amounts
of textual data. **Also, any data stored as cookies is always transmitted to the server
with every HTTP request, even if the data is only interest to the client**

_IndexedDB_

IndexedDB is an asynchronous API to an **object database** that supports **indexing**.

---

> **Storage, Security, and Privacy**

Web browsers often offer to remember web passwords for you, and they store
them safely in **encrypted** form on the device. But none of the forms of
client-side data storage described in this chapter involve encryption: **you should
assume that anything your web applications save resides on the user's device in
unencrypted form**. Stored data is therefore accessible to curious users
who share access to the device and to malicious software(such as spyware) that
exists on the device. For this reason, **no form of client-side storage should
ever be used for passwords, financial account numbers, or other similarly sensitive
information**.

---

#### 15.12.1 `localStorage` and `sessionStorage`

The `localStorage` and `sessionStorage` properties of the **Window object** refer
to Storage objects. They behave much like a regular JavaScript object, except that:

- The property values of Storage objects must be strings.

- The properties stored in a Storage object **persist**.

```js
let name = localStorage.username;

if (!name) {
  name = prompt("What is your name?");
  localStorage.username = name;
}
```

You can use the `delete` operator to remove properties, and you can use a
`for/in` loop or `Object.keys()` to enumerate the properties of a Storage object.
If you want to remove all properties of a storage object, call the `clear()` object.

```js
localStorage.clear();
delete localStorage.username;

for (let prop of Object.keys(localStorage)) {
  console.log(prop);
}
```

Keep in mind that the properties of Storage objects can only store strings (The number will
be automatically converted into string).
If you want to store and retrieve other kinds of data, you'll have to encode
and decode it yourself.

```js
localStorage.x = 10;
let x = parseInt(localStorage.x);

localStorage.lastRead = new Date().toUTCString();
let lastRead = new Date(localStorage.lastRead);

localStorage.data = JSON.stringify(data);
let data = JSON.parse(localStorage.data);
```

##### Storage Lifetime and Scope

The difference between `localStorage` and `sessionStorage` involves the lifetime
and scope of the storage. Data stored through `localStorage` **is permanent** -- until
user asks the browser to delete it or JavaScript deletes it.

`localStorage` is scoped to document origin (protocol + hostname + port). Documents
with different origins can never read or overwrite each other's data even if they're
both running a script from the same third-party server.

`localStorage` is also scoped by browser implementation. You can never access `localStorage` in GoogleChrome in Firefox

Data stored through `sessionStorage` has a different lifetime than data stored through
`localStorage`: it has the same lifetime as the top-level window or browser tab in which
the script that stored it is running. When the window or tab is permanently closed, any data
stored through `sessionStorage` is deleted. (Note, however, that modern
browser have the ability to reopen recently closed tabs and restore the last browsing
session, so the lifetime of these tabs and their associated `sessionStorage` may be
longer than it seems)

Like `localStorage`, `sessionStorage` is scoped to the document origin. **But `sessionStorage` is also scoped on a per-window basis**. **If user has two browser tabs displaying
documents from the same origin, those two tabs have separate `sessionStorage` data**

##### Storage Events

Whenever the data stored in `localStorage` **changes**, the browser
triggers a "storage" event on any other Window objects to which that data
is visible **(but not on the Window that made the change)**.

The event object associated with a "storage" event has some important properties:

_key_

The name or key of the item that was set or removed. If the `clear()` method
was called, this property will be `null`

_newValue_

Holds the new value of the item, if there is one. If `removeItem()` was called,
this property will not be present (`null`)

_oldValue_

Holds the old value of an existing item that changed or was deleted.
If the new property is added, then this property will not be present (`null`).
If the `clear()` method was called, this property will be `null`

_storageArea_

The Storage object that changed. This is usually the `localStorage` object.

_url_

The URL (as a string) of the document whose script made this storage change.

Note that `localStorage` and the `storage` event can serve as a broadcast
mechanism by which a browser sends a message to all windows that are currently
visiting the same website. As example, imagine a web-based image-editing application
that allows the user to display tool palettes in separate windows. When the user selects
a tool, the application can generate a notification to other windows that a new tool
has been selected.

<a id="15.12.2"></a>

#### 15.12.2 Cookies

At the lowest level, cookies are implemented as an extension to the HTTP protocol. Cookie data
is automatically transmitted between the web browser and web server, so server-side
scripts can read and write cookie values that are stored on the client.

---

> **Why "Cookie"**?

The term "cookie" or "magic cookie" has been used to refer to a small chunk of data,
particularly a chunk of **privileged or secret data, akin to a password, that proves
identity or permits access.** In JavaScript, cookies are used to save state and
can establish a kind of identify for a web browser. Cookies in JavaScript do not
use any kind of cryptography, however, and are not secure in any way (although transmitting
them across an `https:` connection helps)

---

##### Reading Cookies

When you read the `document.cookie` property, it returns a string that contains
**all the cookies that apply to the current document**. The string is a list
of name/value pairs separated from each other by a **semicolon (; ) AND a space**.
The cookie value is just the value itself and **does not include any of the attributes
that may be associated with that cookie**.

```js
function getCookie() {
  let cookies = new Map();
  let all = document.cookie;
  let list = all.split("; ");

  for (let cookie of list) {
    if (!cookie.includes("=")) continue;

    let p = cookie.indexOf("=");
    let name = cookie.substring(0, p);
    let value = cookie.substring(p + 1);
    value = decodeURIComponent(value); // Assume that cookie values are encoded with encodeURIComponent()

    cookies.set(name, value);
  }
  return cookies;
}
```

##### Cookie Attributes: Lifetime And Scope

Cookies are transient by default: the values they store last for the duration of
the web browser session, they are lost when the user exits the browser.
If you want a cookie to last beyond a single browsing session, you must
tell the browser how long (in seconds) you would like it to retain the cookie
by specifying a **`max-age` attribute**.

Cookie visibility is scoped by document origin as `localStorage` and `sessionStorage` are,
and `domain`
but also by document path. This scope is configurable through cookie attributes `path`
and `domain`. By default, a cookie is associated with, and accessible to,
the web page (server, not client!) that created it and any other web pages in the same directory
or any subdirectories of that directory: If the web page `example.com/catalog/index.html`
creates a cookie, for example, that cookie is also visible to `example.com/catalog/order.html`
and `example.com/catalog/widgets/index.html` but not `example.com/about.html`.

This default visibility behavior is often exactly what you want. To allow
using cookie values throughout a website, regardless of which page creates the cookie,
you specify a `path` for the cookie. Then, any web page from the same web server whose
URL begins with the path prefix you specified can share the cookie.

For example, if a cookie set by `example.com/catalog/widgets/index.html` has its
path set to "/catalog", the cookie is also visible to "example.com/catalog/order.html".

By default, cookies are scoped by **document origin**. Large websites may want
cookies to be shared across subdomains: `order.example.com` wants to read cookie
value set from `catalog.example.com`. If we set cookie `domain` attribute
to ".example.com", then this cookie is available to all web pages end with `.example.com`. Note that you cannot set the domain of a cookie to a domain other than
a parent domain of your server.

![Trying Dot Com](images/trying-dot-com.png)

![Trying Google](images/trying-google.png)

**Important**: Cookies will not be sent to different origin (protocol, hostname, port). To allow receiving & sending cookies by a CORS request (e.g http://localhost:8080 visits http://localhost:8082) successfully, do
the below:

- In server side

  - set `Access-Control-Allow-Credentials` to `true`

  - Note that after setting `Access-Control-Allow-Credentials` to `true`, the Access-Control-Allow-Origin is not allowed to be wildcard

    ![Credentials With CORS](images/credentials-with-cors.png)

  - Set SameSite = None

    ![Samesite None](images/samesite-none.png)

  - Set Secure = true

    ![Secure True](images/secure-true.png)

- In client side, use `{ credentials: 'include' }` in second argument of `fetch()`

`secure` boolean attribute specifies how cookie values are transmitted over the network.
By default, cookies are insecure, which means that they are transmitted over a normal,
insecure HTTP connection. If a cookie is marked secure, however, **it is transmittedi
only when the browser and server are connected via HTTPS or another secure protocol** (except on localhost).

![Secure Cookie](images/secure-cookie.png)

---

> **Cookie Limitations**

Cookies are intended for storage of small amounts of data by server-side scripts,
and that data is transferred to the server each time a relevant URL is requested.
The standard that defines cookies encourages browser manufacturers to allow unlimited
numbers of cookies of unrestricted size but does not require browsers to retain
more than 300 cookies total, 20 cookies per web server, or 4 KB of
data per cookie (both name and value count toward this 4KB limit). In practice,
browser allow many more than 300 cookies total, but the 4 KB size limit may
still be enforced by some.

---

##### Storing Cookies

Cookie values cannot include semicolons, commas, or whitespace. For this reason,
**you may want to use the core JavaScript global function `encodeURIComponent()`
to encode the value before storing it in the cookie**.

The cookies stored by default expires when session end, and the domain is where the current hostname
on webpage. To create a cookie that last across browser session, specify its lifetime
(in seconds) with a `max-age` attribute.

```js
document.cookie = "name=value;max-age=seconds";
document.cookie = "name=value;max-age=60";
```

```js
function setCookie(name, value, daysToLive = null) {
  let cookie = `${name}=${encodeURIComponent(value)}`;
  if (daysToLive !== null) {
    cookie += `; max-age=${daysToLive * 60 * 60 * 24}`;
  }
  document.cookie = cookie;
}
```

Similarly, you can set the `path` and `domain` attributes of a cookie by appending strings
of the form `;path=value` or `;domain=value` to the string. To set the `secure` property,
simply append `;secure`.

To delete a cookie, set it again using the same name, path, and domain,
specifying an arbitrary (or empty) value, and a `max-age` attribute of 0.

#### 15.12.3 IndexedDB

Web application architecture has traditionally featured HTML, CSS, and JavaScript
on the client and a database on the server. You may find it surprising, therefore,
to learn that the web platform includes a simple object database with a JavaScript API
for persistently storing JavaScript objects on the user's computer and retrieving them as needed.

IndexedDB is an object database, not a relational database, and it is much simpler than database that support SQL queries. If is more powerful, efficient, and robust than the
key/value storage provide by the `localStorage`, however. Like the `localStorage` ,
IndexedDB databases **are scoped to origin of the containing document**: two web pages
with the same origin **can access each other's data, but web pages from different
origins cannot**.

Each origin can have any number of IndexedDB databases. Each one has a unique name within
the origin. In the IndexedDB API, a database is simply a collection of named _object stores_. Objects are serialized into the object store using the structured clone algorithm, which means the objects you store **can have properties whose values are Maps, Sets, or
Typed arrays**. Keys must have a natural ordering so that they can be sorted -- JavaScript
string, numbers, and Date objects are valid keys. An IndexedDB database can automatically
generate a unique key for each object you insert into the database.

```js
let open = indexedDB.open("zipcodes", 1); // request v1 of the database
open.onerror = console.error;
open.onsuccess = () => {
  let db = request.result;
  let transaction = db.transaction(["zipcodes"], "readwrite");
  transaction.onerror = console.error;
  let store = transaction.objectStore("zipcodes");

  // put scenario
  store.put(record);
  transaction.oncomplete = () => {
    callback(db);
  };

  // get scenario
  let request = store.get(zip);
  request.onerror = console.error;
  request.onsuccess = () => {
    let record = request.result;
    if (record) {
      callback(`${record.city}, ${record.state}`);
    } else {
      callback(null);
    }
  };
};

open.onupgradeneeded = () => {
  initdb(open.result, callback);
};
```

**IndexedDB provides atomicity guarantees: queries and updates to the database
are grouped within a _transaction_ so that they all succeed together or all fail
together** and never leave the database in an undefined, partially updated state.

```js
let transaction = db.transaction(["zipcodes"], "readwrite");
transaction.onerror = console.error;
transaction.oncomplete = () => {
  callback(db);
};
```

The IndexedDB API is asynchronous (so that web apps can use it without blocking
the browser's main UI thread). IndexedDB was defined before Promises were widely supported,
so the API is event-based rather than Promise-based, which means that it does not
work with `async` and `await`.

Creating transactions, and looking up object stores and indexes are synchronous operations.
But opening a database, updating an object store and querying a store or index are
all asynchronous operations. These asynchronous methods all immediately return a request
object. The browser triggers a success or error event on the request object when
the request succeeds or fails. Another useful event is the "complete" event dispatched
on transaction objects when a transaction has completed successfully.

```js
let transaction = db.transaction(["zipcodes"], "readwrite");
transaction.onerror = console.error;
transaction.oncomplete = () => {
  callback(db);
};
let store = transaction.objectStore("zipcodes");

// put scenario
store.put(record);
transaction.oncomplete = () => {
  callback(db);
};

// get scenario
let request = store.get(zip);
request.onerror = console.error;
request.onsuccess = () => {
  let record = request.result;
  if (record) {
    callback(`${record.city}, ${record.state}`);
  } else {
    callback(null);
  }
};
```

One convenient feature of this asynchronous API is that it simplifies transaction
management. The IndexedDB API forces you to create a transaction object in order
to get the object store on which you can perform queries and updates. In synchronous
API, you would expect to explicitly mark the end of the transaction by calling a `commit()`
method. But with IndexedDB, transactions are automatically committed (if you do not
explicitly abort them) when all the `onsuccess` event handlers have run and there are
no more pending asynchronous requests that refer to that transaction.

There is one more event that is important to the IndexedDB API. When you open a database
for the first time, or when you increment the version number of an existing database.
IndexedDB fires an `upgradeneeded` event on the request object returned by the
`indexedDB.open()` call. The job of the event handler for "upgradeneeded" events is to
define or update the schema for the new database, this means creating object stores
and defining indexes on those object stores. **And in fact, the only time the
IndexedDB API allows you to create an object store or an index is in response to
an "upgradeneeded" event.**

**Tip**: `db.deleteObjectStore(name)` can be only invoked in `request.onupgradeneeded`

**Tip**: `indexedDB.deleteDatabase(name)` will be blocked until all databases close

```js
function withDB(callback) {
  let request = indexedDB.open("zipcodes", 1);
  request.onerror = console.error;
  request.onsuccess = () => {
    let db = request.result;
    callback(db);
  };

  // if version 1 of the database does not yet exist, then this
  // event handler will be triggered
  request.onupgradeneeded = () => {
    initdb(request.result, callback);
  };
}

function initdb(db, callback) {
  let store = db.createObjectStore("zipcodes", { keyPath: "zipcode" }); // keypath specifying the property name of the key field for this store

  // Index the object store by city name as well as by zipcode.
  // With this method the key path string is passed directly as a
  // required argument rather than as part of an options object.
  store.createIndex("cities", "city");

  fetch("zipcodes.json")
    .then((response) => response.json())
    .then((zipcodes) => {
      // Specify which object stores we'll be using and doing
      let transaction = db.transaction(["zipcodes"], "readwrite");
      transaction.onerror = console.error;

      let store = transaction.objectStore("zipcodes");
      for (let record of zipcodes) {
        store.put(record);
      }

      transaction.oncomplete = () => {
        callback(db);
      };
    });
}

// Given a zip code, use the IndexedDB API to asynchronously look up the city
// with that zip code, and pass it to the specified callback, or pass null if
// no city is found
function lookupCity(zip, callback) {
  withDB((db) => {
    let transaction = db.transaction(["zipcodes"]);
    let zipcodes = transaction.objectStore("zipcodes");
    let request = zipcodes.get(zip);
    request.onerror = console.error;
    request.onsuccess = () => {
      let record = request.result;
      if (record) {
        callback(`${record.city}, ${record.state}`);
      } else {
        callback(null);
      }
    };
  });
}

function lookupZipcodes(city, callback) {
  withDB((db) => {
    let transaction = db.transaction(["zipcodes"]);
    let store = transaction.objectStore("zipcodes");
    let index = store.index("cities");
    let request = index.getAll(city);
    request.onerror = console.error;
    request.onsuccess = () => {
      callback(request.result);
    };
  });
}
```

![Indexeddb1](images/indexeddb1.png)

![Indexeddb2](images/indexeddb2.png)

![Indexeddb3](images/indexeddb3.png)

<a id="15.13"></a>

### 15.13 Worker Threads and Messaging

Web browsers very carefully relax the single-thread requirement with the Worker class:
instances of this class represent threads that run concurrently with the main thread and
the event loop. **Workers live in a self-contained execution environment with
a completely independent global object and no access to the Window or Document objects**. They can communicate with the main thread **only through asynchronous message passing**. Creating a new worker is not a heavyweight operation like opening a new browser window,
but they are not flyweight "fibers" either, and it does not make sense to create new
workers to perform trivial operations. Complex web applications may find it useful
to create tens of workers, but it is unlikely that an application with hundreds or
thousands of workers would be practical.

Workers are useful when your application needs to perform computationally intensive
tasks, such as image processing. Workers are also useful when you have to perform frequent
moderately intensive computations. Suppose, for example, that implementing a simple
in-browser code editor and want to include syntax highlighting.

There are two parts to the Worker API. The first is the Worker object: this is what
a worker looks like from the outside, to the thread that creates it. The second is the
WorkerGlobalScope: this is the global object for a new worker.

#### 15.13.1 Worker Objects

To create a new worker:

```js
let dataCruncher = new Worker("utils/cruncher.js");
```

**Tip**: If you specify an absolute URL, **it must have the same origin as that containing document**.

![Origin Worker](images/origin-worker.png)

Once you have a Worker object, you can send data to it with `postMessage()`. The value
you pass to `postMessage()` will be copied using the structured clone algorithm, and
the resulting copy will be delivered to the worker via a message event:

```js
dataCruncher.postMessage("/api/data/to/crunch");
```

You can receive messages from a worker by listening for "message" events:

```js
dataCruncher.onmessage = function (e) {
  let stats = e.data;
  console.log(`Average: ${stats.mean}`);
};
```

In addition to `postMessage()`, Worker objects have just one other method,
`terminate()`, which forces a worker thread to stop running.

```js
dataCruncher.terminate();
```

#### 15.13.2 The Global Object in Workers

When you create a new worker with the `Worker()` constructor, you specify
the URL of a file of JavaScript code. **That code is executed in a new, pristine JavaScript
execution environment, isolated from the script that created the worker**. The
global object for that new execution environment is WorkerGlobalScope object.

The WorkerGlobalScope object has a `postMessage()` method and
an `onmessage` event handler property that are just like those of the Worker object.
Because the WorkerGlobalScope is the global object for a worker, `postMessage()` and
`onmessage` look like a global function and global variable to worker code.

```js
postMessage("hi");
onmessage(console.log);
```

If you pass an object as the second argument to the `Worker()` constructor, and
if that object has a `name` property, then the value of that property becomes the
value of the `name` property in the worker's global object. A worker might include
this name in any messages it prints with `console.warn()` or `console.error()`

```js
console.log(name);
```

The `close()` function allows a worker to terminate itself.

**Tip**: `postMessage` to closed worker does not throw any exception.

WorkerGlobalScope has all of the properties of the core JavaScript global object,
such as the JSON object `isNaN()` function, and the `Date()` constructor. In
addition, however, WorkerGlobalScope also has the following properties of the client-side
Window object:

- `self` is a reference to the global object itself.
- The timer methods `setTimeout()`, `clearTimeout()`, `setInterval()` and `clearInterval()`
- A `location` property that describes the URL (location of worker.js) that was passed to the `Worker()` constructor.
  This property refers to a Location object, just as the `location` property of a Window
  does. However, in a worker, these properties are read-only.
- A `navigator` property that refers to an object with properties like those of the
  Navigator object of a window. A worker's Navigator object has the properties `appName`,
  `appVersion`, `appVersion`, `platform`, `userAgent`, and `onLine`.

  ![Worker Navigator](images/worker-navigator.png)

- The usual event target methods `addEventListener()` and `removeEventListener()`

And Console object, the `fetch()` function, and the IndexedDB API. And `Worker()` constructor
which means that worker threads can create their own workers.

#### 15.13.3 Importing Code Into a Worker

Workers were defined in web browsers before JavaScript had a module system, so workers have a unique
system for including additional code.
WorkerGlobalScope defines `importScripts()` as a global function that all workers
have access to:

```js
importScripts("utils/Histograms.js", "utils/BitSet.js");
```

**Important**: Relative URLs are resolved relative to the URL that was passed to the
`Worker()` constructor (not relative to the containing document).

`importScripts()` synchronously loads and executes these files one after the other,
in the order in which they were specified. If loading a script causes an error,
none of the subsequent scripts are loaded or executed (nor the script after `importScript()` in Worker. A script loaded with `importScripts()` can itself call `importScripts()`
to load the files it depends on. **Note, however, that `importScripts()` does not
try to keep track of what scripts have already loaded and does nothing to prevent
dependency cycles** -- unlike JavaScript module.

---

> **Modules In Workers**

In order to use modules in workers, you must pass a second argument to the `Worker()`
constructor. This second argument must be an object with a `type` property set to the string
"module" -- so that the `import` declarations are allowed.

When a worker loads a module instead of a traditional script, **the WorkerGlobalScope
does not define the `importScripts()` function**.

**Reminder**: `import` are hoisted at top

**Reminder**: import of no export script

```js
import "xxx.js";
```

---

#### 15.13.4 Worker Execution Model

If a worker registers a "message" event handler, **it will never exit as long as
there is a possibility that message events will still arrive.** But if a
worker doesn't listen for messages, it will run until there are no further pending
tasks (such as `fetch()` promises and timers) and all task-related callbacks
have been called. A worker can also explicitly shut itself down by calling
the global `close()` function. Note that there are no properties or methods
on the Worker object that specify whether a worker thread is still running or not,
so workers should not close themselves without somehow coordinating this with
their parent thread.

##### Errors in Workers

If an exception occurs in a worker and is not caught by any `catch` clause, then
an "error" event is triggered on the **global object of the worker**, and propagates to Worker object.
The error propagation ends when we calls the `preventDefault()` in **Worker object** (global object of the worker does not have this method).

```js
// Worker
self.onerror = function (e) {
  console.log(e);
};

// Outside the worker
worker.onerror = function (e) {
  console.log(`Error in worker at ${e.filename}:${e.lineno}:${e.message}`);
  e.preventDefault();
};
```

Like windows, workers can register a handler to be invoked when a Promise is
rejected and there is no `.catch()` function to handle in (onerror won't catch rejected Promise)
Within a worker you can detect this by defining a `self.onunhandledrejection` function
or by using `addEventListener()` to register an `unhandledrejection` event.
The event object to this handler will have a `promise` property whose value is the
Promise object that rejected and a `reason` property whose value is what would have been
passed to a `.catch()` function.

```js
self.onunhandledrejection = (promise) => {
  console.log(promise);
  console.log(promise.reason);
  promise.preventDefault();
};
```

<a id="15.13.5"></a>

#### 15.13.5 `postMessage()`, MessagePorts, and MessageChannels

The `postMessage()` method of the Worker object and the global `postMessage()`
function defined inside a worker both work by invoking the `postMessage()` methods
of a pair of MessagePort objects that are automatically created along with the worker.
Clientside JavaScript can't directly access these automatically created MessagePort
objects, but it can create new pairs of connected ports with `MessageChannel()` constructor.

```js
let channel = new MessageChannel();
let myPort = channel.port1;
let yourPort = channel.port2;

myPort.postMessage("Can you hear me?");
yourPort.onmessage = (e) => console.log(e.data);
```

**Tip**: Messages sent to a port are **queued until** the `onmessage` property is defined
or until the `start()` method is called on the port. **It is important to use `start()`
method if you use `addEventListener()` method.**

```js
yourPort.addEventListener("message", (e) => console.log(e.data));
yourPort.start();
```

All the `postMessage()` calls we've seen so far have taken a single message
argument. But the method also accepts an optional second argument:
**an array of items that are to be transferred to the other end of the channel
instead of having a copy sent across the channel**. Values that can be **transferred
instead of copied** are MessagePorts and ArrayBuffers. If the first argument to `postMessage()`
includes a MessagePort, **then that MessagePort must also appear in the second argument**.
If you do this, then the MessagePort will become available to the other end of the channel
**and will immediately become non functional on your end (it will not cause any error though)**.

Below is the error if the second argument does not include the message port.

![Failed Sending Post Message](images/failed-sending-post-message.png)

Below shows the example if we created a worker and want to have two channels
for communicating with it: one channel for ordinary data exchange and one channel
for high-priority message.

```js
let worker = new Worker("testmultiport1.js");

let urgentChannel = new MessageChannel();
let urgentPort = urgentChannel.port1;
worker.postMessage({ command: "SetUrgentPort", value: urgentChannel.port2 }, [
  urgentChannel.port2,
]);

urgentPort.onmessage = handlerUrgentMessage;
function handlerUrgentMessage(e) {
  console.log(e.data);
}

// Send urgent messages like this
urgentPort.postMessage("test");

// Implement a low-priority message here
```

```js
// Worker
onmessage = (e) => {
  let urgentPort = e.data.value;
  urgentPort.onmessage = (e) => {
    console.log(e.data);
  };
  urgentPort.postMessage("connected");
};
```

Message Channels are also useful if you create two workers and want to allow them
to communicate directly with each other rather than requiring code on the
main thread to relay messages between them.

```js
// Main Thread
let worker1 = new Worker("./scripts/workerconnect1.js", { name: "Worker1" });
let worker2 = new Worker("./scripts/workerconnect1.js", { name: "Worker2" });
let channel = new MessageChannel();
worker1.postMessage({ command: "SetMessagePort", value: channel.port1 }, [
  channel.port1,
]);
worker2.postMessage({ command: "SetMessagePort", value: channel.port2 }, [
  channel.port2,
]);

// In worker:
onmessage = (e) => {
  let messagePort = e.data.value;
  messagePort.onmessage = (e) => {
    console.log(name + " receive: ", e.data);
  };
  setTimeout(() => {
    messagePort.postMessage("hello");
  }, 1000);
};
```

The other use of the second argument to `postMessage()` is to **transfer
ArrayBuffers between workers without having to copy them**. This is an
important performance enhancement for large ArrayBuffers like those used
to hold image data. When an ArrayBuffer is transferred over a MessagePort,
the ArrayBuffer becomes unusable in the original thread so there is no
possibility of concurrent access to its contents.
If the first argument to `postMessage()` includes an Array Buffer,
then that buffer may appear as an array element in the second `postMessage()`
argument. If it does appear, **then it will be transferred without copying.**
Else, it will be copied rather than transferred.

[See The Mandelbrot Set example for demonstration](#15.14)

<a id="15.13.6"></a>

#### 15.13.6 Cross-Origin Messaging with `postMessage()`

There is another use case for the `postMessage()` method in client side
JavaScript. It involves windows instead of workers, but there are
enough similarities between the two cases that we will describe the `postMessage()`
method of the Window object here.

When a document contains an `<iframe>` element, the element acts as
an embedded but independent window. The Element object that represents
the `<iframe>` has a `contentWindow` property that is the Window object
for the embedded document. And for the scripts running within
that nested iframe, the `window.parent` property refers to the containing Window object.
When two windows display documents with the same origin, then scripts in each
of those windows have access to the contents of the other window.
But when the documents have different origins, the browser's same-origin policy
prevents JavaScript in one window from accessing the content of another window.

![Iframe Cross Origin](images/iframe-cross-origin.png)

Example of usage:

```js
// main
let iframe = window.querySelector("iframe");
console.log(iframe.contentWindow);

// iframe
console.log(window.parent);
```

Even if the same-origin policy prevents your script from seeing the content of another
window, you can still call `postMessage()` on that window, and doing so will
cause a "message" event to be triggered on that window, where it can be seen by
the event handlers in that window's scripts.

The `postMessage()` method of a Window is a little different than the
`postMessage()` method of a Worker.

- The first argument is still an arbitrary message that will be copied by the
  structured clone algorithm

- An origin that specifies who you expect to be receiving the message

  - If you pass the string "https://good.example.com", but the window you are
    posting the message to actually contains the content from `https://malware.example.com`
    then the message you posted will not be delivered.

  - ![Cross Origin Iframe With Post Message](images/cross-origin-iframe-with-post-message.png)

  - If you are willing to send your message to content from any origin, then you can pass
    the wildcard "\*".

- The optional third argument listing objects to be transferred instead of copied

Besides `data` field in `onmessage` event, addition properties `source`, and `origin`
are defined.

- `source` - specifies the Window object that sent the event,
  you can use e.source.postMessage() to send a reply

- `origin` - specifies the origin of the content in the source window.
  **When you receive a "message" event, you will typically want to verify that it is from
  an origin you expect**.

```js
// main
let frame = document.querySelector("iframe");
window.onmessage = (e) => {
  console.log(e.data);
};
setTimeout(() => {
  frame.contentWindow.postMessage(
    "a message from parent",
    "http://localhost:8080",
  );
}, 1000);

// iframe
window.onmessage = (e) => {
  console.log(e.data);
  // Do origin verification here
  console.log(e.origin);
  e.source.postMessage("from child");
};
```

<a id="15.14"></a>

### 15.14 Example: The Mandelbrot Set

Worker:

```js
onmessage = function (message) {
  const { tile, x0, y0, perPixel, maxIterations } = message.data;
  const { width, height } = tile;
  const imageData = new ImageData(width, height);
  const iteration = new Uint32Array(imageData.data.buffer);
  let index = 0,
    max = 0,
    min = maxIterations;
  for (let row = 0, y = y0; row < height; row++, y += perPixel) {
    for (let column = 0, x = x0; column < width; column++, x += perPixel) {
      let n;
      let r = x,
        i = y;
      for (n = 0; n < maxIterations; n++) {
        let rr = r * r,
          ii = i * i;
        if (rr + ii > 4) {
          break;
        }
        i = 2 * r * i + y;
        r = rr - ii + x;
      }
      iterations[index++] = n;
      if (n > max) max = n;
      if (n < min) min = n;
    }
  }

  postMessage({ tile, imageData, min, max }, [imageData.data.buffer]);
};
```

Web application for displaying and exploring the Mandelbrot set

```js
class Tile {
  constructor(x, y, width, height) {
    this.x = x;
    this.y = y;
    this.width = width;
    this.height = height;
  }

  static *tiles(width, height, numRows, numCols) {
    let columnWidth = Math.ceil(width / numCols);
    let rowHeight = Math.ceil(height / numRows);

    for (let row = 0; row < numRows; row++) {
      let tileHeight =
        row < numRows - 1 ? rowHeight : height - rowHeight * (numRows - 1);
      for (let col = 0; col < numCols; col++) {
        let tileWidth =
          col < numCols - 1 ? columnWidth : width - columnWidth * (numCols - 1);
        yield new Tile(
          col * columnWidth,
          row * rowHeight,
          tileWidth,
          tileHeight,
        );
      }
    }
  }
}

class WorkerPool {
  constructor(numWorkers, workerSource) {
    this.idleWorkers = [];
    this.workQueue = [];
    this.workerMap = new Map();

    for (let i = 0; i < numWorkers; i++) {
      let worker = new Worker(workerSource);
      worker.onmessage = (message) => {
        this._workerDone(worker, null, message.data);
      };
      worker.onerror = (error) => {
        this._workerDone(worker, error, null);
      };
      this.idleWorkers[i] = worker;
    }
  }

  _workerDone(worker, error, response) {
    let [resolver, rejector] = this.workerMap.get(worker);
    this.workerMap.delete(worker);
    // If there is no queued work, put this worker back
    // in the list of idle workers. Otherwise, take work
    // from the queue and send it to this worker
    if (this.workQueue.length === 0) {
      this.idleWorkers.push(worker);
    } else {
      let [work, resolver, rejector] = this.workQueue.shift();
      this.workerMap.set(worker, [resolver, rejector]);
      worker.postMessage(work);
    }
    error === null ? resolver(response) : rejector(error);
  }

  addWork(work) {
    return new Promise((resolve, reject) => {
      if (this.idleWorkers.length > 0) {
        let worker = this.idleWorkers.pop();
        this.workerMap.set(worker, [resolve, reject]);
        worker.postMessage(work);
      } else {
        this.workQueue.push([work, resolve, reject]);
      }
    });
  }
}

class PageState {
  static initialState() {
    let s = new PageState();
    s.cx = -0.5;
    s.cy = 0;
    s.perPixel = 3 / window.innerHeight;
    s.maxIterations = 500;
    return s;
  }

  static fromURL(url) {
    let s = new PageState();
    let u = new URL(url);
    s.cx = parseFloat(u.searchParams.get("cx"));
    s.cy = parseFloat(u.searchParams.get("cy"));
    s.perPixel = parseFloat(u.searchParams.get("pp"));
    s.maxIterations = parseFloat(u.searchParams.get("it"));
    return isNaN(s.cx) ||
      isNaN(s.cy) ||
      isNaN(s.perPixel) ||
      isNaN(s.maxIterations)
      ? null
      : s;
  }

  toURL() {
    let u = new URL(window.location);
    u.searchParams.set("cx", this.cx);
    u.searchParams.set("cy", this.cy);
    u.searchParams.set("pp", this.perPixel);
    u.searchParams.set("it", this.maxIterations);
    return u.href;
  }
}

const ROWS = 3,
  COLS = 4,
  NUMWORKERS = navigator.hardwareConcurrency || 2;

class MandelbrotCanvas {
  constructor(canvas) {
    this.canvas = canvas;
    this.context = canvas.getContext("2d");
    this.workerPool = new WorkerPool(NUMWORKERS, "mandelbrotWorker.js");
    this.tiles = null;
    this.pendingRender = null;
    this.wantsRerender = false;
    this.resizeTimer = null; // prevents us from resizing too frequently
    this.colorTable = null; // For converting raw data to pixel values

    this.canvas.addEventListener("pointerdown", (e) => {
      this.handlePointer(e);
    });
    window.addEventListener("keydown", (e) => {
      this.handleKey(e);
    });
    window.addEventListener("resize", (e) => {
      this.handleResize(e);
    });
    window.addEventListener("popstate", (e) => {
      this.setState(e.state, false);
    });

    this.state = PageState.fromURL(window.location) || PageState.initialState();
    history.replaceState(this.state, "", this.state.toURL());
    this.setSize();
    this.render();
  }

  setSize() {
    this.width = this.canvas.width = window.innerWidth;
    this.height = this.canvas.height = window.innerHeight;
    this.tiles = [...Tile.tiles(this.width, this.height, ROWS, COLS)];
  }

  setState(f, save = true) {
    if (typeof f === "function") {
      f(this.state);
    } else {
      for (let property in f) {
        this.state[property] = f[property];
      }
    }
    this.render();

    if (save) {
      history.pushState(this.state, "", this.state.toURL());
    }
  }

  render() {
    if (this.pendingRender) {
      this.wantsRerender = true;
      return;
    }

    let { cx, cy, perPixel, maxIterations } = this.state;
    let x0 = cx - (perPixel * this.width) / 2;
    let y0 = cy - (perPixel * this.height) / 2;

    let promises = this.tiles.map((tile) => {
      return this.workerPool.addWork({
        tile: tile,
        x0: x0 + tile.x * perPixel,
        y0: y0 + tile.y * perPixel,
        perPixel: perPixel,
        maxIterations: maxIterations,
      });
    });
    this.pendingRender = Promise.all(promises)
      .then((responses) => {
        let min = maxIterations,
          max = 0;
        for (let r of responses) {
          if (r.min < min) min = r.min;
          if (r.max > max) max = r.max;
        }

        if (!this.colorTable || this.colorTable.length !== maxIterations + 1) {
          this.colorTable = new Uint32Array(maxIterations + 1);
        }

        if (min === max) {
          if (min === maxIterations) {
            this.colorTable[min] = 0xff000000;
          } else {
            this.colorTable[min] = 0;
          }
        } else {
          let maxlog = Math.log(1 + max - min);
          for (let i = min; i <= max; i++) {
            this.colorTable[i] =
              Math.ceil((Math.log(1 + i - min) / maxlog) * 255) << 24;
          }
        }

        for (let r of responses) {
          let iterations = new Uint32Array(r.imageData.data.buffer);
          for (let i = 0; i < iterations.length; i++) {
            iterations[i] = this.colorTable[iterations[i]];
          }
        }

        this.canvas.style.transform = "";
        for (let r of responses) {
          this.context.putImageData(r.imageData, r.tile.x, r.tile.y);
        }
      })
      .catch((reason) => {
        console.error("Promise rejected in render():", reason);
      })
      .finally(() => {
        this.pendingRender = null;
        if (this.wantsRerender) {
          this.wantsRerender = false;
          this.render();
        }
      });
  }

  handleResize(event) {
    if (this.resizeTimer) clearTimeout(this.resizeTimer);
    this.resizeTimer = setTimeout(() => {
      this.resizeTimer = null;
      this.setSize();
      this.render();
    }, 200);
  }

  handleKey(event) {
    switch (event.key) {
      case "Escape":
        this.setState(PageState.initialState());
        break;
      case "+":
        this.setState((s) => {
          s.maxIterations = Math.round(s.maxIterations * 1.5);
        });
        break;
      case "-":
        this.setState((s) => {
          s.maxIterations = Math.round(s.maxIterations / 1.5);
          if (s.maxIterations < 1) s.maxIterations = 1;
        });
        break;
      case "o":
        this.setState((s) => (s.perPixel *= 2));
        break;
      case "ArrowUp":
        this.setState((s) => (s.cy -= (this.height / 10) * s.perPixel));
        break;
      case "ArrowDown":
        this.setState((s) => (s.cy += (this.height / 10) * s.perPixel));
        break;
      case "ArrowLeft":
        this.setState((s) => (s.cx -= (this.width / 10) * s.perPixel));
        break;
      case "ArrowRight":
        this.setState((s) => (s.cx += (this.width / 10) * s.perPixel));
        break;
    }
  }

  handlePointer(event) {
    const x0 = event.clientX,
      y0 = event.clientY,
      t0 = Date.now();
    const pointerMoveHandler = (event) => {
      let dx = event.clientX - x0,
        dy = event.clientY - y0,
        dt = Date.now() - t0;

      if (dx > 10 || dy > 10 || dt > 500) {
        this.canvas.style.transform = `translate(${dx}px, ${dy}px)`;
      }
    };

    const pointerUpHandler = (event) => {
      this.canvas.removeEventListener("pointermove", pointerMoveHandler);
      this.canvas.removeEventListener("pointerup", pointerUpHandler);

      const dx = event.clientX - x0,
        dy = event.clientY - y0,
        dt = Date.now() - t0;
      const { cx, cy, perPixel } = this.state;
      if (dx > 10 || dy > 10 || dt > 500) {
        this.setState({ cx: cx - dx * perPixel, cy: cy - dy * perPixel });
      } else {
        let cdx = x0 - this.width / 2;
        let cdy = y0 - this.height / 2;
        this.canvas.style.transform = `translate(${-cdx * 2}px, ${
          -cdy * 2
        }px) scale(2)`;
        this.setState((s) => {
          s.cx += cdx * s.perPixel;
          s.cy += cdy * s.perPixel;
          s.perPixel /= 2;
        });
      }
    };

    this.canvas.addEventListener("pointermove", pointerMoveHandler);
    this.canvas.addEventListener("pointerup", pointerUpHandler);
  }
}

let canvas = document.createElement("canvas");
document.body.append(canvas);
document.body.style = "margin:0";
canvas.style.width = "100%";
canvas.style.height = "100%";
new MandelbrotCanvas(canvas);
```

### 15.15 Suggestion for Further Reading

You can't learn about a new API if you don't know that it exists, so
the short sections that follow end the chapter with a quick list of
web platform features that you might want to investigate in the future.

#### 15.15.1 HTML and CSS

Two topics worth paying particular attention to in this area are
accessibility (including ARIA attributes) and internationalization
(including support for right-to-left writing directions)

#### 15.15.2 Performance

It is hard to optimize things that you can't measure, however, so it is worth familiarizing
yourself with the Performance API. The `performance` property of the window object
is the main entry point to this API. It includes a high-resolution time source `performance.now()` and methods `performance.mark()` and `performance.measure()` for marking
critical points in your code and measuring the elapsed time between time. Calling
these methods creates PerformanceEntry objects that you can access with `performance.getEntries()`. Browsers add their own PerformanceEntry objects any time the browser loads
a new page or fetches a file over the network, and these automatically create PerformanceEntry object
include granular timing details of your application's network performance.
The related PerformanceObserver class allows you to specify a function to be invoked
when new PerformanceEntry objects are created.

#### 15.15.3 Security

You may want to spend some time learning more about web security. In addition
to XSS, it is worth learning about the `Content-Security-Policy` HTTP header and
understanding how CSP allows you to ask the web browser to restrict the capabilities it
grants to JavaScript code. Understanding CORS (Cross-Origin Resource Sharing) is
also important.

#### 15.15.4 WebAssembly

WebAssembly (or "wasm") is a low-level virtual machine bytecode
format that is designed to integrate well with JavaScript interpreters in
web browsers. There are compilers that allow you to compile C, C++,
and Rust programs to WebAssembly bytecode and to run those programs
in web browsers at close to native speed, without breaking the browser sandbox
or security model. WebAssembly can export functions that can be called by
JavaScript programs.

A typical use case for it would be to compile the standard C-language zlib compression
library so that JavaScript code has access to high-speed compression and decompression
algorithms.

#### 15.15.5 More Document and Window Features

Window:

- `alert()`, `confirm()`, and `prompt()` methods that display simple modal
  dialogues to the user. These methods block the main thread. These are not suitable
  for production use but can be useful for simple projects and prototypes.

- `navigator` and `screen` have some features that were not described here that you may
  find useful.

- `requestAnimationFrame()` method takes a function as its argument and will
  execute that function when the browser is preparing to render the next frame.
  When you are making visual changes (especially repeated or animated ones),
  wrapping your code within a call to `requestAnimationFrame()` can help to
  ensure that the changes are rendered smoothly and in a way that is
  optimized by the browser.

Element:

- The `requestFullScreen()` method of any Element object requests that that
  element (a `<video>` or `<canvas>` element, for example) be displayed in fullscreen mode

- If the user selects text within your document, you can obtain details of that selection with
  the Window method `getSelection()` and get the selected text with `getSelection().toString()`.
  In some browsers, `navigator.clipboard` is an object with an async API for reading
  and setting the content of the system clipboard to enable copy-and-paste interactions
  with applications outside of the browser.

  ```js
  setInterval(() => {
    console.log(getSelection().toString());
  }, 1000);
  ```

- HTML elements with a `contenteditable="true"` attribute allow their content to be edited.
  The `document.execCommand()` method enables rich-text editing features for editable content.

- A MutationObserver allow JavaScript to monitor changes to, or beneath,
  a specified element in the document. Create a MutationObserver with the `MutationObserver()` constructor, passing the callback function that should be called when changes
  are made. Then call the `observe()` method of the MutationObserver to specify
  which parts of which elements are to be monitored.

- An IntersectionObserver allows JavaScript to determine which document elements
  are on the screen and which are close to being on the screen. It is particularly
  useful for applications that want to dynamically load content on the demand as the user scrolls.

#### 15.15.6 Events

Some more events that you may find useful:

- browser fires "online" and "offline" at the Window object
- browser fires a "visibilitychange" event at the Document object when
  a document becomes visible or invisible (usually because a user has switched tabs).
  JavaScript can check `document.visibilityState` to determine whether its document
  is currently "visible" or "hidden"

  ```js
  document.addEventListener("visibilitychange", () => {
    let title = document.querySelector("title");
    if (document.visibilityState === "visible") {
      title.textContent = "You are here!";
    } else {
      title.textContent = "Please come back :(";
    }
  });
  ```

- browser supports a complicated API to support drag-and-drop UIs and
  to support data exchange with applications outside the browser.
  This API involves a number of events, including "dragstart", "dragover",
  "dragend", and "drop". This API is tricky to use correctly but
  useful when you need it. It is an important API to know if you want
  to
  enable users to drag files from their desktop into your web application.

- The Pointer Lock API enables JavaScript to hide the mouse pointer and get
  raw mouse events as relative movement amounts rather than absolute positions
  on the screen. This is typically useful for games. Call `requestPointerLock()`
  on the element if you want all mouse events directed to. After you do this,
  "mousemove" events delivered to that element will have `movementX` and `movementY`
  properties

- The Gamepad API adds support for game controllers. Use
  `navigator.getGamepads()` to get connected Gamepad object, and listen for
  "gamepadconnected" events on the Window object to be notified when a new controller
  is plugged in. The Gamepad object defines an API for querying the current state of
  the buttons on the controller.

#### 15.15.7 Progressive Web Apps and Service Workers

The term Progressive Web Apps, or PWAs, is a buzzword that
describes web applications that are built using a few key technologies.
It is worth noting that powerful modern APIs like these are typically
**designed to work only on secure HTTPS connections**.

- A ServiceWorker is kind of worker thread with the ability to
  intercept, inspect, and respond to network requests from the
  web application that it "services". When a web application
  registers a service worker, the worker's code becomes persistent
  in the browser's local storage, and when the user visits the associated
  website again, the service worker is reactivated. Service workers can
  cache network responses, which means that web applications that use
  service workers can effectively install themselves onto the user's
  computer **for rapid startup and offline use**. The Service Worker
  Cookbook at https://serviceworke.rs is a valuable resource for learning
  about service workers and their related technologies.

- The CacheAPI is designed for use by service workers (but is also available
  to regular JavaScript code outside of workers).
  It works with the Request and Response objects defined by the `fetch()`
  API and implements a cache of Request/Response pairs. The
  Cache API enables a service worker to cache the scripts and other assets
  of the web app it serves and can also help to enable offline use of the web app
  (which is particularly important for mobile devices)

- A Web Manifest is a JSON-formatted file that describes a web application including
  a name, a URL, and links to icons in various sizes.
  If your web app uses a service worker and includes a `<link rel="manifest">` tag
  references a `.webmanifest` file, then browsers (particularly browsers on
  mobile devices) may give you the option to **add an icon** for the webapp to
  your desktop or home screen. You may refer MDN for reference.

- The Notifications API allows web apps to display notifications using the
  native OS notification system on both mobile and desktop devices. Notifications
  can include an image and text, and your code can receive an event if the user
  clicks on the notification. Using this API is complicated by the fact that you
  must first request the user's permission to display notifications.

- The Push API allows web applications that have a service worker (and that
  have the user's permission) to subscribe to notifications from a server,
  and to display those notifications even when the application itself is not running.
  Push notifications are common on mobile devices, and the Push API brings web apps
  closer to feature parity with native apps on mobile.

#### 15.15.8 Mobile Device APIs

There are a number of web APIs that are primarily useful for web apps
running on mobile devices. (Unfortunately, a number of these APIs only work
on Android devices and not iOS devices.)

- The Geolocation API allow JavaScript (with the user's permission)
  to determine the user's physical location. It is well supported on desktop
  and mobile devices, including iOS devices. Use `navigator.geolocation.getCurrentPosition()` to request the user's current position and use `navigator.geolocation.watchPosition()`
  to register a callback to be called when the user's position changes.

- The `navigator.vibrate()` method causes a mobile device (but not iOS) to vibrate.
  Often this is only allowed in response to a user gesture, but calling this method will
  allow your app to provide silent feedback that a gesture has been recognized.

- The ScreenOrientation API enables a web application to query the current orientation
  of a mobile device screen and also to lock themselves to landscape or portrait orientation.

- The "devicemotion" and "deviceorientation" event on the window object report
  report accelerometer and magnetometer data for the device, enabling you to determine how
  the device is accelerating and how the user is orienting it in space. (These events do
  work on iOS.)

- The Sensor API is not yet widely supported beyond Chrome on Android devices, but
  it enables JavaScript access to the full suite of mobile device sensors, including
  accelerometer, gyroscope, magnetometer, and ambient light sensor. These
  sensors enable JavaScript to determine which direction a user is facing or to detect
  when the user shakes their phone, for example.

#### 15.15.9 Binary APIs

For security reasons, JavaScript can't just read local files. But if the user
selects a file for upload (using a `<input type="file">` form element) or use
drag-and-drop to drop a file into your web application, then JavaScript can access that
file as a File object.

File is a subclass of Blob, and as such, it is opaque representation of a chunk of data.
You can use a FileReader class to asynchronously get the content of a file
as a ArrayBuffer or string. (In some browsers, you can skip the FileReader and instead
use the Promise-based `text()` and `arrayBuffer()` methods defined by the
Blob class, or the `stream()` method for streaming access to the file contents)

**When working with binary data, especially streaming binary data, you may need
to decode bytes into text or encode text as bytes**. The
TextEncoder and TextDecoder classes help with this task.

#### 15.15.10 Media APIs

The `navigator.mediaDevices.getUserMedia()` (HTTP only) function allows JavaScript to request access
to the user's microphone and/or video camera. A successful request results in a
MediaStream object.

Video streams can be displayed in a `<video>` tag (by setting the `srcObject` property
to the stream). Still frames of the video can be captured into an offscreen `<canvas>`
with the canvas `drawImage()` function resulting in a relatively low-resolution
photograph. Audio and video streams returned by `getUserMedia()` can be recorded and
encoded to a Blob with a Media Recorder object.

The more complex WebRTC API enables the transmission and reception of MediaStreams
over the network, enabling peer-to-peer video conferencing, for example.

#### 15.15.11 Cryptography and Related APIS

The `crypto` property of the Window object exposes a `getRandomValues()`
method for cryptographically secure pseudorandom numbers.

Other methods for encryption, decryption, key generation, digital signatures, and
so on are available through `crypto.subtle`. The name of this property is
a warning to everyone who uses these methods that properly using
cryptographic algorithms is difficult and that you should not use those
methods unless you really know what you are doing.
Also, the methods of `crypto.subtle` are only available to JavaScript code running
within documents that were loaded over a secure HTTPs connection.

The Credential Management API and the Web Authentication API allow JavaScript
to generate, store, and retrieve public key (and other types of) credentials
and enables account creation and login without passwords. The JavaScript API
consists primarily of the functions `navigator.credentials.create()` and
`navigator.credentials.get()`, but substantial infrastructure is required
on the server side to make these methods work. These APIs
are not universally supported yet, but have the potential to revolutionize the
way we login to websites.

The Payment Request API adds browser support for making credit card payments
on the web. It allows users to store their payment details securely in the
browser so that they don't have to type their credit card number
each time they make a purchase. Web applications that want to request a payment
create a PaymentRequest object and call its `show()` method to display the request
to users.

## 16 - Server-Side JavaScript with Node

### 16.1 - Node Programming Basics

#### 16.1.1 - Console Output

`console.error()` is distinguished from output displayed with `console.log()`,
because `console.error()` writes to the stderr stream.

#### 16.1.2 - Command-Line Arguments and Environment Variables

Node program can read its command-line arguments from the array of strings
`process.argv`. The first element of this array is always the path to the Node
executable. The second argument is the path to the file of JavaScript code that Node
is executing. Any remaining elements in this array are the space-separated arguments
that you passed on the command-line when you invoke Node.

```js
console.log(process.argv);
```

```bash
$ node --trace-uncaught argv.js --arg1 --arg2 filename
[
  '/home/soon/.nvm/versions/node/v18.15.0/bin/node',
  '/home/soon/Repos/learning/javascript-the-definitive-guide/server/argv.js',
  '--arg1',
  '--arg2',
  'filename'
]
```

There are a couple of things to note here:

- Command-line arguments that are intended for and interpreted by the Node executable itself
  are consumed by the Node executable and do not appear in `process.argv` (e.g. --trace-uncaught). **Any arguments that appear after the name of JavaScript file will appear in
  `process.argv`**.

Node programs can also take input from Unix-style environment variables. Node
makes these available though the `process.env`. **The
property values are always strings.**

```bash
# Partial list of environment variable on my system.
$ node -p -e 'process.env'
{
  PULSE_SERVER: 'unix:/mnt/wslg/PulseServer',
  PAGER: 'less',
  SHLVL: '2',
  NVM_INC: '/home/soon/.nvm/versions/node/v18.15.0/include/node',
  DISPLAY: ':0',
  P9K_SSH: '0'
}
```

#### 16.1.3 Program Life Cycle

Fundamentally, Node executes the JavaScript code in the specified file from top to bottom.
Some Node programs exit when they are done executing the last line of code in the file.
Often, however, a Node program will keep running long after the initial file has been
executed.

As we'll discuss in the following sections, Node programs are often asynchronous and based on callbacks and event handlers.
Node programs do not exit until they are done running the initial file and until
all callbacks have been called and there are no more pending events. A Node-based
server program that listens for incoming network connections will theoretically run
forever because **it will always be waiting for more events**.

A program can force itself to exit by calling `process.exit()`. A program
can ignore Ctrl-C by registering a signal handler function with `process.on("SIGINT", () => {})`

If code in your program throws an exception and no `catch` clause catches it,
the program will print a stack trace and exit. Because of Node's asynchronous nature,
exceptions that occur in **callbacks or event handlers** must be handled locally or not handled at all.

If you don't want these exceptions to cause your program to completely crash, register
a global handler function that will be invoked instead of crashing:

```js
process.setUncaughtExceptionCaptureCallback((e) => {
  console.error("Uncaught exception:", e);
});
```

A similar situation arises if a Promise created by your program is rejected and there
is no `.catch()` invocation to handle it. As of Node 13, this is not a fatal
error that causes your program to exit. In some future version of Node, unhandled
Promise rejections are expected to become fatal errors.

If you do not want unhandled rejections, to print error messages or terminate
your program, register a global handler function:

```js
process.on("unhandledRejection", (reason, promise) => {
  // reason is whatever value would have been passed to a .catch() function
  // promise is the Promise object that rejected
});
```

**Tip**: if the `unhandledRejection` is not registered, then it will trigger
the setUncaughtExceptionCaptureCallback, this is the example message:

```js
Uncaught exception: [UnhandledPromiseRejection: This error originated either by
throwing inside of an async function without a catch block, or by rejecting a pr
omise which was not handled with .catch(). The promise rejected with the reason
"hello".] {
  code: 'ERR_UNHANDLED_REJECTION'
}
```

#### 16.1.4 Node Modules

Node needs to know -- before it loads a module -- whether that module will be using
`require()` and `module.exports` or if it will be using `import` and `export` (visit [Chapter 10](#10) for more details).

The simplest way to tell Node what kind of module it is loading is to encode
this information in the file extension. If you save your JavaScript code in a
file ends with `.mjs`, then Node will always load it as an ES6 module. If you save
your code in a file that ends with `.cjs`, then Node will always treat it as
a CommonJS module.

For files that do not have an explicit `.mjs` or `.cjs` extension, Node looks
for a file named `package.json` in the same directory as the file and then in each
of the containing directories. Once the nearest `package.json` file is found,
Node checks for a top-level `type` property in the JSON object.

```json
// Common JS
"type": "commonjs"

// ES6
"type": "module"
```

When no such file is found, or no `type` property is found, Node defaults to using
CommonJS modules.

Note that you do not need to have a `package.json` file to run Node programs. This
`package.json` trick only becomes necessary if you want to use ES6 modules
with Node and do not want to use the `.mjs` file extension.

Because there is an enormous amount of existing Node code written
using CommonJS module format, Node allows ES6 modules to load CommonJS
modules using the `import` keyword. The reverse it not true, however:
a CommonJS module cannot use `require()` to load an ES6 module.

#### 16.1.5 The Node Package Manager

When you install Node, you typically get a program named `npm` as well.
This is the Node Package Manager, and it helps you download and manage libraries that
your program depends on. npm keeps track of those dependencies (as well as other
information about your program) in a file named `package.json` in the root directory
of your project.

To get a project started, you might create a directory for your project, and then,
in that directory type `npm init`.

To start using an external package (e.g Express), you'd type `npm install express`.
This tells npm to download the Express library along with all of its dependencies
and install all the packages **in a local `node_modules/directory`**

You could give another programmer a copy of your code and your `package.json`,
and they could simply type `npm install` to automatically download and install
all of the libraries that your program needs in order to run.

### 16.2 Node Is Asynchronous by Default

Node achieves high levels of concurrency while maintaining a single-threaded programming
model by making its API asynchronous and nonblocking by default. Node takes its nonblocking
approach very seriously and to an extreme that may surprise you.

Some functions in Node's API are synchronous but non blocking: they run to completion
and return without ever needing to block. But most of the interesting functions
perform some kind of input or output, and these are asynchronous functions
so they can avoid even the tiniest amount of blocking.

Node was created before JavaScript had a Promise class, so asynchronous Node
APIs are callback-based. Generally, the last argument you pass to an asynchronous
Node function is a callback. Node uses `error-first` callbacks, which are typically
invoked with two arguments. The first argument to an error-first callback is normally `null`
in the case where no error occurred, and the second argument is whatever data or response
was produced by the original asynchronous function you called.

The following code demonstrates how to use the nonblocking `readFile()` function
to read a configuration file, parse it as JSON, and then pass the parsed configuration
object to another callback:

```js
const fs = require("fs");

function readConfigFile(path, callback) {
  fs.readFile(path, "utf8", (err, text) => {
    if (err) {
      console.error(err);
      callback(null);
      return;
    }
    let data = null;
    try {
      data = JSON.parse(text);
    } catch (e) {
      console.error(e);
    }
    callback(data);
  });
}

readConfigFile("./package.json", (e) => {
  console.log(e);
});
```

Node predates standardized promises, but because it is fairly consistent
about its error-first callbacks, it is easy to create Promise-based variants of
its callback-based APIs using the `util.promisify()` wrapper:

```js
const util = require("util");
const fs = require("fs");
const pfs = {
  readFile: util.promisify(fs.readFile),
};

function readConfigFile(path) {
  return pfs.readFile(path, "utf8").then((text) => {
    return JSON.parse(text);
  });
}

readConfigFile("./package.json").then(console.log);
```

We can also simply the preceding Promise-based function using `async` and `await`

```js
async function readConfigFile(path) {
  let text = await pfs.readFile(path, "utf8");
  return JSON.parse(text);
}

readConfigFile("./package.json").then(console.log);
```

The `util.promisify()` wrapper can produce a Promise-based version of
many Node functions. In Node 10 and later, the `fs.promises` object has
a number of predefined Promise-based functions for working with the filesystem.
We'll discuss them later in this chapter, but note that in the preceding code, we could
replace `pfs.readFile()` with `fs.promises.readFile()`

We had said that Node's programming model is async-by-default. But for programmer
convenience, Node does define blocking, synchronous variants of many of its functions,
especially in the filesystem module. **These functions typically have names that are clearly
labeled with `Sync` at the end**.

```js
const fs = require("fs");
function readConfigFileSync(path) {
  let text = fs.readFileSync(path, "utf8");
  return JSON.parse(text);
}

console.log(readConfigFileSync("./package.json"));
```

In addition to its error-first two-argument callbacks, Node also has a number
of APIs that use event-based asynchrony, typically for handling streaming data.
We'll cover Node events in more detail later.

**Node's built-in nonblocking functions work using the operating system's version
of callbacks and event handlers**. When you call one of these functions, Node takes
action to get the operation started, then registers some kind of event handler with
the operating system so that it will get notified when the operation is complete. **The
callback you passed to the Node function get stored internally so that Node can
invoke your callback when the operating system sends the appropriate event to Node**.

This kind of concurrency is often called event-based concurrency. At its core,
**Node has a single thread that run an "event loop"**. When a Node program starts,
it runs whatever code you've told it to run. This code presumably calls at least
one nonblocking function causing a callback or event handler to be registered with the
operating system. (If not, then you've written a synchronous Node program, and Node
simply exits when it reaches the end). When Node reaches the end of your program,
it blocks until an event happens, at which time the OS starts it running again.
Node maps the OS event to the JavaScript callback you registered and then invokes
that function. Your callback function may invoke more nonblocking Node functions,
causing more OS event handlers to be registered. Once your callback function is done
running. **Node goes back to sleep again and the cycle repeats**.

For web servers and other I/O-intensive application that spend most of their time
waiting for input and output, this style of event-based concurrency is efficient and
effective. A web server can concurrently handle requests from 50 different clients
without needing 50 different threads as long as it uses nonblocking APIs and there
is some kind of internal mapping from network sockets to JavaScript functions
to invoke when activity occurs on those sockets.

### 16.3 Buffers

One of the datatypes you're likely to use frequently in Node -- especially when
reading data from files or from the network -- is the Buffer class. Node was created
before core JavaScript supported typed arrays and there was no Uint8Array to
represent an array of unsigned bytes. Node defined the Buffer class to fill that need.
Now that Uint8Array is part of the JavaScript language. Node's Buffer class is a
subclass of Uint8Array.

Node's Buffer class has methods that perform both encoding and decoding, and
you can recognize these methods because they expect an `encoding` argument that specifies
the encoding to be used.

Encodings in Node are specified by name, as strings. The supported encodings are:

`"utf8"`

- This is the **default** when no encoding is specified, and is the Unicode
  encoding you are most likely to use

`"utf16le"`

- Two-byte Unicode characters, with little-endian ordering. Code points above `uffff`
  are encoded as a pair of two-byte sequences

`"latin1"`

- The one-byte-per-character ISO-8859-1 encoding that defines a character
  set suitable for many Western European languages.

`"ascii"`

- The 7-bit English-only ASCII encoding, a strict subset of the "utf8" encoding

`"hex"`

- This encoding converts each byte to a pair of ASCII hexadecimal digits.

`"base64"`

- This encoding converts each sequence of three bytes into a sequence of four ascii characters.

Here is some example code that demonstrates how to work with Buffers
and how to convert to and from strings:

```js
let b = Buffer.from([0x41, 0x42, 0x43]); // => <Buffer 41 42 43>
b.toString(); // => "ABC"
b.toString("hex"); // => "414243"

let computer = Buffer.from("IBM3111", "ascii"); // Convert string to buffer
for (let i = 0; i < computer.length; i++) {
  computer[i]--;
}

computer.toString("ascii"); // => "HAL2000"
computer
  .subarray(0, 3)
  .map((x) => x + 1)
  .toString(); // => "IBM"

let zeros = Buffer.alloc(1024);
let ones = Buffer.alloc(128, 1);
let dead = Buffer.alloc(1024, "DEADBEEF", "hex"); // repeating pattern of bytes

dead.readUInt32BE(0); // => 0xDEADBEEF
dead.readUInt32BE(1); // => 0xADBEEFDE
dead.readBigUInt64BE(6); // => 0xBEEFDEADBEEFDEAD
dead.readUInt32LE(1020); // => 0xEFBEADDE
```

If you write a Node program that actually manipulates binary data, you may
find yourself using the Buffer class extensively. On the other hand,
if you are just working with text that is read from or written
to a file or the network, then you may only encounter Buffer as an **intermediate**
representation of your data. A number of Node APIs can take input or return
output as either strings or Buffer objects. Typically, if you pass a string,
or expect a string to be returned, from one of these APIs, **you'll need to
specify the name of the text encoding you want to use. And if you do this, then
you may not need to use a Buffer object at all**

### 16.4 Events and EventEmitter

In Node, objects that emit events are instances of EventEmitter or a
subclass of EventEmitter.

```js
const EventEmitter = require("events");
const net = require("net");

let server = new net.Server();
server instanceof EventEmitter; // => true, it emits a "listening" event when it first starts listening for connections,
// it emits a "connection" event every time a client connects.
// and emits a "close" event when it has been closed and is no longer listening
```

To register an event handler, call the `on()` method, passing the name of
the event type and the function that should be invoked when an event of that type occurs.

```js
const net = require("net");
let server = new net.Server();
server.on("connection", (socket) => {
  // Send some data to the client and disconnect
  socket.end("Hello World", "utf8");
});
```

If you prefer more explicit method names for registering event listeners, you can also use
`addListener()`. And you can remove a previously registered event listener with `off()`
or `removeListener()`. As a special case, you can register an event listener that
will be **automatically removed after it is triggered for the first time by calling
`once()` instead of `on()`**

When an event of a particular type occurs for a particular EventEmitter object,
Node invokes all of the handler functions that are currently registered on that
EventEmitter for events of that type. They are invoked in order from the first registered
to the last registered. If there is more than one handler function, they are invoked
sequentially on a single thread: there is no parallelism in Node, remember. And
importantly, event handling functions are invoked synchronously, not asynchronously.
What this means is that the `emit()` method does not queue up event handlers to be
invoked at some later time. `emit()` invokes all the registered handlers, one after
the other, **and does not return until the last event handler has returned**.

If your program is one -- like a network server -- that needs to be responsive, then
it is important that you keep your event handler functions nonblocking and fast. If
you need to do a lot of computation when an event occurs, it is often best to use
the handler to schedule that computation asynchronously using `setTimeout()`.
Node also defines `setImmediate()`, which schedules a function to be invoked
immediately after all pending callbacks and events have been handled.

```js
function asyncFunction() {
  console.log("Executing asynchronously");
}
// Using setImmediate to schedule the function to run asynchronously
setImmediate(asyncFunction);
console.log("Scheduled function to run asynchronously with setImmediate");

// Result
// Scheduled function to run asynchronously with setImmediate
// Executing asynchronously
```

The EventEmitter class also defines an `emit()` method that causes the
registered event handler functions to be invoked. This is useful if you are
defining your own event-based API, but is not commonly use when
you're just programming with existing APIs.

`emit()` must be invoked with the name of the event type as its first argument.
Any additional arguments that are passed to `emit()` become arguments to
the registered event handler functions. The handler functions are
also invoked with the `this` value set to the EventEmitter object itself, which is
often convenient (not in arrow functions).

```js
import { EventEmitter } from "events";

class A extends EventEmitter {
  constructor() {
    super();
  }

  myCustomMethod() {
    this.emit("customEvent", "Custom data", "Custom data 2");
  }
}

let a = new A();
a.on("customEvent", (data1, data2) => {
  console.log(data1, data2);
});

a.myCustomMethod();
```

Any value returned by an event handler function is ignored. **If an
event handler function throws an exception, however, it propagates out from the
`emit()` call and prevents the execution of any handler functions
that were registered after the one that threw the exception**.

```js
import { EventEmitter } from "events";

class A extends EventEmitter {
  constructor() {
    super();
  }

  myCustomMethod() {
    try {
      this.emit("customEvent", "Custom data", "Custom data 2");
    } catch (e) {
      console.log(e);
    }
  }
}

let a = new A();
a.on("customEvent", (data1, data2) => {
  console.log(data1, data2);
  throw "it is an error";
});

a.myCustomMethod();

// result:
// Custom data Custom data 2
// it is an error
```

Recall that Node's callback-based APIs use error-first callbacks, and **it
is important that you always check the first callback argument to see if
an error occurred.**; With event-based APIs, the equivalent is "error" events.
Most EventEmitter define an "error" event that they emit when an error occurs.
**Whenever you use an event-based API, you should make it a habit to register
a handler for "error" events.** Since the error occurs asynchronously, there is
no way for you to handle the exception in a `catch block`, so this kind of error
typically causes your program to exit.

### 16.5 Streams

When implementing an algorithm to process data, it is almost always
easiest to read all the data into memory, do the processing, and then
write the data out. For example, you could write a Node function to copy a file like this (Node define a `fs.copyFile()` function that you would actually use in practice):

```js
const fs = require("fs");

function copyFile(sourceFilename, destinationFilename, callback) {
  fs.readFile(sourceFilename, (err, buffer) => {
    if (err) {
      callback(err);
    } else {
      fs.writeFile(destinationFilename, buffer, callback);
    }
  });
}
```

This `copyFile()` function uses asynchronous functions and callbacks,
so it does not block and is suitable for use in concurrent programs like
severs. But notice that it must allocate enough memory to hold the entire
contents of the file in memory at once. This may be fine in some use cases,
**but it fails if the files to be copied are very large, or if your program
is highly concurrent and there may be many files being copied at the same time.**
Another shortcoming of this `copyFile()` implementation is that it cannot start writing
the new file until it has finished reading the old file.

The solution to these problems is to use streaming algorithms.

Node's networking APIs are stream-based and Node's filesystem module defines streaming
APIs for reading and writing files, so you are likely to use a streaming API in many
of the Node programs that you write.

Node supports four basic stream types;

_Readable_

Readable streams are sources of data. The stream returned by `fs.createReadStream()`,
for example, is a stream from which the content of a specified file can be read.
`process.stdin` is another Readable stream that returns data from standard input.

_Writable_

Writable streams are sinks or destination for data. The return value of
`fs.createWriteStream()`, for example, is a Writable stream: it allows
data to be written to it in chunks, and outputs all of that data to a specified file.

_Duplex_

Duplex streams combine a Readable stream and a Writable stream into one object.
The Socket objects returned by `net.connect()` and other Node networking APIs, for
example, are Duplex streams.

_Transform_

Transform streams are also readable and writable, but they differ from Duplex streams in
an important way: data written to a Transform stream becomes readable -- usually
in some transformed form -- from the same stream. The `zlib.createGzip()` function,
for example, returns a Transform stream that compresses (with the _gzip_ algorithm)
the data written to it. In a similar way, the `crypto.createCipheriv()` function
returns a Transform stream that encrypts or decrypts data that is written to it.

By default, streams read and write buffers. If you call the `setEncoding()` method
of a Readable stream, it will return decoded strings to you instead of Buffer objects.
And if you write a string to Writable buffer, it will be automatically encoded
using the buffer's default encoding or whatever encoding you specify.
Node's stream API also supports an "object mode" where streams read and
write objects more complex than buffers and strings. None of Node's core
APIs use this object mode, but you may encounter it in other libraries.

```js
process.stdin.setEncoding("utf8")...
```

The tricky thing about stream-based APIs is that the two ends of the stream
will almost always flow at different speeds. Stream implementations almost
always include an internal buffer to hold data that has been written
but not yet read. Buffering helps to ensure that there is data available to read
when it's requested, and that there is space to hold data when it is written. But
neither of these things can ever be guaranteed, and it is the nature of stream-based
programming that readers will sometimes have to wait for data to be written (because
the stream buffer is empty), and writers will sometimes have to wait for
data to be read (because the stream buffer is full).

The need to coordinate stream readability (buffer not empty) and
writability (buffer not full) via events makes Node’s stream APIs
somewhat complicated. This is compounded by the fact that these APIs
have evolved and changed over the years: for Readable streams, there
are two completely distinct APIs that you can use. Despite the
complexity, it is worth understanding and mastering Node’s streaming
APIs because they enable high-throughput I/O in your programs.
The subsections that follow demonstrate how to read and write from
Node’s stream classes.

#### 16.5.1 Pipes

Imagine, for example, that you are writing a simple HTTP server that serves a directory of static
files. In this case, you will need to read data from a file input stream
and write it out to a network socket. You can simply connect the two sockets
together as a "pipe" and let Node handle the complexities for you.
Simply pass the Writable stream to the `pipe()` method of the Readable stream.

```js
const fs = require("fs");
function pipeFileToSocket(filename, socket) {
  fs.createReadStream(filename).pipe(socket);
}

const net = require("net");
let server = new net.Server();
server.on("connection", (socket) => {
  pipeFileToSocket("./myfile", socket);
});
```

The following utility function pipes one stream to another and invokes
a callback when done or when an error occurs:

```js
function pipe(readable, writable, callback) {
  function handleError(err) {
    readable.close();
    writable.close();
    callback(err);
  }

  readable
    .on("error", handleError)
    .pipe(writable)
    .on("error", handleError)
    .on("finish", callback);
}

const fs = require("fs");
let filefrom = fs.createReadStream("./aaa");
let filedest = fs.createWriteStream("./ccc");
pipe(filefrom, filedest, (e) => console.log(e));
```

Transform streams are particularly useful with pipes, and create
pipelines that involve more than two streams:

```js
const fs = require("fs");
const zlib = require("zlib");

function gzip(filename, callback) {
  let source = fs.createReadStream(filename);
  let destination = fs.createWriteStream(filename + ".gz");
  let gzipper = zlib.createGzip();
  source
    .on("error", callback)
    .pipe(gzipper)
    .pipe(destination)
    .on("error", callback)
    .on("finish", callback);
}

gzip("./aaa", (e) => {
  console.log(e);
});
```

Using the pipe() method to copy data from a Readable stream to a
Writable stream is easy, but in practice, you often need to process the
data somehow as it streams through your program. One way to do this
is to implement your own Transform stream to do that processing:

```js
const stream = require("stream");

class GrepStream extends stream.Transform {
  constructor(pattern) {
    super({ decodeStrings: false }); // Don't convert strings back to buffers
    this.pattern = pattern;
    this.incompleteLine = "";
  }

  _transform(chunk, encoding, callback) {
    if (typeof chunk !== "string") {
      callback(new Error("Expected a string but got a buffer"));
      return;
    }
    let lines = (this.incompleteLine + chunk).split("\n");
    this.incompleteLine = lines.pop();

    let output = lines.filter((l) => this.pattern.test(l)).join("\n");
    if (output) {
      output += "\n";
    }

    callback(null, output);
  }

  // This is called right before the stream is closed.
  // It is our chance to write out any last data.
  _flush(callback) {
    if (this.pattern.test(this.incompleteLine)) {
      callback(null, this.incompleteLine + "\n");
    }
  }
}

let pattern = new RegExp(process.argv[2]);
process.stdin
  .setEncoding("utf8")
  .pipe(new GrepStream(pattern))
  .pipe(process.stdout)
  .on("error", () => process.exit());
```

Tips: Run `node xx.js 'yourRegexp'`

#### 16.5.2 Asynchronous Iteration

In Node 12 and later, **Readable streams are asynchronous iterators**,
which means that within an `async` function you can use a `for/await` loop
to read string or Buffer chunks from a stream using code that is
structured like synchronous code would be.

```js
async function grep(source, destination, pattern, encoding = "utf8") {
  source.setEncoding(encoding);
  destination.on("error", (err) => process.exit());

  let incompleteLine = "";

  for await (let chunk of source) {
    let lines = (incompleteLine + chunk).split("\n");
    incompleteLine = lines.pop();
    for (let line of lines) {
      if (pattern.test(line)) {
        destination.write(line + "\n", encoding);
      }
    }
  }
  if (pattern.test(incompleteLine)) {
    destination.write(incompleteLine + "\n", encoding);
  }
}

let pattern = new RegExp(process.argv[2]);
grep(process.stdin, process.stdout, pattern).catch((err) => {
  console.error(err);
  process.exit();
});
```

#### 16.5.3 Writing to Streams and Handling Backpressure

`write()` method takes a buffer or string as the first argument. Object
streams expect other kinds of objects, but are beyond the scope of this chapter.

If you pass a buffer, the bytes of that buffer will be written directly.
If you pass a string, it will be encoded to a buffer of bytes before being written.
Writable streams have a default encoding that is used when you pass a string
as the only argument to `write()`. The default encoding is typically "utf8",
but you can set it explicitly by **calling `setDefaultEncoding()`** on the Writable
stream. Alternatively, you can pass an encoding name as the second argument to `write()`
method.

`write()` optionally takes a **callback** function as its third argument.
This will be invoked when the data has actually been written and is no longer
in the Writable stream's internal buffer.

The `write()` method has a very important return value. When you call `write()`
on a stream, it will always accept and buffer the chunk of data you have passed.
It return `true` if the internal buffer is not yet full. Or, if the buffer is now full
or overfull, it returns `false`. **Writable streams will enlarge
their internal buffer as much as needed if you keep calling `write()`**. But remember
that the reason to use a streaming API in the first place is to avoid the cost of keeping
lots of data in memory at once.

A return value of `false` from the `write()` method is a form of _backpressure_:
a message from the stream that you have written data more quickly than it can be handled.
The proper response to this kind of backpressure is to stop calling `write()` until
the stream emits a "drain" event, signaling that is once again room in the buffer.

```js
function write(stream, chunk, callback) {
  let hasMoreRoom = stream.write(chunk);

  if (hasMoreRoom) {
    setImmediate(callback);
  } else {
    stream.once("drain", callback);
  }
}
```

The fact that it is sometimes OK to call write() multiple times in a
row and sometimes you have to wait for an event between writes
makes for awkward algorithms. This is one of the reasons that using
the pipe() method is so appealing: **when you use pipe(), Node
handles backpressure for you automatically.**

If you are using `await` and `async` in your program, and are treating Readable
streams as asynchronous iterators, it is straightforward to implement
a Promise-based version of the `write()` utility function above to properly handle
backpressure (however, in the async `grep()` function we just looked at, did not
handle the backpressure).

Example:

```js
function write(stream, chunk) {
  let hasMoreRoom = stream.write(chunk);

  if (hasMoreRoom) {
    return Promise.resolve(null);
  } else {
    return new Promise((resolve) => {
      stream.once("drain", resolve);
    });
  }
}

async function copy(source, destination) {
  destination.on("error", (err) => process.exit());

  for await (let chunk of source) {
    await write(destination, chunk);
  }
}

copy(process.stdin, process.stdout);
```

**Caution**: Failing to respond to backpressure **can cause your program to use
more memory than it should when the internal buffer of a Writable stream
overflows and grows larger and larger**.

If you are writing a network server without handling backpressure, this can be
remotely exploitable security issue:

Suppose you write an HTTP server that delivers files over the network. An attacker
could write an HTTP client that initiates requests for large files but never actually
reads the body of the request. Since the client is not reading the data over the network,
buffers on the server are going to overflow. With enough concurrent connections
from the attacker, this can turn into a denial-of-service attack that
slows your sever down or even crashes it.

#### 16.5.4 Reading Streams with Events

Node's readable streams have two modes, each of which has its own
API for reading. If you can't use pipes or asynchronous iteration in your program,
you will need to pick one of these two event-based APIs for handling streams.

##### FLOWING MODE

In flow mode, handle chunk in "data" event of readable stream, handle writable stream's backpressure, when
`write()` return false, call readable stream's `pause()` to stop reading data; and
when "drain" event was triggered, call readable stream's `resume()` method to continue read. "end" event will be emitted when the end of the read stream is reached.
And, as with all streams, an “error” event is emitted if
an error occurs.

Registering a "data" event handler **switches a stream into flow mode**. Note that newly created
streams do not start off in flowing mode. This means that a stream does not emit "data"
events until you register the first "data" event handler.

```js
const fs = require("fs");

function copyFile(sourceFilename, destinationFilename, callback) {
  let input = fs.createReadStream(sourceFilename);
  let output = fs.createWriteStream(destinationFilename);
  input.on("data", (chunk) => {
    let hasRoom = output.write(chunk);
    if (!hasRoom) {
      input.pause();
    }
  });

  input.on("end", () => {
    output.end();
  });

  input.on("error", (err) => {
    callback(err);
    process.exit();
  });

  output.on("drain", () => {
    input.resume();
  });

  output.on("error", (err) => {
    callback(err);
    process.exit();
  });

  output.on("finish", () => {
    callback(null); // call the callback with no error
  });
}

let from = process.argv[2],
  to = process.argv[3];
console.log(`Copying file from ${from} to ${to}`);
copyFile(from, to, (err) => {
  if (err) {
    console.error(err);
  } else {
    console.log("done.");
  }
});
```

##### PAUSED MODE

This is the mode that streams start in.
If you never register a "data" event handler and never call the `pipe()` method,
then a Readable remains in paused mode. **In paused mode,
the stream does not push data to you in the form of "data" events. Instead,
you pull data from the stream by explicitly calling its `read()` method**.
This is not a blocking call, and if there is no data available to read on the stream,
it will return `null`. A Readable steam in paused mode emits
"readable" events when data becomes available to read on the stream.

**Tip**: If you stop calling `read()` while there is still readable data, you will
not get another `readable` event and your program is likely to hang.

Streams in paused mode emit "end" and "error" events just like flowing mode streams
do. It is **inelegant** if you write a copy file program with paused mode, you may find
that flowing mode (or pipes) is easier in this case.

The following code demonstrates how you can compute a SHA256 hash for the contents
of a specified file.

```js
const fs = require("fs");
const crypto = require("crypto");

function sha256(filename, callback) {
  let input = fs.createReadStream(filename);
  let hasher = crypto.createHash("sha256");

  input.on("readable", () => {
    let chunk;
    while ((chunk = input.read())) {
      hasher.update(chunk);
    }
  });

  input.on("end", () => {
    let hash = hasher.digest("hex");
    callback(null, hash);
  });

  input.on("error", callback);
}

sha256(process.argv[2], (err, hash) => {
  if (err) {
    console.error(err.toString());
  } else {
    console.log(hash);
  }
});
```

### 16.6 Process, CPU, and Operating System Details

The global Process object has a number of useful properties and functions
that generally relate to the state of the **currently running Node process**.
Consult the Node documentation for complete details, but here
are some properties and functions you should be aware of:

```js
process.argv; // An array of command line arguments
process.arch; // The CPU architecture: "x64"
process.cwd(); // current working dir. Note that it is the dir where you type "node", not the file location
process.chdir(); // Sets the current working dir
process.cpuUsage(); // Reports user and system CPU time usage of this process { user: 23314, system: 0 }
process.env; // An object of environment variables
process.execPath; // The absolute filesystem path to the `node` executable
process.exit(); // Terminates the program
process.getuid();
process.hrtime.biging(); // Return a "high-resolution" nanosecond timestamp: 17248072991856n
process.kill(); // Send a signal to another process
process.memoryUsage(); // { rss: 54169600, heapTotal: 7438336, heapUsed: 6036144, external: 1018236, arrayBuffers: 16743 }
process.nextTick(); // Like setImmediate(), invoke a function soon
process.pid; // The process id of the current process
process.ppid; // The parent process id
process.platform; // The OS: "linux", "darwin", or "wind32" for example
process.resourceUsage();
// {
//   userCPUTime: 208503,
//   systemCPUTime: 66341,
//   maxRSS: 54240,
//   sharedMemorySize: 0,
//   unsharedDataSize: 0,
//   unsharedStackSize: 0,
//   minorPageFault: 7711,
//   majorPageFault: 15,
//   swappedOut: 0,
//   fsRead: 1840,
//   fsWrite: 224,
//   ipcSent: 0,
//   ipcReceived: 0,
//   signalsCount: 0,
//   voluntaryContextSwitches: 812,
//   involuntaryContextSwitches: 5
// }
process.setuid(); // Sets the current user, by id or name
process.title; // The process name that appears in ps listing
process.umask(); // Set or return the default permissions for new files
process.umask(0222); // Set or return the default permissions for new files
process.uptime(); // Uptim in seconds
process.version; // Node's version string -- v18.5.0
process.versions; // Version strings for the libraries Node depends on
// {
//   node: '18.15.0',
//   v8: '10.2.154.26-node.25',
//   uv: '1.44.2',
//   zlib: '1.2.13',
//   brotli: '1.0.9',
//   ares: '1.18.1',
//   modules: '108',
//   nghttp2: '1.51.0',
//   napi: '8',
//   llhttp: '6.0.10',
//   uvwasi: '0.0.15',
//   acorn: '8.8.2',
//   simdutf: '3.1.0',
//   undici: '5.20.0',
//   openssl: '3.0.8+quic',
//   cldr: '42.0',
//   icu: '72.1',
//   tz: '2022g',
//   unicode: '15.0',
//   ngtcp2: '0.8.1',
//   nghttp3: '0.7.0'
// }
```

The "os" module (needs to be explicitly loaded with `require()`) provides access to similarly low-level
details about the computer and operating system that Node is running
on. You may never need to use any of these features, but it is worth knowing
that Node makes them available:

```js
const os = require("os")
os.arch() // Returns CPU architecture. "x64" or "arm" for example
os.constants // Useful constants such as os.constants.signals.SIGINT
os.cpus() // Data about system CPU cores, including usage times.
os.endianness() // The CPU's native endianness "BE" or "LE"
os.EOL // The OS native line terminator: "\n" or "\r\n"
os.freemem() // The amount of free RAM in bytes
os.getPriority() // Returns the OS scheduling priority of a process
os.homedir() // Returns the current user's home directory.
os.hostname() // Returns the hostname of the computer
os.loadavg() // Return the 1, 5, and 15 minute load averages. (See `top` linux command detail for more)
os.networkInterfaces() // Returns details about available network connections.
os.platform() // Return OS: "linux", "darwin", or "wind32", for example
os.release() // Returns the version number of the OS (e.g 5.15.146.1-microsoft-standard-WSL2)
os.setPriority() // Attempts to set the scheduling priority for a process
os.tmpdir() // returns the default temporary directory e.g '/tmp'
os.totalmem() // Return the total amount of RAM in bytes
os.type() // Returns OS: "Linux", "Darwin", or "Windows_NT", e.g.
os.uptime() // Returns the system uptime in seconds
os.userInfo() // Returns uid, username, home and shell of current user
```

### 16.7 Working with Files

Node's "fs" module is a comprehensive API for working with files and
directories. It is complemented by "path" module, which defines
utility functions for working with file and directory names.

#### 16.7.1 Paths, File Descriptors, and FileHandles

Files are most often specified by _path_, which means the name of the file itself,
plus the hierarchy of directories in which the file appears.
Working with paths can be a little tricky because different operating
system use different characters to separate directory names. Node's "path"
module and a couple of other important Node features help

```js
__filename; // Absolute path (from dir to filename)
__dirname; // Absolute path to dir name

const path = require("path");
path.sep; // Either "/" or "\" depending on your OS

let p = "src/pkg/test.js";
path.basename(p); // => "test.js"
path.extname(p); // => ".js"
path.dirname(p); // => "src/pkg"

// normalize() cleans up paths:
path.normalize("a/b/c/../d/"); // => "a/b/d"
path.normalize("a/./b"); // => "a/b"
path.normalize("//a//b/"); // => "/a/b/"

path.join("src", "pkg", "t.js"); // => "src/pkg/t.js"
path.join("src/pkg/", "t.js"); // => "src/pkg/t.js"
path.join("src/pkg", "t.js"); // => "src/pkg/t.js"

// resolve() takes one or more path segments and returns an absolute
// path.
// It starts with the last argument and works backward until it has meet
// an absolute path, or resolving against process.cwd()
path.resolve(); // => process.cwd()
path.resolve("t.js"); // => path.js(process.cwd(), "t.js")
path.resolve("/tmp", "t.js"); // => "/tmp/t.js"
path.resolve("tmp", "t.js"); // => path.js(process.cwd(), "t.js")
path.resolve("/a", "/b", "t.js"); // => "/b/t.js"
```

Note that `path.normalize()` is simply a string manipulation
function that has no access to the actual filesystem.
The `fs.realpath()` and `fs.realpathSync()` functions perform
filesystem-aware canonicalization: they resolve symbolic links and
interpret relative pathnames relative to the current work directory

```js
$ ll
lrwxrwxrwx 1 soon soon   38 Mar 19 23:00 samplesymbolic -> /home/soon/Repos/Download/demosymbolic
$ node <<- EOF
const fs = require("fs")
console.log(fs.realpathSync("samplesymbolic"))
console.log(fs.realpathSync("."))
EOF

/home/soon/Repos/Download/demosymbolic
/home/soon/Repos/learning/javascript-the-definitive-guide/server
```

If you want work with Unix-style path even when on a Windows system, then
use `path.posix` instead of `path`. And conversely, if you want to work with
Windows paths even when on a Unix system, `path.win32`. These define the
same properties and functions as `path` it self.

```js
path.win32.join("src/pkg", "4"); // => 'src\\pkg\\4'
```

Some of the "fs" functions we'll be covering in the next section s expect a _file descriptor_
instead of a file name. File descriptors are integers used as OS-level
references to "open" files. You obtain a descriptor for a given name by calling
the `fs.open()` (or `fs.openSync()`) function. **Processes are only allowed to have
a limited number of files open at one time**, so it is important that you call `fs.close()`
on your file descriptors when you are done with them. It is only really makes sense
to use the descriptor-based functions if you were going to open the file to read
or write anyway.

Finally, in the Promise-based API defined by `fs.promises`, the equivalent of `fs.open()`
is `fs.promises.open()`, which returns a Promise that resolves to a FileHandle object.
This FileHandle object serves the same purpose as a file descriptor. Again,
however, unless you need to use the lowest-level `read()` and `write()` methods of
a FileHandle, there is really no reason to create one. And if you do create
a FileHandle, you should remember to call its `close()` method once you are done with it.

#### 16.7.2 Reading Files

If your files are small, or if memory usage and performance are not the highest
priority, then it is often easiest to read the entire content of a file
with a single call. You can do this synchronously, with a callback,
or with a Promise. **By default, you'll get the bytes of the file as a buffer,
but if you specify an encoding, you'll get a decoded string instead.**

```js
const fs = require("fs");
let buffer = fs.readFileSync("index.html"); // returns buffer
let text = fs.readFileSync("index.html", "utf8"); // returns string

// Read the bytes of the file asynchronously
fs.readFile("index.html", (err, buffer) => {
  if (err) {
    // Handle the error here
  } else {
    // The bytes of the file are in buffer
  }
});

// Promise-based asynchronous read
fs.promises
  .readFile("index.html", "utf8")
  .then(processFileText)
  .catch(handleReadError);

// Or use the Promise API with await inside an async function
async function processFileText(filename, encoding = "utf8") {
  let text = await fs.promises.readFile(filename, encoding);
  // ... process the text here
}
```

If you are able to process the contents of a file sequentially and **do not need
to have the entire content of the file in memory at the same time**, then reading
a file via a stream may be the most efficient approach.

```js
const fs = require("fs");

function printFile(filename, encoding = "utf8") {
  fs.createReadStream(filename, encoding).pipe(process.stdout);
}
```

Finally, if you need low-level control over exactly what bytes you read from
a file and when you read them, you can **open** a file to get a descriptor
and then use `fs.read()`, `fs.readSync()`, or `fs.promises.read()` to read a specified
number of bytes from a source location of file into a specified buffer at the specified
destination position:

```js
const fs = require("fs");

fs.open("index.html", (err, fd) => {
  // tip: fd is a number
  if (err) {
    // Report error somehow
    return;
  }

  try {
    // Read bytes 20 through 420 into a newly allocated buffer. // fd: number, buffer: TBuffer, offset: number, length: number, position: ReadPosition | null, callback: (err: NodeJS.ErrnoException | null, bytesRead: number, buffer: TBuffer) => void
    fs.read(fd, Buffer.alloc(400), 0, 400, 20, (err, n, b) => {
      console.log(b);
    });
  } finally {
    fs.close(fd);
  }
});
```

The callback-based `read()` API is awkward to use if you need to read more
than one chunk of data from a file. If you can use the
synchronous API (or the Promise-based API with `await`), it becomes
easy **to read multiple chunks from a file**.

```js
const fs = require("fs");

function readData(filename) {
  let fd = fs.openSync(filename);
  try {
    let header = Buffer.alloc(12);
    fs.readSync(fd, header, 0, 12, 0);

    // Verify the file's magic number
    let magicNumber = header.readInt32LE(0);
    if (magicNumber !== 0xdadafeed) {
      console.log(magicNumber.toString(16));
      throw new Error("File is of wrong type");
    }

    // Now get the offset and length of the data from the header
    let offset = header.readInt32LE(4);
    let length = header.readInt32LE(8);

    // And read those bytes from the file
    let data = Buffer.alloc(length);
    fs.readSync(fd, data, 0, length, offset);
    return data;
  } finally {
    fs.closeSync(fd);
  }
}
```

#### 16.7.3 Writing Files

There are three basic ways to write files in Node. If you have the
entire of the file in a string or a buffer, you can write the entire string
in one call with `fs.writeFileSync()`, `fs.writeFile()` (callback-based), or
`fs.promises.writeFile()` (Promise-based).

**Important**: You create a new file simply by writing to a filename that does not
already exist.

**Important**: Every `writeFile` methods overwrite existing file

```js
const fs = require("fs");
const path = require("path");

const settings = { hello: "world" };

fs.writeFileSync(
  path.resolve(__dirname, "settings.json"),
  JSON.stringify(settings),
);
```

If you want to use an encoding other than "utf8", pass the encoding as an optional third
argument.

The related functions `fs.appendFile()`, `fs.appendFileSync()`, and `fs.promises.appendFile()` are similar, **but when the specified file already exists, they append their
data to the end rather than overwriting the existing file content**.

**Important**: You create a new file simply by appending to a filename that does not
already exist.

```js
fs.appendFileSync(
  path.resolve(__dirname, "settings2.json"),
  JSON.stringify(settings),
);
```

If the data you want to write to a file is not all in one chunk, or if it
is not all in memory at the same time, then use a Writable stream is a
good approach.

```js
const fs = require("fs");
let output = fs.createWriteStream("numbers.txt");
for (let i = 0; i < 100; i++) {
  output.write(`${i}\n`);
}
output.end();
```

Finally, if you want to write data to a file in a multiple chunks, and
you want to be able to control the exactly position within the file at which
each chunk is written, then you can open the file with `fs.open()`, `fs.openSync()`,
or `fs.promises.open()` and then use the resulting file descriptor with the
`write()` or `fs.writeSync()` functions. These functions come in different forms
for strings and buffers.

```js
// String variant
const fs = require("fs");
fs.open("settings.json", "w", (err, fd) => {
  if (err) {
    return;
  }
  const filePositionToWrite = 50;
  try {
    fs.writeSync(fd, "hello", filePositionToWrite, "utf8");
  } finally {
    fs.close(fd);
  }
});

// Buffer variant
const fs = require("fs");
fs.open("settings4.json", "w+", (err, fd) => {
  if (err) {
    return;
  }
  let buffer = Buffer.from("hello", "utf8");
  const offset = 0;
  const length = 2;
  const filePositionToWrite = 0;
  try {
    fs.writeSync(fd, buffer, offset, length, filePositionToWrite);
  } finally {
    fs.close(fd);
  }
});
```

If you have an **array** of Buffer objects that you want to write, you can do this
with a single `fs.writev()` or `fs.writevSync()`.

```js
const fs = require("fs");
fs.open("settings4.json", "w+", (err, fd) => {
  if (err) {
    return;
  }
  let buffer = Buffer.from("hello", "utf8");
  try {
    fs.writevSync(fd, [buffer, buffer]);
  } finally {
    fs.close(fd);
  }
});
```

Similar low-level functions exist for writing buffers and strings using
`fs.promises.open()` and the FileHandle object it produces.

```js
const fs = require("fs");
fs.promises.open("settings.json", "w").then((fh) => {
  let buffer = Buffer.from("hello", "utf8");
  try {
    const offset = 0;
    const length = 2;
    const filePositionToWrite = 0;
    fh.write(buffer, offset, length, filePositionToWrite);
  } finally {
    fh.close();
  }
});
```

> **File Mode Strings**

To low leveled write file, you must specify a second string argument to `fs.open()`
or `fs.openSync()` that specifies how you intend to use the file descriptor.

| File mode | Description                                                                       |
| --------- | --------------------------------------------------------------------------------- |
| **"w"**   | Open the file for writing (create if non exist)                                   |
| **"w+"**  | Open for writing and reading (create if non exist)                                |
| **"wx"**  | Open for creating a new file; fails if the named file already exists              |
| **"wx+"** | Open for creation, and also allow reading; fails if the named file already exists |
| **"a"**   | Open the file for appending (create if non exist)                                 |
| **"a+"**  | Open for appending, but also allow reading (create if non exist)                  |

**If you do not pass one of these flag strings to `fs.open()` or `fs.openSync()` ,
they use the default "r" flag, making the file descriptor read-only**.

Note that it can also be useful to pass these flags to other file-writing methods:

```js
const fs = require("fs");
fs.writeFileSync("messages.log", "hello", { flag: "a" });

const ws = fs.createWriteStream("messages.log", { flags: "a" });
process.stdin.pipe(ws);
```

**Important**: Note that it is `flags` in `createWriteStream`, but `flag` in `writeFileSync`

You can chop off the end of file with `fs.truncate()`, `fs.truncateSync()`, or
`fs.promises.truncate()`. These functions take a path as their first argument and a length
as their second, and **modify the file so that file has the specified length**. Specify
zero length to make the file empty. If the length is longer than the file has, then the
file is extended with zero bytes to the new size. If you have already opened the file,
you can use the `fs.ftruncate()` or `fs.ftruncateSync()` with the file descriptor or FileHandle.

```js
const fs = require("fs");
fs.truncateSync("settings.json", 10);
fs.open("settings.json", "w", (err, fd) => {
  if (err) {
    console.error(err);
    return;
  }
  fs.ftruncateSync(fd, 0);
});
```

The various file-writing functions described here return or invoke their callback
or resolve their Promise when the data has been "written" in the sense that Node
**has handed it off to the operating system**. But this does not necessarily
mean that the data has actually been written to persistent storage yet. If you
want to force your data out to disk so you know for sure that it has been safely
saved, use `fs.fsync()` or `fs.fsyncSync()`. These functions only work with file
descriptors: there is no path-based version.

```js
fs.fsync(fd, (err) => {
  if (err) console.error(err);
  // ...
});
```

#### 16.7.4 File Operations

`fs.copyFile()` and `fs.copyFileSync()` takes the name of the original file and the
name of the copy as their first two arguments. These can be specified as strings or
as URL (file schema) or Buffer objects. An optional third argument is an integer whose bit specify
flags that control details of `copy` operation.

```js
const fs = require("fs");
fs.copyFile("./aaa", "bbb", fs.constants.COPYFILE_EXCL, (err) => {
  // this callback will be called when done
  if (err) console.error(err);
});

// COPYFILE_FICLONE -- copy-on-write clone -- means that no additional storage space will be required
// until either the original or the copy is modified
fs.promises
  .copyFile(
    "aaa",
    `aaa${new Date().toISOString()}`,
    fs.constants.COPYFILE_EXCL | fs.constants.COPYFILE_FICLONE,
  )
  .then(() => {
    console.log("Backup done");
  })
  .catch(console.err);
```

`fs.rename()` function (along with the usual synchronous and Promise-based variants)
moves and/or renames a file. There is no flags argument.

```js
fs.renameSync("bbb", "ccc");
```

`fs.link()` and `fs.symlink()` and their variants have the same signatures as `fs.rename()`
and behave something like `fs.copyFile()` except that they create hard links
and symbolic links.

Finally, `fs.unlink()`, `fs.unlinkSync()` and `fs.promises.unlink()` are Node's functions
for **deleting a file**.

#### 16.7.5 File Metadata

The `fs.stat()`, `fs.statSync()` and `fs.promises.stat()` functions allow you
to obtain metadata for a specified file or directory.

```js
const fs = require("fs");
let stats = fs.statSync("ccc");
stats.isFile(); // => true
stats.isDirectory(); // => false
stats.size; // bytes size
stats.atime; // access time: Date when it was last read
stats.mtime; // modification time: Date when its content was last written
stats.uid; // The user id of the file's owner
stats.gid; // The group id of the file's owner
stats.mode.toString(8); // the file's permissions, as an octal string
console.log(stats);
// Stats {
//   dev: 2080,
//   mode: 33188,
//   nlink: 1,
//   uid: 1000,
//   gid: 1000,
//   rdev: 0,
//   blksize: 4096,
//   ino: 124186,
//   size: 14,
//   blocks: 8,
//   atimeMs: 1711169128597.1738,
//   mtimeMs: 1711169128597.1738,
//   ctimeMs: 1711169575417.2607,
//   birthtimeMs: 1711169128597.1738,
//   atime: 2024-03-23T04:45:28.597Z,
//   mtime: 2024-03-23T04:45:28.597Z,
//   ctime: 2024-03-23T04:52:55.417Z, <-- change time
//   birthtime: 2024-03-23T04:45:28.597Z
// }
```

`fs.lstat()` and its variants work just like `fs.stat()`, except that
**if the specified file is a symbolic link, Node will return metadata
for link itself rather than following the link**.

```js
const fs = require("fs");
let stats = fs.lstatSync("ccclink");
stats.isSymbolicLink(); // true

let stats = fs.statSync("ccclink");
stats.isSymbolicLink(); // false
```

If you have opened a file to produce a file descriptor or a FileHandle object,
then you can use `fs.stat()` or its variants to get meta information for the
opened file without having to specify the filename again.

```js
fs.open("ccc", fs.constants.O_RDONLY, (err, fd) => {
  if (err) console.error(err);
  console.log(fs.fstatSync(fd));
});
```

`fs.chmod()`, `fs.lchmod()` and `fs.fchmod()` (along with synchronous and Promise-based
versions) set the "mode" or permissions of a file or directory.

```js
fs.chmodSync("ccc", 0o777);
```

`fs.chown()`, `fs.lchown()` and `fs.fchown()` (along with synchronous and Promise-based
versions) set the owner and groups (as IDs) for a file or directory.

```js
fs.chownSync("ccc", 0, 0);
```

Finally, `fs.utimes()` and `fs.futimes()` and their variants set the access time and modification time
of a file or directory.

```js
fs.utimesSync("ccc", new Date(0), new Date(0));
```

#### 16.7.6 Working with Directories

To create a new directory in Node, use `fs.mkdir()`, `fs.mkdirSync()` or `fs.promises.mkdir()`. The optional second argument is the mode (permission bits). Or you can pass
an object with optional `mode` and `recursive` properties. If `recursive` is `true`,
then this function will create any directories in the path that do not already exist:

**Tip**: If there these are existing directories, the files inside them will not be removed

```js
fs.mkdirSync("a/b/c/d/e", { mode: 0o777, recursive: true });
```

`fs.mkdtemp()` and its variants take a path prefix you provide, **append some random
characters to it** (this is important for security), create a directory
with that name, and return (or pass to a callback) the directory path to you.

```js
let dir;
try {
  dir = fs.mkdtempSync(path.join(os.tempdir(), "temp111"));
  console.log(dir); // => temp111QuigOj
  console.log(fs.statSync(dir).isDirectory()); // true
} finally {
  fs.rmdirSync(dir);
}
```

`fs.rmdir()` or one of its variants to delete a directory. **Note that directories
must be empty before they can be deleted** (it will throw error otherwise):

The "fs" module provides two distinct APIs for listing the contents of a directory.
First, `fs.readdir()`, `fs.readdirSync()` and `fs.promises.readdir()` read the entire
directory all at once and give you an array of strings or array of Dirent objects that
specify the name and types (file or directory) of each item.

```js
let tempFiles = fs.readdirSync("/tmp"); // => array of strings
tempFiles = fs.readdirSync("/tmp", { withFileTypes: true }); // => array of Dirent
fs.promises
  .readdir(".", { withFileTypes: true })
  .then((entries) => {
    entries
      .filter((entry) => entry.isDirectory())
      .map((entry) => entry.name)
      .forEach((entry) => console.log(path.join("/tmp", entry)));
  })
  .catch(console.error);
```

**If you anticipate needing to list directories that might have thousands of entries**,
you might prefer the streaming approach of `fs.opendir()` and its variants.
These functions return a Dir object representing the specified directory. You
can use the `read()` or `readSync()` methods of the Dir object to read Dirent at a time.
If you pass a callback function to `read()`, it will call the callback.
And if you omit the callback argument, it will return a Promise. **When there
are no more directory entries, you'll get `null` instead of a Dirent object**.

The easiest way to use Dir objects is as async iterators with a `for/await` loop.

```js
async function listDirectory(dirpath) {
  let dir = await fs.promises.opendir(dirpath);
  for await (let entry of dir) {
    let name = entry.name;
    if (entry.isDirectory()) {
      name += path.sep; // Add a trailing slash to subdirectories
    }
    let stats = await fs.promises.stat(path.join(dirpath, name));
    let size = stats.size;
    console.log(String(size).padStart(10), name);
  }
}

listDirectory("/tmp");
//    4096 a-XXXXXXpRt50S/
//    4096 gobuild2695/
//    4824 gopls-fas-21207278
//    4096 a-XXXXZCel/
//    4096 a-XXXXAKT0/
```

### 16.8 HTTP Clients and Servers

Node's "http", "https", and "https" modules are full-featured but
relatively low-level implementations of the HTTP protocols. Before these
APIs are relatively low-level, there is not room in this chapter to cover all
these features. But the examples that follow demonstrate how to write basic
clients and servers.

The simplest way to make a basic HTTP GET request is with `http.get()` or
`https.get()` (or `https`). The second arguments is a callback that will be invoked
with an IncomingMessage object when the server's response has started to arrive,
but the body may not be ready yet. The IncomingMessage object is a Readable stream.

```js
const http = require("http");

function getJSON(url) {
  return new Promise((resolve, reject) => {
    request = http.get(url, (res) => {
      if (res.statusCode !== 200) {
        reject(new Error(`HTTP status ${res.statusCode}`));
        res.resume();
      } else if (res.headers["content-type"] !== "application/json") {
        reject(new Error(`Invalid content-type`));
        res.resume();
      } else {
        let body = "";
        response.setEncoding("utf-8");
        response.on("data", (chunk) => {
          body += chunk;
        });
        response.on("end", () => {
          try {
            let parsed = JSON.parse(body);
            resolve(parsed);
          } catch (e) {
            reject(e);
          }
        });
      }
    });
    request.on("error", (error) => {
      reject(error);
    });
  });
}
```

`http.get()` and `https.get()` are slightly simplified variant of the more general
`http.request()` and `https.request()` functions. The following `postJSON()`
function demonstrates how to use `http.request` to make an HTTP POST request.

```js
function postJSON(host, endpoint, body, port, username, password) {
  return new Promise((resolve, reject) => {
    let bodyText = JSON.stringify(body);
    let requestOptions = {
      method: "POST",
      host: host,
      path: endpoint,
      headers: {
        "Content-Type": "application/json",
        "Content-Length": Buffer.byteLength(bodyText),
      },
    };
    if (port) {
      requestOptions.port = port;
    }
    if (username && password) {
      requestOptions.auth = `${username}:${password}`;
    }

    let request = http.request(requestOptions);
    request.write(bodyText);
    request.end();

    request.on("error", (e) => reject(e));
    request.on("response", (response) => {
      if (response.statusCode !== 200) {
        reject(new Error(`HTTP status ${response.statusCode}`));
        response.resume();
        return;
      }
      let body = "";
      response.setEncoding("utf-8");
      response.on("data", (chunk) => {
        body += chunk;
      });
      response.on("end", () => {
        try {
          let parsed = JSON.parse(body);
          resolve(parsed);
        } catch (e) {
          reject(e);
        }
      });
    });
  });
}

postJSON("localhost", "/post", { hoho: "hehe" }, 8081)
  .then(console.log)
  .catch(console.error);
```

To write servers that respond to those request, the basic approach is as follows:

- Create a new Server object.
- Call its `listen()` method to begin listening for requests on a specified port.
- Register an event handler for "request" events, use that handler to read the
  client's request, and write your response.

Simple static files server and implementing a debugging endpoint that responds
to a client's request by echoing that request.

```js
const http = require("http");
const url = require("url");
const path = require("path");
const fs = require("fs");

function serve(rootDirectory, port) {
  let server = new http.Server();
  server.listen(port);
  console.log("Listening on port", port);

  server.on("request", (request, response) => {
    let endpoint = url.parse(request.url).pathname;

    if (endpoint === "/test/mirror") {
      response.setHeader("Content-Type", "text/plain;charset=UTF-8");
      // Begin the response body with the request
      response.write(
        `${request.method} ${request.url} HTTP/${request.httpVersion}\r\n`,
      );

      // Output the request headers
      let headers = request.rawHeaders;
      for (let i = 0; i < headers.length; i++) {
        response.write(`${headers[i]} ${headers[i + 1]}\r\n`);
      }

      // End headers with an extra blank line
      response.write("\r\n");

      // Now we need to copy any request body to the response body
      // Since they are both streams, we can use a pipe
      request.pipe(response);
    } else {
      let filename = endpoint.substring(1);
      // Don't allow "../" in the path because it would be a security hole
      filename = filename.replace(/\.\.\//g, "");
      filename = path.resolve(rootDirectory, filename);

      let type;
      switch (path.extname(filename)) {
        case ".html":
        case ".htm":
          type = "text/html";
          break;
        case ".js":
          type = "text/javascript";
          break;
        case ".css":
          type = "text/css";
          break;
        case ".png":
          type = "image/png";
          break;
        case ".txt":
          type = "text/plain";
          break;
        default:
          type = "application/octet-stream";
      }
      let stream = fs.createReadStream(filename);
      stream.once("readable", () => {
        response.setHeader("Content-Type", type);
        response.writeHead(200);
        stream.pipe(response);
      });

      stream.on("error", (err) => {
        response.setHeader("Content-Type", "text/plain; charset=UTF-8");
        response.writeHead(404);
        response.end(err.message);
      });
    }
  });
}

serve(process.argv[2] || "/tmp", parseInt(process.argv[3]) || 8082);
```

Node's built-in modules are all you need to write simple HTTP and HTTPs servers.
Note, however, that production servers are not typically built directly
on top of these modules. Instead, most nontrivial servers are implemented
using external libraries -- such as the [Express framework](https://expressjs.com/) -- that provide
"middleware" and other higher-level utilities that backend web developers
have come to expect.

### 16.9 Non-HTTP Network Servers and Clients

Even though Node has a reputation as a good environment for writing web servers,
Node also has full support for writing other types of network servers and clients.

The "net" module defines Server and Socket classes. To create a server, call `net.createServer()` and then call the `listen()` method of the resulting object to tell the
server what port to listen on for connections. The Server object will generate "connection"
events when a client connects on that port, and the value passed to the event listener
will be a Socket object. **The Socket object is a Duplex stream**, and you can use it
to read data from the client and write data to the client. Call `end()` on the Socket
to disconnect.

Writing a client is even easier: pass a port number and hostname to `net.createConnection()`
to create a socket to communicate with whatever server is running on that host and
listening on that port. Then use that socket to read and write data from and to the server.

TCP Server:

```js
const net = require("net");
const readline = require("readline");

let server = net.createServer();
server.listen(6789, () => console.log("Delivering laughs on 6798"));

server.on("connection", (socket) => {
  tellJoke(socket)
    .then(() => {
      socket.end();
    })
    .catch((err) => {
      console.error(err);
      socket.end();
    });
});

const jokes = {
  Boo: "Don't cry... it's only a joke!",
  Lettuce: "Let us in! It's freezing out here!",
  "A little old lady": "Wow, I didn't know you could yodel!",
};

async function tellJoke(socket) {
  let randomElement = (a) => a[Math.floor(Math.random() * a.length)];
  let who = randomElement(Object.keys(jokes));
  let punchline = jokes[who];
  let lineReader = readline.createInterface({
    input: socket,
    output: socket,
    prompt: ">> ",
  });

  function output(text, prompt = true) {
    socket.write(`${text}\r\n`);
    if (prompt) lineReader.prompt();
  }

  let stage = 0;
  output("Knock knock!");

  for await (let inputLine of lineReader) {
    if (stage === 0) {
      if (inputLine.toLowerCase() === "who's there?") {
        output(who);
        stage = 1;
      } else {
        output('Please type "Who\'s there?".');
      }
    } else if (stage === 1) {
      if (inputLine.toLowerCase() === `${who.toLowerCase()} who?`) {
        output(`${punchline}`, false);
        return;
      } else {
        output(`Please type "${who} who?".`);
      }
    }
  }
}
```

Simple text-based servers like this do not typically need a custom client. `nc` utility
is enough:

```js
$ nc localhost 6789
Knock knock!
>> Who's there?
A little old lady
>> A little old lady who?
Wow, I didn't know you could yodel!
```

We can also write a client:

```js
const net = require("net");

let socket = net.createConnection(6789, process.argv[2]);
process.stdin.pipe(socket).pipe(process.stdout);
socket.on("close", () => {
  process.exit();
});
```

In addition to supporting TCP-based servers, Node's "net" module also supports
interprocess communication over "Unix domain socket" that are identified by a
filesystem path rather than by a port number. We are not going to cover that kind of
socket in this chapter, but the Node documentation has details.

Other Node features that we don't have space to cover here include the "dgram"
module for UDP-based clients and servers and the "tls" module that is to "net"
as "https" is to "http". The `tls.Server` and `tls.TLSSocket` classes allow
the creation of TCP servers (like the knock-knock joke server) that use SSL-encrypted
connections like HTTPs servers do.

### 16.10 Working with Child Processes

In Node, the "child_process" module defines a number of functions for running
other programs as child processes.

#### 16.10.1 execSync() and execFileSync()

`execSync()` takes the command to run as its first argument. It creates a child
process, runs a shell in that process, and uses the shell to execute
the command you passed. If the command exits with an error (status other than 0), then `execSync()` throws an exception
Otherwise, `execSync()` returns whatever output the command writes to its stdout stream.
By default this return value is a buffer, but you can specify an encoding in an optional
second argument to get a string instead. **If the command writes any output to
stderr, that output just get passed through the parent process's stderr stream**

```bash
echo "this will be the output 1"
echo "this will be the output 2"
echo "this will be the output 3"
echo "this will be just console.error" >&2
```

```js
const child_process = require("child_process");
let output = child_process.execSync("./childprocess.sh", { encoding: "utf8" });
console.log(output);

// $ node childprocess.js
// this will be just console.error
// this will be the output 1
// this will be the output 2
// this will be the output 3
```

**Important**: You must be careful to never pass a command to `execSync()` if
any portion of that command is user input or comes from a similar untrusted source. Imagine
the input is `; rm *` and appended into the command.

If you don't need the features of a shell, you can avoid the overhead of starting
a shell by using `child_process.execFileSync()`. **This function executes
a program directly, without invoking a shell**. But since
no shell is involved, it can't parse a command line, and you must pass the
executable as the first argument and an array of command-line arguments as the second
argument

```js
const childprocess = require("child_process");
let output1 = childprocess.execFileSync("ls", ["-l"], {
  encoding: "utf8",
});
```

> **Child Process Options**

Many other `child_process` functions have a second or third optional argument that specifies
additional details how the child process is to run.
Other than `encoding`, there are other important properties that you can specify include
the following (note that not all options are available to all child process functions):

- `cwd` specifies the working directory for the child process, the default is `process.cwd()`
- `env` specifies the environment variables that child process will have access to. By default, child processes simply inherit `process.env`
- `input` specifies a string or buffer of input data that should be used as the
  standard input to the child process. This option is only available to the synchronous functions that do not return a ChildProcess object
- `maxBuffer` specifies the maximum number of bytes of output that will be collected by
  the `exec` functions. (It does not apply to `spawn()` and `fork()`, which use streams).
  **If a child process produces more output than this, it will be killed and will exit with
  an error**.
- `shell` specifies the path to a shell executable or `true`.
- `timeout` specifies the timeout in millisecond (This option applies to the `exec` functions
  but not to `spawn()` or `fork()`)
- `uid` specifies the user ID (a number) under which the program should be run. If
  the parent process is running in a privileged account, it can use this option
  to run the child with reduced privileges.

#### 16.10.2 exec() and execFile()

If you're writing a program that needs to accomplish a number of tasks, and
those tasks don't depend on each other in any way, then you may want to parallelize
them and run multiple commands at the same time. You can do this with the
asynchronous functions `child_process.exec()` and `child_process.execFile()`

`exec()` and `execFile()` returns a ChildProcess object that represent the running
child process. And they accept callback as the final argument. This callback is
invoked when the child process exits, with three arguments: error, stdout, and stderr.

```js
const childprocess = require("child_process");
let cp = childprocess.exec("./childprocess.sh", (err, stdout, stderr) => {
  console.log(err);
  console.log("stdout:");
  console.log(stdout);
  console.log("stderr:");
  console.log(stderr);
});
```

The ChildProcess object returned by `exec()` and `execFile()` allows you
to terminate the child process, and to write data to it (which it can
then read from its standard input). We'll cover them in [children_process.spawn()](#spawn)

If you plan to execute multiple child processes at the same time, then it
may be easiest to use the "promisified" version of `exec()` which returns
a Promise object which, if the child process exits without error, resolves
to an object with `stdout` and `stderr` properties.

```js
const childprocess = require("child_process");
const util = require("util");
const execP = util.promisify(childprocess.exec);

function parallelExec(commands) {
  let promises = commands.map((command) =>
    execP(command, { encoding: "utf8" }),
  );
  return Promise.all(promises).then((outputs) =>
    outputs.map((out) => out.stdout),
  );
}

parallelExec(["ls", "ls -l", "invalid"])
  .then((arr) => console.log(arr))
  .catch(console.error);
```

<a id="spawn"></a>

#### 16.10.3 spawn()

The various `exec` functions described so far return the process output in a single
batch, only after the process has exited.

The `child_process.spawn()` function allows you streaming access to the output
of the child process, while the process is still running. It also allows
you to write data to the child process (which will see the data as input on its
standard input stream).

`spawn()` does not use a shell by default, so you must invoke it like `execFile()`
with the executable to be run and separate array of command-line arguments to
pass to it. `spawn()` returns a ChildProcess object like `execFile()` does,
but it does not take a callback argument. Instead of using a callback function, you
listen to events on the ChildProcess object and on its streams.

"exit" event to be notified when the child process exits. An ChildProcess
object also has three **stream properties**. `stdout` and `stderr` are Readable streams;
`stdin` is a WritableStream.

```bash
while true; do
  echo "Enter one or more value >> "
	read -r
	echo "REPLY = '$REPLY'"
done
```

```js
let p1 = childprocess.spawn("./childprocess.sh", []);
p1.stdout.pipe(process.stdout);
process.stdin.pipe(p1.stdin);
```

The ChildProcess object also defines a `pid` property and that specifies the process
id of the child. And it defines a `kill()` method that you can use to terminate a child process.

#### 16.10.4 fork()

`child_process.fork()` is a specialized function for running a module of
JavaScript code in a child Node process. `fork()` expects the same arguments as
`spawn()`, but the first argument should specify the path to a **file of JavaScript
code instead of an executable binary file**.

`fork()` enables another much easier (in addition to stdout and stdin in `spawn()`)
communication channel between the parent and child processes.

When you create a child process with `fork()`, you can use the `send()` method
of the returned ChildProcess object to send **a copy of an object** to the child process.
And you can listen for the "message" event on the ChildProcess to receive messages from
the child. The code running in the child process can use `process.send()` to send
a message to the parent and can listen for "message" events on `process` to receive
messages from the parent.

Parent:

```js
let child = childprocess.fork("./childprocess_child.js");
child.send({ x: 4, y: 3 });

child.on("message", (message) => {
  console.log(message.hypotenuse);
  child.disconnect();
});
```

Child:

```js
process.on("message", (message) => {
  process.send({ hypotenuse: Math.hypot(message.x, message.y) });
});
```

**Starting child process is an expensive operation**. If you are
writing a program that needs to be very responsive to incoming events and also
needs to perform time-consuming computations, then you might consider using a
separate child process to perform the computations so that they don't block the
event loop and reduce the responsiveness of the parent process.
(Though a thread -- see [16.11](#16.11) -- may be a better choice than a child
process in this scenario)

The first argument to `send()` will be serialized with `JSON.stringify()` and
deserialized in the child process with `JSON.parse()`, so you should only
include values that are supported by the JSON format.
`send()` has a special second argument, however, that allows you to transfer Socket
and Server objects (from the "net" module) to a child process.

<a id="16.11"></a>

### 16.11 Worker Threads

Multithreaded programming has a well-deserved reputation for being difficult. This
is almost entirely because of the need to carefully synchronize access by threads
to shared memory. But JavaScript threads (in both Node and browsers) do not share
memory by default, so the dangers and difficulties of using threads do not apply to
these "workers" in JavaScript.

Instead of using shared memory, JavaScript's worker threads communicate by
message passing. The main thread can send a message to a worker thread by
calling the `postMessage()` method of the Worker object that represents that
thread. The worker thread can receive messages from its parent by listening for "message"
events. And workers can send messages to the main thread with their own version of
`postMessage()`, which the parent can receive with its own "message" event handler.

There are three reasons why you might want to use worker threads in a Node application:

- Your application actually needs to do more computation than one CPU core can handle,
  then threads allow you to distribute work across the multiple cores.

- Even if your application is not using the full power of one CPU, you may still want
  to use threads to maintain the responsiveness of the main thread. Consider a server
  that handles large but relatively infrequent requests. Suppose it gets only one
  request a second, but needs to spend about half a second of (blocking CPU-bound)
  computation to precess each request. On average, it will be idle 50% of the time.
  But when two requests arrive within a few milliseconds of each other, the
  server will not even be able to begin a response to the second request until the
  computation of the first response is complete.

- In general, workers allow us to turn blocking synchronous operations into non
  blocking asynchronous operations.

Worker threads are not nearly as heavyweight as child processes, but they are
not lightweight. It does not generally make sense to create a worker unless
**you have significant work for it to do**. And, generally speaking, if your program
is not CPU-bound and is not having responsiveness problems, **then you probably
do not need worker threads**.

#### 16.11.1 Creating Workers and Passing Messages

The Node module that defines workers is known as "worker_threads".

```js
const threads = require("worker_threads");
```

```js
const threads = require("worker_threads");

if (threads.isMainThread) {
  module.exports = function reticulateSplines(splines) {
    return new Promise((resolve, reject) => {
      let reticulator = new threads.Worker(__filename);
      reticulator.postMessage(splines);
      reticulator.on("message", resolve);
      reticulator.on("error", reject);
    });
  };
} else {
  threads.parentPort.once("message", (splines) => {
    for (let spline of splines) {
      spline.reticulator ? spline.reticulate() : (spline.reticulated = true);
    }
    threads.parentPort.postMessage(splines);
  });
}
```

Run the worker or main thread

```js
const reticulateSplines = require("./worker.js");
const splines = [{ hello: "world" }];
reticulateSplines(splines).then(console.log);
```

**Important**: If you specify a relative path, it is relative to `process.cwd()`,
not relative to the currently running module. If you want a path relative to the current module, use something like `path.resolve(__dirname, 'workers/reticulator.js')`

The `Worker()` constructor can also accept an object as its second argument, and
the properties of this object provide optional configuration for worker. We'll
cover a number of these options later. But for now note that if you pass `{eval: true}`,
then the first argument to `Worker()` is interpreted as a string of JavaScript code
to be evaluated instead of filename.

```js
const threads = require("worker_threads");
new threads.Worker(
  `
const threads = require("worker_threads")
threads.parentPort.postMessage(threads.isMainThread);
`,
  { eval: true },
).on("message", console.log);
```

Node makes a copy of object passed to `postMessage()` rather than sharing it
directly with the worker thread. This prevents the worker thread and the main thread
from sharing memory. Node borrows a more robust technique known as the structured
clone algorithm from web browsers.

This algorithm enables serialization of most JavaScript types, including Map,
Set, Date, and RegExp objects and typed arrays, but it cannot, in general,
copy types defined by the Node host environment such as sockets and streams.
Note, however, that Buffer objects are partially supported: if you pass
a Buffer to `postMessage()` it will be received as a Uint8Array, and can be
converted back into Buffer with `Buffer.from()`

#### 16.11.2 The Worker Execution Environment

- `threads.isMainThread` is `true` in the main thread

- In the main thread, `thread.parentPort` is always `null`.

- In a worker thread, `threads.workerData` is set to a copy of the
  `workerData` property of the second argument to the `Worker()` constructor.
  In the main thread, this property is always `null`. It use structured clone algorithm as well

- By default, `process.env` in a worker thread is a copy of `process.env` in the
  parent thread but you can specify a custom set of environment variable by setting the
  `env` property of the second argument to the `Worker()` constructor. As a special (and potentially dangerous) case, the parent
  thread can set `env` property to `threads.SHARE_ENV` which will cause the two threads to
  share a single set of environment variables so that a change in one thread is visible
  in the other.

  - ```js
    const threads = require("worker_threads");
    const w = new threads.Worker(
      `
    const threads = require("worker_threads")
    process.env.invalidenvvar = "ICHANGED"
    threads.parentPort.postMessage(threads.isMainThread);
    `,
      { eval: true, env: threads.SHARE_ENV },
    );
    w.on("message", console.log);
    setTimeout(() => {
      console.log(process.env.invalidenvvar);
    }, 2000);
    ```

- By default, the `process.stdin` stream in a worker never has any readable data on it.
  You can change this default by passing `stdin: true` in the second argument to the
  `Worker()` constructor. Any data that parent writes to `worker.stdin` becomes readable
  on `process.stdin` in the worker

  - ```js
    const w = new threads.Worker(
      `
    const threads = require("worker_threads")
    threads.parentPort.postMessage(threads.isMainThread);
    process.stdin.pipe(process.stdout)
    `,
      { eval: true, stdin: true },
    );
    w.on("message", console.log);
    process.stdin.pipe(w.stdin);
    ```

- By default, the `process.stdout` and `process.stderr` streams in the worker
  are simply piped to the corresponding streams in the parent thread. You can override
  this default by passing `stdout: true` and `stderr: true` in the second argument
  to the `Worker()` constructor. If you do this, `worker.stdout` and `worker.stderr`
  become readable streams.

- If a worker thread calls `process.exit()`, only the thread exits, not the entire process.

- Worker threads are not allowed to change shared state of the process they are part of.
  Functions like `process.chdir()` and `process.setuid()` will throw Exceptions when invoked
  from a worker.

- Operating system signals (like `SIGINT`) are only delivered to the main thread; they
  cannot be received or handled in worker threads.

#### 16.11.3 Communication Channels and MessagePorts

The worker thread API also allows the creation of custom communication channels using
the MessageChannel API defined by web browser and covered in [15.13.5](#15.13.5).

```js
const threads = require("worker_threads");

const ch1 = new threads.MessageChannel();
ch1.port2.on("message", console.log);
ch1.port1.postMessage("hello");
```

You can also call `close()` on either port to break the connection between the two
ports and to signal no more messages will be exchanged. When `close()` is called on
either port, a "close" event is delivered to both ports.

See below section to understand how to transfer one of the two ports from the thread to
the worker.

#### 16.11.4 Transferring MessagePorts and TypedArarys

The `postMessage()` function can handle MessagePort objects, but only as a special
case using a special technique. The `postMessage()` method (of a Worker object,
of `threads.parentPort`, or of any MessagePort object) takes an optional second argument. This
argument called `transferList`, is an array of objects that are to be **transferred** between
threads rather than being copied.

A MessagePort object cannot be copied by the structured clone algorithm, but it can be
transferred. If the first argument to `postMessage()` has include one or more
MessagePorts, then those MessagePort objects **must also appear as members of the
array passed as the second argument**. The key thing to understand, however, **once a
value is transferred, it can no longer be used in the thread that called `postMessage()`**

```js
const threads = require("worker_threads");

const ch2 = new threads.MessageChannel();
const worker2 = new threads.Worker(
  `
  const threads = require("worker_threads");
  let port;
  threads.parentPort.on("message", (data) => {
    if (data.command === "changeChannel") {
      if (port) port.close()
      port = data.data;
      port.on("message", (msg) => {
        console.log(msg);
        port.postMessage("Yes I can");
      });
    }
  });
`,
  { eval: true },
);

worker2.postMessage({ command: "changeChannel", data: ch2.port1 }, [ch2.port1]);
ch2.port2.postMessage("Can you hear me now?");
ch2.port2.on("message", console.log);
```

MessagePort objects are not the only ones that can be transferred. If you call
`postMessage()` with a typed array as the message (or contained within message), and
include the ArrayBuffer that backs the array in the second argument, it would transfer
a typed array instead of copying it.

```js
const worker3 = new threads.Worker(
  `
  const threads = require("worker_threads");
  threads.parentPort.on("message", (pixel) => {
    pixel[0] = 1;
    threads.parentPort.postMessage(pixel, [pixel.buffer])
  });
`,
  { eval: true },
);
let pixels = new Uint32Array(1);
worker3.on("message", (pixel) => {
  console.log(pixel);
});
worker3.postMessage(pixels, [pixels.buffer]);
```

As with transferred MessagePorts, a transferred typed array becomes unusable once
transferred. No exceptions are thrown if you attempt to use a MessagePort or typed
array that has been transferred; these objects simply stop doing anything when you interact
with them.

#### 16.11.5 Sharing TypeArrays Between Threads

In addition to transferring typed arrays between threads, it is actually
possible to share a typed array between threads. Simply create a SharedArrayBuffer
of the desired size and then use that buffer to create a typed array. When
a typed array that is backed by a SharedArrayBuffer is passed via `postMessage()`,
the underlying memory will be shared between the threads. **You should not include
the shared buffer in the second argument to `postMessage()` in this case**.

**You really should not do this because JavaScript was never designed with thread
safety in mind and multithreaded programming is very difficult to get right**.

```js
let sharedBuffer = new SharedArrayBuffer(4);
let sharedArray = new Int32Array(sharedBuffer);
const worker4 = new threads.Worker(
  `
  const threads = require("worker_threads");
  threads.parentPort.on("message", (buf) => {
    buf[0] = 1;
  });
`,
  { eval: true },
);
worker4.postMessage(sharedArray);
setTimeout(() => console.log(sharedArray), 1000);
```

One scenario in which it might be reasonable to use a SharedArrayBuffer is
when the two threads operate on entirely **separate sections** of the shared memory.
A parallel merge sort could be done like this. Or some kinds of image-processing
algorithms are also suitable for this approach: multiple threads working on disjoint
regions of the image.

If you really must allow multiple threads to access the same region of a shared array,
you can take one step toward thread safety with the functions defined by the Atomic object.
Atomics was added to JavaScript when SharedArrayBuffer was to define atomic operations
on the elements of a shared array. They ensures that no other thread can read or write
the value while the operation is taking place.

```js
const threads = require("worker_threads");

if (threads.isMainThread) {
  let sharedBuffer1 = new SharedArrayBuffer(4);
  let sharedArray1 = new Int32Array(sharedBuffer1);
  let worker5 = new threads.Worker(__filename, { workerData: sharedArray1 });
  worker5.on("online", () => {
    for (let i = 0; i < 10_000_000; i++) {
      Atomics.add(sharedArray1, 0, 1);
    }
  });
  worker5.on("message", (msg) => {
    console.log(Atomics.load(sharedArray1, 0));
  });
} else {
  let sharedArray = threads.workerData;
  for (let i = 0; i < 10_000_000; i++) {
    Atomics.add(sharedArray, 0, 1);
  }
  threads.parentPort.postMessage("done");
}
```

Result: 20000000

In most real-world programs, multiple array elements are often related to one another
and some kind of higher-level thread synchronization is required. The low-level
`Atomics.wait()` and `Atomics.notify()` function can help with this, but
a discussion of their use is out of scope for this book.

## 17 - JavaScript Tools and Extensions

### 17.1 Linting With ESLint

In programming, the term _lint_ refers to code that, while technically correct, is
unsightly, or a possible bug, or suboptimal in some way. A _linter_ is tool
for detecting lint in your code, and _linting_ is the process of
running a linter on your code (and then fixing your code to remove the lint so that the
linter no longer complains).

The most commonly used linter for JavaScript today is ESLint.

```bash
$ npm install -g eslint
$ eslint linty.js
```

### 17.2 JavaScript Formatting with Prettier

One of the reasons that some projects use linters is to enforce consistent coding
style so that when a team of programmers is working on a shared codebase,
they use compatible code conventions. This includes code indentation rules,
but can also include things like what kind of quotation marks are preferred and
whether there should be a space between the `for` keyword and the open parenthesis
that follows it.

A modern alternative to enforcing code formatting rules via a linter is
to adopt a tool like Prettier to automatically parse and reformat all of your code.

```js
$ prettier factorial.js
```

If you invoke Prettier with the `--write` option, it will simply reformat the
specified file in place rather than printing a reformatted version
If you use `git` to manage your source code, you can invoke Prettier with
the `--write` option in a commit hook so that code is automatically formatted before
being checked in.

Prettier is particularly powerful if you configure your code editor to run it
automatically every time you save a file.

Prettier is configurable, but it only has a few options.
You can select the maximum line length, the indentation amount, whether semicolons
should be used, whether strings should be single- or double-quoted, and
a few other things. In general, Prettier's default options are quite reasonable.

### 17.3 Unit Testing with Jest

There are a lot of test tools and libraries for JavaScript, and many are written
in a modular way so that it is possible to pick one library as your test runner,
another library for assertions, and a third for mocking. In this section, however,
we'll describe Jest, which is popular framework that includes everything you need
in a single package.

Example function:

```js
const getJSON = require("./getJSON.js");

module.exports = async function getTemperature(city) {
  let c = await getJSON(
    `https://globaltemps.example.com/api/city/${city.toLowerCase()}`,
  );
  return (c * 5) / 9 + 32;
};
```

Below code defines a mock implementation of `getJSON()` so that
the test does not actually make a network request.
And because `getTemperature()` is an async function, the tests are
async as well -- it can be tricky to test asynchronous functions,
but Jest makes it relatively easy:

> main.js

```js
const getJSON = require("./getJSON.js");

module.exports = async function getTemperature(city) {
  let c = await getJSON(
    `https://globaltemps.example.com/api/city/${city.toLowerCase()}`,
  );
  return (c * 9) / 5 + 32;
};
```

> x.test.js

```js
const getTemperature = require("./main.js");

// mock the getJSON() module that getTemperature() depends on
jest.mock("./getJSON");
const getJSON = require("./getJSON.js");

// Tell the mock getJSON() function to return an already resolved Promise
// with fulfillment value 0
getJSON.mockResolvedValue(0);

describe("getTemperature()", () => {
  test("Invokes the correct API", async () => {
    let expectedURL = "https://globaltemps.example.com/api/city/vancouver";
    let t = await getTemperature("Vancouver");
    expect(getJSON).toHaveBeenCalledWith(expectedURL);
  });
  test("Converts C to F correctly", async () => {
    getJSON.mockResolvedValue(0); // if getJSON returns 0C
    expect(await getTemperature("x")).toBe(32); // We expect 32F

    getJSON.mockResolvedValue(100);
    expect(await getTemperature("x")).toBe(212); // We expect 212F
  });
});
```

Result:

```bash
$ npx jest
 PASS  ./jest.test.js
  getTemperature()
    ✓ Invokes the correct API (2 ms)
    ✓ Converts C to F correctly (1 ms)

Test Suites: 1 passed, 1 total
Tests:       2 passed, 2 total
Snapshots:   0 total
Time:        0.277 s, estimated 1 s
Ran all test suites.
```

And, as a bonus, if we add the `--coverage` argument when we
invoke `jest`, it will compute and display the code coverage for our tests:

```bash
$ npx jest --coverage
 PASS  ./jest.test.js
  getTemperature()
    ✓ Invokes the correct API (2 ms)
    ✓ Converts C to F correctly

------------|---------|----------|---------|---------|-------------------
File        | % Stmts | % Branch | % Funcs | % Lines | Uncovered Line #s
------------|---------|----------|---------|---------|-------------------
All files   |      24 |        0 |   14.28 |      24 |
 getJSON.js |    9.52 |        0 |       0 |    9.52 | 4-31
 jest.js    |     100 |      100 |     100 |     100 |
------------|---------|----------|---------|---------|-------------------
Test Suites: 1 passed, 1 total
Tests:       2 passed, 2 total
Snapshots:   0 total
Time:        0.365 s, estimated 1 s
Ran all test suites.
```

Running our test gave us 100% code coverage for the module we were testing,
which is exactly what we wanted. It only gave us partial coverage of `getJSON()`,
but we mocked that module and were not trying to test it, so that is expected.

### 17.4 Package Management with npm

A package manager makes it easy to find and install third-party packages. Just
as importantly, a package manager keeps track of what packages your code depends on
and saves this information into a file so that when someone else wants
to try your program, they can download your code and your list of dependencies,
then use their own package manager to install all the third-party packages that
your code needs.

npm is the package manager that is bundled with Node.

If you are trying out someone else's JavaScript project, then one of the first things
you will often do after downloading their code is to type `npm install`.
This reads the dependencies listed in the _package.json_ file and downloads the third-party packages that the project needs and saves them into a _node_modules/_ directory.

You can also type `npm install <package-name>` to install a particular package
to your project.

```bash
$ npm install express
```

The other kind of dependency is on developer tools that are needed by developers
who want to work on your project, but aren't actually needed to run the code: Prettier, Jest, etc.
They are "dev dependency", and you can install and record one of these with
`--save-dev`

```bash
$ npm install --save-dev prettier
```

**Tip**: use `npx prettier` to run locally installed tools

**Tip**: If you use `npx` to invoke a tool that has not been installed yet, it will
install it for you.

Sometimes you might want to install developer tools globally so that
they are accessible anywhere even for code that is not part of a formal
project with a _package.json_ file and a _node_modules/_ directory:

```bash
$ npm install -g eslint jest
```

In addition to the "install" command, npm supports "uninstall" and "update"
commands, which do what their names say. npm also has an interesting
"audit" command that you can use to find and fix vulnerabilities in your dependencies

```bash
$ npm audit --fix
```

The company behind npm also maintains the _https://npmjs.com_ package repository,
which holds hundreds of thousands of open source packages. But you don't have to
use the npm package manager to access this repository of packages. Alternatives
include yarn and pnpm.

### 17.5 Code Bundling

Programmers use a code-bundler tool that starts at the main entry point of the
program and follows the tree of `import` directives to find all modules that
program depends on. In then combines all of those individual module files into a single
bundle of JavaScript code and rewrites the `import` and `export` directives to make
the code work in this new form. The result is a single file of code that can be
loaded into a web browser that does not support modules.

ES6 modules are nearly universally supported by web browsers today, but web developers
still tend to use code bundlers, at least when releasing production code. Developers
find that user experience is best when a single medium-sized bundle of code is loaded
when a user first visits a website than when many small modules are loaded.

---

> **Note**

Keep in mind that there is one variable that is completely under your control to improve
web performance -- code size. Less JavaScript code will always load and run faster
than more JavaScript code!

---

There are a number of good JavaScript bundler tools available. Commonly used bundlers
include webpack, Rollup and Parcel. The basic features of bundlers are more or less the same,
and they are differentiated based on how configurable they are or how easy they are
to use.

In addition to performing basic bundling, bundler tools can also provide some
additional features:

- Some programs have more than one entry point. Bundlers generally allow you to create
  one bundle per entry point or to create a single bundle that supports multiple entry points.

- Programs can use `import()` [10.3.6](#10.3.6) in its functional form instead of
  its static form to dynamically load modules. Doing this is often a good way to
  improve the startup time for your program.
  Bundler tools that support `import()` may be able to produce multiple output bundles:
  one to load at startup time, and one or more that are loaded dynamically when needed.
  If the dynamically loaded modules share dependencies then it becomes tricky to figure
  out how many bundles to produce, and you are likely to have to manually
  configure your bundler to sort this out.

- Bundlers can generally output a _source map_ file that defines a mapping between
  the lines of code in the bundle and the corresponding lines in the original source files.
  This allows browser developer tools to automatically display JavaScript errors at their
  original unbundled locations.

- Sometimes when you import a module into your program, you only use a few of its
  features. A good bundler tool can analyze the code to determine which parts
  are unused and can be omitted from the bundles.

- Bundlers typically have a plug-in-based architecture and support plug-ins that
  allow importing and bundling "modules" that are not actually files of JavaScript code.
  e.g. Code bundlers can be configured to allow you to move a json data structure into
  a seperate JSON file and import it into your program with a declaration like `import
widgets from "./big-widget-list.json"`. Similarly, web developers who embed CSS into their
  JavaScript programs can use bundler plug-ins that allow them to import CSS files with an `import` directive.
  **Note, however, that if you import anything other than a JavaScript file, you are
  using a nonstandard JavaScript extension and making your code dependent on the
  bundler tool**.

- In a language like JavaScript that does not require compilation, running a bundler
  tool feels like a compilation step, and it is frustrating to have to run a bundler after
  every code edit before you can run the code in your browser. **Bundle typically
  support filesystem watchers that detect edits to any files in a project directory
  and automatically regenerate the necessary bundles. With this feature in place you
  can typically save your code and then immediately reload your web browser window
  to try it out.**

- Some bundlers also support a "hot module replacement" mode for developers where
  each time a bundle is regenerated, it is automatically loaded into the browser.
  When this works, it is a magical experience for developers, but there are some
  tricks going on under the hood to make it work, and it is not suitable for all projects.

### 17.6 Transpilation with Babel

Babel is a tool that compiles JavaScript written using modern language features
into JavaScript that does not use those modern language features. Because it compiles
JavaScript to JavaScript, Babel is sometimes called a "transpiler".
Babel was created so that web developers could use the new language features of ES6
and later while still targeting web browsers that only supported ES5.

In general, the code output by Babel is not meant to be human readable. Like bundler
tools, however, Babel can produce source maps that map transformed code locations
back to their original source locations.

Like most of the other tools described in this chapter, you can install Babel with
npm and run it with npx. Babel reads a _.babelrc_ configuration file
that tells it how you would like your JavaScript code transformed. One
of Babel's interesting presets is for code compression by minification (stripping
comments and white space, renaming variables, and so on).

If you use Babel and a code-bundling tool, you may be able to set up the code bundler
to automatically run Babel on your JavaScript files as it builds the bundle for you.
If so, this can be convenient option because it simplifies the process of producing
runnable code. Webpack, for example, supports a "babel-loader" module that you can
install and configure to run Babel on each JavaScript module as it is bundled up.

Even though there is less need to transform the core JavaScript language today,
**Babel is still commonly used to support nonstandard extensions to the language,
and we'll describe two of these language extensions in the sections that follow**.

### 17.7 JSX: Markup Expressions in JavaScript

---

> Setting up a React

```bash
$ npm install react react-dom react-scripts
$ cat package.json
{
  "name": "react.dev",
  "version": "0.0.0",
  "main": "/src/index.js",
  "scripts": {
    "start": "react-scripts start",
    "build": "react-scripts build",
    "test": "react-scripts test --env=jsdom",
    "eject": "react-scripts eject"
  },
  "dependencies": {
    "react": "^18.0.0",
    "react-dom": "^18.0.0",
    "react-scripts": "^5.0.0"
  },
  "devDependencies": {}
}
$ tree -I node_modules
.
├── babel.config.json
├── package.json
├── package-lock.json
── public
│   ├── index.html
└── src
    ├── App.js
    ├── index.js
    └── styles.css
$ npm start
```

---

JSX is an extension to core JavaScript that uses HTML-style syntax to define
a tree of elements. JSX is most closely associated with the React framework for
user interfaces on the web.

If you use JSX, you will need to use Babel (or similar tool) to compile
JSX expressions into regular JavaScript. The transformation is simple enough
that some developers choose to use React without using JSX. Babel transform
the JSX expression in this assignment statement into a simple function call:

```js
// JSX
let line = <hr />;

// Babel transformed
let line = React.createElement("hr", null);
```

JSX syntax is HTML-like, and like HTML elements, React elements can have attributes
like these:

```js
let image = <img src="logo.png" alt="The JSX logo" hidden />;
```

When an element has one or more attributes, they become properties of
an object passed as the second argument to `createElement()`

```js
let image = React.createElement("img", {
  src: "logo.png",
  alt: "The JSX logo",
  hidden: true,
});
```

Like HTML elements, JSX elements can have strings and other elements as children:

```js
let sidebar = (
  <div className="sidebar">
    <h1>Title</h1>
    <hr />
    <p>This is the sidebar content</p>
  </div>
);
```

Those children are passed as the third and subsequent arguments in `createElement()`

```js
let sidebar = React.createElement(
  "div",
  { className: "sidebar" },
  React.createElement("h1", null, "Title"),
  React.createElement("hr", null),
  React.createElement("p", null, "This is the side content"),
);
```

Within a JSX expression, text within curly braces is interpreted as plain JavaScript.
These nested expressions are allowed as attribute values and as child elements.

```js
function sidebar(className, title, content, drawLine = true) {
  return (
    <div className={className}>
      <h1>{title}</h1>
      {drawLine && <hr />}
      <p>{content}</p>
    </div>
  );
}
```

This example code is translated by Babel into the following:

```js
function sidebar(className, title, content, drawLine = true) {
  return React.createElement(
    "div",
    { className: className },
    React.createElement("h1", null, title),
    drawLine && React.createElement("hr", null),
    React.createElement("p", null, content),
  );
}
```

**Tip**: If you pass `false`, then no `<hr/>` element is ever created.
**This use of the && operator is a common idiom in JSX to conditionally include
or exclude a child element depending on the value of some other expression.**

You are not limited to simple values like the string and boolean values in the
preceding example. In fact, it is quite common in React programming to use
objects, arrays, and functions.

```js
function list(items, callback) {
  return (
    <ul style={{ padding: 10, border: "solid red 4px" }}>
      {items.map((item, index) => (
        <li onClick={() => callback(index)} key={index}>
          {item}
        </li>
      ))}
    </ul>
  );
}
```

Note that double curly braces are required to use an object literal as the value
of the `style` attribute on the `<ul>` element. The JSX code
compiles to the following pure JavaScript code:

```js
function list(items, callback) {
  return React.createElement(
    "ul",
    { style: { padding: 10, border: "solid red 4px" } },
    items.map((item, index) =>
      React.createElement(
        "li",
        { onClick: () => callback(index), key: index },
        item,
      ),
    ),
  );
}
```

One other use of object expression in JSX is with the object spread operator to specify
multiple attributes at once.

```js
let hebrew = { lang: "he", dir: "rtl" };
let shalom = (
  <span className="emphasis" {...hebrew}>
    שלום
  </span>
);
```

Babel compiles this to use an `_extends()` function (omitted here) that combines
that `className` attribute with the attributes contained in the `hebrew` object.

```js
let shalom = React.createElement(
  "span",
  _extends({ className: "emphasis" }, hebrew),
  "\u05E9\u05DC\u05D5\u05DD",
);
```

Finally, there is one more important feature of JSX that we have not covered yet.
**All JSX elements begin with an identifier immediately after the open angle bracket.
If the first letter of this identifier is lowercase, then the identifier is passed
to `createElement()` as a string**. But if the first letter of the identifier is uppercase,
then it is treated as an actual identifier, and it is the JavaScript value of that identifier that is passed as the first argument to `createElement()`. This means that the
JSX expression `<Math/>` compiles to JavaScript code that passes the global Math
object to `React.createElement()`

For React, this ability to pass non-string values as the first argument to `createElement()`
enables the creation of components. The simplest way to define a new component in React
is to write a function that takes a "props object" as its argument and returns a
JSX expression. A _props object_ is simply a JavaScript object that represents attribute
values, like the objects that are passed as the second argument to `createElement`.

```js
function Sidebar(props) {
  return (
    <div>
      <h1>{props.title}</h1>
      {props.drawLine && <hr />}
      <p>{props.content}</p>
    </div>
  );
}
```

```js
let sidebar = <Sidebar title="Something snappy" content="Something wise" />;
```

The `<Sidebar/>` element compiles like this:

```js
let sidebar = React.createElement(Sidebar, {
  title: "Something snappy",
  content: "Something wise",
});
```

### 17.8 Type Checking with Flow

Flow is a language extension that allows you to annotate your JavaScript code with type
information, and a tool for checking your JavaScript code for type errors.
To use Flow, you start writing code using the Flow language extension to add
type annotations. Then you run the Flow tool to analyze your code and report type errors.
Once you have fixed the errors and are ready to run the code, you use Babel to
strip the Flow type annotations out of your code.

---

> Typescript Versus Flow

TypeScript is a very popular alternative to Flow. TypeScript is an extension of
JavaScript that adds types as well as other language features. The TypeScript
comppiler "tsc" compiles TypeScript programs into JavaScript programs and in the
process analyzes them and reports type errors in much the same the way that Flow
does. `tsc` is not a Babel plugin: it is its own standalone compiler.

TypeScript was released in 2012, before ES6, when JavaScript did not have a `class`
keyword or a `for/of` loop or modules or Promises. Flow is a narrow language extension
that adds type annotations to JavaScript and nothing else. TypeScript, by contrast,
was very much designed as a new language. As its name implies, adding types to
JavaScript is the primary purpose of TypeScript. But types are not the only
feature that TypeScript adds to JavaScript: the TypeScript language has `enum`
and `namespace` keywords that simply do not exist in JavaScript.

---

#### 17.8.1 Installing and Running Flow

```bash
$ npm install -g flow-bin
```

Before using Flow to do type checking, the first time run it as `flow init`
in the root to create a `.flowconfig` configuration file (although you may never need
to add anything to this file, but Flow needs it to know where your project root is)

Flow will only report type errors for the files that have "opted in" to type checking
by adding a `// @flow` comment at the top of the file.

```js
// @flow
let i = { r: 0, i: 1 };
for (i = 0; i < 10; i++) {
  console.log(i);
}
i.r = 1;
```

Run `flow`, some errors prompts for your to fix:

```bash
$ flow
...
Error ┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈ test.js:6:3

Cannot assign 1 to i.r because property r is missing in number [1].
[incompatible-use]

 [1] 3│ for (i = 0; i < 10; i++) {
     4│   console.log(i);
     5│ }
     6│ i.r = 1;
     7│

Found 4 errors
```

We can fix it to:

```js
for (let i = 0; i < 10; i++) {
```

Here is another error:

```js
Error ┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈ test.js:2:15

Missing an annotation on x. [missing-local-annot]

     1│ // @flow
     2│ function size(x) {
     3│   return x.length
     4│ }
     5│ let s = size(1000)
```

#### 17.8.2 Using Type Annotations

```bash
$ npm init -y
$ npm install --save-dev @babel/preset-react @babel/core @babel/cli @babel/preset-flow nodemon
$ cat package.json
{
  "name": "flow",
  "version": "1.0.0",
  "description": "",
  "main": "lib/test.js",
  "scripts": {
    "start": "nodemon --exec 'clear && flow; babel src --out-dir lib && node lib/test.js' --watch src"
  },
  "keywords": [],
  "author": "",
  "license": "ISC",
  "devDependencies": {
    "@babel/cli": "^7.24.1",
    "@babel/core": "^7.24.4",
    "@babel/preset-flow": "^7.24.1",
    "@babel/preset-react": "^7.24.1",
    "nodemon": "^3.1.0"
  }
}
$ cat babel.config.json
{
  "presets": [
    [
      "@babel/preset-react",
      {
        "development": true
      }
    ],
    "@babel/preset-flow"
  ]
}
$ npm start
```

```js
let message: string = "Hello world"
let flag: boolean = false;
let n: number = 42

function returnNothing(): void {
}

const size = (s: string): number => s.length
```

Flow knows both the type of the variable and that you have expressed the intent that the
variable should always be of that type. So if you use the type annotation, Flow
will flag an error if you ever assign a value of a different type to that variable.
Type annotations for variables are also particularly useful if you tend to declare
all your variables up at the top of a function before they are used.

And important thing to understand about Flow is that the JavaScript value `null`
has the Flow type `null` **and the JavaScript value `undefined` has the flow type `void`**.

If you want to allow `null` and `undefined` as legal values for a variable or
function argument, simply prefix the type with a question mark.

```js
function size(s: ?string): number {
  if (s === null || s === undefined) { return -1;
  } else {
    return s.length
  }
}

// or
return s ? s.length : -1;
```

Question marks can also appear after a parameter name to indicate that the
parameter itself is optional. So if we changed the declaration to:

```js
s?: string
```

That would mean it is OK to call `size()` with no arguments (or with `undefined`). But note that in this case,
`null` is not a legal value.

#### 17.8.3 Class Types

In addition to the primitive types that Flow knows about, it also knows about all
of JavaScript's built-in classes and allows you to use class name as types.

```js
export function dateMatches(d: Date, p: RegExp): boolean {
  return p.test(d.toISOString())
}
```

However, in order to make this work, Flow does require you to use type annotations
**in the class**.

```js
declare class Date {
    constructor(): void;
    constructor(value: number | Date | string): void;
    constructor(year: number, month: number, day?: number, hour?: number, minute?: number, second?: number, millisecond?: number): void;
    /** Gets the day-of-the-month, using local time. */
    getDate(): number;
    /** Gets the day of the week, using local time. */
    getDay(): number;
    /** Gets the year, using local time. */
    getFullYear(): number;

    ...
}
```

#### 17.8.4 Object Types

```js
export default function distance(point: {x: number, y:number}): number {
  return Math.hypot(point.x, point.y)
}
```

As with any type, you can add a question mark at the front to indicate that `null`
and `undefined` should also be allowed.

You can follow any of the property names with a question mark to indicate that
that property is optional and may be omitted.

```js
{x: number, y: number, z?: number}
```

Flow tolerates extra properties, if you want Flow to strictly enforce that
an object does not have properties other than those explicitly declared in its type,
you can declare an _exact object type_ by adding vertical bars to the curly braces:

```js
{| x:number, y:number|}
```

![Flow_Object_Exact](images/flow_object_exact.png)

JavaScript's objects are sometimes used as dictionaries or string-to-value maps.
When used like this, the property names are not known in advance. You can declare this
kind of data structure like this:

```js
const cityLocations: {[string]: {longitude: number, latitude:number}} = {
  "Seattle": {longitude: 47.6062, latitude: -122.3321}
}
```

#### 17.8.5 Type Aliases

```js
export type Point = {
  x: number,
  y: number
}

export default function distance(point: Point): number {
  return Math.hypot(point.x, point.y)
}
```

To import the type:

```js
import type Point from "./distance.js"
```

Keep in mind, though, that `import type` is a Flow language extension
and not a real JavaScript import directive. They are stripped out of the code
before it ever runs.

Finally, it is worth noting that instead of defining a name for a Flow object type
that represents a point, it would probably be simpler and cleaner to just define a Point
class and use that class as the type.

#### 17.8.6 Array Types

```js
function average(data: Array<number>) {
  let sum = 0;
  for (let x of data) sum += x;
  return sum / data.length
}
```

We could have written `number[]` instead of `Array<number>`. The author prefers the
angle bracket notation because, as we'll see.

Flow has a different syntax for describing the type of a _tuple_: an array
with a fixed number of elements, each of which may have a different type.

```js
function getStatus(): [number, string] {
  return [getStatusCode(), getStatusMessage()]
}
```

Use destructuring assignment

```js
let [code, message] = getStatus();
```

Flow's type-aliasing capabilities, make tuple easy enough to work
with that you might consider them as an alternative to classes for simple datatypes:

```js
export type Color = [number, number, number, number] // [r, g, b, opacity]

function gray(level: number): Color {
  return [level, level, level, 1]
}

function fade([r,g,b,a]: Color, factor: number): Color {
  return [r, g, b, a/factor]
}

let [r,g,b,a] = fade(gray(75), 3)
```

If we want the function to be able to accept an array of any length without restricting
the type:

```js
function size(s: Array<mixed>): number {
  return s.length
}
```

If our function actually indexed the array and attempted to use
any of those elements. Flow would insist that we use `typeof` checks or other
tests to determine the type of the element before performing any unsafe operation
on it. (**If you are willing to give up on type checking, you can also use `any`
instead of `mixed`**)

#### 17.8.7 Other Parameterized Types

Array is not the only JavaScript class that is parameterized (type specified inside
angle brackets)

> Set

```js
function double(s: Set<number>): Set<number> {
  let doubled: Set<number> = new Set();
  s.forEach((v) => {
    doubled.add(v * 2);
  });
  return doubled;
}

console.log(double(new Set([1,2,3])))
```

> Map

```js
import type {Color} from "./Color.js"

let colorNames: Map<string, Color> = new Map([
  ["red", [1,0,0,1]],
  ["green", [0,1,0,1]],
  ["blue", [0,0,1,1]],
])
```

Flow lets you define type parameters **for your own classes as well**.
When the user of this class declares a variable of type Result, they will specify the
actual types to substitute for E and V:

```js
let result: Result<TypeError, Set<string>>;
```

And here is how the parameterized class is defined:

```js
export class Result<E, V> {
  error: ?E;
  value: ?V;

  constructor(error: ?E, value: ?V) {
    this.error = error;
    this.value = value;
  }

  threw(): ?E {
    return this.error;
  }

  returned(): ?V {
    return this.value;
  }

  get(): V {
    if (this.error) {
      throw this.error;
    } else if (this.value === null || this.value === undefined) {
      throw new TypeError("Error and value must not both be null");
    } else {
      return this.value;
    }
  }
}
```

And you can even define type parameters for functions:

```js
function zip<A, B>(a: Array<A>, b: Array<B>): Array<[?A, ?B]> {
  let result: Array<[?A, ?B]> = [];
  let len = Math.max(a.length, b.length);
  for (let i = 0; i < len; i++) {
    result.push([a[i], b[i]]);
  }
  return result;
}

// [ [ 1, 'a' ], [ 2, 'b' ], [ 3, 'c' ], [ 4, undefined ]
let pairs: Array<[?number, ?string]> = zip([1,2,3,4],["a", "b", "c"])
```

#### 17.8.8 Read-Only Types

Flow defines some special parameterized "utility types" that have names beginning
with $. Most of these types have advanced use cases that we are not going to cover here.
But two of them are quite useful in practice.

If you have object type T and want to make a read-only version of that type, just write
`$ReadOnly<T>`. Similarly, you can write `$ReadOnlyArray<T>` to describe a read-only
array with elements of type T.

The reason to use these types is not because they can offer any guarantee that
object or array can't be modified, but **it allows you to catch bugs caused
by unintentional modifications**.

```js
export type Point = {
  x: number,
  y: number,
};

export default function distance(point: $ReadOnly<Point>): number {
  point.x = 123;
  return Math.hypot(point.x, point.y);
}
```

![Point_Readonly](images/point_readonly.png)

```js
function average(data: $ReadOnlyArray<number>) {
  let sum = 0;
  for (let x of data) sum += x;
  return sum / data.length;
}
```

#### 17.8.9 Function Types

```js
export type FetchTextCallback = (?Error, ?number, ?string) => void

export default function fetchText(url: string, callback: FetchTextCallback) {
  let status = null;
  fetch(url)
    .then(response => {
      status = response.status
      return response.text()
    })
    .then(body => {
      callback(null, status, body)
    })
    .catch(error => {
      callback(error, status, null)
    })
}
```

#### 17.8.10 Union Types

With Flow, we need a way to express a type that allows arrays, Sets and Maps, but
doesn't allow values of any other type.

Flow calls types like this _Union types_ and allows you to express them by
simply listing the desired types and separating them with vertical bar characters:

```js
function size(collection: Array<mixed>|Set<mixed>|Map<mixed,mixed>): number {
  if (Array.isArray(collection)) {
    return collection.length
  } else {
    return collection.size;
  }
}
```

**Tip**: Union types can be read using the word "or"
And now you can see that a `?` prefix is simply a shortcut of adding `|null|void` suffix
to a type.

#### 17.8.11 Enumerated Types and Discriminated Unions

Flow allows you to use **primitive literals** as types that consist of that one
single value.

```js
type Answer = "yes" | "no"
type Digit = 0|1|2|3|4|5|6|7|8|9
```

**Important**: You need to understand that **only literal values are allowed**

```js
let a: Answer = "Yes".toLowerCase(); // Error: can't assign
let d: Digit = 3 + 4; // Error: can't assign
```

A union type of literal types like `Answer` and `Digit` is an example of
an _enumerated type_, or _enum_. A canonical use case for enum types is to
represent the suits of playing cards:

```js
type Suit = "Clubs" | "Diamonds" | "Hearts" | "Spades"
```

A more relevant example might be HTTP status codes:

```js
type HTTPStatus =
  | 200
  | 304
  | 403
  | 404
```

One of the pieces of advice that new programmers often hear is to
avoid using literals in their code and to instead define symbolic constants
to represent those values. With Flow, this advice does not always apply:
if you annotate a variable with the type Suit, and then try to assign
a misspelled suit to it, Flow will alert you to the error.

Another important use for literal types is the creation of _discriminated unions_.
If you want to create a union of Object types, you can make these types easy
to discriminate by using a literal type within each of the individual Object types.

```js
export type ResultMessage = {
  messageType: "result",
  result: Array<ReticulatedSpline>,
}

export type ErrorMessage = {
  messageType: "error",
  error: Error
}

export type StatisticMessage = {
  messageType: "stats",
  splinesReticulated: number,
  splinesPerSecond: number
}

export type WorkerMessage = ResultMessage | ErrorMessage | StatisticMessage

function handleMessageFromReticulator(message: WorkerMessage) {
  if (message.messageType === "result") {
    console.log(message.result)
  } else if (message.messageType === "error") {
    throw message.error
  } else if (message.messageType === "stats") {
    console.info(message.splinesPerSecond)
  }
}
```
