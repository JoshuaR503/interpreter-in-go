## Lexer:

Note that we save l.ch in a local variable before calling l.readChar() again. This way we don’t lose the current character and can safely advance the lexer so it leaves the NextToken() with l.position and l.readPosition in the correct state. If we were to start supporting more two-character tokens in Monkey, we should probably abstract the behaviour away in a method called makeTwoCharToken that peeks and advances if it found the right token. Because those two branches look awfully similar. For now though == and != are the only two-character tokens in Monkey, so let’s leave it as it is and run our tests again to make sure it works:

## REPL

Sometimes the REPL is called “console”, sometimes “interactive mode”. The concept is the same: the REPL reads input, sends it to the interpreter for evaluation, prints the result/output of the interpreter and starts again. Read, Eval, Print, Loop.

## AST

In most interpreters and compilers the data structure used for the internal representation of the source code is called a “syntax tree” or an “abstract syntax tree” (AST for short). The “abstract” is based on the fact that certain details visible in the source code are omitted in the AST. Semicolons, newlines, whitespace, comments, braces, bracket and parentheses – depending on the language and the parser these details are not represented in the AST, but merely guide
the parser when constructing it.

A fact to note is that there is not one true, universal AST format that’s used by every parser. Their implementations are all pretty similar, the concept is the same, but they differ in details. The concrete implementation depends on the programming language being parsed.

## On parsers

So, this is what parsers do. They take source code as input (either as text or tokens) and produce a data structure which represents this source code. While building up the data structure, they unavoidably analyse the input, checking that it conforms to the expected structure. Thus the process of parsing is also called syntactic analysis.

## Parser Generator

Parser generators are tools that, when fed with a formal description of a language, produce parsers as their output. This output is code that can then be compiled/interpreted and itself fed with source code as input to produce a syntax tree.

The majority of them use a context-free grammar (CFG) as their input. A CFG is a set of rules that describe how to form correct (valid according to the syntax) sentences in a language. The most common notational formats of CFGs are the Backus-Naur Form (BNF) or the Extended Backus-Naur Form (EBNF).

Parsing is one of the most well-understood branches of computer science and really smart people have already invested a lot of time into the problems of parsing. The results of their work are CFG, BNF, EBNF, parser generators and advanced parsing techniques used in them. Why shouldn’t you take advantage of that?

## Parsring approaches

There are two main strategies when parsing a programming language: top-down parsing or bottom-up parsing. A lot of slightly different forms of each strategy exist. For example, “recursive descent parsing”, “Early parsing” or “predictive parsing” are all variations of top down parsing.

The parser we are going to write is a recursive descent parser. And in particular, it’s a “top down operator precedence” parser, sometimes called “Pratt parser”, after its inventor Vaughan Pratt.

## Expressions vs Values

Expressions produce values, statements don’t. let x = 5 doesn’t produce a value, whereas 5 does (the value it produces is 5). A return 5; statement doesn’t produce a value, but add(5, 5) does. This distinction - expressions produce values, statements don’t - changes depending on who you ask, but it’s good enough for our needs.

# AST Structure Breakdown

- **`*ast.Program`**

  - Represents the top-level structure of the program's AST.
  - Contains a collection of statements that make up the program.

  - **Statements**

    - A list of statements in the program (e.g., `LetStatement`, `ReturnStatement`, etc.).

    - **`*ast.LetStatement`**

      - Represents a "let" statement, used for variable declarations.

      - **Name**
        - Points to an `Identifier`, representing the variable name in the "let" statement.
      - **Value**

        - Points to an `Expression`, representing the expression that assigns a value to the identifier.

      - **`*ast.Identifier`**

        - Represents the variable name in the "let" statement.
        - Connected to the `Name` field in the `LetStatement`.

      - **`*ast.Expression`**
        - Represents the value or the expression being assigned to the variable.
        - Connected to the `Value` field in the `LetStatement`.

# Understanding Recursive-Descent Parsing

Imagine you're trying to understand a sentence by breaking it down into smaller parts, like words and phrases. Recursive-descent parsing is a technique used by a parser (a part of a computer program) to do something similar with code.

### The Big Picture

1. **Starting Point:**

   - The process begins with a function called `parseProgram`. Think of this as the starting point where the parser begins its work.
   - This function creates the "root" of a tree structure that represents the entire code. This tree is called an Abstract Syntax Tree (AST), and it helps the program understand the structure of the code.

2. **Building the Tree:**

   - After creating the root, the parser moves on to building the "branches" or "child nodes" of the tree. These branches represent the different parts of the code, like statements or expressions.
   - The parser knows what type of branch to add based on the current "token" (a small piece of the code it's analyzing).

3. **Working Recursively:**

   - Sometimes, the parser needs to call itself to figure out more complex parts of the code. This is where it gets "recursive."
   - For example, if the code includes a mathematical expression like `5 + 5`, the parser first understands the `5 +` part. Then, it calls itself to figure out what comes next, which might be another expression, like `5 * 10`.

4. **Decision-Making:**
   - The parser continuously checks what the current token is and decides what to do next. It might call another function to handle a specific part of the code, or it might realize there's an error and stop.

### Why This Matters

- This method is powerful because it allows the parser to handle complex expressions and statements in a structured way.
- Although it sounds complicated, the basic idea is simple: the parser breaks down the code piece by piece, building a tree structure that represents the entire program.

If this explanation makes sense to you, you’re well on your way to understanding how our parser will work. The actual code for our `ParseProgram` method will follow these same principles, making it relatively easy to understand once you see it in action.

So, let’s dive in and start building!
