GOOF----LE-4-2.0b!      ] W 4       h      ] g  guile¤	 ¤	g  define-module*¤	 ¤	 ¤	g  language¤	g  elisp¤	g  parser¤		 ¤	
g  filenameS¤	f  ]/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/elisp/parser.scm¤	g  importsS¤	g  lexer¤	 ¤	 ¤	 ¤	g  exportsS¤	g  
read-elisp¤	 ¤	g  set-current-module¤	 ¤	 ¤	g  error¤	g  parse-error¤	g  
make-fluid¤	g  circular-definitions¤	g  make-hash-table¤	g  make-circular-definitions¤	g  circular-ref¤	f  invalid token for circular-ref¤	g  	hashq-ref¤	 f  undefined circular reference¤	!g  circular-def¤	"f  "invalid token for circular-define!¤	#g  
hashq-set!¤	$g  make-promise¤	%$ ¤	&$ ¤	'g  circular-define!¤	(g  promise?¤	)g  force¤	*g  force-promises!¤	+g  vector-length¤	,g  finish¤	-f  'lexer-buffer is not empty when finished¤	.g  peek¤	/g  get¤	0f  invalid lexer-buffer action¤	1g  make-lexer-buffer¤	2g  square-close¤	3g  paren-close¤	4f  got different token than peeked¤	5g  get-expression¤	6g  get-list¤	7g  dot¤	8g  length¤	9f  &expected exactly one element after dot¤	:g  quote¤	;::¤	<g  	backquote¤	=g  `¤	><=¤	?g  unquote¤	@g  ,¤	A?@¤	Bg  unquote-splicing¤	Cg  ,@¤	DBC¤	E;>AD ¤	Fg  quotation-symbols¤	Gg  eof¤	Hf  end of file during parsing¤	Ig  integer¤	Jg  float¤	Kg  symbol¤	Lg  	character¤	Mg  string¤	Ng  set-source-properties!¤	Og  source-properties¤	Pg  function¤	Qg  assq-ref¤	Rg  
paren-open¤	Sg  square-open¤	Tg  list->vector¤	Uf  expected expression, got¤	Vg  	get-lexer¤C 5   h@    ]4	
5 4 >  "  G         h   é   - 1 3 @       á       g  token
			 g  msg			 g  args				  g  filenamef  ]/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/elisp/parser.scm
	"
		#	 				
	  g  nameg  parse-error CR4i5 R    h   ½   ] 6   µ       g  filenamef  ]/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/elisp/parser.scm
	3
		4	 		
  g  nameg  make-circular-definitions CR       hH   F  ] &  "  4 >  "  G   4[5$  C 6 >      g  token
		G g  id	&	G g  value		2	G  g  filenamef  ]/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/elisp/parser.scm
	6
		7			7			7			8			8			8		&	9		&	9		)	:		2	9		:	;		C	=		G	=	 		G  g  nameg  circular-ref CR!"#&h      ] M C          g  filenamef  ]/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/elisp/parser.scm
	I	 		
   C#  h   ¹   ] NL L 6±       g  
real-value
		  g  filenamef  ]/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/elisp/parser.scm
	J			K			L	 		   C 	   h`   M  ] &  "  4 >  "  G  [ H44O 5>  "  G  O C    E      g  token
		\ g  value	*	\ g  table		*	\ g  id		*	\  g  filenamef  ]/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/elisp/parser.scm
	C
		D			D			D			E			E			E		*	H		*	F		1	I		8	I		H	I	 		\  g  nameg  circular-define! C'R()*+    hÀ   9  ] $  N4 5$   4 5"  4 >  "  G  4 5$   4 5C 6 $  \4 5"  H$  = £45$   45¤"  4>  "  G  "ÿÿºC
"ÿÿ±C    1      g  data
	 ¼ g  len	e º g  i		k ³ g  el		y ±  g  filenamef  ]/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/elisp/parser.scm
	U
		W				V		
	Y	
		Y			Y	
		Y			Z			Z	 		Z		 	Z	
	%	[	
	*	[		/	[	
	8	\	
	=	\		?	\	
	C	\		F	]		K	]	 	M	]		N	]	
	U	^		W	^	
	Z	_		^	V		_	`		e	`		k	a		p	b		t	b		y	c		y	c		|	d	 	d	 	e	& 	e	 	f	 «	g	 ±	g	 ³	a	 *	 ¼  g  nameg  force-promises! C*R,-./0   hP   #  ]	 &  M$  6CM$  "  4L 5 N $  MC $  
MNC 6      g  action
		P g  result	@	H  g  filenamef  ]/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/elisp/parser.scm
	t			u		
	u			v	
		w			w			z		#	{	!	)	{		2	|		@			E 		L 		P 	 		P   C  h   Þ   ]	H O C Ö       g  lex
		 g  
look-ahead		  g  filenamef  ]/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/elisp/parser.scm
	r
		s	 		  g  nameg  make-lexer-buffer C1R.23/456789   hà   ä  ]#4 5$  "  &  &4 5&  "  4>  "  G  C"  4 54 5C$  e	&  Y4 5&  "  4>  "  G  4 54
5$  "  4>  "  G  C"ÿÿy"ÿÿu       Ü      g  lex
	 Ù g  	allow-dot	 Ù g  close-square		 Ù g  next			 Ù g  type		 Ù g  head		U	j g  tail		b	j g  tail	 ¦ Ñ  g  filenamef  ]/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/elisp/parser.scm
 
	 		 			 			 		 		 		 		 	!	  	/	$ 		% 		) 	 	+ 		, 		2 		7 		; 		@ 		I 		O 		U 		X  		b 		i ¡		j 		t 		x 		y 		} 	 	 	  	  	  	  	  	  	 ¦ 	 © 	 ± 	 µ 	
 º 	 À 	 Å 	 Ð 	
 .	 Ù	  g  nameg  get-list C6REFR/GHIJKLMNOP5:<?BQFR6ST!'*U      h  ÿ  ]!4 5$  6$  "  /$  "  !$  "  $  "  	$  ,$  4
45>  "  G  "   C$  54 5 $  4
45>  "  G  "   C$  "  !$  "  $  "  $  <454 5 $  4
45>  "  G  "   C$  24 5$  4
45>  "  G  "   C$  744 55$  4
45>  "  G  "   C$  6$  9454 54>  "  G  4>  "  G  C6       ÷      g  lex
	 g  token		 g  type		 g  result		e  g  result	 £ Ì g  result	; g  result	Mv g  result	¶ g  setter	Õ g  expr	Þ  
g  filenamef  ]/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/elisp/parser.scm
 «
	 ¬		 ¬			 ¬			 ¬		 ­		 ¬		 ´		 ¶		! ¶		* ´		e ¸		e ¸		j ¯		n ¯		o °		t ²		 °	  ´	  º	  º	 £ º	 £ º	 ¨ ¯	 ¬ ¯	 ­ °	 ² ²	 ½ °	 Õ ´	  ¼		 ½	 ¼	 ¼	 ¯	 ¯	 °	! ²	, °	D ´	E ¿	M ¿	R ¯	V ¯	W °	\ ²	g °	 ´	 Á	 Á	 Á	 Á	 ¯	 ¯	 °	 ²	§ °	¿ ´	Å Ã	Î ´	Ï Æ	Õ Æ	Ø Ç	Þ Æ	á È		ó É		 Ì	 Ì	 D	  g  nameg  get-expression C5RV1.G5, 	 hX   ¨  ]!45 Y4 54545&  "  454>  "  G  ZCZF        g  port
		W g  lexer		S g  lexbuf			S g  next		$	S g  result		=	S  g  filenamef  ]/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/elisp/parser.scm
 Ò
	 Ó	&	 Ô		 Ô		 Õ		 Ô		 Ö		" Ö		$ Ö		$ Ô		) ×		+ ×		/ ×		2 Ø	
	7 Ù		= Ù	
	@ Ú		D Ú		I Ú	 		W  g  nameg  
read-elisp CRC         g  m
		,  g  filenamef  ]/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/elisp/parser.scm		
D	"
E	1	N	1
(	3
Ù	6
 	C

4	U
Á	r
¯ 
± ¦	´ ¦
 «
; Ò
 	=
   C6 