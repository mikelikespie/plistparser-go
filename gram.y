%{
package plistparser

import (
  "strconv"
  "encoding/hex"
)
%}

%union{
    tok int
    val interface{}
    str string
    data []byte
    array []interface{}
    pair struct {key string; val interface{}}
    pairs map[string]interface{}
}

%token STRING
%token IDENT

%left '{', '}', '(', ')'
%left '<', '>'
%left '='
%left ';'
%left ','

%type <pair> Pair
%type <pairs> Pairs, Dict

%type <str> STRING, IDENT, String

%type <array> Array ArrayElements 
%type <val> Value ArrayElement
%type <data> Data DataChunks, DataChunk

%%

Start: Value {
        yylex.(*lexer).result = $1
}
;
   
Value: String {
        $$ = $1
      }
      | Array {
        $$ = $1
      }
      | Dict {
        $$ = $1
      }
      | Data {
        $$ = $1
      }
  ;

String: STRING {
  var e error
  if $$, e = strconv.Unquote($1); e != nil {
    yylex.Error("invalid unquoting" + e.Error())
  }
}
| IDENT {
  $$, _ = strconv.Unquote("\"" + $1 + "\"")
}

Array: '(' ')' {
        $$ = []interface{} {}
     }
     | '(' ArrayElements ')'  {
        $$ = $2
     }
  ;

ArrayElements: ArrayElement {
               $$ = []interface{} {$1}
             }
             | ArrayElements ArrayElement{
               $$ = append($1, $2)
             }
  ;

ArrayElement: Value ',' {
            $$ = $1
            }
  ;

Dict: '{' '}' {
      $$ = map[string]interface{} {}
    }
    | '{' Pairs '}' {
      $$ = $2
    }
  ;

Pairs: Pair {
       $$ = map[string]interface{}{$1.key: $1.val}
     }
     | Pairs Pair {
        $$[$2.key] = $2.val
     }
  ;

Pair: String '=' Value ';' {
    $$.key, $$.val = $1, $3
    }
  ;

Data: '<' DataChunks '>' {
        $$ = $2
    }

DataChunk: IDENT {
              var e error

              if len($1) != 4 {
                yylex.Error("Chunks inside < > brackets must be lenght of 4")
              }

              if $$, e = hex.DecodeString($1); e != nil {
                yylex.Error("error decoding hex chunk" + e.Error())
              }
           }
DataChunks: DataChunk {
              $$ = $1
            }
            |
            DataChunks DataChunk {
              $$ = append($1, $2...)
            }
            ;

%%

