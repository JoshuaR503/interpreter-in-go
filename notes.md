## Lexer:

Note that we save l.ch in a local variable before calling l.readChar() again. This way we don’t lose the current character and can safely advance the lexer so it leaves the NextToken() with l.position and l.readPosition in the correct state. If we were to start supporting more two-character tokens in Monkey, we should probably abstract the behaviour away in a method called makeTwoCharToken that peeks and advances if it found the right token. Because those two branches look awfully similar. For now though == and != are the only two-character tokens in Monkey, so let’s leave it as it is and run our tests again to make sure it works:

## REPL

Sometimes the REPL is called “console”, sometimes “interactive mode”. The concept is the same: the REPL reads input, sends it to the interpreter for evaluation, prints the result/output of the interpreter and starts again. Read, Eval, Print, Loop.
