import ply.lex as lex

tokens = (
    'VAR',
    'IF',
    'ELSE',
    'ID',
    'CTE_STRING',
    'INT_TYPE',
    'LESS_GREAT',
    'FLOAT_TYPE',
    'INT',
    'FLOAT',
    'PROGRAM',
    "PRINT",
    'LPAREN',
    'RPAREN',
    'LBRACE',
    'RBRACE',
    'COMMA',
    'COLON',
    'SEMICOLON',
    'LESS_THAN',
    'GREATER_THAN',
    'MINUS',
    'ASSIGN',
    'PLUS',
    'DIV',
    'MULT',
)

t_COMMA = r','
t_LESS_GREAT = r'\<\>'
t_LESS_THAN = r'\<'
t_GREATER_THAN = r'\>'
t_ASSIGN = r'\='
t_LPAREN = r'\('
t_RPAREN = r'\)'
t_LBRACE = r'\{'
t_RBRACE = r'\}'
t_MINUS = r'-'
t_PLUS = r'\+'
t_DIV = r'\/'
t_MULT = r'\*'
t_COLON = r':'
t_SEMICOLON = r';'

RESERVED = {
    'int': "INT_TYPE",
    "float": "FLOAT_TYPE",
    "if": "IF",
    "else": "ELSE",
    "var": "VAR",
    "program": "PROGRAM",
    "print": "PRINT",
}

precedence = (
    ('left', '+', '-'),
    ('left', '*', '/'),
    ('right', 'UMINUS'),
)

# ciri is opinionated, we don't like _ variables
def t_ID(t):
    r"[a-zA-Z][a-zA-Z0-9]*"
    if t.value in RESERVED:
        t.type = RESERVED[t.value]
    else:
        t.type = 'ID'
    return t

t_ignore = ' '

def t_CTE_STRING(t):
    r'\"[a-zA-Z_][a-zA-Z0-9_]*\"' # I think this is right ...
    return t

def t_NEWLINE(t):
    r'\n'
    pass

def t_FLOAT(t):
    r"[+-]?([0-9]*[.])[0-9]+"
    t.type = 'FLOAT'
    return t


def t_INT(t):
    r"[-+]?[0-9]+"
    t.type = 'INT'
    return t


def t_error(t):
    print("invalid character: ", t.value[0])
    print("Skipping")
    t.lexer.skip(1)


lexer = lex.lex()
