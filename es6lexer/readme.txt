# ES6 Compliant Lexer

## Design
The design of this lexer is based on Rob Pike's talk "Lexical Scanning in Go"
from Google Technology User Group given on Tuesday, 30 August 2011 it can be
watched at: https://www.youtube.com/watch?v=HxaD_trXwRE

## Copy of the lexical grammar from:
http://www.ecma-international.org/ecma-262/6.0/#sec-lexical-grammar

A.1 Lexical Grammar

SourceCharacter (See 10.1) ::
  any Unicode code point


InputElementDiv (See clause 11) ::
  WhiteSpace
  LineTerminator
  Comment
  CommonToken
  DivPunctuator
  RightBracePunctuator


InputElementRegExp (See clause 11) ::
  WhiteSpace
  LineTerminator
  Comment
  CommonToken
  RightBracePunctuator
  RegularExpressionLiteral


InputElementRegExpOrTemplateTail (See clause 11) ::
  WhiteSpace
  LineTerminator
  Comment
  CommonToken
  RegularExpressionLiteral
  TemplateSubstitutionTail


InputElementTemplateTail (See clause 11) ::
  WhiteSpace
  LineTerminator
  Comment
  CommonToken
  DivPunctuator
  TemplateSubstitutionTail


WhiteSpace (See 11.2) ::
  <TAB>
  <VT>
  <FF>
  <SP>
  <NBSP>
  <ZWNBSP>
  <USP>


LineTerminator (See 11.3) ::
  <LF>
  <CR>
  <LS>
  <PS>


LineTerminatorSequence (See 11.3) ::
  <LF>
  <CR> [lookahead ≠ <LF> ]
  <LS>
  <PS>
  <CR> <LF>


Comment (See 11.4) ::
  MultiLineComment
  SingleLineComment


MultiLineComment (See 11.4) ::
  /* MultiLineCommentCharsopt */


MultiLineCommentChars (See 11.4) ::
  MultiLineNotAsteriskChar MultiLineCommentCharsopt
  * PostAsteriskCommentCharsopt


PostAsteriskCommentChars (See 11.4) ::
  MultiLineNotForwardSlashOrAsteriskChar MultiLineCommentCharsopt
  * PostAsteriskCommentCharsopt


MultiLineNotAsteriskChar (See 11.4) ::
  SourceCharacter but not *


MultiLineNotForwardSlashOrAsteriskChar (See 11.4) ::
  SourceCharacter but not one of / or *


SingleLineComment (See 11.4) ::
  // SingleLineCommentCharsopt


SingleLineCommentChars (See 11.4) ::
  SingleLineCommentChar SingleLineCommentCharsopt


SingleLineCommentChar (See 11.4) ::
  SourceCharacter but not LineTerminator


CommonToken (See 11.5) ::
  IdentifierName
  Punctuator
  NumericLiteral
  StringLiteral
  Template


IdentifierName (See 11.6) ::
  IdentifierStart
  IdentifierName IdentifierPart


IdentifierStart (See 11.6) ::
  UnicodeIDStart
  $
  _
  \ UnicodeEscapeSequence


IdentifierPart (See 11.6) ::
  UnicodeIDContinue
  $
  _
  \ UnicodeEscapeSequence
  <ZWNJ>
  <ZWJ>


UnicodeIDStart (See 11.6) ::
  any Unicode code point with the Unicode property “ID_Start” or “Other_ID_Start”


UnicodeIDContinue (See 11.6) ::
  any Unicode code point with the Unicode property “ID_Continue”, “Other_ID_Continue”, or “Other_ID_Start”


ReservedWord (See 11.6.2) ::
  Keyword
  FutureReservedWord
  NullLiteral
  BooleanLiteral


Keyword (See 11.6.2.1) ::  one of
  break	do	in	typeof
  case	else	instanceof	var
  catch	export	new	void
  class	extends	return	while
  const	finally	super	with
  continue	for	switch	yield
  debugger	function	this
  default	if	throw
  delete	import	try


FutureReservedWord (See 11.6.2.2) ::
  enum
  await
  await is only treated as a FutureReservedWord when Module is the goal symbol of the syntactic grammar.

The following tokens are also considered to be FutureReservedWords when parsing strict mode code (see 10.2.1).

implements	package	protected
interface	private	public


Punctuator (See 11.7) ::   one of
  {	}	(	)	[	]
  .	;	,	<	>	<=
  >=	==	!=	===	!==
  +	-	*	%	++	--
  <<	>>	>>>	&	|	^
  !	~	&&	||	?	:
  =	+=	-=	*=	%=	<<=
  >>=	>>>=	&=	|=	^=	=>


DivPunctuator (See 11.7) ::  one of
  /
  /=


RightBracePunctuator (See 11.7) ::   one of
  }


NullLiteral (See 11.8.1) ::
  null


BooleanLiteral (See 11.8.2) ::
  true
  false


NumericLiteral (See 11.8.3) ::
  DecimalLiteral
  BinaryIntegerLiteral
  OctalIntegerLiteral
  HexIntegerLiteral


DecimalLiteral (See 11.8.3) ::
  DecimalIntegerLiteral . DecimalDigitsopt ExponentPartopt
  . DecimalDigits ExponentPartopt
  DecimalIntegerLiteral ExponentPartopt


DecimalIntegerLiteral (See 11.8.3) ::
  0
  NonZeroDigit DecimalDigitsopt


DecimalDigits (See 11.8.3) ::
  DecimalDigit
  DecimalDigits DecimalDigit


DecimalDigit (See 11.8.3) ::   one of
  0 1 2 3 4 5 6 7 8 9


NonZeroDigit (See 11.8.3) ::   one of
  1 2 3 4 5 6 7 8 9


ExponentPart (See 11.8.3) ::
  ExponentIndicator SignedInteger


ExponentIndicator (See 11.8.3) ::  one of
  e E


SignedInteger (See 11.8.3) ::
  DecimalDigits
  + DecimalDigits
  - DecimalDigits


BinaryIntegerLiteral (See 11.8.3) ::
  0b BinaryDigits
  0B BinaryDigits


BinaryDigits (See 11.8.3) ::
  BinaryDigit
  BinaryDigits BinaryDigit


BinaryDigit (See 11.8.3) ::  one of
  0 1


OctalIntegerLiteral (See 11.8.3) ::
  0o OctalDigits
  0O OctalDigits


OctalDigits (See 11.8.3) ::
  OctalDigit
  OctalDigits OctalDigit


OctalDigit (See 11.8.3) ::   one of
  0 1 2 3 4 5 6 7


HexIntegerLiteral (See 11.8.3) ::
  0x HexDigits
  0X HexDigits


HexDigits (See 11.8.3) ::
  HexDigit
  HexDigits HexDigit


HexDigit (See 11.8.3) ::   one of
  0 1 2 3 4 5 6 7 8 9 a b c d e f A B C D E F


StringLiteral (See 11.8.4) ::
  " DoubleStringCharactersopt "
  ' SingleStringCharactersopt '


DoubleStringCharacters (See 11.8.4) ::
  DoubleStringCharacter DoubleStringCharactersopt


SingleStringCharacters (See 11.8.4) ::
  SingleStringCharacter SingleStringCharactersopt


DoubleStringCharacter (See 11.8.4) ::
  SourceCharacter but not one of " or \ or LineTerminator
  \ EscapeSequence
  LineContinuation


SingleStringCharacter (See 11.8.4) ::
  SourceCharacter but not one of ' or \ or LineTerminator
  \ EscapeSequence
  LineContinuation


LineContinuation (See 11.8.4) ::
  \ LineTerminatorSequence


EscapeSequence (See 11.8.4) ::
  CharacterEscapeSequence
  0 [lookahead ∉ DecimalDigit]
  HexEscapeSequence
  UnicodeEscapeSequence


CharacterEscapeSequence (See 11.8.4) ::
  SingleEscapeCharacter
  NonEscapeCharacter


SingleEscapeCharacter (See 11.8.4) ::  one of
  ' " \ b f n r t v


NonEscapeCharacter (See 11.8.4) ::
  SourceCharacter but not one of EscapeCharacter or LineTerminator


EscapeCharacter (See 11.8.4) ::
  SingleEscapeCharacter
  DecimalDigit
  x
  u


HexEscapeSequence (See 11.8.4) ::
  x HexDigit HexDigit


UnicodeEscapeSequence (See 11.8.4) ::
  u Hex4Digits
  u{ HexDigits }


Hex4Digits (See 11.8.4) ::
  HexDigit HexDigit HexDigit HexDigit


RegularExpressionLiteral (See 11.8.5) ::
  / RegularExpressionBody / RegularExpressionFlags


RegularExpressionBody (See 11.8.5) ::
  RegularExpressionFirstChar RegularExpressionChars


RegularExpressionChars (See 11.8.5) ::
  [empty]
  RegularExpressionChars RegularExpressionChar


RegularExpressionFirstChar (See 11.8.5) ::
  RegularExpressionNonTerminator but not one of * or \ or / or [
  RegularExpressionBackslashSequence
  RegularExpressionClass


RegularExpressionChar (See 11.8.5) ::
  RegularExpressionNonTerminator but not one of \ or / or [
  RegularExpressionBackslashSequence
  RegularExpressionClass


RegularExpressionBackslashSequence (See 11.8.5) ::
  \ RegularExpressionNonTerminator


RegularExpressionNonTerminator (See 11.8.5) ::
  SourceCharacter but not LineTerminator


RegularExpressionClass (See 11.8.5) ::
  [ RegularExpressionClassChars ]


RegularExpressionClassChars (See 11.8.5) ::
  [empty]
  RegularExpressionClassChars RegularExpressionClassChar


RegularExpressionClassChar (See 11.8.5) ::
  RegularExpressionNonTerminator but not one of ] or \
  RegularExpressionBackslashSequence


RegularExpressionFlags (See 11.8.5) ::
  [empty]
  RegularExpressionFlags IdentifierPart


Template (See 11.8.6) ::
  NoSubstitutionTemplate
  TemplateHead


NoSubstitutionTemplate (See 11.8.6) ::
  ` TemplateCharactersopt `


TemplateHead (See 11.8.6) ::
  ` TemplateCharactersopt ${


TemplateSubstitutionTail (See 11.8.6) ::
  TemplateMiddle
  TemplateTail


TemplateMiddle (See 11.8.6) ::
  } TemplateCharactersopt ${


TemplateTail (See 11.8.6) ::
  } TemplateCharactersopt `


TemplateCharacters (See 11.8.6) ::
  TemplateCharacter TemplateCharactersopt


TemplateCharacter (See 11.8.6) ::
  $ [lookahead ≠ { ]
  \ EscapeSequence
  LineContinuation
  LineTerminatorSequence
  SourceCharacter but not one of ` or \ or $ or LineTerminator
