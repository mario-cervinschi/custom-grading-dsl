const PREC = {
  ASSIGN: 1,
  OR: 2,
  AND: 3,
  RELATIONAL: 4,
  ADD: 5,
  MUL: 6,
  UNARY: 7,
  CALL: 8,
};

module.exports = grammar({
  name: 'mylang',

  extras: $ => [/\s/, /\r?\n/],

  rules: {
    source_file: $ => repeat(choice(
      $.apply,
      $.normal_operation,
      $.when_operation
    )),

    // ---------- keywords ----------
    APPLY: _ => caseInsensitive('APPLY'),
    TO: _ => caseInsensitive('TO'),
    WHEN: _ => caseInsensitive('WHEN'),
    EXISTS: _ => caseInsensitive('EXISTS'),
    EXPLAIN: _ => caseInsensitive('EXPLAIN'),
    IN: _ => caseInsensitive('IN'),

    // ---------- top level ----------
    apply: $ => seq(
      $.APPLY,
      $.lambda,
      $.TO,
      $.list,
      ';'
    ),

    normal_operation: $ => seq(
      optional($.EXPLAIN),
      $.expression,
      optional(seq(',', $.string)),
      ';'
    ),

    when_operation: $ => seq(
      $.WHEN,
      $.list,
      $.EXISTS,
      $.normal_operation
    ),

    // ---------- structures ----------
    list: $ => seq(
      $.identifier,
      repeat(seq(',', $.identifier))
    ),

    lambda: $ => seq(
      $.identifier,
      '=>',
      $.expression
    ),

    // ---------- expressions ----------
    expression: $ => choice(
      $.assignment_expression,
      $.conditional_expression,
      $.simple_expression
    ),

    simple_expression: $ => choice(
      $.iterate_expression,
      $.binary_expression,
      $.unary_expression,
      $.primary_expression
    ),

    assignment_expression: $ => prec.right(PREC.ASSIGN, seq(
      $.identifier,
      '=',
      $.expression
    )),

    iterate_expression: $ => prec.right(seq(
      $.identifier,
      $.IN,
      '[',
      $.constant_list,
      ']',
      optional(seq('?', $.expression, ':', $.expression))
    )),

    conditional_expression: $ => prec.right(seq(
      $.binary_expression,
      '?',
      $.expression,
      ':',
      $.expression
    )),

    binary_expression: $ => choice(
      prec.left(PREC.OR, seq($.simple_expression, '||', $.simple_expression)),
      prec.left(PREC.AND, seq($.simple_expression, '&&', $.simple_expression)),
      prec.left(PREC.RELATIONAL, seq($.simple_expression, '==', $.simple_expression)),
      prec.left(PREC.RELATIONAL, seq($.simple_expression, '!=', $.simple_expression)),
      prec.left(PREC.RELATIONAL, seq($.simple_expression, '>', $.simple_expression)),
      prec.left(PREC.RELATIONAL, seq($.simple_expression, '>=', $.simple_expression)),
      prec.left(PREC.RELATIONAL, seq($.simple_expression, '<', $.simple_expression)),
      prec.left(PREC.RELATIONAL, seq($.simple_expression, '<=', $.simple_expression)),
      prec.left(PREC.ADD, seq($.simple_expression, '+', $.simple_expression)),
      prec.left(PREC.ADD, seq($.simple_expression, '-', $.simple_expression)),
      prec.left(PREC.MUL, seq($.simple_expression, '*', $.simple_expression)),
      prec.left(PREC.MUL, seq($.simple_expression, '/', $.simple_expression))
    ),

    unary_expression: $ => prec(PREC.UNARY, choice(
      seq('!', $.simple_expression),
      seq('-', $.simple_expression)
    )),

    primary_expression: $ => choice(
      $.constant,
      $.identifier,
      $.number,
      $.string,
      $.regex_expression,
      seq('(', $.expression, ')'),
      $.function_call
    ),

    regex_expression: $ => seq(
      $.identifier,
      choice('~', '!~'),
      $.string
    ),

    function_call: $ => prec(PREC.CALL, seq(
      $.function_name,
      '(',
      $.expression,
      repeat(seq(',', $.expression)),
      ')'
    )),

    function_name: $ => choice(
      caseInsensitive('MAX'),
      caseInsensitive('MIN'),
      caseInsensitive('ROUND')
    ),

    constant_list: $ => seq($.constant, repeat(seq(',', $.constant))),

    constant: $ => choice(
      'nothing','fraud','cancelled','invalid','alert',
      'conflict','ungraded','obscured','absent','present',
      'excused','toolow'
    ),

    identifier: _ => /[a-zA-Z][a-zA-Z0-9_]*/,
    number: _ => /\d+(\.\d+)?/,
    string: _ => /"([^"\r\n]|"")*"/,
  }
});

function caseInsensitive(word) {
  return new RegExp(
    word.split('').map(c => `[${c}${c.toLowerCase()}]`).join('')
  );
}
