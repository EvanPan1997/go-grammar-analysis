parser grammar UniversalParser;

options{ tokenVocab = UniversalLexer; }

@members{
    public boolean double_quoted_identifier_enable = false;
}

statements :
    statement (K_SEMICOLON statement)* K_SEMICOLON?;


statement
    : viewTableAndView
    | setStatement
    | insertQueryStatement
    | insertValueStatement
    | query
    | createTableStatement
    | dropTableStatement
    | renameTablePartitions
    | addTablePartitions
    | dropTablePartitions
    | createMaterializedViewStatement
    | updateMaterializedViewStatement
    | setMaterializedViewLifecycleStatement
    | updateMaterializedViewLifecycleStatement
    | dropMaterializedViewStatement
    | dropMaterializedViewPartitionStatement
    | createViewStatement
    | renameViewStatement
    | changeViewStatement
    | dropViewStatement
    ;

setStatement:
    T_SET multiIdentifier K_EQ  (identifier | number)
    ;

viewTableAndView
    : T_DESC T_EXTENDED? name=multiIdentifier partitionSpec?
    ;

createMaterializedViewStatement
    : T_CREATE T_MATERIALIZED T_VIEW ifOption? viewName=multiIdentifier
      lifecycleSpec?
      (T_BUILD T_DEFERRED)?
      viewColumns?
      (T_DISABLE T_REWRITE)?
      commentSpec?
      (T_PARTITIONED (T_ON|T_BY) K_OPEN_PAREN identifierSeq K_CLOSE_PAREN)?
      clusterBy? sortBy? bucketsOption?
      propertiesSpec?
      T_AS query
    ;

updateMaterializedViewStatement
    : T_ALTER T_MATERIALIZED T_VIEW viewName=multiIdentifier T_REBUILD partitionSpec
    ;

setMaterializedViewLifecycleStatement
    : T_ALTER T_MATERIALIZED T_VIEW viewName=multiIdentifier T_SET lifecycleSpec
    ;

updateMaterializedViewLifecycleStatement
    : T_ALTER T_MATERIALIZED T_VIEW viewName=multiIdentifier partitionSpec (T_ENABLE| T_DISABLE) T_LIFECYCLE
    ;

dropMaterializedViewStatement
    : T_DROP T_MATERIALIZED T_VIEW ifOption? tableName=multiIdentifier T_PURGE?
    ;

dropMaterializedViewPartitionStatement
    : T_ALTER T_MATERIALIZED T_VIEW viewName=multiIdentifier T_DROP ifOption? partitionSpec (K_COMMA partitionSpec)*
    ;

createViewStatement
    : T_CREATE replaceOption? T_VIEW ifOption? viewName=multiIdentifier viewColumns? commentSpec T_AS query;

viewColumns
    : K_OPEN_PAREN viewColumnCollection+=identifier commentSpec? (K_COMMA viewColumnCollection+=identifier commentSpec?)* K_CLOSE_PAREN
    ;

renameViewStatement
    : T_ALTER T_VIEW view_name=multiIdentifier T_RENAME T_TO new_view_name=multiIdentifier
    ;

changeViewStatement
    : T_ALTER T_VIEW view_name=multiIdentifier T_CHANGEOWNER T_TO new_owner=identifier
    ;

dropViewStatement:
    T_DROP T_VIEW ifOption? viewName=multiIdentifier;


createTableStatement
    : T_CREATE tempKeyWord=T_TEMPORARY? transKeyWord=T_TRANSACTIONAL? externalKeyWord=T_EXTERNAL?
    T_TABLE existsOption=ifOption table=multiIdentifier
    bodyOption=createTableBodyspec?
    commentOption=commentSpec?
    partitionOption=tablePartitionSpec?
    sortOption=sortSpec?
    skewedOption=skewedBySpec?
    formatOption=formatSpec?
    locationOption=locationSpec?
    propertiesOption=propertiesSpec?
    tableAsOption=tableSelectSpec?
    tableLifecycleOption=lifecycleSpec?
    ;


dropTableStatement:
    T_DROP T_TABLE ifOption? tableName=multiIdentifier T_PURGE?;

createTableBodyspec:
    K_OPEN_PAREN createTableItems K_CLOSE_PAREN;

createTableItems
    :columnItemCollection+=createTableItem (K_COMMA columnItemCollection+=createTableItem)*
    ;

createTableItem
    : columnName=identifier columnType=dataType nullOption? defaultOption? commentSpec?;

tablePartitionSpec
    : T_PARTITIONED T_BY K_OPEN_PAREN partitionItems K_CLOSE_PAREN;

partitionItems
    : partitionItem (K_COMMA partitionItem)*
    ;

partitionItem
    :columnName=identifier tableColumnType=dataType commentSpec?
    ;

sortSpec
    : clusterBy sortBy bucketsOption
    ;

skewedBySpec
    : T_SKEWED T_BY skewedByColumns T_ON skewedByValues storageDirectories?
    ;

skewedByColumns
    : expressionSeqs
    ;

skewedByValues
    : multiSkewedValue | simpleValue
    ;

multiSkewedValue:
    K_OPEN_PAREN simpleValue (K_COMMA simpleValue)* K_CLOSE_PAREN;

simpleValue
    : expressionSeqs
    ;

storageDirectories
    : T_STORED T_AS T_DIRECTORIES
    ;

formatSpec:
    tableRowFormat | storedBy;

tableRowFormat:
    rowFormat storedSpec?;

storedBy:
    T_STORED T_BY className=expression serdeProperties;

serdeProperties
    : T_WITH T_SERDEPROPERTIES propertyList;

storedSpec:
    T_STORED T_AS formatType;

formatType:
      T_SEQUENCEFILE
    | T_TEXTFILE
    | T_RCFILE
    | T_ORC
    | T_PARQUET
    | T_AVRO
    | T_JSONFILE
    | T_INPUTFORMAT inputFormatClassname=expression T_OUTPUTFORMAT outputFormatClassname=expression;

propertiesSpec:
    T_TBLPROPERTIES propertyList;

tableSelectSpec
    : T_AS query
    ;

lifecycleSpec
    : T_LIFECYCLE days=number
    ;


insertQueryStatement:
    cte? insertInto query
    ;

insertValueStatement:
    insertInto insertValues
    ;

dropTablePartitions
    : T_ALTER (T_TABLE | T_VIEW) multiIdentifier T_DROP ifOption partitionSpec (K_COMMA partitionSpec)* T_PURGE?
    ;

addTablePartitions
    : T_ALTER (T_TABLE | T_VIEW) multiIdentifier T_ADD ifOption partitionSpec (K_COMMA partitionSpec)*
    ;

renameTablePartitions:
    T_ALTER (T_TABLE | T_VIEW) multiIdentifier partitionSpec T_RENAME T_TO partitionSpec
    ;

insertInto
    : T_INSERT T_OVERWRITE T_TABLE? table=multiIdentifier partitionSpec? insertTableColumns?                    #insertOverwriteTable
    | T_INSERT T_INTO T_TABLE? table=multiIdentifier  partitionSpec? insertTableColumns?                        #insertIntoTable
    ;

insertValues
    : T_VALUES K_OPEN_PAREN defaultExpression (K_COMMA defaultExpression)* K_CLOSE_PAREN
    ;

insertTableColumns
    : K_OPEN_PAREN identifierSeq K_CLOSE_PAREN;


partitionSpec
    : T_PARTITION K_OPEN_PAREN partitionVal (K_COMMA partitionVal)* K_CLOSE_PAREN
    ;

partitionVal
    : identifier (compareOperator constant)? #standardPartitionVal
    | identifier K_EQ T_DEFAULT #defaultPartitionVal
    ;

ifOption:
    T_IF T_NOT? T_EXISTS;

replaceOption
    : T_OR T_REPLACE;

nullOption:
    T_NOT? T_NULL;

defaultOption:
    T_DEFAULT value=expression;

query
    : cte? queryTerm queryOrganization
    ;

cte
    : T_WITH namedQueryCollettion+=namedQuery (K_COMMA namedQueryCollettion+=namedQuery)*
    ;

namedQuery
    : name=identifier columnAliases? T_AS? K_OPEN_PAREN query K_CLOSE_PAREN
    ;

columnAliases
    : K_OPEN_PAREN identifierSeq K_CLOSE_PAREN
    ;

queryOrganization:
    orderBy? clusterBy? distributeBy? sortBy? windowClause? limitItem? offsetItem?;

queryTerm
    : queryPrimary                                                                                                                      #queryTermDefault
    | <assoc=left> left=queryTerm operator=(T_INTERSECT | T_UNION | T_EXCEPT | T_SETMINUS | T_MINUS) setQuantifier? right=queryTerm     #combineOperation
    ;

queryPrimary
    : querySpecification                                                                             #queryPrimaryDefault
    | K_OPEN_PAREN query K_CLOSE_PAREN                                                               #subquery
    ;

querySpecification
    : selectClause fromClause? whereClause? aggregationClause? havingClause? windowClause?
    ;

selectClause
    : T_SELECT (hints+=hint)* setQuantifier? namedExpressionSeq
    ;

fromClause
    : T_FROM relation (K_COMMA relation)* lateralView* pivotClause? unpivotClause?
    ;

inlineTable
    : T_VALUES expressionSeq aliasSpec?
    ;

functionTable
    : standardFunction aliasSpec?
    ;

temporalClause:
      T_FOR? (T_SYSTEM_VERSION | T_VERSION) T_AS T_OF version
    | T_FOR? (T_SYSTEM_TIME | T_TIMESTAMP) T_AS  T_OF timestamp=valueExpression
    ;

sample
    : T_TABLESAMPLE K_OPEN_PAREN sampleMethod? K_CLOSE_PAREN (T_REPEATABLE K_OPEN_PAREN seed=INTEGER_VALUE K_CLOSE_PAREN)?
    ;

sampleMethod
    : K_MINUS? percentage=(INTEGER_VALUE | DECIMAL_VALUE) T_PERCENTLIT                                                                #sampleByPercentile
    | expression T_ROWS                                                                                                               #sampleByRows
    | T_BUCKET numerator=INTEGER_VALUE T_OUT T_OF INTEGER_VALUE (T_ON (identifier | multiIdentifier K_OPEN_PAREN K_CLOSE_PAREN))?     #sampleByBucket
    | bytes=expression                                                                                                                #sampleByBytes
    ;

relationPrimary
    : table=multiIdentifier expressionSeqs? temporalClause? sample? aliasSpec?           #tableName
    | K_OPEN_PAREN query K_CLOSE_PAREN sample? aliasSpec?                                #aliasedQuery
    ;

relation
    : T_LATERAL? left=relationPrimary relationExtension*
    ;

relationExtension
    : joinRelation
    | pivotClause
    | unpivotClause
    ;

joinRelation
    : T_NATURAL? joinType T_JOIN T_LATERAL? right=relationPrimary joinCondtion?
    ;

joinCondtion
    : T_ON booleanExpression
    | T_USING  K_OPEN_PAREN identifierSeq K_CLOSE_PAREN
    ;

joinType
    : T_INNER?
    | T_CROSS
    | T_LEFT T_OUTER?
    | T_LEFT? T_SEMI
    | T_RIGHT T_OUTER?
    | T_FULL T_OUTER?
    | T_LEFT? T_ANTI
    ;

//列转行
pivotClause
    : T_PIVOT K_OPEN_PAREN
    aggregates+=pivotAggregateItem (K_COMMA aggregates+=pivotAggregateItem)*
    T_FOR pivotColumn T_IN K_OPEN_PAREN pivotValues+=pivotValue (K_COMMA pivotValues+=pivotValue)* K_CLOSE_PAREN K_CLOSE_PAREN
    ;

pivotAggregateItem:
    expression aliasSpec?;

pivotColumn
    : identifier
    | expressionSeqs
    ;

pivotValue:
      expression aliasSpec?
    | expressionSeqs;

//行转列
unpivotClause:
    T_UNPIVOT unpivotOption? K_OPEN_PAREN (unpivotSingleValueColumn | unpivotMultiValueColumn) K_CLOSE_PAREN unpivotColumnAlias?;

unpivotOption
    : opt=(T_INCLUDE | T_EXCLUDE) T_NULLS;

unpivotSingleValueColumn
    : unpivotVauleColumn=identifier T_FOR unpivotNameColumn T_IN K_OPEN_PAREN unpivotColumns+=unpivotColumnItem (K_COMMA unpivotColumns+=unpivotColumnItem)* K_CLOSE_PAREN;

unpivotMultiValueColumn:
    K_OPEN_PAREN identifierSeq  K_CLOSE_PAREN
    T_FOR unpivotNameColumn T_IN K_OPEN_PAREN unpivotColumnSets+=unpivotMultiColumnItem (K_COMMA unpivotColumnSets+=unpivotMultiColumnItem)* K_CLOSE_PAREN
    ;

unpivotNameColumn
    : identifier                                  #singleUnpivotNameColumn
    | K_OPEN_PAREN identifierSeq K_CLOSE_PAREN    #multiUnpivotNameColumn
    ;

unpivotColumnItem:
    unpivotColumn unpivotColumnAlias?;

unpivotMultiColumnItem:
    K_OPEN_PAREN multiIdentifierSeq K_CLOSE_PAREN unpivotColumnAlias?;

unpivotColumnAlias
    :T_AS? unpivotColumnAliasOption
    ;

unpivotColumnAliasOption
    : unpivotColumnAliasValueList
    | unpivotColumnAliasValue
    ;

unpivotColumnAliasValueList
    : K_OPEN_PAREN unpivotColumnAliasValues+=unpivotColumnAliasValue (K_COMMA unpivotColumnAliasValues+=unpivotColumnAliasValue)* K_CLOSE_PAREN
    ;

unpivotColumnAliasValue
    : string
    | identifier
    ;

unpivotColumn:
    multiIdentifier;


lateralView
    : T_LATERAL T_VIEW T_OUTER? functionName = multiIdentifier expressionSeqs tabName=identifier T_AS identifierSeq
    ;

whereClause
    : T_WHERE booleanExpression
    ;

havingClause
    : T_HAVING booleanExpression
    ;

aggregationClause
    : T_GROUP T_BY groupingExpressionsWithGroupingAnalytics+=groupByClause (K_COMMA groupingExpressionsWithGroupingAnalytics+=groupByClause)* #groupExpr
    | T_GROUP T_BY expressionSeq groupOption?                                                                                                 #groupCollection
    ;



groupOption
    : T_WITH kind=(T_ROLLUP | T_CUBE)                                                   #groupWith
    | T_GROUPING T_SETS K_OPEN_PAREN groupingSet (K_COMMA groupingSet)* K_CLOSE_PAREN   #groupWithSet
    ;

groupByClause
    : groupingAnalytics
    | expression
    ;

groupingElement
    : groupingAnalytics
    | groupingSet
    ;

groupingAnalytics
    : kind=(T_ROLLUP | T_CUBE) K_OPEN_PAREN groupingSet (K_COMMA groupingSet)* K_CLOSE_PAREN         #groupCube
    | T_GROUPING T_SETS K_OPEN_PAREN groupingElement (K_COMMA groupingElement)* K_CLOSE_PAREN        #groupSet
    ;

groupingSet
    : expressionSeqs
    | expression
    ;


hint
    : K_HINT_START hintStatements+=hintStatement (K_COMMA? hintStatements+=hintStatement)* K_HINT_END
    ;

hintStatement
    : hintName=identifier
    | hintName=identifier K_OPEN_PAREN parameters+=defaultExpression (K_COMMA parameters+=defaultExpression)* K_CLOSE_PAREN
    ;

namedExpressionSeq
    :namedExpressionCollection+=namedExpression (K_COMMA namedExpressionCollection+=namedExpression)*;

namedExpression
    : allQueryColumns             #queryStars
    | expression aliasSpec?       #queryColumn
    ;

allQueryColumns
    : (tableSchema = identifier K_DOT)? K_MULTIPLY;


expressionSeqs
    : K_OPEN_PAREN expressionSeq K_CLOSE_PAREN
    ;

expressionSeq
    : expressionCollection+=expression (K_COMMA expressionCollection+=expression)*
    ;

expression
    :  booleanExpression
    ;

booleanExpression
    : K_OPEN_PAREN booleanExpression K_CLOSE_PAREN                   #booleanPriority
    | reservedOpt=(T_NOT| K_EXCLAMATION) booleanExpression           #logicalNot
    | T_EXISTS subQueryExpress                                       #exists
    | left=booleanExpression operator=T_AND right=booleanExpression  #logicalBinary
    | left=booleanExpression operator=T_OR right=booleanExpression   #logicalBinary
    | valueExpression predicate                                      #predicated
    | valueExpression                                                #defaultBinary
    ;

predicate:
      T_NOT? kind=T_BETWEEN lower=valueExpression T_AND upper=valueExpression         #betweenPredicated
    | T_NOT? kind=T_IN expressionSeqs                                                 #inValuePredicated
    | T_NOT? kind=T_IN subQueryExpress                                                #inQueryPredicated
    | T_NOT? kind=T_RLIKE pattern=valueExpression           #rlikePredicated
    | T_NOT? kind=T_REGEXP pattern=valueExpression          #regexpPredicated
    | T_NOT? kind=(T_LIKE | T_ILIKE) quantifier=(T_ANY | T_SOME | T_ALL)  K_OPEN_PAREN expressionSeq? K_CLOSE_PAREN #likeQuantifierPredicated
    | T_NOT? kind=(T_LIKE | T_ILIKE) pattern=valueExpression (T_ESCAPE escape=string)? #likePatternPredicated
    | T_IS T_NOT? kind=T_NULL #nullPredicated
    | T_IS T_NOT? kind=(T_TRUE | T_FALSE | T_UNKNOWN) #booleanPredicated
    | T_IS T_NOT? kind=T_DISTINCT T_FROM right=valueExpression #distinctFromPredicated
    ;

valueExpression:
      K_OPEN_PAREN mid=valueExpression K_CLOSE_PAREN                                                    #valuePriority
    | operator=(K_ADD | K_MINUS | K_TILDE) valueExpression                                              #arithmeticUnaryReverse
    | left=valueExpression operator=(K_MULTIPLY | K_DIVIDE | K_PERCENT | T_DIV) right=valueExpression   #arithmeticBinary
    | left=valueExpression operator=(K_ADD | K_MINUS | K_CONCAT_PIPE) right=valueExpression             #arithmeticBinary
    | left=valueExpression operator=K_AMPERSAND right=valueExpression                                   #arithmeticBinary
    | left=valueExpression operator=K_HAT right=valueExpression                                         #arithmeticBinary
    | left=valueExpression operator=K_PIPE right=valueExpression                                        #arithmeticBinary
    | left=valueExpression compareOperator right=valueExpression                                        #compareBinary
    | defaultExpression                                                                                 #defaultUnary
    ;


defaultExpression
    : caseExpress                                                                     #caseExpression
    | function                                                                        #functionExpression
    | identifier                                                                      #columnReference
    | base=identifier K_DOT fieldName=identifier                                      #dereference
    | subQueryExpress                                                                 #subQueryExpression
    | constant                                                                        #fixValue
    | value=defaultExpression K_OPEN_BRACKET index=valueExpression K_CLOSE_BRACKET    #subScript
    ;


caseExpress
    : searchedCase
    | simpleCase
    ;

searchedCase
    : T_CASE whenClause+ elseClause? T_END
    ;

simpleCase
    : T_CASE value=expression whenClause+ elseClause? T_END
    ;

whenClause
    : T_WHEN condition=expression T_THEN result=expression
    ;

elseClause
    : T_ELSE expression;


function
    : specFunction
    | standardFunction
    ;

specFunction
    : castFunction
    | structFunction
    | firstFunction
    | anyValueFunction
    | lastFunction
    | trimFunction
    | extractFunction
    | substringFunction
    | overlayFunction
    | percentileFunction
    | transArrayFunction
    ;

castFunction
    : name=(T_CAST | T_TRY_CAST) K_OPEN_PAREN expression T_AS dataType K_CLOSE_PAREN
    ;

structFunction
    : T_STRUCT K_OPEN_PAREN (argument+=namedExpression (K_COMMA argument+=namedExpression)*)? K_CLOSE_PAREN
    ;

firstFunction
    : T_FIRST K_OPEN_PAREN expression (T_IGNORE T_NULLS)? K_CLOSE_PAREN
    ;

anyValueFunction
    : T_ANY_VALUE K_OPEN_PAREN expression (T_IGNORE T_NULLS)? K_CLOSE_PAREN
    ;

lastFunction
    : T_LAST K_OPEN_PAREN expression (T_IGNORE T_NULLS)? K_CLOSE_PAREN
    ;

subQueryExpress
    : K_OPEN_PAREN query K_CLOSE_PAREN
    ;


trimFunction
    :name=(T_TRIM | T_RTRIM | T_LTRIM) K_OPEN_PAREN option=(T_BOTH | T_LEADING | T_TRAILING)? trimStr=valueExpression? T_FROM srcStr=valueExpression K_CLOSE_PAREN
    ;

extractFunction:
    T_EXTRACT K_OPEN_PAREN field=identifier T_FROM source=valueExpression K_CLOSE_PAREN
    ;

substringFunction
    : name=(T_SUBSTR | T_SUBSTRING) K_OPEN_PAREN str=valueExpression subSplitOne=(T_FROM | K_COMMA) pos=valueExpression
      (subSpiltTwo=(T_FOR | K_COMMA) len=valueExpression)? K_CLOSE_PAREN;

overlayFunction
    : T_OVERLAY K_OPEN_PAREN input=valueExpression T_PLACING replace=valueExpression
      T_FROM position=valueExpression (T_FOR length=valueExpression)? K_CLOSE_PAREN;

percentileFunction:
    name=(T_PERCENTILE_CONT | T_PERCENTILE_DISC) K_OPEN_PAREN percentage=valueExpression K_CLOSE_PAREN
    T_WITHIN T_GROUP K_OPEN_PAREN orderBy K_CLOSE_PAREN filterSpec? overSpec?
    ;

transArrayFunction:
    T_TRANS_ARRAY K_OPEN_PAREN arguments  K_CLOSE_PAREN T_AS K_OPEN_PAREN identifierSeq  K_CLOSE_PAREN;

standardFunction:
    functionName=multiIdentifier K_OPEN_PAREN arguments? K_CLOSE_PAREN filterSpec?
    (nullsOption=(T_IGNORE | T_RESPECT) T_NULLS)? overSpec?
    ;

overSpec:
    T_OVER windowSpec;

filterSpec:
    T_FILTER K_OPEN_PAREN T_WHERE where=booleanExpression K_CLOSE_PAREN
    ;

constant
    : null                 #nullLiteral
    | interval             #intervalLiteral
    | number               #numberLiteral
    | identifier string    #typeLiteral
    | boolean              #booleaniteral
    | string               #stringLiteral;


compareOperator:
      K_EQ
    | K_EQ2
    | K_NEQ
    | K_NEQJ
    | K_LT
    | K_LTE
    | K_LTEJ
    | K_GT
    | K_GTE
    | K_GTEJ
    | K_NSEQ
    | K_ARROW
    ;

dataType:
      name=T_ARRAY K_LT dataType K_GT                                                                       #arrayDataType
    | name=T_MAP K_LT dataType K_COMMA dataType K_GT                                                        #mapDataType
    | name=T_STRUCT (K_LT complexColTypeList? K_GT | K_NEQ)                                                 #structDataType
    | T_INTERVAL from=(T_YEAR | T_MONTH) (T_TO to=T_MONTH)?                                                 #intervalDataType
    | T_INTERVAL from=(T_DAY | T_HOUR | T_MINUTE | T_SECOND) (T_TO to=(T_HOUR | T_MINUTE | T_SECOND))?      #intervalDataType
    | identifier (K_OPEN_PAREN INTEGER_VALUE (K_COMMA INTEGER_VALUE)* K_CLOSE_PAREN)?                       #primitiveDataType
    ;

complexColTypeList:
    complexColType (K_COMMA complexColType)*
    ;

complexColType:
    identifier K_COLON? dataType (T_NOT T_NULL)? commentSpec?
    ;

setQuantifier:
      T_DISTINCT
    | T_ALL
    ;

//日期基本元素定义开始
interval:
    T_INTERVAL (unitToUnitInterval | multiUnitsInterval)
    ;

multiUnitsInterval:
    (expression unit+=multiUnits)+
    ;

unitToUnitInterval:
    value=expression from=unitToUnit T_TO to=unitToUnit
    ;

multiUnits:
      T_NANOSECOND | T_NANOSECONDS | T_MICROSECOND | T_MICROSECONDS | T_MILLISECOND | T_MILLISECONDS
    | T_SECOND | T_SECONDS | T_MINUTE | T_MINUTES | T_HOUR | T_HOURS | T_DAY | T_DAYS | T_WEEK | T_WEEKS
    | T_MONTH | T_MONTHS | T_YEAR | T_YEARS
    ;

unitToUnit:
    T_SECOND | T_MINUTE | T_HOUR | T_DAY | T_MONTH | T_YEAR
    ;

commentSpec:
    T_COMMENT string
    ;

windowClause
    : T_WINDOW namedWindow (K_COMMA namedWindow)*
    ;

namedWindow
    : name=identifier T_AS windowSpec
    ;

windowSpec
    : name=identifier                                                                                          #windowRef
    | K_OPEN_PAREN name=identifier K_CLOSE_PAREN                                                               #windowRef
    | K_OPEN_PAREN ( clusterBy | (partitionBy| distributeBy)? (orderBy| sortBy)?) windowFrame? K_CLOSE_PAREN   #windowDef
    ;

windowFrame
    : frameType=(T_RANGE | T_ROWS) start=frameBound                                    #standardWindowFrame
    | frameType=(T_RANGE | T_ROWS) T_BETWEEN start=frameBound T_AND end=frameBound     #betweenWindowFrame
    ;

frameBound
    : T_UNBOUNDED boundType=(T_PRECEDING | T_FOLLOWING)   #unboundedPreceding
    | boundType=T_CURRENT T_ROW                           #currentRow
    | expression boundType=(T_PRECEDING | T_FOLLOWING)    #preceding
    ;

tableProvider
    : T_USING multiIdentifier
    ;

createFileFormat
    : T_STORED T_AS fileFormat
    | T_STORED T_BY storageHandler
    ;

fileFormat
    : T_INPUTFORMAT inFmt=string T_OUTPUTFORMAT outFmt=string    #tableFileFormat
    | identifier                                                 #genericFileFormat
    ;

storageHandler
    : string (T_WITH T_SERDEPROPERTIES propertyList)?
    ;

rowFormat
    : T_ROW T_FORMAT T_SERDE name=string (T_WITH T_SERDEPROPERTIES props=propertyList)?          #rowFormatSerde
    | T_ROW T_FORMAT T_DELIMITED
      (T_FIELDS T_TERMINATED T_BY fieldsTerminatedBy=string (T_ESCAPED T_BY escapedBy=string)?)?
      (T_COLLECTION T_ITEMS T_TERMINATED T_BY collectionItemsTerminatedBy=string)?
      (T_MAP T_KEYS T_TERMINATED T_BY keysTerminatedBy=string)?
      (T_LINES T_TERMINATED T_BY linesSeparatedBy=string)?
      (T_NULL T_DEFINED T_AS nullDefinedAs=string)?                                         #rowFormatDelimited
    ;

propertyList
    : K_OPEN_PAREN property (K_COMMA property)* K_CLOSE_PAREN
    ;

property
    : key=propertyKey (K_EQ? value=propertyValue)?
    ;

propertyKey
    : identifier (K_DOT identifier)*
    | string
    ;

propertyValue
    : INTEGER_VALUE
    | DECIMAL_VALUE
    | boolean
    | string
    ;

arguments
    : setQuantifier? argumentCollection+=expression (K_COMMA argumentCollection+=expression)*  #standardArgs
    | setQuantifier? K_MULTIPLY                                                                #stars
    ;

orderBy
    : T_ORDER T_BY sortItems
    ;

sortBy
    : sortKeyword=(T_SORT| T_SORTED) T_BY sortItems
    ;

distributeBy:
    T_DISTRIBUTE  T_BY expressionSeq
    ;

partitionBy:
    T_PARTITION  T_BY expressionSeq
    ;

clusterBy
    : T_RANGE? clusterKeyword=(T_CLUSTER | T_CLUSTERED) T_BY expressionSeq ordering=sortType?
    ;

bucketsOption
    : T_INTO num_buckets=number T_BUCKETS
    ;

limitItem
    : T_LIMIT T_ALL? expression
    ;

offsetItem
    : T_OFFSET offset=expression
    ;


sortItems
    : sortCollection+=sortItem (K_COMMA sortCollection+=sortItem)*
    ;

sortItem
    : expression ordering=sortType? (T_NULLS nullOrder=(T_LAST | T_FIRST))?
    ;

sortType
    : T_ASC | T_DESC
    ;

aliasSpec
    : T_AS? identifier
    ;

locationSpec
    : T_LOCATION string
    ;

multiIdentifierSeq
    :multiIdentifierList+=multiIdentifier (K_COMMA multiIdentifierList+=multiIdentifier)*;

multiIdentifier
    : identifier (K_DOT identifier)*
    ;

identifierSeq
    :identifierCollection+=identifier (K_COMMA identifierCollection+=identifier)*
    ;
identifier
    : IDENTIFIER              #unquotedIdentifierAlternative
    | quotedIdentifier        #quotedIdentifierAlternative
    | noReservedKeywords      #unquotedIdentifierAlternative
    ;

quotedIdentifier:
      BACKQUOTED_IDENTIFIER
    | {double_quoted_identifier_enable}? DOUBLE_QUOTED_STRING
    ;

number:
      K_MINUS? EXPONENT_VALUE                    #exponentLiteral
    | K_MINUS? DECIMAL_VALUE                     #decimalLiteral
    | K_MINUS? (EXPONENT_VALUE | DECIMAL_VALUE)  #legacyDecimalLiteral
    | K_MINUS? INTEGER_VALUE                     #integerLiteral
    | K_MINUS? BIGINT_VALUE                      #bigIntLiteral
    | K_MINUS? SMALLINT_VALUE                    #smallIntLiteral
    | K_MINUS? TINYINT_VALUE                     #tinyIntLiteral
    | K_MINUS? DOUBLE_VALUE                      #doubleLiteral
    | K_MINUS? FLOAT_VALUE                       #floatLiteral
    | K_MINUS? BIGDECIMAL_VALUE                  #bigDecimalLiteral
    ;

string:
      SINGLE_QUOTED_STRING
    | DOUBLE_QUOTED_STRING;

boolean:
      T_TRUE
    | T_FALSE;

null:
    T_NULL;

version:
      INTEGER_VALUE
    | string;

noReservedKeywords:
      T_ADD
    | T_AFTER
    | T_ALL
    | T_ALTER
    | T_ALWAYS
    | T_ANALYZE
    | T_ANY
    | T_ANY_VALUE
    | T_ARCHIVE
    | T_ARRAY
    | T_ASC
    | T_AT
    | T_AUTHORIZATION
    | T_BETWEEN
    | T_BOTH
    | T_BUCKET
    | T_BUCKETS
    | T_CACHE
    | T_CASCADE
    | T_CASE
    | T_CAST
    | T_CATALOG
    | T_CATALOGS
    | T_CHANGE
    | T_CHECK
    | T_CLEAR
    | T_CLUSTER
    | T_CLUSTERED
    | T_CODEGEN
    | T_COLLATE
    | T_COLLECTION
    | T_COLUMN
    | T_COLUMNS
    | T_COMMENT
    | T_COMMIT
    | T_COMPACT
    | T_COMPACTIONS
    | T_COMPUTE
    | T_CONCATENATE
    | T_CONSTRAINT
    | T_COST
    | T_CREATE
    | T_CUBE
    | T_CURRENT
    | T_CURRENT_DATE
    | T_CURRENT_TIME
    | T_CURRENT_TIMESTAMP
    | T_CURRENT_USER
    | T_DAY
    | T_DAYS
    | T_DAYOFYEAR
    | T_DATA
    | T_DATABASE
    | T_DATABASES
    | T_DATEADD
    | T_DATEDIFF
    | T_DBPROPERTIES
    | T_DEFAULT
    | T_DEFINED
    | T_DELETE
    | T_DELIMITED
    | T_DESC
    | T_DESCRIBE
    | T_DFS
    | T_DIRECTORIES
    | T_DIRECTORY
    | T_DISTINCT
    | T_DISTRIBUTE
    | T_DIV
    | T_DROP
    | T_ELSE
    | T_END
    | T_ESCAPE
    | T_ESCAPED
    | T_EXCEPT
    | T_EXCHANGE
    | T_EXCLUDE
    | T_EXISTS
    | T_EXPLAIN
    | T_EXPORT
    | T_EXTENDED
    | T_EXTERNAL
    | T_EXTRACT
    | T_FALSE
    | T_FETCH
    | T_FIELDS
    | T_FILTER
    | T_FILEFORMAT
    | T_FIRST
    | T_FOLLOWING
    | T_FOR
    | T_FOREIGN
    | T_FORMAT
    | T_FORMATTED
    | T_FULL
    | T_FUNCTION
    | T_FUNCTIONS
    | T_GENERATED
    | T_GLOBAL
    | T_GRANT
    | T_GROUP
    | T_GROUPING
    | T_HOUR
    | T_HOURS
    | T_IF
    | T_IGNORE
    | T_IMPORT
    | T_IN
    | T_INCLUDE
    | T_INDEX
    | T_INDEXES
    | T_INPATH
    | T_INPUTFORMAT
    | T_INTERSECT
    | T_INTERVAL
    | T_IS
    | T_ITEMS
    | T_KEYS
    | T_LAST
    | T_LATERAL
    | T_LAZY
    | T_LEADING
    | T_LIKE
    | T_ILIKE
    | T_LIMIT
    | T_LINES
    | T_LIST
    | T_LOAD
    | T_LOCAL
    | T_LOCATION
    | T_LOCK
    | T_LOCKS
    | T_LOGICAL
    | T_MACRO
    | T_MAP
    | T_MATCHED
    | T_MERGE
    | T_MICROSECOND
    | T_MICROSECONDS
    | T_MILLISECOND
    | T_MILLISECONDS
    | T_MINUTE
    | T_MINUTES
    | T_MONTH
    | T_MONTHS
    | T_MSCK
    | T_NAMESPACE
    | T_NAMESPACES
    | T_NANOSECOND
    | T_NANOSECONDS
    | T_NATURAL
    | T_NO
    | T_NULLS
    | T_OF
    | T_OFFSET
    | T_ONLY
    | T_OPTION
    | T_OPTIONS
    | T_OUT
    | T_OUTPUTFORMAT
    | T_OVERLAPS
    | T_OVERLAY
    | T_PARTITIONED
    | T_PARTITIONS
    | T_PERCENTILE_CONT
    | T_PERCENTILE_DISC
    | T_PERCENTLIT
    | T_PIVOT
    | T_PLACING
    | T_POSITION
    | T_PRECEDING
    | T_PRIMARY
    | T_PRINCIPALS
    | T_PROPERTIES
    | T_PURGE
    | T_QUARTER
    | T_QUERY
    | T_RANGE
    | T_RECORDREADER
    | T_RECORDWRITER
    | T_RECOVER
    | T_REDUCE
    | T_REFERENCES
    | T_REFRESH
    | T_RENAME
    | T_REPAIR
    | T_REPEATABLE
    | T_REPLACE
    | T_RESET
    | T_RESPECT
    | T_RESTRICT
    | T_REVOKE
    | T_RLIKE
    | T_ROLE
    | T_ROLES
    | T_ROLLBACK
    | T_ROLLUP
    | T_ROW
    | T_ROWS
    | T_SECOND
    | T_SECONDS
    | T_SCHEMA
    | T_SCHEMAS
    | T_SEPARATED
    | T_SERDE
    | T_SERDEPROPERTIES
    | T_SESSION_USER
    | T_SET
    | T_SETMINUS
    | T_SETS
    | T_SHOW
    | T_SKEWED
    | T_SOME
    | T_SORT
    | T_SORTED
    | T_SOURCE
    | T_START
    | T_STATISTICS
    | T_STORED
    | T_STRATIFY
    | T_STRUCT
    | T_SUBSTR
    | T_SUBSTRING
    | T_SYNC
    | T_SYSTEM_TIME
    | T_SYSTEM_VERSION
    | T_TABLE
    | T_TABLES
    | T_TABLESAMPLE
    | T_TARGET
    | T_TBLPROPERTIES
    | T_TEMPORARY
    | T_TERMINATED
    | T_THEN
    | T_TIME
    | T_TIMESTAMP
    | T_TIMESTAMPADD
    | T_TIMESTAMPDIFF
    | T_TO
    | T_TOUCH
    | T_TRAILING
    | T_TRANSACTION
    | T_TRANSACTIONS
    | T_TRANSFORM
    | T_TRIM
    | T_TRUE
    | T_TRUNCATE
    | T_TRY_CAST
    | T_TYPE
    | T_UNARCHIVE
    | T_UNBOUNDED
    | T_UNCACHE
    | T_UNION
    | T_UNIQUE
    | T_UNKNOWN
    | T_UNLOCK
    | T_UNPIVOT
    | T_UNSET
    | T_UPDATE
    | T_USE
    | T_USER
    | T_USING
    | T_VALUES
    | T_VERSION
    | T_VIEW
    | T_VIEWS
    | T_WEEK
    | T_WEEKS
    | T_WHEN
    | T_WINDOW
    | T_WITH
    | T_WITHIN
    | T_YEAR
    | T_YEARS
    | T_ZONE
    | T_ACTION
    | T_ADD2
    | T_ALLOCATE
    | T_ANSI_NULLS
    | T_ANSI_PADDING
    | T_ASSOCIATE
    | T_AUTO_INCREMENT
    | T_AVG
    | T_BATCHSIZE
    | T_BEGIN
    | T_BIGINT
    | T_BINARY_DOUBLE
    | T_BINARY_FLOAT
    | T_BINARY_INTEGER
    | T_BIT
    | T_BODY
    | T_BREAK
    | T_BULK
    | T_BYTE
    | T_CALL
    | T_CALLER
    | T_CASESPECIFIC
    | T_CHAR
    | T_CHARACTER
    | T_CHARSET
    | T_CLIENT
    | T_CLOSE
    | T_CMP
    | T_COLLECT
    | T_CONSTANT
    | T_COMPRESS
    | T_CONCAT
    | T_CONDITION
    | T_CONTINUE
    | T_COPY
    | T_COUNT
    | T_COUNT_BIG
    | T_CREATION
    | T_CREATOR
    | T_CS
    | T_CURRENT_SCHEMA
    | T_CURSOR
    | T_DATE
    | T_DATETIME
    | T_DEC
    | T_DECIMAL
    | T_DECLARE
    | T_DEFERRED
    | T_DEFINER
    | T_DEFINITION
    | T_DELIMITER
    | T_DIAGNOSTICS
    | T_DIR
    | T_DO
    | T_DOUBLE
    | T_DYNAMIC
    | T_ELSEIF
    | T_ELSIF
    | T_ENABLE
    | T_ENGINE
    | T_EXEC
    | T_EXECUTE
    | T_EXCEPTION
    | T_EXCLUSIVE
    | T_EXIT
    | T_FALLBACK
    | T_FILE
    | T_FILES
    | T_FLOAT
    | T_FOUND
    | T_GET
    | T_GO
    | T_HANDLER
    | T_HASH
    | T_HDFS
    | T_HIVE
    | T_HOST
    | T_IDENTITY
    | T_IMMEDIATE
    | T_INITRANS
    | T_INOUT
    | T_INT
    | T_INT2
    | T_INT4
    | T_INT8
    | T_INTEGER
    | T_INVOKER
    | T_ISOPEN
    | T_KEEP
    | T_KEY
    | T_LANGUAGE
    | T_LEAVE
    | T_LOCATOR
    | T_LOCATORS
    | T_LOG
    | T_LOGGED
    | T_LOGGING
    | T_LOOP
    | T_MAX
    | T_MAXTRANS
    | T_MESSAGE_TEXT
    | T_MIN
    | T_MULTISET
    | T_NCHAR
    | T_NEW
    | T_NVARCHAR
    | T_NOCOUNT
    | T_NOCOMPRESS
    | T_NOLOGGING
    | T_NONE
    | T_NOTFOUND
    | T_NUMERIC
    | T_NUMBER
    | T_OBJECT
    | T_OFF
    | T_OPEN
    | T_OWNER
    | T_PACKAGE
    | T_PCTFREE
    | T_PCTUSED
    | T_PLS_INTEGER
    | T_PRECISION
    | T_PRESERVE
    | T_PRINT
    | T_PROC
    | T_PROCEDURE
    | T_QUALIFY
    | T_QUERY_BAND
    | T_QUIT
    | T_QUOTED_IDENTIFIER
    | T_RAISE
    | T_REAL
    | T_REGEXP
    | T_RESIGNAL
    | T_RESULT
    | T_RESULT_SET_LOCATOR
    | T_RETURN
    | T_RETURNS
    | T_REVERSE
    | T_ROWTYPE
    | T_ROW_COUNT
    | T_RR
    | T_RS
    | T_PWD
    | T_SECURITY
    | T_SEGMENT
    | T_SEL
    | T_SESSION
    | T_SESSIONS
    | T_SHARE
    | T_SIGNAL
    | T_SIMPLE_DOUBLE
    | T_SIMPLE_FLOAT
    | T_SIMPLE_INTEGER
    | T_SMALLDATETIME
    | T_SMALLINT
    | T_SQL
    | T_SQLEXCEPTION
    | T_SQLINSERT
    | T_SQLSTATE
    | T_SQLWARNING
    | T_STATS
    | T_STEP
    | T_STORAGE
    | T_STRING
    | T_SUBDIR
    | T_SUM
    | T_SUMMARY
    | T_SYS_REFCURSOR
    | T_TABLESPACE
    | T_TEXTIMAGE_ON
    | T_TINYINT
    | T_TITLE
    | T_TOP
    | T_UR
    | T_VALUE
    | T_VAR
    | T_VARCHAR
    | T_VARCHAR2
    | T_VARYING
    | T_VOLATILE
    | T_WHILE
    | T_WITHOUT
    | T_WORK
    | T_XACT_ABORT
    | T_XML
    | T_YES
    | T_ACTIVITY_COUNT
    | T_CUME_DIST
    | T_DENSE_RANK
    | T_FIRST_VALUE
    | T_LAG
    | T_LAST_VALUE
    | T_LEAD
    | T_MAX_PART_STRING
    | T_MIN_PART_STRING
    | T_MAX_PART_INT
    | T_MIN_PART_INT
    | T_MAX_PART_DATE
    | T_MIN_PART_DATE
    | T_PART_COUNT
    | T_PART_LOC
    | T_RANK
    | T_ROW_NUMBER
    | T_STDEV
    | T_SYSDATE
    | T_VARIANCE
    | T_ABS
    | T_ACOS
    | T_ADD_MONTHS
    | T_ALL_MATCH
    | T_ANY_MATCH
    | T_ATAN2
    | T_APPROX_DISTINCT
    | T_ARG_MAX
    | T_ARG_MIN
    | T_ARRAY_CONTAINS
    | T_ARRAY_DISTINCT
    | T_ARRAY_EXCEPT
    | T_ARRAY_INTERSECT
    | T_ARRAY_JOIN
    | T_ARRAY_MAX
    | T_ARRAY_MIN
    | T_ARRAY_NORMALIZE
    | T_ARRAY_POSITION
    | T_ARRAY_REDUCE
    | T_ARRAY_REMOVE
    | T_ARRAY_REPEAT
    | T_ARRAY_SORT
    | T_ARRAY_UNION
    | T_ARRAYS_OVERLAP
    | T_ARRAYS_ZIP
    | T_ASCII
    | T_ASIN
    | T_ATAN
    | T_BASE64
    | T_BIN
    | T_BITWISE_AND_AGG
    | T_BITWISE_OR_AGG
    | T_CBRT
    | T_CEIL
    | T_CHAR_MATCHCOUNT
    | T_CHR
    | T_CLUSTER_SAMPLE
    | T_COALESCE
    | T_COLLECT_LIST
    | T_COLLECT_SET
    | T_COMBINATIONS
    | T_CONCAT_WS
    | T_CONV
    | T_CORR
    | T_COS
    | T_COSH
    | T_COT
    | T_COUNT_IF
    | T_COVAR_POP
    | T_COVAR_SAMP
    | T_CRC32
    | T_CURRENT_TIMEZONE
    | T_DATE_ADD
    | T_DATE_FORMAT
    | T_DATE_SUB
    | T_DATEPART
    | T_DATETRUNC
    | T_DAYOFMONTH
    | T_DAYOFWEEK
    | T_DECODE
    | T_DECOMPRESS
    | T_DEGREES
    | T_E
    | T_ENCODE
    | T_EXP
    | T_EXPLODE
    | T_FACTORIAL
    | T_FIELD
    | T_FIND_IN_SET
    | T_FLATTEN
    | T_FLOOR
    | T_FORMAT_NUMBER
    | T_FROM_JSON
    | T_FROM_UNIXTIME
    | T_FROM_UTC_TIMESTAMP
    | T_GET_IDCARD_AGE
    | T_GET_IDCARD_BIRTHDAY
    | T_GET_IDCARD_SEX
    | T_GET_JSON_OBJECT
    | T_GET_USER_ID
    | T_GETDATE
    | T_GREATEST
    | T_HEX
    | T_HISTOGRAM
    | T_INLINE
    | T_INITCAP
    | T_INSTR
    | T_IS_ENCODING
    | T_ISDATE
    | T_ISNAN
    | T_JSON_OBJECT
    | T_JSON_ARRAY
    | T_JSON_EXTRACT
    | T_JSON_EXISTS
    | T_JSON_PRETTY
    | T_JSON_TYPE
    | T_JSON_FORMAT
    | T_JSON_PARSE
    | T_JSON_VALID
    | T_JSON_TUPLE
    | T_KEYVALUE
    | T_KEYVALUE_TUPLE
    | T_LAST_DAY
    | T_LASTDAY
    | T_LEAST
    | T_LENGTH
    | T_LENGTHB
    | T_LN
    | T_LOCATE
    | T_LOG10
    | T_LOG2
    | T_LPAD
    | T_LTRIM
    | T_MAP_AGG
    | T_MAP_CONCAT
    | T_MAP_ENTRIES
    | T_MAP_FILTER
    | T_MAP_FROM_ARRAYS
    | T_MAP_FROM_ENTRIES
    | T_MAP_KEYS
    | T_MAP_UNION
    | T_MAP_UNION_SUM
    | T_MAP_VALUES
    | T_MAP_ZIP_WITH
    | T_MASK_HASH
    | T_MAX_BY
    | T_MAX_PT
    | T_MD5
    | T_MEDIAN
    | T_MIN_BY
    | T_MONTHS_BETWEEN
    | T_MULTIMAP_AGG
    | T_MULTIMAP_FROM_ENTRIES
    | T_NAMED_STRUCT
    | T_NEGATIVE
    | T_NEXT_DAY
    | T_NGRAMS
    | T_NOW
    | T_NTILE
    | T_NTH_VALUE
    | T_NULLIF
    | T_NUMERIC_HISTOGRAM
    | T_NVL
    | T_ORDINAL
    | T_PARSE_URL
    | T_PARSE_URL_TUPLE
    | T_PARTITION_EXISTS
    | T_PERCENT_RANK
    | T_PERCENTILE
    | T_PERCENTILE_APPROX
    | T_PI
    | T_POSEXPLODE
    | T_POSITIVE
    | T_POW
    | T_RADIANS
    | T_RAND
    | T_REGEXP_COUNT
    | T_REGEXP_EXTRACT
    | T_REGEXP_EXTRACT_ALL
    | T_REGEXP_INSTR
    | T_REGEXP_REPLACE
    | T_REGEXP_SUBSTR
    | T_REPEAT
    | T_ROUND
    | T_RPAD
    | T_RTRIM
    | T_SAMPLE
    | T_SEQUENCE
    | T_SHA
    | T_SHA1
    | T_SHA2
    | T_SHIFTLEFT
    | T_SHIFTRIGHT
    | T_SHIFTRIGHTUNSIGNED
    | T_SHUFFLE
    | T_SIGN
    | T_SIN
    | T_SINH
    | T_SIZE
    | T_SLICE
    | T_SORT_ARRAY
    | T_SOUNDEX
    | T_SPACE
    | T_SPLIT
    | T_SPLIT_PART
    | T_SQRT
    | T_STACK
    | T_STDDEV
    | T_STDDEV_SAMP
    | T_STR_TO_MAP
    | T_SUBSTRING_INDEX
    | T_SYM_DECRYPT
    | T_SYM_ENCRYPT
    | T_TABLE_EXISTS
    | T_TAN
    | T_TANH
    | T_TO_CHAR
    | T_TO_DATE
    | T_TO_JSON
    | T_TO_MILLIS
    | T_TOLOWER
    | T_TOUPPER
    | T_TRANS_ARRAY
    | T_TRANS_COLS
    | T_TRANSFORM_KEYS
    | T_TRANSFORM_VALUES
    | T_TRANSLATE
    | T_TRUNC
    | T_UNBASE64
    | T_UNHEX
    | T_UNIQUE_ID
    | T_UNIX_TIMESTAMP
    | T_URL_DECODE
    | T_URL_ENCODE
    | T_UUID
    | T_VAR_SAMP
    | T_VAR_POP
    | T_WEEKDAY
    | T_WEEKOFYEAR
    | T_WIDTH_BUCKET
    | T_WM_CONCAT
    | T_ZIP_WITH;
