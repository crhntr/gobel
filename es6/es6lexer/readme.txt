A.1 Lexical Grammar


(See 10.1)
SourceCharacter ::
any Unicode code point



(See clause 11)
InputElementDiv ::
WhiteSpace
LineTerminator
Comment
CommonToken
DivPunctuator
RightBracePunctuator



(See clause 11)
InputElementRegExp ::
WhiteSpace
LineTerminator
Comment
CommonToken
RightBracePunctuator
RegularExpressionLiteral



(See clause 11)
InputElementRegExpOrTemplateTail ::
WhiteSpace
LineTerminator
Comment
CommonToken
RegularExpressionLiteral
TemplateSubstitutionTail



(See clause 11)
InputElementTemplateTail ::
WhiteSpace
LineTerminator
Comment
CommonToken
DivPunctuator
TemplateSubstitutionTail



(See 11.2)
WhiteSpace ::
<TAB>
<VT>
<FF>
<SP>
<NBSP>
<ZWNBSP>
<USP>



(See 11.3)
LineTerminator ::
<LF>
<CR>
<LS>
<PS>



(See 11.3)
LineTerminatorSequence ::
<LF>
<CR> [lookahead ≠ <LF> ]
<LS>
<PS>
<CR> <LF>



(See 11.4)
Comment ::
MultiLineComment
SingleLineComment



(See 11.4)
MultiLineComment ::
/* MultiLineCommentCharsopt */



(See 11.4)
MultiLineCommentChars ::
MultiLineNotAsteriskChar MultiLineCommentCharsopt
* PostAsteriskCommentCharsopt



(See 11.4)
PostAsteriskCommentChars ::
MultiLineNotForwardSlashOrAsteriskChar MultiLineCommentCharsopt
* PostAsteriskCommentCharsopt



(See 11.4)
MultiLineNotAsteriskChar ::
SourceCharacter but not *



(See 11.4)
MultiLineNotForwardSlashOrAsteriskChar ::
SourceCharacter but not one of / or *



(See 11.4)
SingleLineComment ::
// SingleLineCommentCharsopt



(See 11.4)
SingleLineCommentChars ::
SingleLineCommentChar SingleLineCommentCharsopt



(See 11.4)
SingleLineCommentChar ::
SourceCharacter but not LineTerminator



(See 11.5)
CommonToken ::
IdentifierName
Punctuator
NumericLiteral
StringLiteral
Template



(See 11.6)
IdentifierName ::
IdentifierStart
IdentifierName IdentifierPart



(See 11.6)
IdentifierStart ::
UnicodeIDStart
$
_
\ UnicodeEscapeSequence



(See 11.6)
IdentifierPart ::
UnicodeIDContinue
$
_
\ UnicodeEscapeSequence
<ZWNJ>
<ZWJ>



(See 11.6)
UnicodeIDStart ::
any Unicode code point with the Unicode property “ID_Start” or “Other_ID_Start”



(See 11.6)
UnicodeIDContinue ::
any Unicode code point with the Unicode property “ID_Continue”, “Other_ID_Continue”, or “Other_ID_Start”



(See 11.6.2)
ReservedWord ::
Keyword
FutureReservedWord
NullLiteral
BooleanLiteral



(See 11.6.2.1)
Keyword :: one of
break	do	in	typeof
case	else	instanceof	var
catch	export	new	void
class	extends	return	while
const	finally	super	with
continue	for	switch	yield
debugger	function	this
default	if	throw
delete	import	try



(See 11.6.2.2)
FutureReservedWord ::
enum
await
await is only treated as a FutureReservedWord when Module is the goal symbol of the syntactic grammar.

The following tokens are also considered to be FutureReservedWords when parsing strict mode code (see 10.2.1).

implements	package	protected
interface	private	public



(See 11.7)
Punctuator :: one of
{	}	(	)	[	]
.	;	,	<	>	<=
>=	==	!=	===	!==
+	-	*	%	++	--
<<	>>	>>>	&	|	^
!	~	&&	||	?	:
=	+=	-=	*=	%=	<<=
>>=	>>>=	&=	|=	^=	=>



(See 11.7)
DivPunctuator :: one of
/
/=



(See 11.7)
RightBracePunctuator :: one of
}



(See 11.8.1)
NullLiteral ::
null



(See 11.8.2)
BooleanLiteral ::
true
false



(See 11.8.3)
NumericLiteral ::
DecimalLiteral
BinaryIntegerLiteral
OctalIntegerLiteral
HexIntegerLiteral



(See 11.8.3)
DecimalLiteral ::
DecimalIntegerLiteral . DecimalDigitsopt ExponentPartopt
. DecimalDigits ExponentPartopt
DecimalIntegerLiteral ExponentPartopt



(See 11.8.3)
DecimalIntegerLiteral ::
0
NonZeroDigit DecimalDigitsopt



(See 11.8.3)
DecimalDigits ::
DecimalDigit
DecimalDigits DecimalDigit



(See 11.8.3)
DecimalDigit :: one of
0 1 2 3 4 5 6 7 8 9



(See 11.8.3)
NonZeroDigit :: one of
1 2 3 4 5 6 7 8 9



(See 11.8.3)
ExponentPart ::
ExponentIndicator SignedInteger



(See 11.8.3)
ExponentIndicator :: one of
e E



(See 11.8.3)
SignedInteger ::
DecimalDigits
+ DecimalDigits
- DecimalDigits



(See 11.8.3)
BinaryIntegerLiteral ::
0b BinaryDigits
0B BinaryDigits



(See 11.8.3)
BinaryDigits ::
BinaryDigit
BinaryDigits BinaryDigit



(See 11.8.3)
BinaryDigit :: one of
0 1



(See 11.8.3)
OctalIntegerLiteral ::
0o OctalDigits
0O OctalDigits



(See 11.8.3)
OctalDigits ::
OctalDigit
OctalDigits OctalDigit



(See 11.8.3)
OctalDigit :: one of
0 1 2 3 4 5 6 7



(See 11.8.3)
HexIntegerLiteral ::
0x HexDigits
0X HexDigits



(See 11.8.3)
HexDigits ::
HexDigit
HexDigits HexDigit



(See 11.8.3)
HexDigit :: one of
0 1 2 3 4 5 6 7 8 9 a b c d e f A B C D E F



(See 11.8.4)
StringLiteral ::
" DoubleStringCharactersopt "
' SingleStringCharactersopt '



(See 11.8.4)
DoubleStringCharacters ::
DoubleStringCharacter DoubleStringCharactersopt



(See 11.8.4)
SingleStringCharacters ::
SingleStringCharacter SingleStringCharactersopt



(See 11.8.4)
DoubleStringCharacter ::
SourceCharacter but not one of " or \ or LineTerminator
\ EscapeSequence
LineContinuation



(See 11.8.4)
SingleStringCharacter ::
SourceCharacter but not one of ' or \ or LineTerminator
\ EscapeSequence
LineContinuation



(See 11.8.4)
LineContinuation ::
\ LineTerminatorSequence



(See 11.8.4)
EscapeSequence ::
CharacterEscapeSequence
0 [lookahead ∉ DecimalDigit]
HexEscapeSequence
UnicodeEscapeSequence



(See 11.8.4)
CharacterEscapeSequence ::
SingleEscapeCharacter
NonEscapeCharacter



(See 11.8.4)
SingleEscapeCharacter :: one of
' " \ b f n r t v



(See 11.8.4)
NonEscapeCharacter ::
SourceCharacter but not one of EscapeCharacter or LineTerminator



(See 11.8.4)
EscapeCharacter ::
SingleEscapeCharacter
DecimalDigit
x
u



(See 11.8.4)
HexEscapeSequence ::
x HexDigit HexDigit



(See 11.8.4)
UnicodeEscapeSequence ::
u Hex4Digits
u{ HexDigits }



(See 11.8.4)
Hex4Digits ::
HexDigit HexDigit HexDigit HexDigit



(See 11.8.5)
RegularExpressionLiteral ::
/ RegularExpressionBody / RegularExpressionFlags



(See 11.8.5)
RegularExpressionBody ::
RegularExpressionFirstChar RegularExpressionChars



(See 11.8.5)
RegularExpressionChars ::
[empty]
RegularExpressionChars RegularExpressionChar



(See 11.8.5)
RegularExpressionFirstChar ::
RegularExpressionNonTerminator but not one of * or \ or / or [
RegularExpressionBackslashSequence
RegularExpressionClass



(See 11.8.5)
RegularExpressionChar ::
RegularExpressionNonTerminator but not one of \ or / or [
RegularExpressionBackslashSequence
RegularExpressionClass



(See 11.8.5)
RegularExpressionBackslashSequence ::
\ RegularExpressionNonTerminator



(See 11.8.5)
RegularExpressionNonTerminator ::
SourceCharacter but not LineTerminator



(See 11.8.5)
RegularExpressionClass ::
[ RegularExpressionClassChars ]



(See 11.8.5)
RegularExpressionClassChars ::
[empty]
RegularExpressionClassChars RegularExpressionClassChar



(See 11.8.5)
RegularExpressionClassChar ::
RegularExpressionNonTerminator but not one of ] or \
RegularExpressionBackslashSequence



(See 11.8.5)
RegularExpressionFlags ::
[empty]
RegularExpressionFlags IdentifierPart



(See 11.8.6)
Template ::
NoSubstitutionTemplate
TemplateHead



(See 11.8.6)
NoSubstitutionTemplate ::
` TemplateCharactersopt `



(See 11.8.6)
TemplateHead ::
` TemplateCharactersopt ${



(See 11.8.6)
TemplateSubstitutionTail ::
TemplateMiddle
TemplateTail



(See 11.8.6)
TemplateMiddle ::
} TemplateCharactersopt ${



(See 11.8.6)
TemplateTail ::
} TemplateCharactersopt `



(See 11.8.6)
TemplateCharacters ::
TemplateCharacter TemplateCharactersopt



(See 11.8.6)
TemplateCharacter ::
$ [lookahead ≠ { ]
\ EscapeSequence
LineContinuation
LineTerminatorSequence
SourceCharacter but not one of ` or \ or $ or LineTerminator
