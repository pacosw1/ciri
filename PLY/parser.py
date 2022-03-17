#!/usr/bin/python
# -*- coding: utf-8 -*-
import ply.yacc as yacc
from lex import lexer, tokens



def p_program(p):
    ''' program : PROGRAM ID COLON vars bloque'''

def p_vars(p):
    ''' vars : VAR allVars
        |'''


def p_allVars(p):
    '''allVars : nextId COLON tipo SEMICOLON nextVar'''


def p_nextId(p):
    '''nextId :  ID
        | ID COMMA nextId'''


def p_nextVar(p):
    '''nextVar : allVars
       |'''


def p_bloque(p):
    ''' bloque : LBRACE nextStatuto RBRACE '''


def p_nextStatuto(p):
    ''' nextStatuto : estatuto nextStatuto
        |'''


def p_estatuto(p):
    ''' estatuto : assign
       | condition
       | print'''


def p_condition(p):
    ''' condition : IF LPAREN expresion RPAREN bloque elseBlock SEMICOLON '''


def p_elseBlock(p):
    ''' elseBlock : ELSE bloque
       |'''


def p_assign(p):
    ''' assign : ID ASSIGN expresion SEMICOLON '''


def p_print(p):
    ''' print : LPAREN PRINT nextPrintExp nextPrint RPAREN SEMICOLON '''


def p_nextPrintExp(p):
    ''' nextPrintExp : expresion
       | CTE_STRING'''


def p_nextPrint(p):
    ''' nextPrint : COMMA nextPrintExp nextPrint
        |'''


def p_tipo(p):
    ''' tipo : INT_TYPE
       | FLOAT_TYPE'''


def p_varCte(p):
    '''varCte : ID
       | INT
       | FLOAT'''


def p_factor(p):
    '''factor : LPAREN expresion RPAREN
        | cteExp'''


def p_cteExp(p):
    ''' cteExp : varCte
       | PLUS varCte
       | MINUS varCte'''


def p_termino(p):
    ''' termino : nextFactor'''


def p_nextFactor(p):
    ''' nextFactor : factor
       | factor DIV termino
       | factor MULT termino'''


def p_exp(p):
    '''exp : termino nextTerm'''


def p_nextTerm(p):
    ''' nextTerm : PLUS termino nextTerm
        | MINUS termino nextTerm
        |'''

def p_error(p):
    print("error" + p.value)

def p_expresion(p):
    '''expresion : nextExp'''


def p_nextExp(p):
    '''nextExp : exp GREATER_THAN exp
       | exp LESS_THAN exp
       | exp LESS_GREAT exp
       | exp'''


parser = yacc.yacc(debug=False)
